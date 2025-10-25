package ctx

var (
	ProjectPath  string
	OutputPath   string
	CompilerPath string
	LinkerPath   string
	Modules      []string // TODO: Implement dynamic modules
	NoPIC        bool
	NoLoader     bool
	NoStandalone bool
	NoColors     bool // TODO: Implement colors / disable colors
	NoBanner     bool // TODO: Implement banner / disable banner
)
