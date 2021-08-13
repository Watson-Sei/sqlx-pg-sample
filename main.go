package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Tag struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}

func main() {
	db, err := sqlx.Connect("postgres", "host=pg user=admin password=password dbname=sqlx_db sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	var id int64
	err = db.QueryRowx(`INSERT INTO tags (name) VALUES ($1) on conflict (name) do UPDATE SET name=$1 RETURNING id`, "Python").Scan(&id)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(id)
}
