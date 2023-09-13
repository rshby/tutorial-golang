CREATE TABLE `users` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `first_name` varchar(255) NOT NULL,
  `last_name` varchar(255) NOT NULL,
  `identity_id` varchar(255) NOT NULL,
  `gender` varchar(255) NOT NULL,
  `address` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT (current_timestamp)
);

CREATE TABLE `accounts` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `email` varchar(255) UNIQUE NOT NULL,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `otp` varchar(6),
  `expired_otp` datetime,
  `created_at` datetime NOT NULL DEFAULT (current_timestamp),
  `user_id` int NOT NULL
);

CREATE TABLE `contents` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `picture_url` varchar(255) NOT NULL,
  `textfill` varchar(255) NOT NULL,
  `like` int NOT NULL DEFAULT 0,
  `dislike` int NOT NULL DEFAULT 0,
  `average_rating` double NOT NULL DEFAULT 0,
  `created_at` datetime NOT NULL DEFAULT (current_timestamp),
  `account_id` int NOT NULL
);

CREATE TABLE `reviews` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `rating` int NOT NULL,
  `comment` varchar(255),
  `created_at` datetime NOT NULL DEFAULT (current_timestamp),
  `account_id` int NOT NULL,
  `content_id` int NOT NULL
);

CREATE TABLE `likes` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `account_id` int NOT NULL,
  `content_id` int NOT NULL
);

CREATE TABLE `dislikes` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `account_id` int NOT NULL,
  `content_id` int NOT NULL
);

CREATE TABLE `logger` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `ip_address` varchar(255),
  `url_path` varchar(255),
  `method` varchar(255),
  `status_code` int,
  `status` varchar(255),
  `duration` varchar(255),
  `created_at` datetime NOT NULL DEFAULT (current_timestamp)
);

CREATE INDEX `users_index_0` ON `users` (`id`);

CREATE INDEX `users_index_1` ON `users` (`identity_id`);

CREATE INDEX `users_index_2` ON `users` (`first_name`);

CREATE INDEX `users_index_3` ON `users` (`last_name`);

CREATE INDEX `users_index_4` ON `users` (`first_name`, `last_name`);

CREATE INDEX `accounts_index_5` ON `accounts` (`user_id`);

CREATE INDEX `accounts_index_6` ON `accounts` (`email`);

CREATE INDEX `accounts_index_7` ON `accounts` (`username`);

CREATE INDEX `accounts_index_8` ON `accounts` (`username`, `password`);

CREATE INDEX `contents_index_9` ON `contents` (`title`);

CREATE INDEX `contents_index_10` ON `contents` (`id`);

CREATE INDEX `reviews_index_11` ON `reviews` (`account_id`);

CREATE INDEX `reviews_index_12` ON `reviews` (`content_id`);

CREATE INDEX `reviews_index_13` ON `reviews` (`account_id`, `content_id`);

CREATE INDEX `likes_index_14` ON `likes` (`account_id`);

CREATE INDEX `likes_index_15` ON `likes` (`content_id`);

CREATE INDEX `likes_index_16` ON `likes` (`account_id`, `content_id`);

CREATE INDEX `dislikes_index_17` ON `dislikes` (`account_id`);

CREATE INDEX `dislikes_index_18` ON `dislikes` (`content_id`);

CREATE INDEX `dislikes_index_19` ON `dislikes` (`account_id`, `content_id`);

ALTER TABLE `accounts` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `contents` ADD FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`);

ALTER TABLE `reviews` ADD FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`);

ALTER TABLE `reviews` ADD FOREIGN KEY (`content_id`) REFERENCES `contents` (`id`);

ALTER TABLE `likes` ADD FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`);

ALTER TABLE `likes` ADD FOREIGN KEY (`content_id`) REFERENCES `contents` (`id`);

ALTER TABLE `dislikes` ADD FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`);

ALTER TABLE `dislikes` ADD FOREIGN KEY (`content_id`) REFERENCES `contents` (`id`);
