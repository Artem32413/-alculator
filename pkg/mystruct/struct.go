package mystruct

type Expression struct {
	Expression string `json: "expression"`
}
type Result struct {
	Result string `json: "result"`
}
type MyError struct {
	Error string `json: "error"`
}