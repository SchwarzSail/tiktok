CREATE DATABASE IF NOT EXISTS Video;
CREATE TABLE Video.video
(
    id            INT AUTO_INCREMENT PRIMARY KEY,
    user_id       INT,
    video_url     VARCHAR(255),
    cover_url     VARCHAR(255),
    title         VARCHAR(255),
    `description` VARCHAR(255),
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at    TIMESTAMP DEFAULT NULL
) engine = InnoDB
  default charset = utf8mb4;
CREATE INDEX idx_created_at ON Video.video (created_at DESC);
CREATE INDEX idx_user_id ON Video.video (user_id);
