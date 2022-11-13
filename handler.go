package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type helloReq struct {
	Letter string `form:"letter" binding:"required" example:"a, b, or c"`
}

// hello godoc
// @Summary <在不展開 api 的時候就可以看到的描述>
// @Description <在展開的時候才會出現，此 api 詳細的描述>
// @Description <在展開的時候才會出現，此 api 詳細的描述>
// @ID hello
// @Tags <這邊可以幫 api 新增標籤，做出分類的效果>
// @Accept  json
// @Produce  json
// @Param authorization  header string true "type bearer"
// @Param letter query string true "letter"
// @Success 200 {object} HttpResp
// @Failure 400 {object} HttpResp
// @Router /swagger-example/hello [get]
func hello(c *gin.Context) {
	var req helloReq
	// get => req 的 struct 的補充說明要用 form, Param 要用 query, c.shouldBindQuery
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, HttpResp{
			Code:    http.StatusBadRequest,
			Message: "param type is wrong",
		})
		return
	}

	resp := "hello, "
	switch req.Letter {
	case "a":
		resp += "a"
	case "b":
		resp += "b"
	case "c":
		resp += "c"
	default:
	}

	c.JSON(http.StatusOK, HttpResp{
		Code:    http.StatusOK,
		Message: resp,
	})
}

type echoReq struct {
	Word string `json:"word" binding:"required"`
}

// echo godoc
// @Summary <在不展開 api 的時候就可以看到的描述>
// @Description <在展開的時候才會出現，此 api 詳細的描述>
// @Description <在展開的時候才會出現，此 api 詳細的描述>
// @ID echo
// @Tags <這邊可以幫 api 新增標籤，做出分類的效果>
// @Accept  json
// @Produce  json
// @Param authorization  header string true "type bearer"
// @Param word body echoReq true "word"
// @Success 200 {object} HttpResp
// @Failure 400 {object} HttpResp
// @Router /swagger-example/echo [post]
func echo(c *gin.Context) {

	// post => req 的 struct 的補充說明要用 json, Param 要用 body, c.ShouldBindJSON
	var er echoReq
	if err := c.ShouldBindJSON(&er); err != nil {
		c.JSON(http.StatusBadRequest, HttpResp{
			Code:    http.StatusBadRequest,
			Message: "param type is wrong",
		})
		return
	}
	c.JSON(http.StatusOK, HttpResp{
		Code:    http.StatusOK,
		Message: fmt.Sprintf("I will say %v too", er.Word),
	})
}
