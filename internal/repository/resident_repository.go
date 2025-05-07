package repository

import (
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type ResidentRepository interface {
	Create(tx *gorm.DB, data models.Resident) (*models.Resident, error)
	Update(tx *gorm.DB, id string, data models.Resident) (*models.Resident, error)
	FindAll(rt_id, search string) (models.Residents, error)
	Paginated(pagination utils.Pagination, rt_id, search string) (*utils.Pagination, models.Residents, error)
	FindByID(id string) (*models.Resident, error)
	FindByNIK(nik string) (*models.Resident, error)
	Delete(id string) error
}

type residentRepositoryImpl struct {
	db *gorm.DB
}

func NewResidentRepository(db *gorm.DB) ResidentRepository {
	return &residentRepositoryImpl{db: db}
}

func (r *residentRepositoryImpl) Paginated(pagination utils.Pagination, rt_id, search string) (*utils.Pagination, models.Residents, error) {
	var datas models.Residents

	query := r.db.Table("residents").Select("residents.*")

	if rt_id != "" {
		query = query.
			Joins("JOIN users ON users.resident_id = residents.id").
			Joins("JOIN user_rts ON user_rts.user_id = users.id").
			Where("user_rts.rt_id = ?", rt_id)
	}

	if search != "" {
		query = query.Where("LOWER(name) LIKE LOWER(?) OR nik LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	err := query.Scopes(utils.Paginate(datas, &pagination, query)).Find(&datas).Error
	return &pagination, datas, err
}

func (r *residentRepositoryImpl) Create(tx *gorm.DB, data models.Resident) (*models.Resident, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Create(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *residentRepositoryImpl) Update(tx *gorm.DB, id string, data models.Resident) (*models.Resident, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Model(&models.Resident{}).Where("id = ?", id).Updates(data).Error
	return &data, err
}

func (r *residentRepositoryImpl) FindByID(id string) (*models.Resident, error) {
	var data models.Resident

	err := r.db.First(&data, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *residentRepositoryImpl) FindByNIK(nik string) (*models.Resident, error) {
	var data models.Resident

	err := r.db.First(&data, "nik = ?", nik).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &data, err
}

func (r *residentRepositoryImpl) FindAll(rt_id, search string) (models.Residents, error) {
	var data models.Residents

	query := r.db.Table("residents")
	if rt_id != "" {
		query = query.Joins("users on users.resident_id = residents.id ")
		query = query.Joins("user_rts on user_rts.user_id = users.id").Where("user_rts.rt_id = ?", rt_id)
	}

	if search != "" {
		query = query.Where("LOWER(name) like  LOWER(?)  ", "%"+search+"%")
	}

	err := query.Find(&data).Error
	return data, err
}

func (r *residentRepositoryImpl) Delete(id string) error {
	err := r.db.Delete(&models.Resident{}, "id = ?", id).Error
	return err
}
