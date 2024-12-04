package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not open database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createGamesTable := `
  CREATE TABLE IF NOT EXISTS Game (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    color INTEGER NOT NULL DEFAULT 0,
    scoring_mode INTEGER NOT NULL DEFAULT 1
  );`

	_, err := DB.Exec(createGamesTable)
	if err != nil {
		panic("Could not create 'Game' table.")
	}

	createMatchesTable := `
  CREATE TABLE IF NOT EXISTS Match (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    game_id INTEGER NOT NULL,
    notes TEXT,
    dateTime DATETIME NOT NULL,
    location TEXT,
    FOREIGN KEY(game_id)
      REFERENCES Game(id)
      ON DELETE CASCADE
      ON UPDATE CASCADE
  );
  CREATE INDEX IF NOT EXISTS index_Match_game_id ON Match (game_id);`

	_, err = DB.Exec(createMatchesTable)
	if err != nil {
		panic("Could not create 'Match' table")
	}

	createPlayersTable := `
  CREATE TABLE IF NOT EXISTS Player (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    match_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    position INTEGER NOT NULL DEFAULT 0,
    FOREIGN KEY(match_id)
      REFERENCES Match(id)
      ON DELETE CASCADE
      ON UPDATE CASCADE
  );
  CREATE INDEX IF NOT EXISTS index_Player_match_id ON Player (match_id);`

	_, err = DB.Exec(createPlayersTable)
	if err != nil {
		panic("Could not create 'Player' table")
	}

	createCategoriesTable := `
  CREATE TABLE IF NOT EXISTS Category (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    game_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    position INT NOT NULL DEFAULT 0,
    FOREIGN KEY(game_id)
      REFERENCES Game(id)
      ON DELETE CASCADE
      ON UPDATE CASCADE
  );
  CREATE INDEX IF NOT EXISTS index_Category_game_id ON Category (game_id);`

	_, err = DB.Exec(createCategoriesTable)
	if err != nil {
		panic("Could not create categories table.")
	}

	createScoresTable := `
  CREATE TABLE IF NOT EXISTS scores (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    category_id INTEGER NOT NULL,
    player_id INTEGER NOT NULL,
    value DECIMAL,
    FOREIGN KEY(category_id)
      REFERENCES categories(id)
      ON DELETE CASCADE
      ON UPDATE CASCADE,
    FOREIGN KEY(player_id)
      REFERENCES Player(id)
      ON DELETE CASCADE
      ON UPDATE CASCADE
  );
  CREATE INDEX IF NOT EXISTS index_Score_player_id ON Score (player_id);
  CREATE INDEX IF NOT EXISTS index_Score_category_id ON SCORE (category_id);`

	_, err = DB.Exec(createScoresTable)
	if err != nil {
		panic("Could not create scores table.")
	}
}
