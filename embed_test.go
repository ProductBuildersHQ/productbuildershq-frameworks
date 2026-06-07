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

func TestMustFunctions(t *testing.T) {
	// These should not panic since the JSON is valid
	_ = MustAIDora()
	_ = MustAISpace()
	_ = MustASDM()
	_ = MustPBMM()
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
