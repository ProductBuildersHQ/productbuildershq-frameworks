# AI-SPACE: AI-Modified Developer Productivity Framework

AI-SPACE adapts the SPACE framework (Microsoft Research, 2021) for measuring developer productivity in AI-assisted environments, where traditional metrics like "lines of code" become meaningless.

**Version**: 1.1.0

## Overview

AI-SPACE maintains SPACE's five dimensions but fundamentally reimagines what each measures in an AI-assisted context. The framework now includes **25 metrics** (5 per dimension), drawing from:

- Original SPACE research (Forsgren et al., 2021)
- GitHub Well-Architected Engineering System Metrics (2025)
- LinearB's comprehensive SPACE implementation (2025)

## The Five Dimensions

### S — Satisfaction and Well-being (5 metrics)

| Aspect | Standard SPACE | AI-SPACE |
|--------|----------------|----------|
| Focus | Developer happiness | Human-AI collaboration satisfaction |
| Key Metric | Satisfaction survey | AI Tool Satisfaction + Tooling Satisfaction |
| New Consideration | N/A | Toil reduction, creative work ratio, work-life balance impact |

**AI-SPACE Metrics:**

| ID | Metric | Description | Unit |
|----|--------|-------------|------|
| ai-satisfaction | AI Tool Satisfaction | Developer satisfaction with AI assistants (Copilot, Claude) | 1-5 |
| tooling-satisfaction | Engineering Tooling Satisfaction | Satisfaction with overall tool ecosystem | 1-5 |
| toil-reduction | Toil Reduction | % reduction in repetitive work due to AI | % |
| creative-ratio | Creative Work Ratio | Time on creative vs routine tasks | % |
| work-life-balance | AI Work-Life Balance Impact | Whether AI improves or degrades work-life balance | 1-5 |

### P — Performance (5 metrics)

| Aspect | Standard SPACE | AI-SPACE |
|--------|----------------|----------|
| Focus | Quality of outputs | AI-assisted output quality |
| Key Metric | Code quality, bugs | AI code acceptance + quality + security |
| New Consideration | N/A | Human oversight effectiveness, AI defect density |

**AI-SPACE Metrics:**

| ID | Metric | Description | Unit |
|----|--------|-------------|------|
| ai-acceptance | AI Code Acceptance Rate | % of AI suggestions accepted after review | % |
| ai-quality | AI Code Quality Score | Quality rating of AI-generated code | 1-5 |
| human-catch-rate | Human Review Catch Rate | Issues found in human review of AI code | % |
| ai-defect-density | AI Code Defect Density | Defects per KLOC in AI-generated code | defects/KLOC |
| code-security | Code Security and Maintainability | Security/maintainability of AI code | 1-5 |

### A — Activity (5 metrics)

| Aspect | Standard SPACE | AI-SPACE |
|--------|----------------|----------|
| Focus | Count of outputs | AI leverage and orchestration quality |
| Key Metric | PRs, commits, LOC | AI Leverage %, prompt reuse |
| Warning | Can be gamed | Even more gameable with AI |

**AI-SPACE Metrics:**

| ID | Metric | Description | Unit |
|----|--------|-------------|------|
| ai-leverage | AI Leverage | % of work performed with AI assistance | % |
| prompt-reuse | Prompt Reuse Rate | % of prompts templated for reuse | % |
| session-efficiency | AI Session Efficiency | Outcomes per AI session hour | outcomes/hr |
| orchestration-complexity | Orchestration Complexity | Multi-step AI workflows created | avg steps |
| feature-delivery | AI-Assisted Feature Delivery | Features delivered per sprint with AI | features/sprint |

### C — Communication and Collaboration (5 metrics)

| Aspect | Standard SPACE | AI-SPACE |
|--------|----------------|----------|
| Focus | Team communication | Human-AI-Human collaboration |
| Key Metric | PR review time, meetings | AI workflow sharing, prompt libraries |
| New Consideration | N/A | Feedback loop speed, AI-assisted documentation |

**AI-SPACE Metrics:**

