package models

import (
	"time"

	"example.org/rest-api/db"
)

type Event struct {
	Id         int64
	Name       string    `binding:"required"`
	Location   string    `binding:"required"`
	Desciption string    `binding:"required"`
	DateTime   time.Time `binding:"required"`
	UserId     int64
}

func (event Event) CancelRegistration(userId int64) error {
	query := `
	DELETE FROM registrations
	WHERE user_id=? AND event_id=?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(userId, event.Id)
	return err
}

func (event Event) CreateRegistration(userId int64) error {
	query := `
	INSERT INTO registrations (user_id,event_id)
	VALUES(?,?)
	`
	queryStatement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer queryStatement.Close()
	_, err = queryStatement.Exec(userId, event.Id)
	if err != nil {
		return err
	}
	return err
}

func (event Event) DeleteById() error {
	query := `
	DELETE FROM events
	WHERE id=?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.Id)
	return err
}

func (e Event) Update() error {
	query := `
 UPDATE events
 SET name=?,location=?,description=?,dateTime=?
 WHERE id=?
 `
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Desciption, e.Location, e.DateTime, e.Id)
	return err
}

func GetEventById(id int64) (*Event, error) {
	query := "select * from events where id=?"
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.Id, &event.Name, &event.Desciption,
		&event.Location, &event.DateTime, &event.UserId)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (e Event) Save() (*Event, error) {
	query := `
	INSERT INTO events(name,description,location,dateTime,user_id)
	VALUES (?,?,?,?,?)
	`
	queryStatement, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer queryStatement.Close()
	result, err := queryStatement.Exec(e.Name, e.Desciption, e.Location, e.DateTime, e.UserId)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	e.Id = id
	return &e, err
}

func GetEvents() ([]Event, error) {
	query := `select * from events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event
	for rows.Next() {
		var event Event
		rows.Scan(&event.Id, &event.Name, &event.Desciption,
			&event.Location, &event.DateTime, &event.UserId)
		events = append(events, event)
	}
	return events, nil
}
