-- +goose Up
create table columns
(
    id         serial primary key,
    name       varchar                                      not null,
    position   numeric                                      not null,
    board_id   int REFERENCES boards (id) ON DELETE CASCADE not null,
    created_at timestamp                                    not null default now(),
    updated_at timestamp,
    CONSTRAINT columns_board_position_unique UNIQUE (board_id, position)
);

-- +goose Down
drop table columns;