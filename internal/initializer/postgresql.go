package initializer

import (
	"database/sql"
	"fmt"
	"log"
	"plbooking_go_structure1/global"
	db "plbooking_go_structure1/internal/db/sqlc"

	_ "github.com/lib/pq"
)

// initialize Postgres with SQLC
func InitPostgreSQLC() {
	// initialize a new database
	pg := global.Config.PostgreSQL
	dns := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", pg.Username, pg.Password, pg.Host, pg.Port, pg.DBName)
	conn, err := sql.Open("postgres", dns)
	// conn, err
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}

	store := db.NewStore(conn)
	fmt.Println("initialized db successfully")
	global.Pgdbc = *store
}
