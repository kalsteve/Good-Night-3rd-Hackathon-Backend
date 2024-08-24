package repositories

import (
	"github.com/kalsteve/Good-Night-3rd-Hackathon-Backend/models"
	"gorm.io/gorm"
)

/**
 * TODO: 댓글 기능
 * 1. 댓글 생성
 * 2. 소원에 따른 댓글 조회 - 페이지네이션
 * 3. 댓글 소프트 삭제
 * 4. 댓글 하드 삭제
 */

type CommentRepository interface {
	Create(comment *models.Comment) error
	FindAllByID(id uint, page int, size int) ([]models.Comment, error)
	SoftDeleteByID(id uint) error
	HardDeleteByID(id uint) error
	FindListByWishID(wishID uint) ([]models.Comment, error)
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{db: db}
}

func (r *commentRepository) Create(comment *models.Comment) error {
	return r.db.Create(comment).Error
}

func (r *commentRepository) FindAllByID(wishID uint, page int, size int) ([]models.Comment, error) {
	var comments []models.Comment
	err := r.db.Where("wish_id = ?", wishID).Offset((page - 1) * size).Limit(size).Find(&comments).Error
	return comments, err
}

func (r *commentRepository) FindListByWishID(wishID uint) ([]models.Comment, error) {
	var comments []models.Comment
	err := r.db.Where("wish_id = ?", wishID).Find(&comments).Error
	return comments, err
}

func (r *commentRepository) SoftDeleteByID(id uint) error {
	return r.db.Delete(&models.Comment{}, id).Error
}

func (r *commentRepository) HardDeleteByID(id uint) error {
	return r.db.Unscoped().Delete(&models.Comment{}, id).Error
}
