package main

import (
	"context"
	"database/sql"
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
	"github.com/raphaelmb/fullcycle-walletcore-ms/internal/database"
	"github.com/raphaelmb/fullcycle-walletcore-ms/internal/event"
	"github.com/raphaelmb/fullcycle-walletcore-ms/internal/event/handler"
	createaccount "github.com/raphaelmb/fullcycle-walletcore-ms/internal/usecase/create_account"
	createclient "github.com/raphaelmb/fullcycle-walletcore-ms/internal/usecase/create_client"
	createtransaction "github.com/raphaelmb/fullcycle-walletcore-ms/internal/usecase/create_transaction"
	"github.com/raphaelmb/fullcycle-walletcore-ms/internal/web"
	"github.com/raphaelmb/fullcycle-walletcore-ms/internal/web/webserver"
	"github.com/raphaelmb/fullcycle-walletcore-ms/pkg/events"
	"github.com/raphaelmb/fullcycle-walletcore-ms/pkg/kafka"
	"github.com/raphaelmb/fullcycle-walletcore-ms/pkg/uow"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/wallet?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
	}

	kafkaProducer := kafka.NewKafkaProducer(&configMap)

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("TransactionCreated", handler.NewTransactionCreatedKafkaHandler(kafkaProducer))
	eventDispatcher.Register("BalanceUpdated", handler.NewUpdateBalanceKafkaHanlder(kafkaProducer))
	transactionCreatedEvent := event.NewTransactionCreated()
	balanceUpdatedEvent := event.NewBalanceUpdated()

	clientDb := database.NewClientDB(db)
	accountDb := database.NewAccountDB(db)

	ctx := context.Background()
	uow := uow.NewUow(ctx, db)

	uow.Register("AccountDB", func(tx *sql.Tx) any {
		return database.NewAccountDB(db)
	})

	uow.Register("TransactionDB", func(tx *sql.Tx) any {
		return database.NewTransactionDB(db)
	})

	createClientUseCase := createclient.NewCreateClientUseCase(clientDb)
	createAccountUseCase := createaccount.NewCreateAccountUseCase(accountDb, clientDb)
	createTransactionUseCase := createtransaction.NewCreateTransactionUseCase(uow, eventDispatcher, transactionCreatedEvent, balanceUpdatedEvent)

	webserver := webserver.NewWebServer(":8080")

	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUseCase)

	webserver.AddHandler("/clients", clientHandler.CreateClient)
	webserver.AddHandler("/accounts", accountHandler.CreateAccount)
	webserver.AddHandler("/transactions", transactionHandler.CreateTransaction)

	fmt.Println("Server is running")
	webserver.Start()
}
