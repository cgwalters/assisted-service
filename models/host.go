// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Host host
//
// swagger:model host
type Host struct {

	// api vip connectivity
	APIVipConnectivity string `json:"api_vip_connectivity,omitempty" gorm:"type:text"`

	// bootstrap
	Bootstrap bool `json:"bootstrap,omitempty"`

	// The last time the host's agent communicated with the service.
	// Format: date-time
	CheckedInAt strfmt.DateTime `json:"checked_in_at,omitempty" gorm:"type:timestamp with time zone"`

	// The cluster that this host is associated with.
	// Format: uuid
	ClusterID *strfmt.UUID `json:"cluster_id,omitempty" gorm:"foreignkey:Cluster"`

	// connectivity
	Connectivity string `json:"connectivity,omitempty" gorm:"type:text"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty" gorm:"type:timestamp with time zone"`

	// The time that the host was deleted.
	// Format: date-time
	DeletedAt *strfmt.DateTime `json:"deleted_at,omitempty" gorm:"type:timestamp with time zone"`

	// discovery agent version
	DiscoveryAgentVersion string `json:"discovery_agent_version,omitempty"`

	// Additional information about disks, formatted as JSON.
	DisksInfo string `json:"disks_info,omitempty" gorm:"type:text"`

	// The domain name resolution result.
	DomainNameResolutions string `json:"domain_name_resolutions,omitempty" gorm:"type:text"`

	// free addresses
	FreeAddresses string `json:"free_addresses,omitempty" gorm:"type:text"`

	// Self link.
	// Required: true
	Href *string `json:"href"`

	// Unique identifier of the object.
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id" gorm:"primary_key"`

	// Json formatted string containing the user overrides for the host's pointer ignition
	IgnitionConfigOverrides string `json:"ignition_config_overrides,omitempty" gorm:"type:text"`

	// A string which will be used as Authorization Bearer token to fetch the ignition from ignition_endpoint_url.
	IgnitionEndpointToken *string `json:"ignition_endpoint_token,omitempty"`

	// Array of image statuses.
	ImagesStatus string `json:"images_status,omitempty" gorm:"type:text"`

	// The InfraEnv that this host is associated with.
	// Format: uuid
	InfraEnvID strfmt.UUID `json:"infra_env_id,omitempty" gorm:"primary_key;foreignkey:InfraEnvID"`

	// Contains the inventory disk id to install on.
	InstallationDiskID string `json:"installation_disk_id,omitempty"`

	// Contains the inventory disk path, This field is replaced by installation_disk_id field and used for backward compatability with the old UI.
	InstallationDiskPath string `json:"installation_disk_path,omitempty"`

	// installer args
	InstallerArgs string `json:"installer_args,omitempty"`

	// Installer version.
	InstallerVersion string `json:"installer_version,omitempty"`

	// inventory
	Inventory string `json:"inventory,omitempty" gorm:"type:text"`

	// Indicates the type of this object. Will be 'Host' if this is a complete object or 'HostLink' if it is just a link, or
	// 'AddToExistingClusterHost' for host being added to existing OCP cluster, or
	//
	// Required: true
	// Enum: [Host AddToExistingClusterHost]
	Kind *string `json:"kind"`

	// logs collected at
	// Format: datetime
	LogsCollectedAt strfmt.DateTime `json:"logs_collected_at,omitempty" gorm:"type:timestamp with time zone"`

	// The progress of log collection or empty if logs are not applicable
	LogsInfo LogsState `json:"logs_info,omitempty" gorm:"type:varchar(2048)"`

	// logs started at
	// Format: datetime
	LogsStartedAt strfmt.DateTime `json:"logs_started_at,omitempty" gorm:"type:timestamp with time zone"`

	// machine config pool name
	MachineConfigPoolName string `json:"machine_config_pool_name,omitempty"`

	// The configured NTP sources on the host.
	NtpSources string `json:"ntp_sources,omitempty" gorm:"type:text"`

	// progress
	Progress *HostProgressInfo `json:"progress,omitempty" gorm:"embedded;embedded_prefix:progress_"`

	// progress stages
	ProgressStages []HostStage `json:"progress_stages" gorm:"-"`

	// requested hostname
	RequestedHostname string `json:"requested_hostname,omitempty"`

	// role
	Role HostRole `json:"role,omitempty"`

	// Time at which the current progress stage started.
	// Format: date-time
	StageStartedAt strfmt.DateTime `json:"stage_started_at,omitempty" gorm:"type:timestamp with time zone"`

	// Time at which the current progress stage was last updated.
	// Format: date-time
	StageUpdatedAt strfmt.DateTime `json:"stage_updated_at,omitempty" gorm:"type:timestamp with time zone"`

	// status
	// Required: true
	// Enum: [discovering known disconnected insufficient disabled preparing-for-installation preparing-failed preparing-successful pending-for-input installing installing-in-progress installing-pending-user-action resetting-pending-user-action installed error resetting added-to-existing-cluster cancelled binding unbinding unbinding-pending-user-action known-unbound disconnected-unbound insufficient-unbound disabled-unbound discovering-unbound]
	Status *string `json:"status"`

	// status info
	// Required: true
	StatusInfo *string `json:"status_info" gorm:"type:varchar(2048)"`

	// The last time that the host status was updated.
	// Format: date-time
	StatusUpdatedAt strfmt.DateTime `json:"status_updated_at,omitempty" gorm:"type:timestamp with time zone"`

	// suggested role
	SuggestedRole HostRole `json:"suggested_role,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty" gorm:"type:timestamp with time zone"`

	// user name
	UserName string `json:"user_name,omitempty"`

	// JSON-formatted string containing the validation results for each validation id grouped by category (network, hardware, etc.)
	ValidationsInfo string `json:"validations_info,omitempty" gorm:"type:text"`
}

