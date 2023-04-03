package main

import (
	"fmt"

	"github.com/raphaelmb/fullcycle-balance-ms/internal/web/webserver"
)

func main() {
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

	webserver := webserver.NewWebServer(":3003")
	// webserver.AddHandler("/balances/{account_id}", web.BalanceById)

	fmt.Println("Server is running")
	webserver.Start()
}
