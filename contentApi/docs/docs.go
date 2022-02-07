// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate_swagger = `{
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
        "/books/{userid}/{pageno}": {
            "get": {
                "description": "finds and returns books",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Books"
                ],
                "summary": "finds and returns books",
                "parameters": [
                    {
                        "type": "string",
                        "description": "userid",
                        "name": "userid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "pageno",
                        "name": "pageno",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.RespSuccessBooks"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/main.RespError"
                        }
                    }
                }
            }
        },
        "/csv/{userid}": {
            "post": {
                "description": "parse csv file and update data to the database",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "csv"
                ],
                "summary": "parse csv file and update data to the database",
                "parameters": [
                    {
                        "type": "string",
                        "description": "userid",
                        "name": "userid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "file",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.RespSuccess"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/main.RespError"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/main.RespError"
                        }
                    }
                }
            }
        },
        "/getmostliked/{userid}": {
            "get": {
                "description": "finds and returns top content on the basis of number of likes",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TopContent"
                ],
                "summary": "finds and returns top content",
                "parameters": [
                    {
                        "type": "string",
                        "description": "userid",
                        "name": "userid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.RespSuccessML"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/main.RespError"
                        }
                    }
                }
            }
        },
        "/like/{cmd}/{userid}/{bookid}": {
            "post": {
                "description": "Takes the add/remove command from url and updates the like for the respective user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "like"
                ],
                "summary": "Used to update like",
                "parameters": [
                    {
                        "type": "string",
                        "description": "userid",
                        "name": "userid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "bookid",
                        "name": "bookid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "add or remove command",
                        "name": "cmd",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.RespSuccess"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/main.RespError"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/main.RespError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.RespError": {
            "type": "object",
            "required": [
                "message"
            ],
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Error"
                }
            }
        },
        "main.RespSuccess": {
            "type": "object",
            "required": [
                "message"
            ],
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Data Updated"
                }
            }
        },
        "main.RespSuccessBooks": {
            "type": "object",
            "required": [
                "books"
            ],
            "properties": {
                "books": {
                    "$ref": "#/definitions/main.ResponseWithCount"
                }
            }
        },
        "main.RespSuccessML": {
            "type": "object",
            "required": [
                "mostLiked"
            ],
            "properties": {
                "mostLiked": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.Response"
                    }
                }
            }
        },
        "main.Response": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string",
                    "example": "09/09/2001"
                },
                "id": {
                    "type": "string",
                    "example": "507f191e810c19729de860ea"
                },
                "liked": {
                    "type": "boolean",
                    "example": false
                },
                "likes": {
                    "type": "integer",
                    "example": 10
                },
                "story": {
                    "type": "string",
                    "example": "Test Story"
                },
                "title": {
                    "type": "string",
                    "example": "Test Title"
                }
            }
        },
        "main.ResponseWithCount": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer",
                    "example": 20
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.Response"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo_swagger holds exported Swagger Info so clients can modify it
var SwaggerInfo_swagger = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate_swagger,
}

func init() {
	swag.Register(SwaggerInfo_swagger.InstanceName(), SwaggerInfo_swagger)
}