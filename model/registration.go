package model

import "golearn/first-api/db"

type Registration struct {
	eventId int64
	userId  int64
}

func GetRoster(eventId int64) ([]Registration, error) {
	query := "SELECT * FROM registrations WHERE event_id = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var registrations []Registration
	rows, err := stmt.Query(eventId)
	for rows.Next() {
		var registration Registration
		err = rows.Scan(&registration.eventId, &registration.userId)
		if err != nil {
			return nil, err
		}

		registrations = append(registrations, registration)
	}

	return registrations, nil
}
