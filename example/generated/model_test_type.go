// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test Service
//     Version: 0.1.0
package api

// TestType is an object.
type TestType struct {
	// AArray
	AArray []interface{} `json:"aArray,omitempty"`
	// AEnum
	AEnum EnumType `json:"aEnum,omitempty"`
	// AInt
	AInt int32 `json:"aInt,omitempty"`
	// ANullableArray
	ANullableArray []interface{} `json:"aNullableArray,omitempty"`
	// ANullableObject
	ANullableObject map[string]interface{} `json:"aNullableObject,omitempty"`
	// ANumber
	ANumber float32 `json:"aNumber,omitempty"`
	// AObject
	AObject map[string]interface{} `json:"aObject,omitempty"`
	// AObjectWithAdditionalComplexProperties
	AObjectWithAdditionalComplexProperties map[string]SubType `json:"aObjectWithAdditionalComplexProperties,omitempty"`
	// AObjectWithAdditionalProperties
	AObjectWithAdditionalProperties map[string]string `json:"aObjectWithAdditionalProperties,omitempty"`
	// AObjectWithAnAnonymousObjectProperty
	AObjectWithAnAnonymousObjectProperty struct {
		Anon string
	} `json:"aObjectWithAnAnonymousObjectProperty,omitempty"`
	// AOneOf
	AOneOf interface{} `json:"aOneOf,omitempty"`
	// APointer
	APointer *string `json:"aPointer,omitempty"`
	// ARequiredValue
	ARequiredValue string `json:"aRequiredValue"`
	// AString
	AString string `json:"aString,omitempty"`
	// AStringArray
	AStringArray []string `json:"aStringArray,omitempty"`
	// AStringWithDescription: this is some documentation
	AStringWithDescription string `json:"aStringWithDescription,omitempty"`
	// ATypedArray
	ATypedArray []SubType `json:"aTypedArray,omitempty"`
	// ATypedObject
	ATypedObject SubType `json:"aTypedObject,omitempty"`
	// AUint64
	AUint64 uint64 `json:"aUint64,omitempty"`
}

// GetAArray returns the AArray property
func (m TestType) GetAArray() []interface{} {
	return m.AArray
}

// SetAArray sets the AArray property
func (m TestType) SetAArray(val []interface{}) {
	m.AArray = val
}

// GetAEnum returns the AEnum property
func (m TestType) GetAEnum() EnumType {
	return m.AEnum
}

// SetAEnum sets the AEnum property
func (m TestType) SetAEnum(val EnumType) {
	m.AEnum = val
}

// GetAInt returns the AInt property
func (m TestType) GetAInt() int32 {
	return m.AInt
}

// SetAInt sets the AInt property
func (m TestType) SetAInt(val int32) {
	m.AInt = val
}

// GetANullableArray returns the ANullableArray property
func (m TestType) GetANullableArray() []interface{} {
	return m.ANullableArray
}

// SetANullableArray sets the ANullableArray property
func (m TestType) SetANullableArray(val []interface{}) {
	m.ANullableArray = val
}

// GetANullableObject returns the ANullableObject property
func (m TestType) GetANullableObject() map[string]interface{} {
	return m.ANullableObject
}

// SetANullableObject sets the ANullableObject property
func (m TestType) SetANullableObject(val map[string]interface{}) {
	m.ANullableObject = val
}

// GetANumber returns the ANumber property
func (m TestType) GetANumber() float32 {
	return m.ANumber
}

// SetANumber sets the ANumber property
func (m TestType) SetANumber(val float32) {
	m.ANumber = val
}

// GetAObject returns the AObject property
func (m TestType) GetAObject() map[string]interface{} {
	return m.AObject
}

// SetAObject sets the AObject property
func (m TestType) SetAObject(val map[string]interface{}) {
	m.AObject = val
}

// GetAObjectWithAdditionalComplexProperties returns the AObjectWithAdditionalComplexProperties property
func (m TestType) GetAObjectWithAdditionalComplexProperties() map[string]SubType {
	return m.AObjectWithAdditionalComplexProperties
}

// SetAObjectWithAdditionalComplexProperties sets the AObjectWithAdditionalComplexProperties property
func (m TestType) SetAObjectWithAdditionalComplexProperties(val map[string]SubType) {
	m.AObjectWithAdditionalComplexProperties = val
}

// GetAObjectWithAdditionalProperties returns the AObjectWithAdditionalProperties property
func (m TestType) GetAObjectWithAdditionalProperties() map[string]string {
	return m.AObjectWithAdditionalProperties
}

// SetAObjectWithAdditionalProperties sets the AObjectWithAdditionalProperties property
func (m TestType) SetAObjectWithAdditionalProperties(val map[string]string) {
	m.AObjectWithAdditionalProperties = val
}

// GetAObjectWithAnAnonymousObjectProperty returns the AObjectWithAnAnonymousObjectProperty property
func (m TestType) GetAObjectWithAnAnonymousObjectProperty() struct {
	Anon string
} {
	return m.AObjectWithAnAnonymousObjectProperty
}

// SetAObjectWithAnAnonymousObjectProperty sets the AObjectWithAnAnonymousObjectProperty property
func (m TestType) SetAObjectWithAnAnonymousObjectProperty(val struct {
	Anon string
}) {
	m.AObjectWithAnAnonymousObjectProperty = val
}

// GetAOneOf returns the AOneOf property
func (m TestType) GetAOneOf() interface{} {
	return m.AOneOf
}

// SetAOneOf sets the AOneOf property
func (m TestType) SetAOneOf(val interface{}) {
	m.AOneOf = val
}

// GetAPointer returns the APointer property
func (m TestType) GetAPointer() *string {
	return m.APointer
}

// SetAPointer sets the APointer property
func (m TestType) SetAPointer(val *string) {
	m.APointer = val
}

// GetARequiredValue returns the ARequiredValue property
func (m TestType) GetARequiredValue() string {
	return m.ARequiredValue
}

// SetARequiredValue sets the ARequiredValue property
func (m TestType) SetARequiredValue(val string) {
	m.ARequiredValue = val
}

// GetAString returns the AString property
func (m TestType) GetAString() string {
	return m.AString
}

// SetAString sets the AString property
func (m TestType) SetAString(val string) {
	m.AString = val
}

// GetAStringArray returns the AStringArray property
func (m TestType) GetAStringArray() []string {
	return m.AStringArray
}

// SetAStringArray sets the AStringArray property
func (m TestType) SetAStringArray(val []string) {
	m.AStringArray = val
}

// GetAStringWithDescription returns the AStringWithDescription property
func (m TestType) GetAStringWithDescription() string {
	return m.AStringWithDescription
}

// SetAStringWithDescription sets the AStringWithDescription property
func (m TestType) SetAStringWithDescription(val string) {
	m.AStringWithDescription = val
}

// GetATypedArray returns the ATypedArray property
func (m TestType) GetATypedArray() []SubType {
	return m.ATypedArray
}

// SetATypedArray sets the ATypedArray property
func (m TestType) SetATypedArray(val []SubType) {
	m.ATypedArray = val
}

// GetATypedObject returns the ATypedObject property
func (m TestType) GetATypedObject() SubType {
	return m.ATypedObject
}

// SetATypedObject sets the ATypedObject property
func (m TestType) SetATypedObject(val SubType) {
	m.ATypedObject = val
}

// GetAUint64 returns the AUint64 property
func (m TestType) GetAUint64() uint64 {
	return m.AUint64
}

// SetAUint64 sets the AUint64 property
func (m TestType) SetAUint64(val uint64) {
	m.AUint64 = val
}
