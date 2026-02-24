-- +goose Up
create table users
(
    id         serial primary key,
    username   varchar   not null,
    email      varchar   not null,
    password   varchar   not null,
    created_at timestamp not null default now(),
    updated_at timestamp,
    CONSTRAINT users_email_unique UNIQUE (email)
);

-- +goose Down
drop table users;