# AI-SPACE - AI-Modified Developer Productivity

AI-SPACE adapts the [SPACE framework](https://queue.acm.org/detail.cfm?id=3454124) for AI-assisted software development, providing metrics that measure developer productivity in human-AI collaborative environments.

## Overview

SPACE was introduced by Forsgren et al. in 2021 to measure developer productivity across five dimensions. AI-SPACE modifies these dimensions to reflect the unique characteristics of AI-assisted development.

## Dimensions

### S - Satisfaction and Well-being

Developer satisfaction with AI-assisted work environment.

| Metric | Description | Elite | High | Medium |
|--------|-------------|-------|------|--------|
| AI Tool Satisfaction | Satisfaction with AI coding assistants | >4.5 | >4.0 | >3.0 |
| Toil Reduction | % reduction in repetitive work | >70% | >50% | >30% |
| Creative Work Ratio | % time on creative vs routine tasks | >80% | >60% | >40% |
| Work-Life Balance Impact | AI impact on work-life balance | 4.5+ | 4.0+ | 3.0+ |

**Key Insight:** Focus on human-AI collaboration satisfaction, toil reduction, and AI learning curve.

---

### P - Performance

Quality of AI-assisted outputs and human oversight.

| Metric | Description | Elite | High | Medium |
|--------|-------------|-------|------|--------|
| AI Code Acceptance Rate | % of AI suggestions accepted | >80% | >60% | >40% |
| AI Code Quality Score | Quality rating (1-5) | >4.5 | >4.0 | >3.0 |
| Human Review Catch Rate | Issues caught by human review | 5-15% | 15-25% | 25-40% |
| AI Code Defect Density | Defects per KLOC | <1 | <3 | <5 |

**Key Insight:** Human Review Catch Rate has an optimum range - too low suggests rubber-stamping; too high suggests poor AI quality.

---

### A - Activity

Orchestration quality, AI leverage, and workflow effectiveness.

| Metric | Description | Elite | High | Medium |
|--------|-------------|-------|------|--------|
| AI Leverage | % development work with AI | >80% | >60% | >40% |
| Prompt Reuse Rate | % prompts templated for reuse | >60% | >40% | >20% |
| Session Efficiency | Outcomes per AI session hour | >5 | >3 | >1 |
| Orchestration Complexity | Multi-step AI workflows | 10+ steps | 5-10 steps | 3-5 steps |

**Warning:** Traditional activity metrics (LOC, commits) are meaningless with AI. Focus on AI leverage and orchestration quality instead.

---

### C - Communication and Collaboration

Human-AI-Human collaboration, knowledge sharing, and feedback loops.

| Metric | Description | Elite | High | Medium |
|--------|-------------|-------|------|--------|
| AI Workflow Sharing Rate | % workflows shared with team | >70% | >50% | >30% |
| Prompt Library Contributions | Contributions per month | 10+ | 5-10 | 2-5 |
| Cross-team Pattern Adoption | % patterns from other teams | >50% | >30% | >15% |
| Feedback Loop Speed | Time to actionable feedback | <10 min | <30 min | <1 hr |

**Key Insight:** AI-assisted knowledge sharing amplifies team productivity.

---

### E - Efficiency and Flow

AI-augmented flow state, acceleration, and cycle time reduction.

| Metric | Description | Elite | High | Medium |
|--------|-------------|-------|------|--------|
| AI Acceleration Factor | Speed with AI vs without | >5x | >3x | >2x |
| Context Switch Frequency | Switches per day | <3 | <5 | <10 |
| AI-Augmented Flow Hours | Hours per day in flow | >6 | >4 | >2 |
| AI-Assisted Cycle Time | Start to production | <4 hrs | <1 day | <3 days |
| Weekly Time Savings | Hours saved per week | >15 | >10 | >5 |

---

## Anti-Patterns

Metrics to avoid in AI-assisted environments:

| Metric | Why It's Problematic |
|--------|---------------------|
| Lines of Code | AI can generate infinite LOC; meaningless metric |
| Number of Prompts | Easily gamed; more prompts ≠ better outcomes |
| AI Tool Usage Time | More time ≠ better; efficiency matters |
| Raw Commit Frequency | AI can generate infinite commits |
| Code Churn | AI refactoring inflates churn without value correlation |

## Usage

### Go API

```go
import frameworks "github.com/ProductBuildersHQ/productbuildershq-frameworks"

aispace := frameworks.MustAISPACE()

for _, dimension := range aispace.Dimensions {
    fmt.Printf("\n%s - %s\n", dimension.ID, dimension.Name)
    fmt.Printf("  Standard: %s\n", dimension.StandardDescription)
    fmt.Printf("  AI-Modified: %s\n", dimension.AIModification)

    for _, metric := range dimension.Metrics {
        fmt.Printf("  - %s: %s\n", metric.Name, metric.Description)
    }
}
```

### Assessment

To assess your team against AI-SPACE:

1. **Select metrics** from each dimension (2-3 per dimension recommended)
2. **Measure current state** using surveys, tooling data, and observability
3. **Compare against thresholds** for each metric
4. **Identify improvement areas** based on dimension gaps
5. **Avoid anti-patterns** - don't measure meaningless metrics

## Relationship to AI-DORA

AI-SPACE focuses on **developer productivity** while [AI-DORA](../ai-dora/index.md) measures **delivery performance**. Use both for a complete picture:

| AI-SPACE | Measures | AI-DORA | Measures |
|----------|----------|---------|----------|
| Satisfaction | Developer happiness | DF | Delivery speed |
| Performance | Code quality | CFR | Quality at scale |
| Activity | Work patterns | LT | Process efficiency |
| Communication | Collaboration | MTTR | Operational resilience |
| Efficiency | Flow state | - | - |

## References

- [The SPACE of Developer Productivity](https://queue.acm.org/detail.cfm?id=3454124) - Original paper
- [GitHub Well-Architected](https://wellarchitected.github.com/) - AI Leverage metrics
- [LinearB SPACE Framework](https://linearb.io/blog/space-framework) - Implementation guide
- [AI-SPACE Article](https://productbuildershq.com/frameworks/dora-space-ai-age)
