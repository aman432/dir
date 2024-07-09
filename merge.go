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

// applyPatch applies the patch map to the target map.
func applyPatch(targetMap, patchMap map[string]interface{}) map[string]interface{} {
	for key, patchValue := range patchMap {
		if patchValue == nil {
			delete(targetMap, key)
		} else {
			targetMap[key] = patchValue
		}
	}
	return targetMap
}
