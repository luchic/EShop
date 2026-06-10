create table if not exists users (
    id serial primary key,
    first_name text not null,
    second_name text not null,
    email text not null unique,
    created_at timestamp not null default now()
);