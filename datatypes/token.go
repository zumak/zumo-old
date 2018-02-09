package datatypes

// Token is
type Token struct {
	Username  string
	HashedKey string
	Salt      string
	Labels    map[string]string
}
