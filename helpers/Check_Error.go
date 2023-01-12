package helpers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"task.com/typedefs"
)

func CheckErr(err error) {
	fmt.Println(err)
}

func SendErrResponse(tp string, message string, w http.ResponseWriter) {
	response := typedefs.Json_Response{}
	response = typedefs.Json_Response{Type: tp, Message: message}
	json.NewEncoder(w).Encode(response)
}

func LogError(err error) {
	output := fmt.Sprint(err)

	file, err := os.OpenFile("logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	log.SetOutput(file)
	log.Println(output)
}
