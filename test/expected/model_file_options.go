// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// FileOptions is an object. FileOptions determine how the file will be opened and parsed
type FileOptions struct {
	// CommentCharacter: The character used to indicate that a line is a comment, the line must start with this character.
	//
	// Common values include "#" and "//"
	CommentCharacter string `json:"commentCharacter"`
	// DecimalSeparator: FileDecimalSeparator determines how numeric values are parsed, e.g. '10,50'.
	DecimalSeparator FileDecimalSeparator `json:"decimalSeparator"`
	// Delimiter: FileDelimiter is used to determine the column separator in a CSV file.
	//
	// Common values include 'comma', 'semicolon', and 'tab'
	Delimiter string `json:"delimiter"`
	// Encoding: FileEncodingOption determines the encoding used when opening the file.
	Encoding FileEncodingOption `json:"encoding"`
	// HasHeader: Read the first non-empty line as the column headers/names
	HasHeader bool `json:"hasHeader"`
	// IgnoreHeader: If the file has headers, these are ignored when determining the column names
	IgnoreHeader bool `json:"ignoreHeader"`
	// PreferredDateFormat: FilePreferredDateFormat determines how to parse ambiguous dates, such as '02-02-2010'.
	PreferredDateFormat FilePreferredDateFormat `json:"preferredDateFormat"`
	// TrimSpaces: Determine if discovery should trim leading and trailing spaces from values. Note that spaces will always be trimmed from column names.
	TrimSpaces bool `json:"trimSpaces"`
}

// GetCommentCharacter returns the CommentCharacter property
func (m FileOptions) GetCommentCharacter() string {
	return m.CommentCharacter
}

// SetCommentCharacter sets the CommentCharacter property
func (m FileOptions) SetCommentCharacter(val string) {
	m.CommentCharacter = val
}

// GetDecimalSeparator returns the DecimalSeparator property
func (m FileOptions) GetDecimalSeparator() FileDecimalSeparator {
	return m.DecimalSeparator
}

// SetDecimalSeparator sets the DecimalSeparator property
func (m FileOptions) SetDecimalSeparator(val FileDecimalSeparator) {
	m.DecimalSeparator = val
}

// GetDelimiter returns the Delimiter property
func (m FileOptions) GetDelimiter() string {
	return m.Delimiter
}

// SetDelimiter sets the Delimiter property
func (m FileOptions) SetDelimiter(val string) {
	m.Delimiter = val
}

// GetEncoding returns the Encoding property
func (m FileOptions) GetEncoding() FileEncodingOption {
	return m.Encoding
}

// SetEncoding sets the Encoding property
func (m FileOptions) SetEncoding(val FileEncodingOption) {
	m.Encoding = val
}

// GetHasHeader returns the HasHeader property
func (m FileOptions) GetHasHeader() bool {
	return m.HasHeader
}

// SetHasHeader sets the HasHeader property
func (m FileOptions) SetHasHeader(val bool) {
	m.HasHeader = val
}

// GetIgnoreHeader returns the IgnoreHeader property
func (m FileOptions) GetIgnoreHeader() bool {
	return m.IgnoreHeader
}

// SetIgnoreHeader sets the IgnoreHeader property
func (m FileOptions) SetIgnoreHeader(val bool) {
	m.IgnoreHeader = val
}

// GetPreferredDateFormat returns the PreferredDateFormat property
func (m FileOptions) GetPreferredDateFormat() FilePreferredDateFormat {
	return m.PreferredDateFormat
}

// SetPreferredDateFormat sets the PreferredDateFormat property
func (m FileOptions) SetPreferredDateFormat(val FilePreferredDateFormat) {
	m.PreferredDateFormat = val
}

// GetTrimSpaces returns the TrimSpaces property
func (m FileOptions) GetTrimSpaces() bool {
	return m.TrimSpaces
}

// SetTrimSpaces sets the TrimSpaces property
func (m FileOptions) SetTrimSpaces(val bool) {
	m.TrimSpaces = val
}
