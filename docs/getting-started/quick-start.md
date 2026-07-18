# Quick Start

This guide shows common usage patterns for the ProductBuildersHQ Frameworks.

## Loading Frameworks

### Basic Loading

```go
package main

import (
    "fmt"
    "log"

    frameworks "github.com/ProductBuildersHQ/productbuildershq-frameworks"
)

func main() {
    // Safe loading with error handling
    aidlc, err := frameworks.AIDLC()
    if err != nil {
        log.Fatalf("Failed to load AIDLC: %v", err)
    }

    // Or use Must* for panic-on-error
    asdm := frameworks.MustASDM()

    fmt.Printf("AIDLC: %d phases\n", len(aidlc.Phases))
    fmt.Printf("ASDM: %d levels\n", len(asdm.Levels))
}
```

### Raw JSON Access

```go
// Get raw JSON bytes
aidlcJSON := frameworks.AIDLCJSON()
asdmJSON := frameworks.ASDMJSON()

// Parse with custom unmarshaling
var custom map[string]interface{}
json.Unmarshal(aidlcJSON, &custom)
```

## Working with AIDLC

### List Phases and Deliverables

```go
aidlc := frameworks.MustAIDLC()

for _, phase := range aidlc.Phases {
    fmt.Printf("\n%s Phase (%s)\n", phase.Name, phase.ID)
    fmt.Printf("  Human Role: %s\n", phase.HumanRole)
    fmt.Printf("  AI Role: %s\n", phase.AIRole)
    fmt.Println("  Deliverables:")
    for _, d := range phase.Deliverables {
        required := ""
        if d.Required {
            required = " (required)"
        }
        fmt.Printf("    - %s%s\n", d.Name, required)
    }
}
```

Output:
```
Inception Phase (inception)
  Human Role: Strategic direction and approval
  AI Role: Document generation and analysis
  Deliverables:
    - Vision Document (required)
    - Requirements Specification (required)
    - Technical Specification (required)
    - Architecture Specification

Construction Phase (construction)
  Human Role: Review and approval
  AI Role: Plan generation and security analysis
  Deliverables:
    - Implementation Plan (required)
    - Test Plan (required)
    - Integration Plan
    - Security Review (required)

Operations Phase (operations)
  Human Role: Approval and ongoing governance
  AI Role: Operational documentation generation
  Deliverables:
    - Runbook (required)
    - Monitoring Plan (required)
    - Disaster Recovery Plan
    - SLO Document (required)
```

### Get Cost Estimates

```go
aidlc := frameworks.MustAIDLC()

if aidlc.CostEstimates != nil {
    fmt.Printf("Estimated workflow cost: $%.2f\n",
        aidlc.CostEstimates.Totals.EstimatedCostUSD)
    fmt.Printf("  Input tokens: %d\n",
        aidlc.CostEstimates.Totals.InputTokens)
    fmt.Printf("  Output tokens: %d\n",
        aidlc.CostEstimates.Totals.OutputTokens)
}
```

### Build Dependency Graph

```go
aidlc := frameworks.MustAIDLC()

fmt.Println("Document Dependencies:")
for _, dep := range aidlc.Dependencies.Graph {
    fmt.Printf("  %s → %s\n", dep.From, dep.To)
}
```

## Working with ASDM

### Assess Maturity Level

```go
asdm := frameworks.MustASDM()

// Find level by name
func findLevel(name string) *frameworks.ASDMLevel {
    for i := range asdm.Levels {
        if asdm.Levels[i].Name == name {
            return &asdm.Levels[i]
        }
    }
    return nil
}

level := findLevel("AI-Native Workflows")
if level != nil {
    fmt.Printf("Level %d: %s\n", level.Level, level.Name)
    fmt.Printf("  Human Role: %s\n", level.HumanRole)
    fmt.Printf("  Defining Characteristic: %s\n", level.DefiningCharacteristic)
    fmt.Printf("  Human Loop: %s\n", level.HumanLoopPosition)
}
```

