CREATE TABLE tiktok.user
(
    `id`         INT AUTO_INCREMENT PRIMARY KEY,
    `username`   varchar(255),
    `password`   varchar(255),
    `avatar`     varchar(255),
    `otp_secret` varchar(255),
    `created_at` timestamp default CURRENT_TIMESTAMP,
    `updated_at` timestamp default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` timestamp
) engine = InnoDB
  default charset = utf8mb4;