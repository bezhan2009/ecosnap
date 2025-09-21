package models

type Gemini struct {
	Candidates []GeminiCandidateResp `json:"candidates"`
}

type GeminiCandidateReq struct {
	Contents []GeminiContents `json:"contents"`
}

type GeminiCandidateResp struct {
	Content GeminiContents `json:"content"`
}

type GeminiContents struct {
	Parts []GeminiParts `json:"parts"`
	Role  string        `json:"role"`
}

type GeminiParts struct {
	Text       string            `json:"text,omitempty"`
	InlineData *GeminiInlineData `json:"inlineData,omitempty"`
}

type GeminiInlineData struct {
	MimeType string `json:"mimeType"`
	Data     string `json:"data"`
}
