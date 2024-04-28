// Code generated by ogen, DO NOT EDIT.

package openshock

import (
	"fmt"
	"time"

	"github.com/go-faster/jx"
	"github.com/google/uuid"
)

// SetFake set fake values.
func (s *AccountSignUpV2Conflict) SetFake() {
	var unwrapped OpenShockProblem
	{
		unwrapped.SetFake()
	}
	*s = AccountSignUpV2Conflict(unwrapped)
}

// SetFake set fake values.
func (s *AccountSignUpV2Forbidden) SetFake() {
	var unwrapped OpenShockProblem
	{
		unwrapped.SetFake()
	}
	*s = AccountSignUpV2Forbidden(unwrapped)
}

// SetFake set fake values.
func (s *AdminOnlineDeviceResponse) SetFake() {
	{
		{
			s.ID.SetFake()
		}
	}
	{
		{
			s.Name.SetFake()
		}
	}
	{
		{
			s.Owner.SetFake()
		}
	}
	{
		{
			s.FirmwareVersion.SetFake()
		}
	}
	{
		{
			s.Gateway.SetFake()
		}
	}
	{
		{
			s.ConnectedAt.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *AdminOnlineDeviceResponseIEnumerableBaseResponse) SetFake() {
	{
		{
			s.Message.SetFake()
		}
	}
	{
		{
			s.Data.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *BooleanNullableBaseResponse) SetFake() {
	{
		{
			s.Message.SetFake()
		}
	}
	{
		{
			s.Data.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *Control) SetFake() {
	{
		{
			s.ID.SetFake()
		}
	}
	{
		{
			s.Type.SetFake()
		}
	}
	{
		{
			s.Intensity.SetFake()
		}
	}
	{
		{
			s.Duration.SetFake()
		}
	}
	{
		{
			s.Exclusive.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *ControlLogSenderLight) SetFake() {
	{
		{
			s.ID.SetFake()
		}
	}
	{
		{
			s.Name.SetFake()
		}
	}
	{
		{
			s.Image.SetFake()
		}
	}
	{
		{
			s.CustomName.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *ControlRequest) SetFake() {
	{
		{
			s.Shocks.SetFake()
		}
	}
	{
		{
			s.CustomName.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *ControlType) SetFake() {
	*s = ControlTypeStop
}

// SetFake set fake values.
func (s *CreateShareCode) SetFake() {
	{
		{
			s.Permissions.SetFake()
		}
	}
	{
		{
			s.Limits.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *CreateTokenRequest) SetFake() {
	{
		{
			s.Name.SetFake()
		}
	}
	{
		{
			s.Permissions.SetFake()
		}
	}
	{
		{
			s.ValidUntil.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *DeviceEdit) SetFake() {
	{
		{
			s.Name.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *DeviceSelfResponse) SetFake() {
	{
		{
			s.ID.SetFake()
		}
	}
	{
		{
			s.Name.SetFake()
		}
	}
	{
		{
			s.Shockers.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *DeviceSelfResponseBaseResponse) SetFake() {
	{
		{
			s.Message.SetFake()
		}
	}
	{
		{
			s.Data.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *DevicesGetLiveControlGatewayInfoNotFound) SetFake() {
	var unwrapped OpenShockProblem
	{
		unwrapped.SetFake()
	}
	*s = DevicesGetLiveControlGatewayInfoNotFound(unwrapped)
}

// SetFake set fake values.
func (s *DevicesGetLiveControlGatewayInfoPreconditionFailed) SetFake() {
	var unwrapped OpenShockProblem
	{
		unwrapped.SetFake()
	}
	*s = DevicesGetLiveControlGatewayInfoPreconditionFailed(unwrapped)
}

// SetFake set fake values.
func (s *EditTokenRequest) SetFake() {
	{
		{
			s.Name.SetFake()
		}
	}
	{
		{
			s.Permissions.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *GenericIni) SetFake() {
	{
		{
			s.ID.SetFake()
		}
	}
	{
		{
			s.Name.SetFake()
		}
	}
	{
		{
			s.Image.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *GuidBaseResponse) SetFake() {
	{
		{
			s.Message.SetFake()
		}
	}
	{
		{
			s.Data.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *LcgNodeResponse) SetFake() {
	{
		{
			s.Fqdn.SetFake()
		}
	}
	{
		{
			s.Country.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *LcgNodeResponseBaseResponse) SetFake() {
	{
		{
			s.Message.SetFake()
		}
	}
	{
		{
			s.Data.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *LcgResponse) SetFake() {
	{
		{
			s.Gateway.SetFake()
		}
	}
	{
		{
			s.Country.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *LcgResponseBaseResponse) SetFake() {
	{
		{
			s.Message.SetFake()
		}
	}
	{
		{
			s.Data.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *LogEntry) SetFake() {
	{
		{
			s.ID.SetFake()
		}
	}
	{
		{
			s.CreatedOn.SetFake()
		}
	}
	{
		{
			s.Type.SetFake()
		}
	}
	{
		{
			s.ControlledBy.SetFake()
		}
	}
	{
		{
			s.Intensity.SetFake()
		}
	}
	{
		{
			s.Duration.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *LogEntryIEnumerableBaseResponse) SetFake() {
	{
		{
			s.Message.SetFake()
		}
	}
	{
		{
			s.Data.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *Login) SetFake() {
	{
		{
			s.Password.SetFake()
		}
	}
	{
		{
			s.Email.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *LoginV2) SetFake() {
	{
		{
			s.Password = "string"
		}
	}
	{
		{
			s.Email = "string"
		}
	}
	{
		{
			s.TurnstileResponse = "string"
		}
	}
}

// SetFake set fake values.
func (s *MinimalShocker) SetFake() {
	{
		{
			s.ID.SetFake()
		}
	}
	{
		{
			s.RfId.SetFake()
		}
	}
	{
		{
			s.Model.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *NewShocker) SetFake() {
	{
		{
			s.Name.SetFake()
		}
	}
	{
		{
			s.RfId.SetFake()
		}
	}
	{
		{
			s.Device.SetFake()
		}
	}
	{
		{
			s.Model.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *ObjectBaseResponse) SetFake() {
	{
		{
			s.Message.SetFake()
		}
	}
	{
		{
			s.Data = []byte("null")
		}
	}
}

// SetFake set fake values.
func (s *OpenShockProblem) SetFake() {
	{
		{
			s.Type.SetFake()
		}
	}
	{
		{
			s.Title.SetFake()
		}
	}
	{
		{
			s.Status.SetFake()
		}
	}
	{
		{
			s.Detail.SetFake()
		}
	}
	{
		{
			s.Instance.SetFake()
		}
	}
	{
		{
			s.Message.SetFake()
		}
	}
	{
		{
			s.TraceId.SetFake()
		}
	}
	{
		{
			s.AdditionalProps.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *OpenShockProblemAdditional) SetFake() {
	var (
		elem jx.Raw
		m    map[string]jx.Raw = s.init()
	)
	for i := 0; i < 0; i++ {
		m[fmt.Sprintf("fake%d", i)] = elem
	}
}

// SetFake set fake values.
func (s *OptBool) SetFake() {
	var elem bool
	{
		elem = true
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptControlLogSenderLight) SetFake() {
	var elem ControlLogSenderLight
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptControlRequest) SetFake() {
	var elem ControlRequest
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptControlType) SetFake() {
	var elem ControlType
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptCreateShareCode) SetFake() {
	var elem CreateShareCode
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptCreateTokenRequest) SetFake() {
	var elem CreateTokenRequest
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptDateTime) SetFake() {
	var elem time.Time
	{
		elem = time.Now()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptDeviceEdit) SetFake() {
	var elem DeviceEdit
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptDeviceSelfResponse) SetFake() {
	var elem DeviceSelfResponse
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptEditTokenRequest) SetFake() {
	var elem EditTokenRequest
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptGenericIni) SetFake() {
	var elem GenericIni
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptInt) SetFake() {
	var elem int
	{
		elem = int(0)
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptInt32) SetFake() {
	var elem int32
	{
		elem = int32(0)
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptLcgNodeResponse) SetFake() {
	var elem LcgNodeResponse
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptLcgResponse) SetFake() {
	var elem LcgResponse
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptLogin) SetFake() {
	var elem Login
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptLoginV2) SetFake() {
	var elem LoginV2
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptNewShocker) SetFake() {
	var elem NewShocker
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptNilAdminOnlineDeviceResponseArray) SetFake() {
	s.Null = true
	s.Set = true
}

// SetFake set fake values.
func (s *OptNilBool) SetFake() {
	s.Null = true
	s.Set = true
}

// SetFake set fake values.
func (s *OptNilControlArray) SetFake() {
	s.Null = true
	s.Set = true
}

// SetFake set fake values.
func (s *OptNilDate) SetFake() {
	s.Null = true
	s.Set = true
}

// SetFake set fake values.
func (s *OptNilDateTime) SetFake() {
	s.Null = true
	s.Set = true
}

// SetFake set fake values.
func (s *OptNilInt32) SetFake() {
	s.Null = true
	s.Set = true
}

// SetFake set fake values.
func (s *OptNilLogEntryArray) SetFake() {
	s.Null = true
	s.Set = true
}

// SetFake set fake values.
func (s *OptNilMinimalShockerArray) SetFake() {
	s.Null = true
	s.Set = true
}

// SetFake set fake values.
func (s *OptNilOtaItemArray) SetFake() {
	s.Null = true
	s.Set = true
}

// SetFake set fake values.
func (s *OptNilOwnerShockerResponseArrayArray) SetFake() {
	s.Null = true
	s.Set = true
}

// SetFake set fake values.
func (s *OptNilPermissionTypeArray) SetFake() {
	s.Null = true
	s.Set = true
}

// SetFake set fake values.
func (s *OptNilResponseDeviceArray) SetFake() {
	s.Null = true
	s.Set = true
}

// SetFake set fake values.
func (s *OptNilResponseDeviceWithShockersArray) SetFake() {
	s.Null = true
	s.Set = true
}

// SetFake set fake values.
func (s *OptNilShareCodeInfoArray) SetFake() {
	s.Null = true
	s.Set = true
}

// SetFake set fake values.
func (s *OptNilShareInfoArray) SetFake() {
	s.Null = true
	s.Set = true
}

// SetFake set fake values.
func (s *OptNilShareLinkDeviceArray) SetFake() {
	s.Null = true
	s.Set = true
}

// SetFake set fake values.
func (s *OptNilShareLinkResponseArray) SetFake() {
	s.Null = true
	s.Set = true
}

// SetFake set fake values.
func (s *OptNilShareLinkShockerArray) SetFake() {
	s.Null = true
	s.Set = true
}

// SetFake set fake values.
func (s *OptNilSharedDeviceArray) SetFake() {
	s.Null = true
	s.Set = true
}

// SetFake set fake values.
func (s *OptNilSharedShockerArray) SetFake() {
	s.Null = true
	s.Set = true
}

// SetFake set fake values.
func (s *OptNilShockerResponseArray) SetFake() {
	s.Null = true
	s.Set = true
}

// SetFake set fake values.
func (s *OptNilString) SetFake() {
	s.Null = true
	s.Set = true
}

// SetFake set fake values.
func (s *OptNilTokenResponseArray) SetFake() {
	s.Null = true
	s.Set = true
}

// SetFake set fake values.
func (s *OptNilURI) SetFake() {
	s.Null = true
	s.Set = true
}

// SetFake set fake values.
func (s *OptOtaUpdateStatus) SetFake() {
	var elem OtaUpdateStatus
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptPasswordResetProcessData) SetFake() {
	var elem PasswordResetProcessData
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptPauseRequest) SetFake() {
	var elem PauseRequest
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptPublicShareLinkResponse) SetFake() {
	var elem PublicShareLinkResponse
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptRankType) SetFake() {
	var elem RankType
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptResetRequest) SetFake() {
	var elem ResetRequest
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptResponseDeviceWithToken) SetFake() {
	var elem ResponseDeviceWithToken
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptRootResponse) SetFake() {
	var elem RootResponse
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptSelfResponse) SetFake() {
	var elem SelfResponse
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptShareLinkCreate) SetFake() {
	var elem ShareLinkCreate
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptShareLinkEditShocker) SetFake() {
	var elem ShareLinkEditShocker
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptShareLinkResponse) SetFake() {
	var elem ShareLinkResponse
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptShockerLimits) SetFake() {
	var elem ShockerLimits
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptShockerModelType) SetFake() {
	var elem ShockerModelType
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptShockerPermissions) SetFake() {
	var elem ShockerPermissions
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptShockerWithDevice) SetFake() {
	var elem ShockerWithDevice
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptSignUp) SetFake() {
	var elem SignUp
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptSignUpV2) SetFake() {
	var elem SignUpV2
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptStatsResponse) SetFake() {
	var elem StatsResponse
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptTokenResponse) SetFake() {
	var elem TokenResponse
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptUUID) SetFake() {
	var elem uuid.UUID
	{
		elem = uuid.New()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OtaItem) SetFake() {
	{
		{
			s.ID.SetFake()
		}
	}
	{
		{
			s.StartedAt.SetFake()
		}
	}
	{
		{
			s.Status.SetFake()
		}
	}
	{
		{
			s.Version.SetFake()
		}
	}
	{
		{
			s.Message.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *OtaItemIReadOnlyCollectionBaseResponse) SetFake() {
	{
		{
			s.Message.SetFake()
		}
	}
	{
		{
			s.Data.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *OtaUpdateStatus) SetFake() {
	*s = OtaUpdateStatusStarted
}

// SetFake set fake values.
func (s *OwnerShockerResponse) SetFake() {
	{
		{
			s.ID.SetFake()
		}
	}
	{
		{
			s.Name.SetFake()
		}
	}
	{
		{
			s.Devices.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *OwnerShockerResponseIEnumerableIEnumerableBaseResponse) SetFake() {
	{
		{
			s.Message.SetFake()
		}
	}
	{
		{
			s.Data.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *PasswordResetProcessData) SetFake() {
	{
		{
			s.Password.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *PauseReasonBaseResponse) SetFake() {
	{
		{
			s.Message.SetFake()
		}
	}
	{
		{
			s.Data.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *PauseRequest) SetFake() {
	{
		{
			s.Pause.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *PermissionType) SetFake() {
	*s = PermissionTypeShockersUse
}

// SetFake set fake values.
func (s *PublicShareLinkResponse) SetFake() {
	{
		{
			s.ID.SetFake()
		}
	}
	{
		{
			s.Name.SetFake()
		}
	}
	{
		{
			s.CreatedOn.SetFake()
		}
	}
	{
		{
			s.ExpiresOn.SetFake()
		}
	}
	{
		{
			s.Author.SetFake()
		}
	}
	{
		{
			s.Devices.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *PublicShareLinkResponseBaseResponse) SetFake() {
	{
		{
			s.Message.SetFake()
		}
	}
	{
		{
			s.Data.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *RankType) SetFake() {
	*s = RankTypeUser
}

// SetFake set fake values.
func (s *ResetRequest) SetFake() {
	{
		{
			s.Email.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *ResponseDevice) SetFake() {
	{
		{
			s.ID.SetFake()
		}
	}
	{
		{
			s.Name.SetFake()
		}
	}
	{
		{
			s.CreatedOn.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *ResponseDeviceIEnumerableBaseResponse) SetFake() {
	{
		{
			s.Message.SetFake()
		}
	}
	{
		{
			s.Data.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *ResponseDeviceWithShockers) SetFake() {
	{
		{
			s.ID.SetFake()
		}
	}
	{
		{
			s.Name.SetFake()
		}
	}
	{
		{
			s.CreatedOn.SetFake()
		}
	}
	{
		{
			s.Shockers.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *ResponseDeviceWithShockersIEnumerableBaseResponse) SetFake() {
	{
		{
			s.Message.SetFake()
		}
	}
	{
		{
			s.Data.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *ResponseDeviceWithToken) SetFake() {
	{
		{
			s.ID.SetFake()
		}
	}
	{
		{
			s.Name.SetFake()
		}
	}
	{
		{
			s.CreatedOn.SetFake()
		}
	}
	{
		{
			s.Token.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *ResponseDeviceWithTokenBaseResponse) SetFake() {
	{
		{
			s.Message.SetFake()
		}
	}
	{
		{
			s.Data.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *RootResponse) SetFake() {
	{
		{
			s.Version.SetFake()
		}
	}
	{
		{
			s.Commit.SetFake()
		}
	}
	{
		{
			s.CurrentTime.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *RootResponseBaseResponse) SetFake() {
	{
		{
			s.Message.SetFake()
		}
	}
	{
		{
			s.Data.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *SelfResponse) SetFake() {
	{
		{
			s.ID.SetFake()
		}
	}
	{
		{
			s.Name.SetFake()
		}
	}
	{
		{
			s.Email.SetFake()
		}
	}
	{
		{
			s.Image.SetFake()
		}
	}
	{
		{
			s.Rank.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *SelfResponseBaseResponse) SetFake() {
	{
		{
			s.Message.SetFake()
		}
	}
	{
		{
			s.Data.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *ShareCodeInfo) SetFake() {
	{
		{
			s.ID.SetFake()
		}
	}
	{
		{
			s.CreatedOn.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *ShareCodeInfoIEnumerableBaseResponse) SetFake() {
	{
		{
			s.Message.SetFake()
		}
	}
	{
		{
			s.Data.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *ShareInfo) SetFake() {
	{
		{
			s.SharedWith.SetFake()
		}
	}
	{
		{
			s.CreatedOn.SetFake()
		}
	}
	{
		{
			s.Permissions.SetFake()
		}
	}
	{
		{
			s.Limits.SetFake()
		}
	}
	{
		{
			s.Paused.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *ShareInfoIEnumerableBaseResponse) SetFake() {
	{
		{
			s.Message.SetFake()
		}
	}
	{
		{
			s.Data.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *ShareLinkCreate) SetFake() {
	{
		{
			s.Name.SetFake()
		}
	}
	{
		{
			s.ExpiresOn.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *ShareLinkDevice) SetFake() {
	{
		{
			s.ID.SetFake()
		}
	}
	{
		{
			s.Name.SetFake()
		}
	}
	{
		{
			s.Shockers.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *ShareLinkEditShocker) SetFake() {
	{
		{
			s.Permissions.SetFake()
		}
	}
	{
		{
			s.Limits.SetFake()
		}
	}
	{
		{
			s.Cooldown.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *ShareLinkResponse) SetFake() {
	{
		{
			s.ID.SetFake()
		}
	}
	{
		{
			s.Name.SetFake()
		}
	}
	{
		{
			s.CreatedOn.SetFake()
		}
	}
	{
		{
			s.ExpiresOn.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *ShareLinkResponseBaseResponse) SetFake() {
	{
		{
			s.Message.SetFake()
		}
	}
	{
		{
			s.Data.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *ShareLinkResponseIEnumerableBaseResponse) SetFake() {
	{
		{
			s.Message.SetFake()
		}
	}
	{
		{
			s.Data.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *ShareLinkShocker) SetFake() {
	{
		{
			s.ID.SetFake()
		}
	}
	{
		{
			s.Name.SetFake()
		}
	}
	{
		{
			s.Permissions.SetFake()
		}
	}
	{
		{
			s.Limits.SetFake()
		}
	}
	{
		{
			s.Paused.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *ShareLinksAddShockerConflict) SetFake() {
	var unwrapped OpenShockProblem
	{
		unwrapped.SetFake()
	}
	*s = ShareLinksAddShockerConflict(unwrapped)
}

// SetFake set fake values.
func (s *ShareLinksAddShockerNotFound) SetFake() {
	var unwrapped OpenShockProblem
	{
		unwrapped.SetFake()
	}
	*s = ShareLinksAddShockerNotFound(unwrapped)
}

// SetFake set fake values.
func (s *SharedDevice) SetFake() {
	{
		{
			s.ID.SetFake()
		}
	}
	{
		{
			s.Name.SetFake()
		}
	}
	{
		{
			s.Shockers.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *SharedShocker) SetFake() {
	{
		{
			s.ID.SetFake()
		}
	}
	{
		{
			s.Name.SetFake()
		}
	}
	{
		{
			s.IsPaused.SetFake()
		}
	}
	{
		{
			s.Permissions.SetFake()
		}
	}
	{
		{
			s.Limits.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *SharesLinkShareCodeBadRequest) SetFake() {
	var unwrapped OpenShockProblem
	{
		unwrapped.SetFake()
	}
	*s = SharesLinkShareCodeBadRequest(unwrapped)
}

// SetFake set fake values.
func (s *SharesLinkShareCodeNotFound) SetFake() {
	var unwrapped OpenShockProblem
	{
		unwrapped.SetFake()
	}
	*s = SharesLinkShareCodeNotFound(unwrapped)
}

// SetFake set fake values.
func (s *ShockerLimits) SetFake() {
	{
		{
			s.Intensity.SetFake()
		}
	}
	{
		{
			s.Duration.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *ShockerModelType) SetFake() {
	*s = ShockerModelTypeCaiXianlin
}

// SetFake set fake values.
func (s *ShockerPermissions) SetFake() {
	{
		{
			s.Vibrate.SetFake()
		}
	}
	{
		{
			s.Sound.SetFake()
		}
	}
	{
		{
			s.Shock.SetFake()
		}
	}
	{
		{
			s.Live.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *ShockerRegisterShockerBadRequest) SetFake() {
	var unwrapped OpenShockProblem
	{
		unwrapped.SetFake()
	}
	*s = ShockerRegisterShockerBadRequest(unwrapped)
}

// SetFake set fake values.
func (s *ShockerRegisterShockerNotFound) SetFake() {
	var unwrapped OpenShockProblem
	{
		unwrapped.SetFake()
	}
	*s = ShockerRegisterShockerNotFound(unwrapped)
}

// SetFake set fake values.
func (s *ShockerResponse) SetFake() {
	{
		{
			s.ID.SetFake()
		}
	}
	{
		{
			s.RfId.SetFake()
		}
	}
	{
		{
			s.Model.SetFake()
		}
	}
	{
		{
			s.Name.SetFake()
		}
	}
	{
		{
			s.IsPaused.SetFake()
		}
	}
	{
		{
			s.CreatedOn.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *ShockerResponseIEnumerableBaseResponse) SetFake() {
	{
		{
			s.Message.SetFake()
		}
	}
	{
		{
			s.Data.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *ShockerSendControlDEPRECATEDForbidden) SetFake() {
	var unwrapped OpenShockProblem
	{
		unwrapped.SetFake()
	}
	*s = ShockerSendControlDEPRECATEDForbidden(unwrapped)
}

// SetFake set fake values.
func (s *ShockerSendControlDEPRECATEDNotFound) SetFake() {
	var unwrapped OpenShockProblem
	{
		unwrapped.SetFake()
	}
	*s = ShockerSendControlDEPRECATEDNotFound(unwrapped)
}

// SetFake set fake values.
func (s *ShockerSendControlDEPRECATEDPreconditionFailed) SetFake() {
	var unwrapped OpenShockProblem
	{
		unwrapped.SetFake()
	}
	*s = ShockerSendControlDEPRECATEDPreconditionFailed(unwrapped)
}

// SetFake set fake values.
func (s *ShockerSendControlForbidden) SetFake() {
	var unwrapped OpenShockProblem
	{
		unwrapped.SetFake()
	}
	*s = ShockerSendControlForbidden(unwrapped)
}

// SetFake set fake values.
func (s *ShockerSendControlNotFound) SetFake() {
	var unwrapped OpenShockProblem
	{
		unwrapped.SetFake()
	}
	*s = ShockerSendControlNotFound(unwrapped)
}

// SetFake set fake values.
func (s *ShockerSendControlPreconditionFailed) SetFake() {
	var unwrapped OpenShockProblem
	{
		unwrapped.SetFake()
	}
	*s = ShockerSendControlPreconditionFailed(unwrapped)
}

// SetFake set fake values.
func (s *ShockerWithDevice) SetFake() {
	{
		{
			s.ID.SetFake()
		}
	}
	{
		{
			s.RfId.SetFake()
		}
	}
	{
		{
			s.Model.SetFake()
		}
	}
	{
		{
			s.Name.SetFake()
		}
	}
	{
		{
			s.IsPaused.SetFake()
		}
	}
	{
		{
			s.CreatedOn.SetFake()
		}
	}
	{
		{
			s.Device.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *ShockerWithDeviceBaseResponse) SetFake() {
	{
		{
			s.Message.SetFake()
		}
	}
	{
		{
			s.Data.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *SignUp) SetFake() {
	{
		{
			s.Username.SetFake()
		}
	}
	{
		{
			s.Password.SetFake()
		}
	}
	{
		{
			s.Email.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *SignUpV2) SetFake() {
	{
		{
			s.Username.SetFake()
		}
	}
	{
		{
			s.Password.SetFake()
		}
	}
	{
		{
			s.Email.SetFake()
		}
	}
	{
		{
			s.TurnstileResponse = "string"
		}
	}
}

// SetFake set fake values.
func (s *StatsResponse) SetFake() {
	{
		{
			s.DevicesOnline.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *StatsResponseBaseResponse) SetFake() {
	{
		{
			s.Message.SetFake()
		}
	}
	{
		{
			s.Data.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *StringBaseResponse) SetFake() {
	{
		{
			s.Message.SetFake()
		}
	}
	{
		{
			s.Data.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *TokenResponse) SetFake() {
	{
		{
			s.ID.SetFake()
		}
	}
	{
		{
			s.Name.SetFake()
		}
	}
	{
		{
			s.CreatedOn.SetFake()
		}
	}
	{
		{
			s.CreatedByIp.SetFake()
		}
	}
	{
		{
			s.ValidUntil.SetFake()
		}
	}
	{
		{
			s.Permissions.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *TokenResponseBaseResponse) SetFake() {
	{
		{
			s.Message.SetFake()
		}
	}
	{
		{
			s.Data.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *TokenResponseIEnumerableBaseResponse) SetFake() {
	{
		{
			s.Message.SetFake()
		}
	}
	{
		{
			s.Data.SetFake()
		}
	}
}