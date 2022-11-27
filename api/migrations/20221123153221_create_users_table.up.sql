create table if not exists users (
    id varchar(64) not null primary key,
    screen_name varchar(255) not null unique,
    name varchar(255) not null,
    email varchar(255) not null unique,
    password varchar(255) not null,
    description text,
    created_at timestamp,
    updated_at timestamp,
    last_logined_at timestamp
);
