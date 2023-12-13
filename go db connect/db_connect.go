package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
  host     = "localhost"
  port     = 5433
  user     = "postgres"
  dbname   = "postgres"
	password = "1"
)

func connectToDb()(*sql.DB){
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "dbname=%s password=%s sslmode=disable",
    host, port, user, dbname, password)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }

  err = db.Ping()
  if err != nil {
    panic(err)
  }

  fmt.Println("Successfully connected!")

	return db
}

func createTable(db *sql.DB){
	rows, err := db.Query("CREATE TABLE Erentest(id int NOT NULL, name varchar(255), PRIMARY KEY (id) )");

	if(err==nil){
		fmt.Println("created",rows,err)
	}else{
		fmt.Println("failed")
	}
}

func insertVal(db *sql.DB, id int, name string){
	query:= fmt.Sprintf("INSERT INTO Erentest VALUES (%d, '%s')",id,name)
	fmt.Println(query)
	rows, err := db.Query(query);

	if(err==nil){
		fmt.Println("inserted successfully",rows)
	}else{
		fmt.Println("insertion failed",err)
	}
}

func readAll(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM Erentest;")
	if err != nil {
		fmt.Println("read failed", err)
		return
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		fmt.Println("failed to get columns", err)
		return
	}

	// Create a slice of interface{} to hold the values for each row
	values := make([]interface{}, len(columns))
	for i := range values {
		values[i] = new(interface{})
	}

	// Iterate over rows
	for rows.Next() {
		// Scan the values into the interface{} pointers
		err := rows.Scan(values...)
		if err != nil {
			fmt.Println("scan failed", err)
			return
		}

		// Print the values
		for i, col := range columns {
			fmt.Printf("%s: %v\t", col, *values[i].(*interface{}))
		}
		fmt.Println()
	}

	if err := rows.Err(); err != nil {
		fmt.Println("iteration failed", err)
		return
	}
}

func main() {
  db:=connectToDb()
	//createTable(db)
	insertVal(db,2,"ekin")
	readAll(db)

	defer db.Close()
}