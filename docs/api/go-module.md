# Go Module API

The `productbuildershq-frameworks` Go module provides programmatic access to all framework definitions.

## Installation

```bash
go get github.com/ProductBuildersHQ/productbuildershq-frameworks
```

## Import

```go
import frameworks "github.com/ProductBuildersHQ/productbuildershq-frameworks"
```

## Framework Access

Each framework has three accessor functions:

| Function | Returns | Panics on Error |
|----------|---------|-----------------|
| `Framework()` | `(*Type, error)` | No |
| `MustFramework()` | `*Type` | Yes |
| `FrameworkJSON()` | `[]byte` | No |

### AI-DORA

```go
// Get parsed AI-DORA framework
aidora, err := frameworks.AIDora()
if err != nil {
    log.Fatal(err)
}

// Or panic on error (for initialization)
aidora := frameworks.MustAIDora()

// Get raw JSON bytes
jsonBytes := frameworks.AIDoraJSON()
```

**Type:** `AIDoraFramework`

```go
type AIDoraFramework struct {
    Schema      string        `json:"$schema,omitempty"`
    Framework   string        `json:"framework,omitempty"`
    Name        string        `json:"name,omitempty"`
    Description string        `json:"description,omitempty"`
    Version     string        `json:"version,omitempty"`
    BasedOn     *BasedOn      `json:"basedOn,omitempty"`
    Category    string        `json:"category,omitempty"`
    Metrics     []Metric      `json:"metrics,omitempty"`
    References  []Reference   `json:"references,omitempty"`
}
```

### AI-SPACE

```go
aispace, err := frameworks.AISPACE()
aispace := frameworks.MustAISPACE()
jsonBytes := frameworks.AISPACEJSON()
```

**Type:** `AISPACEFramework`

```go
type AISPACEFramework struct {
    Schema       string            `json:"$schema,omitempty"`
    Framework    string            `json:"framework,omitempty"`
    Name         string            `json:"name,omitempty"`
    Description  string            `json:"description,omitempty"`
    Version      string            `json:"version,omitempty"`
    BasedOn      *BasedOn          `json:"basedOn,omitempty"`
    Category     string            `json:"category,omitempty"`
    Dimensions   []Dimension       `json:"dimensions,omitempty"`
    AntiPatterns []AntiPattern     `json:"antiPatterns,omitempty"`
    References   []Reference       `json:"references,omitempty"`
}
```

### ASDM

```go
asdm, err := frameworks.ASDM()
asdm := frameworks.MustASDM()
jsonBytes := frameworks.ASDMJSON()
```

**Type:** `ASDMFramework`

```go
type ASDMFramework struct {
    Schema      string       `json:"$schema,omitempty"`
    Framework   string       `json:"framework,omitempty"`
    Name        string       `json:"name,omitempty"`
    Description string       `json:"description,omitempty"`
    Version     string       `json:"version,omitempty"`
    Category    string       `json:"category,omitempty"`
    Levels      []ASDMLevel  `json:"levels,omitempty"`
    CaseStudies []CaseStudy  `json:"caseStudies,omitempty"`
    References  []Reference  `json:"references,omitempty"`
}
```

### PBMM

```go
pbmm, err := frameworks.PBMM()
pbmm := frameworks.MustPBMM()
jsonBytes := frameworks.PBMMJSON()
```

**Type:** `PBMMFramework`

```go
type PBMMFramework struct {
    Schema            string             `json:"$schema,omitempty"`
    Framework         string             `json:"framework,omitempty"`
    Name              string             `json:"name,omitempty"`
    Description       string             `json:"description,omitempty"`
    Version           string             `json:"version,omitempty"`
    Category          string             `json:"category,omitempty"`
    Paths             *PBMMPaths         `json:"paths,omitempty"`
    Convergence       *PBMMConvergence   `json:"convergence,omitempty"`
    Levels            []PBMMLevel        `json:"levels,omitempty"`
    FrameworkMappings []FrameworkMapping `json:"frameworkMappings,omitempty"`
    References        []Reference        `json:"references,omitempty"`
}
```

### AIDLC

```go
aidlc, err := frameworks.AIDLC()
aidlc := frameworks.MustAIDLC()
jsonBytes := frameworks.AIDLCJSON()
```

**Type:** `AIDLCFramework`

```go
type AIDLCFramework struct {
    Schema        string              `json:"$schema,omitempty"`
    Framework     string              `json:"framework,omitempty"`
    Name          string              `json:"name,omitempty"`
    Description   string              `json:"description,omitempty"`
    Version       string              `json:"version,omitempty"`
    Phases        []AIDLCPhase        `json:"phases,omitempty"`
    Dependencies  *AIDLCDependencies  `json:"dependencies,omitempty"`
    Evaluation    *AIDLCEvaluation    `json:"evaluation,omitempty"`
    CostEstimates *AIDLCCostEstimates `json:"costEstimates,omitempty"`
    References    []Reference         `json:"references,omitempty"`
}
```

### PDLC

```go
pdlc, err := frameworks.PDLC()
pdlc := frameworks.MustPDLC()
jsonBytes := frameworks.PDLCJSON()
```

**Type:** `PDLCFramework`

