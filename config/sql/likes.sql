CREATE TABLE tiktok.likes
(
    id         INT AUTO_INCREMENT PRIMARY KEY,
    user_id    INT,
    video_id   INT,
    comment_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
) engine = InnoDB
  default charset = utf8mb4;

