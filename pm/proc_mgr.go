package main

type ProcMgr interface {
	FindPids(programName string) ([]uint32, error)
	TermProc(pid uint32) error
	WaitForPid(pid uint32) error
	KillProc(pid uint32) error
}

type procMgr struct{}

func NewProcMgr() ProcMgr {
	return &procMgr{}
}

func (p *procMgr) FindPids(programName string) ([]uint32, error) {
	return []uint32{}, nil
}

func (p *procMgr) TermProc(pid uint32) error {
	return nil
}

func (p *procMgr) WaitForPid(pid uint32) error {
	return nil
}

func (p *procMgr) KillProc(pid uint32) error {
	return nil
}
