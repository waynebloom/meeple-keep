package match

import (
	"database/sql"
	"golearn/first-api/db"
	"time"
)

type Match struct {
	ID       int64     `json:"id"`
	OwnerID  int64     `json:"owner_id"`
	GameID   int64     `json:"game_id"`
	Notes    string    `json:"notes"`
	DateTime time.Time `json:"datetime"`
	Location string    `json:"location"`
	// players
}

func Get(id int64) (*Match, error) {
	stmt, err := db.DB.Prepare("SELECT * FROM Match WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var match Match
	err = stmt.QueryRow(id).Scan(
		&match.ID,
		&match.OwnerID,
		&match.GameID,
		&match.Notes,
		&match.DateTime,
		&match.Location,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &match, nil
}

func GetByGameID(gameID int64) ([]Match, error) {
	stmt, err := db.DB.Prepare("SELECT * FROM Match WHERE game_id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(gameID)
	if err != nil {
		return nil, err
	}

	var matches []Match
	for rows.Next() {
		var match Match
		err := rows.Scan(
			&match.ID,
			&match.OwnerID,
			&match.GameID,
			&match.Notes,
			&match.DateTime,
			&match.Location,
		)

		if err != nil {
			return nil, err
		}

		matches = append(matches, match)
	}

	return matches, nil
}

func (m Match) Delete() error {
	stmt, err := db.DB.Prepare("DELETE FROM Match WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(m.ID)
	if err != nil {
		return err
	}

	return nil
}

func (m *Match) UpdateWith(data Match) error {
	stmt, err := db.DB.Prepare(`
    UPDATE Match
    SET notes = ?, datetime = ?, location = ?
    WHERE id = ?
  `)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(data.Notes, data.DateTime, data.Location, m.ID)
	if err != nil {
		return err
	}

	return nil
}

func (m *Match) Save(ownerID int64, gameID int64) error {
	stmt, err := db.DB.Prepare(`
    INSERT INTO Match(owner_id, game_id, notes, datetime, location)
    VALUES (?, ?, ?, ?, ?)
  `)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(ownerID, gameID, m.Notes, m.DateTime, m.Location)
	if err != nil {
		return err
	}

	return nil
}
