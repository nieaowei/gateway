create table service_http_rule
(
    id              bigint auto_increment comment '自增主键'
        primary key,
    service_id      bigint                                  not null comment '服务id',
    rule_type       tinyint       default 0                 not null comment '匹配类型 0=url前缀url_prefix 1=域名domain ',
    rule            varchar(255)  default ''                not null comment 'type=domain表示域名，type=url_prefix时表示url前缀',
    need_https      tinyint       default 0                 not null comment '支持https 1=支持',
    need_strip_uri  tinyint       default 0                 not null comment '启用strip_uri 1=启用',
    need_websocket  tinyint       default 0                 not null comment '是否支持websocket 1=支持',
    url_rewrite     varchar(5000) default ''                not null comment 'url重写功能 格式：^/gatekeeper/test_service(.*) $1 多个逗号间隔',
    header_transfor varchar(5000) default ''                not null comment 'header转换支持增加(add)、删除(del)、修改(edit) 格式: add headname headvalue 多个逗号间隔',
    created_at      datetime      default CURRENT_TIMESTAMP not null comment '新增时间',
    updated_at      datetime      default CURRENT_TIMESTAMP not null comment '更新时间',
    deleted_at      datetime                                null comment '删除时间'
)
    comment '网关路由匹配表';

INSERT INTO go_gateway.service_http_rule (id, service_id, rule_type, rule, need_https, need_strip_uri, need_websocket, url_rewrite, header_transform, created_at, updated_at, deleted_at) VALUES (192, 85, 70, 'do sed sit', 85, 45, 36, 'in anim adipisicing mollit', 'enim', '2020-07-23 03:27:30', '2020-07-23 03:27:30', '2020-07-23 03:28:03');
INSERT INTO go_gateway.service_http_rule (id, service_id, rule_type, rule, need_https, need_strip_uri, need_websocket, url_rewrite, header_transform, created_at, updated_at, deleted_at) VALUES (193, 86, 70, 'do sed sit', 85, 45, 36, 'in anim adipisicing mollit', 'enim', '2020-07-23 03:27:42', '2020-07-23 03:27:42', null);