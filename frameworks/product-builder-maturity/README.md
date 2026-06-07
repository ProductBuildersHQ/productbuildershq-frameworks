# Product Builder Maturity Model

A maturity model for measuring progress from traditional roles to AI-enabled product builders.

## Overview

The Product Builder Maturity Model tracks the journey from **Observer** (L0) to **Platform Leader** (L6), with two parallel paths:

- **PM Path**: Product managers gaining technical fluency
- **Engineer Path**: Engineers gaining product judgment

Both paths converge at L5-L6 as full **Product Builders**.

## Levels

| Level | Name | PM Path | Engineer Path |
|-------|------|---------|---------------|
| L0 | Observer | Uses AI for docs/brainstorming | Uses AI for docs/brainstorming |
| L1 | Operator | Structured AI workflows | Structured AI workflows |
| L2 | Prototype Builder | Ships prototypes before eng | Ships prototypes before PM |
| L3 | Technical Builder | AI Developer (writes code) | Product-Aware Engineer |
| L4 | Production Engineer | AI Product Engineer | Full-Stack Product Engineer |
| L5 | Product Builder | End-to-end ownership | End-to-end ownership |
| L6 | Platform Leader | Enables other builders | Enables other builders |

## Framework Mappings

Product Builder metrics map to both standard and AI-modified frameworks:

| Product Builder Metric | AI-DORA | AI-SPACE |
|------------------------|---------|----------|
| Features shipped without handoff | Deployment Frequency | Activity |
| Time to prototype | Lead Time | Efficiency |
| Production incidents on owned systems | Change Failure Rate | Performance |
| Mean time to resolution | MTTR | Efficiency |
| Prompt reuse rate | - | Activity |
| Stakeholder satisfaction | - | Satisfaction |

## Canonical Source

- **Article**: [Product Builder Maturity Model](https://productbuildershq.com/frameworks/product-builder-maturity-model)
- **PDF**: [Product Builder Maturity Model](https://productbuildershq.com/papers/product-builder-maturity-model.pdf)
- **Machine-readable spec**: [`product-builder-maturity.json`](product-builder-maturity.json)

## Usage with PRISM

The Product Builder Maturity Model is consumed by [PRISM](https://github.com/grokify/prism) for:

- Dashboard visualization
- Progress tracking
- Maturity assessment
- Roadmap planning

## Related Frameworks

- [AI-DORA](../ai-dora/) - AI-modified software delivery metrics
- [AI-SPACE](../ai-space/) - AI-modified developer productivity metrics
