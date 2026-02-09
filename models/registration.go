package models

import (
	"go-project.com/go-project/db"
)

type Registration struct {
	ID      int64
	EventID int64 `json:"-"`
	UserID  int64 `json:"-"`
	Event   Event
	User    User
}

func (r Registration) Save() error {
	var query string = `INSERT INTO registrations(event_id, user_id) VALUES(?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(r.EventID, r.UserID)
	if err != nil {
		return err
	}

	return nil
}

func GetAllRegistrations() ([]Registration, error) {
	query := `
		SELECT r.id, e.id, e.name, e.location, e.dateTime, u.id, u.email
		FROM registrations r
		LEFT JOIN events e ON r.event_id = e.id
		LEFT JOIN users u ON r.user_id = u.id
	`

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var registrations []Registration

	for rows.Next() {
		var reg Registration

		err := rows.Scan(
			&reg.ID,
			&reg.Event.ID,
			&reg.Event.Name,
			&reg.Event.Location,
			&reg.Event.DateTime,
			&reg.User.ID,
			&reg.User.Email,
		)

		if err != nil {
			return nil, err
		}

		registrations = append(registrations, reg)
	}

	return registrations, nil
}

func (r Registration) Delete() error {
	var query string = `DELETE FROM registrations WHERE event_id = ? AND user_id = ?`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(r.EventID, r.UserID)
	if err != nil {
		return err
	}

	return nil
}
