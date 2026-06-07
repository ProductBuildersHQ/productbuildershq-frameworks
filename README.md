# ProductBuildersHQ Frameworks

Machine-readable specifications for maturity models and capability frameworks. These JSON definitions are consumed by [PRISM](https://github.com/grokify/prism) for visualization and tracking.

## Frameworks

### Original Models

| Model | Spec | Documentation |
|-------|------|---------------|
| Product Builder Maturity Model | [`frameworks/product-builder-maturity/`](frameworks/product-builder-maturity/) | [Article](https://productbuildershq.com/frameworks/product-builder-maturity-model) \| [PDF](https://productbuildershq.com/papers/product-builder-maturity-model.pdf) |
| ASDM (Autonomous Software Delivery) | [`frameworks/asdm/`](frameworks/asdm/) | [Article](https://productbuildershq.com/frameworks/software-delivery-autonomy) \| [PDF](https://productbuildershq.com/papers/software-delivery-autonomy.pdf) |

### AI-Adapted Industry Frameworks

| Framework | Spec | Based On | Documentation |
|-----------|------|----------|---------------|
| AI-DORA | [`frameworks/ai-dora/ai-dora.json`](frameworks/ai-dora/ai-dora.json) | [DORA](https://dora.dev) | [Article](https://productbuildershq.com/frameworks/dora-space-ai-age) |
| AI-SPACE | [`frameworks/ai-space/ai-space.json`](frameworks/ai-space/ai-space.json) | [SPACE](https://queue.acm.org/detail.cfm?id=3454124) | [Article](https://productbuildershq.com/frameworks/dora-space-ai-age) |

### Case Studies

| Case Study | ASDM Level | Documentation |
|------------|------------|---------------|
| AWS Project Mantle | Level 5 | [Article](https://productbuildershq.com/case-studies/aws-project-mantle) \| [PDF](https://productbuildershq.com/papers/aws-project-mantle.pdf) |
| StrongDM Software Factory | Level 6 | [Article](https://productbuildershq.com/case-studies/strongdm-software-factory) \| [PDF](https://productbuildershq.com/papers/strongdm-software-factory.pdf) |

## Go Module Usage

This repository is also a Go module providing typed access to all frameworks.

### Installation

```bash
go get github.com/ProductBuildersHQ/productbuildershq-frameworks
```

### Usage

```go
package main

import (
    "fmt"
    "log"

    frameworks "github.com/ProductBuildersHQ/productbuildershq-frameworks"
)

func main() {
    // Get AI-DORA framework
    aidora, err := frameworks.AIDora()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("AI-DORA v%s: %d metrics\n", aidora.Version, len(aidora.Metrics))

    // Get AI-SPACE framework (25 metrics across 5 dimensions)
    aispace := frameworks.MustAISpace() // panics on error
    for _, dim := range aispace.Dimensions {
        fmt.Printf("  %s: %d metrics\n", dim.Name, len(dim.Metrics))
    }

    // Get ASDM (7-level maturity model)
    asdm := frameworks.MustASDM()
    fmt.Printf("ASDM: %d levels\n", len(asdm.Levels))

    // Get PBMM (dual-path maturity model)
    pbmm := frameworks.MustPBMM()
    fmt.Printf("PBMM: %d levels, converges at L%d\n",
        len(pbmm.Levels), pbmm.Convergence.Level)

    // Access raw JSON if needed
    jsonBytes := frameworks.AIDoraJSON()
    fmt.Printf("Raw AI-DORA JSON: %d bytes\n", len(jsonBytes))

    // Access embedded filesystem
    fs := frameworks.FS()
    data, _ := fs.ReadFile("frameworks/ai-dora/ai-dora.json")
    fmt.Printf("From FS: %d bytes\n", len(data))
}
```

### Available Functions

| Function | Description |
|----------|-------------|
| `AIDora()` | Returns `*AIDoraFramework, error` |
| `MustAIDora()` | Returns `*AIDoraFramework` (panics on error) |
| `AISpace()` | Returns `*AISpaceFramework, error` |
| `MustAISpace()` | Returns `*AISpaceFramework` (panics on error) |
| `ASDM()` | Returns `*ASDMFramework, error` |
| `MustASDM()` | Returns `*ASDMFramework` (panics on error) |
| `PBMM()` | Returns `*PBMMFramework, error` |
| `MustPBMM()` | Returns `*PBMMFramework` (panics on error) |
| `AIDoraJSON()` | Returns raw JSON `[]byte` |
| `AISpaceJSON()` | Returns raw JSON `[]byte` |
| `ASDMJSON()` | Returns raw JSON `[]byte` |
| `PBMMJSON()` | Returns raw JSON `[]byte` |
| `FS()` | Returns `embed.FS` with all JSON files |

## Usage with PRISM

These specs are consumed by PRISM for maturity tracking. An SLI can map to multiple frameworks:

```json
{
  "id": "lead-time",
  "frameworkMappings": [
    {"framework": "DORA", "reference": "LT"},
    {"framework": "AI_DORA", "reference": "LT"},
    {"framework": "SPACE", "reference": "E"}
  ]
}
```

See [prism-examples](https://github.com/grokify/prism-examples) for complete implementations.

## Related Projects

| Project | Purpose |
|---------|---------|
| [productbuildershq.github.io](https://productbuildershq.com) | Articles, PDFs, professional documentation |
| [PRISM](https://github.com/grokify/prism) | Platform consuming these specs |
| [VisionSpec](https://github.com/ProductBuildersHQ/visionspec) | Spec-driven design (AWS WB, Google Docs, Lean Startup) |

## License

Apache 2.0
