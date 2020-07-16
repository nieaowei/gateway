create table service_tcp_rule
(
    id         bigint auto_increment comment '自增主键'
        primary key,
    service_id bigint           not null comment '服务id',
    port       int(5) default 0 not null comment '端口号',
    created_at datetime     default current_timestamp not null comment '新增时间',
    updated_at datetime     default current_timestamp not null comment '更新时间',
    deleted_at datetime                                   null comment '删除时间'
)
    comment '网关路由匹配表';

INSERT INTO go_gateway.service_tcp_rule (id, service_id, port) VALUES (171, 41, 8002);
INSERT INTO go_gateway.service_tcp_rule (id, service_id, port) VALUES (172, 42, 8003);
INSERT INTO go_gateway.service_tcp_rule (id, service_id, port) VALUES (173, 43, 8004);
INSERT INTO go_gateway.service_tcp_rule (id, service_id, port) VALUES (174, 38, 8004);
INSERT INTO go_gateway.service_tcp_rule (id, service_id, port) VALUES (175, 45, 8001);
INSERT INTO go_gateway.service_tcp_rule (id, service_id, port) VALUES (176, 46, 8005);
INSERT INTO go_gateway.service_tcp_rule (id, service_id, port) VALUES (177, 50, 8006);
INSERT INTO go_gateway.service_tcp_rule (id, service_id, port) VALUES (178, 51, 8007);
INSERT INTO go_gateway.service_tcp_rule (id, service_id, port) VALUES (179, 52, 8008);
INSERT INTO go_gateway.service_tcp_rule (id, service_id, port) VALUES (180, 55, 8010);
INSERT INTO go_gateway.service_tcp_rule (id, service_id, port) VALUES (181, 57, 8011);