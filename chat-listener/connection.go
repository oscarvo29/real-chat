package main

import (
	"context"

	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// var counts int
// var DB *sql.DB

// var dbTimeout = time.Second * 3

func GetConnection(dsn string) *pgxpool.Pool {
	ctx := context.Background()

	conn, err := pgxpool.New(ctx, dsn)
	if err != nil {
		panic(err)
	}

	return conn
}

// func GetConnection(dsn string) {
// 	for {
// 		connection, err := openConnection(dsn)
// 		if err != nil {
// 			log.Println("Postgress not yet ready..")
// 			counts++
// 		} else {
// 			log.Println("Connected to Postgres!")
// 			DB = connection
// 			return
// 		}

// 		if counts > 10 {
// 			log.Println(err)
// 			return
// 		}

// 		log.Println("back off for three seconds")
// 		time.Sleep(dbTimeout)
// 		continue
// 	}
// }

// func openConnection(dsn string) (*sql.DB, error) {
// 	db, err := sql.Open("pgx", dsn)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = db.Ping()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return db, nil
// }
