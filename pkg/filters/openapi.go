package filters

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/getkin/kin-openapi/openapi3"
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

	// note that we _need_ yaml.v3 or ghodss/yaml here because v2 unmarhsals to map[interface{}]interface{}
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

	// responses are optional
	var responses Entry
	r := components["responses"]
	if r != nil {
		responses, ok = r.(Entry)
		if !ok {
			return filteredSpec, fmt.Errorf("`responses` is invalid type")
		}
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

	// now filter the response schemas and then check
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

	filteredYAML, err := yaml.Marshal(spec)
	if err != nil {
		return filteredSpec, err
	}

	// ensure we don't return an invalid spec
	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(filteredYAML)
	err = swagger.Validate(context.Background())
	if err != nil {
		return filteredSpec, err
	}

	// can't directly marshal to yaml because of extra fields
	outJSON, err := swagger.MarshalJSON()
	if err != nil {
		return filteredSpec, err
	}

	return JSONToYAML(outJSON)
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

func JSONToYAML(j []byte) ([]byte, error) {
	// Convert the JSON to an object.
	var jsonObj interface{}
	// We are using yaml.Unmarshal here (instead of json.Unmarshal) because the
	// Go JSON library doesn't try to pick the right number type (int, float,
	// etc.) when unmarshalling to interface{}, it just picks float64
	// universally. go-yaml does go through the effort of picking the right
	// number type, so we can preserve number type throughout this process.
	err := yaml.Unmarshal(j, &jsonObj)
	if err != nil {
		return nil, err
	}

	var out bytes.Buffer
	enc := yaml.NewEncoder(&out)
	enc.SetIndent(2)
	err = enc.Encode(jsonObj)

	// Marshal this object into YAML.
	return out.Bytes(), err
}
