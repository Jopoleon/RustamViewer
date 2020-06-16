package postgres

import (
	"strings"

	"github.com/Jopoleon/rustamViewer/models"
	"github.com/pkg/errors"
)

func (db *DB) GetCallsAllByCallID(callID string, projectNames []string) ([]models.Calls, error) {
	res := []models.Calls{}
	param := "{" + strings.Join(projectNames, ",") + "}"
	err := db.DB.Select(&res, "SELECT * FROM calls_all WHERE callid = $1 AND project_id = ANY($2);", callID, param)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	return res, nil
}

func (db *DB) GetCallsAllByID(ID string, projectNames []string) (*models.Calls, error) {
	res := models.Calls{}
	param := "{" + strings.Join(projectNames, ",") + "}"
	err := db.DB.Get(&res, "SELECT * FROM calls_all WHERE id = $1 AND project_id = ANY($2);", ID, param)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	return &res, nil
}

func (db *DB) GetCallsOutboundByCallID(callID string, projectNames []string) ([]models.CallsOutbound, error) {
	res := []models.CallsOutbound{}
	param := "{" + strings.Join(projectNames, ",") + "}"
	err := db.DB.Select(&res, "SELECT * FROM calls_outbound WHERE callid = $1 AND  project_id = ANY($2);", callID, param)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	return res, nil
}

func (db *DB) GetCallsAllProjectIDs(projectIDS []string) ([]models.Calls, error) {
	res := []models.Calls{}
	param := "{" + strings.Join(projectIDS, ",") + "}"
	err := db.DB.Select(&res, "SELECT * FROM calls_all WHERE project_id = ANY($1);", param)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	return res, nil
}

func (db *DB) GetCallsOutboundProjectIDs(projectIDS []string) ([]models.CallsOutbound, error) {
	res := []models.CallsOutbound{}
	param := "{" + strings.Join(projectIDS, ",") + "}"
	err := db.DB.Select(&res, "SELECT * FROM calls_outbound WHERE project_id = ANY($1);", param)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	return res, nil
}

func (db *DB) GetCallsAllAdmin() ([]models.Calls, error) {
	res := []models.Calls{}
	err := db.DB.Select(&res, "SELECT * FROM calls_all;")
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	return res, nil
}
func (db *DB) GetCallsOutboundAdmin() ([]models.CallsOutbound, error) {
	res := []models.CallsOutbound{}
	err := db.DB.Select(&res, "SELECT * FROM calls_outbound;")
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	return res, nil
}

func (db *DB) GetCallsAllByCallIDAdmin(callID string) ([]models.Calls, error) {
	res := []models.Calls{}
	err := db.DB.Select(&res, "SELECT * FROM calls_all WHERE callid = $1;", callID)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	return res, nil
}

func (db *DB) GetCallsOutboundByCallIDAdmin(callID string) ([]models.CallsOutbound, error) {
	res := []models.CallsOutbound{}

	err := db.DB.Select(&res, "SELECT * FROM calls_outbound WHERE callid = $1;", callID)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	return res, nil
}
