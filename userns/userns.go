package userns

import "github.com/AnatoleLucet/sudont/user"

// User namespace
type UserNS struct {
	UID   int
	GID   int
	SGIDs []int

	// The original user namespace from which this one was instantiated
	Source *UserNS
}

func New(uid, gid int, sgids []int) (*UserNS, error) {
	source, err := user.Current()
	if err != nil {
		return nil, err
	}

	return &UserNS{
		UID:   uid,
		GID:   gid,
		SGIDs: sgids,

		Source: &UserNS{
			UID:   source.UID,
			GID:   source.GID,
			SGIDs: source.SGIDs,
		},
	}, nil
}

func (ns *UserNS) Apply() error {
	return applyUserNS(ns)
}

func (ns *UserNS) Restore() error {
	if ns.Source == nil {
		return nil
	}

	return applyUserNS(ns.Source)
}
