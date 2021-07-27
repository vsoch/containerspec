package spec

// A Microarchitecture include a vendor, compiler, etc.
type Microarchitecture struct {
	Name       string                `json:"name"`
	From       []string              `json:"from"`
	Vendor     []string              `json:"vendor"`
	Generation int                   `json:"generation"`
	Features   []string              `json:"features"`
	Compilers  map[string][]Compiler `json:"compilers"`
}

// Compiler holds information about compilers
type Compiler struct {
	Name     string   `json:"name"`
	Versions string   `json:"versions"`
	Flags    string   `json:"flags"`
	Family   []string `json:"family"`
	Warnings []string `json:"warnings"`
}

// FeatureAlias holds aliases for features
type FeatureAlias struct {
	Reason   string   `json:reason`
	AnyOf    []string `json:any_of`
	Families []string `json:families`
}

// Conversions that map some platform specific values to canonical values
type Conversion map[string]string
