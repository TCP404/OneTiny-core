package verify

type Handler interface {
	Handle() error
}

var _ Handler = (*pathHandler)(nil)
var _ Handler = (*portHandler)(nil)
var _ Handler = (*upsHandler)(nil)

func NewPathHandler(path string) Handler { return &pathHandler{rootPath: path} }

func NewPortHandler(port int) Handler { return &portHandler{port: port} }

func NewUPSHandler(weight uint8) *upsHandler { return &upsHandler{weight: ups(weight)} }
