// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/comment/createonecomment": {
            "post": {
                "description": "创建一个评论",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "评论相关接口"
                ],
                "summary": "创建一个评论",
                "parameters": [
                    {
                        "description": "评论请求体",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/post.Comment"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/api/comment/deletecomment/:id": {
            "get": {
                "description": "删除一个评论",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "评论相关接口"
                ],
                "summary": "删除一个评论",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "commentid",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/api/comment/deleteonepost/:id": {
            "get": {
                "description": "删除一个树洞",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "树洞相关接口"
                ],
                "summary": "删除一个树洞",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "postid",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/api/comment/getallcomment/:id": {
            "get": {
                "description": "打开树洞下面所有的评论",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "评论相关接口"
                ],
                "summary": "打开树洞下面所有的评论",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "postid",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/api/post/createonepost": {
            "post": {
                "description": "创建一个树洞",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "树洞相关接口"
                ],
                "summary": "创建一个树洞",
                "parameters": [
                    {
                        "description": "树洞请求体",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/post.Post"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/api/post/getallpost/:page": {
            "get": {
                "description": "查看所有树洞",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "树洞相关接口"
                ],
                "summary": "查看所有树洞",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "页数",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/api/post/getpostbyid": {
            "post": {
                "description": "通过userid查树洞",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "树洞相关接口"
                ],
                "summary": "通过userid查树洞",
                "parameters": [
                    {
                        "description": "请求体",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/post.PagePost"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/api/question": {
            "get": {
                "description": "查看注册问题，一次三个随机问题",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "树洞相关接口"
                ],
                "summary": "查看注册问题",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/api/user/banusers": {
            "post": {
                "description": "管理员封禁用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "管理员封禁用户",
                "parameters": [
                    {
                        "description": "封禁请求体",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.BanedRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/api/user/getusername": {
            "post": {
                "description": "通过id得到用户名称",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "得到用户名称",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/api/user/logincheck": {
            "post": {
                "description": "检查用户登陆，验证用户名和密码",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "验证用户名和密码",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/api/user/registercheck": {
            "post": {
                "description": "用户注册",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "用户注册",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/api/user/showbanedlist": {
            "get": {
                "description": "展示已经被ban的用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "展示已经被ban的用户",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.BanedRequest": {
            "type": "object",
            "properties": {
                "post_id": {
                    "type": "integer"
                },
                "reason": {
                    "type": "string"
                }
            }
        },
        "post.Comment": {
            "type": "object",
            "properties": {
                "Content": {
                    "type": "string"
                },
                "Page": {
                    "type": "integer"
                },
                "PostId": {
                    "type": "integer"
                },
                "UserId": {
                    "type": "integer"
                }
            }
        },
        "post.PagePost": {
            "type": "object",
            "properties": {
                "Page": {
                    "type": "integer"
                },
                "UserId": {
                    "type": "integer"
                }
            }
        },
        "post.Post": {
            "type": "object",
            "properties": {
                "Content": {
                    "type": "string"
                },
                "PostId": {
                    "type": "integer"
                },
                "UserId": {
                    "type": "integer"
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
	swag.Register("swagger", &s{})
}