package db

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"os"
)

type DBConfig struct {
	DBUsername *string
	DBPassword *string
	DBName     *string
	DBHost     *string
	DBPort     *int
	DBRestart  *bool
}

var Config DBConfig

var DB *sql.DB

func init() {
	ParsingDBConfig()
}

func ParsingDBConfig() {
	Config = DBConfig{
		DBUsername: flag.String("dbusername", "postgres", "database username of our website"),
		DBPassword: flag.String("dbpassword", "200103287sdu", "database password of our website"),
		DBName:     flag.String("dbname", "online_shop", "database name of our website"),
		DBHost:     flag.String("dbhost", "127.0.0.1", "database host of our website"),
		DBPort:     flag.Int("dbport", 5432, "database port of our website"),
		DBRestart:  flag.Bool("dbrestart", false, "needs to restart to database"),
	}
}

func InitDB() {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
		*Config.DBHost, *Config.DBPort, *Config.DBUsername, *Config.DBPassword, *Config.DBName,
	)
	var err error
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("We have problems with connection database", err)
		os.Exit(1)
	}
	if *Config.DBRestart {
		configDB()
		fmt.Println("Database restarted")
	}
}

func configDB() {
	tables()
	triggers()
}

func tables() {
	st, ioErr := ioutil.ReadFile("db/setup.sql")
	if ioErr != nil {
		fmt.Println("Cannot read data/setup.sql")
		os.Exit(1)
	}
	if _, err := DB.Exec(string(st)); err != nil {
		fmt.Println(err)
	}
}

func triggers() {
	st, ioErr := ioutil.ReadFile("db/triggers.sql")
	if ioErr != nil {
		fmt.Println("Cannot read data/triggers.sql")
		os.Exit(1)
	}
	if _, err := DB.Exec(string(st)); err != nil {
		fmt.Println(err)
	}
}
