package requestsdto

type CreateCompletionRequest struct {
	Image string `json:"image" form:"image"`
}

type CreateCompletionRequestText struct {
	Text string `json:"text" form:"text"`
}
