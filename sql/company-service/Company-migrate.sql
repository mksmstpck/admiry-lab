CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table if not exists "companies" (
    id uuid default uuid_generate_v1() unique not null,
    name varchar(255) not null unique,
    user_i_ds uuid[] default null,
    primary key (id)
    );