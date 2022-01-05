package actionresults

func NewErrorAction(err error) ActionResult {
	return &ErrorActionResult{err}
}

type ErrorActionResult struct {
	error
}

func (action *ErrorActionResult) Execute(*ActionContext) error {
	return action.error
}
