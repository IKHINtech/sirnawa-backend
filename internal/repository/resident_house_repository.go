package repository

import (
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type ResidentHouseRepository interface {
	Create(tx *gorm.DB, data models.ResidentHouse) (*models.ResidentHouse, error)
	Update(tx *gorm.DB, id string, data models.ResidentHouse) (*models.ResidentHouse, error)
	ChangeToPrimary(tx *gorm.DB, id string) error
	ChangeNotPrimaryByResidentID(tx *gorm.DB, residentID string) error
	FindAll() (models.ResidentHouses, error)
	Paginated(pagination utils.Pagination) (*utils.Pagination, models.ResidentHouses, error)
	FindByID(id string) (*models.ResidentHouse, error)
	FindByResidentID(residentID string) (models.ResidentHouses, error)
	Delete(tx *gorm.DB, id string) error
}

type residentHouseRepositoryImpl struct {
	db *gorm.DB
}

func NewResidentHouseRepository(db *gorm.DB) ResidentHouseRepository {
	return &residentHouseRepositoryImpl{db: db}
}

func (r *residentHouseRepositoryImpl) Paginated(pagination utils.Pagination) (*utils.Pagination, models.ResidentHouses, error) {
	var datas models.ResidentHouses
	query := r.db
	err := query.Scopes(utils.Paginate(datas, &pagination, query)).Find(&datas).Error
	return &pagination, datas, err
}

func (r *residentHouseRepositoryImpl) Create(tx *gorm.DB, data models.ResidentHouse) (*models.ResidentHouse, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Create(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *residentHouseRepositoryImpl) ChangeNotPrimaryByResidentID(tx *gorm.DB, residentID string) error {
	if tx == nil {
		tx = r.db
	}
	err := tx.Model(&models.ResidentHouse{}).Select("is_primary").Where("resident_id = ?", residentID).Update("is_primary", false).Error
	return err
}

func (r *residentHouseRepositoryImpl) ChangeToPrimary(tx *gorm.DB, id string) error {
	if tx == nil {
		tx = r.db
	}
	err := tx.Model(&models.ResidentHouse{}).Select("is_primary").Where("id = ?", id).Update("is_primary", true).Error
	return err
}

func (r *residentHouseRepositoryImpl) Update(tx *gorm.DB, id string, data models.ResidentHouse) (*models.ResidentHouse, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Model(&models.ResidentHouse{}).Where("id = ?", id).Updates(data).Error
	return &data, err
}

func (r *residentHouseRepositoryImpl) FindByResidentID(residentID string) (models.ResidentHouses, error) {
	var data models.ResidentHouses

	err := r.db.First(&data, "resident_id = ?", residentID).Error
	if err != nil {
		return nil, err
	}
	return data, err
}

func (r *residentHouseRepositoryImpl) FindByID(id string) (*models.ResidentHouse, error) {
	var data models.ResidentHouse

	err := r.db.First(&data, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *residentHouseRepositoryImpl) FindAll() (models.ResidentHouses, error) {
	var data models.ResidentHouses
	err := r.db.Find(&data).Error
	return data, err
}

func (r *residentHouseRepositoryImpl) Delete(tx *gorm.DB, id string) error {
	if tx == nil {
		tx = r.db
	}
	err := tx.Delete(&models.ResidentHouse{}, "id = ?", id).Error
	return err
}
