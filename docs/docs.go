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
        "/routesearch": {
            "get": {
                "description": "出発駅、到着駅、経由駅、出発時刻もしくは到着時刻を受け取り、経路を検索して返す",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "routes"
                ],
                "summary": "経路検索",
                "parameters": [
                    {
                        "type": "string",
                        "description": "出発駅のnodeID(例: 00004464)",
                        "name": "from",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "到着駅のnodeID(例: 00004254)",
                        "name": "to",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "日付と時刻（例: 2025-05-28T10:00）",
                        "name": "datetime",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "description": "datetimeが到着時刻かどうか(デフォルト: false)",
                        "name": "isArrivalTime",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": "経由地(nodeID)※複数指定可能",
                        "name": "via",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Route"
                        }
                    }
                }
            }
        },
        "/stations": {
            "get": {
                "description": "外部APIから取得した全駅一覧を返す",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "stations"
                ],
                "summary": "駅一覧の取得",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Station"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Route": {
            "type": "object",
            "properties": {
                "steps": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.RouteStep"
                    }
                },
                "total_fare": {
                    "type": "integer",
                    "example": 270
                },
                "total_time": {
                    "type": "integer",
                    "example": 30
                }
            }
        },
        "model.RouteStep": {
            "type": "object",
            "properties": {
                "arrival_station": {
                    "type": "string",
                    "example": "新川崎"
                },
                "departure_station": {
                    "type": "string",
                    "example": "新橋"
                },
                "from_time": {
                    "type": "string",
                    "example": "2025-05-28T10:01:00+09:00"
                },
                "movetype": {
                    "type": "string",
                    "example": "local_train"
                },
                "to_time": {
                    "type": "string",
                    "example": "2025-05-28T10:13:00+09:00"
                }
            }
        },
        "model.Station": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string",
                    "example": "urn:ucode:_00001C0000000000000100000341B4E7"
                },
                "name": {
                    "type": "string",
                    "example": "渋谷"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "NebosukeRouteAPI",
	Description:      "ねぼすけあんないのAPI",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
