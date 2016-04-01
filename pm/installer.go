package main

type Installer interface {
	InstallBin(pkg Package) error
	InstallCfg(pkg Package) error
	InstallAssets(pkg Package) error
	InstallCtl(pkg Package) error
	InstallSpec(pkg Package) error
}

type installer struct{}

func NewInstaller(
	binInstPath string,
	cfgInstPath string,
	assetsInstPath string,
	ctlInstPath string,
	specInstPath string,
) Installer {
	return nil
}

func (i *installer) InstallBin(pkg Package) error {
	return nil
}

func (i *installer) InstallCfg(pkg Package) error {
	return nil

}

func (i *installer) InstallAssets(pkg Package) error {
	return nil

}

func (i *installer) InstallCtl(pkg Package) error {
	return nil

}

func (i *installer) InstallSpec(pkg Package) error {
	return nil

}
