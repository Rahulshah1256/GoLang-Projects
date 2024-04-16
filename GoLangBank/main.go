package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/techschool/simplebank/api"

	"github.com/techschool/simplebank/util"

	db "github.com/techschool/simplebank/db/sqlc"
)

// const (
// 	dbDriver      = "postgres"
// 	dbSource      = "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable"
// 	serverAddress = "localhost:8080"
// )

// func main() {
// 	config, err := util.LoadConfig(".")
// 	if err != nil {
// 		log.Fatal("cannot connect to db:", err)
// 	}

// 	conn, err := sql.Open(config.DBDriver, config.DBSource)
// 	if err != nil {
// 		log.Fatal("cannot connect to db:", err)
// 	}

// 	store := db.NewStore(conn)
// 	runGrpcServer(config, store)

// }

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, *store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
