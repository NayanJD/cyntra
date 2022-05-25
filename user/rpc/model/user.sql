CREATE TABLE IF NOT EXISTS role (
    id  text PRIMARY KEY,
    name text NOT NULL UNIQUE,
    created_at  timestamptz NOT NULL,
    updated_at  timestamptz NOT NULL,
    archived_at  timestamptz NOT NULL
);

CREATE TABLE IF NOT EXISTS scope (
    id  text PRIMARY KEY,
    name text NOT NULL UNIQUE,

    created_at  timestamptz NOT NULL,
    updated_at  timestamptz NOT NULL,
    archived_at  timestamptz NOT NULL
);

CREATE TABLE IF NOT EXISTS role_scope (
    id  serial PRIMARY KEY,
    role_id text NOT NULL,
    scope_id text NOT NULL,
    FOREIGN KEY (role_id) REFERENCES role (id),
    FOREIGN KEY (scope_id) REFERENCES scope (id),

    created_at  timestamptz NOT NULL,
    updated_at  timestamptz NOT NULL,
    archived_at  timestamptz NOT NULL
);

CREATE INDEX role_scope_role_id ON role_scope (role_id);
CREATE INDEX role_scope_scope_id ON role_scope (scope_id);


CREATE TABLE IF NOT EXISTS public.user (
    id text PRIMARY KEY,
    username text NOT NULL UNIQUE,
    hashed_password text NOT NULL,
    first_name  text NOT NULL,
    last_name text NOT NULL,
    gender  text NOT NULL,
    dob timestamptz NOT NULL,
    role_id text NOT NULL,

    created_at  timestamptz NOT NULL,
    updated_at  timestamptz NOT NULL,
    archived_at  timestamptz NOT NULL,

    FOREIGN KEY (role_id) REFERENCES role (id)
);

CREATE INDEX user_role_id ON public.user (role_id);