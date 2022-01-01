create table if not exists users (
  user_uuid uuid not null default gen_random_uuid(),
  first_name varchar not null,
  last_name varchar not null,
  email varchar not null unique,
  encripted_password  varchar not null,
  created_at timestamp not null DEFAULT current_timestamp,
  updated_at timestamp not null DEFAULT current_timestamp,
  constraint user_pkey primary key (user_uuid)
);