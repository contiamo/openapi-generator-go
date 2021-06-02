package filters

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"gopkg.in/yaml.v3"
)

type Any = interface{}
type Entry = map[string]Any
type Nothing = struct{}

// ByPath removes all the endpoints that are not specified in the given list of.
//
// It also, removes all the schema definitions which are not used by the white-listed
// endpoints keeping the schemas which are referenced by other schemas recursively.
//
// For now, it supports only `components/schemas/*` as references (no responses, for example).
func ByPath(file io.Reader, allowedPaths []string) (filteredSpec []byte, err error) {

	var spec Entry
	data, err := io.ReadAll(file)
	if err != nil {
		return filteredSpec, err
	}

	if len(allowedPaths) == 0 {
		return data, nil
	}

	// note that we _need_ yaml.v3 here because v2 unmarhsals to map[interface{}]interface{}
	err = yaml.Unmarshal(data, &spec)
	if err != nil {
		return filteredSpec, err
	}

	filterList := make(map[string]Nothing, len(allowedPaths))
	for _, allowedPath := range allowedPaths {
		filterList[allowedPath] = Nothing{}
	}

	paths, ok := spec["paths"].(Entry)
	if !ok {
		return filteredSpec, fmt.Errorf("`paths` is invalid type: %T", spec["paths"])
	}

	components, ok := spec["components"].(Entry)
	if !ok {
		return filteredSpec, fmt.Errorf("`components` is invalid type")
	}

	schemas, ok := components["schemas"].(Entry)
	if !ok {
		return filteredSpec, fmt.Errorf("`schemas` is invalid type")
	}

	responses, ok := components["responses"].(Entry)
	if !ok {
		return filteredSpec, fmt.Errorf("`responses` is invalid type")
	}

	// filter out extra endpoints
	paths = filter(paths, filterList)
	spec["paths"] = paths

	// find all references from the endpoints
	referenceList := make(map[string]Nothing)
	findReferences(paths, referenceList)

	// now we need to filter out all the unused schemas
	// but they have references too, so we need to find them
	// and put the removed used references back
	newSchemas := filter(schemas, referenceList)
	var count int

	for len(referenceList) != count {
		count = len(referenceList)
		findReferences(newSchemas, referenceList)
		newSchemas = filter(schemas, referenceList)
	}

	// now filter the reponse schemas and then check
	// if there are still any more references we need
	// to include
	newResponses := filter(responses, referenceList)
	findReferences(newResponses, referenceList)
	newSchemas = filter(schemas, referenceList)

	for len(referenceList) != count {
		count = len(referenceList)
		findReferences(newSchemas, referenceList)
		newSchemas = filter(schemas, referenceList)
	}
	components["schemas"] = newSchemas
	components["responses"] = newResponses

	removeOneOf(spec)

	var out bytes.Buffer
	enc := yaml.NewEncoder(&out)
	enc.SetIndent(2)
	err = enc.Encode(spec)

	return out.Bytes(), err
}

func findReferences(entry Entry, results map[string]Nothing) {
	for key := range entry {
		if key == "$ref" {
			refName := getRefName(entry[key])
			if refName == "" {
				continue
			}
			results[refName] = Nothing{}
		}
		subEntry, ok := entry[key].(Entry)
		if ok {
			findReferences(subEntry, results)
		}

		// handles entries like "paramaters"
		subList, ok := entry[key].([]interface{})
		if ok {
			for _, entry := range subList {
				subEntry, ok := entry.(Entry)
				if ok {
					findReferences(subEntry, results)
				}
			}
		}
	}
}

func filter(entry Entry, whiteList map[string]Nothing) (newEntry Entry) {
	newEntry = make(map[string]Any, len(entry))

	for key := range entry {
		if _, ok := whiteList[key]; !ok {
			continue
		}
		newEntry[key] = entry[key]
	}

	return newEntry
}

func removeOneOf(entry Entry) {
	for key := range entry {
		if key == "oneOf" {
			delete(entry, key)
			if entry["type"] == nil {
				entry["type"] = "object"
			}
		}
		subEntry, ok := entry[key].(Entry)
		if ok {
			removeOneOf(subEntry)
		}
	}
}

func getRefName(fullPath Any) string {
	refPath, ok := fullPath.(string)
	if !ok {
		return ""
	}
	segments := strings.Split(refPath, "/")
	l := len(segments)
	if l == 0 {
		return ""
	}
	return segments[l-1]
}
