package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	Path   string
	DB     *sql.DB
	Tables map[string]Table
}

type Table struct {
	DB      *sql.DB
	Name    string
	Columns Columns
}

type Columns []string
type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

const (
	dbPath                   = "products.db"
	prodJSONPath             = "products.json"
	createTableStatement     = "create table %s (%s)"
	insertIntoTableStatement = "insert into %s(%s) values(%s)"
	prodColumns              = "id integer not null primary key, name text, description text, price real"
)

func main() {
	if err := checkDeleteDB(dbPath); err != nil {
		log.Fatal(err)
	}
	db, err := createDB(dbPath)
	if err != nil {
		log.Fatal(err)
	}
	prodColumns := Columns{"id integer not null primary key", "name text", "description text", "price real"}
	db.createTable("products", prodColumns)
	prodJSON, err := loadJSON(prodJSONPath)
	if err != nil {
		log.Fatal(err)
	}
	tb := db.Tables["products"]
	if err := tb.injectProductJSON(prodJSON); err != nil {
		log.Fatal(err)
	}
}

func checkDeleteDB(p string) error {
	if err := os.Remove(dbPath); err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return nil
}

func createDB(p string) (Database, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return Database{}, err
	}
	return Database{dbPath, db, map[string]Table{}}, nil
}

func (d *Database) createTable(n string, c Columns) error {
	stm := fmt.Sprintf(createTableStatement, n, c)
	_, err := d.DB.Exec(stm)
	if err != nil {
		return err
	}
	d.Tables[n] = Table{d.DB, n, c}
	return nil
}

func loadJSON(p string) ([]byte, error) {
	jsonFile, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	bs, err := os.ReadFile(jsonFile.Name())
	if err != nil {
		return nil, err
	}
	return bs, nil
}

func (t *Table) injectProductJSON(bs []byte) error {
	var prods []Product
	if err := json.Unmarshal(bs, &prods); err != nil {
		return err
	}
	tx, err := t.DB.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(fmt.Sprintf(insertIntoTableStatement, t.Name, "id, name, description, price", "?, ?, ?, ?"))
	if err != nil {
		return err
	}
	for _, p := range prods {
		_, err := stmt.Exec(fmt.Sprintf("%d", p.ID), p.Name, p.Description, fmt.Sprintf("%d", p.Price))
		if err != nil {
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (t *Table) query() {

}
