package services

import (
	"github.com/kalsteve/Good-Night-3rd-Hackathon-Backend/models"
	"github.com/kalsteve/Good-Night-3rd-Hackathon-Backend/repositories"
)

type CommentService interface {
	CreateComment(comment *models.Comment) error
	GetCommentListByWishID(wishID uint) ([]models.Comment, error)
	DeleteComment(id uint) error
}

type commentService struct {
	repo repositories.CommentRepository
}

func NewCommentService(repo repositories.CommentRepository) CommentService {
	return &commentService{repo: repo}
}

func (s *commentService) CreateComment(comment *models.Comment) error {
	return s.repo.Create(comment)
}

func (s *commentService) GetCommentListByWishID(wishID uint) ([]models.Comment, error) {
	return s.repo.FindListByWishID(wishID)
}

func (s *commentService) DeleteComment(id uint) error {
	return s.repo.SoftDeleteByID(id)
}
