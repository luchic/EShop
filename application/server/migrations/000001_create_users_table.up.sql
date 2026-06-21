create table if not exists users (
    id serial primary key,
    first_name text not null,
    second_name text not null,
    email text not null unique,
    role text not null,
    password text not null,
    created_at timestamp not null default now()
);

create table if not exists products (
    id serial primary key,
    name text not null,
    description text,
    price numeric(10, 2) not null,
    stock integer default 0,
    image_url text,
    created_at timestamp default now()  
);