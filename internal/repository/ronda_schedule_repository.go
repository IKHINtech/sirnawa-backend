package repository

import (
	"time"

	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type RondaScheduleRepository interface {
	Create(tx *gorm.DB, data models.RondaSchedule) (*models.RondaSchedule, error)
	Update(tx *gorm.DB, id string, data models.RondaSchedule) (*models.RondaSchedule, error)
	FindAll(rtID, groupID string, date *time.Time) (models.RondaSchedules, error)
	Paginated(pagination utils.Pagination, rtID, groupID string, date *time.Time) (*utils.Pagination, models.RondaSchedules, error)
	FindByID(id string) (*models.RondaSchedule, error)
	Delete(id string) error
}

type rondaScheduleRepositoryImpl struct {
	db *gorm.DB
}

func NewRondaScheduleRepository(db *gorm.DB) RondaScheduleRepository {
	return &rondaScheduleRepositoryImpl{db: db}
}

func (r *rondaScheduleRepositoryImpl) Paginated(pagination utils.Pagination, rtID, groupID string, date *time.Time) (*utils.Pagination, models.RondaSchedules, error) {
	var datas models.RondaSchedules
	query := r.db.Preload("Rt").Preload("Group")

	if rtID != "" {
		query = query.Where("rt_id = ?", rtID)
	}

	if groupID != "" {
		query = query.Where("group_id = ?", groupID)
	}

	if date != nil {
		query = query.Where("date = ?", date.Format("2006-01-02"))
	}
	err := query.Scopes(utils.Paginate(datas, &pagination, query)).Find(&datas).Error
	return &pagination, datas, err
}

func (r *rondaScheduleRepositoryImpl) Create(tx *gorm.DB, data models.RondaSchedule) (*models.RondaSchedule, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Create(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *rondaScheduleRepositoryImpl) Update(tx *gorm.DB, id string, data models.RondaSchedule) (*models.RondaSchedule, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.Model(&models.RondaSchedule{}).Where("id = ?", id).Updates(data).Error
	return &data, err
}

func (r *rondaScheduleRepositoryImpl) FindByID(id string) (*models.RondaSchedule, error) {
	var data models.RondaSchedule

	err := r.db.Preload("Rt").Preload("Group").First(&data, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (r *rondaScheduleRepositoryImpl) FindAll(rtID, groupID string, date *time.Time) (models.RondaSchedules, error) {
	var data models.RondaSchedules
	query := r.db.Preload("Rt").Preload("Group")

	if rtID != "" {
		query = query.Where("rt_id = ?", rtID)
	}

	if groupID != "" {
		query = query.Where("group_id = ?", groupID)
	}

	if date != nil {
		query = query.Where("date = ?", date.Format("2006-01-02"))
	}
	err := query.Find(&data).Error
	return data, err
}

func (r *rondaScheduleRepositoryImpl) Delete(id string) error {
	err := r.db.Delete(&models.RondaSchedule{}, "id = ?", id).Error
	return err
}
