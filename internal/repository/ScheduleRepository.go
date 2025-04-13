package repository

import "github.com/EnricRG/openscheduler-backend/internal/domain"

type ScheduleRepository interface {
	// FindById retrieves an schedule identified by its unique identifier.
	FindById(scheduleId uint) *domain.Schedule
	// FindUserSchedules retrieves al schedules from a given user.
	FindUserSchedules(userId uint) []domain.Schedule
	// AddEvent adds an event to a given schedule.
	// Will generate a copy of the given event linked to the given schedule.
	// Returns the result of saving the new copy to persistent storage. Use returned value from now on.
	AddEvent(scheduleId uint, event domain.Event) domain.Event
}