// Validate validates this host
func (m *Host) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCheckedInAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeletedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInfraEnvID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogsCollectedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogsInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogsStartedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProgress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProgressStages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStageStartedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStageUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSuggestedRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Host) validateCheckedInAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CheckedInAt) { // not required
		return nil
	}

	if err := validate.FormatOf("checked_in_at", "body", "date-time", m.CheckedInAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Host) validateClusterID(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterID) { // not required
		return nil
	}

	if err := validate.FormatOf("cluster_id", "body", "uuid", m.ClusterID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Host) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Host) validateDeletedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.DeletedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("deleted_at", "body", "date-time", m.DeletedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Host) validateHref(formats strfmt.Registry) error {

	if err := validate.Required("href", "body", m.Href); err != nil {
		return err
	}

	return nil
}

func (m *Host) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Host) validateInfraEnvID(formats strfmt.Registry) error {

	if swag.IsZero(m.InfraEnvID) { // not required
		return nil
	}

	if err := validate.FormatOf("infra_env_id", "body", "uuid", m.InfraEnvID.String(), formats); err != nil {
		return err
	}

	return nil
}

var hostTypeKindPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Host","AddToExistingClusterHost"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hostTypeKindPropEnum = append(hostTypeKindPropEnum, v)
	}
}

const (

	// HostKindHost captures enum value "Host"
	HostKindHost string = "Host"

	// HostKindAddToExistingClusterHost captures enum value "AddToExistingClusterHost"
	HostKindAddToExistingClusterHost string = "AddToExistingClusterHost"
)

// prop value enum
func (m *Host) validateKindEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, hostTypeKindPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Host) validateKind(formats strfmt.Registry) error {

	if err := validate.Required("kind", "body", m.Kind); err != nil {
		return err
	}

	// value enum
	if err := m.validateKindEnum("kind", "body", *m.Kind); err != nil {
		return err
	}

	return nil
}

func (m *Host) validateLogsCollectedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.LogsCollectedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("logs_collected_at", "body", "datetime", m.LogsCollectedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Host) validateLogsInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.LogsInfo) { // not required
		return nil
	}

	if err := m.LogsInfo.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("logs_info")
		}
		return err
	}

	return nil
}

func (m *Host) validateLogsStartedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.LogsStartedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("logs_started_at", "body", "datetime", m.LogsStartedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Host) validateProgress(formats strfmt.Registry) error {

	if swag.IsZero(m.Progress) { // not required
		return nil
	}

	if m.Progress != nil {
		if err := m.Progress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("progress")
			}
			return err
		}
	}

	return nil
}

