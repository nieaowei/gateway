create table service_load_balance
(
    id                       bigint auto_increment comment '自增主键'
        primary key,
    service_id               bigint        default 0                 not null comment '服务id',
    check_method             tinyint(20)   default 0                 not null comment '检查方法 0=tcpchk,检测端口是否握手成功',
    check_timeout            int(10)       default 0                 not null comment 'check超时时间,单位s',
    check_interval           int           default 0                 not null comment '检查间隔, 单位s',
    round_type               tinyint       default 2                 not null comment '轮询方式 0=random 1=round-robin 2=weight_round-robin 3=ip_hash',
    ip_list                  varchar(2000) default ''                not null comment 'ip列表',
    weight_list              varchar(2000) default ''                not null comment '权重列表',
    forbid_list              varchar(2000) default ''                not null comment '禁用ip列表',
    upstream_connect_timeout int           default 0                 not null comment '建立连接超时, 单位s',
    upstream_header_timeout  int           default 0                 not null comment '获取header超时, 单位s',
    upstream_idle_timeout    int(10)       default 0                 not null comment '链接最大空闲时间, 单位s',
    upstream_max_idle        int           default 0                 not null comment '最大空闲链接数',
    created_at               datetime      default CURRENT_TIMESTAMP not null comment '新增时间',
    updated_at               datetime      default CURRENT_TIMESTAMP not null comment '更新时间',
    deleted_at               datetime                                null comment '删除时间'
)
    comment '网关负载表';

INSERT INTO go_gateway.service_load_balance (id, service_id, check_method, check_timeout, check_interval, round_type, ip_list, weight_list, forbid_list, upstream_connect_timeout, upstream_header_timeout, upstream_idle_timeout, upstream_max_idle, created_at, updated_at, deleted_at) VALUES (196, 85, 19, 123, 51, 18, 'in dolore sint', 'tempor elit esse', '30', 123, 123, 73, 7, '2020-07-23 03:27:30', '2020-07-23 03:27:30', '2020-07-23 03:28:03');
INSERT INTO go_gateway.service_load_balance (id, service_id, check_method, check_timeout, check_interval, round_type, ip_list, weight_list, forbid_list, upstream_connect_timeout, upstream_header_timeout, upstream_idle_timeout, upstream_max_idle, created_at, updated_at, deleted_at) VALUES (197, 86, 19, 123, 51, 18, 'in dolore sint', 'tempor elit esse', '30', 123, 123, 73, 7, '2020-07-23 03:27:43', '2020-07-23 03:27:43', null);