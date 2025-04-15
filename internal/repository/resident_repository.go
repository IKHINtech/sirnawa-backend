package repository

import (
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type ResidentRepository interface {
	Create(tx *gorm.DB, data models.Resident) (*models.Resident, error)
	Update(tx *gorm.DB, id string, data models.Resident) (*models.Resident, error)
	FindAll() (models.Residents, error)
	Paginated(pagination utils.Pagination) (*utils.Pagination, models.Residents, error)
	FindByID(id string) (*models.Resident, error)
	Delete(id string) error
}

type residentRepositoryImpl struct {
	db *gorm.DB
}

func NewResidentRepository(db *gorm.DB) ResidentRepository {
	return &residentRepositoryImpl{db: db}
}

func (r *residentRepositoryImpl) Paginated(pagination utils.Pagination) (*utils.Pagination, models.Residents, error) {
	var datas models.Residents
	query := r.db
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

func (r *residentRepositoryImpl) FindAll() (models.Residents, error) {
	var data models.Residents
	err := r.db.Find(&data).Error
	return data, err
}

func (r *residentRepositoryImpl) Delete(id string) error {
	err := r.db.Delete(&models.Resident{}, "id = ?", id).Error
	return err
}
