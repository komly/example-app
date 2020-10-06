-- +goose Up
-- SQL in this section is executed when the migration is applied.
create table users
(
    id bigserial not null
        constraint users_pkey
            primary key,
    name text not null,
    created_at timestamp with time zone default now() not null
);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
drop table users;