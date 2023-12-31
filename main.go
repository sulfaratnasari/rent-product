package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"rent-product/internal/config"
	"rent-product/internal/entity/constant"
	producthandler "rent-product/internal/handler/product"
	orderrepo "rent-product/internal/repo/order"
	productrepo "rent-product/internal/repo/product"
	stockitemrepo "rent-product/internal/repo/stock_item"
	productuc "rent-product/internal/usecase/product"
	xdb "rent-product/lib/database/xorm"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	cfg := getConfigJSON()
	xormConn, err := getXORMConnection(cfg)
	if err != nil {
		log.Fatalln("err xorm , ", err)
	}

	orderDB := &orderrepo.Conn{
		DB: xormConn,
	}
	orderDB.DB.Master = xormConn.Master
	orderDB.DB.Slave = xormConn.Slave

	stockItemDB := &stockitemrepo.Conn{
		DB: xormConn,
	}
	stockItemDB.DB.Master = xormConn.Master
	stockItemDB.DB.Slave = xormConn.Slave

	productDB := &productrepo.Conn{
		DB: xormConn,
	}
	productDB.DB.Master = xormConn.Master
	productDB.DB.Slave = xormConn.Slave

	productUc := productuc.New(&productuc.Usecase{
		OrderDB:     orderDB,
		StockItemDB: stockItemDB,
		ProductDB:   productDB,
	})
	prodHandler := producthandler.New(&producthandler.ProductHandler{
		ProductUC: productUc,
	})

	r := mux.NewRouter()
	fmt.Println("Starting server ......")
	r.HandleFunc("/product/{id}", prodHandler.ProductAvailability).Methods("GET")
	r.HandleFunc("/product/add", prodHandler.AddProduct).Methods("POST")
	r.HandleFunc("/list/product", prodHandler.ListProduct).Methods("GET")

	log.Fatalln(http.ListenAndServe(":3333", r))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePagesss!")
	fmt.Println("Endpoint Hit: homePage")
}

func getConfigJSON() (cfg *config.ConfigCredential) {
	dir, _ := os.Getwd()
	vaultPath := dir + "/files/"
	//open file
	vaultPath = vaultPath + constant.APP_NAME + ".json"

	vaultFile, err := os.Open(vaultPath)
	if err != nil {
		log.Fatalln("Path not found: ", err)
	}

	configByte, err := ioutil.ReadAll(vaultFile)
	if err != nil {
		log.Fatalln("Path not found: ", err)
	}

	cfg = &config.ConfigCredential{}
	err = json.Unmarshal(configByte, cfg)
	if err != nil {
		log.Fatalln("Failed get config: ", err)
	}
	if cfg.Data == nil {
		log.Fatalln("Config Data nil")
	}
	if cfg.Data.Database == nil {
		log.Fatalln("database config nil")
	}
	return cfg
}

func getXORMConnection(cfg *config.ConfigCredential) (*xdb.Connection, error) {
	dbConfig := xdb.DBConfig{
		MasterDSN: cfg.Data.Database.Master,
	}
	tbPrefix := strings.Split(cfg.Data.Database.Prefix, ",")
	xorm, err := xdb.New(dbConfig, xdb.DriverPostgres, tbPrefix...)
	if err != nil {
		databaseErr := errors.New("main.getXORMConnection" + err.Error())
		log.Fatalln(databaseErr)
		return nil, databaseErr
	}
	return xorm, nil
}
