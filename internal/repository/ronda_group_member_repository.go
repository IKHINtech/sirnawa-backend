package repository

import (
	"errors"

	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type RondaGroupMemberRepository interface {
	Create(tx *gorm.DB, data models.RondaGroupMember) (*models.RondaGroupMember, error)
	Update(tx *gorm.DB, id string, data models.RondaGroupMember) (*models.RondaGroupMember, error)
	FindAll() (models.RondaGroupMembers, error)
	Paginated(pagination utils.Pagination) (*utils.Pagination, models.RondaGroupMembers, error)
	FindByID(id string) (*models.RondaGroupMember, error)
	FindByHouseID(houseID string) (*models.RondaGroupMember, error)
	GetTotalMember(groupID string) (*int64, error)
	Delete(id string) error
}

type rondaGroupMemberRepositoryImpl struct {
	db *gorm.DB
}

func NewRondaGroupMemberRepository(db *gorm.DB) RondaGroupMemberRepository {
	return &rondaGroupMemberRepositoryImpl{db: db}
}

func (r *rondaGroupMemberRepositoryImpl) GetTotalMember(groupID string) (*int64, error) {
	var data int64
	err := r.db.Model(&models.RondaGroupMember{}).Where("group_id = ?", groupID).Count(&data).Error
	return &data, err
}

func (r *rondaGroupMemberRepositoryImpl) Paginated(pagination utils.Pagination) (*utils.Pagination, models.RondaGroupMembers, error) {
	var datas models.RondaGroupMembers
	query := r.db
	err := query.Scopes(utils.Paginate(datas, &pagination, query)).Find(&datas).Error
	return &pagination, datas, err
}

func (r *rondaGroupMemberRepositoryImpl) Create(tx *gorm.DB, data models.RondaGroupMember) (*models.RondaGroupMember, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Create(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *rondaGroupMemberRepositoryImpl) Update(tx *gorm.DB, id string, data models.RondaGroupMember) (*models.RondaGroupMember, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Model(&models.RondaGroupMember{}).Where("id = ?", id).Updates(data).Error
	return &data, err
}

func (r *rondaGroupMemberRepositoryImpl) FindByHouseID(houseID string) (*models.RondaGroupMember, error) {
	var data models.RondaGroupMember

	err := r.db.First(&data, "house_id = ?", houseID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &data, err
}

func (r *rondaGroupMemberRepositoryImpl) FindByID(id string) (*models.RondaGroupMember, error) {
	var data models.RondaGroupMember

	err := r.db.First(&data, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *rondaGroupMemberRepositoryImpl) FindAll() (models.RondaGroupMembers, error) {
	var data models.RondaGroupMembers
	err := r.db.Find(&data).Error
	return data, err
}

func (r *rondaGroupMemberRepositoryImpl) Delete(id string) error {
	err := r.db.Delete(&models.RondaGroupMember{}, "id = ?", id).Error
	return err
}
