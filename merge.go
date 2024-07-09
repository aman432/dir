package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// MergePatch applies a JSON Merge Patch to a target document.
func MergePatch(target, patch []byte) ([]byte, error) {
	var targetMap map[string]interface{}
	var patchMap map[string]interface{}

	if err := json.Unmarshal(target, &targetMap); err != nil {
		return nil, fmt.Errorf("error unmarshalling target: %w", err)
	}

	if err := json.Unmarshal(patch, &patchMap); err != nil {
		return nil, fmt.Errorf("error unmarshalling patch: %w", err)
	}

	mergedMap := applyPatch(targetMap, patchMap)

	return json.Marshal(mergedMap)
}

// applyPatch applies the patch map to the target map, handling nested maps.
func applyPatch(targetMap, patchMap map[string]interface{}) map[string]interface{} {
	for key, patchValue := range patchMap {
		if patchValue == nil {
			delete(targetMap, key)
		} else {
			// If both target and patch values are maps, recursively apply patch
			if targetValue, ok := targetMap[key].(map[string]interface{}); ok {
				if patchValueMap, ok := patchValue.(map[string]interface{}); ok {
					targetMap[key] = applyPatch(targetValue, patchValueMap)
					continue
				}
			}
			// Otherwise, just set the patch value
			targetMap[key] = patchValue
		}
	}
	return targetMap
}

func main() {
	original := []byte(`{
		"name": "John",
		"age": 24,
		"address": {
			"city": "New York",
			"zipcode": "10001"
		},
		"skills": ["Go", "Python"]
	}`)
	patch := []byte(`{
		"name": "Jane",
		"address": {
			"zipcode": "10002"
		},
		"skills": null
	}`)

	modified, err := MergePatch(original, patch)
	if err != nil {
		log.Fatalf("Error applying merge patch: %v", err)
	}

	fmt.Printf("Modified document: %s\n", modified)
}
