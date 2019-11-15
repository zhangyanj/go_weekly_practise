// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-11-15 15:33:59.55595 +0800 CST m=+0.056520227

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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/test": {
            "post": {
                "description": "request和response的示例",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "snow"
                ],
                "summary": "request和response的示例",
                "parameters": [
                    {
                        "description": "test request",
                        "name": "test",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/entities.TestRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entities.TestResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controllers.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.HTTPError"
                        }
                    }
                }
            }
        },
        "/test_validator": {
            "post": {
                "description": "HandleTestValidator的示例",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "snow"
                ],
                "summary": "HandleTestValidator的示例",
                "parameters": [
                    {
                        "description": "example of validator",
                        "name": "testValidator",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/entities.TestValidatorRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entities.TestValidatorRequest"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controllers.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "example": "status bad request"
                }
            }
        },
        "entities.Address": {
            "type": "object",
            "required": [
                "city",
                "phone",
                "planet",
                "street"
            ],
            "properties": {
                "city": {
                    "type": "string",
                    "example": "xiamen"
                },
                "phone": {
                    "type": "string",
                    "example": "snow"
                },
                "planet": {
                    "type": "string",
                    "example": "snow"
                },
                "street": {
                    "type": "string",
                    "example": "huandaodonglu"
                }
            }
        },
        "entities.TestRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "snow"
                },
                "url": {
                    "type": "string",
                    "example": "github.com/qit-team/snow"
                }
            }
        },
        "entities.TestResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "snow"
                },
                "url": {
                    "type": "string",
                    "example": "github.com/qit-team/snow"
                }
            }
        },
        "entities.TestValidatorRequest": {
            "type": "object",
            "required": [
                "age",
                "email",
                "id",
                "mobile",
                "name",
                "test_num",
                "url"
            ],
            "properties": {
                "addresses": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entities.Address"
                    }
                },
                "age": {
                    "type": "integer",
                    "example": 20
                },
                "content": {
                    "type": "string",
                    "example": "snow"
                },
                "email": {
                    "type": "string",
                    "example": "snow@github.com"
                },
                "id": {
                    "description": "tips，因为组件required不管是没传值或者传 0 or \"\" 都通过不了，但是如果用指针类型，那么0就是0，而nil无法通过校验",
                    "type": "integer",
                    "example": 1
                },
                "mobile": {
                    "type": "string",
                    "example": "snow"
                },
                "name": {
                    "type": "string",
                    "example": "snow"
                },
                "range_num": {
                    "type": "integer",
                    "example": 3
                },
                "test_num": {
                    "type": "integer",
                    "example": 7
                },
                "url": {
                    "type": "string",
                    "example": "github.com/qit-team/snow"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        },
        "BasicAuth": {
            "type": "basic"
        },
        "OAuth2AccessCode": {
            "type": "oauth2",
            "flow": "accessCode",
            "authorizationUrl": "https://example.com/oauth/authorize",
            "tokenUrl": "https://example.com/oauth/token",
            "scopes": {
                "admin": " Grants read and write access to administrative information"
            }
        },
        "OAuth2Application": {
            "type": "oauth2",
            "flow": "application",
            "tokenUrl": "https://example.com/oauth/token",
            "scopes": {
                "admin": " Grants read and write access to administrative information",
                "write": " Grants write access"
            }
        },
        "OAuth2Implicit": {
            "type": "oauth2",
            "flow": "implicit",
            "authorizationUrl": "https://example.com/oauth/authorize",
            "scopes": {
                "admin": " Grants read and write access to administrative information",
                "write": " Grants write access"
            }
        },
        "OAuth2Password": {
            "type": "oauth2",
            "flow": "password",
            "tokenUrl": "https://example.com/oauth/token",
            "scopes": {
                "admin": " Grants read and write access to administrative information",
                "read": " Grants read access",
                "write": " Grants write access"
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
	Version:     "1.0",
	Host:        "localhost:8080",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Swagger Example API",
	Description: "This is a sample server celler server.",
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
