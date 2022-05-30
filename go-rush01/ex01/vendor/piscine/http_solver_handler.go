package piscine

import (
	"encoding/json"
	"net/http"
)

type kuromasuSolverHttpRequest struct {
	Request []string `json:"request"`
}

type kuromasuSolverHttpResponse struct {
	Result []string `json:"result"`
}

func HttpKuromasuSolverHandler(writer http.ResponseWriter, req *http.Request) {
	var err error

	if req.Method != "POST" {
		http.Error(writer, "This endpoint only allows POST method", http.StatusMethodNotAllowed)
		return
	}

	// HTTP Request (JSON) -> Struct
	var reqStruct kuromasuSolverHttpRequest
	err = json.NewDecoder(req.Body).Decode(&reqStruct)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	result := Kuromasu(reqStruct.Request, false)

	// Generate Response Struct
	var response kuromasuSolverHttpResponse
	response.Result = make([]string, len(result))
	for index, row := range result {
		response.Result[index] = getRowStr(row, "", true, "", false, SOLVER_WHITE_NUM, SOLVER_BLACK_NUM)
	}

	// Struct -> HTTP Response (JSON)
	err = json.NewEncoder(writer).Encode(response)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
