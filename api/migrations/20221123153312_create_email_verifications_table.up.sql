create table if not exists email_verifications (
    id varchar(64) not null primary key,
    token varchar(255) not null unique,
    email varchar(255) not null,
    user_id varchar(64),
    created_at timestamp,
    updated_at timestamp
);
