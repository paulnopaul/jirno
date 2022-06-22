CREATE ROLE ADMIN WITH LOGIN SUPERUSER PASSWORD 'dbms';

CREATE DATABASE jirno;

\c jirno

CREATE TABLE Users
(
    id         INTEGER PRIMARY KEY,
    first_name VARCHAR(40) NOT NULL,
    last_name  VARCHAR(40) NOT NULL,
    nickname   VARCHAR(40) NOT NULL UNIQUE,
    email      VARCHAR(40) NOT NULL UNIQUE,
    password   VARCHAR(40) NOT NULL UNIQUE,
    phone      VARCHAR(20) UNIQUE
);

CREATE TABLE Tasks
(
    id             UUID PRIMARY KEY,
    creator        INTEGER      NOT NULL,
    executor       INTEGER,
    status         INTEGER      NOT NULL,
    title          VARCHAR(256) NOT NULL,
    description    TEXT,
    is_completed   BOOLEAN      NOT NULL,
    created_date   DATE         NOT NULL,
    start_date     DATE,
    date_to        DATE,
    completed_date DATE
);


CREATE TABLE TaskStatuses
(
    id    INTEGER PRIMARY KEY,
    value VARCHAR(40) UNIQUE
);

ALTER TABLE Tasks
    ADD CONSTRAINT tasks_creator_fk FOREIGN KEY (creator) REFERENCES Users (id),
    ADD CONSTRAINT tasks_executor_fk FOREIGN KEY (executor) REFERENCES Users (id),
    ADD CONSTRAINT tasks_status_fk FOREIGN KEY (executor) REFERENCES TaskStatuses (id);

CREATE TABLE Projects
(
    id           UUID PRIMARY KEY,
    creator      INTEGER NOT NULL,
    title        TEXT    NOT NULL,
    description  TEXT,
    status       INTEGER NOT NULL,
    parent_pid   UUID,
    created_date DATE
);

create table ProjectStatuses
(
    id    INTEGER PRIMARY KEY,
    value VARCHAR(40) UNIQUE
);


ALTER TABLE Projects
    ADD CONSTRAINT projects_parent_pid_fk FOREIGN KEY (parent_pid) REFERENCES Projects (id),
    ADD CONSTRAINT projects_creator_fk FOREIGN KEY (creator) REFERENCES Users (id),
    ADD CONSTRAINT projects_status_fk FOREIGN KEY (status) REFERENCES ProjectStatuses (id);

CREATE TABLE ProjectUsers
(
    pid UUID,
    uid INTEGER
);

ALTER TABLE ProjectUsers
    ADD CONSTRAINT project_users_pair_unique UNIQUE (pid, uid),
    ADD CONSTRAINT project_users_pid_fk FOREIGN KEY (pid) REFERENCES Projects (id),
    ADD CONSTRAINT project_users_uid_fk FOREIGN KEY (uid) REFERENCES Users (id);


CREATE TABLE ProjectTasks
(
    pid UUID,
    tid UUID
);

ALTER TABLE ProjectTasks
    ADD CONSTRAINT project_tasks_pair_unique UNIQUE (pid, tid),
    ADD CONSTRAINT project_tasks_pid_fk FOREIGN KEY (pid) REFERENCES Projects (id),
    ADD CONSTRAINT project_tasks_uid_fk FOREIGN KEY (tid) REFERENCES Tasks (id);
