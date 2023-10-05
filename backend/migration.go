package main

import (
	"log"
	"os"
)

func migrate() {
	log.Println("migrate start")
	db := connectDB()
	defer db.Close()

	files, err := os.ReadDir(SQL_PATH)
	if err != nil {
		panic(err)
	}
	log.Println("migrate files: ", files)

	for _, file := range files {
		log.Println("migrate: " + file.Name())
		data, err := os.ReadFile(SQL_PATH + "/" + file.Name())
		if err != nil {
			panic(err)
		}
		_, err = db.Exec(string(data))
		if err != nil {
			panic(err)
		}
	}

	// indexをposts.created_atにつける
	_, err = db.Exec("CREATE INDEX created_at ON posts (created_at)")
	if err != nil {
		log.Println("index already exists: ", err)
	}

	log.Println("migrate end")
}
