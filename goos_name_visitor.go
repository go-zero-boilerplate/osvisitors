package osvisitors

type GoOSNameVisitor struct{ name string }

func (g *GoOSNameVisitor) VisitWindows() { g.name = "windows" }
func (g *GoOSNameVisitor) VisitLinux()   { g.name = "linux" }
func (g *GoOSNameVisitor) VisitDarwin()  { g.name = "darwin" }
