package models

import (
	"time"

	"example.com/go_events_api/pkgs/db"
)

type Event struct {
	ID          int64
	Name        string    `bindings:"required"`
	Description string    `bindings:"required"`
	Location    string    `bindings:"required"`
	Datetime    time.Time `bindings:"required"`
	UserID      int64
}

func (e *Event) Save() error {
	// Todo: Add it to database
	query := `
	INSERT INTO events (name, description, location, datetime, user_id)
	VALUES (?, ?, ?, ?, ?)
	`

	// Prepare() prepares a SQL statement - this can lead to better performance
	// if the same statement is executed multiple times (potentially with different data for its placeholders).
	// This is only true, if the prepared statement is not closed (stmt.Close()) in between those executions.
	// In that case, there wouldn't be any advantages.
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close() // executed when the Save() func completes

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.Datetime, e.UserID) // Exec method is used when we go for Create, update delete ect

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"

	rows, err := db.DB.Query(query) // Query method is used if wwe are fetching a data

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Datetime, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (Event, error) {
	query := "SELECT * FROM events where id = ?"

	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Datetime, &event.UserID)

	if err != nil {
		return Event{}, err
	}

	return event, nil
}

func (e Event) Update() error {
	query := `
	UPDATE events
	set name = ?, description = ?, location = ?, datetime = ?
	where id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.Datetime, e.ID)

	return err

}

func (e Event) Delete() error {
	query := "DELETE FROM events where id = ?"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID)

	return err
}

func (e Event) Register(userId int64) error {
	query := "INSERT INTO registrations(event_id, user_id) VALUES(?, ?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.ID, userId)

	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	return err
}

func (e Event) CancelRegistration(userId int64) error {
	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.ID, userId)

	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	return err
}
