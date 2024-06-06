CREATE TABLE tiktok.comment
(
    id         INT AUTO_INCREMENT PRIMARY KEY,
    parent_id  INT,
    video_id   INT,
    user_id    INT,
    content    VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
