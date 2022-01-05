package actionresults

import "encoding/json"

func NewJsonAction(data interface{}) ActionResult {
    return &JsonActionResult{ data: data}
}

type JsonActionResult struct {
    data interface{}
}

func (action *JsonActionResult) Execute(ctx *ActionContext) error {
    ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
    encoder := json.NewEncoder(ctx.ResponseWriter)
    return encoder.Encode(action.data)
}
