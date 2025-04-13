package mapper

import (
	"github.com/EnricRG/resteduler-backend/api/view"
	"github.com/EnricRG/resteduler-backend/domain"
	"github.com/life4/genesis/slices"
)

type ScheduleMapper struct{}

func NewScheduleMapper() *ScheduleMapper {
	return &ScheduleMapper{}
}

func (m *ScheduleMapper) ToViews(schedules []domain.Schedule) []view.ScheduleView {
	if schedules == nil {
		return nil
	}
	return slices.Map(schedules, m.ToView)
}

func (m *ScheduleMapper) ToView(schedule domain.Schedule) view.ScheduleView {
	return view.ScheduleView{}
}
