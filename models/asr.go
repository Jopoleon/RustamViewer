package models

import "time"

type ASR struct {
	ID             int       `json:"id" db:"id"`
	ANI            int       `json:"ani" db:"ani"`
	DNIS           int       `json:"dnis" db:"dnis"`
	Profile        string    `json:"profile" db:"profile"`
	Uterance       string    `json:"utterance" db:"utterance"`
	Interpritation string    `json:"interpretation" db:"interpretation"`
	Confidence     float64   `json:"confidence" db:"confidence"`
	WAVRecord      []byte    `json:"wavRecord" db:"waverecord"`
	CreatedOn      time.Time `json:"createdOn" db:"created_on"`
}
