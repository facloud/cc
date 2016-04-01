package dep_installer

type DepInstaller interface {
	Install(dep string) error
}
