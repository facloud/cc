package main

type PackageSpec struct {
	Name         string
	Description  string
	Agents       []string
	Dependencies map[string]string
}

type Package struct {
	Spec PackageSpec
	Path string
}

func ReadInstalled(packageName, specInstPath string) (Package, error) {
	return Package{}, nil
}
