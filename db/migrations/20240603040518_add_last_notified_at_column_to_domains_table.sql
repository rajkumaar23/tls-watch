-- migrate:up
alter table domains add column last_notified_at timestamp not null default (current_timestamp - interval 24 hour);

-- migrate:down
alter table domains drop column last_notified_at;
