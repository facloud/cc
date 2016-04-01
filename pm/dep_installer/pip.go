package dep_installer

type pipInstaller struct{}

func NewPipInstaller() DepInstaller {
	return &pipInstaller{}
}

func (a *pipInstaller) Install(pkg string) error {
	return nil
}
