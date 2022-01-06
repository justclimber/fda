package hasher

// Hasher hashes and validate strings
type Hasher interface {
	// Hash returns a hash for the provided string
	Hash(raw string) (string, error)

	// IsValid checks if a hash and a string match
	IsValid(hash string, raw string) bool
}
