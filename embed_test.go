package frameworks

import (
	"testing"
)

func TestAIDora(t *testing.T) {
	f, err := AIDora()
	if err != nil {
		t.Fatalf("AIDora() error: %v", err)
	}
	if f.Framework != "AI_DORA" {
		t.Errorf("Framework = %q, want %q", f.Framework, "AI_DORA")
	}
	if f.Name != "AI-Modified DORA" {
		t.Errorf("Name = %q, want %q", f.Name, "AI-Modified DORA")
	}
	if len(f.Metrics) != 4 {
		t.Errorf("len(Metrics) = %d, want 4", len(f.Metrics))
	}
	// Verify metric IDs
	expectedMetrics := []string{"DF", "LT", "CFR", "MTTR"}
	for i, m := range f.Metrics {
		if m.ID != expectedMetrics[i] {
			t.Errorf("Metrics[%d].ID = %q, want %q", i, m.ID, expectedMetrics[i])
		}
	}
}

func TestAISpace(t *testing.T) {
	f, err := AISpace()
	if err != nil {
		t.Fatalf("AISpace() error: %v", err)
	}
	if f.Framework != "AI_SPACE" {
		t.Errorf("Framework = %q, want %q", f.Framework, "AI_SPACE")
	}
	if f.Version != "1.1.0" {
		t.Errorf("Version = %q, want %q", f.Version, "1.1.0")
	}
	if len(f.Dimensions) != 5 {
		t.Errorf("len(Dimensions) = %d, want 5", len(f.Dimensions))
	}
	// Verify dimension IDs (SPACE)
	expectedDimensions := []string{"S", "P", "A", "C", "E"}
	for i, d := range f.Dimensions {
		if d.ID != expectedDimensions[i] {
			t.Errorf("Dimensions[%d].ID = %q, want %q", i, d.ID, expectedDimensions[i])
		}
	}
	// Each dimension should have 5 metrics (25 total)
	totalMetrics := 0
	for _, d := range f.Dimensions {
		totalMetrics += len(d.Metrics)
	}
	if totalMetrics != 25 {
		t.Errorf("total metrics = %d, want 25", totalMetrics)
	}
}

func TestASDM(t *testing.T) {
	f, err := ASDM()
	if err != nil {
		t.Fatalf("ASDM() error: %v", err)
	}
	if f.Framework != "ASDM" {
		t.Errorf("Framework = %q, want %q", f.Framework, "ASDM")
	}
	if f.Type != "original" {
		t.Errorf("Type = %q, want %q", f.Type, "original")
	}
	if len(f.Levels) != 7 {
		t.Errorf("len(Levels) = %d, want 7", len(f.Levels))
	}
	// Verify level 1 is Agile/Scrum
	if f.Levels[0].Name != "Agile/Scrum" {
		t.Errorf("Levels[0].Name = %q, want %q", f.Levels[0].Name, "Agile/Scrum")
	}
	// Verify level 7 is Autonomous Operations
	if f.Levels[6].Name != "Autonomous Operations" {
		t.Errorf("Levels[6].Name = %q, want %q", f.Levels[6].Name, "Autonomous Operations")
	}
	// Verify case studies
	if len(f.CaseStudies) != 2 {
		t.Errorf("len(CaseStudies) = %d, want 2", len(f.CaseStudies))
	}
}

func TestPBMM(t *testing.T) {
	f, err := PBMM()
	if err != nil {
		t.Fatalf("PBMM() error: %v", err)
	}
	if f.Framework != "PBMM" {
		t.Errorf("Framework = %q, want %q", f.Framework, "PBMM")
	}
	if f.Type != "original" {
		t.Errorf("Type = %q, want %q", f.Type, "original")
	}
	if len(f.Levels) != 7 {
		t.Errorf("len(Levels) = %d, want 7", len(f.Levels))
	}
	// Verify dual paths exist
	if f.Paths.PM.Name == "" {
		t.Error("Paths.PM.Name is empty")
	}
	if f.Paths.Engineer.Name == "" {
		t.Error("Paths.Engineer.Name is empty")
	}
	// Verify convergence at level 5
	if f.Convergence.Level != 5 {
		t.Errorf("Convergence.Level = %d, want 5", f.Convergence.Level)
	}
	// Verify L5 and L6 are converged
	for _, l := range f.Levels {
		if l.Level >= 5 && !l.Converged {
			t.Errorf("Level %d should be converged", l.Level)
		}
		if l.Level < 5 && l.Converged {
			t.Errorf("Level %d should not be converged", l.Level)
		}
	}
}

func TestAIDLC(t *testing.T) {
	f, err := AIDLC()
	if err != nil {
		t.Fatalf("AIDLC() error: %v", err)
	}
	if f.Framework != "AIDLC" {
		t.Errorf("Framework = %q, want %q", f.Framework, "AIDLC")
	}
	if f.Type != "adapted" {
		t.Errorf("Type = %q, want %q", f.Type, "adapted")
	}
	if len(f.Phases) != 3 {
		t.Errorf("len(Phases) = %d, want 3", len(f.Phases))
	}
	// Verify phase IDs
	expectedPhases := []string{"inception", "construction", "operations"}
	for i, p := range f.Phases {
		if p.ID != expectedPhases[i] {
			t.Errorf("Phases[%d].ID = %q, want %q", i, p.ID, expectedPhases[i])
		}
	}
	// Count total deliverables (should be 12)
	totalDeliverables := 0
	for _, p := range f.Phases {
		totalDeliverables += len(p.Deliverables)
	}
	if totalDeliverables != 12 {
		t.Errorf("total deliverables = %d, want 12", totalDeliverables)
	}
	// Verify cost estimates
	if f.CostEstimates == nil {
		t.Error("CostEstimates is nil")
	} else if f.CostEstimates.Totals.InputTokens != 127000 {
		t.Errorf("CostEstimates.Totals.InputTokens = %d, want 127000", f.CostEstimates.Totals.InputTokens)
	}
}

