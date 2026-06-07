package frameworks

import "encoding/json"

// Documentation contains links to framework documentation.
type Documentation struct {
	Article string `json:"article,omitempty"`
	PDF     string `json:"pdf,omitempty"`
}

// Reference contains a citation or source reference.
type Reference struct {
	Title   string   `json:"title,omitempty"`
	URL     string   `json:"url,omitempty"`
	Authors []string `json:"authors,omitempty"`
	Year    int      `json:"year,omitempty"`
	Date    string   `json:"date,omitempty"`
	Note    string   `json:"note,omitempty"`
}

// BasedOn indicates the source framework for adapted frameworks.
type BasedOn struct {
	Framework string   `json:"framework,omitempty"`
	Name      string   `json:"name,omitempty"`
	Source    string   `json:"source,omitempty"`
	Authors   []string `json:"authors,omitempty"`
	Version   string   `json:"version,omitempty"`
	Year      int      `json:"year,omitempty"`
}

// MetricLevel defines thresholds for a performance level.
type MetricLevel struct {
	Threshold          float64 `json:"threshold,omitempty"`
	Min                float64 `json:"min,omitempty"`
	Max                float64 `json:"max,omitempty"`
	Operator           string  `json:"operator,omitempty"`
	Label              string  `json:"label,omitempty"`
	StandardComparison *struct {
		StandardThreshold  string `json:"standardThreshold,omitempty"`
		AccelerationFactor string `json:"accelerationFactor,omitempty"`
	} `json:"standardComparison,omitempty"`
}

// MetricLevels contains threshold definitions for all performance tiers.
type MetricLevels struct {
	Elite  MetricLevel `json:"elite,omitempty"`
	High   MetricLevel `json:"high,omitempty"`
	Medium MetricLevel `json:"medium,omitempty"`
	Low    MetricLevel `json:"low,omitempty"`
}

// =============================================================================
// AI-DORA Types
// =============================================================================

// AIDoraFramework represents the AI-modified DORA metrics framework.
type AIDoraFramework struct {
	Schema      string        `json:"$schema,omitempty"`
	Framework   string        `json:"framework,omitempty"`
	Name        string        `json:"name,omitempty"`
	Description string        `json:"description,omitempty"`
	Version     string        `json:"version,omitempty"`
	BasedOn     *BasedOn      `json:"basedOn,omitempty"`
	Category    string        `json:"category,omitempty"`
	Metrics     []DoraMetric  `json:"metrics,omitempty"`
	References  []Reference   `json:"references,omitempty"`
}

// DoraMetric represents a single DORA metric with AI modifications.
type DoraMetric struct {
	ID          string       `json:"id,omitempty"`
	Name        string       `json:"name,omitempty"`
	Description string       `json:"description,omitempty"`
	Direction   string       `json:"direction,omitempty"`
	Unit        string       `json:"unit,omitempty"`
	Levels      MetricLevels `json:"levels,omitempty"`
	AIEnablers  []string     `json:"aiEnablers,omitempty"`
}

// =============================================================================
// AI-SPACE Types
// =============================================================================

// AISpaceFramework represents the AI-modified SPACE developer productivity framework.
type AISpaceFramework struct {
	Schema       string           `json:"$schema,omitempty"`
	Framework    string           `json:"framework,omitempty"`
	Name         string           `json:"name,omitempty"`
	Description  string           `json:"description,omitempty"`
	Version      string           `json:"version,omitempty"`
	BasedOn      *BasedOn         `json:"basedOn,omitempty"`
	Category     string           `json:"category,omitempty"`
	Dimensions   []SpaceDimension `json:"dimensions,omitempty"`
	AntiPatterns []AntiPattern    `json:"antiPatterns,omitempty"`
	References   []Reference      `json:"references,omitempty"`
}

// SpaceDimension represents one of the SPACE dimensions (S, P, A, C, E).
type SpaceDimension struct {
	ID                  string        `json:"id,omitempty"`
	Name                string        `json:"name,omitempty"`
	Description         string        `json:"description,omitempty"`
	StandardDescription string        `json:"standardDescription,omitempty"`
	AIModification      string        `json:"aiModification,omitempty"`
	Warning             string        `json:"warning,omitempty"`
	Metrics             []SpaceMetric `json:"metrics,omitempty"`
}

// SpaceMetric represents a metric within a SPACE dimension.
type SpaceMetric struct {
	ID          string       `json:"id,omitempty"`
	Name        string       `json:"name,omitempty"`
	Description string       `json:"description,omitempty"`
	Unit        string       `json:"unit,omitempty"`
	Direction   string       `json:"direction,omitempty"`
	Source      string       `json:"source,omitempty"`
	Note        string       `json:"note,omitempty"`
	Levels      MetricLevels `json:"levels,omitempty"`
}

