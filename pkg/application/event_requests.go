package application

import (
	"github.com/hanna3-14/BackTheMiles/pkg/domain"
)

type EventRequestService struct {
	repo   IEventRepository
	erRepo IEventResultRepository
}

func NewEventRequestService(eventRepo IEventRepository, eventResultRepo IEventResultRepository) (EventRequestService, error) {
	return EventRequestService{repo: eventRepo, erRepo: eventResultRepo}, nil
}

func (r *EventRequestService) GetEvents() ([]domain.Event, error) {
	events, err := r.repo.FindAllEvents()
	if err != nil {
		return []domain.Event{}, err
	}

	eventResults, err := r.erRepo.FindAllEventResults()
	if err != nil {
		return []domain.Event{}, err
	}

	for i := range events {
		events[i].ResultIDs = eventResults[events[i].ID]
	}
	return events, nil
}

func (r *EventRequestService) GetEvent(id int) (domain.Event, error) {
	event, err := r.repo.FindEventByID(id)
	if err != nil {
		return domain.Event{}, err
	}

	eventResults, err := r.erRepo.FindAllEventResults()
	if err != nil {
		return domain.Event{}, err
	}

	event.ResultIDs = eventResults[event.ID]
	return event, nil
}

func (r *EventRequestService) PostEvent(event domain.Event) error {
	return r.repo.CreateEvent(event)
}

func (r *EventRequestService) PatchEvent(id int, modifiedEvent domain.Event) error {
	storedEvent, err := r.repo.FindEventByID(id)
	if err != nil {
		return err
	}
	return r.repo.UpdateEvent(id, storedEvent, modifiedEvent)
}

func (r *EventRequestService) DeleteEvent(id int) error {
	r.erRepo.DeleteEventResult(id)
	return r.repo.DeleteEvent(id)
}
