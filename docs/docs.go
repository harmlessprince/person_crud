// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "email": "realolamilekan@gmail.io"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "get list of all person resource",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Person CRUD"
                ],
                "summary": "List Person",
                "operationId": "index-person",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name to filter person by",
                        "name": "name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.JSONSuccessResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/controllers.PersonDto"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.JSONErrorResult"
                        }
                    }
                }
            },
            "post": {
                "description": "store a new person",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Person CRUD"
                ],
                "summary": "Store User",
                "operationId": "store-person",
                "parameters": [
                    {
                        "description": "Person object to create",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.StorePersonBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.JSONSuccessResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/controllers.PersonDto"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/utils.ValidationError"
                            }
                        }
                    }
                }
            }
        },
        "/{user_id}": {
            "get": {
                "description": "Update person resource by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Person CRUD"
                ],
                "summary": "Update Person",
                "operationId": "show-person",
                "parameters": [
                    {
                        "type": "string",
                        "description": "person ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Person object to update",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.UpdatePersonBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.JSONSuccessResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/controllers.PersonDto"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/utils.JSONErrorResult"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/utils.ValidationError"
                            }
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete Person resource by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Person CRUD"
                ],
                "summary": "Delete Person",
                "operationId": "delete-person",
                "parameters": [
                    {
                        "type": "string",
                        "description": "person ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.JSONSuccessResult"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/utils.JSONErrorResult"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.PersonDto": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "2023-09-11T16:23:57.832758+01:00"
                },
                "deleted_at": {
                    "type": "string",
                    "example": "2023-09-11T16:23:57.832758+01:00"
                },
                "id": {
                    "type": "string",
                    "example": "1"
                },
                "name": {
                    "type": "string",
                    "example": "Mark Essien"
                },
                "updated_at": {
                    "type": "string",
                    "example": "2023-09-11T16:23:57.832758+01:00"
                }
            }
        },
        "controllers.StorePersonBody": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "example": "Mark Essien"
                }
            }
        },
        "controllers.UpdatePersonBody": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "Mark Essien"
                }
            }
        },
        "utils.JSONErrorResult": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "utils.JSONSuccessResult": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "utils.ValidationError": {
            "type": "object",
            "properties": {
                "condition": {
                    "type": "string",
                    "example": "required"
                },
                "error": {
                    "type": "string",
                    "example": "Name field is required"
                },
                "key": {
                    "type": "string",
                    "example": "name"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Swagger Person Crud API",
	Description:      "\"This is a crud server for creating resource.\"",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
