package services

import (
	"errors"

	"github.com/IKHINtech/sirnawa-backend/internal/config"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/response"
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/internal/repository"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type ResidentService interface {
	Create(data request.ResidentCreateRequest) (*response.ResidentResponse, error)
	Update(id string, data request.ResidentUpdateRequset) (*response.ResidentResponse, error)
	FindByID(id string) (*response.ResidentResponse, error)
	Delete(id string) error
	FindAll(rt_id, search string) (response.ResidentResponses, error)
	Paginated(pagination utils.Pagination, rt_id, search string) (*utils.Pagination, *response.ResidentResponses, error)
}

type residentServiceImpl struct {
	repository      repository.ResidentRepository
	userRepostitory repository.UserRepository
	userRt          repository.UserRTRepository
	db              *gorm.DB
}

func NewResidentServices(
	repo repository.ResidentRepository,
	userRepostitory repository.UserRepository,
	userRt repository.UserRTRepository,
	db *gorm.DB,
) ResidentService {
	return &residentServiceImpl{
		repository:      repo,
		userRepostitory: userRepostitory,
		userRt:          userRt,
		db:              db,
	}
}

func (s *residentServiceImpl) withTransaction(fn func(tx *gorm.DB) error) error {
	tx := s.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (s *residentServiceImpl) CheckUserRT(tx *gorm.DB, data request.ResidentCreateRequest, resident models.Resident, user models.User) error {
	// cek user rt
	listUserRT, err := s.userRt.FindByUserID(user.ID)
	if err != nil {
		return err
	}

	// jika user tidak memiliki rt => langsung buatkan
	if len(listUserRT) == 0 {
		userRT := models.UserRT{
			UserID: user.ID,
			Role:   models.Role(data.Role),
			RtID:   data.RtID,
		}
		_, err = s.userRt.Create(tx, userRT)
		if err != nil {
			return err
		}
	} else {
		// ketik user telah memiliki rt cek apakah rt pada payload sudah terdafar
		existRt := false
		var exisitingUserRt *models.UserRT
		for _, rt := range listUserRT {
			if rt.RtID == data.RtID {
				existRt = true
				exisitingUserRt = &rt
				break
			}
		}
		if !existRt {
			userRT := models.UserRT{
				UserID: user.ID,
				Role:   models.Role(data.Role),
				RtID:   data.RtID,
			}
			_, err = s.userRt.Create(tx, userRT)
			if err != nil {
				return err
			}
		} else {
			// jika user sudah terdaftar pada suatu rt
			if models.Role(data.Role) != models.RoleWarga {
				if exisitingUserRt.Role == models.RoleWarga {
					// naik level jadi bukan warga
					userRT := models.UserRT{
						UserID: user.ID,
						Role:   models.Role(data.Role),
						RtID:   data.RtID,
					}
					userRT.ID = exisitingUserRt.ID
					_, err = s.userRt.Update(tx, exisitingUserRt.ID, userRT)
					if err != nil {
						return err
					}
				}
			}
		}
	}
	return err
}

func (s *residentServiceImpl) Create(data request.ResidentCreateRequest) (*response.ResidentResponse, error) {
	var result *models.Resident

	err := s.withTransaction(func(tx *gorm.DB) error {
		// tampungan untuk user
		var existingUser *models.User
		// cari dulu apakah nik sudah terdaftar
		existingNik, err := s.repository.FindByNIK(data.NIK)
		if err != nil {
			return err
		}

		// apakah payload email != nil
		if data.Email != nil {
			// carikan email di database
			existingUser, err = s.userRepostitory.FindByEmail(*data.Email)
			if err != nil {
				return err
			}
		}

		// jika email dan nik di temukan
		if existingUser != nil && existingNik != nil {
			if *existingUser.ResidentID != existingNik.ID {
				return errors.New("user dengan email tersebut terdaftar dengan NIK yang berbeda")
			} else {
				// cek user role
				err = s.CheckUserRT(tx, data, *existingNik, *existingUser)
				if err != nil {
					return err
				}
			}
		}

		// jika nik ditemukan dan email tidak
		if existingUser == nil && existingNik != nil {
			// cari user dengan nik yang dikirim
			userByResidentID, err := s.userRepostitory.FindByResidentID(existingNik.ID)
			if err != nil {
				return err
			}

			if userByResidentID == nil {
				return errors.New("warga tidak memiliki user")
			}
			// jika email != nil maka update user dengan email yang dikirimkan
			if data.Email != nil {
				// update email pada user yang ditemuakn dengan email yang baru dimasukkan
				err = s.userRepostitory.UpdateEmail(tx, userByResidentID.ID, *data.Email)
				if err != nil {
				}
				return err
			}
			// cek user role
			err = s.CheckUserRT(tx, data, *existingNik, *userByResidentID)
			if err != nil {
				return err
			}
		}

		// ketika nik dan email tidak terdaftar (resident baru)
		if existingNik == nil && existingUser == nil {
			// buat resident
			payload := request.ResidentCreateRequestToResidentModel(data)
			created, err := s.repository.Create(tx, payload)
			if err != nil {
				return err
			}
			existingNik = created
			// buat user

			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(config.AppConfig.DEFAULT_PASSWORD), bcrypt.DefaultCost)
			if err != nil {
				return err
			}

			// ada data email pada resident request, buatkan user
			payloadUser := models.User{
				Email:      data.Email,
				Password:   string(hashedPassword),
				ResidentID: &existingNik.ID,
			}

			newUser, err := s.userRepostitory.Create(tx, payloadUser)
			if err != nil {
				return err
			}

			userRT := models.UserRT{
				UserID: newUser.ID,
				Role:   models.Role(data.Role),
				RtID:   data.RtID,
			}
			_, err = s.userRt.Create(tx, userRT)
			if err != nil {
				return err
			}
		}

		result = existingNik
		return nil
	})
	// handling err
	if err != nil {
		return nil, err
	}

	res := response.ResidentModelToResidentResponse(result)
	return res, nil
}

