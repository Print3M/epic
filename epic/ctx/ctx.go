package ctx

var (
	ProjectPath  string
	OutputPath   string
	GccPath      string
	MingwGccPath string
	LinkerPath   string
	PICLinkOnly  bool     //TODO: Implement linking only, without building
	Modules      []string // TODO: Implement dynamic modules
	NoPIC        bool
	NoLoader     bool
	NoStandalone bool
	NoColors     bool // TODO: Implement colors / disable colors
	NoBanner     bool // TODO: Implement banner / disable banner
)
