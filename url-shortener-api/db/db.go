package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"sync"
	"time"
)

const ReconnectDuration = 30 * time.Second

type Config struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

type mysql struct {
	Config Config
	DB     *xorm.Engine
}

var once sync.Once
var instance *mysql

// Init Initializes MySQL Database. This function creates internal instance which can be called through function
// 'GetInstance'. Establishing connection to Database and start background listener for connection failure.
func Init(config Config) {
	if instance == nil {
		once.Do(
			func() {
				instance = &mysql{Config: config}
				instance.connect()
				go instance.connectionFailureListener()
			},
		)
	}
}

// GetInstance Retrieves MySQL instance.
func GetInstance() *mysql {
	if instance == nil {
		log.Println("Database is not initialized. Call 'Init' func before this call.")
		return nil
	}

	return instance
}

// connect Establishing Database connection.
func (p *mysql) connect() {
	var err error
	var dsn = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8",
		p.Config.User,
		p.Config.Pass,
		p.Config.Host,
		p.Config.Port,
		p.Config.Database,
	)
	p.DB, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		log.Println("error on connecting to Database", err.Error())
		return
	}

	err = p.DB.Ping()
	if err != nil {
		log.Println(err.Error())
		return
	}

	p.DB.ShowSQL(true)

	log.Println("successfully connected")
}

// connectionFailureListener Reconnection mechanism on connection failure.
func (p *mysql) connectionFailureListener() {
	ticker := time.NewTicker(ReconnectDuration)
	for ; true; <-ticker.C {
		if err := p.DB.Ping(); err != nil {
			fmt.Println("Database not responding: ", err.Error())
			p.DB.Close()

			fmt.Println("try to reconnect to Database ...")
			p.connect()
		}
	}
}
