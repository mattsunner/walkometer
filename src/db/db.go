package dataParser

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Walk struct {
	walkingDate int
	walkingDist float64
}

func LoadDatabase(engine string, file string) (*sql.DB, error) {
	db, err := sql.Open(engine, file)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func InsertWalk(db *sql.DB, date int, distance float64) error {
	_, err := db.Exec("INSERT INTO walking_distance (walking_date, walking_dist) VALUES (?, ?)", date, distance)
	return err
}

func GetWalks(db *sql.DB) ([]Walk, error) {
	rows, err := db.Query("SELECT * FROM walking_distance")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var walks []Walk
	for rows.Next() {
		var walk Walk
		if err := rows.Scan(&walk.walkingDate, &walk.walkingDist); err != nil {
			return nil, err
		}
		walks = append(walks, walk)
	}
	return walks, nil
}
