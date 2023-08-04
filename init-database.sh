#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    create table if not exists public.poll
    (
        id    uuid default gen_random_uuid() not null
            constraint poll_pk
                primary key,
        title text                           not null
    );


    create table if not exists public.question
    (
        id      uuid    default gen_random_uuid() not null
            constraint question_pk
                primary key,
        poll_id uuid                              not null,
        title   text                              not null,
        vote    integer default 0                 not null
    );

EOSQL