package dao

import (
	"gorm.io/gorm"
	"time"
)

/******sql******
CREATE TABLE `admin` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `username` varchar(255) NOT NULL DEFAULT '' COMMENT '用户名',
  `salt` varchar(50) NOT NULL DEFAULT '' COMMENT '盐',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '密码',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '新增时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1 COMMENT='管理员表'
******sql******/
// Admin 管理员表
type Admin struct {
	gorm.Model
	Username string `gorm:"column:username;type:varchar(255);not null" json:"username"` // 用户名
	Avatar   string `gorm:"column:avatar;type:varchar(255);not null" json:"avatar"`     // 头像
	Salt     string `gorm:"column:salt;type:varchar(50);not null" json:"salt"`          // 盐
	Password string `gorm:"column:password;type:varchar(255);not null" json:"password"` // 密码
}

// TableName get sql table name.获取数据库表名
func (m *Admin) TableName() string {
	return "admin"
}

/******sql******
CREATE TABLE `app` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `app_id` varchar(255) NOT NULL DEFAULT '' COMMENT '租户id',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '租户名称',
  `secret` varchar(255) NOT NULL DEFAULT '' COMMENT '密钥',
  `white_ips` varchar(1000) NOT NULL DEFAULT '' COMMENT 'ip白名单，支持前缀匹配',
  `qpd` bigint(20) NOT NULL DEFAULT '0' COMMENT '日请求量限制',
  `qps` bigint(20) NOT NULL DEFAULT '0' COMMENT '每秒请求量限制',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '新增时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COMMENT='网关租户表'
******sql******/
// App 网关租户表
type App struct {
	gorm.Model
	AppID    string `gorm:"column:app_id;type:varchar(255);not null" json:"app_id"`        // 租户id
	Name     string `gorm:"column:name;type:varchar(255);not null" json:"name"`            // 租户名称
	Secret   string `gorm:"column:secret;type:varchar(255);not null" json:"secret"`        // 密钥
	WhiteIPs string `gorm:"column:white_ips;type:varchar(1000);not null" json:"white_ips"` // ip白名单，支持前缀匹配
	Qpd      int64  `gorm:"column:qpd;type:bigint(20);not null" json:"qpd"`                // 日请求量限制
	QPS      int64  `gorm:"column:qps;type:bigint(20);not null" json:"qps"`                // 每秒请求量限制
}

// TableName get sql table name.获取数据库表名
func (m *App) TableName() string {
	return "app"
}

/******sql******
CREATE TABLE `service_access_control` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `service_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '服务id',
  `open_auth` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否开启权限 1=开启',
  `black_list` varchar(1000) NOT NULL DEFAULT '' COMMENT '黑名单ip',
  `white_list` varchar(1000) NOT NULL DEFAULT '' COMMENT '白名单ip',
  `white_host_name` varchar(1000) NOT NULL DEFAULT '' COMMENT '白名单主机',
  `clientip_flow_limit` int(11) NOT NULL DEFAULT '0' COMMENT '客户端ip限流',
  `service_flow_limit` int(20) NOT NULL DEFAULT '0' COMMENT '服务端限流',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '新增时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1 COMMENT='网关权限控制表'
******sql******/
// ServiceAccessControl 网关权限控制表
type ServiceAccessControl struct {
	gorm.Model
	ServiceID         uint   `gorm:"column:service_id;type:bigint(20);not null" json:"service_id"`                    // 服务id
	OpenAuth          int8   `gorm:"column:open_auth;type:tinyint(4);not null" json:"open_auth" validate:"oneof=0 1"` // 是否开启权限 1=开启
	BlackList         string `gorm:"column:black_list;type:varchar(1000);not null" json:"black_list"`                 // 黑名单ip
	WhiteList         string `gorm:"column:white_list;type:varchar(1000);not null" json:"white_list"`                 // 白名单ip
	WhiteHostName     string `gorm:"column:white_host_name;type:varchar(1000);not null" json:"white_host_name"`       // 白名单主机
	ClientipFlowLimit int    `gorm:"column:clientip_flow_limit;type:int(11);not null" json:"clientip_flow_limit"`     // 客户端ip限流
	ServiceFlowLimit  int    `gorm:"column:service_flow_limit;type:int(20);not null" json:"service_flow_limit"`       // 服务端限流
}

// TableName get sql table name.获取数据库表名
func (m *ServiceAccessControl) TableName() string {
	return "service_access_control"
}

