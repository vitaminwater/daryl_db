create database daryl;

\c daryl;

create table primate (
  id serial,
  slug varchar(10),
  username varchar(30)
);

create unique index primate_name_index on primate (slug);
