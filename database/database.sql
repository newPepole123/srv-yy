CREATE TABLE `ticket_types` (
  `id` bigint(20) unsigned NOT NULL,
  `ticket_type_name` varchar(127) NOT NULL DEFAULT '',
  `ticket_type_description` varchar(255),
  `status` int(10) NOT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
);
CREATE TABLE `tickets` (
  `id` bigint(20) unsigned NOT NULL,
  `ticket_type_id` bigint(20) unsigned NOT NULL,
  `user_id` bigint(20) unsigned NOT NULL,
  `assignee_id` bigint(20) unsigned NOT NULL,
  `ticket_name` varchar(255) NOT NULL DEFAULT '',
  `ticket_description` varchar(255),
  `priority` int(10) NOT NULL, --数值越大，级别越高
  `status` int(10) NOT NULL, --0表示未准备，1表示未读，大于1都是已读
  `deleted_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_ticket_ticket_type_id` (`ticket_type_id`),
  CONSTRAINT `fk_ticket_ticket_type_id` FOREIGN KEY (`ticket_type_id`) REFERENCES `ticket_types` (`id`)
);
