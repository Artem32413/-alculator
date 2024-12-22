package decode

import (
	"encoding/json"
	"net/http"

	er "calc/pkg/decision/decode/errorJson"
	ex "calc/pkg/decision/expense"
	s "calc/pkg/mystruct"
	
)

func Calculate(w http.ResponseWriter, r *http.Request) {
	var str s.Expression
	var res s.Result
	var myErr s.MyError

	decoder := json.NewDecoder(r.Body)
	if resErrDecoding := er.InitFuncDecoding(w, decoder, &str, &myErr); resErrDecoding != true {
		return
	}
	todo, err := json.Marshal(str)
	if resErrMarshalling := er.InitFuncMarshalling(err, w, &myErr); resErrMarshalling != true {
		return
	}
	var root s.Expression
	if resErrUnmarshalling := er.InitFuncUnmarshalling(err, &todo, &root, w, &myErr); resErrUnmarshalling != true {
		return
	}

	resRead, err := ex.Start(root.Expression)
	if FuncExiting := er.Error422(err, w, &myErr); FuncExiting != true {
		return
	} else {
		er.SuccessfulTransition(resRead, res, err, w)
	}
}
