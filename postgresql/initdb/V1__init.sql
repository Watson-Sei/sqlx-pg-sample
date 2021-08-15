create table if not exists articles(
    id int generated always as identity primary key,
    title text not null,
    body text not null,
    user_id int not null,
    created_at timestamptz not null default current_timestamp
);

create table if not exists tags(
    id int generated always as identity primary key,
    name text not null unique,
    created_at timestamptz not null default current_timestamp
);

create table if not exists articles_tags(
    id int generated always as identity primary key,
    article_id int not null,
    tag_id int not null,
    foreign key (article_id) references articles(id),
    foreign key (tag_id) references tags(id)
);