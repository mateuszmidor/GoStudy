-- schema for event store
CREATE TABLE IF NOT EXISTS events (
    id             BIGSERIAL    PRIMARY KEY,
    event_id       UUID         NOT NULL,
    stream_id      VARCHAR      NOT NULL,
    stream_position BIGINT      NOT NULL,
    event_type     VARCHAR      NOT NULL,
    payload        BYTEA        NOT NULL,
    metadata       BYTEA,
    occurred_at    TIMESTAMPTZ  NOT NULL,
    UNIQUE (stream_id, stream_position)
);
CREATE INDEX IF NOT EXISTS idx_events_stream_id ON events (stream_id);

-- schema for event bus
CREATE TABLE IF NOT EXISTS event_subscriptions (
    name     VARCHAR PRIMARY KEY,
    position BIGINT  NOT NULL DEFAULT 0
);
CREATE OR REPLACE FUNCTION eventsourcing_notify_events_inserted() RETURNS trigger AS $$
BEGIN
    PERFORM pg_notify('eventsourcing_events_inserted', '');
    RETURN NULL;
END;
$$ LANGUAGE plpgsql;
DROP TRIGGER IF EXISTS eventsourcing_events_notify ON events;
CREATE TRIGGER eventsourcing_events_notify
    AFTER INSERT ON events
    FOR EACH STATEMENT
    EXECUTE FUNCTION eventsourcing_notify_events_inserted();