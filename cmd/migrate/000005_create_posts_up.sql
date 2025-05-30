
CREATE TABLE posts (
    id VARCHAR(255) PRIMARY KEY  DEFAULT (UUID()),
    author_id VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (author_id) REFERENCES users(id) ON DELETE CASCADE
);