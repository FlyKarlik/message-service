CREATE TABLE IF NOT EXISTS message_status (
    code INT PRIMARY KEY,
    description VARCHAR(50)
);

INSERT INTO message_status VALUES (-1, 'Unprocessed'), (0, 'Processed');
