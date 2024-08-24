package models

import (
	"errors"
	"time"
)

// TODO 제목, 내용, 카테고리, 등록일,  삭제, 승인 상태 정보

type Category uint
type Confirm uint

// 승인 상태 : 예약, 승인, 거절
const (
	pending Confirm = iota + 1
	approve
	reject
)

// 카테고리 : 진로, 건강, 인간관계, 돈, 목표, 학업/성적, 기타
const (
	course Category = iota + 1
	health
	relationship
	money
	objective
	record
	etc
)

var confirmMap = map[string]Confirm{
	"reserve": pending,
	"approve": approve,
	"reject":  reject,
}

var categoryMap = map[string]Category{
	"course":       course,
	"health":       health,
	"relationship": relationship,
	"money":        money,
	"objective":    objective,
	"record":       record,
	"etc":          etc,
}

type Wish struct {
	ID        uint       `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	Category  Category   `json:"category"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Confirm   Confirm    `json:"is_confirm"`
}

func StringToCategory(category string) (Category, error) {
	if c, ok := categoryMap[category]; ok {
		return c, nil
	}
	return 0, errors.New("invalid category")
}

func StringToConfirm(confirm string) (Confirm, error) {
	if c, ok := confirmMap[confirm]; ok {
		return c, nil
	}
	return 0, errors.New("invalid confirm")
}

func IsValidConfirm(confirm Confirm) bool {
	switch confirm {
	case approve, reject:
		return true
	default:
		return false
	}
}
