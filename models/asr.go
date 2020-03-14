package models

import "time"

//create table asrresults
//(
//	id serial not null
//		constraint asrresults_pkey
//			primary key,
//	external_id integer,
//	menu_name varchar(128),
//	project_id varchar(64),
//	ani varchar(128),
//	callid varchar(128),
//	seq integer,
//	utterance varchar(256),
//	interpretation varchar(256),
//	confidence varchar(128),
//	inputmode varchar(16),
//	grammaruri varchar(128),
//	waverecord bytea,
//	created_on timestamp default CURRENT_TIMESTAMP
//);
//
//alter table asrresults owner to postgres;
//type ASR struct {
//	ID             int       `json:"id" db:"id"`
//	ANI            int       `json:"ani" db:"ani"`
//	DNIS           int       `json:"dnis" db:"dnis"`
//	Profile        string    `json:"profile" db:"profile"`
//	Uterance       string    `json:"utterance" db:"utterance"`
//	Interpritation string    `json:"interpretation" db:"interpretation"`
//	Confidence     float64   `json:"confidence" db:"confidence"`
//	WAVRecord      []byte    `json:"wavRecord" db:"waverecord"`
//	CreatedOn      time.Time `json:"createdOn" db:"created_on"`
//}
type ASR struct {
	ID             int       `json:"id" db:"id"`
	ExternalID     int       `json:"external_id" db:"external_id"`
	MenuName       string    `json:"menu_name" db:"menu_name"`
	ProjectID      string    `json:"project_id" db:"project_id"`
	Ani            int       `json:"ani" db:"ani"`
	CallID         string    `json:"callid" db:"callid"`
	Seq            int       `json:"seq" db:"seq"`
	Utterance      string    `json:"utterance" db:"utterance"`
	Interpretation string    `json:"interpretation" db:"interpretation"`
	Confidence     float64   `json:"confidence" db:"confidence"`
	Inputmode      string    `json:"inputmode" db:"inputmode"`
	Grammaruri     string    `json:"grammaruri" db:"grammaruri"`
	Waverecord     []byte    `json:"waverecord" db:"waverecord"`
	CreatedOn      time.Time `json:"created_on" db:"created_on"`
}
