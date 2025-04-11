package urlparser

type URL struct {
	BaseURL       string   // BaseURL is the base URL of the website
	FuzzablePaths []string // Paths is the list of paths of the website
}
