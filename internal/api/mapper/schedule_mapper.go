package mapper

import (
	"github.com/EnricRG/openscheduler-backend/internal/api/view"
	"github.com/EnricRG/openscheduler-backend/internal/domain"
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
