package main

import "database/sql"

func main() {

}

func checkDeleteDB(p string) error {
	return nil
}

func createDB(p string) (*Database, error) {
	return &Database{}, nil
}

func (d *Database) createTable(n string, c Columns) error {
	return nil
}

func loadJSON(p string) ([]byte, error) {
	return []byte{}, nil
}

func (t *Table) injectJSON(p string) error {
	return nil
}

func (t *Table) query() {

}

type Database sql.DB

type Table struct{}

type Columns map[string]string

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}
