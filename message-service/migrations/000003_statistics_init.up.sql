CREATE TABLE IF NOT EXISTS statistics (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    processed_count INT NOT NULL,
    last_processed_message UUID REFERENCES message(id),
    last_update TIMESTAMPTZ
);

CREATE OR REPLACE FUNCTION update_statistics()
RETURNS TRIGGER AS $$
BEGIN
    IF EXISTS (SELECT 1 FROM statistics) THEN
        UPDATE statistics
        SET processed_count = (SELECT COUNT(id) FROM message WHERE status=0),
            last_processed_message = (SELECT id FROM message WHERE status=0 ORDER BY processed_at DESC LIMIT 1),
            last_update = CURRENT_TIMESTAMP
        WHERE id = (SELECT id FROM statistics LIMIT 1); 
    ELSE
        INSERT INTO statistics (processed_count, last_processed_message, last_update)
        VALUES ((SELECT COUNT(id) FROM message WHERE status=0),(SELECT id FROM message WHERE status=0 ORDER BY processed_at DESC LIMIT 1), CURRENT_TIMESTAMP);
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_statistics
AFTER UPDATE OF status ON message
FOR EACH ROW
EXECUTE FUNCTION update_statistics();

