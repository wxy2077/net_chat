-- +migrate Up
-- 用户表
CREATE TABLE IF NOT EXISTS `users` (
     `id` bigint(20) NOT NULL AUTO_INCREMENT,
     `account` varchar(191) DEFAULT '' COMMENT '账号',
     `password` varchar(191) DEFAULT NULL COMMENT '密码',
     `username` varchar(191) DEFAULT NULL COMMENT '昵称',
     `phone` varchar(16) DEFAULT NULL COMMENT '手机号',
     `avatar` varchar(191) DEFAULT NULL COMMENT '头像',
     `email` varchar(191) DEFAULT NULL COMMENT '头像',
     `created_at` datetime DEFAULT NULL,
     `updated_at` datetime DEFAULT NULL,
     `deleted_at` datetime DEFAULT NULL,
     PRIMARY KEY (`id`),
     KEY `idx_users_deleted_at` (`deleted_at`),
     KEY `idx_account` (`account`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `friends` (
   `id` bigint(20) NOT NULL AUTO_INCREMENT,
   `user_id` bigint(20) DEFAULT NULL,
   `friend_user_id` bigint(20) DEFAULT NULL,
   `status` tinyint(4) DEFAULT '0' COMMENT '申请状态 1:申请中 2:同意 3:拒绝',
   `created_at` datetime DEFAULT NULL,
   `updated_at` datetime DEFAULT NULL,
   `deleted_at` datetime DEFAULT NULL,
   PRIMARY KEY (`id`),
   KEY `idx_friends_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- # password 12345
INSERT INTO `chat`.`users`(`id`, `account`, `password`, `username`, `phone`, `avatar`, `email`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 'wukong', '$2a$10$YXWtf0mwdh1yEC3fH22BieiItD/o02bVIxbPhY0pIg9d4i0eTswIu', '悟空', NULL, 'https://avatars.githubusercontent.com/u/33140097', NULL, '2024-06-12 17:09:13', NULL, NULL);
INSERT INTO `chat`.`users`(`id`, `account`, `password`, `username`, `phone`, `avatar`, `email`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, 'peppa', '$2a$10$YXWtf0mwdh1yEC3fH22BieiItD/o02bVIxbPhY0pIg9d4i0eTswIu', 'Peppa', NULL, 'https://avatars.githubusercontent.com/u/12712884', NULL, '2024-06-12 17:09:16', NULL, NULL);
INSERT INTO `chat`.`users`(`id`, `account`, `password`, `username`, `phone`, `avatar`, `email`, `created_at`, `updated_at`, `deleted_at`) VALUES (3, 'qi', '$2a$10$YXWtf0mwdh1yEC3fH22BieiItD/o02bVIxbPhY0pIg9d4i0eTswIu', '小七', NULL, 'https://avatars.githubusercontent.com/u/37442408', NULL, '2024-06-12 17:09:21', NULL, NULL);
INSERT INTO `chat`.`users`(`id`, `account`, `password`, `username`, `phone`, `avatar`, `email`, `created_at`, `updated_at`, `deleted_at`) VALUES (4, 'feifei', '$2a$10$YXWtf0mwdh1yEC3fH22BieiItD/o02bVIxbPhY0pIg9d4i0eTswIu', '飞飞', NULL, 'https://avatars.githubusercontent.com/u/101162076', NULL, '2024-06-12 17:09:25', NULL, NULL);
INSERT INTO `chat`.`users`(`id`, `account`, `password`, `username`, `phone`, `avatar`, `email`, `created_at`, `updated_at`, `deleted_at`) VALUES (5, 'jay', '$2a$10$YXWtf0mwdh1yEC3fH22BieiItD/o02bVIxbPhY0pIg9d4i0eTswIu', 'Jay', NULL, 'https://avatars.githubusercontent.com/u/54275016', NULL, '2024-06-12 17:09:29', NULL, NULL);
INSERT INTO `chat`.`users`(`id`, `account`, `password`, `username`, `phone`, `avatar`, `email`, `created_at`, `updated_at`, `deleted_at`) VALUES (6, 'lufei', '$2a$10$YXWtf0mwdh1yEC3fH22BieiItD/o02bVIxbPhY0pIg9d4i0eTswIu', 'A路飞', NULL, 'https://avatars.githubusercontent.com/u/18071885', NULL, '2024-06-12 17:09:30', NULL, NULL);

INSERT INTO `chat`.`friends`(`id`, `user_id`, `friend_user_id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 1, 2, 1, '2024-06-12 16:38:16', NULL, NULL);
INSERT INTO `chat`.`friends`(`id`, `user_id`, `friend_user_id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, 1, 3, 1, '2024-06-12 16:38:20', NULL, NULL);
INSERT INTO `chat`.`friends`(`id`, `user_id`, `friend_user_id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (3, 1, 4, 1, '2024-06-13 18:44:35', NULL, NULL);
INSERT INTO `chat`.`friends`(`id`, `user_id`, `friend_user_id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (4, 1, 5, 1, '2024-06-13 18:19:57', NULL, NULL);
INSERT INTO `chat`.`friends`(`id`, `user_id`, `friend_user_id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (5, 1, 6, 1, '2024-06-13 18:20:09', NULL, NULL);
INSERT INTO `chat`.`friends`(`id`, `user_id`, `friend_user_id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (6, 2, 3, 1, '2024-06-16 18:57:35', NULL, NULL);
INSERT INTO `chat`.`friends`(`id`, `user_id`, `friend_user_id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (7, 2, 4, 1, '2024-06-16 18:57:43', NULL, NULL);
INSERT INTO `chat`.`friends`(`id`, `user_id`, `friend_user_id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (8, 2, 5, 1, '2024-06-16 18:57:50', NULL, NULL);
INSERT INTO `chat`.`friends`(`id`, `user_id`, `friend_user_id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (9, 2, 6, 1, '2024-06-16 18:57:58', NULL, NULL);
INSERT INTO `chat`.`friends`(`id`, `user_id`, `friend_user_id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (10, 3, 4, 1, '2024-06-16 18:58:07', NULL, NULL);
INSERT INTO `chat`.`friends`(`id`, `user_id`, `friend_user_id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (11, 3, 5, 1, '2024-06-16 18:58:15', NULL, NULL);
INSERT INTO `chat`.`friends`(`id`, `user_id`, `friend_user_id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (12, 3, 6, 1, '2024-06-16 18:58:22', NULL, NULL);
INSERT INTO `chat`.`friends`(`id`, `user_id`, `friend_user_id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (13, 4, 5, 1, '2024-06-16 18:58:32', NULL, NULL);
INSERT INTO `chat`.`friends`(`id`, `user_id`, `friend_user_id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (14, 4, 6, 1, '2024-06-16 18:58:42', NULL, NULL);
INSERT INTO `chat`.`friends`(`id`, `user_id`, `friend_user_id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (15, 5, 6, 1, '2024-06-16 18:58:52', NULL, NULL);

-- +migrate Down
DROP TABLE `users`;
DROP TABLE `friends`;