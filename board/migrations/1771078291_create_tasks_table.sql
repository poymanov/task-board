-- +goose Up
create table tasks
(
    id          serial primary key,
    title       varchar                                       not null,
    description varchar                                       not null,
    assignee    varchar                                       not null,
    position    numeric                                       not null,
    column_id   int REFERENCES columns (id) ON DELETE CASCADE not null,
    created_at  timestamp                                     not null default now(),
    updated_at  timestamp,
    CONSTRAINT tasks_column_position_unique UNIQUE (column_id, position)
);

-- +goose Down
drop table tasks;