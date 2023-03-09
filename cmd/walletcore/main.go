package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/raphaelmb/fullcycle-walletcore/internal/database"
	"github.com/raphaelmb/fullcycle-walletcore/internal/event"
	createaccount "github.com/raphaelmb/fullcycle-walletcore/internal/usecase/create_account"
	createclient "github.com/raphaelmb/fullcycle-walletcore/internal/usecase/create_client"
	createtransaction "github.com/raphaelmb/fullcycle-walletcore/internal/usecase/create_transaction"
	"github.com/raphaelmb/fullcycle-walletcore/internal/web"
	"github.com/raphaelmb/fullcycle-walletcore/internal/web/webserver"
	"github.com/raphaelmb/fullcycle-walletcore/pkg/events"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/wallet?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	eventDispatcher := events.NewEventDispatcher()
	transactionCreatedEvent := event.NewTransactionCreated()
	// eventDispatcher.Register("TransactionCreated", handler)

	clientDb := database.NewClientDB(db)
	accountDb := database.NewAccountDB(db)
	transactionDb := database.NewTransactionDB(db)

	createClientUseCase := createclient.NewCreateClientUseCase(clientDb)
	createAccountUseCase := createaccount.NewCreateAccountUseCase(accountDb, clientDb)
	createTransactionUseCase := createtransaction.NewCreateTransactionUseCase(transactionDb, accountDb, eventDispatcher, transactionCreatedEvent)

	webserver := webserver.NewWebServer(":3000")

	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUseCase)

	webserver.AddHandler("/clients", clientHandler.CreateClient)
	webserver.AddHandler("/accounts", accountHandler.CreateAccount)
	webserver.AddHandler("/transactions", transactionHandler.CreateTransaction)

	webserver.Start()
}
