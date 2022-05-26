DO $$
DECLARE
    role_id text = uuid_generate_v4()::text;
    scope_id text = uuid_generate_v4()::text;
BEGIN
    insert into role (id, name, created_at, updated_at, archived_at) values (role_id, 'default_user', now(), now(), null);
    insert into scope (id, name, created_at, updated_at, archived_at) values (scope_id, 'product:read', now(), now(), null);
    insert into role_scope(role_id, scope_id, created_at, updated_at, archived_at) values (role_id, scope_id, now(), now(), null);
END $$;