package application

type IEventResultRepository interface {
	CreateEventResult(eventId int) error
	FindAllEventResults() (map[int][]int, error)
	UpdateEventResult(id int, resultId int, eventId int) error
	DeleteEventResult(resultId int) error
}
