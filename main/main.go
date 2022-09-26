package main

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	_personHttpDelivery "test/http"
	_personLogic "test/logic"
	_personRepo "test/postgres"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	dbDriver := viper.GetString(`database.driver`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	dbConn, err := sql.Open(dbDriver, "user="+dbUser+" password="+dbPass+" dbname="+dbName+" sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()

	pr := _personRepo.NewPersonRepository(dbConn)
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	pu := _personLogic.NewPersonUsecase(pr, timeoutContext)
	_personHttpDelivery.NewPersonHandler(e, pu)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
