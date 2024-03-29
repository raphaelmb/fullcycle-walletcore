package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/raphaelmb/fullcycle-balance-ms/internal/database"
	getbalance "github.com/raphaelmb/fullcycle-balance-ms/internal/usecase/get_balance"
	"github.com/raphaelmb/fullcycle-balance-ms/internal/web"
	"github.com/raphaelmb/fullcycle-balance-ms/internal/web/webserver"
)

func main() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// db, err := sql.Open("mysql", "root:root@tcp(mysql:3307)/balance?charset=utf8&parseTime=True&loc=Local")
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()
	// topics := []string{"transactions"}
	// configMap := &kafka.ConfigMap{
	// 	"bootstrap.servers": "kafka:29092",
	// 	"group.id":          "wallet",
	// 	"client.id":         "consumer",
	// 	"auto.offset.reset": "earliest",
	// }

	// consumer := ckafka.NewConsumer(configMap, topics)

	// go consumer.Consume()

	balance := database.NewBalanceDB(db)
	getBalanceByID := getbalance.NewGetBalanceByIDUseCase(balance)
	getBalanceHandler := web.NewWebBalanceHandler(*getBalanceByID)

	webserver := webserver.NewWebServer(":3003")
	webserver.AddHandler("/balances", getBalanceHandler.BalanceById)

	fmt.Println("Server is running")
	webserver.Start()
}
