package game

import (
	"database/sql"
	"golearn/first-api/db"
)

type Game struct {
	ID          int64  `json:"id"`
	OwnerID     int64  `json:"owner_id"`
	Name        string `json:"name"`
	Color       int    `json:"color"`
	ScoringMode int    `json:"scoring_mode"`
	// matches
	// categories
}

func Get(id int64) (*Game, error) {
	stmt, err := db.DB.Prepare("SELECT * FROM Game WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var game Game
	err = stmt.QueryRow(id).Scan(
		&game.ID,
		&game.OwnerID,
		&game.Name,
		&game.Color,
		&game.ScoringMode,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &game, nil
}

func (g Game) Delete() error {
	stmt, err := db.DB.Prepare("DELETE FROM Game WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(g.ID)
	if err != nil {
		return err
	}

	return nil
}

func (g *Game) UpdateWith(data Game) error {
	stmt, err := db.DB.Prepare(`
    UPDATE Game
    SET name = ?, color = ?, scoring_mode = ?
    WHERE id = ?
  `)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(data.Name, data.Color, data.ScoringMode, g.ID)
	if err != nil {
		return err
	}

	return nil
}

func (g *Game) Save(ownerID int64) error {
	stmt, err := db.DB.Prepare(`
    INSERT INTO Game(owner_id, name, color, scoring_mode)
    VALUES (?, ?, ?, ?)
  `)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(ownerID, g.Name, g.Color, g.ScoringMode)
	if err != nil {
		return err
	}

	return nil
}
