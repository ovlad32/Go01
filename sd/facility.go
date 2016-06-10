package SD

import (
	"database/sql"
	"log"

	"fmt"
	"os"
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
	prod := NewSimpleTime()
	dataPair := prod.NextValue();
	fmt.Println(dataPair.StringValue);
	dataPair = prod.NextValue();
	fmt.Println(dataPair.StringValue);
}


func Sd2() {
	bb := new(B)
	bb.overFunc = bb.Over;
	bb.Call();

}
func Sd3 (){
	file,err := os.OpenFile("text.txt",os.O_CREATE|os.O_WRONLY,0600)
	if err != nil {
		panic(err);
	}
	defer file.Close();
	for i:=0;i<1000000; i++ {
		file.WriteString("%");
	}
}

type A struct {
	value int;
	overFunc func()
}

type B struct {
	A
}

type Aer interface {
	Over()
	Call()
}

func (a *A) Over()  {
	a.value = 1;
}

func (a *A) Call()  {
	if a.overFunc == nil{
		panic("tada!")
	} else {
		a.overFunc()
	}
	fmt.Println(a.value)
}

type Ber interface {
	Aer
}

func (b *B) Over()  {
	b.value = 2;
}



