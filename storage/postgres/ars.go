package postgres

import (
	"fmt"
	"strings"

	"github.com/Jopoleon/rustamViewer/models"
	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

func (db *DB) GetWaveRecordByID(id int) (*models.ASR, error) {
	res := models.ASR{}
	err := db.DB.Get(&res, "SELECT * FROM asrresults WHERE id=$1;", id)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	return &res, nil
}

func (db *DB) GetWaveRecordByProfileNames(profileNames []string) ([]models.ASR, error) {
	res := []models.ASR{}
	param := "{" + strings.Join(profileNames, ",") + "}"
	err := db.DB.Select(&res, "SELECT * FROM asrresults WHERE project_id = ANY($1) LIMIT 100;", param)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	return res, nil
}

func (db *DB) GetWaveRecordByFilters(profileName string) ([]models.ASR, error) {
	res := []models.ASR{}
	err := db.DB.Select(&res, "SELECT * FROM asrresults WHERE profile=$1;", profileName)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	return res, nil
}

func (db *DB) GetAllAsrResults() ([]models.ASR, error) {
	res := []models.ASR{}
	err := db.DB.Select(&res, "SELECT id,ani,dnis,profile,"+
		"utterance,interpretation,confidence,created_on FROM asrresults;")
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	return res, nil
}

func (db *DB) GetAllAsrWithFilters(filters map[string][]string) ([]models.ASR, error) {
	res := []models.ASR{}
	qq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	q := qq.Select("id,ani,dnis,profile," +
		"utterance,interpretation,confidence,created_on").From("asrresults")
	confidence := ""
	confidenceType := ""
	for k, v := range filters {
		if v[0] != "" {
			switch {
			case k == "ani":
				q = q.Where(sq.Expr("ani=?", v[0]))
			case k == "dnis":
				q = q.Where(sq.Expr("dnis=?", v[0]))
			case k == "profileName":
				q = q.Where(sq.Expr("profile=?", v[0]))
			case k == "utterance":
				q = q.Where(sq.Expr("utterance LIKE ?", "%"+v[0]+"%"))
			case k == "confidenceType":
				confidenceType = v[0]
			case k == "confidence":
				confidence = v[0]
			}
		}
	}
	if confidence != "" && confidenceType != "" {
		q = q.Where(sq.Expr(fmt.Sprintf("confidence %s ?", confidenceType), confidence))
	}
	sss, arr, err := q.ToSql()
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	err = db.DB.Select(&res, sss, arr...)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	return res, nil
}
