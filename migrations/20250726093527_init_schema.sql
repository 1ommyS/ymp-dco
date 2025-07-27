-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE SCHEMA dco;

CREATE TABLE dco.segments
(
    group_id uuid    NOT NULL,
    id       uuid    NOT NULL,
    title    text    not null,
    p        integer not null,
    PRIMARY KEY (id)
);

CREATE TABLE dco.clients_segments
(
    client_id  uuid NOT NULL,
    segment_id uuid NOT NULL,
    PRIMARY KEY (client_id, segment_id)
);

CREATE TABLE dco.groups
(
    id    uuid NOT NULL,
    title text NOT NULL,
    PRIMARY KEY (id)
);
-- Indexes
CREATE INDEX groups_idx_title ON dco.groups using hash (title);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
