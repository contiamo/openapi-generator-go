package models

import (
	"bytes"
	"embed"
	"encoding/csv"

	"github.com/mitchellh/mapstructure"
)

//go:embed validations/*_validation.csv
var dataFolder embed.FS

func load() map[string][]*DerivedRule {
	rulesFiles, err := dataFolder.ReadDir("validations")
	if err != nil {
		panic(err)
	}

	rawRules := make(map[string][]map[string]bool)
	for _, file := range rulesFiles {
		rulesData, err := dataFolder.ReadFile("validations/" + file.Name())
		if err != nil {
			panic(err)
		}

		reader := csv.NewReader(bytes.NewBuffer(rulesData))
		reader.TrimLeadingSpace = true
		records, err := reader.ReadAll()
		if err != nil {
			panic(err)
		}

		// Header
		header := records[0]
		records = records[1:]
		for _, row := range records {
			ozzoValidation := row[0]
			rule := make(map[string]bool)
			for idx, col := range header[1:] {
				value := row[idx+1]
				// Do not set empty values
				if value == "" {
					continue
				}
				rule[col] = row[idx+1] == "y"
			}
			rawRules[ozzoValidation] = append(rawRules[ozzoValidation], rule)
		}
	}

	rulesByType := make(map[string][]*DerivedRule)
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		ErrorUnused: true,
		TagName:     "json",
		Result:      &rulesByType,
	})
	if err != nil {
		panic(err)
	}
	err = decoder.Decode(rawRules)
	if err != nil {
		panic(err)
	}

	for _, rules := range rulesByType {
		for _, rule := range rules {
			rule.MakeConsistent()
		}
	}

	return rulesByType
}

var DerivedRulesByValidationType = load()

type DerivedRule struct {
	//Type related
	IsStruct       bool  `json:"struct,omitempty"`
	IsArray        bool  `json:"array,omitempty"`
	IsMap          bool  `json:"map,omitempty"`
	IsNumber       bool  `json:"number,omitempty"`
	IsInteger      bool  `json:"integer,omitempty"`
	IsString       bool  `json:"string,omitempty"`
	IsEnum         bool  `json:"enum,omitempty"`
	IsEnumWithNull *bool `json:"enum_with_null,omitempty"`
	IsEnumWithZero *bool `json:"enum_with_zero,omitempty"`
	IsPtr          bool  `json:"ptr,omitempty"`

	// Format related
	HasFormat  *bool `json:"format,omitempty"`
	HasPattern *bool `json:"pattern,omitempty"`

	// Validation related
	IsRequired *bool `json:"required,omitempty"`

	HasMin               *bool `json:"min,omitempty"`
	HasMax               *bool `json:"max,omitempty"`
	IsMinGreaterThanZero *bool `json:"min_gt_zero,omitempty"`
	IsMaxLessThanZero    *bool `json:"max_lt_zero,omitempty"`

	HasMinLength *bool `json:"min_length,omitempty"`
	HasMaxLength *bool `json:"max_length,omitempty"`

	HasMinItems   *bool `json:"min_items,omitempty"`
	HasMaxItems   *bool `json:"max_items,omitempty"`
	IsUniqueItems *bool `json:"unique_items,omitempty"`

	HasMinProps *bool `json:"min_props,omitempty"`
	HasMaxProps *bool `json:"max_props,omitempty"`
}

func (v *DerivedRule) MakeConsistent() {
	if v.IsMinGreaterThanZero != nil && *v.IsMinGreaterThanZero {
		v.HasMin = new(bool)
		*v.HasMin = true
	}
	if v.IsMaxLessThanZero != nil && *v.IsMaxLessThanZero {
		v.HasMax = new(bool)
		*v.HasMax = true
	}
}
