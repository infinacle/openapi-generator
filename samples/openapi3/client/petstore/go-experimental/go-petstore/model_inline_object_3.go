/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"os"
	"time"
	"encoding/json"
	"errors"
)
// InlineObject3 struct for InlineObject3
type InlineObject3 struct {
	// None
	Integer *int32 `json:"integer,omitempty"`

	// None
	Int32 *int32 `json:"int32,omitempty"`

	// None
	Int64 *int64 `json:"int64,omitempty"`

	// None
	Number *float32 `json:"number,omitempty"`

	// None
	Float *float32 `json:"float,omitempty"`

	// None
	Double *float64 `json:"double,omitempty"`

	// None
	String *string `json:"string,omitempty"`

	// None
	PatternWithoutDelimiter *string `json:"pattern_without_delimiter,omitempty"`

	// None
	Byte *string `json:"byte,omitempty"`

	// None
	Binary **os.File `json:"binary,omitempty"`

	// None
	Date *string `json:"date,omitempty"`

	// None
	DateTime *time.Time `json:"dateTime,omitempty"`

	// None
	Password *string `json:"password,omitempty"`

	// None
	Callback *string `json:"callback,omitempty"`

}

// GetInteger returns the Integer field if non-nil, zero value otherwise.
func (o *InlineObject3) GetInteger() int32 {
	if o == nil || o.Integer == nil {
		var ret int32
		return ret
	}
	return *o.Integer
}

// GetIntegerOk returns a tuple with the Integer field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetIntegerOk() (int32, bool) {
	if o == nil || o.Integer == nil {
		var ret int32
		return ret, false
	}
	return *o.Integer, true
}

// HasInteger returns a boolean if a field has been set.
func (o *InlineObject3) HasInteger() bool {
	if o != nil && o.Integer != nil {
		return true
	}

	return false
}

// SetInteger gets a reference to the given int32 and assigns it to the Integer field.
func (o *InlineObject3) SetInteger(v int32) {
	o.Integer = &v
}

// GetInt32 returns the Int32 field if non-nil, zero value otherwise.
func (o *InlineObject3) GetInt32() int32 {
	if o == nil || o.Int32 == nil {
		var ret int32
		return ret
	}
	return *o.Int32
}

// GetInt32Ok returns a tuple with the Int32 field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetInt32Ok() (int32, bool) {
	if o == nil || o.Int32 == nil {
		var ret int32
		return ret, false
	}
	return *o.Int32, true
}

// HasInt32 returns a boolean if a field has been set.
func (o *InlineObject3) HasInt32() bool {
	if o != nil && o.Int32 != nil {
		return true
	}

	return false
}

// SetInt32 gets a reference to the given int32 and assigns it to the Int32 field.
func (o *InlineObject3) SetInt32(v int32) {
	o.Int32 = &v
}

// GetInt64 returns the Int64 field if non-nil, zero value otherwise.
func (o *InlineObject3) GetInt64() int64 {
	if o == nil || o.Int64 == nil {
		var ret int64
		return ret
	}
	return *o.Int64
}

// GetInt64Ok returns a tuple with the Int64 field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetInt64Ok() (int64, bool) {
	if o == nil || o.Int64 == nil {
		var ret int64
		return ret, false
	}
	return *o.Int64, true
}

// HasInt64 returns a boolean if a field has been set.
func (o *InlineObject3) HasInt64() bool {
	if o != nil && o.Int64 != nil {
		return true
	}

	return false
}

// SetInt64 gets a reference to the given int64 and assigns it to the Int64 field.
func (o *InlineObject3) SetInt64(v int64) {
	o.Int64 = &v
}

