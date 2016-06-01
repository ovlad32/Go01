package SD

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	_ "github.com/lib/pq"

	"database/sql"
	"log"
	"encoding/json"
)

var aa int;




func Start() {
	DbInit();

	r := mux.NewRouter();
	r.HandleFunc("/",HomeHandler);
	r.HandleFunc("/solver/{data}",solverHandler)
	r.HandleFunc("/dict/currencies",currenciesAllHandler)
	r.HandleFunc("/dict/currencies/{code}",currenciesHandler)



	http.Handle("/",r);
	http.ListenAndServe(":8080",nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"GO!");
	for _,name := range sql.Drivers() {
		fmt.Fprintf(w,"%v ",name);
	}
}

func solverHandler(w http.ResponseWriter, r *http.Request) {
	rVars := mux.Vars(r)
	rData := rVars["data"]

	fmt.Fprintf(w,"solver...%v",rData);


	rows,err := Db.Query("SELECT table_name from tables_for_export");
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	
	for rows.Next() {
            var table_name string
            if err := rows.Scan(&table_name); err != nil {
                    log.Fatal(err)
            }
            fmt.Fprintln(w,table_name)
    }

}

func currenciesAllHandler(w http.ResponseWriter, r *http.Request) {
	allCurrencies := getAllCurrencies()
	data,err := json.Marshal(allCurrencies)
	if err != nil  {
		log.Fatal("Marshal",err)
	}
	w.Write(data);
}
func currenciesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		rVars := mux.Vars(r)
		rData := rVars["code"]
		if rData == "" {
			fmt.Fprint(w,"all")
		} else {
			fmt.Fprint(w,"not all")
		}
	} else {

	}
}


