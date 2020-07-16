create table service_info
(
    id           bigint unsigned auto_increment comment '自增主键'
        primary key,
    load_type    tinyint      default 0                     not null comment '负载类型 0=http 1=tcp 2=grpc',
    service_name varchar(255) default ''                    not null comment '服务名称 6-128 数字字母下划线',
    service_desc varchar(255) default ''                    not null comment '服务描述',
    created_at datetime     default current_timestamp not null comment '新增时间',
    updated_at datetime     default current_timestamp not null comment '更新时间',
    deleted_at datetime                                   null comment '删除时间'
)
    comment '网关基本信息表';

INSERT INTO go_gateway.service_info (id, load_type, service_name, service_desc, created_at, updated_at, deleted_at) VALUES (34, 0, 'websocket_test', 'websocket_test', '2020-04-13 01:31:47', '1971-01-01 00:00:00', null);
INSERT INTO go_gateway.service_info (id, load_type, service_name, service_desc, created_at, updated_at, deleted_at) VALUES (35, 0, 'test_grpc', 'test_grpc', '2020-04-13 01:34:32', '1971-01-01 00:00:00', null);
INSERT INTO go_gateway.service_info (id, load_type, service_name, service_desc, created_at, updated_at, deleted_at) VALUES (36, 2, 'test_httpe', 'test_httpe', '2020-04-11 21:12:48', '1971-01-01 00:00:00', null);
INSERT INTO go_gateway.service_info (id, load_type, service_name, service_desc, created_at, updated_at, deleted_at) VALUES (38, 0, 'service_name', '11111', '2020-04-15 07:49:45', '2020-04-11 23:59:39', null);
INSERT INTO go_gateway.service_info (id, load_type, service_name, service_desc, created_at, updated_at, deleted_at) VALUES (41, 0, 'service_name_tcp', '11111', '2020-04-13 01:38:01', '2020-04-12 01:06:09', null);
INSERT INTO go_gateway.service_info (id, load_type, service_name, service_desc, created_at, updated_at, deleted_at) VALUES (42, 0, 'service_name_tcp2', '11111', '2020-04-13 01:38:06', '2020-04-12 01:13:24', null);
INSERT INTO go_gateway.service_info (id, load_type, service_name, service_desc, created_at, updated_at, deleted_at) VALUES (43, 1, 'service_name_tcp4', 'service_name_tcp4', '2020-04-15 07:49:44', '2020-04-12 01:13:50', null);
INSERT INTO go_gateway.service_info (id, load_type, service_name, service_desc, created_at, updated_at, deleted_at) VALUES (44, 0, 'websocket_service', 'websocket_service', '2020-04-15 07:49:43', '2020-04-13 01:20:08', null);
INSERT INTO go_gateway.service_info (id, load_type, service_name, service_desc, created_at, updated_at, deleted_at) VALUES (45, 1, 'tcp_service', 'tcp_desc', '2020-04-15 07:49:41', '2020-04-13 01:46:27', null);
INSERT INTO go_gateway.service_info (id, load_type, service_name, service_desc, created_at, updated_at, deleted_at) VALUES (46, 1, 'grpc_service', 'grpc_desc', '2020-04-13 01:54:12', '2020-04-13 01:53:14', null);
INSERT INTO go_gateway.service_info (id, load_type, service_name, service_desc, created_at, updated_at, deleted_at) VALUES (47, 0, 'testsefsafs', 'werrqrr', '2020-04-13 01:59:14', '2020-04-13 01:57:49', null);
INSERT INTO go_gateway.service_info (id, load_type, service_name, service_desc, created_at, updated_at, deleted_at) VALUES (48, 0, 'testsefsafs1', 'werrqrr', '2020-04-13 01:59:11', '2020-04-13 01:58:14', null);
INSERT INTO go_gateway.service_info (id, load_type, service_name, service_desc, created_at, updated_at, deleted_at) VALUES (49, 0, 'testsefsafs1222', 'werrqrr', '2020-04-13 01:59:08', '2020-04-13 01:58:23', null);
INSERT INTO go_gateway.service_info (id, load_type, service_name, service_desc, created_at, updated_at, deleted_at) VALUES (50, 2, 'grpc_service_name', 'grpc_service_desc', '2020-04-15 07:49:40', '2020-04-13 02:01:00', null);
INSERT INTO go_gateway.service_info (id, load_type, service_name, service_desc, created_at, updated_at, deleted_at) VALUES (51, 2, 'gresafsf', 'wesfsf', '2020-04-15 07:49:39', '2020-04-13 02:01:57', null);
INSERT INTO go_gateway.service_info (id, load_type, service_name, service_desc, created_at, updated_at, deleted_at) VALUES (52, 2, 'gresafsf11', 'wesfsf', '2020-04-13 02:03:41', '2020-04-13 02:02:55', null);
INSERT INTO go_gateway.service_info (id, load_type, service_name, service_desc, created_at, updated_at, deleted_at) VALUES (53, 2, 'tewrqrw111', '123313', '2020-04-13 02:03:38', '2020-04-13 02:03:20', null);
INSERT INTO go_gateway.service_info (id, load_type, service_name, service_desc, created_at, updated_at, deleted_at) VALUES (54, 2, 'test_grpc_service1', 'test_grpc_service1', '2020-04-15 07:49:37', '2020-04-15 07:38:43', null);
INSERT INTO go_gateway.service_info (id, load_type, service_name, service_desc, created_at, updated_at, deleted_at) VALUES (55, 1, 'test_tcp_service1', 'redis服务代理', '2020-04-15 07:49:35', '2020-04-15 07:46:35', null);
INSERT INTO go_gateway.service_info (id, load_type, service_name, service_desc, created_at, updated_at, deleted_at) VALUES (56, 0, 'test_http_service', '测试HTTP代理', '2020-04-16 00:54:45', '2020-04-15 07:55:07', null);
INSERT INTO go_gateway.service_info (id, load_type, service_name, service_desc, created_at, updated_at, deleted_at) VALUES (57, 1, 'test_tcp_service', '测试TCP代理', '2020-04-19 14:03:09', '2020-04-15 07:58:39', null);
INSERT INTO go_gateway.service_info (id, load_type, service_name, service_desc, created_at, updated_at, deleted_at) VALUES (58, 2, 'test_grpc_service', '测试GRPC服务', '2020-04-21 07:20:16', '2020-04-15 07:59:46', null);
INSERT INTO go_gateway.service_info (id, load_type, service_name, service_desc, created_at, updated_at, deleted_at) VALUES (59, 0, 'test.com:8080', '测试域名接入', '2020-04-18 22:54:14', '2020-04-18 20:29:13', null);
INSERT INTO go_gateway.service_info (id, load_type, service_name, service_desc, created_at, updated_at, deleted_at) VALUES (60, 0, 'test_strip_uri', '测试路径接入', '2020-04-21 06:55:26', '2020-04-18 22:56:37', null);
INSERT INTO go_gateway.service_info (id, load_type, service_name, service_desc, created_at, updated_at, deleted_at) VALUES (61, 0, 'test_https_server', '测试https服务', '2020-04-19 12:22:33', '2020-04-19 12:17:04', '2020-07-16 21:42:59');