create table users
(
    id serial not null,
    login varchar(36) not null,
    email varchar(60) not null,
    password varchar(60) not null,
    is_admin boolean default false not null,
    first_name varchar(36),
    second_name varchar(36),
    company_name varchar(36),
    company_id integer
);

alter table users owner to genesys;

create unique index users_email_uindex
    on users (email);

create unique index users_id_uindex
    on users (id);

create unique index users_login_uindex
    on users (login);

create table users_sessions
(
    id serial not null
        constraint users_sessions_pk
            primary key,
    user_id integer not null,
    session_token text not null,
    created_at timestamp default timenow() not null,
    updated_at timestamp not null
);

alter table users_sessions owner to genesys;

create unique index users_sessions_id_uindex
    on users_sessions (id);

create unique index users_sessions_session_token_uindex
    on users_sessions (session_token);

create unique index users_sessions_user_id_uindex
    on users_sessions (user_id);

create table users_projects
(
    id serial not null
        constraint users_apps_pk
            primary key,
    user_id integer default 0 not null,
    project_id varchar(36)
);

alter table users_projects owner to genesys;

create table companies
(
    id serial not null
        constraint companies_pk
            primary key,
    name varchar(36)
);

alter table companies owner to genesys;

create unique index companies_id_uindex
    on companies (id);

create table project_companies
(
    id serial not null
        constraint apps_to_companies_pk
            primary key,
    project_id varchar(64),
    company_id integer
);

alter table project_companies owner to genesys;

create unique index apps_to_companies_id_uindex
    on project_companies (id);

create table projects
(
    id serial not null
        constraint applications_pk
            primary key,
    project_id varchar(64),
    "desc" text
);

alter table projects owner to genesys;

create unique index applications_column_1_uindex
    on projects (id);

create unique index projects_project_id_uindex
    on projects (project_id);

