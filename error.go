package xenAPI

import (
	"fmt"
)

const (
	ERR_MESSAGE_DEPRECATED = "MESSAGE_DEPRECATED"
	ERR_MESSAGE_REMOVED    = "MESSAGE_REMOVED"

	ERR_PERMISSION_DENIED                  = "PERMISSION_DENIED"
	ERR_INTERNAL_ERROR                     = "INTERNAL_ERROR"
	ERR_MAP_DUPLICATE_KEY                  = "MAP_DUPLICATE_KEY"
	ERR_DB_UNIQUENESS_CONSTRAINT_VIOLATION = "DB_UNIQUENESS_CONSTRAINT_VIOLATION"
	ERR_LOCATION_NOT_UNIQUE                = "LOCATION_NOT_UNIQUE"
	ERR_MESSAGE_METHOD_UNKNOWN             = "MESSAGE_METHOD_UNKNOWN"
	ERR_MESSAGE_PARAMETER_COUNT_MISMATCH   = "MESSAGE_PARAMETER_COUNT_MISMATCH"
	ERR_VALUE_NOT_SUPPORTED                = "VALUE_NOT_SUPPORTED"
	ERR_INVALID_VALUE                      = "INVALID_VALUE"
	ERR_MEMORY_CONSTRAINT_VIOLATION        = "MEMORY_CONSTRAINT_VIOLATION"
	ERR_FIELD_TYPE_ERROR                   = "FIELD_TYPE_ERROR"
	ERR_SESSION_AUTHENTICATION_FAILED      = "SESSION_AUTHENTICATION_FAILED"
	ERR_SESSION_INVALID                    = "SESSION_INVALID"
	ERR_CHANGE_PASSWORD_REJECTED           = "CHANGE_PASSWORD_REJECTED"
	ERR_USER_IS_NOT_LOCAL_SUPERUSER        = "USER_IS_NOT_LOCAL_SUPERUSER"
	ERR_CANNOT_CONTACT_HOST                = "CANNOT_CONTACT_HOST"
	ERR_NOT_SUPPORTED_DURING_UPGRADE       = "NOT_SUPPORTED_DURING_UPGRADE"

	ERR_HANDLE_INVALID = "HANDLE_INVALID"
	ERR_UUID_INVALID   = "UUID_INVALID"

	ERR_VM_HVM_REQUIRED  = "VM_HVM_REQUIRED"
	ERR_VM_NO_VCPUS      = "VM_NO_VCPUS"
	ERR_VM_TOOMANY_VCPUS = "VM_TOO_MANY_VCPUS"
	ERR_VM_IS_PROTECTED  = "VM_IS_PROTECTED"

	ERR_HOST_IN_USE                     = "HOST_IN_USE"
	ERR_HOST_IN_EMERGENCY_MODE          = "HOST_IN_EMERGENCY_MODE"
	ERR_HOST_CANNOT_READ_METRICS        = "HOST_CANNOT_READ_METRICS"
	ERR_HOST_DISABLED                   = "HOST_DISABLED"
	ERR_HOST_DISABLED_UNTIL_REBOOT      = "HOST_DISABLED_UNTIL_REBOOT"
	ERR_HOST_NOT_DISABLED               = "HOST_NOT_DISABLED"
	ERR_HOST_NOT_LIVE                   = "HOST_NOT_LIVE"
	ERR_HOST_IS_LIVE                    = "HOST_IS_LIVE"
	ERR_HOST_POWER_ON_MODE_DISABLED     = "HOST_POWER_ON_MODE_DISABLED"
	ERR_HOST_NOT_ENOUGH_FREE_MEMORY     = "HOST_NOT_ENOUGH_FREE_MEMORY"
	ERR_NO_HOSTS_AVAILABLE              = "NO_HOSTS_AVAILABLE"
	ERR_HOST_OFFLINE                    = "HOST_OFFLINE"
	ERR_HOST_CANNOT_DESTROY_SELF        = "HOST_CANNOT_DESTROY_SELF"
	ERR_HOST_IS_SLAVE                   = "HOST_IS_SLAVE"
	ERR_HOST_NAME_INVALID               = "HOST_NAME_INVALID"
	ERR_HOST_HAS_RESIDENT_VMS           = "HOST_HAS_RESIDENT_VMS"
	ERR_HOSTS_FAILED_TO_ENABLE_CACHING  = "HOSTS_FAILED_TO_ENABLE_CACHING"
	ERR_HOSTS_FAILED_TO_DISABLE_CACHING = "HOSTS_FAILED_TO_DISABLE_CACHING"
	ERR_HOST_CANNOT_SEE_SR              = "HOST_CANNOT_SEE_SR"

	// (* HOST ERRORS WHICH EXPLAIN WHY THE HOST IS IN EMERGENCY MODE *)
	ERR_HOST_ITS_OWN_SLAVE = "HOST_ITS_OWN_SLAVE"
	ERR_HOST_STILL_BOOTING = "HOST_STILL_BOOTING"

	// (* LICENSE *)
	ERR_HOST_HAS_NO_MANAGEMENT_IP    = "HOST_HAS_NO_MANAGEMENT_IP"
	ERR_HOST_MASTER_CANNOT_TALK_BACK = "HOST_MASTER_CANNOT_TALK_BACK"
	ERR_HOST_UNKNOWN_TO_MASTER       = "HOST_UNKNOWN_TO_MASTER"

	// (* SHOULD BE FENCED *)
	ERR_HOST_BROKEN                         = "HOST_BROKEN"
	ERR_INTERFACE_HAS_NO_IP                 = "INTERFACE_HAS_NO_IP"
	ERR_INVALID_IP_ADDRESS_SPECIFIED        = "INVALID_IP_ADDRESS_SPECIFIED"
	ERR_PIF_HAS_NO_NETWORK_CONFIGURATION    = "PIF_HAS_NO_NETWORK_CONFIGURATION"
	ERR_PIF_HAS_NO_V6_NETWORK_CONFIGURATION = "PIF_HAS_NO_V6_NETWORK_CONFIGURATION"

	ERR_DEVICE_ATTACH_TIMEOUT  = "DEVICE_ATTACH_TIMEOUT"
	ERR_DEVICE_DETACH_TIMEOUT  = "DEVICE_DETACH_TIMEOUT"
	ERR_DEVICE_DETACH_REJECTED = "DEVICE_DETACH_REJECTED"

	ERR_OPERATION_NOT_ALLOWED                                                       = "OPERATION_NOT_ALLOWED"
	ERR_OPERATION_BLOCKED                                                           = "OPERATION_BLOCKED"
	ERR_NETWORK_ALREADY_CONNECTED                                                   = "NETWORK_ALREADY_CONNECTED"
	ERR_CANNOT_DESTROY_SYSTEM_NETWORK                                               = "CANNOT_DESTROY_SYSTEM_NETWORK"
	ERR_PIF_IS_PHYSICAL                                                             = "PIF_IS_PHYSICAL"
	ERR_PIF_IS_VLAN                                                                 = "PIF_IS_VLAN"
	ERR_PIF_VLAN_EXISTS                                                             = "PIF_VLAN_EXISTS"
	ERR_PIF_VLAN_STILL_EXISTS                                                       = "PIF_VLAN_STILL_EXISTS"
	ERR_PIF_DEVICE_NOT_FOUND                                                        = "PIF_DEVICE_NOT_FOUND"
	ERR_PIF_ALREADY_BONDED                                                          = "PIF_ALREADY_BONDED"
	ERR_PIF_CANNOT_BOND_CROSS_HOST                                                  = "PIF_CANNOT_BOND_CROSS_HOST"
	ERR_PIF_BOND_NEEDS_MORE_MEMBERS                                                 = "PIF_BOND_NEEDS_MORE_MEMBERS"
	ERR_PIF_CONFIGURATION_ERROR                                                     = "PIF_CONFIGURATION_ERROR"
	ERR_PIF_IS_MANAGEMENT_IFACE                                                     = "PIF_IS_MANAGEMENT_INTERFACE"
	ERR_PIF_INCOMPATIBLE_PRIMARY_ADDRESS_TYPE                                       = "PIF_INCOMPATIBLE_PRIMARY_ADDRESS_TYPE"
	ERR_REQUIRED_PIF_IS_UNPLUGGED                                                   = "REQUIRED_PIF_IS_UNPLUGGED"
	ERR_PIF_DOES_NOT_ALLOW_UNPLUG                                                   = "PIF_DOES_NOT_ALLOW_UNPLUG"
	ERR_PIF_UNMANAGED                                                               = "PIF_UNMANAGED"
	ERR_CANNOT_PLUG_BOND_SLAVE                                                      = "CANNOT_PLUG_BOND_SLAVE"
	ERR_CANNOT_ADD_VLAN_TO_BOND_SLAVE                                               = "CANNOT_ADD_VLAN_TO_BOND_SLAVE"
	ERR_CANNOT_ADD_TUNNEL_TO_BOND_SLAVE                                             = "CANNOT_ADD_TUNNEL_TO_BOND_SLAVE"
	ERR_CANNOT_CHANGE_PIF_PROPERTIES                                                = "CANNOT_CHANGE_PIF_PROPERTIES"
	ERR_INCOMPATIBLE_PIF_PROPERTIES                                                 = "INCOMPATIBLE_PIF_PROPERTIES"
	ERR_SLAVE_REQUIRES_MANAGEMENT_IFACE                                             = "SLAVE_REQUIRES_MANAGEMENT_INTERFACE"
	ERR_VIF_IN_USE                                                                  = "VIF_IN_USE"
	ERR_CANNOT_PLUG_VIF                                                             = "CANNOT_PLUG_VIF"
	ERR_MAC_STILL_EXISTS                                                            = "MAC_STILL_EXISTS"
	ERR_MAC_DOES_NOT_EXIST                                                          = "MAC_DOES_NOT_EXIST"
	ERR_MAC_INVALID                                                                 = "MAC_INVALID"
	ERR_DUPLICATE_PIF_DEVICE_NAME                                                   = "DUPLICATE_PIF_DEVICE_NAME"
	ERR_COULD_NOT_FIND_NETWORK_INTERFACE_WITH_SPECIFIED_DEVICE_NAME_AND_MAC_ADDRESS = "COULD_NOT_FIND_NETWORK_INTERFACE_WITH_SPECIFIED_DEVICE_NAME_AND_MAC_ADDRESS"

	ERR_OPENVSWITCH_NOT_ACTIVE                                                   = "OPENVSWITCH_NOT_ACTIVE"
	ERR_TRANSPORT_PIF_NOT_CONFIGURED                                             = "TRANSPORT_PIF_NOT_CONFIGURED"
	ERR_IS_TUNNEL_ACCESS_PIF                                                     = "IS_TUNNEL_ACCESS_PIF"
	ERR_PIF_TUNNEL_STILL_EXISTS                                                  = "PIF_TUNNEL_STILL_EXISTS"
	ERR_BRIDGE_NOT_AVAILABLE                                                     = "BRIDGE_NOT_AVAILABLE"
	ERR_VLAN_TAG_INVALID                                                         = "VLAN_TAG_INVALID"
	ERR_VM_BAD_POWER_STATE                                                       = "VM_BAD_POWER_STATE"
	ERR_VM_IS_TEMPLATE                                                           = "VM_IS_TEMPLATE"
	ERR_VM_IS_SNAPSHOT                                                           = "VM_IS_SNAPSHOT"
	ERR_OTHER_OPERATION_IN_PROGRESS                                              = "OTHER_OPERATION_IN_PROGRESS"
	ERR_VBD_NOT_REMOVABLE_MEDIA                                                  = "VBD_NOT_REMOVABLE_MEDIA"
	ERR_VBD_NOT_UNPLUGGABLE                                                      = "VBD_NOT_UNPLUGGABLE"
	ERR_VBD_NOT_EMPTY                                                            = "VBD_NOT_EMPTY"
	ERR_VBD_IS_EMPTY                                                             = "VBD_IS_EMPTY"
	ERR_VBD_TRAY_LOCKED                                                          = "VBD_TRAY_LOCKED"
	ERR_VBD_MISSING                                                              = "VBD_MISSING"
	ERR_VM_NO_EMPTY_CD_VBD                                                       = "VM_NO_EMPTY_CD_VBD"
	ERR_VM_SNAPSHOT_FAILED                                                       = "VM_SNAPSHOT_FAILED"
	ERR_VM_SNAPSHOT_WITH_QUIESCE_FAILED                                          = "VM_SNAPSHOT_WITH_QUIESCE_FAILED"
	ERR_VM_SNAPSHOT_WITH_QUIESCE_TIMEOUT                                         = "VM_SNAPSHOT_WITH_QUIESCE_TIMEOUT"
	ERR_VM_SNAPSHOT_WITH_QUIESCE_PLUGIN_DOES_NOT_RESPOND                         = "VM_SNAPSHOT_WITH_QUIESCE_PLUGIN_DEOS_NOT_RESPOND"
	ERR_VM_SNAPSHOT_WITH_QUIESCE_NOT_SUPPORTED                                   = "VM_SNAPSHOT_WITH_QUIESCE_NOT_SUPPORTED"
	ERR_XEN_VSS_REQ_ERROR_INIT_FAILED                                            = "XEN_VSS_REQ_ERROR_INIT_FAILED"
	ERR_XEN_VSS_REQ_ERROR_PROV_NOT_LOADED                                        = "XEN_VSS_REQ_ERROR_PROV_NOT_LOADED"
	ERR_XEN_VSS_REQ_ERROR_NO_VOLUMES_SUPPORTED                                   = "XEN_VSS_REQ_ERROR_NO_VOLUMES_SUPPORTED"
	ERR_XEN_VSS_REQ_ERROR_START_SNAPSHOT_SET_FAILED                              = "XEN_VSS_REQ_ERROR_START_SNAPSHOT_SET_FAILED"
	ERR_XEN_VSS_REQ_ERROR_ADDING_VOLUME_TO_SNAPSET_FAILED                        = "XEN_VSS_REQ_ERROR_ADDING_VOLUME_TO_SNAPSET_FAILED"
	ERR_XEN_VSS_REQ_ERROR_PREPARING_WRITERS                                      = "XEN_VSS_REQ_ERROR_PREPARING_WRITERS"
	ERR_XEN_VSS_REQ_ERROR_CREATING_SNAPSHOT                                      = "XEN_VSS_REQ_ERROR_CREATING_SNAPSHOT"
	ERR_XEN_VSS_REQ_ERROR_CREATING_SNAPSHOT_XML_STRING                           = "XEN_VSS_REQ_ERROR_CREATING_SNAPSHOT_XML_STRING"
	ERR_VM_REVERT_FAILED                                                         = "VM_REVERT_FAILED"
	ERR_VM_CHECKPOINT_SUSPEND_FAILED                                             = "VM_CHECKPOINT_SUSPEND_FAILED"
	ERR_VM_CHECKPOINT_RESUME_FAILED                                              = "VM_CHECKPOINT_RESUME_FAILED"
	ERR_VM_UNSAFE_BOOT                                                           = "VM_UNSAFE_BOOT"
	ERR_VM_REQUIRES_SR                                                           = "VM_REQUIRES_SR"
	ERR_VM_REQUIRES_VDI                                                          = "VM_REQUIRES_VDI"
	ERR_VM_REQUIRES_NET                                                          = "VM_REQUIRES_NETWORK"
	ERR_VM_REQUIRES_GPU                                                          = "VM_REQUIRES_GPU"
	ERR_VM_REQUIRES_VGPU                                                         = "VM_REQUIRES_VGPU"
	ERR_VM_REQUIRES_IOMMU                                                        = "VM_REQUIRES_IOMMU"
	ERR_VM_HOST_INCOMPATIBLE_VERSION                                             = "VM_HOST_INCOMPATIBLE_VERSION"
	ERR_VM_HOST_INCOMPATIBLE_VIRTUAL_HARDWARE_PLATFORM_VERSION                   = "VM_HOST_INCOMPATIBLE_VIRTUAL_HARDWARE_PLATFORM_VERSION"
	ERR_VM_HAS_PCI_ATTACHED                                                      = "VM_HAS_PCI_ATTACHED"
	ERR_VM_HAS_VGPU                                                              = "VM_HAS_VGPU"
	ERR_HOST_CANNOT_ATTACH_NETWORK                                               = "HOST_CANNOT_ATTACH_NETWORK"
	ERR_VM_NO_SUSPEND_SR                                                         = "VM_NO_SUSPEND_SR"
	ERR_VM_NO_CRASHDUMP_SR                                                       = "VM_NO_CRASHDUMP_SR"
	ERR_VM_MIGRATE_FAILED                                                        = "VM_MIGRATE_FAILED"
	ERR_VM_MISSING_PV_DRIVERS                                                    = "VM_MISSING_PV_DRIVERS"
	ERR_VM_FAILED_SHUTDOWN_ACK                                                   = "VM_FAILED_SHUTDOWN_ACKNOWLEDGMENT"
	ERR_VM_OLD_PV_DRIVERS                                                        = "VM_OLD_PV_DRIVERS"
	ERR_VM_LACKS_FEATURE_SHUTDOWN                                                = "VM_LACKS_FEATURE_SHUTDOWN"
	ERR_VM_LACKS_FEATURE_SUSPEND                                                 = "VM_LACKS_FEATURE_SUSPEND"
	ERR_VM_LACKS_FEATURE_VCPU_HOTPLUG                                            = "VM_LACKS_FEATURE_VCPU_HOTPLUG"
	ERR_VM_CANNOT_DELETE_DEFAULT_TEMPLATE                                        = "VM_CANNOT_DELETE_DEFAULT_TEMPLATE"
	ERR_VM_MEMORY_SIZE_TOO_LOW                                                   = "VM_MEMORY_SIZE_TOO_LOW"
	ERR_VM_MEMORY_TARGET_WAIT_TIMEOUT                                            = "VM_MEMORY_TARGET_WAIT_TIMEOUT"
	ERR_VM_SHUTDOWN_TIMEOUT                                                      = "VM_SHUTDOWN_TIMEOUT"
	ERR_VM_DUPLICATE_VBD_DEVICE                                                  = "VM_DUPLICATE_VBD_DEVICE"
	ERR_ILLEGAL_VBD_DEVICE                                                       = "ILLEGAL_VBD_DEVICE"
	ERR_VM_NOT_RESIDENT_HERE                                                     = "VM_NOT_RESIDENT_HERE"
	ERR_VM_CRASHED                                                               = "VM_CRASHED"
	ERR_VM_REBOOTED                                                              = "VM_REBOOTED"
	ERR_VM_HALTED                                                                = "VM_HALTED"
	ERR_VM_ATTACHED_TO_MORE_THAN_ONE_VDI_WITH_TIMEOFFSET_MARKED_AS_RESET_ON_BOOT = "VM_ATTACHED_TO_MORE_THAN_ONE_VDI_WITH_TIMEOFFSET_MARKED_AS_RESET_ON_BOOT"
	ERR_VMS_FAILED_TO_COOPERATE                                                  = "VMS_FAILED_TO_COOPERATE"
	ERR_VM_PV_DRIVERS_IN_USE                                                     = "VM_PV_DRIVERS_IN_USE"
	ERR_DOMAIN_EXISTS                                                            = "DOMAIN_EXISTS"
	ERR_CANNOT_RESET_CONTROL_DOMAIN                                              = "CANNOT_RESET_CONTROL_DOMAIN"
	ERR_NOT_SYSTEM_DOMAIN                                                        = "NOT_SYSTEM_DOMAIN"
	ERR_ONLY_PROVISION_TEMPLATE                                                  = "PROVISION_ONLY_ALLOWED_ON_TEMPLATE"
	ERR_ONLY_REVERT_SNAPSHOT                                                     = "REVERT_ONLY_ALLOWED_ON_SNAPSHOT"
	ERR_PROVISION_FAILED_OUT_OF_SPACE                                            = "PROVISION_FAILED_OUT_OF_SPACE"
	ERR_BOOTLOADER_FAILED                                                        = "BOOTLOADER_FAILED"
	ERR_UNKNOWN_BOOTLOADER                                                       = "UNKNOWN_BOOTLOADER"
	ERR_OBJECT_NOLONGER_EXISTS                                                   = "OBJECT_NOLONGER_EXISTS"
	ERR_SR_ATTACH_FAILED                                                         = "SR_ATTACH_FAILED"
	ERR_SR_FULL                                                                  = "SR_FULL"
	ERR_SR_HAS_PBD                                                               = "SR_HAS_PBD"
	ERR_SR_REQUIRES_UPGRADE                                                      = "SR_REQUIRES_UPGRADE"
	ERR_SR_IS_CACHE_SR                                                           = "SR_IS_CACHE_SR"
	ERR_VDI_IN_USE                                                               = "VDI_IN_USE"
	ERR_VDI_IS_SHARABLE                                                          = "VDI_IS_SHARABLE"
	ERR_VDI_READONLY                                                             = "VDI_READONLY"
	ERR_VDI_TOO_SMALL                                                            = "VDI_TOO_SMALL"
	ERR_VDI_NOT_SPARSE                                                           = "VDI_NOT_SPARSE"
	ERR_VDI_IS_A_PHYSICAL_DEVICE                                                 = "VDI_IS_A_PHYSICAL_DEVICE"
	ERR_VDI_IS_NOT_ISO                                                           = "VDI_IS_NOT_ISO"
	ERR_VBD_CDS_MUST_BE_READONLY                                                 = "VBD_CDS_MUST_BE_READONLY"

	// (* CA-83260 *)
	ERR_DISK_VBD_MUST_BE_READWRITE_FOR_HVM           = "DISK_VBD_MUST_BE_READWRITE_FOR_HVM"
	ERR_HOST_CD_DRIVE_EMPTY                          = "HOST_CD_DRIVE_EMPTY"
	ERR_VDI_NOT_AVAILABLE                            = "VDI_NOT_AVAILABLE"
	ERR_VDI_HAS_RRDS                                 = "VDI_HAS_RRDS"
	ERR_VDI_LOCATION_MISSING                         = "VDI_LOCATION_MISSING"
	ERR_VDI_CONTENT_ID_MISSING                       = "VDI_CONTENT_ID_MISSING"
	ERR_VDI_MISSING                                  = "VDI_MISSING"
	ERR_VDI_INCOMPATIBLE_TYPE                        = "VDI_INCOMPATIBLE_TYPE"
	ERR_VDI_NOT_MANAGED                              = "VDI_NOT_MANAGED"
	ERR_VDI_IO_ERROR                                 = "VDI_IO_ERROR"
	ERR_VDI_ON_BOOT_MODE_INCOMPATIBLE_WITH_OPERATION = "VDI_ON_BOOT_MODE_INCOMPATIBLE_WITH_OPERATION"
	ERR_VDI_NOT_IN_MAP                               = "VDI_NOT_IN_MAP"
	ERR_CANNOT_CREATE_STATE_FILE                     = "CANNOT_CREATE_STATE_FILE"

	ERR_OPERATION_PARTIALLY_FAILED = "OPERATION_PARTIALLY_FAILED"

	ERR_SR_UUID_EXISTS             = "SR_UUID_EXISTS"
	ERR_SR_NO_PBDS                 = "SR_HAS_NO_PBDS"
	ERR_SR_HAS_MULTIPLE_PBDS       = "SR_HAS_MULTIPLE_PBDS"
	ERR_SR_BACKEND_FAILURE         = "SR_BACKEND_FAILURE"
	ERR_SR_UNKNOWN_DRIVER          = "SR_UNKNOWN_DRIVER"
	ERR_SR_VDI_LOCKING_FAILED      = "SR_VDI_LOCKING_FAILED"
	ERR_SR_NOT_EMPTY               = "SR_NOT_EMPTY"
	ERR_SR_DEVICE_IN_USE           = "SR_DEVICE_IN_USE"
	ERR_SR_OPERATION_NOT_SUPPORTED = "SR_OPERATION_NOT_SUPPORTED"
	ERR_SR_NOT_SHARABLE            = "SR_NOT_SHARABLE"
	ERR_SR_INDESTRUCTIBLE          = "SR_INDESTRUCTIBLE"
	ERR_SR_DETACHED_ON_MASTER      = "SR_DETACHED_ON_MASTER"
	ERR_SR_ATTACHED_ON_SLAVE       = "SR_ATTACHED_ON_SLAVE"
	ERR_CLUSTERED_SR_DEGRADED      = "CLUSTERED_SR_DEGRADED"

	ERR_SM_PLUGIN_COMMUNICATION_FAILURE = "SM_PLUGIN_COMMUNICATION_FAILURE"

	ERR_PBD_EXISTS = "PBD_EXISTS"

	ERR_NOT_IMPLEMENTED                  = "NOT_IMPLEMENTED"
	ERR_DEVICE_ALREADY_ATTACHED          = "DEVICE_ALREADY_ATTACHED"
	ERR_DEVICE_ALREADY_DETACHED          = "DEVICE_ALREADY_DETACHED"
	ERR_DEVICE_ALREADY_EXISTS            = "DEVICE_ALREADY_EXISTS"
	ERR_DEVICE_NOT_ATTACHED              = "DEVICE_NOT_ATTACHED"
	ERR_NETWORK_CONTAINS_PIF             = "NETWORK_CONTAINS_PIF"
	ERR_NETWORK_CONTAINS_VIF             = "NETWORK_CONTAINS_VIF"
	ERR_GPU_GROUP_CONTAINS_VGPU          = "GPU_GROUP_CONTAINS_VGPU"
	ERR_GPU_GROUP_CONTAINS_PGPU          = "GPU_GROUP_CONTAINS_PGPU"
	ERR_GPU_GROUP_CONTAINS_NO_PGPUS      = "GPU_GROUP_CONTAINS_NO_PGPUS"
	ERR_INVALID_DEVICE                   = "INVALID_DEVICE"
	ERR_EVENTS_LOST                      = "EVENTS_LOST"
	ERR_EVENT_SUBSCRIPTION_PARSE_FAILURE = "EVENT_SUBSCRIPTION_PARSE_FAILURE"
	ERR_EVENT_FROM_TOKEN_PARSE_FAILURE   = "EVENT_FROM_TOKEN_PARSE_FAILURE"
	ERR_SESSION_NOT_REGISTERED           = "SESSION_NOT_REGISTERED"

	ERR_PGPU_IN_USE_BY_VM                          = "PGPU_IN_USE_BY_VM"
	ERR_PGPU_NOT_COMPATIBLE_WITH_GPU_GROUP         = "PGPU_NOT_COMPATIBLE_WITH_GPU_GROUP"
	ERR_PGPU_INSUFFICIENT_CAPACITY_FOR_VGPU        = "PGPU_INSUFFICIENT_CAPACITY_FOR_VGPU"
	ERR_VGPU_TYPE_NOT_ENABLED                      = "VGPU_TYPE_NOT_ENABLED"
	ERR_VGPU_TYPE_NOT_SUPPORTED                    = "VGPU_TYPE_NOT_SUPPORTED"
	ERR_VGPU_TYPE_NOT_COMPATIBLE_WITH_RUNNING_TYPE = "VGPU_TYPE_NOT_COMPATIBLE_WITH_RUNNING_TYPE"

	ERR_IMPORT_ERROR_GENERIC                  = "IMPORT_ERROR"
	ERR_IMPORT_ERROR_PREMATURE_EOF            = "IMPORT_ERROR_PREMATURE_EOF"
	ERR_IMPORT_ERROR_SOME_CHECKSUMS_FAILED    = "IMPORT_ERROR_SOME_CHECKSUMS_FAILED"
	ERR_IMPORT_ERROR_CANNOT_HANDLE_CHUNKED    = "IMPORT_ERROR_CANNOT_HANDLE_CHUNKED"
	ERR_IMPORT_ERROR_FAILED_TO_FIND_OBJECT    = "IMPORT_ERROR_FAILED_TO_FIND_OBJECT"
	ERR_IMPORT_ERROR_ATTACHED_DISKS_NOT_FOUND = "IMPORT_ERROR_ATTACHED_DISKS_NOT_FOUND"
	ERR_IMPORT_ERROR_UNEXPECTED_FILE          = "IMPORT_ERROR_UNEXPECTED_FILE"
	ERR_IMPORT_INCOMPATIBLE_VERSION           = "IMPORT_INCOMPATIBLE_VERSION"

	ERR_RESTORE_INCOMPATIBLE_VERSION         = "RESTORE_INCOMPATIBLE_VERSION"
	ERR_RESTORE_TARGET_MISSING_DEVICE        = "RESTORE_TARGET_MISSING_DEVICE"
	ERR_RESTORE_TARGET_MGMT_IF_NOT_IN_BACKUP = "RESTORE_TARGET_MGMT_IF_NOT_IN_BACKUP"

	ERR_POOL_NOT_IN_EMERGENCY_MODE                                = "NOT_IN_EMERGENCY_MODE"
	ERR_POOL_HOSTS_NOT_COMPATIBLE                                 = "HOSTS_NOT_COMPATIBLE"
	ERR_POOL_HOSTS_NOT_HOMOGENEOUS                                = "HOSTS_NOT_HOMOGENEOUS"
	ERR_POOL_JOINING_HOST_CANNOT_CONTAIN_SHARED_SRS               = "JOINING_HOST_CANNOT_CONTAIN_SHARED_SRS"
	ERR_POOL_JOINING_HOST_CANNOT_CONTAIN_NETWORK_BOND             = "JOINING_HOST_CANNOT_CONTAIN_NETWORK_BOND"
	ERR_POOL_JOINING_HOST_CANNOT_HAVE_RUNNING_OR_SUSPENDED_VMS    = "JOINING_HOST_CANNOT_HAVE_RUNNING_OR_SUSPENDED_VMS"
	ERR_POOL_JOINING_HOST_CANNOT_HAVE_RUNNING_VMS                 = "JOINING_HOST_CANNOT_HAVE_RUNNING_VMS"
	ERR_POOL_JOINING_HOST_CANNOT_HAVE_VMS_WITH_CURRENT_OPERATIONS = "JOINING_HOST_CANNOT_HAVE_VMS_WITH_CURRENT_OPERATIONS"
	ERR_POOL_JOINING_HOST_CANNOT_BE_MASTER_OF_OTHER_HOSTS         = "JOINING_HOST_CANNOT_BE_MASTER_OF_OTHER_HOSTS"
	ERR_POOL_JOINING_HOST_CONNECTION_FAILED                       = "JOINING_HOST_CONNECTION_FAILED"
	ERR_POOL_JOINING_HOST_SERVICE_FAILED                          = "JOINING_HOST_SERVICE_FAILED"
	ERR_POOL_JOINING_HOST_MUST_HAVE_PHYSICAL_MANAGEMENT_NIC       = "POOL_JOINING_HOST_MUST_HAVE_PHYSICAL_MANAGEMENT_NIC"
	ERR_POOL_JOINING_EXTERNAL_AUTH_MISMATCH                       = "POOL_JOINING_EXTERNAL_AUTH_MISMATCH"
	ERR_POOL_JOINING_HOST_MUST_HAVE_SAME_PRODUCT_VERSION          = "POOL_JOINING_HOST_MUST_HAVE_SAME_PRODUCT_VERSION"

	// (*WORKLOAD BALANCING*)
	ERR_WLB_NOT_INITIALIZED                 = "WLB_NOT_INITIALIZED"
	ERR_WLB_DISABLED                        = "WLB_DISABLED"
	ERR_WLB_CONNECTION_REFUSED              = "WLB_CONNECTION_REFUSED"
	ERR_WLB_UNKNOWN_HOST                    = "WLB_UNKNOWN_HOST"
	ERR_WLB_TIMEOUT                         = "WLB_TIMEOUT"
	ERR_WLB_AUTHENTICATION_FAILED           = "WLB_AUTHENTICATION_FAILED"
	ERR_WLB_MALFORMED_REQUEST               = "WLB_MALFORMED_REQUEST"
	ERR_WLB_MALFORMED_RESPONSE              = "WLB_MALFORMED_RESPONSE"
	ERR_WLB_XENSERVER_CONNECTION_REFUSED    = "WLB_XENSERVER_CONNECTION_REFUSED"
	ERR_WLB_XENSERVER_UNKNOWN_HOST          = "WLB_XENSERVER_UNKNOWN_HOST"
	ERR_WLB_XENSERVER_TIMEOUT               = "WLB_XENSERVER_TIMEOUT"
	ERR_WLB_XENSERVER_AUTHENTICATION_FAILED = "WLB_XENSERVER_AUTHENTICATION_FAILED"
	ERR_WLB_XENSERVER_MALFORMED_RESPONSE    = "WLB_XENSERVER_MALFORMED_RESPONSE"
	ERR_WLB_INTERNAL_ERROR                  = "WLB_INTERNAL_ERROR"
	ERR_WLB_URL_INVALID                     = "WLB_URL_INVALID"
	ERR_WLB_CONNECTION_RESET                = "WLB_CONNECTION_RESET"

	ERR_SR_NOT_SHARED = "SR_NOT_SHARED"

	ERR_DEFAULT_SR_NOT_FOUND = "DEFAULT_SR_NOT_FOUND"

	ERR_TASK_CANCELLED = "TASK_CANCELLED"

	ERR_TOO_MANY_PENDING_TASKS = "TOO_MANY_PENDING_TASKS"
	ERR_TOO_BUSY               = "TOO_BUSY"

	ERR_OUT_OF_SPACE                               = "OUT_OF_SPACE"
	ERR_INVALID_PATCH                              = "INVALID_PATCH"
	ERR_INVALID_PATCH_WITH_LOG                     = "INVALID_PATCH_WITH_LOG"
	ERR_PATCH_ALREADY_EXISTS                       = "PATCH_ALREADY_EXISTS"
	ERR_PATCH_IS_APPLIED                           = "PATCH_IS_APPLIED"
	ERR_CANNOT_FIND_PATCH                          = "CANNOT_FIND_PATCH"
	ERR_CANNOT_FETCH_PATCH                         = "CANNOT_FETCH_PATCH"
	ERR_PATCH_ALREADY_APPLIED                      = "PATCH_ALREADY_APPLIED"
	ERR_PATCH_PRECHECK_FAILED_UNKNOWN_ERROR        = "PATCH_PRECHECK_FAILED_UNKNOWN_ERROR"
	ERR_PATCH_PRECHECK_FAILED_PREREQUISITE_MISSING = "PATCH_PRECHECK_FAILED_PREREQUISITE_MISSING"
	ERR_PATCH_PRECHECK_FAILED_WRONG_SERVER_VERSION = "PATCH_PRECHECK_FAILED_WRONG_SERVER_VERSION"
	ERR_PATCH_PRECHECK_FAILED_WRONG_SERVER_BUILD   = "PATCH_PRECHECK_FAILED_WRONG_SERVER_BUILD"
	ERR_PATCH_PRECHECK_FAILED_VM_RUNNING           = "PATCH_PRECHECK_FAILED_VM_RUNNING"
	ERR_PATCH_PRECHECK_FAILED_OUT_OF_SPACE         = "PATCH_PRECHECK_FAILED_OUT_OF_SPACE"
	ERR_PATCH_PRECHECK_TOOLS_ISO_MOUNTED           = "PATCH_PRECHECK_FAILED_ISO_MOUNTED"
	ERR_PATCH_APPLY_FAILED                         = "PATCH_APPLY_FAILED"
	ERR_PATCH_APPLY_FAILED_BACKUP_FILES_EXIST      = "PATCH_APPLY_FAILED_BACKUP_FILES_EXIST"
	ERR_CANNOT_FIND_OEM_BACKUP_PARTITION           = "CANNOT_FIND_OEM_BACKUP_PARTITION"
	ERR_ONLY_ALLOWED_ON_OEM_EDITION                = "ONLY_ALLOWED_ON_OEM_EDITION"
	ERR_NOT_ALLOWED_ON_OEM_EDITION                 = "NOT_ALLOWED_ON_OEM_EDITION"
	ERR_CANNOT_FIND_STATE_PARTITION                = "CANNOT_FIND_STATE_PARTITION"
	ERR_BACKUP_SCRIPT_FAILED                       = "BACKUP_SCRIPT_FAILED"
	ERR_RESTORE_SCRIPT_FAILED                      = "RESTORE_SCRIPT_FAILED"

	ERR_LICENSE_EXPIRED                  = "LICENSE_EXPIRED"
	ERR_LICENSE_RESTRICTION              = "LICENCE_RESTRICTION"
	ERR_LICENSE_DOES_NOT_SUPPORT_POOLING = "LICENSE_DOES_NOT_SUPPORT_POOLING"
	ERR_LICENSE_HOST_POOL_MISMATCH       = "LICENSE_HOST_POOL_MISMATCH"
	ERR_LICENSE_PROCESSING_ERROR         = "LICENSE_PROCESSING_ERROR"
	ERR_LICENSE_CANNOT_DOWNGRADE_IN_POOL = "LICENSE_CANNOT_DOWNGRADE_WHILE_IN_POOL"
	ERR_LICENSE_DOES_NOT_SUPPORT_XHA     = "LICENSE_DOES_NOT_SUPPORT_XHA"

	ERR_V6D_FAILURE                = "V6D_FAILURE"
	ERR_INVALID_EDITION            = "INVALID_EDITION"
	ERR_MISSING_CONNECTION_DETAILS = "MISSING_CONNECTION_DETAILS"
	ERR_LICENSE_CHECKOUT_ERROR     = "LICENSE_CHECKOUT_ERROR"
	ERR_LICENSE_FILE_DEPRECATED    = "LICENSE_FILE_DEPRECATED"
	ERR_ACTIVATION_WHILE_NOT_FREE  = "ACTIVATION_WHILE_NOT_FREE"

	ERR_FEATURE_RESTRICTED = "FEATURE_RESTRICTED"

	ERR_XMLRPC_UNMARSHAL_FAILURE = "XMLRPC_UNMARSHAL_FAILURE"

	ERR_DUPLICATE_VM       = "DUPLICATE_VM"
	ERR_DUPLICATE_MAC_SEED = "DUPLICATE_MAC_SEED"

	ERR_CLIENT_ERROR = "CLIENT_ERROR"

	ERR_BALLOONING_DISABLED = "BALLOONING_DISABLED"

	ERR_HA_HOST_IS_ARMED         = "HA_HOST_IS_ARMED"
	ERR_HA_IS_ENABLED            = "HA_IS_ENABLED"
	ERR_HA_NOT_ENABLED           = "HA_NOT_ENABLED"
	ERR_HA_ENABLE_IN_PROGRESS    = "HA_ENABLE_IN_PROGRESS"
	ERR_HA_DISABLE_IN_PROGRESS   = "HA_DISABLE_IN_PROGRESS"
	ERR_HA_NOT_INSTALLED         = "HA_NOT_INSTALLED"
	ERR_HA_HOST_CANNOT_SEE_PEERS = "HA_HOST_CANNOT_SEE_PEERS"
	ERR_HA_TOO_FEW_HOSTS         = "HA_TOO_FEW_HOSTS"
	ERR_HA_SHOULD_BE_FENCED      = "HA_SHOULD_BE_FENCED"

	ERR_HA_ABORT_NEW_MASTER = "HA_ABORT_NEW_MASTER"

	ERR_HA_NO_PLAN                              = "HA_NO_PLAN"
	ERR_HA_LOST_STATEFILE                       = "HA_LOST_STATEFILE"
	ERR_HA_POOL_IS_ENABLED_BUT_HOST_IS_DISABLED = "HA_POOL_IS_ENABLED_BUT_HOST_IS_DISABLED"

	ERR_HA_HEARTBEAT_DAEMON_STARTUP_FAILED = "HA_HEARTBEAT_DAEMON_STARTUP_FAILED"
	ERR_HA_HOST_CANNOT_ACCESS_STATEFILE    = "HA_HOST_CANNOT_ACCESS_STATEFILE"

	ERR_HA_FAILED_TO_FORM_LIVESET = "HA_FAILED_TO_FORM_LIVESET"

	ERR_HA_CANNOT_CHANGE_BOND_STATUS_OF_MGMT_IFACE = "HA_CANNOT_CHANGE_BOND_STATUS_OF_MGMT_IFACE"

	// (* CA-16480: PREVENT CONFIGURATION ERRORS WHICH NULLIFY XHA GOODNESS *)
	ERR_HA_CONSTRAINT_VIOLATION_SR_NOT_SHARED      = "HA_CONSTRAINT_VIOLATION_SR_NOT_SHARED"
	ERR_HA_CONSTRAINT_VIOLATION_NETWORK_NOT_SHARED = "HA_CONSTRAINT_VIOLATION_NETWORK_NOT_SHARED"

	ERR_HA_OPERATION_WOULD_BREAK_FAILOVER_PLAN = "HA_OPERATION_WOULD_BREAK_FAILOVER_PLAN"

	ERR_INCOMPATIBLE_STATEFILE_SR         = "INCOMPATIBLE_STATEFILE_SR"
	ERR_INCOMPATIBLE_CLUSTER_STACK_ACTIVE = "INCOMPATIBLE_CLUSTER_STACK_ACTIVE"

	ERR_CANNOT_EVACUATE_HOST      = "CANNOT_EVACUATE_HOST"
	ERR_HOST_EVACUATE_IN_PROGRESS = "HOST_EVACUATE_IN_PROGRESS"

	ERR_SYSTEM_STATUS_RETRIEVAL_FAILED    = "SYSTEM_STATUS_RETRIEVAL_FAILED"
	ERR_SYSTEM_STATUS_MUST_USE_TAR_ON_OEM = "SYSTEM_STATUS_MUST_USE_TAR_ON_OEM"

	ERR_XAPI_HOOK_FAILED = "XAPI_HOOK_FAILED"

	ERR_NO_LOCAL_STORAGE = "NO_LOCAL_STORAGE"

	ERR_XENAPI_MISSING_PLUGIN = "XENAPI_MISSING_PLUGIN"
	ERR_XENAPI_PLUGIN_FAILURE = "XENAPI_PLUGIN_FAILURE"

	ERR_SR_ATTACHED     = "SR_ATTACHED"
	ERR_SR_NOT_ATTACHED = "SR_NOT_ATTACHED"

	ERR_DOMAIN_BUILDER_ERROR = "DOMAIN_BUILDER_ERROR"

	ERR_AUTH_ALREADY_ENABLED                         = "AUTH_ALREADY_ENABLED"
	ERR_AUTH_UNKNOWN_TYPE                            = "AUTH_UNKNOWN_TYPE"
	ERR_AUTH_IS_DISABLED                             = "AUTH_IS_DISABLED"
	ERR_AUTH_ENABLE_FAILED                           = "AUTH_ENABLE_FAILED"
	ERR_AUTH_ENABLE_FAILED_WRONG_CREDENTIALS         = "AUTH_ENABLE_FAILED_WRONG_CREDENTIALS"
	ERR_AUTH_ENABLE_FAILED_PERMISSION_DENIED         = "AUTH_ENABLE_FAILED_PERMISSION_DENIED"
	ERR_AUTH_ENABLE_FAILED_DOMAIN_LOOKUP_FAILED      = "AUTH_ENABLE_FAILED_DOMAIN_LOOKUP_FAILED"
	ERR_AUTH_ENABLE_FAILED_UNAVAILABLE               = "AUTH_ENABLE_FAILED_UNAVAILABLE"
	ERR_AUTH_ENABLE_FAILED_INVALID_OU                = "AUTH_ENABLE_FAILED_INVALID_OU"
	ERR_AUTH_DISABLE_FAILED                          = "AUTH_DISABLE_FAILED"
	ERR_AUTH_DISABLE_FAILED_WRONG_CREDENTIALS        = "AUTH_DISABLE_FAILED_WRONG_CREDENTIALS"
	ERR_AUTH_DISABLE_FAILED_PERMISSION_DENIED        = "AUTH_DISABLE_FAILED_PERMISSION_DENIED"
	ERR_POOL_AUTH_ALREADY_ENABLED                    = "POOL_AUTH_ALREADY_ENABLED"
	ERR_POOL_AUTH_ENABLE_FAILED                      = "POOL_AUTH_ENABLE_FAILED"
	ERR_POOL_AUTH_ENABLE_FAILED_WRONG_CREDENTIALS    = "POOL_AUTH_ENABLE_FAILED_WRONG_CREDENTIALS"
	ERR_POOL_AUTH_ENABLE_FAILED_PERMISSION_DENIED    = "POOL_AUTH_ENABLE_FAILED_PERMISSION_DENIED"
	ERR_POOL_AUTH_ENABLE_FAILED_DOMAIN_LOOKUP_FAILED = "POOL_AUTH_ENABLE_FAILED_DOMAIN_LOOKUP_FAILED"
	ERR_POOL_AUTH_ENABLE_FAILED_UNAVAILABLE          = "POOL_AUTH_ENABLE_FAILED_UNAVAILABLE"
	ERR_POOL_AUTH_ENABLE_FAILED_INVALID_OU           = "POOL_AUTH_ENABLE_FAILED_INVALID_OU"
	ERR_POOL_AUTH_ENABLE_FAILED_DUPLICATE_HOSTNAME   = "POOL_AUTH_ENABLE_FAILED_DUPLICATE_HOSTNAME"
	ERR_POOL_AUTH_DISABLE_FAILED                     = "POOL_AUTH_DISABLE_FAILED"
	ERR_POOL_AUTH_DISABLE_FAILED_WRONG_CREDENTIALS   = "POOL_AUTH_DISABLE_FAILED_WRONG_CREDENTIALS"
	ERR_POOL_AUTH_DISABLE_FAILED_PERMISSION_DENIED   = "POOL_AUTH_DISABLE_FAILED_PERMISSION_DENIED"
	ERR_SUBJECT_CANNOT_BE_RESOLVED                   = "SUBJECT_CANNOT_BE_RESOLVED"
	ERR_AUTH_SERVICE_ERROR                           = "AUTH_SERVICE_ERROR"
	ERR_SUBJECT_ALREADY_EXISTS                       = "SUBJECT_ALREADY_EXISTS"
	ERR_ROLE_NOT_FOUND                               = "ROLE_NOT_FOUND"
	ERR_ROLE_ALREADY_EXISTS                          = "ROLE_ALREADY_EXISTS"
	ERR_RBAC_PERMISSION_DENIED                       = "RBAC_PERMISSION_DENIED"

	ERR_CERTIFICATE_DOES_NOT_EXIST  = "CERTIFICATE_DOES_NOT_EXIST"
	ERR_CERTIFICATE_ALREADY_EXISTS  = "CERTIFICATE_ALREADY_EXISTS"
	ERR_CERTIFICATE_NAME_INVALID    = "CERTIFICATE_NAME_INVALID"
	ERR_CERTIFICATE_CORRUPT         = "CERTIFICATE_CORRUPT"
	ERR_CERTIFICATE_LIBRARY_CORRUPT = "CERTIFICATE_LIBRARY_CORRUPT"
	ERR_CRL_DOES_NOT_EXIST          = "CRL_DOES_NOT_EXIST"
	ERR_CRL_ALREADY_EXISTS          = "CRL_ALREADY_EXISTS"
	ERR_CRL_NAME_INVALID            = "CRL_NAME_INVALID"
	ERR_CRL_CORRUPT                 = "CRL_CORRUPT"

	ERR_VMPP_HAS_VM                            = "VMPP_HAS_VM"
	ERR_VMPP_ARCHIVE_MORE_FREQUENT_THAN_BACKUP = "VMPP_ARCHIVE_MORE_FREQUENT_THAN_BACKUP"
	ERR_VM_ASSIGNED_TO_PROTECTION_POLICY       = "VM_ASSIGNED_TO_PROTECTION_POLICY"

	ERR_SSL_VERIFY_ERROR = "SSL_VERIFY_ERROR"

	ERR_CANNOT_ENABLE_REDO_LOG = "CANNOT_ENABLE_REDO_LOG"
	ERR_REDO_LOG_IS_ENABLED    = "REDO_LOG_IS_ENABLED"

	ERR_VM_BIOS_STRINGS_ALREADY_SET = "VM_BIOS_STRINGS_ALREADY_SET"

	ERR_INVALID_FEATURE_STRING            = "INVALID_FEATURE_STRING"
	ERR_CPU_FEATURE_MASKING_NOT_SUPPORTED = "CPU_FEATURE_MASKING_NOT_SUPPORTED"

	ERR_FEATURE_REQUIRES_HVM = "FEATURE_REQUIRES_HVM"

	// (* DISASTER RECOVERY *)
	ERR_VDI_CONTAINS_METADATA_OF_THIS_POOL       = "VDI_CONTAINS_METADATA_OF_THIS_POOL"
	ERR_NO_MORE_REDO_LOGS_ALLOWED                = "NO_MORE_REDO_LOGS_ALLOWED"
	ERR_COULD_NOT_IMPORT_DATABASE                = "COULD_NOT_IMPORT_DATABASE"
	ERR_VM_INCOMPATIBLE_WITH_THIS_HOST           = "VM_INCOMPATIBLE_WITH_THIS_HOST"
	ERR_CANNOT_DESTROY_DISASTER_RECOVERY_TASK    = "CANNOT_DESTROY_DISASTER_RECOVERY_TASK"
	ERR_VM_IS_PART_OF_AN_APPLIANCE               = "VM_IS_PART_OF_AN_APPLIANCE"
	ERR_VM_TO_IMPORT_IS_NOT_NEWER_VERSION        = "VM_TO_IMPORT_IS_NOT_NEWER_VERSION"
	ERR_SUSPEND_VDI_REPLACEMENT_IS_NOT_IDENTICAL = "SUSPEND_VDI_REPLACEMENT_IS_NOT_IDENTICAL"
	ERR_VDI_COPY_FAILED                          = "VDI_COPY_FAILED"

	ERR_VDI_NEEDS_VM_FOR_MIGRATE  = "VDI_NEEDS_VM_FOR_MIGRATE"
	ERR_VM_HAS_TOO_MANY_SNAPSHOTS = "VM_HAS_TOO_MANY_SNAPSHOTS"
	ERR_VM_HAS_CHECKPOINT         = "VM_HAS_CHECKPOINT"

	ERR_MIRROR_FAILED                 = "MIRROR_FAILED"
	ERR_TOO_MANY_STORAGE_MIGRATES     = "TOO_MANY_STORAGE_MIGRATES"
	ERR_SR_DOES_NOT_SUPPORT_MIGRATION = "SR_DOES_NOT_SUPPORT_MIGRATION"
	ERR_UNIMPLEMENTED_IN_SM_BACKEND   = "UNIMPLEMENTED_IN_SM_BACKEND"

	ERR_VM_CALL_PLUGIN_RATE_LIMIT = "VM_CALL_PLUGIN_RATE_LIMIT"
)

// Error represents errors returned on xmlrpc request.
type Error struct {
	code    string
	objtype string
	uuid    string
}

// Error() method implements Error interface
func (e *Error) Error() string {
	return fmt.Sprintf("API Error: %s %s %s", e.code, e.objtype, e.uuid)
}

// Code ...
func (e *Error) Code() string {
	return e.code
}

// Type ...
func (e *Error) Type() string {
	return e.objtype
}

// UUID ...
func (e *Error) UUID() string {
	return e.uuid
}
