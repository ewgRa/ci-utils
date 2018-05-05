package hunspell_parser

type SuggestResponse struct {
	Line int `json:"line"`
	Word string `json:"word"`
	Alternative string `json:"alternative"`
}
