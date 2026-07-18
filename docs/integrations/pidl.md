# PIDL Integration

[PIDL](https://github.com/grokify/pidl) (Protocol Interaction Description Language) is a specification format for describing processes with steps, inputs, outputs, cost models, and latency budgets.

## Overview

PIDL provides:

- **Process specifications** - Step-by-step workflow definitions
- **Data ports** - Typed inputs and outputs for each step
- **Cost models** - LLM token usage and cost estimation
- **Latency budgets** - P50/P95/P99 timing expectations
- **Failure modes** - Error handling and retry strategies

## AIDLC PIDL Spec

The AIDLC framework includes a comprehensive PIDL specification at:

```
frameworks/aidlc/aidlc-workflow.pidl.json
```

### Structure

```json
{
  "metadata": {
    "name": "aidlc-workflow",
    "version": "1.0.0",
    "description": "AWS AI-Driven Development Lifecycle workflow"
  },
  "process": {
    "steps": [...]
  }
}
```

### Step Definition

Each document generation step includes:

```json
{
  "id": "generate-vision-document",
  "name": "Generate Vision Document",
  "type": "llm_generation",
  "inputs": [
    {
      "name": "product_brief",
      "dataType": "document",
      "required": true
    }
  ],
  "outputs": [
    {
      "name": "vision_document",
      "dataType": "document"
    }
  ],
  "costModel": {
    "provider": "anthropic",
    "model": "claude-3-sonnet",
    "estimatedInputTokens": 5000,
    "estimatedOutputTokens": 3000,
    "estimatedCostUSD": 0.06
  },
  "latencyBudget": {
    "p50Ms": 15000,
    "p95Ms": 30000,
    "p99Ms": 60000
  },
  "failureModes": [
    {
      "type": "rate_limit",
      "retryStrategy": "exponential_backoff",
      "maxRetries": 3
    }
  ]
}
```

## Using PIDL Specs

### Load and Parse

```go
import (
    "encoding/json"
    "os"
)

type PIDLSpec struct {
    Metadata PIDLMetadata `json:"metadata"`
    Process  PIDLProcess  `json:"process"`
}

// Load PIDL spec
data, _ := os.ReadFile("frameworks/aidlc/aidlc-workflow.pidl.json")
var spec PIDLSpec
json.Unmarshal(data, &spec)

// Access steps
for _, step := range spec.Process.Steps {
    fmt.Printf("%s: %s\n", step.ID, step.Name)
    fmt.Printf("  Cost: $%.2f\n", step.CostModel.EstimatedCostUSD)
}
```

### Calculate Total Cost

```go
func totalCost(spec PIDLSpec) float64 {
    var total float64
    for _, step := range spec.Process.Steps {
        if step.CostModel != nil {
            total += step.CostModel.EstimatedCostUSD
        }
    }
    return total
}
```

### Validate Dependencies

```go
func validateDependencies(spec PIDLSpec) error {
    stepIDs := make(map[string]bool)
    for _, step := range spec.Process.Steps {
        stepIDs[step.ID] = true
    }

    for _, step := range spec.Process.Steps {
        for _, dep := range step.DependsOn {
            if !stepIDs[dep] {
                return fmt.Errorf("step %s depends on unknown step %s", step.ID, dep)
            }
        }
    }
    return nil
}
```

## PIDL Features

### Step Types

| Type | Description |
|------|-------------|
| `llm_generation` | LLM-based document generation |
| `human_review` | Human review gate |
| `automated_check` | Automated validation |
| `integration` | External system integration |

### Data Ports

```json
{
  "inputs": [
    {
      "name": "requirements_spec",
      "dataType": "document",
      "format": "markdown",
      "required": true,
      "description": "Approved requirements specification"
    }
  ],
  "outputs": [
    {
      "name": "technical_spec",
      "dataType": "document",
      "format": "markdown",
      "schema": "technical-spec-v1"
    }
  ]
}
```

### Cost Models

```json
{
  "costModel": {
    "provider": "anthropic",
    "model": "claude-3-sonnet",
    "estimatedInputTokens": 12000,
    "estimatedOutputTokens": 10000,
    "estimatedCostUSD": 0.19,
    "currency": "USD"
  }
}
```

### Latency Budgets

```json
{
  "latencyBudget": {
    "p50Ms": 20000,
    "p95Ms": 45000,
    "p99Ms": 90000,
    "timeoutMs": 120000
  }
}
```

### Failure Modes

```json
{
  "failureModes": [
    {
      "type": "rate_limit",
      "retryStrategy": "exponential_backoff",
      "maxRetries": 3,
      "initialDelayMs": 1000
    },
    {
      "type": "context_overflow",
      "retryStrategy": "chunk_and_retry",
      "maxRetries": 2
    },
    {
      "type": "quality_threshold",
      "retryStrategy": "regenerate",
      "maxRetries": 2
    }
  ]
}
```

## Workflow Execution

PIDL specs can drive workflow execution:

```go
type WorkflowExecutor struct {
    spec   PIDLSpec
    state  map[string]interface{}
    status map[string]StepStatus
}

func (e *WorkflowExecutor) ExecuteStep(stepID string) error {
    step := e.findStep(stepID)

    // Check dependencies
    for _, dep := range step.DependsOn {
        if e.status[dep] != StatusCompleted {
            return fmt.Errorf("dependency %s not completed", dep)
        }
    }

    // Gather inputs
    inputs := e.gatherInputs(step)

    // Execute based on step type
    switch step.Type {
    case "llm_generation":
        return e.executeLLMGeneration(step, inputs)
    case "human_review":
        return e.executeHumanReview(step, inputs)
    }

    return nil
}
```

## Best Practices

1. **Use PIDL for complex workflows** - Simple processes may not need full PIDL
2. **Include realistic cost estimates** - Update estimates based on actual usage
3. **Set appropriate latency budgets** - Account for LLM variability
4. **Define failure modes** - All steps should have retry strategies
5. **Validate dependencies** - Ensure dependency graph is acyclic

## References

- [PIDL Repository](https://github.com/grokify/pidl)
- [AIDLC Framework](../frameworks/aidlc/index.md)
- [VisionSpec Integration](visionspec.md)
