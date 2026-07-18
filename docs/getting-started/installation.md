# Installation

## Go Module

### Requirements

- Go 1.21 or later

### Installation

```bash
go get github.com/ProductBuildersHQ/productbuildershq-frameworks
```

### Verify Installation

```go
package main

import (
    "fmt"
    frameworks "github.com/ProductBuildersHQ/productbuildershq-frameworks"
)

func main() {
    // List all available frameworks
    aidlc := frameworks.MustAIDLC()
    asdm := frameworks.MustASDM()
    aidora := frameworks.MustAIDora()
    aispace := frameworks.MustAISpace()
    pbmm := frameworks.MustPBMM()

    fmt.Println("Frameworks loaded successfully:")
    fmt.Printf("  AIDLC: %s v%s\n", aidlc.Name, aidlc.Version)
    fmt.Printf("  ASDM: %s v%s\n", asdm.Name, asdm.Version)
    fmt.Printf("  AI-DORA: %s v%s\n", aidora.Name, aidora.Version)
    fmt.Printf("  AI-SPACE: %s v%s\n", aispace.Name, aispace.Version)
    fmt.Printf("  PBMM: %s v%s\n", pbmm.Name, pbmm.Version)
}
```

## Raw JSON Access

If you prefer to work with raw JSON files, you can:

### Clone the Repository

```bash
git clone https://github.com/ProductBuildersHQ/productbuildershq-frameworks.git
cd productbuildershq-frameworks
```

### Access JSON Files

```
frameworks/
├── aidlc/
│   ├── aidlc-framework.json      # AIDLC framework definition
│   └── aidlc-workflow.pidl.json  # PIDL process specification
├── asdm/
│   └── asdm.json                 # ASDM maturity model
├── ai-dora/
│   └── ai-dora.json              # AI-DORA metrics
├── ai-space/
│   └── ai-space.json             # AI-SPACE framework
└── product-builder-maturity/
    └── product-builder-maturity.json  # PBMM
```

### Use with jq

```bash
# List AIDLC phases
jq '.phases[].name' frameworks/aidlc/aidlc-framework.json

# List ASDM levels
jq '.levels[] | {level, name}' frameworks/asdm/asdm.json

# Get AI-DORA metrics
jq '.metrics[].id' frameworks/ai-dora/ai-dora.json
```

## Embedded Filesystem

The Go module embeds all JSON files, accessible via `embed.FS`:

```go
package main

import (
    "fmt"
    frameworks "github.com/ProductBuildersHQ/productbuildershq-frameworks"
)

func main() {
    fs := frameworks.FS()

    // Read a specific file
    data, err := fs.ReadFile("frameworks/aidlc/aidlc-framework.json")
    if err != nil {
        panic(err)
    }
    fmt.Printf("Read %d bytes\n", len(data))

    // List all files
    entries, _ := fs.ReadDir("frameworks")
    for _, entry := range entries {
        fmt.Printf("  %s/\n", entry.Name())
    }
}
```

## Related Tools

For full functionality, you may also want to install:

### PIDL (Protocol Visualization)

```bash
go install github.com/grokify/pidl@latest

# Generate workflow diagrams
pidl generate -f mermaid frameworks/aidlc/aidlc-workflow.pidl.json
pidl generate -f svg frameworks/aidlc/aidlc-workflow.pidl.json -o workflow.svg
```

### VisionSpec (Spec-Driven Design)

```bash
go install github.com/ProductBuildersHQ/visionspec/cmd/visionspec@latest

# Use AIDLC for document generation
visionspec init --methodology aidlc
```

### PRISM (Maturity Tracking)

See [PRISM Integration](../integrations/prism.md) for details on using frameworks with PRISM.

## Next Steps

- [Quick Start](quick-start.md) - Basic usage examples
- [Frameworks Overview](../frameworks/index.md) - Learn about each framework
- [API Reference](../api/go-module.md) - Complete API documentation
