package container

import "runtime"

func Init() {
	// Make sure we are running on a single thread.
	// This is required for user namespace updates to properly apply to all future operations
	// since namespaces changes don't propagate across threads.
	runtime.GOMAXPROCS(1)
	runtime.LockOSThread()
}
