package frameworks

import (
	"embed"
	"encoding/json"
	"fmt"
	"sync"
)

//go:embed frameworks/ai-dora/ai-dora.json
var aiDoraJSON []byte

//go:embed frameworks/ai-space/ai-space.json
var aiSpaceJSON []byte

//go:embed frameworks/asdm/asdm.json
var asdmJSON []byte

//go:embed frameworks/product-builder-maturity/product-builder-maturity.json
var pbmmJSON []byte

//go:embed frameworks
var frameworksFS embed.FS

var (
	aiDoraOnce   sync.Once
	aiSpaceOnce  sync.Once
	asdmOnce     sync.Once
	pbmmOnce     sync.Once

	aiDoraCached   *AIDoraFramework
	aiSpaceCached  *AISpaceFramework
	asdmCached     *ASDMFramework
	pbmmCached     *PBMMFramework

	aiDoraErr  error
	aiSpaceErr error
	asdmErr    error
	pbmmErr    error
)

// AIDora returns the AI-DORA framework.
// The result is cached after the first call.
func AIDora() (*AIDoraFramework, error) {
	aiDoraOnce.Do(func() {
		aiDoraCached = &AIDoraFramework{}
		aiDoraErr = json.Unmarshal(aiDoraJSON, aiDoraCached)
		if aiDoraErr != nil {
			aiDoraErr = fmt.Errorf("failed to parse AI-DORA: %w", aiDoraErr)
		}
	})
	return aiDoraCached, aiDoraErr
}

// MustAIDora returns the AI-DORA framework or panics on error.
func MustAIDora() *AIDoraFramework {
	f, err := AIDora()
	if err != nil {
		panic(err)
	}
	return f
}

// AISpace returns the AI-SPACE framework.
// The result is cached after the first call.
func AISpace() (*AISpaceFramework, error) {
	aiSpaceOnce.Do(func() {
		aiSpaceCached = &AISpaceFramework{}
		aiSpaceErr = json.Unmarshal(aiSpaceJSON, aiSpaceCached)
		if aiSpaceErr != nil {
			aiSpaceErr = fmt.Errorf("failed to parse AI-SPACE: %w", aiSpaceErr)
		}
	})
	return aiSpaceCached, aiSpaceErr
}

// MustAISpace returns the AI-SPACE framework or panics on error.
func MustAISpace() *AISpaceFramework {
	f, err := AISpace()
	if err != nil {
		panic(err)
	}
	return f
}

// ASDM returns the Autonomous Software Delivery Model framework.
// The result is cached after the first call.
func ASDM() (*ASDMFramework, error) {
	asdmOnce.Do(func() {
		asdmCached = &ASDMFramework{}
		asdmErr = json.Unmarshal(asdmJSON, asdmCached)
		if asdmErr != nil {
			asdmErr = fmt.Errorf("failed to parse ASDM: %w", asdmErr)
		}
	})
	return asdmCached, asdmErr
}

// MustASDM returns the ASDM framework or panics on error.
func MustASDM() *ASDMFramework {
	f, err := ASDM()
	if err != nil {
		panic(err)
	}
	return f
}

// PBMM returns the Product Builder Maturity Model framework.
// The result is cached after the first call.
func PBMM() (*PBMMFramework, error) {
	pbmmOnce.Do(func() {
		pbmmCached = &PBMMFramework{}
		pbmmErr = json.Unmarshal(pbmmJSON, pbmmCached)
		if pbmmErr != nil {
			pbmmErr = fmt.Errorf("failed to parse PBMM: %w", pbmmErr)
		}
	})
	return pbmmCached, pbmmErr
}

// MustPBMM returns the PBMM framework or panics on error.
func MustPBMM() *PBMMFramework {
	f, err := PBMM()
	if err != nil {
		panic(err)
	}
	return f
}

// AIDoraJSON returns the raw JSON bytes for AI-DORA.
func AIDoraJSON() []byte {
	return aiDoraJSON
}

// AISpaceJSON returns the raw JSON bytes for AI-SPACE.
func AISpaceJSON() []byte {
	return aiSpaceJSON
}

// ASDMJSON returns the raw JSON bytes for ASDM.
func ASDMJSON() []byte {
	return asdmJSON
}

// PBMMJSON returns the raw JSON bytes for PBMM.
func PBMMJSON() []byte {
	return pbmmJSON
}

// FS returns the embedded filesystem containing all framework JSON files.
func FS() embed.FS {
	return frameworksFS
}