func (s *residentServiceImpl) Update(id string, data request.ResidentUpdateRequset) (*response.ResidentResponse, error) {
	var result *models.Resident

	err := s.withTransaction(func(tx *gorm.DB) error {
		existing, err := s.repository.FindByID(id)
		if err != nil {
			return err
		}

		existingNik, err := s.repository.FindByNIK(data.NIK)
		if err != nil {
			return err
		}

		if existingNik != nil && existingNik.ID != existing.ID {
			return errors.New("NIK tersebut sudah terdaftar")
		}

		payload := request.ResidentUpdateRequsetToResidentModel(data)
		payload.ID = existing.ID

		updated, err := s.repository.Update(tx, payload.ID, payload)
		if err != nil {
			return err
		}

		result = updated
		return nil
	})
	// handle error
	if err != nil {
		return nil, err
	}

	res := response.ResidentModelToResidentResponse(result)
	return res, nil
}

func (s *residentServiceImpl) FindAll(rt_id, search string) (response.ResidentResponses, error) {
	result, err := s.repository.FindAll(rt_id, search)
	if err != nil {
		return nil, err
	}

	resp := response.ResidentListToResponse(result)
	return resp, nil
}

func (s *residentServiceImpl) FindByID(id string) (*response.ResidentResponse, error) {
	result, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	resp := response.ResidentModelToResidentResponse(result)
	return resp, err
}

func (s *residentServiceImpl) Paginated(pagination utils.Pagination, rt_id, search string) (*utils.Pagination, *response.ResidentResponses, error) {
	paginated, data, err := s.repository.Paginated(pagination, rt_id, search)
	if err != nil {
		return nil, nil, err
	}

	resp := response.ResidentListToResponse(data)
	return paginated, &resp, err
}

func (s *residentServiceImpl) Delete(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}
