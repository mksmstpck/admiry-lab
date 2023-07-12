CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table if not exists "permissions" (
    id uuid default uuid_generate_v1() unique not null,
    name varchar(255) not null unique,
    allowed_to varchar(255)[] default null,
    primary key (id)
);