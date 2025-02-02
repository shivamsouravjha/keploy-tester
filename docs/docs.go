// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/events": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Fetch event logs from the last 2 hours",
                "produces": [
                    "application/json"
                ],
                "summary": "Get active event logs",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.EventLog"
                            }
                        }
                    }
                }
            }
        },
        "/events/archived": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Fetch logs that are older than 2 hours but still within 48 hours",
                "produces": [
                    "application/json"
                ],
                "summary": "Get archived event logs",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.EventLog"
                            }
                        }
                    }
                }
            }
        },
        "/events/purge": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Deletes logs older than 48 hours",
                "summary": "Purge old event logs",
                "responses": {
                    "200": {
                        "description": "Old events purged",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Authenticates a user and returns a JWT token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "description": "User Login Credentials",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.AuthInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Creates a new user with a username and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "User Registration Data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.AuthInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/triggers": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Fetch all stored triggers",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Triggers"
                ],
                "summary": "Get all triggers",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Trigger"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Creates a scheduled or API trigger",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Triggers"
                ],
                "summary": "Create a new trigger",
                "parameters": [
                    {
                        "description": "Trigger object",
                        "name": "trigger",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Trigger"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Trigger"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/triggers/test/api": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "This endpoint sends an API request with a **test payload** to a specified endpoint without saving it as a trigger.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Testing API"
                ],
                "summary": "Test a one-time API trigger",
                "parameters": [
                    {
                        "description": "API endpoint and payload to test",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.TriggerRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Trigger executed successfully",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Error in executing API trigger",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/triggers/test/scheduled": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "This endpoint allows users to test a scheduled event trigger **without saving it permanently**.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Testing API"
                ],
                "summary": "Test a one-time scheduled trigger",
                "parameters": [
                    {
                        "description": "Request body for scheduled trigger test",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.ScheduledTriggerTestRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Trigger executed successfully",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/triggers/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Fetch a trigger by ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Get a specific trigger",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Trigger ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Trigger"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Modify an existing trigger",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update a trigger",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Trigger ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated Trigger object",
                        "name": "trigger",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Trigger"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Trigger"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Removes a trigger from the system",
                "summary": "Delete a trigger",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Trigger ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Trigger deleted",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/triggers/{id}/execute": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Triggers an event immediately for testing",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Triggers"
                ],
                "summary": "Manually execute a trigger",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Trigger ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.AuthInput": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "controllers.ScheduledTriggerTestRequest": {
            "type": "object",
            "properties": {
                "delay": {
                    "description": "Delay in minutes",
                    "type": "integer"
                }
            }
        },
        "controllers.TriggerRequest": {
            "type": "object",
            "properties": {
                "endpoint": {
                    "type": "string"
                },
                "payload": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                }
            }
        },
        "models.EventLog": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "payload": {
                    "type": "object"
                },
                "status": {
                    "description": "\"active\", \"archived\"",
                    "type": "string"
                },
                "triggerID": {
                    "type": "string"
                },
                "triggeredAt": {
                    "type": "string"
                },
                "type": {
                    "description": "\"scheduled\" or \"api\"",
                    "type": "string"
                }
            }
        },
        "models.Trigger": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "endpoint": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "one_time": {
                    "type": "boolean"
                },
                "payload": {
                    "type": "object"
                },
                "schedule": {
                    "type": "string"
                },
                "type": {
                    "description": "\"scheduled\" or \"api\"",
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:4000",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Event Trigger API",
	Description:      "This API allows users to create and manage event triggers",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
