package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_mahasiswaRepo "github.com/me/golang-clean_architecture-gin_gonic/mahasiswa/repository"
	_mahasiswaUsecase "github.com/me/golang-clean_architecture-gin_gonic/mahasiswa/usecase"
	_mahasiswaHttpDeliver "github.com/me/golang-clean_architecture-gin_gonic/mahasiswa/delivery/http"
	"github.com/spf13/viper"
	"log"
	"net/url"
	"os"
	"time"
)

func init(){
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)
	if err != nil && viper.GetBool("debug") {
		fmt.Println(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer dbConn.Close()
	engine := gin.Default()
	//middL := middleware.InitMiddleware()
	//e.Use(middL.CORS)
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	mhsRepo := _mahasiswaRepo.NewMysqlArticleRepository(dbConn)
	mhsU := _mahasiswaUsecase.NewArticleUsecase(mhsRepo, timeoutContext)
	_mahasiswaHttpDeliver.NewMahasiswaHttpHandler(engine, mhsU)

	engine.Run(viper.GetString("server.address"))
}
