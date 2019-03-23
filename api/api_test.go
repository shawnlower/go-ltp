package api

import (
	"testing"
)

// Test the NormalizeIri() function
func TestNormalizeIri(t *testing.T) {
	validUris := []string{
		"http://golang.org/",
		"https://golang.org/foo#abc",
		"https://schema.org:443/Restaurant",
		"<http://xmlns.com/foaf/0.1/>",
	}
	invalidUris := []string{
		"Hello, 世界!",
		" ",
		"",
		"google.com/",
		"schema.org:Book",
		"<purl.org:Person>",
	}
	cURIEs := []string{
		"<schema:Person>",
		"foaf:name",
	}

	// Test valid URIs
	for _, s := range validUris {
		u, err := NormalizeIri(IRI(s))
		if err != nil {
			t.Fatalf("Failed with valid URI: `%s'. (got url=`%s', err=`%s'.\n",
				s, u, err)
		}
	}

	// Test invalid URIs
	for _, s := range invalidUris {
		u, err := NormalizeIri(IRI(s))
		if err == nil {
			t.Fatalf("Failed to fail on invalid URI: `%s' got `%s'.", s, u)
		}
	}

	// Test valid CURIEs
	for _, s := range cURIEs {
		u, err := NormalizeIri(IRI(s))
		if err != nil {
			t.Fatalf("Failed on valid CURIE: `%s' (got url=`%s', err=`%s'.\n",
				s, u, err)
		}
	}
}

func TestCreateItem(t *testing.T) {
	i, err := NewItem(IRI("schema:Book"))
	if err != nil {
		t.Fatal(err)
	}
	i.ToRequest()
}

func TestEmptyNewItem(t *testing.T) {
	_, err := NewItem(IRI(""))
	if err == nil {
		t.Fatal("Empty NewItemRequest should fail.")
	}
}
