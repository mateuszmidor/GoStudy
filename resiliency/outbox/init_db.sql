create table if not exists outbox
(
    id             uuid                  not null,
    event_name     text                  not null,
    event_data     jsonb default '{}'::jsonb not null,
    occurred_at    timestamp with time zone default now() not null,
    processed_at   timestamp with time zone,
    fail_count     integer default 0     not null,
    failed         boolean default false not null,
    failure_reason text,
    trace_id       text,
    constraint outbox_pk primary key (id)
);