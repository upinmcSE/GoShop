-- CREATE TABLE IF NOT EXISTS users (
--     user_id INT AUTO_INCREMENT PRIMARY KEY,
--     uuid CHAR(36) NOT NULL DEFAULT (UUID()),
--     name VARCHAR(50) NOT NULL,
--     email VARCHAR(100) UNIQUE NOT NULL,
--     created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
-- );

-- Create table
CREATE TABLE IF NOT EXISTS users (
    user_id         INT AUTO_INCREMENT PRIMARY KEY,
    user_uuid       CHAR(36) NOT NULL DEFAULT (UUID()) UNIQUE,
    user_email      VARCHAR(150) NOT NULL UNIQUE,
    user_password   VARCHAR(90) NOT NULL,
    user_fullname   VARCHAR(100) NOT NULL,
    user_age        INT CHECK (user_age >= 1 AND user_age <= 150) COMMENT 'User age, must be between 1 and 150',
    user_status     INT NOT NULL DEFAULT 1 CHECK (user_status IN (1, 2, 3)) COMMENT 'User status: 1 - Active, 2 - Inactive, 3 - Banned',
    user_level      INT NOT NULL DEFAULT 1 CHECK (user_level IN (1, 2, 3)) COMMENT 'User level: 1 - Administrator, 2 - Moderator, 3 - Member',
    user_deleted_at TIMESTAMP NULL DEFAULT NULL COMMENT 'Soft delete timestamp: NULL means not deleted',
    user_created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    user_updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- Indexes
CREATE INDEX idx_users_status       ON users(user_status);
CREATE INDEX idx_users_level        ON users(user_level);
CREATE INDEX idx_users_created_at   ON users(user_created_at);
CREATE INDEX idx_user_deleted_at    ON users(user_deleted_at);
CREATE INDEX idx_users_email_status ON users(user_email, user_status);