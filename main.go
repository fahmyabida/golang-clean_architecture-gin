package main

import (
	"database/sql"
	"fmt"
	_pemesanRepo "github.com/fahmyabida/golang-clean_architecture-gin/domain/pemesan/repository"
	_orderRepo "github.com/fahmyabida/golang-clean_architecture-gin/domain/order/repository"
	_jenisPesananRepo "github.com/fahmyabida/golang-clean_architecture-gin/domain/jenis_pesanan/repository"
	_menuRepo "github.com/fahmyabida/golang-clean_architecture-gin/domain/menu/repository"
	_invoiceUsecase "github.com/fahmyabida/golang-clean_architecture-gin/domain/invoice/usecase"
	_workOrderUsecase "github.com/fahmyabida/golang-clean_architecture-gin/domain/work_order/usecase"
	_invoiceHttpDeliver "github.com/fahmyabida/golang-clean_architecture-gin/domain/invoice/delivery/http"
	_workOrderHttpDeliver "github.com/fahmyabida/golang-clean_architecture-gin/domain/work_order/delivery/http"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
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

	jpRepo 	:= _jenisPesananRepo.NewJenisPesananRepository(dbConn)
	mRepo 	:= _menuRepo.NewMenuRepository(dbConn)
	oRepo 	:= _orderRepo.NewOrderRepository(dbConn)
	pRepo 	:= _pemesanRepo.NewPemesanRepository(dbConn)

	iU := _invoiceUsecase.NewInvoiceMenuUsecase(jpRepo,mRepo,oRepo,pRepo,timeoutContext)
	woU := _workOrderUsecase.NewWorkOrderUsecase(jpRepo,mRepo,oRepo,pRepo,timeoutContext)

	_invoiceHttpDeliver.NewInvoiceHttpHandler(engine, iU)
	_workOrderHttpDeliver.NewWorkOrderHttpHandler(engine, woU)

	engine.Run(viper.GetString("server.address"))
}
