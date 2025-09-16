package hash

import (
	"strings"
	"testing"
)

func TestHash(t *testing.T) {
	testData := []byte("hello world")

	tests := []struct {
		algorithm HashAlgorithm
		expected  string
	}{
		{MD5, "5eb63bbbe01eeed093cb22bb8f5acdc3"},
		{SHA1, "2aae6c35c94fcfb415dbe95f408b9ce91ee846ed"},
		{SHA256, "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"},
		{SHA512, "309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f"},
	}

	for _, test := range tests {
		result, err := Hash(testData, test.algorithm)
		if err != nil {
			t.Errorf("Hash(%s) returned error: %v", test.algorithm, err)
		}
		if result != test.expected {
			t.Errorf("Hash(%s) = %s, want %s", test.algorithm, result, test.expected)
		}
	}
}

func TestHashString(t *testing.T) {
	testString := "hello world"

	tests := []struct {
		algorithm HashAlgorithm
		expected  string
	}{
		{MD5, "5eb63bbbe01eeed093cb22bb8f5acdc3"},
		{SHA1, "2aae6c35c94fcfb415dbe95f408b9ce91ee846ed"},
		{SHA256, "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"},
		{SHA512, "309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f"},
	}

	for _, test := range tests {
		result, err := HashString(testString, test.algorithm)
		if err != nil {
			t.Errorf("HashString(%s) returned error: %v", test.algorithm, err)
		}
		if result != test.expected {
			t.Errorf("HashString(%s) = %s, want %s", test.algorithm, result, test.expected)
		}
	}
}

func TestConvenienceFunctions(t *testing.T) {
	testData := []byte("hello world")
	testString := "hello world"

	// Test MD5
	md5Result := MD5Hash(testData)
	md5StringResult := MD5HashString(testString)
	expectedMD5 := "5eb63bbbe01eeed093cb22bb8f5acdc3"

	if md5Result != expectedMD5 {
		t.Errorf("MD5Hash() = %s, want %s", md5Result, expectedMD5)
	}
	if md5StringResult != expectedMD5 {
		t.Errorf("MD5HashString() = %s, want %s", md5StringResult, expectedMD5)
	}

	// Test SHA256
	sha256Result := SHA256Hash(testData)
	sha256StringResult := SHA256HashString(testString)
	expectedSHA256 := "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"

	if sha256Result != expectedSHA256 {
		t.Errorf("SHA256Hash() = %s, want %s", sha256Result, expectedSHA256)
	}
	if sha256StringResult != expectedSHA256 {
		t.Errorf("SHA256HashString() = %s, want %s", sha256StringResult, expectedSHA256)
	}
}

func TestUnsupportedAlgorithm(t *testing.T) {
	_, err := Hash([]byte("test"), HashAlgorithm("unsupported"))
	if err == nil {
		t.Error("Expected error for unsupported algorithm, got nil")
	}
}

func TestHashWithReader(t *testing.T) {
	testString := "hello world"

	tests := []struct {
		algorithm HashAlgorithm
		expected  string
	}{
		{MD5, "5eb63bbbe01eeed093cb22bb8f5acdc3"},
		{SHA1, "2aae6c35c94fcfb415dbe95f408b9ce91ee846ed"},
		{SHA256, "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"},
		{SHA512, "309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f"},
	}

	for _, test := range tests {
		// Create a new reader for each test since readers are consumed
		reader := strings.NewReader(testString)
		result, err := HashWithReader(reader, test.algorithm)
		if err != nil {
			t.Errorf("HashWithReader(%s) returned error: %v", test.algorithm, err)
		}
		if result != test.expected {
			t.Errorf("HashWithReader(%s) = %s, want %s", test.algorithm, result, test.expected)
		}
	}
}

func TestHashWithReaderConvenienceFunctions(t *testing.T) {
	testString := "hello world"
	reader := strings.NewReader(testString)

	// Test MD5
	md5Result, err := MD5HashWithReader(reader)
	if err != nil {
		t.Errorf("MD5HashWithReader() returned error: %v", err)
	}
	expectedMD5 := "5eb63bbbe01eeed093cb22bb8f5acdc3"
	if md5Result != expectedMD5 {
		t.Errorf("MD5HashWithReader() = %s, want %s", md5Result, expectedMD5)
	}

	// Test SHA256
	reader = strings.NewReader(testString) // Create new reader since previous was consumed
	sha256Result, err := SHA256HashWithReader(reader)
	if err != nil {
		t.Errorf("SHA256HashWithReader() returned error: %v", err)
	}
	expectedSHA256 := "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"
	if sha256Result != expectedSHA256 {
		t.Errorf("SHA256HashWithReader() = %s, want %s", sha256Result, expectedSHA256)
	}
}

func TestHashWithReaderUnsupportedAlgorithm(t *testing.T) {
	reader := strings.NewReader("test")
	_, err := HashWithReader(reader, HashAlgorithm("unsupported"))
	if err == nil {
		t.Error("Expected error for unsupported algorithm, got nil")
	}
}
