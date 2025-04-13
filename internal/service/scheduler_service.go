package service

import "github.com/EnricRG/openscheduler-backend/internal/domain"

type SchedulerService interface {
	GetUserSchedules(userId uint) ([]domain.Schedule, error)
}
