package models

import "time"

//create table calls_inbound
//(
//id serial not null
//constraint calls_inbound_pkey
//primary key,
//interaction_id numeric(19),
//source_address varchar(255),
//target_address varchar(255),
//interaction_type varchar(32),
//media_type varchar(32),
//start_time timestamp,
//end_time timestamp,
//project_id varchar(64),
//customer_data varchar(255),
//callid varchar(128),
//profilename varchar(128),
//created_on timestamp default CURRENT_TIMESTAMP
//);
type Calls struct {
	ID              int       `json:"id" db:"id"`
	InteractionID   int       `json:"interaction_id" db:"interaction_id"`
	SourceAddress   string    `json:"source_address" db:"source_address"`
	TargetAddress   string    `json:"target_address" db:"target_address"`
	InteractionType string    `json:"interaction_type" db:"interaction_type"`
	MediaType       string    `json:"media_type" db:"media_type"`
	StartTime       time.Time `json:"start_time" db:"start_time"`
	EndTime         time.Time `json:"end_time" db:"end_time"`
	ProjectID       int       `json:"project_id" db:"project_id"`
	CustomerData    string    `json:"customer_data" db:"customer_data"`
	CallID          string    `json:"callid" db:"callid"`
	ProfileName     string    `json:"profilename" db:"profilename"`
	CreatedOn       time.Time `json:"created_on" db:"created_on"`
}

//create table calls_outbound
//(
//id serial not null
//contact_attempt_fact_key numeric(19),
//contact_info varchar(128),
//media_type varchar(32),
//dialing_mode varchar(32),
//campaign varchar(64),
//call_result varchar(32),
//record_type varchar(32),
//record_status varchar(32),
//calling_list varchar(64),
//contact_info_type varchar(32),
//time_zone varchar(32),
//callid varchar(128),
//start_time timestamp,
//end_time timestamp,
//record_id integer,
//chain_id integer,
//chain_n integer,
//attempt integer,
//daily_from timestamp,
//daily_till timestamp,
//dial_sched_time timestamp,
//project_id varchar(64),
//customer_data varchar(255),
//created_on timestamp default CURRENT_TIMESTAMP
//);
//
type CallsOutbound struct {
	ID                    int       `json:"id" db:"id"`
	ContactAttemptFactKey int       `json:"contact_attempt_fact_key" db:"contact_attempt_fact_key"`
	ContactInfo           string    `json:"contact_info" db:"contact_info"`
	MediaType             string    `json:"media_type" db:"media_type"`
	DialingMode           string    `json:"dialing_mode" db:"dialing_mode"`
	Campaign              string    `json:"campaing" db:"campaing"`
	CallResult            string    `json:"call_result" db:"call_result"`
	RecordType            string    `json:"record_type" db:"record_type"`
	RecordStatus          string    `json:"record_status" db:"record_status"`
	CallingList           string    `json:"calling_list" db:"calling_list"`
	ContactInfoType       string    `json:"contact_info_type" db:"contact_info_type"`
	TimeZone              string    `json:"time_zone" db:"time_zone"`
	Callid                string    `json:"callid" db:"callid"`
	StartTime             time.Time `json:"start_time" db:"start_time"`
	EndTime               time.Time `json:"end_time" db:"end_time"`
	RecordID              int       `json:"record_id" db:"record_id"`
	ChainID               int       `json:"chain_id" db:"chain_id"`
	ChainN                int       `json:"chain_n" db:"chain_n"`
	Attempt               int       `json:"attempt" db:"attempt"`
	DailyFrom             time.Time `json:"daily_from" db:"daily_from"`
	DailyTill             time.Time `json:"daily_till" db:"daily_till"`
	DialSchedTime         time.Time `json:"dial_sched_time" db:"dial_sched_time"`
	ProjectID             string    `json:"project_id" db:"project_id"`
	CustomerData          string    `json:"customer_data" db:"customer_data"`
	CreatedOn             time.Time `json:"created_on" db:"created_on"`
}
