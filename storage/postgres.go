package storage

import (
	"database/sql"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"

	"golang.org/x/crypto/bcrypt"

	"github.com/Jopoleon/rustamViewer/models"

	"github.com/Jopoleon/rustamViewer/config"
	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	//sq "github.com/Masterminds/squirrel"
)

type DB struct {
	Logger *logrus.Logger
	DB     *sqlx.DB
}

type Storager interface {
}

func NewPostgres(cfg config.Config, logger *logrus.Logger) (*DB, error) {
	str := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName)
	db, err := sqlx.Connect("postgres", str)
	if err != nil {

		logger.Errorf("could not establish connection to ", str, err)
		return nil, errors.WithStack(err)
	}
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)
	return &DB{DB: db, Logger: logger}, nil
}

func (db *DB) CreateUser() (*models.User, error) {
	login := uuid.NewV4()

	password := RandomString(10)
	bpas, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return &models.User{}, errors.WithStack(err)
	}
	_, err = db.DB.Exec("INSERT INTO users (login, email, password) VALUES ($1,$2,$3)",
		login.String(), "", string(bpas))
	if err != nil {
		return &models.User{}, errors.WithStack(err)
	}
	return &models.User{
		Login:    login.String(),
		Password: password,
	}, nil
}

func (db *DB) GetUserByID(id int) (*models.User, error) {
	res := []models.User{}
	err := db.DB.Select(&res, "SELECT * FROM users WHERE id=$1;", id)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if len(res) < 1 {
		return nil, sql.ErrNoRows
	}

	return &res[0], nil
}

func (db *DB) GetUserByEmailOrLogin(email string) (*models.User, error) {
	res := []models.User{}
	err := db.DB.Select(&res, "SELECT * FROM users WHERE email=$1 OR login=$1;", email)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if len(res) < 1 {
		return nil, sql.ErrNoRows
	}

	return &res[0], nil
}

func (db *DB) GetWaveRecordByID(id int) (*models.ASR, error) {
	res := []models.ASR{}
	err := db.DB.Select(&res, "SELECT * FROM asrresults WHERE id=$1;", id)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &res[0], nil
}

func (db *DB) GetWaveRecord(id int) (*models.ASR, error) {
	res := []models.ASR{}
	err := db.DB.Select(&res, "SELECT * FROM asrresults WHERE id=$1;", id)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &res[0], nil
}

func (db *DB) GetWaveRecordByProfileName(profileName string) ([]models.ASR, error) {
	res := []models.ASR{}
	err := db.DB.Select(&res, "SELECT * FROM asrresults WHERE profile=$1 LIMIT 100;", profileName)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return res, nil
}

func (db *DB) GetWaveRecordByFilters(profileName string) ([]models.ASR, error) {
	res := []models.ASR{}

	err := db.DB.Select(&res, "SELECT * FROM asrresults WHERE profile=$1;", profileName)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return res, nil
}

func (db *DB) GetAllAsrResults() ([]models.ASR, error) {
	res := []models.ASR{}
	//	order := &models.Order{}
	err := db.DB.Select(&res, "SELECT id,ani,dnis,profile,"+
		"utterance,interpretation,confidence,created_on FROM asrresults;")
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return res, nil
}

func (db *DB) SetUserSession(user_id int, session string) error {
	_, err := db.DB.Exec("INSERT INTO users_sessions (user_id, session_token, created_at, updated_at) "+
		"VALUES ($1,$2,$3,$4 ) ON CONFLICT (user_id) DO UPDATE SET session_token=excluded.session_token;",
		user_id, session, time.Now(), time.Now())
	//DO UPDATE SET session_token=$2;
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (db *DB) GetUserSession(user_id int) (string, error) {
	var session string
	err := db.DB.Get(&session, "SELECT (session_token) FROM users_session WHERE user_id = $1;", user_id)
	if err != nil {
		return session, errors.WithStack(err)
	}
	return session, nil
}
