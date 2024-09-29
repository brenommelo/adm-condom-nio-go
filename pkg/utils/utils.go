package utils

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ParseJsonRequest(r *http.Request, dst interface{}) error {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	return decoder.Decode(dst)
}

func WriteErrorResponse(ctx *gin.Context, code int, msg string) {
	// ctx.Header("value", "result")
	ctx.JSON(code, gin.H{
		"error":     msg,
		"errorCode": code,
	})
}

func WriteJSONResponse(ctx *gin.Context, statusCode int, message string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.Header("message", message)
	ctx.JSON(statusCode, gin.H{
		"data": data,
	})
}

func WriteJSONResponseMarshal(w http.ResponseWriter, statusCode int, data interface{}) {
	dat, err := json.Marshal(data)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(dat)
}
