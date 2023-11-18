CREATE TABLE `user`
(
    `id`                INT(10) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '自增Id',

    `created_at`        TIMESTAMP    NOT NULL DEFAULT current_timestamp COMMENT '创建时间',
    `updated_at`        TIMESTAMP    NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp COMMENT '更新时间',
    `deleted_at`        DATETIME     NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '删除时间'
) COMMENT ='用户表';