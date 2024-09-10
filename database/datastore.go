package database

type Row []interface{}

type DataStore struct {
	rows []Row
}

func (d *DataStore) Append(r Row) {
	d.rows = append(d.rows, r)
}

func (d *DataStore) Filter(filters map[int]interface{}) []Row {
	result := []Row{}

	for _, row := range d.rows {
		rowMatched := true
		for idx, filter := range filters {
			if row[idx] != filter {
				rowMatched = false
			}
		}

		if rowMatched {
			result = append(result, row)
		}
	}

	return result
}
