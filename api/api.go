package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
    "regexp"
	"time"
)

const (
	DEFAULT_ITEM_TYPE = "http://schema.org/Thing"
)

// Returns a new item object of the type specified.
//
// Used by the CreateItem implementation, or anywhere that the Item model
// is used. Does not assign a URL, or commit to the store.
func NewItem(itemTypeStr string) (i *Item, e error) {
	// return Item{}, ErrUnimplemented
	itemType := ItemType{
		Uri: itemTypeStr,
	}
	itemTypes := []*ItemType{&itemType}

	item := &Item{
		Uri:       "",
		ItemTypes: itemTypes,
	}

	return item, nil
}

func NewCreateItemRequest(itemTypeStr string) (*CreateItemRequest, error) {

    if itemTypeStr == "" {
        return nil, errors.New("NewItemRequest called with empty type string.")
    }

    var itemTypes []*ItemType
    if itemTypeStr != "" {
        uri, err := NormalizeUri(itemTypeStr)
        if err != nil {
            return nil, errors.New(
                fmt.Sprintf("Unable to validate type of item: `%s'. Expected a url, e.g. http://schema.org/Book. Error: %v", itemTypeStr, err))
        }
        itemTypes = append(itemTypes, &ItemType{Uri: uri.String()})
    }

    req := &CreateItemRequest{
        ItemTypes: itemTypes,
    }

    return req, nil
}

// Creates a new statement of the type specified
func NewStatement(s string, p string, o string, c *Scope) (i *Statement, e error) {
    return &Statement{
        Subject: s,
        Predicate: p,
        Object: o,
        Scope: c,
    }, nil
}


// Ensure that a URL is valid, returning it as a url.URL object
//
// Normalization also performs the following:
//  - Expansion of Compact URIs (CURIEs)
//    Both '<schema:Book>' and 'schema:Book' are permitted
//    The list of namespace prefixes supported can be retrieved via
//    GetNamespacePrefixes() (or ltpcli list namespaces)
//    See: https://lov.linkeddata.es/dataset/lov/
func NormalizeUri(uriString string) (*url.URL, error) {

    var re *regexp.Regexp
    var err error

    // Extract the substring of a bracketed <url>, if necessary
    // when the URL matches http/https prefix
    re, err = regexp.Compile("^<?(https?://.*)>?$")
	if err != nil {
		return nil, err
	}
    if m := re.FindStringSubmatch(uriString); len(m) > 1 {
        // We have a URI
        return url.Parse(m[1]) // Return the parsed url, or the error
    }

    // Try for a CURIE
    uri, err := ExpandCurie(uriString)
    if err != nil {
        return nil, err
    } else {
        return uri, nil
    }
}

// Expandeds a CURIE (e.g. <schema:Person> or foaf:name) into a
// qualified name, e.g. https://schema.org/Person
func ExpandCurie(curieString string) (*url.URL, error) {

    // Regex splits input into 3 groups:
    // 0: Left-most match
    // 1: Prefix e.g. `schema'
    // 2: Suffix e.g. `Person'
    re, err := regexp.Compile("^<?([^. <>]+):([^>]+)>?$")
	if err != nil {
		return nil, err
	}
    if m := re.FindStringSubmatch(curieString); len(m) == 3 {
        // We have a CURIE
        prefix := m[1]
        suffix := m[2]
        uriPrefixes := map[string]string{
            "schema": "https://schema.org/",
        }
        ns := uriPrefixes[prefix]
        if prefix != "" {
            return url.Parse(fmt.Sprintf("%s%s", ns, suffix))
        } else {
            return nil, &ErrInvalidUri{Uri: curieString}
        }
    }

    return nil, &ErrInvalidUri{Uri: curieString}
}

// Returns the scope as a URI.
// The URI does not contain all attributes of a scope.
// To serialize an entire scope object, use the GetScopeJson() method
func (s Scope) GetScopeURI() (uri string, err error) {
	url := fmt.Sprintf("%s.%s", s.Agent, s.AssertionTime.String())

	return url, nil
}

// Returns the scope as a JSON decoder
func (s Scope) GetScopeJSON() (json *json.Decoder, err error) {
	return nil, ErrUnimplemented
}

// An activity is used to group the items currently being viewed or interacted
// with by the user, in the completion of a task. Synonymous with 'workspace'
type Activity struct {
	items   []Item
	created time.Time
	updated time.Time
}
