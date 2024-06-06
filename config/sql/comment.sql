CREATE DATABASE IF NOT EXISTS Interaction;
CREATE TABLE Interaction.comment
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
CREATE INDEX idx_video_id_created_at_desc ON Interaction.comment(video_id, created_at DESC);
CREATE INDEX  idx_parent_id_created_at_desc ON Interaction.comment(parent_id, created_at DESC);
