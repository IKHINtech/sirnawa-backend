package services

import (
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/response"
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/internal/repository"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"gorm.io/gorm"
)

type PostCommentService interface {
	Create(data request.PostCommentCreateRequest) (*response.PostCommentResponse, error)
	Update(id string, data request.PostCommentUpdateRequset) (*response.PostCommentResponse, error)
	FindByID(id string) (*response.PostCommentResponse, error)
	Delete(id string) error
	FindAll() (response.PostCommentResponses, error)
	Paginated(pagination utils.Pagination) (*utils.Pagination, *response.PostCommentResponses, error)
}

type postCommentServiceImpl struct {
	repository repository.PostCommentRepository
	db         *gorm.DB
}

func NewPostCommentServices(repo repository.PostCommentRepository, db *gorm.DB) PostCommentService {
	return &postCommentServiceImpl{repository: repo, db: db}
}

func (s *postCommentServiceImpl) withTransaction(fn func(tx *gorm.DB) error) error {
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

func (s *postCommentServiceImpl) Create(data request.PostCommentCreateRequest) (*response.PostCommentResponse, error) {
	var result *models.PostComment

	err := s.withTransaction(func(tx *gorm.DB) error {
		payload := request.PostCommentCreateRequestToPostCommentModel(data)

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

	res := response.PostCommentModelToPostCommentResponse(result)
	return res, nil
}

func (s *postCommentServiceImpl) Update(id string, data request.PostCommentUpdateRequset) (*response.PostCommentResponse, error) {
	var result *models.PostComment

	err := s.withTransaction(func(tx *gorm.DB) error {
		existing, err := s.repository.FindByID(id)
		if err != nil {
			return err
		}

		payload := request.PostCommentUpdateRequsetToPostCommentModel(data)
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

	res := response.PostCommentModelToPostCommentResponse(result)
	return res, nil
}

func (s *postCommentServiceImpl) FindAll() (response.PostCommentResponses, error) {
	result, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	resp := response.PostCommentListToResponse(result)
	return resp, nil
}

func (s *postCommentServiceImpl) FindByID(id string) (*response.PostCommentResponse, error) {
	result, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	resp := response.PostCommentModelToPostCommentResponse(result)
	return resp, err
}

func (s *postCommentServiceImpl) Paginated(pagination utils.Pagination) (*utils.Pagination, *response.PostCommentResponses, error) {
	paginated, data, err := s.repository.Paginated(pagination)
	if err != nil {
		return nil, nil, err
	}

	resp := response.PostCommentListToResponse(data)
	return paginated, &resp, err
}

func (s *postCommentServiceImpl) Delete(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}
