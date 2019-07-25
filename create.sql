CREATE DATABASE IF NOT EXISTS `cloudmusic`;
CREATE TABLE IF NOT EXISTS `cloudmusic`.`hot_comments` (
  `id` INT UNSIGNED AUTO_INCREMENT,
  `content` TEXT NOT NULL,
  `nickname` VARCHAR(50) NOT NULL,
  `liked_count` INT UNSIGNED NOT NULL,
   PRIMARY KEY ( `id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;