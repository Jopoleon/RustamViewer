package models

import (
	"database/sql"
	"database/sql/driver"
	"strconv"
	"time"
)

const MyTimeFormat = "2006-01-02T15:04:05"

type MyTime sql.NullTime

func (t MyTime) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(sql.NullTime(t).Time.Format(MyTimeFormat))), nil
	//return []byte(sql.NullTime(t).Time.Format(time.RFC3339)), nil
}

func (t MyTime) String() string {
	return sql.NullTime(t).Time.Format(MyTimeFormat)
	//return sql.NullTime(t).Time.Format(time.RFC3339)
}

// Scan implements the Scanner interface.
func (nt *MyTime) Scan(value interface{}) error {
	nt.Time, nt.Valid = value.(time.Time)
	return nil
}

// Value implements the driver Valuer interface.
func (nt MyTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}
