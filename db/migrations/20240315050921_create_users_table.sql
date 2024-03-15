-- migrate:up
create table users (
  id serial primary key,
  oidc_subject varchar(255) unique not null,
  name varchar(255),
  picture varchar(255),
  created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp on update current_timestamp
);

-- migrate:down
drop table users;
