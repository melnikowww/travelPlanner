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
        "/equip": {
            "get": {
                "description": "Получить список всего снаряжения или конкретный экземпляр по ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Equipment"
                ],
                "summary": "Получить снаряжение",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID снаряжения",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Список снаряжения",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Equipment"
                            }
                        }
                    },
                    "400": {
                        "description": "Неверный ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Снаряжение не найдено",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Создать новое снаряжение",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Equipment"
                ],
                "summary": "Создать снаряжение",
                "parameters": [
                    {
                        "description": "Данные снаряжения",
                        "name": "equipment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Equipment"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Созданное снаряжение",
                        "schema": {
                            "$ref": "#/definitions/models.Equipment"
                        }
                    },
                    "400": {
                        "description": "Ошибка валидации",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "409": {
                        "description": "Конфликт данных",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Удалить снаряжение по ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Equipment"
                ],
                "summary": "Удалить снаряжение",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID снаряжения",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Снаряжение удалено"
                    },
                    "400": {
                        "description": "Неверный ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "409": {
                        "description": "Ошибка удаления",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "description": "Частичное обновление данных снаряжения",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Equipment"
                ],
                "summary": "Обновить снаряжение",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID снаряжения",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "Обновляемые поля",
                        "name": "equipment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Equipment"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Обновленное снаряжение",
                        "schema": {
                            "$ref": "#/definitions/models.Equipment"
                        }
                    },
                    "400": {
                        "description": "Неверные данные",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Снаряжение не найдено",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "409": {
                        "description": "Ошибка обновления",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/points": {
            "get": {
                "description": "Получить все точки или конкретную точку по ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Points"
                ],
                "summary": "Получить точку/точки",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID точки",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешный ответ для одной точки",
                        "schema": {
                            "$ref": "#/definitions/models.Point"
                        }
                    },
                    "400": {
                        "description": "Неверный ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Точка не найдена",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Добавление новой географической точки",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Points"
                ],
                "summary": "Создать новую точку",
                "parameters": [
                    {
                        "description": "Данные точки",
                        "name": "point",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Point"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Созданная точка",
                        "schema": {
                            "$ref": "#/definitions/models.Point"
                        }
                    },
                    "400": {
                        "description": "Неверный формат данных",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "409": {
                        "description": "Конфликт данных",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Удаление точки по ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Points"
                ],
                "summary": "Удалить точку",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID точки",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Точка удалена"
                    },
                    "400": {
                        "description": "Неверный ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Точка не найдена",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "description": "Частичное или полное обновление данных точки",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Points"
                ],
                "summary": "Обновить точку",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID точки",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "Обновленные данные",
                        "name": "point",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Point"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Обновленная точка",
                        "schema": {
                            "$ref": "#/definitions/models.Point"
                        }
                    },
                    "400": {
                        "description": "Неверные данные",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Точка не найдена",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "409": {
                        "description": "Конфликт данных",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "Возвращает конкретного пользователя по ID или всех пользователей, если ID не указан",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Получить пользователя/пользователей",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID пользователя",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Данные пользователя",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Некорректный ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Пользователь не найден",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Регистрация нового пользователя в системе",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Создать пользователя",
                "parameters": [
                    {
                        "description": "Данные пользователя",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Созданный пользователь",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Некорректные входные данные",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "409": {
                        "description": "Пользователь уже существует",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Ошибка хеширования пароля",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Удаление пользователя по ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Удалить пользователя",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID пользователя",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Пользователь удален"
                    },
                    "400": {
                        "description": "Некорректный ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Пользователь не найден",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "description": "Частичное обновление данных пользователя. Пароль будет автоматически хеширован.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Обновить пользователя",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID пользователя",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "Обновляемые данные",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Обновленные данные",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Некорректные данные",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Пользователь не найден",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "409": {
                        "description": "Конфликт данных",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Equipment": {
            "type": "object",
            "properties": {
                "expedition_id": {
                    "type": "integer",
                    "example": 1
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "GPS Navigator"
                }
            }
        },
        "models.Point": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "location": {
                    "type": "string",
                    "example": "69.164529, 35.138287"
                },
                "name": {
                    "type": "string",
                    "example": "Teriberka"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "alex@example.com"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "Alex"
                },
                "password": {
                    "type": "string",
                    "example": "password"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8081",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Travel Planner",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
