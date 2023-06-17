package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Data struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
	DateTime  string    `json:"datetime"`
}

var (
	queryCreateTable = `CREATE TABLE yourls_log (
		"click_id" int NOT NULL AUTO_INCREMENT,
		"click_time" datetime NOT NULL,
		"shorturl" varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
		"referrer" varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
		"user_agent" varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
		"ip_address" varchar(41) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
		"country_code" char(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
		PRIMARY KEY ("click_id"),
		KEY "shorturl" ("shorturl")
	  ) ENGINE=InnoDB`
	queryInsert = `insert into query_perform_table(name,address,datetime)values(?,?,?)`
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	for i := 10000; i < 50000; i++ {
		initial := strconv.Itoa(i)
		Insert(db, queryInsert, "ade mawan "+initial, "cibeureum "+initial, time.Now().Format("2006-01-02 15:04:05"))
	}
	fmt.Println("SUCCESS")

}

func CreateTable(db *sql.DB, query string) {
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}

}
func Insert(db *sql.DB, query string, arg ...any) {
	_, err := db.Exec(query, arg...)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

}
