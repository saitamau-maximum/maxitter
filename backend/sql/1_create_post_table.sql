CREATE TABLE IF NOT EXISTS `posts` (
  `id` varchar(36) NOT NULL COMMENT '投稿ID',
  `body` text NOT NULL COMMENT '投稿の本文',
  `created_at` datetime NOT NULL COMMENT '投稿日時',
  `user_id` varchar(36) NOT NULL COMMENT '投稿作成者のユーザーID'
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
