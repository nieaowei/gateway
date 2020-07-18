create table app
(
    id         bigint unsigned auto_increment comment '自增id'
        primary key,
    app_id     varchar(255)  default ''                    not null comment '租户id',
    name       varchar(255)  default ''                    not null comment '租户名称',
    secret     varchar(255)  default ''                    not null comment '密钥',
    white_ips  varchar(1000) default ''                    not null comment 'ip白名单，支持前缀匹配',
    qpd        bigint        default 0                     not null comment '日请求量限制',
    qps        bigint        default 0                     not null comment '每秒请求量限制',
    created_at datetime     default current_timestamp not null comment '新增时间',
    updated_at datetime     default current_timestamp not null comment '更新时间',
    deleted_at datetime                                   null comment '删除时间'
)
    comment '网关租户表';

INSERT INTO go_gateway.app (id, app_id, name, secret, white_ips, qpd, qps, created_at, updated_at, deleted_at) VALUES (31, 'app_id_a', '租户A', '449441eb5e72dca9c42a12f3924ea3a2', 'white_ips', 100000, 100, '2020-04-15 20:55:02', '2020-04-21 07:23:34', null);
INSERT INTO go_gateway.app (id, app_id, name, secret, white_ips, qpd, qps, created_at, updated_at, deleted_at) VALUES (32, 'app_id_b', '租户B', '8d7b11ec9be0e59a36b52f32366c09cb', '', 20, 0, '2020-04-15 21:40:52', '2020-04-21 07:23:27', null);
INSERT INTO go_gateway.app (id, app_id, name, secret, white_ips, qpd, qps, created_at, updated_at, deleted_at) VALUES (33, 'app_id', '租户名称', '', '', 0, 0, '2020-04-15 22:02:23', '2020-04-15 22:06:51', null);
INSERT INTO go_gateway.app (id, app_id, name, secret, white_ips, qpd, qps, created_at, updated_at, deleted_at) VALUES (34, 'app_id45', '名称', '07d980f8a49347523ee1d5c1c41aec02', '', 0, 0, '2020-04-15 22:06:38', '2020-04-15 22:06:49', null);