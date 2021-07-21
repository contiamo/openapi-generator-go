// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"

	"time"
)

// Person is an object.
type Person struct {
	// Address:
	Address Address `json:"address,omitempty"`
	// Age:
	Age float32 `json:"age,omitempty"`
	// Base64:
	Base64 string `json:"base64,omitempty"`
	// Date:
	Date string `json:"date,omitempty"`
	// Datetime:
	Datetime time.Time `json:"datetime,omitempty"`
	// Email:
	Email string `json:"email,omitempty"`
	// FavoriteColors:
	FavoriteColors []Color `json:"favoriteColors"`
	// Gender:
	Gender Gender `json:"gender"`
	// Hostname:
	Hostname string `json:"hostname,omitempty"`
	// Ip:
	Ip string `json:"ip,omitempty"`
	// Ipv4:
	Ipv4 string `json:"ipv4,omitempty"`
	// Ipv6:
	Ipv6 string `json:"ipv6,omitempty"`
	// Name:
	Name string `json:"name"`
	// RequestURI:
	RequestURI string `json:"requestURI,omitempty"`
	// SecondGender:
	SecondGender *string `json:"secondGender,omitempty"`
	// Uri:
	Uri string `json:"uri,omitempty"`
	// Url:
	Url string `json:"url,omitempty"`
	// Uuid:
	Uuid string `json:"uuid,omitempty"`
}

// Validate implements basic validation for this model
func (m Person) Validate() error {
	return validation.Errors{
		"address": validation.Validate(
			m.Address,
		),
		"age": validation.Validate(
			m.Age, validation.Min(float32(18)), validation.Max(float32(120)),
		),
		"base64": validation.Validate(
			m.Base64, is.Base64,
		),
		"email": validation.Validate(
			m.Email, is.EmailFormat,
		),
		"favoriteColors": validation.Validate(
			m.FavoriteColors, validation.Required, validation.Length(1, 0),
		),
		"gender": validation.Validate(
			m.Gender, validation.Required,
		),
		"hostname": validation.Validate(
			m.Hostname, is.Host,
		),
		"ip": validation.Validate(
			m.Ip, is.IP,
		),
		"ipv4": validation.Validate(
			m.Ipv4, is.IPv4,
		),
		"ipv6": validation.Validate(
			m.Ipv6, is.IPv6,
		),
		"name": validation.Validate(
			m.Name, validation.Required, validation.Length(2, 32),
		),
		"requestURI": validation.Validate(
			m.RequestURI, is.RequestURL.Error("must be valid URI with scheme"),
		),
		"uri": validation.Validate(
			m.Uri, is.RequestURI,
		),
		"url": validation.Validate(
			m.Url, is.URL.Error("must be a valid URL with HTTP or HTTPS scheme"),
		),
		"uuid": validation.Validate(
			m.Uuid, is.UUID,
		),
	}.Filter()
}

// GetAddress returns the Address property
func (m Person) GetAddress() Address {
	return m.Address
}

// SetAddress sets the Address property
func (m *Person) SetAddress(val Address) {
	m.Address = val
}

// GetAge returns the Age property
func (m Person) GetAge() float32 {
	return m.Age
}

// SetAge sets the Age property
func (m *Person) SetAge(val float32) {
	m.Age = val
}

// GetBase64 returns the Base64 property
func (m Person) GetBase64() string {
	return m.Base64
}

// SetBase64 sets the Base64 property
func (m *Person) SetBase64(val string) {
	m.Base64 = val
}

// GetDate returns the Date property
func (m Person) GetDate() string {
	return m.Date
}

// SetDate sets the Date property
func (m *Person) SetDate(val string) {
	m.Date = val
}

// GetDatetime returns the Datetime property
func (m Person) GetDatetime() time.Time {
	return m.Datetime
}

// SetDatetime sets the Datetime property
func (m *Person) SetDatetime(val time.Time) {
	m.Datetime = val
}

// GetEmail returns the Email property
func (m Person) GetEmail() string {
	return m.Email
}

// SetEmail sets the Email property
func (m *Person) SetEmail(val string) {
	m.Email = val
}

// GetFavoriteColors returns the FavoriteColors property
func (m Person) GetFavoriteColors() []Color {
	return m.FavoriteColors
}

// SetFavoriteColors sets the FavoriteColors property
func (m *Person) SetFavoriteColors(val []Color) {
	m.FavoriteColors = val
}

// GetGender returns the Gender property
func (m Person) GetGender() Gender {
	return m.Gender
}

// SetGender sets the Gender property
func (m *Person) SetGender(val Gender) {
	m.Gender = val
}

// GetHostname returns the Hostname property
func (m Person) GetHostname() string {
	return m.Hostname
}

// SetHostname sets the Hostname property
func (m *Person) SetHostname(val string) {
	m.Hostname = val
}

// GetIp returns the Ip property
func (m Person) GetIp() string {
	return m.Ip
}

// SetIp sets the Ip property
func (m *Person) SetIp(val string) {
	m.Ip = val
}

// GetIpv4 returns the Ipv4 property
func (m Person) GetIpv4() string {
	return m.Ipv4
}

// SetIpv4 sets the Ipv4 property
func (m *Person) SetIpv4(val string) {
	m.Ipv4 = val
}

// GetIpv6 returns the Ipv6 property
func (m Person) GetIpv6() string {
	return m.Ipv6
}

// SetIpv6 sets the Ipv6 property
func (m *Person) SetIpv6(val string) {
	m.Ipv6 = val
}

// GetName returns the Name property
func (m Person) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *Person) SetName(val string) {
	m.Name = val
}

// GetRequestURI returns the RequestURI property
func (m Person) GetRequestURI() string {
	return m.RequestURI
}

// SetRequestURI sets the RequestURI property
func (m *Person) SetRequestURI(val string) {
	m.RequestURI = val
}

// GetSecondGender returns the SecondGender property
func (m Person) GetSecondGender() *string {
	return m.SecondGender
}

// SetSecondGender sets the SecondGender property
func (m *Person) SetSecondGender(val *string) {
	m.SecondGender = val
}

// GetUri returns the Uri property
func (m Person) GetUri() string {
	return m.Uri
}

// SetUri sets the Uri property
func (m *Person) SetUri(val string) {
	m.Uri = val
}

// GetUrl returns the Url property
func (m Person) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *Person) SetUrl(val string) {
	m.Url = val
}

// GetUuid returns the Uuid property
func (m Person) GetUuid() string {
	return m.Uuid
}

// SetUuid sets the Uuid property
func (m *Person) SetUuid(val string) {
	m.Uuid = val
}
