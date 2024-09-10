package database

import "errors"

type Database struct {
	tablesMap map[string]Table
}

func NewDatabase() Database {
	return Database{
		tablesMap: map[string]Table{},
	}
}

func (d *Database) CreateTable(name string, columns []Column) {
	table := NewTable(columns)
	d.tablesMap[name] = table
}

func (d *Database) GetTable(name string) (*Table, error) {
	if val, ok := d.tablesMap[name]; ok {
		return &val, nil
	}

	return nil, errors.New("invalid table name")
}
