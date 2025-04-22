CREATE TABLE IF NOT EXISTS completed_tasks (
    id UUID PRIMARY KEY,
    input TEXT NOT NULL,
    result TEXT,
    error TEXT,
    status TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL
);
