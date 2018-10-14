// Copyright © 2018 Shawn Lower <shawn@shawnlower.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package api defines the API used by the client and server component.
// Serialization and marshalling to protobuf is done in the 'proto' package
package api

import (
    // "encoding/json"
    "errors"
    "fmt"
    "net/url"
    "regexp"

	"github.com/shawnlower/go-ltp/api/proto"

)

const (
	DEFAULT_ITEM_TYPE = "http://schema.org/Thing"
)

// The graph is the backing store containing all item relationships, and
// any literal data for items.
//
//
// Client:
// i := NewItem('<schema:Book>') // RDF type is Book
// i.AddProperty('<schema:name>', 'The Indispensable Calvin and Hobbes')
//
// c := NewClient()
// req := NewCreateItemRequest(i)
// c.CreateItem(req)

// Server:
// store := NewStore(prefix='http://shawnlower.net/i/')
// url := store.MakeUrl()    // return prefix + uuid()
//
// scope := req.SessionID()

// store.AddQuad(scope, url)
// req.Item
// g.AddQuad(req.


// An Item is the basic type used for describing resources.
// It is composed of
// - The URL that can be used to retrieve the resource
// - One or more IRIs that define the type of item (e.g. RDF Type)
// - A list of additional properties that apply to the item

// Example:
//   i := &Item{
//     IRI("") // URL initially undefined
//     ItemType("schema:Movie") // Type
//   }
//   i.AddProperty(IRI("schema:name"), String("Transformers"))

type Item struct {
    IRI IRI
    ItemTypes []IRI
    statements []Statement
}

// Returns a new item object of the type specified.

// Used by the CreateItem implementation, or anywhere that the Item model
// is used. Does not assign a URL, or commit to the store.
func NewItem(itemType IRI) (Item, error) {

    itemTypes := []IRI{itemType}

    item := Item{
		IRI:       "",
		ItemTypes: itemTypes,
	}

	return item, nil
}

func (i Item) GetStatements() ([]Statement, error) {
    return i.statements, nil
}

// AddStatement: Append a statement (s,p,o,l) to an item
func (i *Item) AddStatement(s Statement) error {
    i.statements = append(i.statements, s)
    return nil
}

// AddProperty: Append a property=value pair to an item
// Example:
//   i.AddProperty(IRI("schema:name"), String("Transformers"))
func (i *Item) AddProperty(p Property, v Value) error {
    s := Statement{
        Subject: Value(i.IRI),
        Predicate: Value(p.IRI),
        Object: Value(v),
        Label: IRI(""),
    }

    i.statements = append(i.statements, s)
    i.AddStatement(s)
    return nil
}

// Add a type to an Item
// If a single empty-string type exists, it will be replaced with the
// type specified
func (i Item) AddType(iri IRI) error {
    if len(i.ItemTypes) == 1 && i.ItemTypes[0] == "" {
        i.ItemTypes[0] = iri
    }
    i.ItemTypes = append(i.ItemTypes, iri)
    return nil
}

type Statement struct {
    Subject Value
    Predicate Value
    Object Value
    Label Value
}

type Property struct {
    IRI IRI
    Name string
    Comment string
}

func NewProperty(iri IRI) *Property {
    return &Property{
        IRI: iri,
    }

}

func (p Property) String() string {
    return string(p.IRI)
}

func (p Property) Native() interface{} {
    return p
}

// Interface for IRIs, strings, etc

// See also:
// https://github.com/cayleygraph/cayley/blob/master/quad/value.go
type Value interface {
    String() string
    Native() interface{} // Return closest go type
}

// IRI: An Internationalized Resource Identifier, similar to a URI
// https://tools.ietf.org/html/rfc3987
type IRI string

func (s IRI) String() string {
    return string(s)
}

func (s IRI) Native() interface{} {
    return s
}

type String string

func (s String) String() string {
    return string(s)
}

func (s String) Native() interface{} {
    return s
}

// Submit a resource to a remote store
//     1) Enforce that a type is specified
//     2) Enforce that at least one property from the type is specified
//     3) Optionally upload a payload. The payload URL may be the property

// Example:
//   i := NewItem('<schema:Book>')
//   p := NewProperty('<schema:name>', 'Kitchen Confidential')
//   i.AddProperty(p)
//   i.Validate() // Validate item and properties
//   c := GetClient()

//   c.CreateItem(i.ToRequest())

func (i Item) ToRequest() (*proto.CreateItemRequest, error) {

    if len(i.ItemTypes) == 0 {
        return nil, errors.New("Item type is empty")
    }

    var itemTypes []string
    for _, itemType := range i.ItemTypes {
        iri, err := NormalizeIri(itemType)
        if err != nil {
            return nil, errors.New(
                fmt.Sprintf("Unable to validate type of item: `%s'. Expected a url, e.g. http://schema.org/Book. Error: %v", i, err))
        }
        itemTypes = append(itemTypes, iri.String())
    }

    req := &proto.CreateItemRequest{
        ItemTypes: itemTypes,
    }

    statements, _ := i.GetStatements()
    for _, statement := range statements {
        statementPb := &proto.Statement{
            Subject: statement.Subject.String(),
            Predicate: statement.Predicate.String(),
            Object: statement.Object.String(),
            Label: statement.Label.String(),
        }
        req.Statements = append(req.Statements, statementPb)
    }
    return req, nil
}

// Ensure that a URL is valid, returning it as a url.URL object
//
// Normalization also performs the following:
//  - Expansion of Compact URIs (CURIEs)
//    Both '<schema:Book>' and 'schema:Book' are permitted
//    The list of namespace prefixes supported can be retrieved via
//    GetNamespacePrefixes() (or ltpcli list namespaces)
//    See: https://lov.linkeddata.es/dataset/lov/
func NormalizeIri(iri IRI) (IRI, error) {

    var re *regexp.Regexp
    var err error

    // Extract the substring of a bracketed <url>, if necessary
    // when the URL matches http/https prefix
    if re, err = regexp.Compile("^<?(https?://.*)>?$"); err != nil {
		return "", err
	}
    if m := re.FindStringSubmatch(iri.String()); len(m) > 1 {
        // We have a URI
        return iri, nil
    }

    // Try for a CURIE
    uri, err := ExpandCurie(iri.String())
    if err != nil {
        return "", err
    } else {
        return IRI(uri.String()), nil
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
// func (s Scope) GetScopeURI() (uri string, err error) {
// 	url := fmt.Sprintf("%s.%s", s.Agent, s.AssertionTime.String())

// 	return url, nil
// }

// // Returns the scope as a JSON decoder
// func (s Scope) GetScopeJSON() (json *json.Decoder, err error) {
// 	return nil, ErrUnimplemented
// }

