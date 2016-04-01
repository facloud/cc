package dep_installer

type aptInstaller struct{}

func NewAptInstaller() DepInstaller {
	return &aptInstaller{}
}

func (a *aptInstaller) Install(pkg string) error {
	return nil
}
