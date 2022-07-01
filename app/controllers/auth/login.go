package auth

import (
	"go-commerce/app/utils"
	"net/http"
)

func Login(writer http.ResponseWriter, request *http.Request)  {
	utils.JsonResponse(writer, "From login", 200)
}
