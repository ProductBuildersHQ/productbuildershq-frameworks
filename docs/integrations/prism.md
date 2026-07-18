# PRISM Integration

[PRISM](https://github.com/grokify/prism-roadmap) is a maturity tracking tool that consumes framework definitions to track organizational progress against metrics.

## Overview

PRISM provides:

- **Roadmap tracking** across multiple maturity frameworks
- **Metric scoring** against framework thresholds
- **Progress visualization** over time
- **Gap analysis** between current and target states

## Framework Compatibility

All frameworks in this repository are PRISM-compatible:

| Framework | Category | PRISM Support |
|-----------|----------|---------------|
| AI-DORA | DevOps Metrics | Full |
| AI-SPACE | Developer Productivity | Full |
| ASDM | Maturity Model | Full |
| PBMM | Maturity Model | Full |
| AIDLC | Methodology | Partial |

## Configuration

### Importing Frameworks

```json
{
  "prismConfig": {
    "frameworks": [
      {
        "source": "github.com/ProductBuildersHQ/productbuildershq-frameworks",
        "frameworks": ["AI_DORA", "AI_SPACE", "ASDM", "PBMM"]
      }
    ]
  }
}
```

### Mapping Capabilities to Metrics

```json
{
  "capabilities": [
    {
      "id": "deployment-automation",
      "name": "Deployment Automation",
      "frameworkMappings": [
        {
          "framework": "AI_DORA",
          "metric": "deployment-frequency"
        },
        {
          "framework": "ASDM",
          "level": 4
        }
      ]
    }
  ]
}
```

## Metric Scoring

PRISM scores capabilities against framework thresholds:

```json
{
  "assessment": {
    "capabilityId": "deployment-automation",
    "scores": [
      {
        "framework": "AI_DORA",
        "metric": "deployment-frequency",
        "value": 8,
        "unit": "per day",
        "level": "high",
        "targetLevel": "elite"
      }
    ]
  }
}
```

### Threshold Evaluation

| Framework | Direction | Scoring Logic |
|-----------|-----------|---------------|
| AI-DORA | `higher_is_better` | Value >= threshold |
| AI-DORA | `lower_is_better` | Value <= threshold |
| AI-SPACE | `optimum` | Min <= Value <= Max |

## Example: Tracking ASDM Progress

```json
{
  "roadmap": {
    "name": "Autonomous Delivery Journey",
    "framework": "ASDM",
    "currentLevel": 3,
    "targetLevel": 5,
    "milestones": [
      {
        "level": 4,
        "name": "AI-Native Workflows",
        "targetDate": "2025-Q2",
        "practices": [
          {"name": "Steering files", "status": "in_progress"},
          {"name": "Governance-by-policy", "status": "planned"}
        ]
      },
      {
        "level": 5,
        "name": "Agentic Engineering",
        "targetDate": "2025-Q4",
        "practices": [
          {"name": "Extended autonomous sessions", "status": "planned"},
          {"name": "Digital Twin infrastructure", "status": "planned"}
        ]
      }
    ]
  }
}
```

## Dashboard Integration

PRISM dashboards display framework metrics:

```yaml
dashboard:
  panels:
    - type: metric-grid
      title: AI-DORA Metrics
      framework: AI_DORA
      metrics:
        - deployment-frequency
        - lead-time
        - change-failure-rate
        - mttr

    - type: maturity-ladder
      title: ASDM Progress
      framework: ASDM
      showPractices: true
```

## API Usage

### Fetch Framework Definition

```bash
curl -s https://api.prism.dev/frameworks/AI_DORA | jq '.metrics[].name'
```

### Submit Assessment

```bash
curl -X POST https://api.prism.dev/assessments \
  -H "Content-Type: application/json" \
  -d '{
    "teamId": "platform-team",
    "framework": "AI_DORA",
    "metrics": [
      {"id": "deployment-frequency", "value": 12, "unit": "per day"},
      {"id": "lead-time", "value": 10, "unit": "minutes"}
    ]
  }'
```

## Best Practices

1. **Start with one framework** - Begin with AI-DORA or ASDM before adding others
2. **Map incrementally** - Add framework mappings as capabilities mature
3. **Track trends** - PRISM value comes from tracking progress over time
4. **Use realistic targets** - Set achievable level targets based on current state
5. **Automate collection** - Integrate metrics collection with CI/CD and observability tools

## References

- [PRISM Documentation](https://prism.dev/docs)
- [PRISM Roadmap Repository](https://github.com/grokify/prism-roadmap)
