package main

import (
	"log"

	"github.com/rizqdwan/e-wallet/internal/config"
	"github.com/rizqdwan/e-wallet/internal/db"
	"github.com/rizqdwan/e-wallet/pkg/logger"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	log := logger.New()

	database, err := db.NewDatabase(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	log.Println("Database connected successfully")


	// cfg := config.Load()
	// dbConn, err := db.NewDatabase(cfg)
	// if err != nil {
	// 	log.Fatalf("could not initialize database connection: %s", err)
	// }


	// userRepo := repository.NewUserRepository(dbConn)
  // walletRepo := repository.NewWalletRepository(dbConn)

  // walletService := service.NewWalletService(walletRepo)
  // walletHandler := handler.NewWalletHandler(walletService)

  // router := http.NewRouter(walletHandler)
  // router.Run()
}