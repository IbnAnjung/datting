create database if not exists ks;

use ks;

-- datting.users definition
CREATE TABLE IF NOT EXISTS `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(25) NOT NULL,
  `password` text NOT NULL,
  `fullname` varchar(100) NOT NULL DEFAULT '',
  `age` smallint DEFAULT NULL,
  `gender` enum('L','P') DEFAULT NULL,
  `is_premium` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=183 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- datting.user_swaps definition
CREATE TABLE IF NOT EXISTS `user_swaps` (
  `user_id` bigint unsigned NOT NULL,
  `swapped_user_id` bigint unsigned NOT NULL,
  `pass_count` int DEFAULT '0',
  `like_count` int DEFAULT '0',
  PRIMARY KEY (`user_id`,`swapped_user_id`),
  KEY `swapped_user_id` (`swapped_user_id`),
  CONSTRAINT `user_swaps_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  CONSTRAINT `user_swaps_ibfk_2` FOREIGN KEY (`swapped_user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;