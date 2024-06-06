CREATE DATABASE IF NOT EXISTS Interaction;
CREATE TABLE Interaction.likes
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
#因为sql语句有包含对各自字段对单独查询，所以没有建立联合索引
CREATE INDEX idx_user_id ON Interaction.likes(user_id);
CREATE INDEX idx_video_id ON Interaction.likes(video_id);
CREATE INDEX idx_comment_id ON Interaction.likes(comment_id);
