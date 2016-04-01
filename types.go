package cc

import "net"

type Credentials struct {
	AccessToken string
	SecretKey   string
}

type VPCSepc struct {
	Name string
	CIDR net.IPNet
}

type VolumeSpec struct {
	Name        string
	SizeInBytes uint64
	Persistent  bool
}

type ResourcesSpec struct {
	CPUInShares   uint32
	MemoryInBytes uint64
	DiskInBytes   uint64
}

type ProcessSpec struct {
	Path string
	Args []string
	User string
	Dir  string
	Env  []string
}

type VolumeID string

type VPCID string

type ContainerSpec struct {
	Image     string
	VPCID     VPCID
	Volumes   []VolumeID
	Resources ResourcesSpec
	Process   ProcessSpec
}
