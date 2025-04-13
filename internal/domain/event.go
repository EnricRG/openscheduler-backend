package domain

// An schedulable event
type Event interface {
	Id() uint
	ScheduleId() uint
	Name() string
}
