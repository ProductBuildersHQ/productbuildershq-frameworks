# PDLC (Product Development Lifecycle)

A six-stage development lifecycle spanning product definition through operations, native to ProductBuildersHQ. PDLC is the **canonical stage-ID hub**: other lifecycle and posture frameworks (AIDLC, and consumer-owned overlays like Threat Model Spec's ASPM security-posture domains) map their stages onto PDLC's, so any two frameworks can be crosswalked through PDLC without a pairwise mapping.

## Overview

PDLC splits each of AWS AI-DLC's three phases into two parallel lenses — **product** and **builder** — so every stage has exactly one accountable role, while the two stages within a phase can proceed together rather than gating each other:

```
┌──────────────────────┬──────────────────────┬──────────────────────────────┐
│      INCEPTION        │     CONSTRUCTION      │          OPERATION           │
├───────────┬───────────┼───────────┬───────────┼───────────────┬──────────────┤
│ product-  │ builder-  │implementa-│deployment │  builder-     │  product-    │
│definition │definition │tion       │           │ operations    │ operations   │
│ (product) │ (builder) │(builder)  │(builder)  │  (builder)    │  (product)   │
└───────────┴───────────┴───────────┴───────────┴───────────────┴──────────────┘
                                                    └─── run in parallel ───┘
```

## Stages

### 1. Product Definition (`product-definition`)

Discovery through baseline handoff — the existing pdlc specification's seven detailed sub-stages: Discovery, Definition, Experience & Prototype, API, Documentation, Localization, Baseline & Handoff. Governed in full by the [pdlc repository](https://github.com/ProductBuildersHQ/pdlc); this framework entry references it rather than duplicating its normative content.

**Owner:** Product person. **Gate:** Baseline Approval (readiness report green, revision-pinned).

### 2. Builder Definition (`builder-definition`)

Consumes the approved Product Baseline and produces the authoritative technical design. Key deliverable pattern mirrors Product Definition's normative-spec-plus-advisory-evidence shape: the finalized **API Contract** (OpenAPI, refined from Product Definition's advisory draft) is normative, and a **Reference SDK Client** auto-generated from that contract is advisory — used for reference implementation and acceptance testing, proving the contract is implementable without itself being a requirement.

**Owner:** Builder person. **Gate:** Technical Design Review.

### 3. Implementation (`implementation`)

Code is written against the Builder Definition technical contract. Overlaid by Application Security Posture Management (ASPM) domains 1–5 (git posture, code security, secret/PII scan, open source security, SBOM) — the overlay is owned by consuming security-analysis tooling (e.g. Threat Model Spec), not by this framework.

**Owner:** Builder person. **Gate:** Code Review / CI Gate.

### 4. Deployment (`deployment`)

Built artifacts are released to target environments. Overlaid by ASPM domains 6–9 (IaC scan, CI/CD posture, container security, artifact security).

**Owner:** Builder person. **Gate:** Release Gate.

### 5. Builder Operations (`builder-operations`)

Infrastructure, security, and reliability operations for the deployed system: monitoring, incident response, cloud security posture. Overlaid by ASPM domain 10 (cloud context) plus dynamic testing (DAST, penetration testing, red teaming), which sits alongside ASPM rather than within it. **Runs concurrently with Product Operations**, not sequentially after it.

**Owner:** Builder person.

### 6. Product Operations (`product-operations`)

Adoption, usage, and feedback operations for the shipped product: activation, retention, feature usage, PMF signal, and customer feedback synthesis — closing the loop back into the next Product Definition baseline revision. **Runs concurrently with Builder Operations.**

**Owner:** Product person. **Gate:** Baseline Revision Trigger (on significant definition-vs-reality drift).

## Files in This Directory

| File | Purpose |
|------|---------|
| [`pdlc-framework.json`](pdlc-framework.json) | Framework definition: phases, deliverables, gates, dependency graph, AIDLC crosswalk |

A PIDL process model and per-phase deep-dive docs (mirroring [`frameworks/aidlc/phases/`](../aidlc/phases/)) are a documented follow-up — the pdlc repository already owns a PIDL model for its Product Definition sub-stages (`model/pdlc-lifecycle.pidl.json`), and duplicating it here risks drift. This entry is intentionally the coarse six-stage catalog view; detailed sub-stage content stays normatively owned by the pdlc repository.

## Relationship to AIDLC

Every PDLC stage maps to exactly one AIDLC phase, and every AIDLC phase is claimed by exactly two PDLC stages:

| AIDLC phase | PDLC stages |
|-------------|-------------|
| Inception | Product Definition, Builder Definition |
| Construction | Implementation, Deployment |
| Operation | Builder Operations, Product Operations (parallel) |

See `relatedFrameworks[0].stageMapping` in [`pdlc-framework.json`](pdlc-framework.json) for the machine-readable crosswalk.

## Relationship to ASPM

Application Security Posture Management is not cataloged in this repository — it is security-domain-specific and owned by [Threat Model Spec](https://github.com/grokify/threat-model-spec), which maps its 10 ASPM domains onto the three builder-side stages here (Implementation, Deployment, Builder Operations). This keeps `productbuildershq-frameworks` scoped to general product/engineering-lifecycle frameworks.

## References

- [PDLC Specification](https://github.com/ProductBuildersHQ/pdlc)
- [AWS AI-Driven Development Lifecycle](https://docs.aws.amazon.com/prescriptive-guidance/latest/ai-driven-software-development/)
- [AIDLC framework entry](../aidlc/) in this repository
