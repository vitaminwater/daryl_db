package daryl_db

import (
	"database/sql"
)

func Insert(s interface{}) (sql.Result, error) {
	q, err := InsertQuery(s)
	if err != nil {
		return nil, err
	}
	return NamedExec(q, s)
}
