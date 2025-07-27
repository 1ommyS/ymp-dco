-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

alter table dco.groups
    alter column id set default gen_random_uuid();

alter table dco.segments
    alter column id set default gen_random_uuid();

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
