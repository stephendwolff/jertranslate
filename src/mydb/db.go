package mydb

import (
	"database/sql"
//	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type connection struct {
	db *sql.DB
}

type dbRow map[string]interface{}

var (
	conn *connection
	dsn  string
)

func dbScan(rows *sql.Rows) (dbRow, error) {
	r := dbRow{}

	cols, _ := rows.Columns()
	c := len(cols)
	vals := make([]interface{}, c)
	valPtrs := make([]interface{}, c)

	for i := range cols {
		valPtrs[i] = &vals[i]
	}

	var err = rows.Scan(valPtrs...)
	if err != nil {
		return nil, err
	}

	for i := range cols {
		if val, ok := vals[i].([]byte); ok {
			r[cols[i]] = string(val)
		} else {
			r[cols[i]] = vals[i]
		}
	}

	return r, nil
}

func getConn() (*connection, error) {
	if conn == nil {
		conn = new(connection)
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			return nil, err
		}
		conn.db = db
	}

	return conn, nil
}

func Select(query string, args ...interface{}) ([]dbRow, error) {
	c, err := getConn()
	if err != nil {
		return nil, err
	}

	rows, err := c.db.Query(query, args...)

	if err != nil {
		return nil, err
	}

	res := make([]dbRow, 0)

	for rows.Next() {
		scan, err := dbScan(rows)
		if err != nil {
			return nil, err
		}
		res = append(res, scan)
	}

	return res, nil
}

func Exec(query string, args ...interface{}) (int64, int64, error) {
	c, err := getConn()
	if err != nil {
		return 0, 0, err
	}

	stmt, err := c.db.Prepare(query)
	if err != nil {
		return 0, 0, err
	}

	res, err := stmt.Exec(args...)
	if err != nil {
		return 0, 0, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, 0, err
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		return 0, 0, err
	}

	return lastId, rowCount, nil
}

func Close() error {
	c, err := getConn()
	if err != nil {
		return err
	}

	return c.db.Close()
}

/*
func Transaction(f func(*sql.Tx) error) error {
	c, err := getConn()
	if err != nil {
		return err
	}

	tx, err := c.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			switch p := p.(type) {
			case error:
				err = p
			default:
				err = fmt.Errorf("%s", p)
			}
		}

		if err != nil {
			tx.Rollback()
			return
		}

		err = tx.Commit()
	}()

	return f(tx)
}
*/

func SetDSN(new_dsn string) {
	dsn = new_dsn
}
