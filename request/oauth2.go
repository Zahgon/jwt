package request

func stripBearerPrefixFromTokenString(tok string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

var AuthorizationHeaderExtractor = &PostExtractionFilter{
	HeaderExtractor{"Authorization"},
	stripBearerPrefixFromTokenString,
}

var OAuth2Extractor = &MultiExtractor{
	AuthorizationHeaderExtractor,
	ArgumentExtractor{"access_token"},
}
