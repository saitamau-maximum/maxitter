CREATE TABLE IF NOT EXISTS `users` (
  `id` varchar(36) NOT NULL COMMENT 'ユーザーID',
  `username` varchar(255) NOT NULL COMMENT 'ユーザー名',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `profile_image_url` varchar(255) DEFAULT NULL COMMENT 'プロフィール画像のURL',
  `bio` varchar(1024) DEFAULT NULL COMMENT '自己紹介文',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
