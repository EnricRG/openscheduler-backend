package schedule

import "github.com/EnricRG/openscheduler-backend/paging"

type UserSchedulesRetriever interface {
	Get(userId string, pageSort paging.PagedSort) ([]ScheduleDefinition, error)
	GetById(userId, id string) (ScheduleDefinition, error)
}