// GetNumber returns the Number field if non-nil, zero value otherwise.
func (o *InlineObject3) GetNumber() float32 {
	if o == nil || o.Number == nil {
		var ret float32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetNumberOk() (float32, bool) {
	if o == nil || o.Number == nil {
		var ret float32
		return ret, false
	}
	return *o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *InlineObject3) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given float32 and assigns it to the Number field.
func (o *InlineObject3) SetNumber(v float32) {
	o.Number = &v
}

// GetFloat returns the Float field if non-nil, zero value otherwise.
func (o *InlineObject3) GetFloat() float32 {
	if o == nil || o.Float == nil {
		var ret float32
		return ret
	}
	return *o.Float
}

// GetFloatOk returns a tuple with the Float field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetFloatOk() (float32, bool) {
	if o == nil || o.Float == nil {
		var ret float32
		return ret, false
	}
	return *o.Float, true
}

// HasFloat returns a boolean if a field has been set.
func (o *InlineObject3) HasFloat() bool {
	if o != nil && o.Float != nil {
		return true
	}

	return false
}

// SetFloat gets a reference to the given float32 and assigns it to the Float field.
func (o *InlineObject3) SetFloat(v float32) {
	o.Float = &v
}

// GetDouble returns the Double field if non-nil, zero value otherwise.
func (o *InlineObject3) GetDouble() float64 {
	if o == nil || o.Double == nil {
		var ret float64
		return ret
	}
	return *o.Double
}

// GetDoubleOk returns a tuple with the Double field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetDoubleOk() (float64, bool) {
	if o == nil || o.Double == nil {
		var ret float64
		return ret, false
	}
	return *o.Double, true
}

// HasDouble returns a boolean if a field has been set.
func (o *InlineObject3) HasDouble() bool {
	if o != nil && o.Double != nil {
		return true
	}

	return false
}

// SetDouble gets a reference to the given float64 and assigns it to the Double field.
func (o *InlineObject3) SetDouble(v float64) {
	o.Double = &v
}

// GetString returns the String field if non-nil, zero value otherwise.
func (o *InlineObject3) GetString() string {
	if o == nil || o.String == nil {
		var ret string
		return ret
	}
	return *o.String
}

// GetStringOk returns a tuple with the String field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetStringOk() (string, bool) {
	if o == nil || o.String == nil {
		var ret string
		return ret, false
	}
	return *o.String, true
}

// HasString returns a boolean if a field has been set.
func (o *InlineObject3) HasString() bool {
	if o != nil && o.String != nil {
		return true
	}

	return false
}

// SetString gets a reference to the given string and assigns it to the String field.
func (o *InlineObject3) SetString(v string) {
	o.String = &v
}

// GetPatternWithoutDelimiter returns the PatternWithoutDelimiter field if non-nil, zero value otherwise.
func (o *InlineObject3) GetPatternWithoutDelimiter() string {
	if o == nil || o.PatternWithoutDelimiter == nil {
		var ret string
		return ret
	}
	return *o.PatternWithoutDelimiter
}

// GetPatternWithoutDelimiterOk returns a tuple with the PatternWithoutDelimiter field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetPatternWithoutDelimiterOk() (string, bool) {
	if o == nil || o.PatternWithoutDelimiter == nil {
		var ret string
		return ret, false
	}
	return *o.PatternWithoutDelimiter, true
}

// HasPatternWithoutDelimiter returns a boolean if a field has been set.
func (o *InlineObject3) HasPatternWithoutDelimiter() bool {
	if o != nil && o.PatternWithoutDelimiter != nil {
		return true
	}

	return false
}

// SetPatternWithoutDelimiter gets a reference to the given string and assigns it to the PatternWithoutDelimiter field.
func (o *InlineObject3) SetPatternWithoutDelimiter(v string) {
	o.PatternWithoutDelimiter = &v
}

// GetByte returns the Byte field if non-nil, zero value otherwise.
func (o *InlineObject3) GetByte() string {
	if o == nil || o.Byte == nil {
		var ret string
		return ret
	}
	return *o.Byte
}

// GetByteOk returns a tuple with the Byte field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetByteOk() (string, bool) {
	if o == nil || o.Byte == nil {
		var ret string
		return ret, false
	}
	return *o.Byte, true
}

// HasByte returns a boolean if a field has been set.
func (o *InlineObject3) HasByte() bool {
	if o != nil && o.Byte != nil {
		return true
	}

	return false
}

// SetByte gets a reference to the given string and assigns it to the Byte field.
func (o *InlineObject3) SetByte(v string) {
	o.Byte = &v
}

// GetBinary returns the Binary field if non-nil, zero value otherwise.
func (o *InlineObject3) GetBinary() *os.File {
	if o == nil || o.Binary == nil {
		var ret *os.File
		return ret
	}
	return *o.Binary
}

