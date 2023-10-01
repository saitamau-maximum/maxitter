package main

var (
	CREATE_POSTS_TABLE = /* sql */ `
CREATE TABLE IF NOT EXISTS posts (
  id varchar(36) NOT NULL COMMENT '投稿ID',
  body text NOT NULL COMMENT '投稿の本文',
  created_at datetime NOT NULL COMMENT '投稿日時',
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
`
)
