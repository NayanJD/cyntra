DO $$
DECLARE
    role_id text = uuid_generate_v4()::text;
    product_read_scope_id text = uuid_generate_v4()::text;
    order_create_scope_id text = uuid_generate_v4()::text;
BEGIN
    insert into role (id, name, created_at, updated_at, archived_at) values (role_id, 'default_user', now(), now(), null);
    insert into scope (id, name, created_at, updated_at, archived_at) values (product_read_scope_id, 'product:read', now(), now(), null);
    insert into scope (id, name, created_at, updated_at, archived_at) values (order_create_scope_id, 'order:create', now(), now(), null);
    insert into role_scope(role_id, scope_id, created_at, updated_at, archived_at) values (role_id, product_read_scope_id, now(), now(), null);
    insert into role_scope(role_id, scope_id, created_at, updated_at, archived_at) values (role_id, order_create_scope_id, now(), now(), null);
END $$;


select u.id, u.username, u.first_name, 
    (select string_agg(s.name, ',')  
    from public.user u join 
    role_scope rs on u.role_id = rs.role_id 
    join scope s on rs.scope_id = s.id 
    where u.id = 'dcc2c95d-e8d2-4099-a6cf-bc2637fcc78b') as scopes 
from public.user u  where u.id = 'dcc2c95d-e8d2-4099-a6cf-bc2637fcc78b';