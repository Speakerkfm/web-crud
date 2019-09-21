create table student
(
  id      int auto_increment,
  name    varchar(255) not null,
  surname varchar(255) not null,
  constraint student_id_uindex
  unique (id)
);

alter table student
  add primary key (id);