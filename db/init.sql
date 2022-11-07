create table posts
(
    id         varchar(36)             not null,
    title      varchar(255)            not null,
    content    text                    not null,
    created_at timestamp default now() not null,
    updated_at timestamp default now() null on update now(),
    deleted_at timestamp default null  null,
    constraint posts_pk
        primary key (id)
);

