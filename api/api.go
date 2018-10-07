package api

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

// Creates a new item of the type specified. One or more properties may also
// be specified.
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

// A semantic 'statement' about the world. Can generally be viewed as
// an RDF triple, with an additional 'Scope', used for provenance.
// This can be implemented as a named graph, where the graph name (or label)
// is used for this purpose.
//
// See also:
// <http://patterns.dataincubator.org/book/named-graphs.html>
// type Statement struct {
//    Subject url.URL
//    Predicate url.URL
//    Object url.URL
//    Scope Scope
//}

// Creates a new statement of the type specified
//func NewStatement(s url.URL, p url.URL, o url.URL, c Scope) (i *Statement, e error) {
//    return &Statement{
//        Subject: s,
//        Predicate: p,
//        Object: o,
//        Scope: c,
//    }, nil
//}

// The scope bounds the set of statements or assertions being made by
// an agent.
// Example: scope := &Scope{Time.now(), "ltp_client.shawnlower.net", nil}
type Scope struct {
	AssertionTime time.Time
	Agent         url.URL
	onBehalfOf    url.URL
}

// Returns the scope as a URI.
// The URI does not contain all attributes of a scope.
// To serialize an entire scope object, use the GetScopeJson() method
func (s Scope) GetScopeURI() (uri *url.URL, err error) {
	url, err := url.Parse(fmt.Sprintf("%s.%s",
		s.Agent.String(), s.AssertionTime.String()))

	return url, err
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
