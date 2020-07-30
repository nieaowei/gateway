create table service_info
(
    id           bigint unsigned auto_increment comment '自增主键'
        primary key,
    load_type    tinyint      default 0                 not null comment '负载类型 0=http 1=tcp 2=grpc',
    service_name varchar(255) default ''                not null comment '服务名称 6-128 数字字母下划线',
    service_desc varchar(255) default ''                not null comment '服务描述',
    created_at   datetime     default CURRENT_TIMESTAMP not null comment '新增时间',
    updated_at   datetime     default CURRENT_TIMESTAMP not null comment '更新时间',
    deleted_at   datetime                               null comment '删除时间'
)
    comment '网关基本信息表';

INSERT INTO go_gateway.service_info (id, load_type, service_name, service_desc, created_at, updated_at, deleted_at) VALUES (85, 0, '1233333', 'est magna nulla Excepteur', '2020-07-23 03:27:30', '2020-07-23 03:27:30', '2020-07-23 03:28:03');
INSERT INTO go_gateway.service_info (id, load_type, service_name, service_desc, created_at, updated_at, deleted_at) VALUES (86, 0, '3321231', 'est magna nulla Excepteur', '2020-07-23 03:27:42', '2020-07-23 03:27:42', null);