// GetBinaryOk returns a tuple with the Binary field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetBinaryOk() (*os.File, bool) {
	if o == nil || o.Binary == nil {
		var ret *os.File
		return ret, false
	}
	return *o.Binary, true
}

// HasBinary returns a boolean if a field has been set.
func (o *InlineObject3) HasBinary() bool {
	if o != nil && o.Binary != nil {
		return true
	}

	return false
}

// SetBinary gets a reference to the given *os.File and assigns it to the Binary field.
func (o *InlineObject3) SetBinary(v *os.File) {
	o.Binary = &v
}

// GetDate returns the Date field if non-nil, zero value otherwise.
func (o *InlineObject3) GetDate() string {
	if o == nil || o.Date == nil {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetDateOk() (string, bool) {
	if o == nil || o.Date == nil {
		var ret string
		return ret, false
	}
	return *o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *InlineObject3) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *InlineObject3) SetDate(v string) {
	o.Date = &v
}

// GetDateTime returns the DateTime field if non-nil, zero value otherwise.
func (o *InlineObject3) GetDateTime() time.Time {
	if o == nil || o.DateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.DateTime
}

// GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetDateTimeOk() (time.Time, bool) {
	if o == nil || o.DateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.DateTime, true
}

// HasDateTime returns a boolean if a field has been set.
func (o *InlineObject3) HasDateTime() bool {
	if o != nil && o.DateTime != nil {
		return true
	}

	return false
}

// SetDateTime gets a reference to the given time.Time and assigns it to the DateTime field.
func (o *InlineObject3) SetDateTime(v time.Time) {
	o.DateTime = &v
}

// GetPassword returns the Password field if non-nil, zero value otherwise.
func (o *InlineObject3) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetPasswordOk() (string, bool) {
	if o == nil || o.Password == nil {
		var ret string
		return ret, false
	}
	return *o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *InlineObject3) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *InlineObject3) SetPassword(v string) {
	o.Password = &v
}

// GetCallback returns the Callback field if non-nil, zero value otherwise.
func (o *InlineObject3) GetCallback() string {
	if o == nil || o.Callback == nil {
		var ret string
		return ret
	}
	return *o.Callback
}

// GetCallbackOk returns a tuple with the Callback field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetCallbackOk() (string, bool) {
	if o == nil || o.Callback == nil {
		var ret string
		return ret, false
	}
	return *o.Callback, true
}

// HasCallback returns a boolean if a field has been set.
func (o *InlineObject3) HasCallback() bool {
	if o != nil && o.Callback != nil {
		return true
	}

	return false
}

// SetCallback gets a reference to the given string and assigns it to the Callback field.
func (o *InlineObject3) SetCallback(v string) {
	o.Callback = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o InlineObject3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Integer != nil {
		toSerialize["integer"] = o.Integer
	}
	if o.Int32 != nil {
		toSerialize["int32"] = o.Int32
	}
	if o.Int64 != nil {
		toSerialize["int64"] = o.Int64
	}
	if o.Number == nil {
		return nil, errors.New("Number is required and not nullable, but was not set on InlineObject3")
	}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	if o.Float != nil {
		toSerialize["float"] = o.Float
	}
	if o.Double == nil {
		return nil, errors.New("Double is required and not nullable, but was not set on InlineObject3")
	}
	if o.Double != nil {
		toSerialize["double"] = o.Double
	}
	if o.String != nil {
		toSerialize["string"] = o.String
	}
	if o.PatternWithoutDelimiter == nil {
		return nil, errors.New("PatternWithoutDelimiter is required and not nullable, but was not set on InlineObject3")
	}
	if o.PatternWithoutDelimiter != nil {
		toSerialize["pattern_without_delimiter"] = o.PatternWithoutDelimiter
	}
	if o.Byte == nil {
		return nil, errors.New("Byte is required and not nullable, but was not set on InlineObject3")
	}
	if o.Byte != nil {
		toSerialize["byte"] = o.Byte
	}
	if o.Binary != nil {
		toSerialize["binary"] = o.Binary
	}
	if o.Date != nil {
		toSerialize["date"] = o.Date
	}
	if o.DateTime != nil {
		toSerialize["dateTime"] = o.DateTime
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.Callback != nil {
		toSerialize["callback"] = o.Callback
	}
	return json.Marshal(toSerialize)
}


