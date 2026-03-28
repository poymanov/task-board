-- +goose Up
create table outbox_events
(
    id           serial primary key,
    entity_type  varchar   not null,
    entity_id    int       not null,
    payload      json      not null,
    created_at   timestamp not null default now(),
    processed_at timestamp
);

-- +goose Down
drop table outbox_events;