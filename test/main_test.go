package test

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/LeMinh0706/lala-song/internal/db"
	"github.com/LeMinh0706/lala-song/util"

	_ "github.com/lib/pq"
)

var testQueries *db.Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../")
	if err != nil {
		log.Fatal("can't load config:", err)
	}
	fmt.Println("Run:", config.DBDriver)
	testDB, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("can't open db:", err)
	}
	testQueries = db.New(testDB)
	os.Exit(m.Run())
}
