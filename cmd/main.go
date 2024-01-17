package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	logger "github.com/Antoha2/auth/cmd/logger"
	"github.com/Antoha2/auth/config"
	"github.com/Antoha2/auth/repository"
	"github.com/Antoha2/auth/services"

	// "github.com/Antoha2/auth/transport/grpc"
	transport "github.com/Antoha2/auth/transport/http"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	Run()
}

func Run() {

	cfg := config.MustLoad()
	log := logger.SetupLogger(logger.EnvLocal)

	gormDB, err := initDb(cfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	authRep := repository.NewRepAuth(gormDB)
	authServ := services.NewServAuth(*authRep, log)
	authTrans := transport.NewWeb(authServ, log, cfg.HTTP.HostAddr)

	go authTrans.StartHTTP()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
	authTrans.Stop()

}

func initDb(cfg *config.Config) (*gorm.DB, error) {

	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.DBConfig.User,
		cfg.DBConfig.Password,
		cfg.DBConfig.Host,
		cfg.DBConfig.Port,
		cfg.DBConfig.Dbname,
		cfg.DBConfig.Sslmode,
	)

	connConfig, err := pgx.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("1 failed to parse config: %v", err)
	}

	// Make connections
	dbx, err := sqlx.Open("pgx", stdlib.RegisterConnConfig(connConfig))
	if err != nil {
		return nil, fmt.Errorf("2 failed to create connection db: %v", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: dbx,
	}), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("3 gorm.Open(): %v", err)
	}

	err = dbx.Ping()
	if err != nil {
		return nil, fmt.Errorf("4 error to ping connection pool: %v", err)
	}
	log.Printf("Подключение к базе данных на http://127.0.0.1:%d\n", cfg.DBConfig.Port)
	return gormDB, nil
}
