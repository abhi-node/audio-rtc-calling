DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'request_status') THEN
        CREATE TYPE request_status AS ENUM ('accepted', 'outgoing', 'denied');
    END IF;
END
$$;

CREATE TABLE IF NOT EXISTS friend_requests (
    sender_id UUID NOT NULL,
    receiver_id UUID NOT NULL,
    status request_status DEFAULT 'outgoing',
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (sender_id, receiver_id),
    FOREIGN KEY (sender_id) REFERENCES users (id),
    FOREIGN KEY (receiver_id) REFERENCES users (id)
);

CREATE INDEX IF NOT EXISTS idx_receiver_id_friend_requests ON friend_requests (receiver_id);
