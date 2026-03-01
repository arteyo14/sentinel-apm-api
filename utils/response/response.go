package response

import "github.com/gin-gonic/gin"

type Pagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Total int `json:"total"`
}

type Response struct {
	Status bool        `json:"status"`
	Code   int         `json:"code"`
	Data   any         `json:"data,omitempty"`
	Meta   *Pagination `json:"meta,omitempty"`
	Error  any         `json:"error,omitempty"`
}

func HandleResponse(c *gin.Context, response Response) {
	if response.Status == true {
		c.JSON(response.Code, Response{
			Status: response.Status,
			Code:   response.Code,
			Data:   response.Data,
			Meta:   response.Meta,
		})
	} else {
		c.JSON(response.Code, Response{
			Status: response.Status,
			Code:   response.Code,
			Error:  response.Error,
		})
	}
}
