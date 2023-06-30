CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table if not exists "users" (
    id uuid default uuid_generate_v1() not null,
    username varchar(255) not null,
    full_name varchar(255) not null,
    email varchar(255) not null,
    password varchar(255) not null,
    primary key (id),
    unique(username, email)
    );