### List Practices by Level

```go
asdm := frameworks.MustASDM()

for _, level := range asdm.Levels {
    fmt.Printf("\nLevel %d: %s\n", level.Level, level.Name)
    if len(level.Practices) > 0 {
        fmt.Println("  Practices:")
        for _, p := range level.Practices {
            if p.Maturity != "" {
                fmt.Printf("    - %s (%s)\n", p.Name, p.Maturity)
            } else {
                fmt.Printf("    - %s\n", p.Name)
            }
        }
    }
}
```

## Working with AI-DORA

### Display Metrics and Thresholds

```go
aidora := frameworks.MustAIDora()

fmt.Println("AI-DORA Metrics:")
for _, m := range aidora.Metrics {
    fmt.Printf("\n%s (%s)\n", m.Name, m.ID)
    fmt.Printf("  Direction: %s\n", m.Direction)
    fmt.Printf("  Unit: %s\n", m.Unit)
    fmt.Printf("  Elite threshold: %v\n", m.Levels.Elite.Threshold)
    if len(m.AIEnablers) > 0 {
        fmt.Printf("  AI Enablers: %v\n", m.AIEnablers)
    }
}
```

## Working with AI-SPACE

### Explore Dimensions

```go
aispace := frameworks.MustAISpace()

fmt.Printf("AI-SPACE: %d dimensions\n", len(aispace.Dimensions))
for _, dim := range aispace.Dimensions {
    fmt.Printf("\n%s - %s\n", dim.ID, dim.Name)
    fmt.Printf("  %s\n", dim.Description)
    fmt.Printf("  AI Modification: %s\n", dim.AIModification)
    fmt.Printf("  Metrics: %d\n", len(dim.Metrics))
}
```

## Working with PBMM

### Track Dual-Path Progression

```go
pbmm := frameworks.MustPBMM()

fmt.Printf("PBMM Paths:\n")
fmt.Printf("  PM Path: %s\n", pbmm.Paths.PM.Description)
fmt.Printf("  Engineer Path: %s\n", pbmm.Paths.Engineer.Description)
fmt.Printf("  Converge at Level: %d\n", pbmm.Convergence.Level)

for _, level := range pbmm.Levels {
    convergence := ""
    if level.Converged {
        convergence = " (converged)"
    }
    fmt.Printf("\nLevel %d: %s%s\n", level.Level, level.Name, convergence)
    if !level.Converged {
        fmt.Printf("  PM: %s\n", level.PMPath)
        fmt.Printf("  Eng: %s\n", level.EngineerPath)
    }
}
```

## Common Patterns

### Framework Comparison

```go
func compareFrameworks() {
    aidlc := frameworks.MustAIDLC()
    asdm := frameworks.MustASDM()

    fmt.Println("Framework Comparison:")
    fmt.Printf("  AIDLC: %s (v%s) - %d phases\n",
        aidlc.Name, aidlc.Version, len(aidlc.Phases))
    fmt.Printf("  ASDM: %s (v%s) - %d levels\n",
        asdm.Name, asdm.Version, len(asdm.Levels))

    // Find AIDLC reference in ASDM
    for _, level := range asdm.Levels {
        for _, example := range level.Examples {
            if strings.Contains(example, "AI-DLC") {
                fmt.Printf("  AIDLC referenced at ASDM Level %d\n", level.Level)
            }
        }
    }
}
```

### Export to JSON

```go
func exportFramework(filename string) error {
    aidlc := frameworks.MustAIDLC()

    data, err := json.MarshalIndent(aidlc, "", "  ")
    if err != nil {
        return err
    }

    return os.WriteFile(filename, data, 0644)
}
```

## Next Steps

- [AIDLC Framework](../frameworks/aidlc/index.md) - Deep dive into AIDLC
- [ASDM Framework](../frameworks/asdm/index.md) - Explore maturity levels
- [API Reference](../api/go-module.md) - Complete API documentation
- [Integrations](../integrations/prism.md) - Use with PRISM and other tools