| ID | Metric | Description | Unit |
|----|--------|-------------|------|
| workflow-sharing | AI Workflow Sharing Rate | % of AI workflows shared with team | % |
| prompt-contributions | Prompt Library Contributions | Monthly contributions to shared prompt libraries | count/month |
| pattern-adoption | Cross-team Pattern Adoption | AI patterns adopted from other teams | % |
| feedback-loop-speed | AI-Enabled Feedback Loop Speed | Time to actionable feedback | minutes |
| ai-documentation | AI-Assisted Documentation Quality | Completeness of AI-maintained docs | % |

### E — Efficiency and Flow (5 metrics)

| Aspect | Standard SPACE | AI-SPACE |
|--------|----------------|----------|
| Focus | Uninterrupted work time | AI-augmented flow state |
| Key Metric | Flow state hours | AI acceleration factor, cycle time |
| New Consideration | N/A | Context switching to/from AI |

**AI-SPACE Metrics:**

| ID | Metric | Description | Unit |
|----|--------|-------------|------|
| acceleration-factor | AI Acceleration Factor | Task speed with AI vs without | x |
| context-switches | Context Switch Frequency | Disruptive human ↔ AI switches | /day |
| augmented-flow | AI-Augmented Flow Hours | Hours per day in AI-augmented flow | hours/day |
| cycle-time | AI-Assisted Cycle Time | Work start to production time | hours |
| time-savings | Weekly AI Time Savings | Hours saved per week due to AI | hours/week |

## Key Differences from Standard SPACE

| Dimension | Standard Focus | AI-SPACE Focus |
|-----------|---------------|----------------|
| **S**atisfaction | Am I happy at work? | Is AI reducing my toil and improving my work? |
| **P**erformance | Is my code good? | Is AI-generated code good? Am I reviewing effectively? |
| **A**ctivity | How much am I producing? | What % is AI-assisted? How well am I orchestrating? |
| **C**ommunication | Am I collaborating with humans? | Am I sharing AI workflows? Are feedback loops fast? |
| **E**fficiency | Am I in flow? | How much is AI accelerating me? What's my cycle time? |

## Anti-Patterns: What NOT to Measure

| Anti-Pattern | Why It Fails |
|--------------|-------------|
| Lines of code | AI can generate infinite LOC; meaningless |
| Number of prompts | Easily gamed; more prompts ≠ better |
| AI tool usage time | More time ≠ better; efficiency matters |
| Raw commit frequency | AI can generate infinite commits |
| Code churn | AI refactoring inflates churn artificially |

## Maturity Levels for AI-SPACE

| Level | Description | Example Metrics |
|-------|-------------|-----------------|
| **L0** | No AI usage | Standard SPACE applies |
| **L1** | Ad-hoc AI | AI satisfaction > 3, some toil reduction |
| **L2** | Structured AI | Prompt reuse > 30%, workflow sharing begins |
| **L3** | Optimized AI | Acceleration factor > 3x, high quality scores |
| **L4** | AI-Native | AI leverage > 80%, < 3 context switches/day |

## Framework Definition Summary

```json
{
  "framework": "AI_SPACE",
  "name": "AI-Modified SPACE",
  "version": "1.1.0",
  "dimensions": [
    {"id": "S", "name": "Satisfaction and Well-being", "metrics": 5},
    {"id": "P", "name": "Performance", "metrics": 5},
    {"id": "A", "name": "Activity", "metrics": 5},
    {"id": "C", "name": "Communication and Collaboration", "metrics": 5},
    {"id": "E", "name": "Efficiency and Flow", "metrics": 5}
  ],
  "totalMetrics": 25,
  "antiPatterns": 5
}
```

## References

- [The SPACE of Developer Productivity](https://queue.acm.org/detail.cfm?id=3454124) - Forsgren et al., 2021 (original SPACE paper)
- [GitHub Well-Architected: Engineering System Metrics](https://wellarchitected.github.com/library/productivity/recommendations/engineering-system-metrics/) - GitHub, 2025
- [LinearB SPACE Framework Guide](https://linearb.io/blog/space-framework) - LinearB, 2025

## Canonical Source

The full framework definition is in [`ai-space.json`](./ai-space.json).
