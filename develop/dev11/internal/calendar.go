package internal

import (
	"sync"
	"time"
)

type Calendar interface {
	CreateEvent(ev Event)
	UpdateEvent(ev Event)
	DeleteEvent(evId int)
	GetEventPerDay(date time.Time) ([]Event, error)
	GetEventPerWeek(date time.Time) ([]Event, error)
	GetEventPerMonth(date time.Time) ([]Event, error)
}

type calendar struct {
	events  map[int]Event
	idCount int
	sync.RWMutex
}

// NewCalendar - конструктор календаря
func NewCalendar() Calendar {
	return &calendar{
		events:  make(map[int]Event),
		idCount: 0,
	}
}

// CreateEvent - создание события
func (c *calendar) CreateEvent(ev Event) {
	c.Lock()
	ev.ID = c.idCount
	c.events[ev.ID] = ev
	c.idCount++
	c.Unlock()
}

// UpdateEvent - обновление события
func (c *calendar) UpdateEvent(ev Event) {
	c.Lock()
	c.events[ev.ID] = ev
	c.Unlock()
}

// DeleteEvent - удаление события
func (c *calendar) DeleteEvent(evId int) {
	c.Lock()
	delete(c.events, evId)
	c.Unlock()
}

// GetEventPerDay - получение событий на день
func (c *calendar) GetEventPerDay(date time.Time) ([]Event, error) {
	result := make([]Event, 0)

	for _, e := range c.events {
		if e.Date == date {
			result = append(result, e)
		}
	}

	return result, nil
}

// GetEventPerWeek - получение событий на неделю
func (c *calendar) GetEventPerWeek(date time.Time) ([]Event, error) {
	result := make([]Event, 0)

	for _, e := range c.events {
		if e.Date.After(date) && e.Date.Before(date.AddDate(0, 0, 7)) {
			result = append(result, e)
		}
	}

	return result, nil
}

// GetEventPerMonth - получение событий на месяц
func (c *calendar) GetEventPerMonth(date time.Time) ([]Event, error) {
	result := make([]Event, 0)

	for _, e := range c.events {
		if e.Date.After(date) && e.Date.Before(date.AddDate(0, 1, 0)) {
			result = append(result, e)
		}
	}

	return result, nil
}
