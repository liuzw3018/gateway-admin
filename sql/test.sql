CREATE TABLE `area` (
                        `id` bigint(20) NOT NULL AUTO_INCREMENT,
                        `area_name` varchar(255) NOT NULL,
                        `city_id` int(11) NOT NULL,
                        `user_id` int(11) NOT NULL,
                        `update_at` datetime NOT NULL,
                        `create_at` datetime NOT NULL,
                        `delete_at` datetime NOT NULL,
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='area';
INSERT INTO `area` (`id`, `area_name`, `city_id`, `user_id`, `update_at`, `create_at`, `delete_at`) VALUES (NULL, 'area_name', '1', '2', '2019-06-15 00:00:00', '2019-06-15 00:00:00', '2019-06-15 00:00:00');
