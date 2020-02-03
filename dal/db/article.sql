-- 文章表
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '文章id',
    `category_id` bigint(20) UNSIGNED NOT NULL COMMENT '分类id',
    `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文章内容',
    `title` VARCHAR(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文章标题',
    `view_count` int(255) UNSIGNED NOT NULL DEFAULT 0 COMMENT '阅读次数',
    `comment_count` int(255) UNSIGNED NOT NULL DEFAULT 0 COMMENT '评论次数',
    `username` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '作者',
    `status` int(10) UNSIGNED NOT NULL DEFAULT 1 COMMENT '状态',
    `summary` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文章摘要',
    `create_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发布时间',
    `update_time` timestamp(0) NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    INDEX `idx_view_count`(`view_count`) USING BTREE COMMENT '阅读次数索引',
    INDEX `idx_comment_count`(`comment_count`) USING BTREE COMMENT '评论数索引',
    INDEX `idx_category_id`(`category_id`) USING BTREE COMMENT '分类id索引'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;


-- 分类表
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `category_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '分类名字',
    `category_no` int(10) UNSIGNED NOT NULL COMMENT '分类排序',
    `create_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp(0) NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0),
    PRIMARY KEY (`id`) USING BTREE
)ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- 评论表
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
    `id` bigint(20)  NOT NULL AUTO_INCREMENT COMMENT '评论id',
    `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '评论内容',
    `username` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '评论作者',
    `create_time` timestamp(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '评论发布时间',
    `status` int(255) UNSIGNED NOT NULL COMMENT '评论状态：0，删除；1，正常',
    `article_id` bigint(20) UNSIGNED NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- 留言表
DROP TABLE IF EXISTS `message_board`;
CREATE TABLE `leaves` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `email` VARCHAR(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `create_time` TIMESTAMP(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` TIMESTAMP(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0),
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;


