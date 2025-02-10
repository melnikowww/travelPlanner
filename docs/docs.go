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
        "/cars": {
            "get": {
                "description": "Получить список всех автомобилей или конкретный автомобиль по ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Автомобили"
                ],
                "summary": "Получить автомобиль(и)",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Идентификатор автомобиля (необязательно)",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешное получение автомобиля",
                        "schema": {
                            "$ref": "#/definitions/models.Car"
                        }
                    },
                    "400": {
                        "description": "Некорректный запрос",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Автомобиль не найден",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Добавить новый автомобиль в базу данных",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Автомобили"
                ],
                "summary": "Создать автомобиль",
                "parameters": [
                    {
                        "description": "Данные автомобиля",
                        "name": "car",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Car"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Автомобиль успешно создан",
                        "schema": {
                            "$ref": "#/definitions/models.Car"
                        }
                    },
                    "400": {
                        "description": "Некорректные входные данные",
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
                "description": "Удаление автомобиля по идентификатору",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Автомобили"
                ],
                "summary": "Удалить автомобиль",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Идентификатор автомобиля",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Автомобиль успешно удален"
                    },
                    "400": {
                        "description": "Некорректный ID",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "description": "Частичное обновление данных автомобиля",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Автомобили"
                ],
                "summary": "Обновить автомобиль",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Идентификатор автомобиля",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "Данные для обновления",
                        "name": "car",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Car"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Автомобиль успешно обновлен",
                        "schema": {
                            "$ref": "#/definitions/models.Car"
                        }
                    },
                    "400": {
                        "description": "Некорректный ID или данные",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Автомобиль не найден",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/equip": {
            "get": {
                "description": "Получить список всего снаряжения или конкретный экземпляр по ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Снаряжение"
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
                    "Снаряжение"
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
                    "Снаряжение"
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
                    "Снаряжение"
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
        "/goods": {
            "get": {
                "description": "Получение списка всех продуктов или конкретной позиции по ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Продукты"
                ],
                "summary": "Получить список продуктов",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID продукта",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Конкретная позиция",
                        "schema": {
                            "$ref": "#/definitions/models.Good"
                        }
                    },
                    "400": {
                        "description": "Некорректный запрос",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Не найдено",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Ошибка сервера",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Создание новой позиции продуктов",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Продукты"
                ],
                "summary": "Создать новую позицию",
                "parameters": [
                    {
                        "description": "Данные продукта",
                        "name": "good",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Good"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Созданная позиция",
                        "schema": {
                            "$ref": "#/definitions/models.Good"
                        }
                    },
                    "400": {
                        "description": "Некорректные данные",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "409": {
                        "description": "Конфликт - позиция уже существует",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Удаление позиции продукта по ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Продукты"
                ],
                "summary": "Удалить позицию",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID продукта",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Успешное удаление"
                    },
                    "400": {
                        "description": "Некорректный ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Не найдено",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Ошибка сервера",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "description": "Обновление существующей позиции (частичное обновление)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Продукты"
                ],
                "summary": "Обновить позицию",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID продукта",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "Данные для обновления",
                        "name": "good",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Good"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Обновленная позиция",
                        "schema": {
                            "$ref": "#/definitions/models.Good"
                        }
                    },
                    "400": {
                        "description": "Некорректный запрос",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Не найдено",
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
        "/points": {
            "get": {
                "description": "Получить все точки или конкретную точку по ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Контрольные точки"
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
                    "Контрольные точки"
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
                    "Контрольные точки"
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
                    "Контрольные точки"
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
                    "Пользователи"
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
                    "Пользователи"
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
                    "Пользователи"
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
                    "Пользователи"
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
        "models.Car": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "Toyota Land Cruiser 100"
                },
                "owner_id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
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
        "models.Good": {
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
                    "example": "First Aid Kit"
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
