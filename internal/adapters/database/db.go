package database

import (
	"database/sql"
	"fmt"
	"log"
	
	"github.com/Masterminds/squirrel"
	_ "github.com/mattn/go-sqlite3"
)

// Adapter implements the DbPort interface
type Adapter struct {
	db    *sql.DB
	limit uint64
}

// NewAdapter creates a new Adapter
func NewAdapter(driverName, dataSourceName string, paginationLimit uint64) (*Adapter, error) {
	// connect
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("db connection failur: %v", err)
	}
	
	// test db connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("db ping failure: %v", err)
	}
	
	return &Adapter{db: db, limit: paginationLimit}, nil
}

// CloseDbConnection closes the db  connection
func (da *Adapter) CloseDbConnection() error {
	if err := da.db.Close(); err != nil {
		return fmt.Errorf("db close failure: %v", err)
	}
	return nil
}

// GetCustomers pulls all the customers from the database
func (da *Adapter) GetCustomers(page uint64) ([]map[string]interface{}, error) {
	query := squirrel.Select("id", "name", "phone").From("customer")
	
	//Simple pagination
	if da.limit != 0 && page != 0 {
		query = query.Offset((page - 1) * da.limit).Limit(da.limit)
	}
	rows, err := query.RunWith(da.db).Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var customers []map[string]interface{}
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}
		if err := rows.Scan(columnPointers...); err != nil {
			return nil, err
		}
		row := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			row[colName] = *val
		}
		customers = append(customers, row)
	}
	
	return customers, nil
}
