package SD

import (
	"errors"
	_ "strings"
)

type DataType string
type ConstraintType string
type ServerType string

const (
	NUMERIC   DataType = "NU"
	INTEGER   DataType = "IN"
	DATE      DataType = "DA"
	DATETIME  DataType = "DT"
	TIMESTAMP DataType = "TS"
	CHAR      DataType = "CR"
	VARCHAR   DataType = "VC"
)
const (
	PRIMARY_KEY ConstraintType = "P"
	FOREIGN_KEY ConstraintType = "R"
	UNIQUE_KEY  ConstraintType = "U"
	CHECK       ConstraintType = "C"
)
const (
	ORACLE ServerType = "ORACLE"
	MSSQL  ServerType = "MSSQL"
	SYBASE ServerType = "SYBASE"
)

type Column struct {
	name      string
	datatype  DataType
	position  int
	length    int
	precision int
	scale     int
	nullable  bool
}

type Constraint struct {
	name              string
	owner             string
	tableName         string
	cType             ConstraintType
	columns           *[]Column
	refOwner          string
	refConstraintName string
	refTableName      string
	refColumns        *[]Column
	checkExpression   string
}

type Table struct {
	owner   string
	name    string
	server  *Server
	columns []*Column
	keys    []*Constraint
	refKeys []*Constraint
}
type ServerCredential struct{
	login    string
	password string
}
type Server struct {
	stype    ServerType
	name     string
	driver   string
	url      string
	cred     ServerCredential
}

func newTable(server *Server, owner string, name string) (err error, result Table) {
	if server == nil {
		return errors.New("Server reference is empty!"), result
	}

	result = Table{
		server: server,
		owner:  owner,
		name:   name,
	}
	return nil, result
}
func newServer(stype ServerType,name string, driver string, url string) (err error, result Server) {
	result = Server{
		stype : stype,
		name:name,
		driver:driver,
		url:url,
	}
	return nil,result
}

func (server *Server) setCredential( cred ServerCredential) {
	server.cred = cred
}



