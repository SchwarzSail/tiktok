CREATE DATABASE IF NOT EXISTS Social;
CREATE TABLE Social.follow
(
    id          INT AUTO_INCREMENT PRIMARY KEY,
    user_id     INT,
    follower_id INT,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at  TIMESTAMP DEFAULT NULL
) engine = InnoDB
  default charset = utf8mb4;
CREATE INDEX idx_user_id ON Social.follow(user_id);
CREATE INDEX idx_follower_id ON Social.follow(follower_id);