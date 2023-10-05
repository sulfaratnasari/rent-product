package xorm

import (
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"xorm.io/core"
	"xorm.io/xorm"
)


const (
	DriverMySQL    = "mysql"
	DriverPostgres = "postgres"
	PingTick       = 30
)

type Connection struct {
	Master *xorm.Engine
	Slave *xorm.Engine
}

// DB Config
type DBConfig struct {
	SlaveDSN string
	MasterDSN string
	DBEncKey string
	ConnMaxLifetime time.Duration
	RetryInterval int
	MaxIdleConn int
	MaxConn int
	isTest bool
}

// DB Config
type DB struct {
	DBConnection *xorm.Engine
	DBString string
	ConnMaxLifetime time.Duration
	RetryInterval int
	MaxIdleConn int
	MaxConn int
	doneChannel chan bool
	isTest bool
}

var (
	Master *DB
	Slave *DB
	dbTicker *time.Ticker
)

func New(cfg DBConfig, dbDriver string, tbPrefix ...string) (*Connection, error) {
	masterDSN := cfg.MasterDSN

	Master = &DB{
		DBString: masterDSN,
		RetryInterval: 30,
		MaxIdleConn: 10,
		MaxConn: 200,
		ConnMaxLifetime: 8,
		doneChannel: make(chan bool),
	}
	err := Master.ConnectionAndMonitor(dbDriver, tbPrefix...)
	if err != nil {
		return nil, err
	}
	dbTicker = time.NewTicker(time.Second * 2)
	return &Connection{Master: Master.DBConnection}, nil
}

func (d *DB) ConnectionAndMonitor(driver string, tbPrefix ...string) error {
	err := d.Connect(driver, tbPrefix...)
	if err != nil {
		fmt.Println("error ConnectionAndMonitor : ", err)
		return err
	} else {
		fmt.Println("Success connecting to database")
	}
	ticker := time.NewTicker(time.Duration(d.RetryInterval) * time.Second)
	go func() error {
		for {
			select {
			case <-ticker.C:
				err := d.DBConnection.Ping()
				if err != nil {
					fmt.Println("db reconnect error : ", err.Error())
					continue
				}
			case <-d.doneChannel:
				return nil
			}
		}
	}()
	return nil
}

func (d *DB) Connect(driver string, tbPrefix ...string) error {
	var db *xorm.Engine
	var err error
	db, err = xorm.NewEngine(driver, d.DBString)
	if err != nil {
		return err
	}
	if tbPrefix != nil {
		tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, tbPrefix[0])
		db.SetTableMapper(tbMapper)
		db.SetColumnMapper(core.SameMapper{})
	}
	d.DBConnection = db

	if d.MaxConn > 0 {
		db.SetMaxOpenConns(d.MaxConn)
	}

	if d.MaxIdleConn > 0 {
		db.SetMaxIdleConns(d.MaxIdleConn)
	}

	if d.ConnMaxLifetime > 0 {
		db.SetConnMaxLifetime(d.ConnMaxLifetime)
	}
	return err
}