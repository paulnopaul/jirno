create table Users
(
    id       integer,
    name     text,
    nickname text,
    email    text,
    password text
);

create table Tasks
(
    id             varchar(40),
    uid            integer,
    pid            varchar(40),
    title          text,
    description    text,
    additional     text,
    is_completed   integer,
    created_date   integer,
    completed_date integer,
    date_to        integer
);


create table Projects
(
    id             varchar(40),
    title          text,
    description    text,
    additional     text,
    is_completed   integer,
    parent_pid     varchar(40),
    created_date   integer,
    completed_date integer,
    date_to        integer
);

create table ProjectUsers
(
    pid varchar(40),
    uid integer
);

-- LOCAL STORAGE --
create table LocalStorage (
    field varchar(40),
    value text
);

