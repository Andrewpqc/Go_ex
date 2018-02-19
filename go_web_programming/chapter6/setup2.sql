drop table posts cascade if exists
drop table comments if exists

create table posts(
    id serial primary key,
    content text,
    author varchar(225)
);

create table comments(
    id serial primary key,
    content text,
    author varchar(225),
    post_id integer references posts(id)
);