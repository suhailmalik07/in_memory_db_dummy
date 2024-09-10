package database

import (
	"errors"
	"fmt"
)

type Table struct {
	columnsMap map[string]Column
	columns    []Column
	data       DataStore
}

func NewTable(columns []Column) Table {
	columnsMap := map[string]Column{}

	for _, col := range columns {
		columnsMap[col.name] = col
	}

	return Table{
		columnsMap: columnsMap,
		columns:    columns,
		data:       DataStore{},
	}
}

func (t *Table) Insert(rec map[string]interface{}) error {
	// Create Row
	row := Row{}

	for _, col := range t.columns {
		valForCol := rec[col.name]

		if err := col.Validate(valForCol); err != nil {
			return errors.New(fmt.Sprintf("Error: %s", err.Error()))
		}

		row = append(row, valForCol)
	}

	t.data.Append(row)

	return nil
}

func (t *Table) FetchAll() []Row {
	return t.data.rows
}

func (t *Table) FindBy(filters map[string]interface{}) []Row {
	colIdxToFilter := map[int]interface{}{}

	for idx, col := range t.columns {
		if filter, ok := filters[col.name]; ok {
			colIdxToFilter[idx] = filter
		}
	}

	return t.data.Filter(colIdxToFilter)
}
