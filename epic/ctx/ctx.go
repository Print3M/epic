package ctx

var (
	ProjectPath  string
	OutputPath   string
	CompilerPath string
	LinkerPath   string
	ObjcopyPath  string
	PICLinkOnly  bool //TODO: Implement linking only, without building
	Modules      []string
	NoPIC        bool
	NoLoader     bool
	NoNonPIC     bool
	NoColors     bool // TODO: Implement colors / disable colors
	NoBanner     bool // TODO: Implement banner / disable banner
	Debug        bool
)
