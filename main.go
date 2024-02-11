// Walkometer - A command line application for entering and tracking walking data
// from my standing desk treadmill. CLI is used to be a quick, seamless entry point
// for data to be entered into a persistent data store.
//
// Author: Matthew Sunner, 2024

package main

import (
	"flag"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"

	dp "github.com/mattsunner/walkometer/src/db"
)

func main() {
	modeFlag := flag.String("mode", "query", "Query or Post Mode Flag")
	datePtr := flag.Int("date", 0, "Date of Walk (YYYYMMDD)")
	distPtr := flag.Float64("dist", 0, "Distance of Walk")

	flag.Parse()

	db, err := dp.LoadDatabase("sqlite3", "walks.db")
	if err != nil {
		log.Fatal(err)
	}

	if *modeFlag == "query" {
		walks, err := dp.GetWalks(db)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(walks)

	} else if *modeFlag == "post" {
		err := dp.InsertWalk(db, *datePtr, *distPtr)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		log.Fatal(err)
	}
}
