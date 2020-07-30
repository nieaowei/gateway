create table service_tcp_rule
(
    id         bigint auto_increment comment '自增主键'
        primary key,
    service_id bigint                             not null comment '服务id',
    port       int(5)   default 0                 not null comment '端口号',
    created_at datetime default CURRENT_TIMESTAMP not null comment '新增时间',
    updated_at datetime default CURRENT_TIMESTAMP not null comment '更新时间',
    deleted_at datetime                           null comment '删除时间'
)
    comment '网关路由匹配表';

