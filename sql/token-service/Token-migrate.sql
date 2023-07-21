CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table if not exists "tokens" (
    token uuid not null,
    expires_at integer not null,
    user_id uuid not null,
    primary key (token)

)