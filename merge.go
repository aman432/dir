package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Patcher defines an interface for patching JSON documents.
type Patcher interface {
	Patch(target, patch map[string]interface{}) map[string]interface{}
}

// MergePatcher implements the Patcher interface for JSON Merge Patch.
type MergePatcher struct{}

// Patch applies the merge patch to the target document.
func (p *MergePatcher) Patch(target, patch map[string]interface{}) map[string]interface{} {
	for key, patchValue := range patch {
		if patchValue == nil {
			delete(target, key)
		} else {
			if targetValue, ok := target[key].(map[string]interface{}); ok {
				if patchValueMap, ok := patchValue.(map[string]interface{}); ok {
					target[key] = p.Patch(targetValue, patchValueMap)
					continue
				}
			}
			target[key] = patchValue
		}
	}
	return target
}

// JSONMerger handles merging JSON documents using the provided Patcher strategy.
type JSONMerger struct {
	patcher Patcher
}

// NewJSONMerger creates a new JSONMerger with the given Patcher strategy.
func NewJSONMerger(patcher Patcher) *JSONMerger {
	return &JSONMerger{patcher: patcher}
}

// Merge applies the patch to the target JSON document.
func (m *JSONMerger) Merge(target, patch []byte) ([]byte, error) {
	var targetMap, patchMap map[string]interface{}

	if err := json.Unmarshal(target, &targetMap); err != nil {
		return nil, fmt.Errorf("error unmarshalling target: %w", err)
	}

	if err := json.Unmarshal(patch, &patchMap); err != nil {
		return nil, fmt.Errorf("error unmarshalling patch: %w", err)
	}

	mergedMap := m.patcher.Patch(targetMap, patchMap)
	return json.Marshal(mergedMap)
}
