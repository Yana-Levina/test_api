package main

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"log"
	"net/url"
	_ "time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
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
	//dbHost := viper.GetString(`database.host`)
	//dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	//connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	//dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	//dbConn, err := sql.Open(`postgres`, dsn)
	dbConn, err := sql.Open("postgres", "user="+dbUser+" password="+dbPass+" dbname="+dbName+" sslmode=disable")

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
	//middL := _personHttpDeliveryMiddleware.InitMiddleware()
	//e.Use(middL.CORS)
	//
	//person := _personRepo.NewMysqlPersonRepository(dbConn)
	//
	//timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	//au := _articleUcase.NewArticleUsecase(ar, authorRepo, timeoutContext)
	//_articleHttpDelivery.NewArticleHandler(e, au)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
