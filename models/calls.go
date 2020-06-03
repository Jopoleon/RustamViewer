package models

import "fmt"

const FileNameTimeFormat = "2006-01-02_15-04-05"

type Calls struct {
	ID              int     `json:"id" db:"id"`
	InteractionID   int     `json:"interaction_id" db:"interaction_id"`
	SourceAddress   *string `json:"source_address" db:"source_address"`
	TargetAddress   *string `json:"target_address" db:"target_address"`
	InteractionType *string `json:"interaction_type" db:"interaction_type"`
	MediaType       *string `json:"media_type" db:"media_type"`
	StartTime       MyTime  `json:"start_time" db:"start_time"`
	EndTime         MyTime  `json:"end_time" db:"end_time"`
	ProjectID       *int    `json:"project_id" db:"project_id"`
	CustomerData    *string `json:"customer_data" db:"customer_data"`
	CallID          *string `json:"callid" db:"callid"`
	ProfileName     *string `json:"profilename" db:"profilename"`
	RecordingFileID *string `json:"recording_file_id" db:"recording_file_id"`
	EndTimeOriginal MyTime  `json:"end_time_original" db:"end_time_original"`
	UpdateFlag      *int    `json:"updated_flag" db:"updated_flag"`
	RecordTrans     *string `json:"recordtrans" db:"recordtrans"`
	CreateAuditKey  *int    `json:"create_audit_key" db:"create_audit_key"`
	CreatedOn       MyTime  `json:"created_on" db:"created_on"`
}

func (c *Calls) ToFileName() string {
	//start_time _ project_id _ interaction_type _ source_address _ target_address _ callid
	//2020-03-20_07-49-01_4953080492_Inbound_79167013970_4953080492_3bae508b-2c43-4142-8a9b-899255b4da9f
	if c.RecordTrans != nil &&
		c.ProjectID != nil &&
		c.InteractionType != nil &&
		c.SourceAddress != nil &&
		c.TargetAddress != nil &&
		c.CallID != nil {
		return fmt.Sprintf("%s_%s_%d_%s_%s_%s_%s",
			c.StartTime.Time.Format(FileNameTimeFormat),
			*c.RecordTrans,
			*c.ProjectID,
			*c.InteractionType,
			*c.SourceAddress,
			*c.TargetAddress,
			*c.CallID)
	}
	return ""
}

type CallsOutbound struct {
	ID                    int     `json:"id" db:"id"`
	ContactAttemptFactKey int     `json:"contact_attempt_fact_key" db:"contact_attempt_fact_key"`
	ContactInfo           *string `json:"contact_info" db:"contact_info"`
	MediaType             *string `json:"media_type" db:"media_type"`
	DialingMode           *string `json:"dialing_mode" db:"dialing_mode"`
	Campaign              *string `json:"campaing" db:"campaing"`
	CallResult            *string `json:"call_result" db:"call_result"`
	RecordType            *string `json:"record_type" db:"record_type"`
	RecordStatus          *string `json:"record_status" db:"record_status"`
	CallingList           *string `json:"calling_list" db:"calling_list"`
	ContactInfoType       *string `json:"contact_info_type" db:"contact_info_type"`
	TimeZone              *string `json:"time_zone" db:"time_zone"`
	Callid                *string `json:"callid" db:"callid"`
	StartTime             MyTime  `json:"start_time" db:"start_time"`
	EndTime               MyTime  `json:"end_time" db:"end_time"`
	RecordID              *int    `json:"record_id" db:"record_id"`
	ChainID               *int    `json:"chain_id" db:"chain_id"`
	ChainN                *int    `json:"chain_n" db:"chain_n"`
	Attempt               *int    `json:"attempt" db:"attempt"`
	DailyFrom             MyTime  `json:"daily_from" db:"daily_from"`
	DailyTill             MyTime  `json:"daily_till" db:"daily_till"`
	DialSchedTime         MyTime  `json:"dial_sched_time" db:"dial_sched_time"`
	ProjectID             *string `json:"project_id" db:"project_id"`
	CustomerData          *string `json:"customer_data" db:"customer_data"`
	CreatedOn             MyTime  `json:"created_on" db:"created_on"`
	CreateAuditKey        *int    `json:"create_audit_key" db:"create_audit_key"`
}
