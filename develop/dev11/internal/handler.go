package internal

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"
)

type Handler struct {
	calendar Calendar
}

func (h *Handler) InitRoutes(calendar Calendar) {
	h.calendar = calendar

	// Обработка endpoints
	http.Handle("/create_event", Logging(http.HandlerFunc(h.CreateEvent)))
	http.Handle("/update_event", Logging(http.HandlerFunc(h.UpdateEvent)))
	http.Handle("/delete_event", Logging(http.HandlerFunc(h.DeleteEvent)))
	http.Handle("/events_for_day", Logging(http.HandlerFunc(h.EventsPerDay)))
	http.Handle("/events_for_week", Logging(http.HandlerFunc(h.EventsPerWeek)))
	http.Handle("/events_for_month", Logging(http.HandlerFunc(h.EventsPerMonth)))
}

func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "bad request"}`))

		return
	}

	var event Event

	body, err := io.ReadAll(r.Body)
	parseJSON := make(map[string]string)
	json.Unmarshal(body, &parseJSON)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "bad request - incorrect body"}`))

		return
	}

	event.UserID, err = strconv.Atoi(parseJSON["user_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "bad request - incorrect user id"}`))

		return
	}

	event.Date, err = time.Parse("2006-02-02", parseJSON["date"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "bad request - incorrect date"}`))

		return
	}

	event.Event = parseJSON["event"]

	h.calendar.CreateEvent(event)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"result": "created"}`))
}

func (h *Handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "bad request"}`))

		return
	}

	var event Event

	body, err := io.ReadAll(r.Body)
	parseJSON := make(map[string]string)
	json.Unmarshal(body, &parseJSON)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "bad request - incorrect body"}`))

		return
	}

	event.UserID, err = strconv.Atoi(parseJSON["user_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "bad request - incorrect user id"}`))

		return
	}

	event.Date, err = time.Parse("2006-02-02", parseJSON["date"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "bad request - incorrect date"}`))

		return
	}

	event.Event = parseJSON["event"]

	event.ID, err = strconv.Atoi(parseJSON["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "bad request - incorrect id"}`))

		return
	}

	h.calendar.UpdateEvent(event)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"result": "updated"}`))
}

func (h *Handler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "bad request"}`))

		return
	}

	var eventId int

	body, err := io.ReadAll(r.Body)
	parseJSON := make(map[string]string)
	json.Unmarshal(body, &parseJSON)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "bad request - incorrect body"}`))

		return
	}

	eventId, err = strconv.Atoi(parseJSON["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "bad request - incorrect id"}`))

		return
	}

	h.calendar.DeleteEvent(eventId)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"result": "deleted"}`))
}

func (h *Handler) EventsPerDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "bad request"}`))

		return
	}

	param := r.URL.Query()
	dateStr := param.Get("date")

	date, err := time.Parse("2006-02-02", dateStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "bad request - incorrect date"}`))

		return
	}

	events, err := h.calendar.GetEventPerDay(date)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(`{"error": "server error"}`))

		return
	}

	if len(events) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "event not found"}`))

		return
	}

	resJSON, err := json.Marshal(events)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(`{"error": "server error"}`))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resJSON)
}

func (h *Handler) EventsPerWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "bad request"}`))

		return
	}

	param := r.URL.Query()
	dateStr := param.Get("date")

	date, err := time.Parse("2006-02-02", dateStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "bad request - incorrect date"}`))

		return
	}

	events, err := h.calendar.GetEventPerWeek(date)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(`{"error": "server error"}`))

		return
	}

	if len(events) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "event not found"}`))

		return
	}

	resJSON, err := json.Marshal(events)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(`{"error": "server error"}`))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resJSON)
}

func (h *Handler) EventsPerMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "bad request"}`))

		return
	}

	param := r.URL.Query()
	dateStr := param.Get("date")

	date, err := time.Parse("2006-02-02", dateStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "bad request - incorrect date"}`))

		return
	}

	events, err := h.calendar.GetEventPerMonth(date)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(`{"error": "server error"}`))

		return
	}

	if len(events) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "event not found"}`))

		return
	}

	resJSON, err := json.Marshal(events)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(`{"error": "server error"}`))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resJSON)
}
