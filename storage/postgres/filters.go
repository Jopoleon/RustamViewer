package postgres

import (
	"strconv"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

func (db *DB) BuildQueryFromFilters(tableName string, projectIDS []string, filters map[string][]string) (query string, args []interface{}, err error) {
	qq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	ids := "{" + strings.Join(projectIDS, ",") + "}"
	q := qq.Select("*").From(tableName).
		Where(sq.Expr("project_id = ANY(?)", ids))

	var page, limit int
	for k, v := range filters {
		if v[0] != "" {
			switch k {
			case "menu_name":
				q = q.Where(sq.Expr("menu_name=?", v[0]))
			case "project_id":
				q = q.Where(sq.Expr("project_id=?", v[0]))
			case "utterance":
				q = q.Where(sq.Expr("utterance LIKE ?", "%"+v[0]+"%"))
			case "limit":
				var err error
				limit, err = strconv.Atoi(v[0])
				if err != nil {
					db.Logger.Error(errors.WithStack(err))
					return "", nil, errors.WithStack(err)
				}
			case "page":
				var err error
				page, err = strconv.Atoi(v[0])
				if err != nil {
					db.Logger.Error(errors.WithStack(err))
					return "", nil, errors.WithStack(err)
				}

			}
		}
	}
	if limit != 0 && page != 0 {
		offset := limit * (page - 1)
		q = q.Limit(uint64(limit)).Offset(uint64(offset))
	}
	return q.ToSql()
}
