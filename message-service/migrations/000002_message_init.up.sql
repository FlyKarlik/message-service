CREATE TABLE IF NOT EXISTS message (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    content TEXT NOT NULL,
    status INT REFERENCES message_status(code) DEFAULT -1,
    created_at TIMESTAMPTZ DEFAULT now(),
    processed_at TIMESTAMPTZ
);
