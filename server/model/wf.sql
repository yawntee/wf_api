create table user
(
    id       int auto_increment primary key,
    username varchar(255) not null,
    password varchar(255) not null
);

create table account
(
    id   bigint unsigned auto_increment primary key,
    user int      not null,
    name tinytext not null,
    data blob     not null
);
