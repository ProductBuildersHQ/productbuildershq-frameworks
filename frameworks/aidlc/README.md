# AWS AI-Driven Development Lifecycle (AIDLC)

A three-phase development lifecycle for AI-native software delivery, adapted from the AWS AI DLC methodology.

## Overview

AIDLC structures AI-assisted software development into three distinct phases, each with specific deliverables, human review gates, and quality evaluation criteria.

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                              AIDLC Workflow                                  │
├───────────────────┬───────────────────────┬─────────────────────────────────┤
│    INCEPTION      │     CONSTRUCTION      │          OPERATIONS             │
├───────────────────┼───────────────────────┼─────────────────────────────────┤
│ Vision Document   │ Implementation Plan   │ Runbook                         │
│        ↓          │        ↓    ↓         │    ↓                            │
│ Requirements Spec │ Test Plan  Integration│ Disaster Recovery Plan          │
│    ↓       ↓      │    ↓       Plan       │                                 │
│ Technical  Arch   │ Monitoring            │ Monitoring Plan                 │
│ Spec      Spec    │ Plan                  │    ↓                            │
│        ↓          │                       │ SLO Document                    │
│   [Human Gate]    │ Security Review       │                                 │
│                   │        ↓              │        ↓                        │
│                   │   [Human Gate]        │   [Human Gate]                  │
└───────────────────┴───────────────────────┴─────────────────────────────────┘
```

## Phases

### Phase 1: Inception

Discovery and requirements definition. Establishes the strategic foundation for the project.

| Document | Required | Description |
|----------|----------|-------------|
| Vision Document | Yes | Strategic vision, problem space, target users, success criteria |
| Requirements Specification | Yes | Functional/non-functional requirements, acceptance criteria |
| Technical Specification | Yes | APIs, data models, system interfaces |
| Architecture Specification | No | Component design, scalability, deployment topology |

**Gate:** Human review and approval before proceeding to Construction.

### Phase 2: Construction

Implementation planning and quality assurance. Prepares for development execution.

| Document | Required | Description |
|----------|----------|-------------|
| Implementation Plan | Yes | Milestones, task breakdown, resource allocation |
| Test Plan | Yes | Test strategy, test cases, automation approach |
| Integration Plan | No | External system integration, rollout strategy |
| Security Review | Yes | Threat modeling, access control, compliance |

**Gate:** Human review and approval before proceeding to Operations.

### Phase 3: Operations

Deployment and operational readiness. Ensures production-ready systems.

| Document | Required | Description |
|----------|----------|-------------|
| Runbook | Yes | Deployment procedures, troubleshooting, maintenance |
| Monitoring Plan | Yes | Metrics, alerts, dashboards |
| Disaster Recovery Plan | No | RTO/RPO targets, backup, failover procedures |
| SLO Document | Yes | SLIs, SLOs, error budgets, reporting |

**Gate:** Human review and approval before production deployment.

## Files in This Directory

| File | Purpose |
|------|---------|
| [`aidlc-framework.json`](aidlc-framework.json) | PRISM-compatible framework definition |
| [`aidlc-workflow.pidl.json`](aidlc-workflow.pidl.json) | PIDL process specification for visualization |
| [`phases/inception.md`](phases/inception.md) | Detailed Inception phase documentation |
| [`phases/construction.md`](phases/construction.md) | Detailed Construction phase documentation |
| [`phases/operations.md`](phases/operations.md) | Detailed Operations phase documentation |

## Cost Estimates

Estimated LLM costs for generating all 12 documents (at Claude Sonnet pricing):

| Metric | Value |
|--------|-------|
| Total Input Tokens | 127,000 |
| Total Output Tokens | 95,000 |
| Input Cost | $0.38 |
| Output Cost | $1.43 |
| **Total Estimated Cost** | **$1.81** |

## Evaluation

AIDLC uses LLM-as-Judge evaluation with a 1-5 scoring scale:

| Score | Label | Description |
|-------|-------|-------------|
| 5 | Excellent | Exceeds all criteria |
| 4 | Good | Meets all criteria with minor gaps |
| 3 | Average | Meets basic criteria |
| 2 | Below Average | Significant gaps |
| 1 | Poor | Does not meet criteria |

### Pass Criteria

**Default:**
- Minimum weighted score: 70%
- Required category minimum: 3
- Maximum critical issues: 0
- Maximum high issues: 2

**Strict (for Security Review):**
- Minimum weighted score: 85%
- Required category minimum: 4
- Maximum critical/high issues: 0

## PIDL Visualization

Generate workflow diagrams using the PIDL CLI:

```bash
# Mermaid sequence diagram
pidl generate -f mermaid aidlc-workflow.pidl.json

# SVG workflow diagram
pidl generate -f svg aidlc-workflow.pidl.json -o aidlc-workflow.svg

# Animated SVG
pidl generate -f svg-animated aidlc-workflow.pidl.json

# Security analysis
pidl analyze aidlc-workflow.pidl.json

# Infographic for presentations
pidl generate -f infographic --ig-title="AWS AIDLC" aidlc-workflow.pidl.json
```

## Integration with VisionSpec

AIDLC integrates with [VisionSpec](https://github.com/ProductBuildersHQ/visionspec) for:

- Bidirectional sync between `.visionspec/` and `aidlc-docs/` directories
- Workflow progress tracking and visualization
- LLM-as-Judge document evaluation
- Rubric-based quality scoring

```go
import "github.com/ProductBuildersHQ/visionspec/pkg/aidlc"

// Create default AIDLC workflow
workflow := aidlc.DefaultWorkflow()

// Track execution progress
ctx, _ := aidlc.NewWorkflowExecutionContext(workflow)
ctx.StartStep("vision_document")
ctx.CompleteStep("vision_document", &score)

// Get metrics
metrics := ctx.GetMetrics()
fmt.Printf("Progress: %.1f%%\n", metrics.ProgressPercent)
```

## Relationship to ASDM

AIDLC is the recommended methodology at **ASDM Level 4 (AI-Native Workflows)**. It represents a process redesigned around AI capabilities:

| ASDM Level | Methodology | Human Role |
|------------|-------------|------------|
| Level 3 | AI-Assisted | Primary implementer |
| **Level 4** | **AIDLC** | **Orchestrator** |
| Level 5 | Agentic Engineering | Batch reviewer |

## References

- [AWS AI-Driven Development Lifecycle](https://docs.aws.amazon.com/prescriptive-guidance/latest/ai-driven-software-development/)
- [AWS Project Mantle Case Study](https://productbuildershq.com/case-studies/aws-project-mantle)
- [ASDM Framework](../asdm/)
