package daryl_db

import (
	sq "github.com/Masterminds/squirrel"
	log "github.com/sirupsen/logrus"
)

func InsertQuery(s interface{}) (string, error) {
	fields, _ := ListField(s, "i")
	qi, args, err := sq.Insert("test").Columns(fields...).Values(ToExpr(fields...)...).ToSql()
	log.Info(args)
	if err != nil {
		return "", err
	}
	return qi, nil
}
