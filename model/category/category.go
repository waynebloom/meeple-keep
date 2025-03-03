package category

import (
	"database/sql"
	"golearn/first-api/db"
)

type Category struct {
	ID       int64  `json:"id"`
	OwnerID  int64  `json:"owner_id"`
	GameID   int64  `json:"game_id"`
	Name     string `json:"name"`
	Position int    `json:"position"`
}

func Get(id int64) (*Category, error) {
	stmt, err := db.DB.Prepare("SELECT * FROM Category WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var score Category
	err = stmt.QueryRow(id).Scan(
		&score.ID,
		&score.OwnerID,
		&score.GameID,
		&score.Name,
		&score.Position,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &score, nil
}

func GetByGameID(gameID int64) ([]Category, error) {
	stmt, err := db.DB.Prepare("SELECT * FROM Category WHERE game_id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(gameID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var categories []Category
	for rows.Next() {
		var score Category
		err := rows.Scan(
			&score.ID,
			&score.OwnerID,
			&score.GameID,
			&score.Name,
			&score.Position,
		)

		if err != nil {
			return nil, err
		}

		categories = append(categories, score)
	}

	return categories, nil
}

func (c Category) Delete() error {
	stmt, err := db.DB.Prepare("DELETE FROM Category WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(c.ID)
	if err != nil {
		return err
	}

	return nil
}

func (c *Category) UpdateWith(data Category) error {
	stmt, err := db.DB.Prepare("UPDATE Category SET name = ?, position = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(data.Name, data.Position, c.ID)
	if err != nil {
		return err
	}

	return nil
}

func (c *Category) Save(ownerID int64, gameID int64) error {
	stmt, err := db.DB.Prepare(`
    INSERT INTO Category(owner_id, game_id, name, position)
    VALUES (?, ?, ?, ?)
  `)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(ownerID, gameID, c.Name, c.Position)
	if err != nil {
		return err
	}

	return nil
}