func TestPDLC(t *testing.T) {
	f, err := PDLC()
	if err != nil {
		t.Fatalf("PDLC() error: %v", err)
	}
	if f.Framework != "PDLC" {
		t.Errorf("Framework = %q, want %q", f.Framework, "PDLC")
	}
	if f.Type != "original" {
		t.Errorf("Type = %q, want %q", f.Type, "original")
	}
	if len(f.Phases) != 6 {
		t.Errorf("len(Phases) = %d, want 6", len(f.Phases))
	}

	// Verify phase IDs and order.
	expectedPhases := []string{
		"product-definition", "builder-definition", "implementation",
		"deployment", "builder-operations", "product-operations",
	}
	for i, p := range f.Phases {
		if p.ID != expectedPhases[i] {
			t.Errorf("Phases[%d].ID = %q, want %q", i, p.ID, expectedPhases[i])
		}
		if p.Order != i+1 {
			t.Errorf("Phases[%d].Order = %d, want %d", i, p.Order, i+1)
		}
	}

	// Product Definition carries the seven pdlc sub-stages.
	if len(f.Phases[0].SubStages) != 7 {
		t.Errorf("Product Definition SubStages = %d, want 7", len(f.Phases[0].SubStages))
	}

	// Role split: exactly 2 product-owned, 4 builder-owned.
	roleCounts := map[string]int{}
	for _, p := range f.Phases {
		roleCounts[p.Role]++
	}
	if roleCounts["product"] != 2 {
		t.Errorf("product-role phases = %d, want 2", roleCounts["product"])
	}
	if roleCounts["builder"] != 4 {
		t.Errorf("builder-role phases = %d, want 4", roleCounts["builder"])
	}

	// Builder Operations and Product Operations are parallel, not sequential.
	builderOps := f.Phases[4]
	productOps := f.Phases[5]
	if builderOps.ID != "builder-operations" || len(builderOps.ParallelWith) == 0 || builderOps.ParallelWith[0] != "product-operations" {
		t.Errorf("builder-operations.ParallelWith = %v, want [product-operations]", builderOps.ParallelWith)
	}
	if productOps.ID != "product-operations" || len(productOps.ParallelWith) == 0 || productOps.ParallelWith[0] != "builder-operations" {
		t.Errorf("product-operations.ParallelWith = %v, want [builder-operations]", productOps.ParallelWith)
	}

	// AIDLC crosswalk: every PDLC phase maps to exactly one AIDLC phase, and
	// each AIDLC phase (inception, construction, operations) claims exactly 2.
	if len(f.RelatedFrameworks) != 1 || f.RelatedFrameworks[0].Framework != "AIDLC" {
		t.Fatalf("RelatedFrameworks = %+v, want a single AIDLC relation", f.RelatedFrameworks)
	}
	aiDlcCounts := map[string]int{}
	for _, c := range f.RelatedFrameworks[0].StageMapping {
		if len(c.MapsTo) != 1 {
			t.Errorf("StageMapping[%s].MapsTo = %v, want exactly 1 target", c.Stage, c.MapsTo)
			continue
		}
		aiDlcCounts[c.MapsTo[0]]++
	}
	for _, phase := range []string{"inception", "construction", "operations"} {
		if aiDlcCounts[phase] != 2 {
			t.Errorf("AIDLC phase %q claimed by %d PDLC stages, want 2", phase, aiDlcCounts[phase])
		}
	}

	// Dependency graph closes the loop back to product-definition.
	if f.Dependencies == nil || len(f.Dependencies.Graph) == 0 {
		t.Fatal("Dependencies.Graph is empty")
	}
	var closesLoop bool
	for _, d := range f.Dependencies.Graph {
		if d.To == "product-definition" {
			closesLoop = true
		}
	}
	if !closesLoop {
		t.Error("dependency graph never feeds back into product-definition")
	}
}

func TestMustFunctions(t *testing.T) {
	// These should not panic since the JSON is valid
	_ = MustAIDora()
	_ = MustAISpace()
	_ = MustASDM()
	_ = MustPBMM()
	_ = MustAIDLC()
	_ = MustPDLC()
}

func TestRawJSON(t *testing.T) {
	if len(AIDoraJSON()) == 0 {
		t.Error("AIDoraJSON() is empty")
	}
	if len(AISpaceJSON()) == 0 {
		t.Error("AISpaceJSON() is empty")
	}
	if len(ASDMJSON()) == 0 {
		t.Error("ASDMJSON() is empty")
	}
	if len(PBMMJSON()) == 0 {
		t.Error("PBMMJSON() is empty")
	}
	if len(AIDLCJSON()) == 0 {
		t.Error("AIDLCJSON() is empty")
	}
	if len(PDLCJSON()) == 0 {
		t.Error("PDLCJSON() is empty")
	}
}

func TestFS(t *testing.T) {
	fs := FS()

	// Test reading a file from the embedded FS
	data, err := fs.ReadFile("frameworks/ai-dora/ai-dora.json")
	if err != nil {
		t.Fatalf("fs.ReadFile() error: %v", err)
	}
	if len(data) == 0 {
		t.Error("fs.ReadFile() returned empty data")
	}
}
