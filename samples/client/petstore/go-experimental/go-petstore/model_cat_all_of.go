/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package petstore

import (
	"bytes"
	"encoding/json"
)

// CatAllOf struct for CatAllOf
type CatAllOf struct {
	Declawed *bool `json:"declawed,omitempty"`
}

// GetDeclawed returns the Declawed field value if set, zero value otherwise.
func (o *CatAllOf) GetDeclawed() bool {
	if o == nil || o.Declawed == nil {
		var ret bool
		return ret
	}
	return *o.Declawed
}

// GetDeclawedOk returns a tuple with the Declawed field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CatAllOf) GetDeclawedOk() (bool, bool) {
	if o == nil || o.Declawed == nil {
		var ret bool
		return ret, false
	}
	return *o.Declawed, true
}

// HasDeclawed returns a boolean if a field has been set.
func (o *CatAllOf) HasDeclawed() bool {
	if o != nil && o.Declawed != nil {
		return true
	}

	return false
}

// SetDeclawed gets a reference to the given bool and assigns it to the Declawed field.
func (o *CatAllOf) SetDeclawed(v bool) {
	o.Declawed = &v
}

type NullableCatAllOf struct {
	Value CatAllOf
	ExplicitNull bool
}

func (v NullableCatAllOf) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}	
}

func (v *NullableCatAllOf) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

