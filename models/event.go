package models

import (
	"eventBooking/db"
	"fmt"
	"time"
)

type Event struct {
	ID          int64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time
	UserID      int64
}

var events = []Event{}

func (e *Event) Save() error {
	query :=
		`INSERT INTO events(name, description, location, dateTime, user_id) 
		 VALUES (?, ?, ?, ?, ?)`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println(err)
		fmt.Println("HERE1")
		return err
	}
	defer statement.Close()
	result, err := statement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	if err != nil {
		fmt.Println(err)
		fmt.Println("HERE2")
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		fmt.Println("HERE3")
		return err
	}
	e.ID = id
	fmt.Println("Event Inserted successfully", e.ID)
	return err
}

func GetAllEvents() ([]Event, error) { // return type is event slice
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		fmt.Println(err)
		fmt.Println("HERE5")
		return nil, err
	}
	defer rows.Close()
	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			fmt.Println("HERE6")
			fmt.Println(err)
			return nil, err
		}
		fmt.Println(event)
		events = append(events, event)
	}

	return events, nil

}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id=?"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}
	return &event, nil

}

func (event Event) Update() error {
	query :=
		`UPDATE events 
		 SET name = ?, description = ?, location = ?, dateTime = ?
		 WHERE id=?`

	statement, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err

}

func (event Event) Delete() error {
	query := "DELETE FROM events WHERE id=?"
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(event.ID)

	return err
}

func (event Event) Register(userId int64) error {

	query := `INSERT INTO registrations(event_id, user_id) VALUES (?, ?)`

	statement, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(event.ID, userId)

	return err
}

func (event Event) CancelRegistration(userId int64) error {

	query := `DELETE FROM registrations WHERE event_id=? AND user_id=?`

	statement, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(event.ID, userId)

	return err

}
