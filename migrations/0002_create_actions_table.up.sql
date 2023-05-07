CREATE TABLE IF NOT EXISTS actions (
    id              uuid PRIMARY KEY,
    name            varchar(500) NOT NULL,
    description     text NULL,
    done            boolean NOT NULL,
    created         timestamp with time zone NULL,
    updated         timestamp with time zone NULL
);