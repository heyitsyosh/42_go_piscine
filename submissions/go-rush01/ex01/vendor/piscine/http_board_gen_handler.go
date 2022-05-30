package piscine

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type boardGenHttpRequest struct {
	Height int `json:"height"`
	Width  int `json:"width"`
	Black  int `json:"black"`
	Num    int `json:"num"`
}

type boardGenHttpResponse struct {
	Expected []string `json:"expected"`
	NoBlack  []string `json:"no-black"`
}

func HttpBoardGeneratorHandler(writer http.ResponseWriter, req *http.Request) {
	var err error

	if req.Method != "POST" {
		http.Error(writer, "This endpoint only allows POST method", http.StatusMethodNotAllowed)
		return
	}

	// HTTP Request (JSON) -> Struct
	var reqStruct boardGenHttpRequest
	err = json.NewDecoder(req.Body).Decode(&reqStruct)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	if reqStruct.Height < 2 || reqStruct.Width < 2 || reqStruct.Black < 0 || reqStruct.Num < 0 {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(writer, "Invalid Arguments (Height:%d, Width:%d, Black:%d, Num:%d)", reqStruct.Height, reqStruct.Width, reqStruct.Black, reqStruct.Num)
		return
	}

	result, err := BoardGenerator(reqStruct.Height, reqStruct.Width, reqStruct.Black, reqStruct.Num)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	// Generate Response Struct
	var response boardGenHttpResponse
	response.Expected = make([]string, len(result))
	response.NoBlack = make([]string, len(result))
	for index, row := range result {
		rowStr := getRowStr(row, "", true, "", false, GENERATOR_WHITE_NUM, GENERATOR_BLACK_NUM)
		response.Expected[index] = rowStr
		response.NoBlack[index] = strings.Replace(rowStr, "B", ".", -1)
	}

	// Struct -> HTTP Response (JSON)
	err = json.NewEncoder(writer).Encode(response)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
