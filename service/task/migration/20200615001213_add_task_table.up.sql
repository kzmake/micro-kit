BEGIN;

CREATE TABLE `tasks` (
  `id` varchar(26) NOT NULL,
  `description` varchar(255) NOT NULL,
  `created_at` datetime(6) NULL DEFAULT NULL,
  `updated_at` datetime(6) NULL DEFAULT NULL,
  `deleted_at` datetime(6) NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_tasks_deleted_at` (`deleted_at`)
) ENGINE=InnoDB CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

COMMIT;
