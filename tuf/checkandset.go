package tuf

import (
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/docker/notary/tuf/data"
)

// SetRoot parses the Signed object into a SignedRoot object, sets
// the keys and roles in the KeyDB, and sets the Repo.Root field
// to the SignedRoot object.
func (tr *Repo) CheckAndSetRoot(s *data.Signed) error {
	root, err := data.RootFromSigned(s)
	if err != nil {
		return err
	}
	for _, key := range root.Signed.Keys {
		logrus.Debug("Adding key ", key.ID())
		tr.keysDB.AddKey(key)
	}
	for roleName, role := range root.Signed.Roles {
		logrus.Debugf("Adding role %s with keys %s", roleName, strings.Join(role.KeyIDs, ","))
		baseRole, err := data.NewRole(
			roleName,
			role.Threshold,
			role.KeyIDs,
			nil,
			nil,
		)
		if err != nil {
			return err
		}
		err = tr.keysDB.AddRole(baseRole)
		if err != nil {
			return err
		}
	}
	tr.Root = root
	return nil
}

// SetTimestamp parses the Signed object into a SignedTimestamp object
// and sets the Repo.Timestamp field.
func (tr *Repo) CheckAndSetTimestamp(s *data.Signed) error {
	ts, err := data.TimestampFromSigned(s)
	if err != nil {
		return err
	}
	tr.Timestamp = ts
	return nil
}

// SetSnapshot parses the Signed object into a SignedSnapshots object
// and sets the Repo.Snapshot field.
func (tr *Repo) CheckAndSetSnapshot(s *data.Signed) error {
	snap, err := data.SnapshotFromSigned(s)
	if err != nil {
		return err
	}
	tr.Snapshot = snap
	return nil
}

// SetTargets parses the Signed object into a SignedTargets object,
// reads the delegated roles and keys into the KeyDB, and sets the
// SignedTargets object agaist the role in the Repo.Targets map.
func (tr *Repo) CheckAndSetTargets(role string, s *data.Signed) error {
	t, err := data.TargetsFromSigned(s)
	if err != nil {
		return err
	}
	for _, k := range t.Signed.Delegations.Keys {
		tr.keysDB.AddKey(k)
	}
	for _, r := range t.Signed.Delegations.Roles {
		tr.keysDB.AddRole(r)
	}
	tr.Targets[role] = t
	return nil
}
