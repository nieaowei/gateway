create table service_access_control
(
    id                  bigint auto_increment comment '自增主键'
        primary key,
    service_id          bigint        default 0                 not null comment '服务id',
    open_auth           tinyint       default 0                 not null comment '是否开启权限 1=开启',
    black_list          varchar(1000) default ''                not null comment '黑名单ip',
    white_list          varchar(1000) default ''                not null comment '白名单ip',
    white_host_name     varchar(1000) default ''                not null comment '白名单主机',
    clientip_flow_limit int           default 0                 not null comment '客户端ip限流',
    service_flow_limit  int(20)       default 0                 not null comment '服务端限流',
    created_at          datetime      default CURRENT_TIMESTAMP not null comment '新增时间',
    updated_at          datetime      default CURRENT_TIMESTAMP not null comment '更新时间',
    deleted_at          datetime                                null comment '删除时间'
)
    comment '网关权限控制表';

INSERT INTO go_gateway.service_access_control (id, service_id, open_auth, black_list, white_list, white_host_name, clientip_flow_limit, service_flow_limit, created_at, updated_at, deleted_at) VALUES (196, 85, 69, 'officia adipisicing', 'deserunt anim', '千较量导相林', 14, 66, '2020-07-23 03:27:30', '2020-07-23 03:27:30', '2020-07-23 03:28:03');
INSERT INTO go_gateway.service_access_control (id, service_id, open_auth, black_list, white_list, white_host_name, clientip_flow_limit, service_flow_limit, created_at, updated_at, deleted_at) VALUES (197, 86, 69, 'officia adipisicing', 'deserunt anim', '千较量导相林', 14, 66, '2020-07-23 03:27:43', '2020-07-23 03:27:43', null);