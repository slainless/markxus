package resty

import (
	"fmt"

	r "github.com/go-resty/resty/v2"
)

type RestyError struct {
	Response *r.Response
}

func (e *RestyError) Error() string {
	body := e.Response.String()
	if len(body) == 0 {
		return fmt.Sprintf(
			"%v %v - (%v)",
			e.Response.Request.Method,
			e.Response.StatusCode(),
			e.Response.Request.URL,
		)
	}

	return fmt.Sprintf(
		"%v %v - %v (%v)",
		e.Response.Request.Method,
		e.Response.StatusCode(),
		body,
		e.Response.Request.URL,
	)
}
