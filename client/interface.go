package client

import (
	"github.com/docker/notary/client/changelist"
	"github.com/docker/notary/tuf/data"
)

type Repository interface {
	// General management operations
	Initialize(rootKeyIDs []string, serverManagedRoles ...string) error
	Publish() error
	DeleteTrustData(deleteRemote bool)

	// Target Operations
	AddTarget(target *Target, roles ...string) error
	RemoveTarget(targetName string, roles ...string) error
	ListTargets(roles ...string) ([]*TargetWithRole, error)
	GetTargetByName(name string, roles ...string) (*TargetWithRole, error)
	GetAllTargetMetadataByName(name string) ([]TargetSignedStruct, error)

	// Changelist operations
	GetChangelist() (changelist.Changelist, error)

	// Role operations
	ListRoles() ([]RoleWithSignatures, error)
	GetDelegationRoles() ([]data.Role, error)
	AddDelegation(name string, delegationKeys []data.PublicKey, paths []string) error
	AddDelegationRoleAndKeys(name string, delegationKeys []data.PublicKey) error
	AddDelegationPaths(name string, paths []string) error
	RemoveDelegationKeysAndPaths(name string, keyIDs, paths []string) error
	RemoveDelegationRole(name string) error
	RemoveDelegationPaths(name string, paths []string) error
	RemoveDelegationKeys(name string, keyIDs []string) error
	ClearDelegationPaths(name string) error

	// Witness and other re-signing operations
	Witness(roles ...string) ([]string, error)

	// Key Operations
	RotateKey(role string, serverManagesKey bool, keyList []string) error
}
