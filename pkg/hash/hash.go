package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"io"
)

// HashAlgorithm represents the supported hash algorithms
type HashAlgorithm string

const (
	MD5    HashAlgorithm = "md5"
	SHA1   HashAlgorithm = "sha1"
	SHA256 HashAlgorithm = "sha256"
	SHA512 HashAlgorithm = "sha512"
)

// Hash calculates the hash of the input data using the specified algorithm
func Hash(data []byte, algorithm HashAlgorithm) (string, error) {
	var h hash.Hash

	switch algorithm {
	case MD5:
		h = md5.New()
	case SHA1:
		h = sha1.New()
	case SHA256:
		h = sha256.New()
	case SHA512:
		h = sha512.New()
	default:
		return "", fmt.Errorf("unsupported hash algorithm: %s", algorithm)
	}

	h.Write(data)
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// HashString calculates the hash of a string using the specified algorithm
func HashString(data string, algorithm HashAlgorithm) (string, error) {
	return Hash([]byte(data), algorithm)
}

// HashWithReader calculates the hash of data from an io.Reader using the specified algorithm
func HashWithReader(reader io.Reader, algorithm HashAlgorithm) (string, error) {
	var h hash.Hash

	switch algorithm {
	case MD5:
		h = md5.New()
	case SHA1:
		h = sha1.New()
	case SHA256:
		h = sha256.New()
	case SHA512:
		h = sha512.New()
	default:
		return "", fmt.Errorf("unsupported hash algorithm: %s", algorithm)
	}

	_, err := io.Copy(h, reader)
	if err != nil {
		return "", fmt.Errorf("failed to read from reader: %w", err)
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// MD5Hash calculates MD5 hash of the input data
func MD5Hash(data []byte) string {
	h := md5.New()
	h.Write(data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// MD5HashString calculates MD5 hash of a string
func MD5HashString(data string) string {
	return MD5Hash([]byte(data))
}

// SHA1Hash calculates SHA1 hash of the input data
func SHA1Hash(data []byte) string {
	h := sha1.New()
	h.Write(data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// SHA1HashString calculates SHA1 hash of a string
func SHA1HashString(data string) string {
	return SHA1Hash([]byte(data))
}

// SHA256Hash calculates SHA256 hash of the input data
func SHA256Hash(data []byte) string {
	h := sha256.New()
	h.Write(data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// SHA256HashString calculates SHA256 hash of a string
func SHA256HashString(data string) string {
	return SHA256Hash([]byte(data))
}

// SHA512Hash calculates SHA512 hash of the input data
func SHA512Hash(data []byte) string {
	h := sha512.New()
	h.Write(data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// SHA512HashString calculates SHA512 hash of a string
func SHA512HashString(data string) string {
	return SHA512Hash([]byte(data))
}

// MD5HashWithReader calculates MD5 hash of data from an io.Reader
func MD5HashWithReader(reader io.Reader) (string, error) {
	return HashWithReader(reader, MD5)
}

// SHA1HashWithReader calculates SHA1 hash of data from an io.Reader
func SHA1HashWithReader(reader io.Reader) (string, error) {
	return HashWithReader(reader, SHA1)
}

// SHA256HashWithReader calculates SHA256 hash of data from an io.Reader
func SHA256HashWithReader(reader io.Reader) (string, error) {
	return HashWithReader(reader, SHA256)
}

// SHA512HashWithReader calculates SHA512 hash of data from an io.Reader
func SHA512HashWithReader(reader io.Reader) (string, error) {
	return HashWithReader(reader, SHA512)
}
