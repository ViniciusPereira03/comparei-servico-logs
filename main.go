package main

import (
	"comparei-servico-logs/config"
	"comparei-servico-logs/internal/app"
	customHTTP "comparei-servico-logs/internal/infrastructure/http"
	"comparei-servico-logs/internal/infrastructure/repository"
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		log.Fatal("Erro ao carregar configurações")
	}

	// Configuração da conexão com o MySQL usando variáveis de ambiente
	dsn := os.Getenv("MYSQL_USER") + ":" + os.Getenv("MYSQL_PASSWORD") + "@tcp(" + os.Getenv("MYSQL_HOST") + ")/" + os.Getenv("MYSQL_DB")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Verificar a conexão com o MySQL
	if err := db.Ping(); err != nil {
		log.Fatal("Não foi possível conectar ao MySQL: ", err)
	}

	mysqlRepo := repository.NewMySQLRepository(db)
	if mysqlRepo == nil {
		log.Fatal("mysqlRepo está nil")
	}

	logService := app.NewLogService(mysqlRepo)
	customHTTP.IniHandlers(logService)

	router := customHTTP.NewRouter(logService)

	log.Println("Servidor iniciado na porta " + os.Getenv("PORT"))
	http.ListenAndServe(":"+os.Getenv("PORT"), router)
}
