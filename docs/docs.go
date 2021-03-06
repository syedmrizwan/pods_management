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
            "email": "syedmrizwan@outlook.com"
        },
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/create_pod": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API"
                ],
                "summary": "Create Pod",
                "parameters": [
                    {
                        "description": "description",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.PodBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Pod"
                        }
                    }
                }
            }
        },
        "/api/v1/delete_pod": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API"
                ],
                "summary": "Delete Pods based on user request",
                "parameters": [
                    {
                        "description": "description",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "integer"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/update_pod": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API"
                ],
                "summary": "Update Pod Configurtion",
                "parameters": [
                    {
                        "description": "description",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.PodConfiguration"
                            }
                        }
                    }
                ]
            }
        }
    },
    "definitions": {
        "model.Cluster": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "vcenter_id": {
                    "type": "integer"
                }
            }
        },
        "model.ConfigurationInfo": {
            "type": "object",
            "properties": {
                "cluster_id": {
                    "type": "integer"
                },
                "cluster_name": {
                    "type": "string"
                },
                "datastore_id": {
                    "type": "integer"
                },
                "datastore_name": {
                    "type": "string"
                },
                "ip_address": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "template_name": {
                    "type": "string"
                },
                "type_name": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                },
                "vcenter_id": {
                    "type": "integer"
                }
            }
        },
        "model.Datastore": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "vcenter_id": {
                    "type": "integer"
                }
            }
        },
        "model.Pod": {
            "type": "object",
            "properties": {
                "cluster": {
                    "type": "object",
                    "$ref": "#/definitions/model.Cluster"
                },
                "cluster_id": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "datastore": {
                    "type": "object",
                    "$ref": "#/definitions/model.Datastore"
                },
                "datastore_id": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "ip_address": {
                    "type": "string"
                },
                "is_expired": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "student": {
                    "type": "object",
                    "$ref": "#/definitions/model.Student"
                },
                "student_id": {
                    "type": "integer"
                },
                "subscription_type_id": {
                    "type": "integer"
                },
                "task_id": {
                    "type": "integer"
                }
            }
        },
        "model.PodBody": {
            "type": "object",
            "properties": {
                "ip_address": {
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
        "model.PodConfiguration": {
            "type": "object",
            "properties": {
                "configuration": {
                    "type": "object",
                    "$ref": "#/definitions/model.ConfigurationInfo"
                },
                "pod_id": {
                    "type": "integer"
                },
                "pod_name": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "model.RefType": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "trainingContents": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.TrainingContent"
                    }
                },
                "type_name": {
                    "type": "string"
                },
                "vapp_template_name": {
                    "type": "string"
                }
            }
        },
        "model.Response": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "model.Root": {
            "type": "object",
            "properties": {
                "account_id": {
                    "type": "integer"
                }
            }
        },
        "model.Student": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "tenant": {
                    "type": "object",
                    "$ref": "#/definitions/model.Tenant"
                },
                "tenant_id": {
                    "type": "integer"
                }
            }
        },
        "model.SubscriptionType": {
            "type": "object",
            "properties": {
                "expiry_time": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "pods": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Pod"
                    }
                },
                "ref_type": {
                    "type": "object",
                    "$ref": "#/definitions/model.RefType"
                },
                "ref_type_id": {
                    "type": "integer"
                },
                "tenant": {
                    "type": "object",
                    "$ref": "#/definitions/model.Tenant"
                },
                "tenant_id": {
                    "type": "integer"
                }
            }
        },
        "model.Tenant": {
            "type": "object",
            "properties": {
                "activate_later": {
                    "type": "boolean"
                },
                "activation_time": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "root": {
                    "type": "object",
                    "$ref": "#/definitions/model.Root"
                },
                "root_account_id": {
                    "type": "integer"
                },
                "subscriptionTypes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.SubscriptionType"
                    }
                }
            }
        },
        "model.TrainingContent": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "ref_type_id": {
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
	Version:     "1.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Pods Management",
	Description: "Pods Management Blueprint.",
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