/******sql******
CREATE TABLE `service_grpc_rule` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `service_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '服务id',
  `port` int(5) NOT NULL DEFAULT '0' COMMENT '端口',
  `header_transfor` varchar(5000) NOT NULL DEFAULT '' COMMENT 'header转换支持增加(add)、删除(del)、修改(edit) 格式: add headname headvalue 多个逗号间隔',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '新增时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COMMENT='网关路由匹配表'
******sql******/
// ServiceGrpcRule 网关路由匹配表
type ServiceGrpcRule struct {
	gorm.Model
	ServiceID         uint   `gorm:"column:service_id;type:bigint(20);not null" json:"service_id"`                    // 服务id
	Port              int    `gorm:"column:port;type:int(5);not null" json:"port"`                                    // 端口
	MetadataTransform string `gorm:"column:metadata_transform;type:varchar(5000);not null" json:"metadata_transform"` // header转换支持增加(add)、删除(del)、修改(edit) 格式: add headname headvalue 多个逗号间隔
}

// TableName get sql table name.获取数据库表名
func (m *ServiceGrpcRule) TableName() string {
	return "service_grpc_rule"
}

/******sql******
CREATE TABLE `service_http_rule` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `service_id` bigint(20) NOT NULL COMMENT '服务id',
  `rule_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '匹配类型 0=url前缀url_prefix 1=域名domain ',
  `rule` varchar(255) NOT NULL DEFAULT '' COMMENT 'type=domain表示域名，type=url_prefix时表示url前缀',
  `need_https` tinyint(4) NOT NULL DEFAULT '0' COMMENT '支持https 1=支持',
  `need_strip_uri` tinyint(4) NOT NULL DEFAULT '0' COMMENT '启用strip_uri 1=启用',
  `need_websocket` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否支持websocket 1=支持',
  `url_rewrite` varchar(5000) NOT NULL DEFAULT '' COMMENT 'url重写功能 格式：^/gatekeeper/test_service(.*) $1 多个逗号间隔',
  `header_transfor` varchar(5000) NOT NULL DEFAULT '' COMMENT 'header转换支持增加(add)、删除(del)、修改(edit) 格式: add headname headvalue 多个逗号间隔',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '新增时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=196 DEFAULT CHARSET=latin1 COMMENT='网关路由匹配表'
******sql******/
// ServiceHTTPRule 网关路由匹配表
type ServiceHTTPRule struct {
	gorm.Model
	ServiceID       uint   `gorm:"column:service_id;type:bigint(20);not null" json:"service_id"`                              // 服务id
	RuleType        int8   `gorm:"column:rule_type;type:tinyint(4);not null" json:"rule_type" validate:"oneof=0 1"`           // 匹配类型 0=url前缀url_prefix 1=域名domain
	Rule            string `gorm:"column:rule;type:varchar(255);not null" json:"rule"`                                        // type=domain表示域名，type=url_prefix时表示url前缀
	NeedHTTPs       int8   `gorm:"column:need_https;type:tinyint(4);not null" json:"need_https" validate:"oneof=0 1"`         // 支持https 1=支持
	NeedStripURI    int8   `gorm:"column:need_strip_uri;type:tinyint(4);not null" json:"need_strip_uri" validate:"oneof=0 1"` // 启用strip_uri 1=启用
	NeedWebsocket   int8   `gorm:"column:need_websocket;type:tinyint(4);not null" json:"need_websocket" validate:"oneof=0 1"` // 是否支持websocket 1=支持
	URLRewrite      string `gorm:"column:url_rewrite;type:varchar(5000);not null" json:"url_rewrite"`                         // url重写功能 格式：^/gatekeeper/test_service(.*) $1 多个逗号间隔
	HeaderTransform string `gorm:"column:header_transform;type:varchar(5000);not null" json:"header_transform"`               // header转换支持增加(add)、删除(del)、修改(edit) 格式: add headname headvalue 多个逗号间隔
}

// TableName get sql table name.获取数据库表名
func (m *ServiceHTTPRule) TableName() string {
	return "service_http_rule"
}

/******sql******
CREATE TABLE `service_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `load_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '负载类型 0=http 1=tcp 2=grpc',
  `service_name` varchar(255) NOT NULL DEFAULT '' COMMENT '服务名称 6-128 数字字母下划线',
  `service_desc` varchar(255) NOT NULL DEFAULT '' COMMENT '服务描述',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '新增时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=89 DEFAULT CHARSET=latin1 COMMENT='网关基本信息表'
******sql******/
// ServiceInfo 网关基本信息表
type ServiceInfo struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	LoadType    int8           `gorm:"column:load_type;type:tinyint(4);not null" json:"load_type" validate:"oneof=0 1 2"`           // 负载类型 0=http 1=tcp 2=grpc
	ServiceName string         `gorm:"column:service_name;type:varchar(255);not null" json:"service_name" validate:"min=6,max=128"` // 服务名称 6-128 数字字母下划线
	ServiceDesc string         `gorm:"column:service_desc;type:varchar(255);not null" json:"service_desc" validate:"required"`      // 服务描述
}

