package expense

import(
	app "github.com/Artem32413/Aplication"
)

func Start(expression string)(float64, error){
	znach, err:= app.Calc(expression)
	if err!=nil{
        return 0, err
    }
	return znach, nil
}
