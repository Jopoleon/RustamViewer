
alter table users add column created_on timestamptz;
alter table users add column updated_on timestamptz;
alter table users add column deleted_on timestamptz;
alter table users add column  action_by_id int;
alter table users add column action_by_login varchar(36);

alter table projects add column created_on timestamptz;
alter table projects add column updated_on timestamptz;
alter table projects add column deleted_on timestamptz;
alter table projects add column  action_by_id int;
alter table projects add column action_by_login varchar(36);

alter table project_companies add column created_on timestamptz;
alter table project_companies add column updated_on timestamptz;
alter table project_companies add column deleted_on timestamptz;
alter table project_companies add column  action_by_id int;
alter table project_companies add column action_by_login varchar(36);

alter table user_projects add column created_on timestamptz;
alter table user_projects add column updated_on timestamptz;
alter table user_projects add column deleted_on timestamptz;
alter table user_projects add column  action_by_id int;
alter table user_projects add column action_by_login varchar(36);