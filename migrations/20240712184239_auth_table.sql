-- +goose Up
create table User (
    id serial primary key,
    username text not null,
    password text not null,
    role int not null,
    created_at timestamp not null default now(),
    updated_at timestamp
)

-- +goose Down
drop table User;
