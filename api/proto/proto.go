package proto

import (
)

const (
	DEFAULT_ITEM_TYPE = "http://schema.org/Thing"
)

// func NewCreateItemRequest(itemTypeStr string) (*CreateItemRequest, error) {
// 
//     if itemTypeStr == "" {
//         return nil, errors.New("NewItemRequest called with empty type string.")
//     }
// 
//     var itemTypes []*ItemType
//     if itemTypeStr != "" {
//         uri, err := NormalizeUri(itemTypeStr)
//         if err != nil {
//             return nil, errors.New(
//                 fmt.Sprintf("Unable to validate type of item: `%s'. Expected a url, e.g. http://schema.org/Book. Error: %v", itemTypeStr, err))
//         }
//         itemTypes = append(itemTypes, &ItemType{Uri: uri.String()})
//     }
// 
//     req := &CreateItemRequest{
//         ItemTypes: itemTypes,
//     }
// 
//     return req, nil
// }

