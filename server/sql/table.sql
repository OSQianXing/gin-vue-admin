

CREATE TABLE IF NOT EXISTS `fragment` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '碎片视频名称',
  `url` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '碎片视频存储地址',
  `author` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '碎片视频作者',
  `subtitles` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '碎片视频字幕',
  `uploaded_at` datetime DEFAULT NULL COMMENT '碎片视频上传时间',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '碎片视频审核状态（0 未审核 1 初审中 3 初审完成 4 复审审核中 5 复审完成）',
  `deleted_at` datetime DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '碎片创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='碎片视频表';



CREATE TABLE IF NOT EXISTS `fragment_audit` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL COMMENT '审核人 ID',
  `fragment_id` bigint NOT NULL COMMENT '碎片视频ID',
  `type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '审核类型  0 初审一审  1 初审二审 2 复审',
  `clipping_range` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '剪辑范围（掐头去尾）',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '审核结果  0 未审核 1 不通过 2 neededit 3 pending  4 ok  5 great',
  `reason` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '审核失败原因',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '审核创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='碎片审核表';


CREATE TABLE IF NOT EXISTS `fragment2words` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `fragment_id` bigint NOT NULL COMMENT '碎片视频ID',
  `words` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '单词',
  `removed` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否移除（0-默认 1-已移除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='碎片单词映射表';





CREATE TABLE IF NOT EXISTS `operator_log` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;




CREATE TABLE IF NOT EXISTS `play_list` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '播单ID',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '播单名称',
  `online_version` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '上线版本',
  `deleted_at` datetime DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '播单创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='播单表';




CREATE TABLE IF NOT EXISTS `play_list_video` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `play_list_id` bigint unsigned NOT NULL COMMENT '播单ID',
  `version` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '版本',
  `url` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '合成视频地址',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '短视频状态 0-待审核 1-审核通过 2-审核不通过',
  `audit_reason` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '审核不通过原因',
  `user_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '审核人ID',
  `audited_at` datetime DEFAULT NULL,
  `prod_id` int unsigned NOT NULL DEFAULT '0' COMMENT '成品ID',
  `prod_status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '成品状态 0-待审核 1-审核通过 2-审核不通过',
  `prod_audit_reason` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '成品审核不通过原因',
  `prod_user_id` bigint unsigned NOT NULL DEFAULT '0',
  `prod_audited_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='播单视频表';
