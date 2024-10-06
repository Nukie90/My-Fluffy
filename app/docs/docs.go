// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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

        "/payments": {
            "post": {
                "description": "Create a new payment",
                "consumes": [
                    "application/json"
                ],

        "/notifications": {
            "get": {
                "description": "Get all notifications for current user",

                "produces": [
                    "application/json"
                ],
                "tags": [

                    "payments"
                ],
                "summary": "Create a new payment",
                "parameters": [
                    {
                        "description": "Payment information",
                        "name": "payment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreatePayment"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Payment created successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/payments/user/{userID}": {
            "get": {
                "description": "Get payments from specific user",
                "consumes": [
                    "application/json"
                ],

                    "notifications"
                ],
                "summary": "Get all notifications",
                "responses": {
                    "200": {
                        "description": "List of notifications",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Notification"
                            }
                        }
                    }
                }
            }
        },
        "/notifications/{id}": {
            "delete": {
                "description": "Delete a notification",

                "produces": [
                    "application/json"
                ],
                "tags": [

                    "payments"
                ],
                "summary": "Get payments from specific user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "userID",

                    "notifications"
                ],
                "summary": "Delete a notification",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Notification ID",
                        "name": "id",

                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {

                        "description": "Payments from specific user",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Payment"
                            }

                        "description": "Notification deleted",
                        "schema": {
                            "type": "string"

                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/posts": {
            "post": {
                "description": "Create a new post",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "Create a new post",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Post title",
                        "name": "title",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Post content",
                        "name": "content",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Post picture",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Post reward",
                        "name": "reward",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Post created successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/posts/feed": {
            "get": {
                "description": "Get paginated posts for the feed, 10 at a time",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "Get paginated posts",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Paginated posts",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Post"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/posts/found": {
            "put": {
                "description": "Found pet",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "Found pet",
                "parameters": [
                    {
                        "description": "Post ID",
                        "name": "id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.FoundPost"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Pet found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/posts/user": {
            "get": {
                "description": "Get posts from specific user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "Get posts from specific user",
                "responses": {
                    "200": {
                        "description": "Posts from specific user",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/savedPost/saved_posts": {
            "post": {
                "description": "Create a new saved post",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "saved_posts"
                ],
                "summary": "Create a new saved post",
                "parameters": [
                    {
                        "description": "Saved post information",
                        "name": "saved_post",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SavedPost"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Saved post created successfully",
                        "schema": {
                            "$ref": "#/definitions/model.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/saved_posts": {
            "get": {
                "description": "Get all saved posts by user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "saved_posts"
                ],
                "summary": "Get all saved posts by user",
                "responses": {
                    "200": {
                        "description": "Saved posts with details",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Post"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users": {
            "post": {
                "description": "Create a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "Signup information",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Signup"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User created successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users/all": {
            "get": {
                "description": "Get all users",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get all users",
                "responses": {
                    "200": {
                        "description": "List of users",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.User"
                            }
                        }
                    }
                }
            }
        },
        "/users/login": {
            "post": {
                "description": "Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Login information",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Login successful",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "description": "Get user by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get user by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User details",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {

        "model.CreatePayment": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "transaction": {
                    "type": "string"
                },
                "user_id": {

        "model.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {

                    "type": "string"
                }
            }
        },
        "model.FoundPost": {
            "type": "object",
            "properties": {
                "found_id": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "model.Login": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string",
                    "example": "password"
                },
                "username": {
                    "type": "string",
                    "example": "john_doe"
                }
            }
        },

        "model.Payment": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "id": {
                    "type": "string"
                },
                "transaction": {
                    "type": "string"
                },

        "model.Notification": {
            "type": "object",
            "properties": {
                "create_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "owner_id": {
                    "type": "string"
                }
            }
        },
        "model.Post": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "found_id": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "owner_id": {
                    "type": "string"
                },
                "picture": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "reward": {
                    "type": "number"
                },
                "status": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.SavedPost": {
            "type": "object",
            "properties": {
                "post_id": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "model.Signup": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string",
                    "example": "password"
                },
                "role": {
                    "type": "string",
                    "example": ""
                },
                "username": {
                    "type": "string",
                    "example": "john_doe"
                }
            }
        },
        "model.SuccessResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "SDA My Fluffy API",
	Description:      "This is the API documentation for the SDA My Fluffy API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
