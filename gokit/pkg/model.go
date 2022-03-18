package pkg

import "net/http"

type HomeResponse struct {
}

type HomeRequest struct {
}

func (r HomeResponse) StatusCode() int {
	return http.StatusOK
}
