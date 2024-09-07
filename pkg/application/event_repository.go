package application

import "github.com/hanna3-14/BackTheMiles/pkg/domain"

type IEventRepository interface {
	CreateEvent(event domain.Event) error
	FindAllEvents() ([]domain.Event, error)
	FindEventByID(id int) (domain.Event, error)
	UpdateEvent(id int, event, modifiedEvent domain.Event) error
	DeleteEvent(id int) error
}
