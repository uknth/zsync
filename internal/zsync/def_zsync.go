package zsync

import "fmt"

type defaultZsyncer struct{}

func (dz *defaultZsyncer) Sync() error {
	fmt.Println("Hello World!!!")
	return nil
}

func NewDefaultSync() (Syncer, error) {
	return &defaultZsyncer{}, nil
}
