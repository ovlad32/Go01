package SD

import (
	"database/sql"
	"log"

	"fmt"
)

var Db *sql.DB;
func DbInit(){
	var err error;
	Db,err = sql.Open("postgres", "user=gouser password=11 dbname=fin host=52.29.37.253 port=5432 sslmode=disable")
	if err != nil {
		log.Fatal(err);
	}
}

func  Sd1()  {
	_,prod := NewBoundInt64Sequential(1, 100, 1, 0, false);
	dataPair := prod.NextValue();
	fmt.Println(dataPair.StringValue);
}
