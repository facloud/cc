package main

import (
	"flag"
	"log"
	"os"
	"sync"
)

func downloadPackage(
	logger *log.Logger, wg *sync.WaitGroup, downlaoder Downloader,
	packageURL string, packageChan chan *Package,
) {
	defer wg.Done()

	logger.Printf("Downloading package from '%s'", packageURL)
	packageChan <- nil
}

func killProcess(
	logger *log.Logger, wg *sync.WaitGroup, procMgr ProcMgr, timeout int,
	pid uint32,
) {
	defer wg.Done()

	logger.Printf("Killing process with PID %d", pid)
}

func runInstallationStep(
	logger *log.Logger, wg *sync.WaitGroup, cb func(Package) error, pkg Package,
) {
	defer wg.Done()

	logger.Print("Installation step")
}

func main() {
	packageName := flag.String("package", "", "Package name")
	packageURL := flag.String("url", "", "Package download URL")
	timeout := flag.Int("timeout", 10,
		"Time in seconds that cc-pm waits before killing the running process")
	flag.Parse()

	logger := log.New(os.Stdout, "[cc-pm] ", 0)
	procMgr := NewProcMgr()
	downloader := NewDownloader()
	installer := NewInstaller(
		"/tmp/binInstPath",
		"/tmp/cfgInstPath",
		"/tmp/assetsInstPath",
		"/tmp/ctlInstPath",
		"/tmp/specInstPath",
	)

	pids, err := procMgr.FindPids(*packageName)
	if err != nil {
		logger.Panicf("getting runnig process pids: %s", err)
	}

	wg := new(sync.WaitGroup)
	packageChan := make(chan *Package, 1)

	wg.Add(len(pids) + 1)
	go downloadPackage(logger, wg, downloader, *packageURL, packageChan)
	for _, pid := range pids {
		go killProcess(logger, wg, procMgr, *timeout, pid)
	}
	wg.Wait()

	pkg := <-packageChan
	if pkg != nil {
		wg.Add(5)
		go runInstallationStep(logger, wg, installer.InstallAssets, *pkg)
		go runInstallationStep(logger, wg, installer.InstallBin, *pkg)
		go runInstallationStep(logger, wg, installer.InstallCfg, *pkg)
		go runInstallationStep(logger, wg, installer.InstallSpec, *pkg)
		go runInstallationStep(logger, wg, installer.InstallCtl, *pkg)
		wg.Wait()
	}
}
