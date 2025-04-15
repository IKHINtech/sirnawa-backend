package repository

import (
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type RondaAttendanceRepository interface {
	Create(tx *gorm.DB, data models.RondaAttendance) (*models.RondaAttendance, error)
	Update(tx *gorm.DB, id string, data models.RondaAttendance) (*models.RondaAttendance, error)
	FindAll() (models.RondaAttendances, error)
	Paginated(pagination utils.Pagination) (*utils.Pagination, models.RondaAttendances, error)
	FindByID(id string) (*models.RondaAttendance, error)
	Delete(id string) error
}

type rondaAttendanceRepositoryImpl struct {
	db *gorm.DB
}

func NewRondaAttendanceRepository(db *gorm.DB) RondaAttendanceRepository {
	return &rondaAttendanceRepositoryImpl{db: db}
}

func (r *rondaAttendanceRepositoryImpl) Paginated(pagination utils.Pagination) (*utils.Pagination, models.RondaAttendances, error) {
	var datas models.RondaAttendances
	query := r.db
	err := query.Scopes(utils.Paginate(datas, &pagination, query)).Find(&datas).Error
	return &pagination, datas, err
}

func (r *rondaAttendanceRepositoryImpl) Create(tx *gorm.DB, data models.RondaAttendance) (*models.RondaAttendance, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Create(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *rondaAttendanceRepositoryImpl) Update(tx *gorm.DB, id string, data models.RondaAttendance) (*models.RondaAttendance, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Model(&models.RondaAttendance{}).Where("id = ?", id).Updates(data).Error
	return &data, err
}

func (r *rondaAttendanceRepositoryImpl) FindByID(id string) (*models.RondaAttendance, error) {
	var data models.RondaAttendance

	err := r.db.First(&data, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *rondaAttendanceRepositoryImpl) FindAll() (models.RondaAttendances, error) {
	var data models.RondaAttendances
	err := r.db.Find(&data).Error
	return data, err
}

func (r *rondaAttendanceRepositoryImpl) Delete(id string) error {
	err := r.db.Delete(&models.RondaAttendance{}, "id = ?", id).Error
	return err
}
