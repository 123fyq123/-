// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/admin/admin_info": {
            "get": {
                "description": "管理员信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员接口"
                ],
                "summary": "管理员信息",
                "operationId": "/admin/admin_info",
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/middleware.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.AdminInfoOutput"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/admin/change_pwd": {
            "post": {
                "description": "修改密码",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员接口"
                ],
                "summary": "修改密码",
                "operationId": "/admin/change_pwd",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ChangePwdInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/middleware.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/admin_login/login": {
            "post": {
                "description": "管理员登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员接口"
                ],
                "summary": "管理员登录",
                "operationId": "/admin_login/login",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.AdminLoginInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/middleware.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.AdminLoginOutput"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/admin_login/logout": {
            "get": {
                "description": "管理员退出",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员接口"
                ],
                "summary": "管理员退出",
                "operationId": "/admin_login/logout",
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/middleware.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/service/service_add_http": {
            "post": {
                "description": "添加HTTP服务",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "服务管理"
                ],
                "summary": "添加HTTP服务",
                "operationId": "/service/service_add_http",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ServiceAddHTTPInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/middleware.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/service/service_delete": {
            "get": {
                "description": "服务删除",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "服务管理"
                ],
                "summary": "服务删除",
                "operationId": "/service/service_delete",
                "parameters": [
                    {
                        "type": "string",
                        "description": "服务ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/middleware.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/service/service_list": {
            "get": {
                "description": "服务列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "服务管理"
                ],
                "summary": "服务列表",
                "operationId": "/service/service_list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "关键词",
                        "name": "info",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页个数",
                        "name": "page_size",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "当前页数",
                        "name": "page_no",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/middleware.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.ServiceListOutput"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/service/service_update_http": {
            "post": {
                "description": "修改HTTP服务",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "服务管理"
                ],
                "summary": "修改HTTP服务",
                "operationId": "/service/service_update_http",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ServiceUpdateHTTPInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/middleware.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.AdminInfoOutput": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "introduction": {
                    "type": "string"
                },
                "login_time": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "roles": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "dto.AdminLoginInput": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "description": "密码",
                    "type": "string",
                    "example": "123456"
                },
                "username": {
                    "description": "用户名",
                    "type": "string",
                    "example": "admin"
                }
            }
        },
        "dto.AdminLoginOutput": {
            "type": "object",
            "properties": {
                "token": {
                    "description": "token",
                    "type": "string",
                    "example": "token"
                }
            }
        },
        "dto.ChangePwdInput": {
            "type": "object",
            "required": [
                "password"
            ],
            "properties": {
                "password": {
                    "description": "密码",
                    "type": "string",
                    "example": "123456"
                }
            }
        },
        "dto.ServiceAddHTTPInput": {
            "type": "object",
            "required": [
                "ip_list",
                "rule",
                "service_desc",
                "service_name",
                "weight_list"
            ],
            "properties": {
                "black_list": {
                    "description": "黑名单ip",
                    "type": "string"
                },
                "clientip_flow_limit": {
                    "description": "客户端ip限流",
                    "type": "integer"
                },
                "header_transfor": {
                    "description": "header转换",
                    "type": "string"
                },
                "ip_list": {
                    "description": "ip列表",
                    "type": "string"
                },
                "need_https": {
                    "description": "支持https",
                    "type": "integer"
                },
                "need_strip_uri": {
                    "description": "启用strip_uri",
                    "type": "integer"
                },
                "need_websocket": {
                    "description": "是否支持websocket",
                    "type": "integer"
                },
                "open_auth": {
                    "description": "关键词",
                    "type": "integer"
                },
                "round_type": {
                    "description": "轮询方式",
                    "type": "integer"
                },
                "rule": {
                    "description": "域名或者前缀",
                    "type": "string"
                },
                "rule_type": {
                    "description": "接入类型",
                    "type": "integer"
                },
                "service_desc": {
                    "description": "服务描述",
                    "type": "string"
                },
                "service_flow_limit": {
                    "description": "服务端限流",
                    "type": "integer"
                },
                "service_name": {
                    "description": "服务名",
                    "type": "string"
                },
                "upstream_connect_timeout": {
                    "description": "建立连接超时, 单位s",
                    "type": "integer"
                },
                "upstream_header_timeout": {
                    "description": "获取header超时, 单位s",
                    "type": "integer"
                },
                "upstream_idle_timeout": {
                    "description": "链接最大空闲时间, 单位s",
                    "type": "integer"
                },
                "upstream_max_idle": {
                    "description": "最大空闲链接数",
                    "type": "integer"
                },
                "url_rewrite": {
                    "description": "url重写功能",
                    "type": "string"
                },
                "weight_list": {
                    "description": "权重列表",
                    "type": "string"
                },
                "white_list": {
                    "description": "白名单ip",
                    "type": "string"
                }
            }
        },
        "dto.ServiceListItemOutput": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "id",
                    "type": "integer"
                },
                "load_type": {
                    "description": "类型",
                    "type": "integer"
                },
                "qpd": {
                    "description": "qpd",
                    "type": "integer"
                },
                "qps": {
                    "description": "qps",
                    "type": "integer"
                },
                "service_addr": {
                    "description": "服务地址",
                    "type": "string"
                },
                "service_desc": {
                    "description": "服务描述",
                    "type": "string"
                },
                "service_name": {
                    "description": "服务名称",
                    "type": "string"
                },
                "total_node": {
                    "description": "节点数",
                    "type": "integer"
                }
            }
        },
        "dto.ServiceListOutput": {
            "type": "object",
            "properties": {
                "list": {
                    "description": "列表",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.ServiceListItemOutput"
                    }
                },
                "total": {
                    "description": "总数",
                    "type": "integer"
                }
            }
        },
        "dto.ServiceUpdateHTTPInput": {
            "type": "object",
            "required": [
                "id",
                "ip_list",
                "rule",
                "service_desc",
                "service_name",
                "weight_list"
            ],
            "properties": {
                "black_list": {
                    "description": "黑名单ip",
                    "type": "string"
                },
                "clientip_flow_limit": {
                    "description": "客户端ip限流",
                    "type": "integer"
                },
                "header_transfor": {
                    "description": "header转换",
                    "type": "string"
                },
                "id": {
                    "description": "服务ID",
                    "type": "integer",
                    "example": 62
                },
                "ip_list": {
                    "description": "ip列表",
                    "type": "string",
                    "example": "127.0.0.1:80"
                },
                "need_https": {
                    "description": "支持https",
                    "type": "integer"
                },
                "need_strip_uri": {
                    "description": "启用strip_uri",
                    "type": "integer"
                },
                "need_websocket": {
                    "description": "是否支持websocket",
                    "type": "integer"
                },
                "open_auth": {
                    "description": "关键词",
                    "type": "integer"
                },
                "round_type": {
                    "description": "轮询方式",
                    "type": "integer"
                },
                "rule": {
                    "description": "域名或者前缀",
                    "type": "string"
                },
                "rule_type": {
                    "description": "接入类型",
                    "type": "integer"
                },
                "service_desc": {
                    "description": "服务描述",
                    "type": "string"
                },
                "service_flow_limit": {
                    "description": "服务端限流",
                    "type": "integer"
                },
                "service_name": {
                    "description": "服务名",
                    "type": "string",
                    "example": "string"
                },
                "upstream_connect_timeout": {
                    "description": "建立连接超时, 单位s",
                    "type": "integer"
                },
                "upstream_header_timeout": {
                    "description": "获取header超时, 单位s",
                    "type": "integer"
                },
                "upstream_idle_timeout": {
                    "description": "链接最大空闲时间, 单位s",
                    "type": "integer"
                },
                "upstream_max_idle": {
                    "description": "最大空闲链接数",
                    "type": "integer"
                },
                "url_rewrite": {
                    "description": "url重写功能",
                    "type": "string"
                },
                "weight_list": {
                    "description": "权重列表",
                    "type": "string",
                    "example": "1"
                },
                "white_list": {
                    "description": "白名单ip",
                    "type": "string"
                }
            }
        },
        "middleware.Response": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "errmsg": {
                    "type": "string"
                },
                "errno": {
                    "type": "integer"
                },
                "stack": {
                    "type": "object"
                },
                "trace_id": {
                    "type": "object"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
