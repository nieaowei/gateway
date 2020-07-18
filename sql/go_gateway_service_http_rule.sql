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
    created_at      datetime      default current_timestamp not null comment '新增时间',
    updated_at      datetime      default current_timestamp not null comment '更新时间',
    deleted_at      datetime                                null comment '删除时间'
)
    comment '网关路由匹配表';

INSERT INTO go_gateway.service_http_rule (id, service_id, rule_type, rule, need_https, need_strip_uri, need_websocket,
                                          url_rewrite, header_transfor)
VALUES (165, 35, 1, '', 0, 0, 0, '', '');
INSERT INTO go_gateway.service_http_rule (id, service_id, rule_type, rule, need_https, need_strip_uri, need_websocket,
                                          url_rewrite, header_transfor)
VALUES (168, 34, 0, '', 0, 0, 0, '', '');
INSERT INTO go_gateway.service_http_rule (id, service_id, rule_type, rule, need_https, need_strip_uri, need_websocket,
                                          url_rewrite, header_transfor)
VALUES (170, 36, 0, '', 0, 0, 0, '', '');
INSERT INTO go_gateway.service_http_rule (id, service_id, rule_type, rule, need_https, need_strip_uri, need_websocket,
                                          url_rewrite, header_transfor)
VALUES (171, 38, 0, '/abc', 1, 0, 1, '^/abc $1', 'add head1 value1');
INSERT INTO go_gateway.service_http_rule (id, service_id, rule_type, rule, need_https, need_strip_uri, need_websocket,
                                          url_rewrite, header_transfor)
VALUES (172, 43, 0, '/usr', 1, 1, 0, '^/afsaasf $1,^/afsaasf $1', '');
INSERT INTO go_gateway.service_http_rule (id, service_id, rule_type, rule, need_https, need_strip_uri, need_websocket,
                                          url_rewrite, header_transfor)
VALUES (173, 44, 1, 'www.test.com', 1, 1, 1, '', '');
INSERT INTO go_gateway.service_http_rule (id, service_id, rule_type, rule, need_https, need_strip_uri, need_websocket,
                                          url_rewrite, header_transfor)
VALUES (174, 47, 1, 'www.test.com', 1, 1, 1, '', '');
INSERT INTO go_gateway.service_http_rule (id, service_id, rule_type, rule, need_https, need_strip_uri, need_websocket,
                                          url_rewrite, header_transfor)
VALUES (175, 48, 1, 'www.test.com', 1, 1, 1, '', '');
INSERT INTO go_gateway.service_http_rule (id, service_id, rule_type, rule, need_https, need_strip_uri, need_websocket,
                                          url_rewrite, header_transfor)
VALUES (176, 49, 1, 'www.test.com', 1, 1, 1, '', '');
INSERT INTO go_gateway.service_http_rule (id, service_id, rule_type, rule, need_https, need_strip_uri, need_websocket,
                                          url_rewrite, header_transfor)
VALUES (177, 56, 0, '/test_http_service', 1, 1, 1, '^/test_http_service/abb/(.*) /test_http_service/bba/$1',
        'add header_name header_value');
INSERT INTO go_gateway.service_http_rule (id, service_id, rule_type, rule, need_https, need_strip_uri, need_websocket,
                                          url_rewrite, header_transfor)
VALUES (178, 59, 1, 'test.com', 0, 1, 1, '', 'add headername headervalue');
INSERT INTO go_gateway.service_http_rule (id, service_id, rule_type, rule, need_https, need_strip_uri, need_websocket,
                                          url_rewrite, header_transfor)
VALUES (179, 60, 0, '/test_strip_uri', 0, 1, 0, '^/aaa/(.*) /bbb/$1', '');
INSERT INTO go_gateway.service_http_rule (id, service_id, rule_type, rule, need_https, need_strip_uri, need_websocket,
                                          url_rewrite, header_transfor)
VALUES (180, 61, 0, '/test_https_server', 1, 1, 0, '', '');