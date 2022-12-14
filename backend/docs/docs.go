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
            "name": "Hema Mathiyazhagan",
            "email": "mhema2795@gmail.com"
        },
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/albums": {
            "post": {
                "description": "creates new album if not present",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "albums"
                ],
                "summary": "create a album",
                "parameters": [
                    {
                        "description": "Add Album",
                        "name": "album",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AlbumJSON"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.SuccessMessage"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/albums/{albumName}": {
            "delete": {
                "description": "deletes the specified album if present",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "albums"
                ],
                "summary": "delete a album",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Delete Album",
                        "name": "albumName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.SuccessMessage"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorMessage"
                        }
                    }
                }
            },
        },
        "/albums/{albumName}/images": {
            "get": {
                "description": "returns list of images of the specified album",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "images"
                ],
                "summary": "list of images",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Album name",
                        "name": "albumName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ListImageJSON"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorMessage"
                        }
                    }
                }
            },
            "post": {
                "description": "adds an image to the specified album",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "images"
                ],
                "summary": "add a image",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Album name",
                        "name": "albumName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Add Image",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.SuccessMessage"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/albums/{albumName}/images/{imageName}": {
            "get": {
                "description": "get the specified image from the server",
                "produces": [
                    "application/octet-stream"
                ],
                "tags": [
                    "images"
                ],
                "summary": "get a image",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Album name",
                        "name": "albumName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Image name",
                        "name": "imageName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "file"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorMessage"
                        }
                    }
                }
            },
            "delete": {
                "description": "deletes a image from the specified album",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "images"
                ],
                "summary": "delete a image",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Album name",
                        "name": "albumName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Image name",
                        "name": "imageName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.SuccessMessage"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorMessage"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.AlbumJSON": {
            "type": "object",
            "required": [
                "albumName"
            ],
            "properties": {
                "albumName": {
                    "type": "string"
                }
            }
        },
        "model.ErrorMessage": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "model.ImageJSON": {
            "type": "object",
            "required": [
                "imageName"
            ],
            "properties": {
                "imageName": {
                    "type": "string"
                }
            }
        },
        "model.ListAlbumJSON": {
            "type": "object",
            "properties": {
                "albumList": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "model.ListImageJSON": {
            "type": "object",
            "properties": {
                "imageList": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "model.SuccessMessage": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
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