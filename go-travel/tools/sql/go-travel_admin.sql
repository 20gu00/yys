-- go-travel_manager --
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin` (
                        `id` bigint NOT NULL AUTO_INCREMENT,
                        `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        `user_name` varchar(255) NOT NULL DEFAULT '' COMMENT '用户名',
                        `salt` varchar(50) NOT NULL DEFAULT '' COMMENT  '盐',
                        `introduction` varchar(50) NOT NULL DEFAULT '',
                        `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
			`user_id` int(100) NOT NULL,
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='admin';
