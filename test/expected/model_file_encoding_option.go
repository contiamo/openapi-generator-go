// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// FileEncodingOption is an enum. FileEncodingOption determines the encoding used when opening the file.
type FileEncodingOption string

var (
	FileEncodingOptionEUCJP    FileEncodingOption = "EUC-JP"
	FileEncodingOptionEUCKR    FileEncodingOption = "EUC-KR"
	FileEncodingOptionGB18030  FileEncodingOption = "GB18030"
	FileEncodingOptionGBK      FileEncodingOption = "GBK"
	FileEncodingOptionISO88591 FileEncodingOption = "ISO-8859-1"
	FileEncodingOptionISO88592 FileEncodingOption = "ISO-8859-2"
	FileEncodingOptionKOI8R    FileEncodingOption = "KOI8-R"
	FileEncodingOptionShiftJIS FileEncodingOption = "Shift-JIS"
	FileEncodingOptionUTF8     FileEncodingOption = "UTF-8"

	// KnownFileEncodingOption is the list of valid FileEncodingOption
	KnownFileEncodingOption = []FileEncodingOption{
		FileEncodingOptionEUCJP,
		FileEncodingOptionEUCKR,
		FileEncodingOptionGB18030,
		FileEncodingOptionGBK,
		FileEncodingOptionISO88591,
		FileEncodingOptionISO88592,
		FileEncodingOptionKOI8R,
		FileEncodingOptionShiftJIS,
		FileEncodingOptionUTF8,
	}
	// KnownFileEncodingOptionString is the list of valid FileEncodingOption as string
	KnownFileEncodingOptionString = []string{
		string(FileEncodingOptionEUCJP),
		string(FileEncodingOptionEUCKR),
		string(FileEncodingOptionGB18030),
		string(FileEncodingOptionGBK),
		string(FileEncodingOptionISO88591),
		string(FileEncodingOptionISO88592),
		string(FileEncodingOptionKOI8R),
		string(FileEncodingOptionShiftJIS),
		string(FileEncodingOptionUTF8),
	}

	// InKnownFileEncodingOption is an ozzo-validator for FileEncodingOption
	InKnownFileEncodingOption = validation.In(
		FileEncodingOptionEUCJP,
		FileEncodingOptionEUCKR,
		FileEncodingOptionGB18030,
		FileEncodingOptionGBK,
		FileEncodingOptionISO88591,
		FileEncodingOptionISO88592,
		FileEncodingOptionKOI8R,
		FileEncodingOptionShiftJIS,
		FileEncodingOptionUTF8,
	)
)
