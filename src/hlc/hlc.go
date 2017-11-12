package hlc

type HlcRequest struct {
	Text   string //Input text to translate
	Cookie string
}

type HlcResponse struct {
}

type HlcTranslator interface {
	Search(info HlcRequest) string
}