// AntiPattern describes a metric that should be avoided.
type AntiPattern struct {
	Metric string `json:"metric,omitempty"`
	Reason string `json:"reason,omitempty"`
}

// =============================================================================
// ASDM Types
// =============================================================================

// ASDMFramework represents the Autonomous Software Delivery Model.
type ASDMFramework struct {
	Schema        string       `json:"$schema,omitempty"`
	Framework     string       `json:"framework,omitempty"`
	Name          string       `json:"name,omitempty"`
	Description   string       `json:"description,omitempty"`
	Version       string       `json:"version,omitempty"`
	Category      string       `json:"category,omitempty"`
	Type          string       `json:"type,omitempty"`
	Documentation Documentation `json:"documentation,omitempty"`
	Levels        []ASDMLevel  `json:"levels,omitempty"`
	CaseStudies   []CaseStudy  `json:"caseStudies,omitempty"`
	References    []Reference  `json:"references,omitempty"`
}

// ASDMLevel represents a level in the ASDM progression.
type ASDMLevel struct {
	Level                   int            `json:"level,omitempty"`
	Name                    string         `json:"name,omitempty"`
	HumanRole               string         `json:"humanRole,omitempty"`
	DefiningCharacteristic  string         `json:"definingCharacteristic,omitempty"`
	HumanLoopPosition       string         `json:"humanLoopPosition,omitempty"`
	HumanLoopDescription    string         `json:"humanLoopDescription,omitempty"`
	Practices               []ASDMPractice `json:"practices,omitempty"`
	Artifacts               []string       `json:"artifacts,omitempty"`
	Rituals                 []string       `json:"rituals,omitempty"`
	Examples                []string       `json:"examples,omitempty"`
}

// ASDMPractice represents a practice with maturity level.
// For earlier levels, practices may be simple strings represented as Name only.
type ASDMPractice struct {
	Name     string `json:"name,omitempty"`
	Maturity string `json:"maturity,omitempty"`
}

// UnmarshalJSON handles both string and object practice formats.
func (p *ASDMPractice) UnmarshalJSON(data []byte) error {
	// Try string format first
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		p.Name = s
		return nil
	}
	// Try object format
	type practice ASDMPractice
	return json.Unmarshal(data, (*practice)(p))
}

// CaseStudy represents a real-world implementation example.
type CaseStudy struct {
	Name          string        `json:"name,omitempty"`
	Level         int           `json:"level,omitempty"`
	Date          string        `json:"date,omitempty"`
	Documentation Documentation `json:"documentation,omitempty"`
}

// =============================================================================
// Product Builder Maturity Model Types
// =============================================================================

// PBMMFramework represents the Product Builder Maturity Model.
type PBMMFramework struct {
	Schema            string             `json:"$schema,omitempty"`
	Framework         string             `json:"framework,omitempty"`
	Name              string             `json:"name,omitempty"`
	Description       string             `json:"description,omitempty"`
	Version           string             `json:"version,omitempty"`
	Category          string             `json:"category,omitempty"`
	Type              string             `json:"type,omitempty"`
	Documentation     Documentation      `json:"documentation,omitempty"`
	Paths             PBMMPaths          `json:"paths,omitempty"`
	Convergence       PBMMConvergence    `json:"convergence,omitempty"`
	Levels            []PBMMLevel        `json:"levels,omitempty"`
	FrameworkMappings []FrameworkMapping `json:"frameworkMappings,omitempty"`
	References        []Reference        `json:"references,omitempty"`
}

// PBMMPaths defines the dual paths in the maturity model.
type PBMMPaths struct {
	PM       PBMMPath `json:"pm,omitempty"`
	Engineer PBMMPath `json:"engineer,omitempty"`
}

// PBMMPath describes a single path in the model.
type PBMMPath struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// PBMMConvergence describes where the paths converge.
type PBMMConvergence struct {
	Level       int    `json:"level,omitempty"`
	Description string `json:"description,omitempty"`
}

// PBMMLevel represents a level in the Product Builder progression.
type PBMMLevel struct {
	Level           int      `json:"level,omitempty"`
	Name            string   `json:"name,omitempty"`
	PMPath          string   `json:"pmPath,omitempty"`
	EngineerPath    string   `json:"engineerPath,omitempty"`
	Converged       bool     `json:"converged,omitempty"`
	Characteristics []string `json:"characteristics,omitempty"`
}

// FrameworkMapping maps PBMM metrics to other frameworks.
type FrameworkMapping struct {
	Metric  string  `json:"metric,omitempty"`
	AIDora  *string `json:"aiDora,omitempty"`
	AISpace *string `json:"aiSpace,omitempty"`
}
