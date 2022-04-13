CREATE TABLE IF NOT EXISTS "user"
(
    id         uuid                     not null
        constraint user_pkey
            primary key,
    email      varchar                  not null
        constraint user_email_unq
            unique,
    password   varchar                  not null,
    first_name varchar                  not null,
    last_name  varchar                  not null,
    nickname   varchar                  not null,
    country    varchar                  not null,
    created_at timestamp with time zone not null,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);

