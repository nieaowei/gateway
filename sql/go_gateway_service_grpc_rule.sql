create table service_grpc_rule
(
    id              bigint auto_increment comment '自增主键'
        primary key,
    service_id      bigint        default 0  not null comment '服务id',
    port            int(5)        default 0  not null comment '端口',
    header_transfor varchar(5000) default '' not null comment 'header转换支持增加(add)、删除(del)、修改(edit) 格式: add headname headvalue 多个逗号间隔'
    created_at datetime     default current_timestamp not null comment '新增时间',
    updated_at datetime     default current_timestamp not null comment '更新时间',
    deleted_at datetime                                   null comment '删除时间'
)
    comment '网关路由匹配表';

INSERT INTO go_gateway.service_grpc_rule (id, service_id, port, header_transfor) VALUES (171, 53, 8009, '');
INSERT INTO go_gateway.service_grpc_rule (id, service_id, port, header_transfor) VALUES (172, 54, 8002, 'add metadata1 datavalue,edit metadata2 datavalue2');
INSERT INTO go_gateway.service_grpc_rule (id, service_id, port, header_transfor) VALUES (173, 58, 8012, 'add meta_name meta_value');
INSERT INTO go_gateway.service_grpc_rule (id, service_id, port, header_transfor) VALUES (174, 50, 8888, '');