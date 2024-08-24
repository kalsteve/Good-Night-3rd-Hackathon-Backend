package repositories

import (
	"github.com/kalsteve/Good-Night-3rd-Hackathon-Backend/models"
	"gorm.io/gorm"
)

/**
 * TODO
 * 1. 소원 생성
 * 2. 소원 소프트 삭제
 * 3. 소원 하드 삭제
 * 4. 소원 단일 조회
 * 5. 소원 목록 조회 - 페이지네이션
 * 6. (추가) 소원 검색
 */

type WishRepository interface {
	Create(wish *models.Wish) error
	SoftDeleteByID(id uint) error
	HardDeleteByID(id uint) error
	FindByID(id uint) (*models.Wish, error)
	FindListByConfirm(confirm models.Confirm, page int, size int) ([]models.Wish, error)
	UpdateByConfirm(wish *models.Wish, confirm models.Confirm) error
}

type wishRepository struct {
	db *gorm.DB
}

func NewWishRepository(db *gorm.DB) WishRepository {
	return &wishRepository{db: db}
}

func (r *wishRepository) Create(post *models.Wish) error {
	return r.db.Create(post).Error
}

func (r *wishRepository) SoftDeleteByID(id uint) error {
	return r.db.Delete(&models.Wish{}, id).Error
}

func (r *wishRepository) HardDeleteByID(id uint) error {
	return r.db.Unscoped().Delete(&models.Wish{}, id).Error
}

func (r *wishRepository) FindByID(id uint) (*models.Wish, error) {
	var wish models.Wish
	err := r.db.First(&wish, id).Error
	return &wish, err
}

func (r *wishRepository) FindListByConfirm(confirm models.Confirm, page int, size int) ([]models.Wish, error) {
	var wishes []models.Wish
	query := r.db.Where("confirm = ?", confirm)

	if page > 0 && size > 0 {
		query = query.Offset((page - 1) * size).Limit(size)
	}

	if err := query.Find(&wishes).Error; err != nil {
		return nil, err
	}

	return wishes, nil
}

func (r *wishRepository) UpdateByConfirm(wish *models.Wish, confirm models.Confirm) error {
	return r.db.Model(wish).Update("confirm", confirm).Error
}
