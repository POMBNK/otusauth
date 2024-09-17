-- +goose Up
-- +goose StatementBegin
create table if not exists users
(
    id         serial
        primary key,
    last_name  varchar(255) not null,
    first_name varchar(255) not null,
    email      varchar(255)
        unique,
    phone      varchar(255) not null
        unique,
    username   varchar(255) not null
        unique,
    password   varchar(255)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
Drop table users if exists;
-- +goose StatementEnd