func (m *Host) validateProgressStages(formats strfmt.Registry) error {

	if swag.IsZero(m.ProgressStages) { // not required
		return nil
	}

	for i := 0; i < len(m.ProgressStages); i++ {

		if err := m.ProgressStages[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("progress_stages" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *Host) validateRole(formats strfmt.Registry) error {

	if swag.IsZero(m.Role) { // not required
		return nil
	}

	if err := m.Role.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("role")
		}
		return err
	}

	return nil
}

func (m *Host) validateStageStartedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.StageStartedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("stage_started_at", "body", "date-time", m.StageStartedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Host) validateStageUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.StageUpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("stage_updated_at", "body", "date-time", m.StageUpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

var hostTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["discovering","known","disconnected","insufficient","disabled","preparing-for-installation","preparing-failed","preparing-successful","pending-for-input","installing","installing-in-progress","installing-pending-user-action","resetting-pending-user-action","installed","error","resetting","added-to-existing-cluster","cancelled","binding","unbinding","unbinding-pending-user-action","known-unbound","disconnected-unbound","insufficient-unbound","disabled-unbound","discovering-unbound"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hostTypeStatusPropEnum = append(hostTypeStatusPropEnum, v)
	}
}

const (

	// HostStatusDiscovering captures enum value "discovering"
	HostStatusDiscovering string = "discovering"

	// HostStatusKnown captures enum value "known"
	HostStatusKnown string = "known"

	// HostStatusDisconnected captures enum value "disconnected"
	HostStatusDisconnected string = "disconnected"

	// HostStatusInsufficient captures enum value "insufficient"
	HostStatusInsufficient string = "insufficient"

	// HostStatusDisabled captures enum value "disabled"
	HostStatusDisabled string = "disabled"

	// HostStatusPreparingForInstallation captures enum value "preparing-for-installation"
	HostStatusPreparingForInstallation string = "preparing-for-installation"

	// HostStatusPreparingFailed captures enum value "preparing-failed"
	HostStatusPreparingFailed string = "preparing-failed"

	// HostStatusPreparingSuccessful captures enum value "preparing-successful"
	HostStatusPreparingSuccessful string = "preparing-successful"

	// HostStatusPendingForInput captures enum value "pending-for-input"
	HostStatusPendingForInput string = "pending-for-input"

	// HostStatusInstalling captures enum value "installing"
	HostStatusInstalling string = "installing"

	// HostStatusInstallingInProgress captures enum value "installing-in-progress"
	HostStatusInstallingInProgress string = "installing-in-progress"

	// HostStatusInstallingPendingUserAction captures enum value "installing-pending-user-action"
	HostStatusInstallingPendingUserAction string = "installing-pending-user-action"

	// HostStatusResettingPendingUserAction captures enum value "resetting-pending-user-action"
	HostStatusResettingPendingUserAction string = "resetting-pending-user-action"

	// HostStatusInstalled captures enum value "installed"
	HostStatusInstalled string = "installed"

	// HostStatusError captures enum value "error"
	HostStatusError string = "error"

	// HostStatusResetting captures enum value "resetting"
	HostStatusResetting string = "resetting"

	// HostStatusAddedToExistingCluster captures enum value "added-to-existing-cluster"
	HostStatusAddedToExistingCluster string = "added-to-existing-cluster"

	// HostStatusCancelled captures enum value "cancelled"
	HostStatusCancelled string = "cancelled"

	// HostStatusBinding captures enum value "binding"
	HostStatusBinding string = "binding"

	// HostStatusUnbinding captures enum value "unbinding"
	HostStatusUnbinding string = "unbinding"

	// HostStatusUnbindingPendingUserAction captures enum value "unbinding-pending-user-action"
	HostStatusUnbindingPendingUserAction string = "unbinding-pending-user-action"

	// HostStatusKnownUnbound captures enum value "known-unbound"
	HostStatusKnownUnbound string = "known-unbound"

	// HostStatusDisconnectedUnbound captures enum value "disconnected-unbound"
	HostStatusDisconnectedUnbound string = "disconnected-unbound"

	// HostStatusInsufficientUnbound captures enum value "insufficient-unbound"
	HostStatusInsufficientUnbound string = "insufficient-unbound"

	// HostStatusDisabledUnbound captures enum value "disabled-unbound"
	HostStatusDisabledUnbound string = "disabled-unbound"

	// HostStatusDiscoveringUnbound captures enum value "discovering-unbound"
	HostStatusDiscoveringUnbound string = "discovering-unbound"
)

// prop value enum
func (m *Host) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, hostTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Host) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *Host) validateStatusInfo(formats strfmt.Registry) error {

	if err := validate.Required("status_info", "body", m.StatusInfo); err != nil {
		return err
	}

	return nil
}

func (m *Host) validateStatusUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.StatusUpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("status_updated_at", "body", "date-time", m.StatusUpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Host) validateSuggestedRole(formats strfmt.Registry) error {

	if swag.IsZero(m.SuggestedRole) { // not required
		return nil
	}

	if err := m.SuggestedRole.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("suggested_role")
		}
		return err
	}

	return nil
}

func (m *Host) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Host) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Host) UnmarshalBinary(b []byte) error {
	var res Host
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
