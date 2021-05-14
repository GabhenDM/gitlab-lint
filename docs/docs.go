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
        "contact": {
            "name": "API Support",
            "url": "https://github.com/globocom/gitlab-lint"
        },
        "license": {
            "name": "BSD-3-Clause License",
            "url": "https://opensource.org/licenses/BSD-3-Clause"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/levels": {
            "get": {
                "description": "get levels",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show levels",
                "operationId": "get-levels",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/api.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Data": {
                                            "type": "object",
                                            "additionalProperties": {
                                                "type": "integer"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/projects": {
            "get": {
                "description": "get all projects",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show projects",
                "operationId": "get-projects",
                "parameters": [
                    {
                        "type": "string",
                        "description": "fuzzy search projects",
                        "name": "q",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/api.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Data": {
                                            "$ref": "#/definitions/rules.Project"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/projects/{id}": {
            "get": {
                "description": "get project by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show project by id",
                "operationId": "get-projects-by-id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Project ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/api.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Data": {
                                            "type": "object",
                                            "additionalProperties": true
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/rules": {
            "get": {
                "description": "get all projects",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show rules",
                "operationId": "get-rules",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/api.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Data": {
                                            "$ref": "#/definitions/rules.Rule"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/rules/{id}": {
            "get": {
                "description": "get rule by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show rule by id",
                "operationId": "get-rules-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Rule ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/api.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Data": {
                                            "type": "object",
                                            "additionalProperties": true
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/stats": {
            "get": {
                "description": "get stats",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show stats",
                "operationId": "get-stats",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/api.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Data": {
                                            "$ref": "#/definitions/rules.Stats"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/status": {
            "get": {
                "description": "get status",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show status services",
                "operationId": "get-status",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Status"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.MetaResponse": {
            "type": "object",
            "properties": {
                "currentPage": {
                    "type": "integer"
                },
                "perPage": {
                    "type": "integer"
                },
                "totalOfItems": {
                    "type": "integer"
                },
                "totalOfPages": {
                    "type": "integer"
                }
            }
        },
        "api.Response": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "meta": {
                    "type": "object",
                    "$ref": "#/definitions/api.MetaResponse"
                }
            }
        },
        "api.Service": {
            "type": "object",
            "properties": {
                "elapsed": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "api.Status": {
            "type": "object",
            "properties": {
                "project": {
                    "type": "string"
                },
                "services": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.Service"
                    }
                },
                "status": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "rules.Project": {
            "type": "object",
            "properties": {
                "rules": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "integer"
                    }
                }
            }
        },
        "rules.Rule": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "level": {
                    "type": "string"
                },
                "namespaceId": {
                    "type": "integer"
                },
                "namespacePath": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                },
                "pathWithNamespace": {
                    "type": "string"
                },
                "projectId": {
                    "type": "integer"
                },
                "ruleId": {
                    "type": "string"
                },
                "webUrl": {
                    "type": "string"
                }
            }
        },
        "rules.Stats": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "gitlabProjectsCount": {
                    "type": "integer"
                },
                "levelsCount": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "integer"
                    }
                },
                "projectsCount": {
                    "type": "integer"
                },
                "registeredRulesCount": {
                    "type": "integer"
                },
                "rulesCount": {
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
	Version:     "0.1.0",
	Host:        "localhost:3000",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "gitlab-lint API",
	Description: "gitlab-lint API data",
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
