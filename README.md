# ProductBuildersHQ Frameworks

[![Go CI][go-ci-svg]][go-ci-url]
[![Go Lint][go-lint-svg]][go-lint-url]
[![Go SAST][go-sast-svg]][go-sast-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![Docs][docs-mkdoc-svg]][docs-mkdoc-url]
[![Visualization][viz-svg]][viz-url]
[![License][license-svg]][license-url]

 [go-ci-svg]: https://github.com/ProductBuildersHQ/productbuildershq-frameworks/actions/workflows/go-ci.yaml/badge.svg?branch=main
 [go-ci-url]: https://github.com/ProductBuildersHQ/productbuildershq-frameworks/actions/workflows/go-ci.yaml
 [go-lint-svg]: https://github.com/ProductBuildersHQ/productbuildershq-frameworks/actions/workflows/go-lint.yaml/badge.svg?branch=main
 [go-lint-url]: https://github.com/ProductBuildersHQ/productbuildershq-frameworks/actions/workflows/go-lint.yaml
 [go-sast-svg]: https://github.com/ProductBuildersHQ/productbuildershq-frameworks/actions/workflows/go-sast-codeql.yaml/badge.svg?branch=main
 [go-sast-url]: https://github.com/ProductBuildersHQ/productbuildershq-frameworks/actions/workflows/go-sast-codeql.yaml
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/ProductBuildersHQ/productbuildershq-frameworks
 [docs-godoc-url]: https://pkg.go.dev/github.com/ProductBuildersHQ/productbuildershq-frameworks
 [docs-mkdoc-svg]: https://img.shields.io/badge/Go-dev%20guide-blue.svg
 [docs-mkdoc-url]: https://productbuildershq.com/productbuildershq-frameworks
 [viz-svg]: https://img.shields.io/badge/Go-visualizaton-blue.svg
 [viz-url]: https://mango-dune-07a8b7110.1.azurestaticapps.net/?repo=ProductBuildersHQ%2Fproductbuildershq-frameworks
 [loc-svg]: https://tokei.rs/b1/github/ProductBuildersHQ/productbuildershq-frameworks
 [repo-url]: https://github.com/ProductBuildersHQ/productbuildershq-frameworks
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/ProductBuildersHQ/productbuildershq-frameworks/blob/main/LICENSE

Machine-readable specifications for maturity models and capability frameworks. These JSON definitions are consumed by [PRISM](https://github.com/grokify/prism) for visualization and tracking.

## Frameworks

### Original Models

| Model | Spec | Documentation |
|-------|------|---------------|
| Product Builder Maturity Model | [`frameworks/product-builder-maturity/`](frameworks/product-builder-maturity/) | [Article](https://productbuildershq.com/frameworks/product-builder-maturity-model) \| [PDF](https://productbuildershq.com/papers/product-builder-maturity-model.pdf) |
| ASDM (Autonomous Software Delivery) | [`frameworks/asdm/`](frameworks/asdm/) | [Article](https://productbuildershq.com/frameworks/software-delivery-autonomy) \| [PDF](https://productbuildershq.com/papers/software-delivery-autonomy.pdf) |

### Development Lifecycle Frameworks

| Framework | Spec | Based On | Documentation |
|-----------|------|----------|---------------|
| AIDLC (AI-Driven Development Lifecycle) | [`frameworks/aidlc/`](frameworks/aidlc/) | [AWS AI DLC](https://docs.aws.amazon.com/prescriptive-guidance/latest/ai-driven-software-development/) | [Article](https://productbuildershq.com/frameworks/aws-aidlc) |
| PDLC (Product Development Lifecycle) | [`frameworks/pdlc/`](frameworks/pdlc/) | Native — splits each AIDLC phase into parallel product/builder lenses | [Specification](https://github.com/ProductBuildersHQ/pdlc) |

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

    // Get AIDLC (3-phase development lifecycle)
    aidlc := frameworks.MustAIDLC()
    fmt.Printf("AIDLC: %d phases, %d total deliverables\n",
        len(aidlc.Phases), countDeliverables(aidlc))

    // Get PDLC (6-stage lifecycle; the canonical stage-ID hub other
    // frameworks and consumers, e.g. Threat Model Spec's ASPM overlay
    // and specification-workflow-spec's spec-type tagging, map into)
    pdlc := frameworks.MustPDLC()
    fmt.Printf("PDLC: %d phases\n", len(pdlc.Phases))
    for _, p := range pdlc.Phases {
        fmt.Printf("  %s (%s, %s)\n", p.Name, p.Role, p.AIDLCPhase)
    }

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
| `AIDLC()` | Returns `*AIDLCFramework, error` |
| `MustAIDLC()` | Returns `*AIDLCFramework` (panics on error) |
| `PDLC()` | Returns `*PDLCFramework, error` |
| `MustPDLC()` | Returns `*PDLCFramework` (panics on error) |
| `AIDoraJSON()` | Returns raw JSON `[]byte` |
| `AISpaceJSON()` | Returns raw JSON `[]byte` |
| `ASDMJSON()` | Returns raw JSON `[]byte` |
| `PBMMJSON()` | Returns raw JSON `[]byte` |
| `AIDLCJSON()` | Returns raw JSON `[]byte` |
| `PDLCJSON()` | Returns raw JSON `[]byte` |
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
