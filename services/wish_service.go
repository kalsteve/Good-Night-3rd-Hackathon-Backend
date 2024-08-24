package services

import (
	"fmt"
	"github.com/kalsteve/Good-Night-3rd-Hackathon-Backend/models"
	"github.com/kalsteve/Good-Night-3rd-Hackathon-Backend/repositories"
)

type WishService interface {
	GetWishByID(id uint) (*models.Wish, error)
	CreateWish(wish *models.Wish) error
	GetWishListByConfirm(confirm models.Confirm, page int, size int) ([]models.Wish, error)
	DeleteWish(id uint) error
	UpdateWish(id uint, confirm models.Confirm) error
}

type wishService struct {
	repo repositories.WishRepository
}

func NewWishService(repo repositories.WishRepository) WishService {
	return &wishService{repo: repo}
}

func (s *wishService) GetWishByID(id uint) (*models.Wish, error) {
	wish, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return wish, nil
}

func (s *wishService) CreateWish(wish *models.Wish) error {
	return s.repo.Create(wish)
}

func (s *wishService) GetWishListByConfirm(confirm models.Confirm, page int, size int) ([]models.Wish, error) {
	return s.repo.FindListByConfirm(confirm, page, size)
}

func (s *wishService) DeleteWish(id uint) error {
	return s.repo.SoftDeleteByID(id)
}

func (s *wishService) UpdateWish(id uint, confirm models.Confirm) error {

	wish, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("failed to find wish by id: %w", err)
	}

	if !models.IsValidConfirm(confirm) {
		return fmt.Errorf("invalid confirm status: %v", confirm)
	}

	return s.repo.UpdateByConfirm(wish, confirm)
}