// TableName get sql table name.获取数据库表名
func (m *ServiceInfo) TableName() string {
	return "service_info"
}

/******sql******
CREATE TABLE `service_load_balance` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `service_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '服务id',
  `check_method` tinyint(20) NOT NULL DEFAULT '0' COMMENT '检查方法 0=tcpchk,检测端口是否握手成功',
  `check_timeout` int(10) NOT NULL DEFAULT '0' COMMENT 'check超时时间,单位s',
  `check_interval` int(11) NOT NULL DEFAULT '0' COMMENT '检查间隔, 单位s',
  `round_type` tinyint(4) NOT NULL DEFAULT '2' COMMENT '轮询方式 0=random 1=round-robin 2=weight_round-robin 3=ip_hash',
  `ip_list` varchar(2000) NOT NULL DEFAULT '' COMMENT 'ip列表',
  `weight_list` varchar(2000) NOT NULL DEFAULT '' COMMENT '权重列表',
  `forbid_list` varchar(2000) NOT NULL DEFAULT '' COMMENT '禁用ip列表',
  `upstream_connect_timeout` int(11) NOT NULL DEFAULT '0' COMMENT '建立连接超时, 单位s',
  `upstream_header_timeout` int(11) NOT NULL DEFAULT '0' COMMENT '获取header超时, 单位s',
  `upstream_idle_timeout` int(10) NOT NULL DEFAULT '0' COMMENT '链接最大空闲时间, 单位s',
  `upstream_max_idle` int(11) NOT NULL DEFAULT '0' COMMENT '最大空闲链接数',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '新增时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=200 DEFAULT CHARSET=latin1 COMMENT='网关负载表'
******sql******/
// ServiceLoadBalance 网关负载表
type ServiceLoadBalance struct {
	gorm.Model
	ServiceID              uint   `gorm:"column:service_id;type:bigint(20);not null" json:"service_id"`                          // 服务id
	CheckMethod            int8   `gorm:"column:check_method;type:tinyint(20);not null" json:"check_method"`                     // 检查方法 0=tcpchk,检测端口是否握手成功
	CheckTimeout           int    `gorm:"column:check_timeout;type:int(10);not null" json:"check_timeout"`                       // check超时时间,单位s
	CheckInterval          int    `gorm:"column:check_interval;type:int(11);not null" json:"check_interval"`                     // 检查间隔, 单位s
	RoundType              int8   `gorm:"column:round_type;type:tinyint(4);not null" json:"round_type" validate:"oneof=0 1 2 3"` // 轮询方式 0=random 1=round-robin 2=weight_round-robin 3=ip_hash
	IPList                 string `gorm:"column:ip_list;type:varchar(2000);not null" json:"ip_list"`                             // ip列表
	WeightList             string `gorm:"column:weight_list;type:varchar(2000);not null" json:"weight_list"`                     // 权重列表
	ForbidList             string `gorm:"column:forbid_list;type:varchar(2000);not null" json:"forbid_list"`                     // 禁用ip列表
	UpstreamConnectTimeout int    `gorm:"column:upstream_connect_timeout;type:int(11);not null" json:"upstream_connect_timeout"` // 建立连接超时, 单位s
	UpstreamHeaderTimeout  int    `gorm:"column:upstream_header_timeout;type:int(11);not null" json:"upstream_header_timeout"`   // 获取header超时, 单位s
	UpstreamIDleTimeout    int    `gorm:"column:upstream_idle_timeout;type:int(10);not null" json:"upstream_idle_timeout"`       // 链接最大空闲时间, 单位s
	UpstreamMaxIDle        int    `gorm:"column:upstream_max_idle;type:int(11);not null" json:"upstream_max_idle"`               // 最大空闲链接数
}

// TableName get sql table name.获取数据库表名
func (m *ServiceLoadBalance) TableName() string {
	return "service_load_balance"
}

/******sql******
CREATE TABLE `service_tcp_rule` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `service_id` bigint(20) NOT NULL COMMENT '服务id',
  `port` int(5) NOT NULL DEFAULT '0' COMMENT '端口号',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '新增时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COMMENT='网关路由匹配表'
******sql******/
// ServiceTCPRule 网关路由匹配表
type ServiceTCPRule struct {
	gorm.Model
	ServiceID uint `gorm:"column:service_id;type:bigint(20);not null" json:"service_id"` // 服务id
	Port      int  `gorm:"column:port;type:int(5);not null" json:"port"`                 // 端口号
}

// TableName get sql table name.获取数据库表名
func (m *ServiceTCPRule) TableName() string {
	return "service_tcp_rule"
}
