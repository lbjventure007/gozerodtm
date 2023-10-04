CREATE TABLE `user` (
                        `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
                        `username` varchar(50) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
                        `password` varchar(50) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户密码，MD5加密',
                        `phone` varchar(20) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机号',
                        `question` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '找回密码问题',
                        `answer` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '找回密码答案',
                        `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                        `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                        `balance` decimal(16,2) DEFAULT NULL,
                        `forzen_balance` decimal(16,2) DEFAULT NULL,
                        PRIMARY KEY (`id`),
                        KEY `ix_update_time` (`update_time`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户表';