```go
type PDLCFramework struct {
    Schema            string              `json:"$schema,omitempty"`
    Framework         string              `json:"framework,omitempty"`
    Name              string              `json:"name,omitempty"`
    Description       string              `json:"description,omitempty"`
    Version           string              `json:"version,omitempty"`
    Category          string              `json:"category,omitempty"`
    Type              string              `json:"type,omitempty"`
    Documentation     PDLCDocumentation   `json:"documentation,omitempty"`
    Phases            []PDLCPhase         `json:"phases,omitempty"`
    Dependencies      *PDLCDependencies   `json:"dependencies,omitempty"`
    Tooling           *PDLCTooling        `json:"tooling,omitempty"`
    RelatedFrameworks []FrameworkRelation `json:"relatedFrameworks,omitempty"`
    References        []Reference         `json:"references,omitempty"`
}

type PDLCPhase struct {
    ID                  string            `json:"id,omitempty"`
    Name                string            `json:"name,omitempty"`
    Order               int               `json:"order,omitempty"`
    Description         string            `json:"description,omitempty"`
    Role                string            `json:"role,omitempty"` // "product" | "builder"
    AIDLCPhase          string            `json:"aiDlcPhase,omitempty"`
    ParallelWith        []string          `json:"parallelWith,omitempty"`
    HumanRole           string            `json:"humanRole,omitempty"`
    SecurityOverlayNote string            `json:"securityOverlayNote,omitempty"`
    SubStages           []PDLCSubStage    `json:"subStages,omitempty"`
    Deliverables        []PDLCDeliverable `json:"deliverables,omitempty"`
    Gates               []PDLCGate        `json:"gates,omitempty"`
}
```

`PDLCPhase` stage IDs (`product-definition`, `builder-definition`, `implementation`, `deployment`, `builder-operations`, `product-operations`) are stable and imported by consumers — e.g. `visionspec` tags each spec type with a PDLC stage, and Threat Model Spec's ASPM security-posture domains map onto the three builder-side stages. See the [PDLC framework page](../frameworks/pdlc/index.md) for the stage-by-stage guide.

## Common Types

### Metric

```go
type Metric struct {
    ID          string       `json:"id,omitempty"`
    Name        string       `json:"name,omitempty"`
    Description string       `json:"description,omitempty"`
    Unit        string       `json:"unit,omitempty"`
    Direction   string       `json:"direction,omitempty"`
    Levels      MetricLevels `json:"levels,omitempty"`
    AIEnablers  []string     `json:"aiEnablers,omitempty"`
}
```

### MetricLevel

```go
type MetricLevel struct {
    Threshold   interface{} `json:"threshold,omitempty"`
    Min         interface{} `json:"min,omitempty"`
    Max         interface{} `json:"max,omitempty"`
    Label       string      `json:"label,omitempty"`
    Description string      `json:"description,omitempty"`
}
```

### Reference

```go
type Reference struct {
    Title string `json:"title,omitempty"`
    URL   string `json:"url,omitempty"`
    Year  int    `json:"year,omitempty"`
    Note  string `json:"note,omitempty"`
}
```

### FrameworkRelation / StageCrosswalk

Generic cross-framework crosswalk types, not specific to any one framework. `PDLCFramework.RelatedFrameworks` uses these to declare its AIDLC stage mapping; any framework can adopt the same pattern to crosswalk against another without a bespoke type.

```go
type FrameworkRelation struct {
    Framework    string           `json:"framework,omitempty"`
    Relationship string           `json:"relationship,omitempty"`
    StageMapping []StageCrosswalk `json:"stageMapping,omitempty"`
}

type StageCrosswalk struct {
    Stage  string   `json:"stage,omitempty"`
    MapsTo []string `json:"mapsTo,omitempty"`
}
```

## Example Usage

### Iterate Over Metrics

```go
aidora := frameworks.MustAIDora()

for _, metric := range aidora.Metrics {
    fmt.Printf("%s (%s)\n", metric.Name, metric.ID)
    fmt.Printf("  Direction: %s\n", metric.Direction)
    fmt.Printf("  Unit: %s\n", metric.Unit)
    fmt.Printf("  Elite: %v\n", metric.Levels.Elite.Threshold)
}
```

### Check ASDM Level

```go
asdm := frameworks.MustASDM()

for _, level := range asdm.Levels {
    if level.Level == 5 {
        fmt.Printf("Level 5: %s\n", level.Name)
        fmt.Printf("  Human Role: %s\n", level.HumanRole)
        for _, practice := range level.Practices {
            fmt.Printf("  - %s (%s)\n", practice.Name, practice.Maturity)
        }
    }
}
```

### Use Raw JSON

```go
// For custom unmarshaling or validation
jsonBytes := frameworks.AIDLCJSON()

var custom map[string]interface{}
if err := json.Unmarshal(jsonBytes, &custom); err != nil {
    log.Fatal(err)
}
```

## Thread Safety

All accessor functions are thread-safe. The `Must*` variants use `sync.Once` for lazy initialization and caching.

## Error Handling

```go
// Explicit error handling
aidora, err := frameworks.AIDora()
if err != nil {
    // Handle JSON parsing error
    log.Printf("Failed to parse AI-DORA: %v", err)
    return
}

// Or panic on error (for initialization)
func init() {
    aidora := frameworks.MustAIDora() // Will panic if JSON is invalid
}
```
