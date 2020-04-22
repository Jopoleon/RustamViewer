package postgres

import (
	"strconv"
	"strings"

	sq "github.com/Masterminds/squirrel"

	"github.com/Jopoleon/rustamViewer/models"
	"github.com/pkg/errors"
)

func (db *DB) GetVarsByProjectIDs(projectIDS []string) ([]models.VAR, error) {
	res := []models.VAR{}
	param := "{" + strings.Join(projectIDS, ",") + "}"
	err := db.DB.Select(&res, "SELECT * FROM var WHERE project_id = ANY($1);", param)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	return res, nil
}

func (db *DB) GetVarsByCallID(callID string, projectNames []string) ([]models.VAR, error) {
	res := []models.VAR{}
	param := "{" + strings.Join(projectNames, ",") + "}"
	err := db.DB.Select(&res, "SELECT * FROM var WHERE callid = $1 AND project_id = ANY($2);", callID, param)
	if err != nil {

		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	return res, nil
}

func (db *DB) GetVarsByFilters(projectIDS []string,
	filters map[string][]string) ([]models.VAR, error) {
	res := []models.VAR{}
	//Menu Name	ProjectName	ANI	CallID
	qq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	ids := "{" + strings.Join(projectIDS, ",") + "}"
	q := qq.Select("*").From("var").
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
					return nil, errors.WithStack(err)
				}

			case "page":
				var err error
				page, err = strconv.Atoi(v[0])
				if err != nil {
					db.Logger.Error(errors.WithStack(err))
					return nil, errors.WithStack(err)
				}

			}
		}
	}

	if limit != 0 {
		offset := limit * (page - 1)
		q = q.Limit(uint64(limit)).Offset(uint64(offset))
	}

	sss, arr, err := q.ToSql()
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	//pp.Println(sss)
	//pp.Println(arr)
	err = db.DB.Select(&res, sss, arr...)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}

	return res, nil
}
