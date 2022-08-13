CREATE TABLE todo (
    id TEXT NOT NULL UNIQUE,
    title character varying(255) NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp with time zone DEFAULT NOW(),
    deleted_at timestamp with time zone
);

ALTER TABLE
    todo
ADD
    PRIMARY KEY (id);