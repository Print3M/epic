package ctx

type CompilePicCtx struct {
	ProjectPath string
	OutputPath  string
	Modules     []string
}

type LinkPicCtx struct {
	ObjectsPath string
	OutputPath  string
	Modules     []string
}

type LoaderCtx struct {
	PayloadPath string
	OutputPath  string
}

type MonolithCtx struct {
	ProjectPath string
	OutputPath  string
}

var (
	NoColor          bool
	NoBanner         bool
	Debug            bool
	Version          string
	MingwGccPath     string
	MingwObjcopyPath string
	MingwLdPath      string
	CompilePIC       CompilePicCtx
	LinkPIC          LinkPicCtx
	Loader           LoaderCtx
	Monolith         MonolithCtx
)
