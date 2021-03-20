DROP DATABASE IF EXISTS `my_blog`;
CREATE DATABASE `my_blog`;

CREATE TABLE IF NOT EXISTS `blog_tag` (
    `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` varchar(100) NOT NULL COMMENT '标签名称',
    `created_at` timestamp DEFAULT NULL COMMENT '创建时间',
    `created_by` varchar(100) NOT NULL COMMENT '创建人',
    `updated_at` timestamp DEFAULT NULL COMMENT '修改时间',
    `updated_by` varchar(100) NOT NULL COMMENT '修改人',
    `deleted_at` timestamp DEFAULT NULL COMMENT '删除时间',
    `is_deleted` tinyint(3) UNSIGNED DEFAULT '0' COMMENT '是否删除 0为未删除，1为已删除',
    `state` tinyint(3) UNSIGNED DEFAULT '1' COMMENT '状态 0为禁用，1为启用',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章标签管理';


CREATE TABLE IF NOT EXISTS `blog_article` (
    `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    -- `tag_id` int(10) UNSIGNED DEFAULT '0' COMMENT '标签ID',
    `title` varchar(100) DEFAULT '' COMMENT '文章标题',
    `description` varchar(255) DEFAULT '' COMMENT '文章简述',
    `cover_url` varchar(255) DEFAULT '' COMMENT '封面图片地址', -- 新增封面图片
    `content` longtext COMMENT '文章内容',
    `created_at` timestamp DEFAULT NULL COMMENT '创建时间',
    `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
    `updated_at` timestamp DEFAULT NULL COMMENT '修改时间',
    `updated_by` varchar(100) DEFAULT '' COMMENT '修改人',
    `deleted_at` timestamp DEFAULT NULL COMMENT '删除时间',
    `is_deleted` tinyint(3) UNSIGNED DEFAULT '0' COMMENT '是否删除 0为未删除，1为已删除',
    `state` tinyint(3) UNSIGNED DEFAULT '1' COMMENT '状态0为禁用，1为启用',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章管理';


CREATE TABLE IF NOT EXISTS `blog_author` (
    `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `app_key` varchar(20) DEFAULT '' COMMENT 'Key',
    `app_secret` varchar(50) DEFAULT '' COMMENT 'Secret',
    `created_at` timestamp DEFAULT NULL COMMENT '创建时间',
    `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
    `updated_at` timestamp DEFAULT NULL COMMENT '修改时间',
    `updated_by` varchar(100) DEFAULT '' COMMENT '修改人',
    `deleted_at` timestamp DEFAULT NULL COMMENT '删除时间',
    `is_deleted` tinyint(3) UNSIGNED DEFAULT '0' COMMENT '是否删除 0为未删除，1为已删除',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='认证管理';

CREATE TABLE IF NOT EXISTS `blog_article_tag` (
    `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `article_id` int(10) NOT NULL COMMENT '文章ID',
    `tag_id` int(10) NOT NULL DEFAULT '0' COMMENT '标签ID',
    `created_at` timestamp DEFAULT NULL COMMENT '创建时间',
    `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
    `updated_at` timestamp DEFAULT NULL COMMENT '修改时间',
    `updated_by` varchar(100) DEFAULT '' COMMENT '修改人',
    `deleted_at` timestamp DEFAULT NULL COMMENT '删除时间',
    `is_deleted` tinyint(3) UNSIGNED DEFAULT '0' COMMENT '是否删除 0为未删除，1为已删除',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章标签关联表';