package errorjson

import (
	s "calc/pkg/mystruct"
	"encoding/json"
	"log"
	"net/http"
)

func InitFuncDecoding(w http.ResponseWriter, decoder *json.Decoder, str *s.Expression, myErr *s.MyError) bool {
	if err := decoder.Decode(str); err != nil {
		log.Println("Error decoding JSON:", err)
		myErr.Error = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		response, err := json.MarshalIndent(myErr, "", "  ")
		if err != nil {
			http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
			return false
		}
		w.Write(response)
		return false
	}
	return true
}
func InitFuncMarshalling(err error, w http.ResponseWriter, myErr *s.MyError) bool {
	if err != nil {
		log.Println("Error marshalling")
		myErr.Error = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		response, err := json.MarshalIndent(myErr, "", "  ")
		if err != nil {
			http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
			return false
		}
		w.Write(response)
		return false
	}
	return true
}
func InitFuncUnmarshalling(err error, todo *[]byte, root *s.Expression, w http.ResponseWriter, myErr *s.MyError) bool {
	if err = json.Unmarshal(*todo, &root); err != nil {
		log.Println("Error unmarshalling")
		myErr.Error = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		response, err := json.MarshalIndent(myErr, "", "  ")
		if err != nil {
			http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
			return false
		}
		w.Write(response)
		return false
	}
	return true
}
func Error422(err error, w http.ResponseWriter, myErr *s.MyError) bool {
	if err != nil {
		myErr.Error = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		response, err := json.MarshalIndent(myErr, "", "  ")
		if err != nil {
			http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
			return false
		}
		w.Write(response)
		return false
	}
	return true
}
func SuccessfulTransition(resRead float64, res s.Result, err error, w http.ResponseWriter) {
	
	res.Result = resRead
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Write(response)
}
