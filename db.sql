CREATE TABLE `voices` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '属性ID - 使用Snowflake算法生成',
  `voice_id` varchar(50) NOT NULL COMMENT '声线ID（唯一标识）',
  `scenario_id` bigint NOT NULL COMMENT '场景ID - 使用Snowflake算法生成',
  `name` varchar(100) NOT NULL COMMENT '声线名称',
  `icon_url` varchar(500) DEFAULT '' COMMENT '声线图标URL',
  `description` text COMMENT '声线描述',
  `preview_url` varchar(500) DEFAULT '' COMMENT '试听地址URL',
  `gender` varchar(20) DEFAULT 'neutral' COMMENT '性别：male=男性，female=女性，neutral=中性',
  `age_group` varchar(20) DEFAULT 'adult' COMMENT '年龄组：child=儿童，young=青年，adult=成年，elderly=老年',
  `style` varchar(100) DEFAULT '' COMMENT '声线风格（如：温柔、活泼、专业等）',
  `language` varchar(20) DEFAULT 'zh-CN' COMMENT '支持语言',
  `sort_order` int DEFAULT '0' COMMENT '排序权重',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态：1=启用，0=禁用',
  `created_at` bigint NOT NULL COMMENT '创建时间',
  `updated_at` bigint NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_voice_id` (`voice_id`),
  KEY `idx_scenario_id` (`scenario_id`),
  KEY `idx_status` (`status`),
  KEY `idx_sort_order` (`sort_order`),
  CONSTRAINT `fk_voices_scenario` FOREIGN KEY (`scenario_id`) REFERENCES `scenarios` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='声线配置表';






