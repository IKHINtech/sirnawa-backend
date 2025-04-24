package services

import (
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/response"
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/internal/repository"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type PostService interface {
	Create(data request.PostCreateRequest) (*response.PostResponse, error)
	Update(id string, data request.PostUpdateRequset) (*response.PostResponse, error)
	FindByID(id string) (*response.PostResponse, error)
	Delete(id string) error
	FindAll(rtID string) (response.PostResponses, error)
	Paginated(pagination utils.Pagination, rtID string) (*utils.Pagination, *response.PostResponses, error)
}

type postServiceImpl struct {
	repository repository.PostRepository
	db         *gorm.DB
}

func NewPostServices(repo repository.PostRepository, db *gorm.DB) PostService {
	return &postServiceImpl{repository: repo, db: db}
}

func (s *postServiceImpl) withTransaction(fn func(tx *gorm.DB) error) error {
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

func (s *postServiceImpl) Create(data request.PostCreateRequest) (*response.PostResponse, error) {
	var result *models.Post

	err := s.withTransaction(func(tx *gorm.DB) error {
		payload := request.PostCreateRequestToPostModel(data)

		created, err := s.repository.Create(tx, payload)
		if err != nil {
			return err
		}

		result = created
		return nil
	})
	// handling err
	if err != nil {
		return nil, err
	}

	res := response.PostModelToPostResponse(result)
	return res, nil
}

func (s *postServiceImpl) Update(id string, data request.PostUpdateRequset) (*response.PostResponse, error) {
	var result *models.Post

	err := s.withTransaction(func(tx *gorm.DB) error {
		existing, err := s.repository.FindByID(id)
		if err != nil {
			return err
		}

		payload := request.PostUpdateRequsetToPostModel(data)
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

	res := response.PostModelToPostResponse(result)
	return res, nil
}

func (s *postServiceImpl) FindAll(rtID string) (response.PostResponses, error) {
	result, err := s.repository.FindAll(rtID)
	if err != nil {
		return nil, err
	}

	resp := response.PostListToResponse(result)
	return resp, nil
}

func (s *postServiceImpl) FindByID(id string) (*response.PostResponse, error) {
	result, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	resp := response.PostModelToPostResponse(result)
	return resp, err
}

func (s *postServiceImpl) Paginated(pagination utils.Pagination, rtID string) (*utils.Pagination, *response.PostResponses, error) {
	paginated, data, err := s.repository.Paginated(pagination, rtID)
	if err != nil {
		return nil, nil, err
	}

	resp := response.PostListToResponse(data)
	return paginated, &resp, err
}

func (s *postServiceImpl) Delete(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}
