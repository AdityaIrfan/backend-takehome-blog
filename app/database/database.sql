CREATE TABLE
    users (
        id VARCHAR(26) PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        email VARCHAR(100) NOT NULL UNIQUE,
        password_hash VARCHAR(255) NOT NULL,
        password_salt VARCHAR(255) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT NULL,
        INDEX idx_email (email)
    );

CREATE TABLE
    posts (
        id VARCHAR(26) PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        content TEXT NOT NULL,
        author_id VARCHAR(26) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT NULL,
        deleted_at TIMESTAMP DEFAULT NULL,
        INDEX idx_author_id (author_id),
        INDEX idx_created_at (created_at),
        FOREIGN KEY (author_id) REFERENCES users (id) ON DELETE CASCADE
    );

CREATE TABLE
    comments (
        id VARCHAR(26) PRIMARY KEY,
        post_id VARCHAR(26) NOT NULL,
        author_id VARCHAR(26) NOT NULL,
        content TEXT NOT NULL,
        parent_id VARCHAR(26) DEFAULT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        deleted_at TIMESTAMP DEFAULT NULL,
        INDEX idx_post_id (post_id),
        INDEX idx_parent_id (parent_id),
        INDEX idx_created_at (created_at),
        FOREIGN KEY (post_id) REFERENCES posts (id) ON DELETE CASCADE,
        FOREIGN KEY (parent_id) REFERENCES comments (id) ON DELETE CASCADE,
        FOREIGN KEY (author_id) REFERENCES users (id) ON DELETE CASCADE
    );