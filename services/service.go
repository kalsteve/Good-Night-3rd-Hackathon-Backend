package services

import "github.com/kalsteve/Good-Night-3rd-Hackathon-Backend/repositories"

// Service groups all individual services.
type Service struct {
	WishService    WishService
	CommentService CommentService
}

// NewService creates a new instance of Service with all required services.
func NewService(repo *repositories.Repository) *Service {
	return &Service{
		WishService:    NewWishService(repo.WishRepository),
		CommentService: NewCommentService(repo.CommentRepository),
	}
}
