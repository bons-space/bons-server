CREATE DATABASE bons;

use bons;

create table if not exists user(
    id int(10) not null,
    name varchar(255),
    password varchar(255),
    primary key (id)
);

