-- migrate:up
create table domains (
  id serial primary key,
  user_id int not null,
  domain varchar(255) not null,
  created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp on update current_timestamp,
  constraint uc_user_domains unique (user_id, domain)
);
create index idx_user_domains on domains (user_id, domain);

-- migrate:down
drop index idx_user_domains on domains;
drop table domains;
