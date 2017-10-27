package client

import (
	"github.com/theupdateframework/notary/client/changelist"
	"github.com/theupdateframework/notary/tuf/data"
	"github.com/theupdateframework/notary/tuf/signed"
)

// Repository represents the set of options that must be supported over a TUF repo.
type Repository interface {
	// General management operations

	// Initialize performs the necessary setup operations for a new, empty notary repository
	Initialize(rootKeyIDs []string, serverManagedRoles ...data.RoleName) error

	// InitializeWithCertificate erforms the necessary setup operations for a new, empty repo using the provided certificates
	// as the public components for the root keys. This permits use of external CAs to be used
	// with trust pinning.
	InitializeWithCertificate(rootKeyIDs []string, rootCerts []data.PublicKey, serverManagedRoles ...data.RoleName) error

	// Publish the repository to the relevant location. This will be implementation
	// specific and could be a local filesystem, or a remote Notary Server.
	Publish() error

	// Target Operations

	// AddTarget adds a single target to the specified roles.
	AddTarget(target *Target, roles ...data.RoleName) error

	// RemoveTarget deletes the target corresponding to targetName from the specified roles.
	RemoveTarget(targetName string, roles ...data.RoleName) error

	// ListTargets returns a list of all targets using the provided roles
	// as starting points from which to perform a priority based traversal
	// of the delegations.
	// The priority ordering is implementation specific but it is expected
	// that no role will be inspected more than once even when one role
	// in the roles is a descendent of another.
	ListTargets(roles ...data.RoleName) ([]*TargetWithRole, error)

	// GetTargetByName returns a single Target corresponding to name.
	// It performs a priority based traversal of the delegations using the provided
	// roles list as starting points.
	// The priority ordering is implementation specific but it is expected
	// that no role will be inspected more than once even when one role
	// in the roles is a descendent of another.
	GetTargetByName(name string, roles ...data.RoleName) (*TargetWithRole, error)

	// GetAllTargetMetadataByName returns all instances of a target referenced
	// by name across all delegations. The returned TargetSignedStruct includes
	// the signatures associated with the delegations the target was found in.
	GetAllTargetMetadataByName(name string) ([]TargetSignedStruct, error)

	// Changelist operations

	// GetChangelist returns the pending changelist for the current Repository.
	GetChangelist() (changelist.Changelist, error)

	// Role operations

	// ListRoles returns all roles in the Repository
	ListRoles() ([]RoleWithSignatures, error)

	// GetDelegationRoles returns all delegations in the Repository
	GetDelegationRoles() ([]data.Role, error)

	// AddDelegation creates a new delegation, or updates an existing one
	// with the provided keys and paths if the role already exists
	AddDelegation(name data.RoleName, delegationKeys []data.PublicKey, paths []string) error

	// AddDelegationRoleAndKeys updates an existing delegation with the provided keys
	AddDelegationRoleAndKeys(name data.RoleName, delegationKeys []data.PublicKey) error

	// AddDelegationPaths updates an existing delegation with the provided paths
	AddDelegationPaths(name data.RoleName, paths []string) error

	// RemoveDelegationKeysAndPaths removes the give keyIDs and paths from the
	// delegation referenced by name.
	RemoveDelegationKeysAndPaths(name data.RoleName, keyIDs, paths []string) error

	// RemoveDelegationRole deletes the delegation referenced by name
	RemoveDelegationRole(name data.RoleName) error

	// RemoveDelegationPaths removes paths from the delegation referenced by name
	RemoveDelegationPaths(name data.RoleName, paths []string) error

	// RemoveDelegationKeys removes keyIDs from the delegation referenced by name
	RemoveDelegationKeys(name data.RoleName, keyIDs []string) error

	// ClearDelegationPaths removes all paths from the delegation referenced by name
	ClearDelegationPaths(name data.RoleName) error

	// Witness and other re-signing operations

	// Witness is used to recover delegation roles that have become invalid
	// due to, for example, expiry or the keys matching the current signatures
	// having been removed as valid signing keys.
	Witness(roles ...data.RoleName) ([]data.RoleName, error)

	// Key Operations

	// RotateKey is used to change the key for one of the four primary TUF
	// roles: root, targets, snapshot, or timestamp.
	RotateKey(role data.RoleName, serverManagesKey bool, keyList []string) error

	// GetCryptoService returns the CryptoService in use by this Repository
	GetCryptoService() signed.CryptoService

	// SetLegacyVersions sets the number of old root versions that should
	// be parsed for out of date root keys before signing a new root file.
	// This is to support some older clients and will be removed at some point.
	SetLegacyVersions(int)

	// GetGUN returns the Globally Unique Name (GUN) this Repository references.
	GetGUN() data.GUN
}
