create table user
(
    id       int auto_increment primary key,
    username varchar(255) not null,
    password varchar(255) not null
);

create table game_user
(
    id   bigint unsigned auto_increment primary key,
    user int      not null,
    name tinytext not null,
    channel tinyint not null,
    data blob     not null
);