// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
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
        "/api/v1/article/alternative": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Create alternative article for record by record_id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "article"
                ],
                "summary": "Create alternative article",
                "parameters": [
                    {
                        "description": "Article info for create",
                        "name": "article",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ArticleCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Article created",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Article not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "Unprocessable entity",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/article/{record_id}/all": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get articles for record by record_id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "article"
                ],
                "summary": "Get articles",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Record id",
                        "name": "record_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.ArticleDto"
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Articles not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "Unprocessable entity",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/article/{record_id}/is_main": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Set is_main for article by record_id and article_id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "article"
                ],
                "summary": "Set is_main",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Record id",
                        "name": "record_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Article id",
                        "name": "article_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Article set as main",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Article not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "Unprocessable entity",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/article/{record_id}/main": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get main article for record by record_id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "article"
                ],
                "summary": "Get main article",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Record id",
                        "name": "record_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Article",
                        "schema": {
                            "$ref": "#/definitions/models.ArticleDto"
                        }
                    },
                    "404": {
                        "description": "Article not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "Unprocessable entity",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/record/all": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get all records(not articles) for current user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "record"
                ],
                "summary": "Get records by user",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Records",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.RecordDto"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "Unprocessable entity",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/record/create": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Create record",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "record"
                ],
                "summary": "Create record",
                "parameters": [
                    {
                        "description": "Record create info",
                        "name": "record",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RecordCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Created record",
                        "schema": {
                            "$ref": "#/definitions/models.RecordDto"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "Unprocessable entity",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/record/published": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get all published records(not articles)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "record"
                ],
                "summary": "Get published records",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Records",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.RecordDto"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "Unprocessable entity",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/record/{record_id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get record by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "record"
                ],
                "summary": "Get record by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Record id",
                        "name": "record_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Record",
                        "schema": {
                            "$ref": "#/definitions/models.RecordDto"
                        }
                    },
                    "403": {
                        "description": "Hidden record",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Record not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "Unprocessable entity",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Delete record",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "record"
                ],
                "summary": "Delete record",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Record id",
                        "name": "record_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Record not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "Unprocessable entity",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/record/{record_id}/published_status": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Set published status",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "record"
                ],
                "summary": "Set published status",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Record id",
                        "name": "record_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "description": "Published status",
                        "name": "published",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Record published",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Record not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "Unprocessable entity",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/user/create": {
            "post": {
                "description": "Create user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Create user",
                "parameters": [
                    {
                        "description": "User create info",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserCreate"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created token for user",
                        "schema": {
                            "$ref": "#/definitions/models.TokenResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "Unprocessable entity",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/user/login": {
            "post": {
                "description": "Login user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "User login info",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Token",
                        "schema": {
                            "$ref": "#/definitions/models.TokenResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "Unprocessable entity",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.ArticleCreate": {
            "description": "Article create model",
            "type": "object",
            "required": [
                "body",
                "record_id"
            ],
            "properties": {
                "body": {
                    "type": "string",
                    "format": "html",
                    "example": "{html page}"
                },
                "record_id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "models.ArticleDto": {
            "description": "Article dto model",
            "type": "object",
            "required": [
                "body",
                "created_at",
                "id",
                "is_main"
            ],
            "properties": {
                "body": {
                    "type": "string",
                    "format": "html",
                    "example": "{html page}"
                },
                "created_at": {
                    "type": "string",
                    "example": "2021-01-01T00:00:00Z"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "is_main": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "models.RecordCreate": {
            "description": "Record create model",
            "type": "object",
            "required": [
                "settings",
                "video_link"
            ],
            "properties": {
                "settings": {
                    "$ref": "#/definitions/models.RecordSettings"
                },
                "video_link": {
                    "type": "string",
                    "format": "url",
                    "example": "https://www.youtube.com/watch?v=4O3UGW-Bbbw"
                }
            }
        },
        "models.RecordDto": {
            "description": "Record dto model",
            "type": "object",
            "required": [
                "id",
                "preview_picture",
                "published",
                "title",
                "video_link"
            ],
            "properties": {
                "annotation_length": {
                    "type": "integer",
                    "example": 200
                },
                "article_length": {
                    "type": "integer",
                    "example": 1000
                },
                "end_timecode": {
                    "type": "string",
                    "example": "00:10:00"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "preview_picture": {
                    "type": "string",
                    "format": "url",
                    "example": "picture url"
                },
                "published": {
                    "type": "boolean",
                    "example": false
                },
                "screenshot_timing": {
                    "type": "integer",
                    "example": 3
                },
                "start_timecode": {
                    "type": "string",
                    "example": "00:00:00"
                },
                "title": {
                    "type": "string",
                    "example": "title"
                },
                "video_link": {
                    "type": "string",
                    "format": "url",
                    "example": "https://www.youtube.com/watch?v=4O3UGW-Bbbw"
                }
            }
        },
        "models.RecordSettings": {
            "type": "object",
            "properties": {
                "annotation_length": {
                    "type": "integer",
                    "example": 200
                },
                "article_length": {
                    "type": "integer",
                    "example": 1000
                },
                "end_timecode": {
                    "type": "string",
                    "example": "00:10:00"
                },
                "screenshot_timing": {
                    "type": "integer",
                    "example": 3
                },
                "start_timecode": {
                    "type": "string",
                    "example": "00:00:00"
                }
            }
        },
        "models.TokenResponse": {
            "description": "Token response model",
            "type": "object",
            "required": [
                "token",
                "token_type"
            ],
            "properties": {
                "token": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
                },
                "token_type": {
                    "type": "string",
                    "example": "Bearer"
                }
            }
        },
        "models.UserCreate": {
            "description": "User create model",
            "type": "object",
            "required": [
                "email",
                "fio",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "format": "email",
                    "example": "test@test.ru"
                },
                "fio": {
                    "type": "string",
                    "example": "Мисосов Мисос Мисосович"
                },
                "password": {
                    "type": "string",
                    "example": "test"
                }
            }
        },
        "models.UserLogin": {
            "description": "User login model",
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "format": "email",
                    "example": "test@test.ru"
                },
                "password": {
                    "type": "string",
                    "example": "test"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "localhost:8000",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Profbuh API",
	Description:      "This is a sample server for Profbuh API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	// LeftDelim:        "{{",
	// RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
