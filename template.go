package elit

// Template is root object
type Template struct {
	Template string   `json:"template"`
	Settings Settings `json:"settings"`
	Mappings Mappings `json:"mappings"`
}

// Mappings has property
type Mappings struct {
	Properties map[string]Property `json:"properties"`
}

// Settings node settings
type Settings struct {
	NumberOfShards   uint      `json:"number_of_shards"`
	NumberOfReplicas uint      `json:"number_of_replicas,omitempty"`
	Analysis         *Analysis `json:"analysis,omitempty"`
}

// Analysis settings
type Analysis struct {
	Tokenizer map[string]Tokenizer `json:"tokenizer,omitempty"`
	Analyzer  map[string]Analyzer  `json:"analyzer,omitempty"`
	Filter    map[string]Filter    `json:"filter,omitempty"`
}

// Analyzer .
type Analyzer struct {
	Type       *string  `json:"type,omitempty"`
	Tokenizer  string   `json:"tokenizer"`
	Filter     []string `json:"filter"`
	CharFilter []string `json:"char_filter"`
}

// Tokenizer .
type Tokenizer struct {
	Type string `json:"type"`
}

// Filter for analysis
type Filter struct {
	Type        FilterType `json:"type"`
	Format      string     `json:"format,omitempty"`
	SynonymPath string     `json:"synonym_path,omitempty"`
	Synonyms    []Synonym  `json:"synonyms,omitempty"`
}

// Source .
type Source struct {
	Enabled bool `json:"enabled"`
}

// All .
type All struct {
	Enabled bool `json:"enabled"`
}

// Property .
type Property struct {
	All         *All                `json:"_all,omitempty"`
	Type        PropertyType        `json:"type,omitempty"`
	Analyzer    string              `json:"analyzer,omitempty"`
	Format      string              `json:"format,omitempty"`
	FieldData   bool                `json:"fielddata,omitempty"`
	Fields      map[string]Property `json:"fields,omitempty"`
	Properies   map[string]Property `json:"properties,omitempty"`
	IgnoreAbove int                 `json:"ignore_above,omitempty"`
}

// PropertyType .
type PropertyType string

// FilterType .
type FilterType string

const (
	PropertyTypeDate      PropertyType = "date"
	PropertyTypeInteger32 PropertyType = "integer"
	PropertyTypeInteger64 PropertyType = "long"
	PropertyTypeFloat32   PropertyType = "float"
	PropertyTypeFloat64   PropertyType = "double"
	PropertyTypeText      PropertyType = "text"
	PropertyTypeGeoPoint  PropertyType = "geo_point"
	PropertyTypeNested    PropertyType = "nested"
	PropertyTypeKeyword   PropertyType = "keyword"
	PropertyTypeBoolean   PropertyType = "boolean"

	FilterTypeSynonym FilterType = "synonym"
)
