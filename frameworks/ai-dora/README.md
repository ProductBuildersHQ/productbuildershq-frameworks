# AI-DORA: AI-Modified DevOps Research and Assessment

AI-DORA adapts the industry-standard DORA metrics for AI-assisted software delivery, providing updated baselines that reflect the acceleration AI enables.

## The Four Key Metrics

### 1. Deployment Frequency

| Level | Standard DORA | AI-DORA | AI Acceleration |
|-------|---------------|---------|-----------------|
| Elite | On-demand (multiple/day) | On-demand (10+/day) | 3-5x |
| High | Daily to weekly | Multiple per day | 5-10x |
| Medium | Weekly to monthly | Daily | 7x |
| Low | Monthly+ | Weekly | 4x |

**Why AI changes this**: AI-assisted code generation, automated testing, and intelligent CI/CD reduce the friction of each deployment.

### 2. Lead Time for Changes

| Level | Standard DORA | AI-DORA | AI Acceleration |
|-------|---------------|---------|-----------------|
| Elite | < 1 day | < 2 hours | 12x |
| High | 1 day - 1 week | < 1 day | 7x |
| Medium | 1 week - 1 month | 1-3 days | 10x |
| Low | 1 month+ | 1 week | 4x |

**Why AI changes this**: AI code generation eliminates boilerplate time, AI-assisted code review accelerates PR cycles, and AI testing catches issues pre-merge.

### 3. Change Failure Rate

| Level | Standard DORA | AI-DORA | AI Improvement |
|-------|---------------|---------|----------------|
| Elite | 0-5% | 0-2% | 2.5x better |
| High | 5-10% | 2-5% | 2x better |
| Medium | 10-15% | 5-8% | 1.8x better |
| Low | 15%+ | 8-12% | 1.5x better |

**Why AI changes this**: AI code review catches bugs before merge, AI-generated tests improve coverage, and AI assists with safer deployment patterns.

### 4. Mean Time to Recovery (MTTR)

| Level | Standard DORA | AI-DORA | AI Acceleration |
|-------|---------------|---------|-----------------|
| Elite | < 1 hour | < 15 minutes | 4x |
| High | < 1 day | < 2 hours | 12x |
| Medium | 1 day - 1 week | < 1 day | 7x |
| Low | 1 week+ | 1-3 days | 3x |

**Why AI changes this**: AI-assisted debugging identifies root causes faster, AI generates fix suggestions, and AI-powered runbooks accelerate remediation.

## Framework Definition

```json
{
  "framework": "AI_DORA",
  "name": "AI-Modified DORA",
  "version": "1.0.0",
  "basedOn": {
    "framework": "DORA",
    "source": "https://dora.dev"
  },
  "metrics": [
    {
      "id": "DF",
      "name": "Deployment Frequency",
      "description": "How often code is deployed to production",
      "levels": {
        "elite": {"value": "10+/day", "comparison": "3-5x standard"},
        "high": {"value": "multiple/day", "comparison": "5-10x standard"},
        "medium": {"value": "daily", "comparison": "7x standard"},
        "low": {"value": "weekly", "comparison": "4x standard"}
      }
    },
    {
      "id": "LT",
      "name": "Lead Time for Changes",
      "description": "Time from commit to production",
      "levels": {
        "elite": {"value": "< 2 hours", "comparison": "12x faster"},
        "high": {"value": "< 1 day", "comparison": "7x faster"},
        "medium": {"value": "1-3 days", "comparison": "10x faster"},
        "low": {"value": "1 week", "comparison": "4x faster"}
      }
    },
    {
      "id": "CFR",
      "name": "Change Failure Rate",
      "description": "Percentage of deployments causing failures",
      "levels": {
        "elite": {"value": "0-2%", "comparison": "2.5x better"},
        "high": {"value": "2-5%", "comparison": "2x better"},
        "medium": {"value": "5-8%", "comparison": "1.8x better"},
        "low": {"value": "8-12%", "comparison": "1.5x better"}
      }
    },
    {
      "id": "MTTR",
      "name": "Mean Time to Recovery",
      "description": "Time to recover from failures",
      "levels": {
        "elite": {"value": "< 15 min", "comparison": "4x faster"},
        "high": {"value": "< 2 hours", "comparison": "12x faster"},
        "medium": {"value": "< 1 day", "comparison": "7x faster"},
        "low": {"value": "1-3 days", "comparison": "3x faster"}
      }
    }
  ]
}
```

## AI Enablers for Each Metric

| Metric | AI Capabilities That Enable Improvement |
|--------|----------------------------------------|
| Deployment Frequency | AI code gen, auto-testing, intelligent CI/CD |
| Lead Time | AI pair programming, AI code review, AI test generation |
| Change Failure Rate | AI static analysis, AI-generated tests, AI security scanning |
| MTTR | AI debugging, AI root cause analysis, AI-powered runbooks |

## Measuring Your AI Acceleration Factor

Compare your metrics against both standard DORA and AI-DORA:

```
AI Acceleration Factor = Standard DORA Target / Your Actual Performance
```

If your lead time is 4 hours:
- vs Standard Elite (1 day): You're 6x better than traditional elite
- vs AI-DORA Elite (2 hours): You're at 50% of AI-assisted elite

## References

- [DORA Research](https://dora.dev)
- [Accelerate Book](https://itrevolution.com/product/accelerate/)
- [State of DevOps Report](https://cloud.google.com/devops/state-of-devops)
