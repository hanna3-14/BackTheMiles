package domain

type IEventRepository interface {
	CreateEvent(event Event) error
	FindAllEvents() ([]Event, error)
	FindEventByID(id int) (Event, error)
	UpdateEvent(id int, event, modifiedEvent Event) error
	DeleteEvent(id int) error
}
