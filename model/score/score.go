package score

import (
	"database/sql"
	"golearn/first-api/db"
)

type Score struct {
	ID         int64  `json:"id"`
	OwnerID    int64  `json:"owner_id"`
	PlayerID   int64  `json:"player_id"`
	CategoryID int64  `json:"category_id"`
	Value      string `json:"value"`
}

func Get(id int64) (*Score, error) {
	stmt, err := db.DB.Prepare("SELECT * FROM Score WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var score Score
	err = stmt.QueryRow(id).Scan(
		&score.ID,
		&score.OwnerID,
		&score.CategoryID,
		&score.PlayerID,
		&score.Value,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &score, nil
}

func GetByPlayerID(playerID int64) ([]Score, error) {
	stmt, err := db.DB.Prepare("SELECT * FROM Score WHERE player_id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(playerID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var scores []Score
	for rows.Next() {
		var score Score
		err := rows.Scan(
			&score.ID,
			&score.OwnerID,
			&score.CategoryID,
			&score.PlayerID,
			&score.Value,
		)

		if err != nil {
			return nil, err
		}

		scores = append(scores, score)
	}

	return scores, nil
}

func (s Score) Delete() error {
	stmt, err := db.DB.Prepare("DELETE FROM Score WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(s.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Score) UpdateWith(data Score) error {
	stmt, err := db.DB.Prepare("UPDATE Score SET value = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(data.Value, s.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Score) Save(ownerID int64, playerID int64) error {
	stmt, err := db.DB.Prepare(`
    INSERT INTO Score(owner_id, player_id, category_id, value)
    VALUES (?, ?, ?, ?)
  `)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(ownerID, playerID, s.CategoryID, s.Value)
	if err != nil {
		return err
	}

	return nil
}
