-- migrate:up
create table notification_settings (
    id serial primary key,
    user_id int not null,
    enabled boolean default false,
    provider enum('telegram', 'email') not null,
    provider_user_id varchar(255) not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    constraint uc_user_notification_provider unique (user_id, provider)
);

-- migrate:down
drop table notification_settings;
