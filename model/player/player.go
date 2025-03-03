package player

import (
	"database/sql"
	"golearn/first-api/db"
)

type Player struct {
	ID       int64  `json:"id"`
	OwnerID  int64  `json:"owner_id"`
	MatchID  int64  `json:"match_id"`
	Name     string `json:"name"`
	Position int    `json:"position"`
	// score
}

func Get(id int64) (*Player, error) {
	stmt, err := db.DB.Prepare("SELECT * FROM Player WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var player Player
	err = stmt.QueryRow(id).Scan(
		&player.ID,
		&player.OwnerID,
		&player.MatchID,
		&player.Name,
		&player.Position,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &player, nil
}

func GetByMatchID(matchID int64) ([]Player, error) {
	stmt, err := db.DB.Prepare("SELECT * FROM Player WHERE match_id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(matchID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var players []Player
	for rows.Next() {
		var player Player
		err := rows.Scan(
			&player.ID,
			&player.OwnerID,
			&player.MatchID,
			&player.Name,
			&player.Position,
		)

		if err != nil {
			return nil, err
		}

		players = append(players, player)
	}

	return players, nil
}

func (p Player) Delete() error {
	stmt, err := db.DB.Prepare("DELETE FROM Player WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.ID)
	if err != nil {
		return err
	}

	return nil
}

func (p *Player) UpdateWith(data Player) error {
	stmt, err := db.DB.Prepare(`
    UPDATE Player
    SET name = ?, position = ?
    WHERE id = ?
  `)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(data.Name, data.Position, p.ID)
	if err != nil {
		return err
	}

	return nil
}

func (p *Player) Save(ownerID int64, matchID int64) error {
	stmt, err := db.DB.Prepare(`
    INSERT INTO Player(owner_id, match_id, name, position)
    VALUES (?, ?, ?, ?)
  `)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(ownerID, matchID, p.Name, p.Position)
	if err != nil {
		return err
	}

	return nil
}
