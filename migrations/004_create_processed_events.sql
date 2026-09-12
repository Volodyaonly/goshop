CREATE TABLE IF NOT EXISTS processed_events (
    service TEXT NOT NULL,
    event_id TEXT NOT NULL,
    processed_at TIMESTAMP NOT NULL DEFAULT NOW(),

    PRIMARY KEY (service, event_id)
);