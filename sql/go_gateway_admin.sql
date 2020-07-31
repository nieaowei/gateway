create table admin
(
    id         bigint auto_increment comment '自增id'
        primary key,
    username   varchar(255) default ''                not null comment '用户名',
    salt       varchar(50)  default ''                not null comment '盐',
    password   varchar(255) default ''                not null comment '密码',
    created_at datetime     default CURRENT_TIMESTAMP not null comment '新增时间',
    updated_at datetime     default CURRENT_TIMESTAMP not null comment '更新时间',
    deleted_at datetime                               null comment '删除时间'
)
    comment '管理员表';

INSERT INTO go_gateway.admin (id, username, salt, password, created_at, updated_at, deleted_at) VALUES (1, 'admin', '123', '6d8b2aadeecc1a9504b396ad74697f5675aca7d6751c42747ac42403cb3b9ef7', '2020-04-10 16:42:05', '2020-07-12 16:42:18', null);