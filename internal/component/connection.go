package component

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/qaultsabit/bengkel/internal/config"
)

func GetDatabaseConnection(conf config.Config) *sql.DB {
	dsn := fmt.Sprintf(""+
		"host=%s "+
		"port=%s "+
		"user=%s "+
		"password=%s "+
		"dbname=%s "+
		"sslmode=disable ",
		conf.DB.Host,
		conf.DB.Port,
		conf.DB.User,
		conf.DB.Pass,
		conf.DB.Name,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("error connecting to database: %s", err.Error())
	}
	return db
}
