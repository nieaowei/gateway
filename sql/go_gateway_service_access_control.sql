create table service_access_control
(
    id                  bigint auto_increment comment '自增主键'
        primary key,
    service_id          bigint        default 0  not null comment '服务id',
    open_auth           tinyint       default 0  not null comment '是否开启权限 1=开启',
    black_list          varchar(1000) default '' not null comment '黑名单ip',
    white_list          varchar(1000) default '' not null comment '白名单ip',
    white_host_name     varchar(1000) default '' not null comment '白名单主机',
    clientip_flow_limit int           default 0  not null comment '客户端ip限流',
    service_flow_limit  int(20)       default 0  not null comment '服务端限流',
    created_at datetime     default current_timestamp not null comment '新增时间',
    updated_at datetime     default current_timestamp not null comment '更新时间',
    deleted_at datetime                                   null comment '删除时间'
)
    comment '网关权限控制表';

INSERT INTO go_gateway.service_access_control (id, service_id, open_auth, black_list, white_list, white_host_name, clientip_flow_limit, service_flow_limit) VALUES (162, 35, 1, '', '', '', 0, 0);
INSERT INTO go_gateway.service_access_control (id, service_id, open_auth, black_list, white_list, white_host_name, clientip_flow_limit, service_flow_limit) VALUES (165, 34, 0, '', '', '', 0, 0);
INSERT INTO go_gateway.service_access_control (id, service_id, open_auth, black_list, white_list, white_host_name, clientip_flow_limit, service_flow_limit) VALUES (167, 36, 0, '', '', '', 0, 0);
INSERT INTO go_gateway.service_access_control (id, service_id, open_auth, black_list, white_list, white_host_name, clientip_flow_limit, service_flow_limit) VALUES (168, 38, 1, '111.11', '22.33', '11.11', 12, 12);
INSERT INTO go_gateway.service_access_control (id, service_id, open_auth, black_list, white_list, white_host_name, clientip_flow_limit, service_flow_limit) VALUES (169, 41, 1, '111.11', '22.33', '11.11', 12, 12);
INSERT INTO go_gateway.service_access_control (id, service_id, open_auth, black_list, white_list, white_host_name, clientip_flow_limit, service_flow_limit) VALUES (170, 42, 1, '111.11', '22.33', '11.11', 12, 12);
INSERT INTO go_gateway.service_access_control (id, service_id, open_auth, black_list, white_list, white_host_name, clientip_flow_limit, service_flow_limit) VALUES (171, 43, 0, '111.11', '22.33', '11.11', 12, 12);
INSERT INTO go_gateway.service_access_control (id, service_id, open_auth, black_list, white_list, white_host_name, clientip_flow_limit, service_flow_limit) VALUES (172, 44, 0, '', '', '', 0, 0);
INSERT INTO go_gateway.service_access_control (id, service_id, open_auth, black_list, white_list, white_host_name, clientip_flow_limit, service_flow_limit) VALUES (173, 45, 0, '', '', '', 0, 0);
INSERT INTO go_gateway.service_access_control (id, service_id, open_auth, black_list, white_list, white_host_name, clientip_flow_limit, service_flow_limit) VALUES (174, 46, 0, '', '', '', 0, 0);
INSERT INTO go_gateway.service_access_control (id, service_id, open_auth, black_list, white_list, white_host_name, clientip_flow_limit, service_flow_limit) VALUES (175, 47, 0, '', '', '', 0, 0);
INSERT INTO go_gateway.service_access_control (id, service_id, open_auth, black_list, white_list, white_host_name, clientip_flow_limit, service_flow_limit) VALUES (176, 48, 0, '', '', '', 0, 0);
INSERT INTO go_gateway.service_access_control (id, service_id, open_auth, black_list, white_list, white_host_name, clientip_flow_limit, service_flow_limit) VALUES (177, 49, 0, '', '', '', 0, 0);
INSERT INTO go_gateway.service_access_control (id, service_id, open_auth, black_list, white_list, white_host_name, clientip_flow_limit, service_flow_limit) VALUES (178, 50, 0, '', '', '', 0, 0);
INSERT INTO go_gateway.service_access_control (id, service_id, open_auth, black_list, white_list, white_host_name, clientip_flow_limit, service_flow_limit) VALUES (179, 51, 0, '', '', '', 0, 0);
INSERT INTO go_gateway.service_access_control (id, service_id, open_auth, black_list, white_list, white_host_name, clientip_flow_limit, service_flow_limit) VALUES (180, 52, 0, '', '', '', 0, 0);
INSERT INTO go_gateway.service_access_control (id, service_id, open_auth, black_list, white_list, white_host_name, clientip_flow_limit, service_flow_limit) VALUES (181, 53, 0, '', '', '', 0, 0);
INSERT INTO go_gateway.service_access_control (id, service_id, open_auth, black_list, white_list, white_host_name, clientip_flow_limit, service_flow_limit) VALUES (182, 54, 1, '127.0.0.3', '127.0.0.2', '', 11, 12);
INSERT INTO go_gateway.service_access_control (id, service_id, open_auth, black_list, white_list, white_host_name, clientip_flow_limit, service_flow_limit) VALUES (183, 55, 1, '127.0.0.2', '127.0.0.1', '', 45, 34);
INSERT INTO go_gateway.service_access_control (id, service_id, open_auth, black_list, white_list, white_host_name, clientip_flow_limit, service_flow_limit) VALUES (184, 56, 0, '192.168.1.0', '', '', 0, 0);
INSERT INTO go_gateway.service_access_control (id, service_id, open_auth, black_list, white_list, white_host_name, clientip_flow_limit, service_flow_limit) VALUES (185, 57, 0, '', '127.0.0.1,127.0.0.2', '', 0, 0);
INSERT INTO go_gateway.service_access_control (id, service_id, open_auth, black_list, white_list, white_host_name, clientip_flow_limit, service_flow_limit) VALUES (186, 58, 1, '', '', '', 0, 0);
INSERT INTO go_gateway.service_access_control (id, service_id, open_auth, black_list, white_list, white_host_name, clientip_flow_limit, service_flow_limit) VALUES (187, 59, 1, '127.0.0.1', '', '', 2, 3);
INSERT INTO go_gateway.service_access_control (id, service_id, open_auth, black_list, white_list, white_host_name, clientip_flow_limit, service_flow_limit) VALUES (188, 60, 1, '', '', '', 0, 0);
INSERT INTO go_gateway.service_access_control (id, service_id, open_auth, black_list, white_list, white_host_name, clientip_flow_limit, service_flow_limit) VALUES (189, 61, 0, '', '', '', 0, 0);