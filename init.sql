create table if not exists tags(
    id int generated always as identity primary key,
    name text not null unique,
    created_at timestamptz not null default current_timestamp
);