# ASDM: Autonomous Software Delivery Model

A seven-level progression from traditional Agile teams to autonomous software delivery and operations.

## Overview

ASDM (Autonomous Software Delivery Model) defines the methodology for achieving autonomous coding, review, validation, and operations. It provides the practices, artifacts, and infrastructure required for each level of software delivery autonomy.

## The Seven Levels

| Level | Name | Human Role | Defining Characteristic |
|-------|------|------------|------------------------|
| 1 | Agile/Scrum | Execute work | Team coordination optimization |
| 2 | DevOps | Build + operate | Pipeline automation (CI/CD) |
| 3 | AI-Assisted | Primary implementer | AI augments existing process |
| 4 | AI-Native Workflows | Orchestrator | Process redesigned around AI |
| 5 | Agentic Engineering | Batch reviewer | Extended autonomous sessions |
| 6 | Autonomous Coding & Review | Specification owner | No human coding or code review |
| 7 | Autonomous Operations | Governor | Autonomous production operations |

## Human Loop Position

| Levels | Position | Description |
|--------|----------|-------------|
| 1-3 | Human-in-the-loop | Interactive, every step |
| 4 | Human-in-the-loop | But process is AI-native |
| 5 | Human-on-the-loop | Batch review after autonomous sessions |
| 6 | Human-over-the-code-loop | Humans define intent and scenarios, not code |
| 7 | Human-over-the-operations-loop | Governance, incidents, and exceptions only |

## ASDM Practices

| Practice | Level 4 | Level 5 | Level 6 | Level 7 |
|----------|---------|---------|---------|---------|
| Steering files (AGENTS.md) | Introduce | Mature | Required | Required |
| Extended autonomous sessions | — | Core | Core | Core |
| Scenario-based validation | Experiment | Partial | Required | Required |
| Digital Twin infrastructure | — | Experiment | Required | Required |
| Zero human code review | — | — | Required | Required |
| Governance-by-policy | Introduce | Develop | Required | Required |
| Autonomous operations | — | — | Experiment | Required |

## Key Artifacts by Level

| Level | Artifacts | Rituals |
|-------|----------|---------|
| 4 | Steering files, prompt libraries, context specs | Mob Elaboration, Mob Construction, prompt reviews |
| 5 | Execution logs, batch review queues, agent traces | Queue-before-sleep, morning review, high-bandwidth standups |
| 6 | Scenario libraries, Digital Twins, satisfaction dashboards | Scenario authoring, DTU maintenance, policy governance |
| 7 | Operational policy engines, production guardrails, remediation playbooks | SLO governance, incident simulation, exception review |

## Real-World Examples

| Level | Example |
|-------|---------|
| Level 4 | AWS AI-Driven Development Lifecycle (AI-DLC) |
| Level 5 | AWS Project Mantle, Spotify Honk |
| Level 6 | StrongDM Software Factory |
| Level 7 | Emerging frontier (no public examples yet) |

## Canonical Source

The full ASDM methodology is documented in:

- **Article**: [Software Delivery Autonomy Levels](https://productbuildershq.com/frameworks/software-delivery-autonomy)
- **PDF**: [Software Delivery Autonomy](https://productbuildershq.com/papers/software-delivery-autonomy.pdf)
- **Machine-readable spec**: [`asdm.json`](asdm.json)

## References

- AWS AI-Driven Development Lifecycle (AI-DLC), January 2026
- AWS Project Mantle, January 2026
- Spotify Honk, February 2026
- StrongDM Software Factory, February 2026
