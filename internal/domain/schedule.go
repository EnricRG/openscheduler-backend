package domain

// An schedule that contains events and its constrains
type Schedule struct {
	id     uint
	userId uint
	name   string
	events []Event
}

func (s *Schedule) Id() uint {
	return s.id
}

func (s *Schedule) Name() string {
	return s.name
}

func (s *Schedule) Events() []Event {
	return s.events
}
