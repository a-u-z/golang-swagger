package main

type HttpResp struct {
	Code    int         `json:"code" example:"200"`
	Message string      `json:"message" example:"success"`
	Data    interface{} `json:"data"`
}
