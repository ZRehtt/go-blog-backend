CREATE TABLE IF NOT EXISTS `blog_tag` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(100) DEFAULT '' COMMENT '标签名称',
    `created_at` timestamp DEFAULT NULL COMMENT '创建时间',
    `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
    `updated_at` timestamp DEFAULT NULL COMMENT '修改时间',
    `updated_by` varchar(100) DEFAULT '' COMMENT '修改人',
    `deleted_at` timestamp DEFAULT NULL COMMENT '删除时间',
    `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态0为禁用，1为启用',
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章标签管理';


CREATE TABLE IF NOT EXISTS `blog_article` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `tag_id` int(10) unsigned DEFAULT '0' COMMENT '标签ID',
    `title` varchar(100) DEFAULT '' COMMENT '文章标题',
    `description` varchar(255) DEFAULT '' COMMENT '简述',
    `content` text COMMENT '内容',
    `created_at` timestamp DEFAULT NULL COMMENT '创建时间',
    `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
    `updated_at` timestamp DEFAULT NULL COMMENT '修改时间',
    `updated_by` varchar(100) DEFAULT '' COMMENT '修改人',
    `deleted_at` timestamp DEFAULT NULL COMMENT '删除时间',
    `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态0为禁用，1为启用',
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章管理';


CREATE TABLE IF NOT EXISTS `blog_user` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `username` varchar(50) DEFAULT '' COMMENT '账号',
    `password_hash` varchar(50) DEFAULT '' COMMENT '密码',
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户管理';

INSERT INTO `my_blog`.`blog_user` (`id`, `username`, `password_hash`) VALUES (null, 'test', 'test123456');