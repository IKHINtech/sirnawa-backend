package repository

import (
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type RondaActivityRepository interface {
	Create(tx *gorm.DB, data models.RondaActivity) (*models.RondaActivity, error)
	Update(tx *gorm.DB, id string, data models.RondaActivity) (*models.RondaActivity, error)
	FindAll() (models.RondaActivitys, error)
	Paginated(pagination utils.Pagination) (*utils.Pagination, models.RondaActivitys, error)
	FindByID(id string) (*models.RondaActivity, error)
	Delete(id string) error
}

type rondaActivityRepositoryImpl struct {
	db *gorm.DB
}

func NewRondaActivityRepository(db *gorm.DB) RondaActivityRepository {
	return &rondaActivityRepositoryImpl{db: db}
}

func (r *rondaActivityRepositoryImpl) Paginated(pagination utils.Pagination) (*utils.Pagination, models.RondaActivitys, error) {
	var datas models.RondaActivitys
	query := r.db
	err := query.Scopes(utils.Paginate(datas, &pagination, query)).Find(&datas).Error
	return &pagination, datas, err
}

func (r *rondaActivityRepositoryImpl) Create(tx *gorm.DB, data models.RondaActivity) (*models.RondaActivity, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Create(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *rondaActivityRepositoryImpl) Update(tx *gorm.DB, id string, data models.RondaActivity) (*models.RondaActivity, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Model(&models.RondaActivity{}).Where("id = ?", id).Updates(data).Error
	return &data, err
}

func (r *rondaActivityRepositoryImpl) FindByID(id string) (*models.RondaActivity, error) {
	var data models.RondaActivity

	err := r.db.First(&data, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *rondaActivityRepositoryImpl) FindAll() (models.RondaActivitys, error) {
	var data models.RondaActivitys
	err := r.db.Find(&data).Error
	return data, err
}

func (r *rondaActivityRepositoryImpl) Delete(id string) error {
	err := r.db.Delete(&models.RondaActivity{}, "id = ?", id).Error
	return err
}
