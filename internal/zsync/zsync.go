package zsync

type Syncer interface {
	Sync() error
}
