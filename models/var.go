package models

//create table var
//(
//	id serial not null
//		constraint var_pkey
//			primary key,
//	external_id integer,
//	menu_name varchar(128),
//	project_id varchar(64),
//	ani varchar(128),
//	callid varchar(128),
//	seq integer,
//	action_status varchar(256),
//	action_description varchar(256),
//	enter_menu_time timestamp,
//	leave_menu_time timestamp,
//	action_time timestamp,
//	created_on timestamp default CURRENT_TIMESTAMP
//);
//
type VAR struct {
	ID                int     `json:"id" db:"id"`
	ExternalID        *string `json:"external_id" db:"external_id"`
	MenuName          *string `json:"menu_name" db:"menu_name"`
	ProjectID         *string `json:"project_id" db:"project_id"`
	ANI               *string `json:"ani" db:"ani"`
	CallID            *string `json:"callid" db:"callid"`
	Seq               *string `json:"seq" db:"seq"`
	ActionStatus      *string `json:"action_status" db:"action_status"`
	ActionDescription *string `json:"action_description" db:"action_description"`
	EnterMenuName     MyTime  `json:"enter_menu_time" db:"enter_menu_time"`
	LeaveMenuTime     MyTime  `json:"leave_menu_time" db:"leave_menu_time"`
	ActionTime        MyTime  `json:"action_time" db:"action_time"`
	CreatedOn         MyTime  `json:"created_on" db:"created_on"`
}
