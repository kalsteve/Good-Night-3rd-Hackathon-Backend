package models

import "time"

// TODO 댓글 - 소원 번호, 댓글 내용, 등록일, 삭제
type Comment struct {
	ID        uint       `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Content   string     `json:"content"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	WishID    uint       `json:"wish_id"`
	Wish      Wish       `json:"wish" gorm:"foreignKey:WishID"`
}
