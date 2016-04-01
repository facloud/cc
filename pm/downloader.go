package main

type Downloader interface {
	Download(url string) (Package, error)
}

type downloader struct{}

func NewDownloader() Downloader {
	return &downloader{}
}

func (pd *downloader) Download(url string) (Package, error) {
	return Package{}, nil
}
