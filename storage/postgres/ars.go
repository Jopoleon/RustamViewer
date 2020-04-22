package postgres

import (
	"fmt"
	"strings"

	"github.com/Jopoleon/rustamViewer/models"
	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

func (db *DB) GetWaveRecordByID(id int, projectNames []string) (*models.ASR, error) {
	res := models.ASR{}
	param := "{" + strings.Join(projectNames, ",") + "}"
	err := db.DB.Get(&res, "SELECT * FROM asrresults WHERE id=$1 AND project_id = ANY($2);", id, param)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	return &res, nil
}

func (db *DB) GetWaveRecordByCallID(callid string, projectNames []string) (*models.ASR, error) {
	res := models.ASR{}
	param := "{" + strings.Join(projectNames, ",") + "}"
	err := db.DB.Get(&res, "SELECT * FROM asrresults WHERE call_id=$1 AND project_id = ANY($2);", callid, param)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	return &res, nil
}

func (db *DB) GetWaveRecordByProfileNames(projectIDS []string) ([]models.ASR, error) {
	res := []models.ASR{}
	param := "{" + strings.Join(projectIDS, ",") + "}"
	err := db.DB.Select(&res, "SELECT * FROM asrresults WHERE project_id = ANY($1);", param)
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

func (db *DB) GetAllAsrWithFilters(projectIDS []string, filters map[string][]string) ([]models.ASR, error) {
	res := []models.ASR{}
	qq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	ids := "{" + strings.Join(projectIDS, ",") + "}"

	q := qq.Select("id, menu_name, project_id, ani, callid, " +
		"seq, utterance, interpretation, confidence, inputmode,grammaruri,created_on").
		From("asrresults").Where(sq.Expr("project_id = ANY(?)", ids))
	confidence := ""
	confidenceType := ""
	for k, v := range filters {
		if v[0] != "" {
			switch {
			case k == "project_id":
				q = q.Where(sq.Expr("project_id=?", v[0]))
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
