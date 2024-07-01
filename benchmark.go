package main

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/evanphx/json-patch"
)

// Sample JSON object to patch
var jsonObject = []byte(`
{
  "name": "John Doe",
  "age": 30,
  "address": {
    "city": "New York",
    "zip": "10001"
  }
}
`)

// Sample JSON patch to apply
var jsonPatchOperations = []byte(`
[
  { "op": "replace", "path": "/name", "value": "Jane Doe" },
  { "op": "add", "path": "/gender", "value": "female" },
  { "op": "remove", "path": "/address/zip" }
]
`)

func BenchmarkApplyJsonPatch(b *testing.B) {
	// Unmarshal initial JSON object
	var initialObj map[string]interface{}
	err := json.Unmarshal(jsonObject, &initialObj)
	if err != nil {
		b.Fatalf("Error unmarshalling initial JSON: %v", err)
	}

	// Unmarshal JSON patch operations
	var patchOps []jsonpatch.Operation
	err = json.Unmarshal(jsonPatchOperations, &patchOps)
	if err != nil {
		b.Fatalf("Error unmarshalling JSON patch operations: %v", err)
	}

	b.ResetTimer()

	// Run the benchmark
	for i := 0; i < b.N; i++ {
		_, err := jsonpatch.ApplyPatch(jsonObject, patchOps)
		if err != nil {
			b.Fatalf("Error applying JSON patch: %v", err)
		}
	}
}

func main() {
	// Run the benchmark
	result := testing.Benchmark(BenchmarkApplyJsonPatch)
	fmt.Printf("Benchmark results:\n%s\n", result)

	// Sleep for a moment to see output
	time.Sleep(1 * time.Second)
}
