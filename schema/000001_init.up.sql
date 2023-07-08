CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null,
    email varchar(255) not null unique,
    username varchar(255) not null unique,
    password_hash varchar(255) not null,
    avatar varchar(255),
    bio text
);

CREATE TABLE workspaces
(
    id serial not null unique,
    name varchar(255) not null,
    description text,
    avatar varchar(255)
);

CREATE TABLE workspace_users
(
    id serial not null unique,
    workspace_id int references workspaces(id) on delete cascade not null,
    user_id int references users(id) on delete cascade not null,
    role int not null
);

CREATE TABLE workspace_invites
(
    id serial not null unique,
    workspace_id int references workspaces(id) on delete cascade not null,
    user_id int references users(id) on delete cascade not null
);

CREATE TABLE boards
(
    id serial not null unique,
    workspace_id int references workspaces(id) on delete cascade not null,
    name varchar(255) not null,
    color varchar(10) not null
);

CREATE TABLE tasks
(
    id serial not null unique,
    board_id int references boards(id) on delete cascade not null,
    number int not null,
    name varchar(255) not null,
    description text,
    time_estimated int,
    assignee_id int references users(id) on delete set null,
    created_by int references users(id) on delete set null
);

CREATE TABLE task_attachments
(
    id serial not null unique,
    file varchar(255) not null,
    task_id int references tasks(id) on delete cascade not null
);

CREATE TABLE task_time_tracks
(
    id serial not null unique,
    task_id int references tasks(id) on delete cascade not null,
    user_id int references users(id) on delete cascade not null,
    time_spent int not null
);

CREATE TABLE task_comments
(
    id serial not null unique,
    task_id int references tasks(id) on delete cascade not null,
    user_id int references users(id) on delete cascade not null,
    body text not null
);