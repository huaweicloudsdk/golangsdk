package gophercloud

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
)

//server error Message
const (
	AuthRequired          = "Authentication required"
	PoilcyNotAllow        = "Policy doesn't allow .*. to be performed"
	TokenRoleEmpty        = "token role is empty, forbidden to perform this action"
	TokenRoleForbidden    = "token role * is forbidden to perform this action"
	ErrorRoleToPerform    = "do not have the required roles, forbbiden to perform this action"
	QuotaExceeded         = "Quota exceeded"
	PortNumberExceed      = "Maximum number of ports exceeded"
	VolumeNumberOver      = "Volume number is over limit"
	BlockImageNotFound    = "Block Device Mapping is Invalid: failed to get image.*."
	ImageNotFound         = "Image * could not be found."
	FlavorNotFound        = "Flavor .*. could not be found"
	NetworkNotFound       = "Network.*.could not be found"
	BlockDevInvalid       = "Block Device Mapping is Invalid"
	AZUnavailable         = "The requested availability zone is not available"
	SecurityGroupNotFound = "Security group .*. not found"
	KeyPairNotFound       = "Keypair .*. not found for user *"
	InstanceGroupNotFound = "Instance group .*. could not be found"
	InvalidMetadata       = "Invalid metadata.*"
	UserDataBase64        = "User data needs to be valid base 64"
	UserDataTooLarge      = "User data too large. User data must be no larger than .*"
	InstanceDiskExceed    = "The created instance's disk would be too small"
	FlavorMemoryNotEnough = "Flavor's memory is too small for requested image"
	InstanceNotFound      = "Instance .*. could not be found"
	InstanceIsLocked      = "Instance .*. is locked"
	InstCantBeOperated    = "Cannot .*. instance .*. while it is in .*"
	UnexpectedApiERROR    = "Unexpected API Error"
	ServerCantComply      = "The server could not comply with the request since it is either malformed.*."
	InvalidFlavorRef      = "Invalid flavorRef provided"
	InvalidKeyName        = "Invalid key_name provided"
	InvalidInputField     = "Invalid input for field/attribute"
)

//server StatusCode
const (
	Ecs1499 = "Ecs.1499"
	Ecs1500 = "Ecs.1500"
	Ecs1501 = "Ecs.1501"
	Ecs1502 = "Ecs.1502"
	Ecs1503 = "Ecs.1503"
	Ecs1511 = "Ecs.1511"
	Ecs1512 = "Ecs.1512"
	Ecs1513 = "Ecs.1513"
	Ecs1514 = "Ecs.1514"
	Ecs1515 = "Ecs.1515"
	Ecs1516 = "Ecs.1516"
	Ecs1517 = "Ecs.1517"
	Ecs1518 = "Ecs.1518"
	Ecs1519 = "Ecs.1519"
	Ecs1520 = "Ecs.1520"
	Ecs1521 = "Ecs.1521"
	Ecs1522 = "Ecs.1522"
	Ecs1523 = "Ecs.1523"
	Ecs1544 = "Ecs.1544"
	Ecs1545 = "Ecs.1545"
	Ecs1546 = "Ecs.1546"
	Ecs1599 = "Ecs.1599"
)

//common error code
const (
	Com1000           = "Com.1000"
	Com1001           = "Com.1001"
	Com1002           = "Com.1002"
	Com1003           = "Com.1003"
	Com1004           = "Com.1004"
	Com1005           = "Com.1005"
	DefaultCommonCode = "Com.1100"
)

//common error message
const (
	//Com1000
	MissingInput1     = "Missing input for argument"
	MissingInput2     = "Exactly one of * and * must be provided"
	MissingInput3     = "At least one of * and * must be provided"
	MissingInput4     = "Must have 1 and only 1 key-value pair"
	APINotFound       = "API not found"
	FlavorParamsError = "One and only one of the flavor ID and the flavor name must be provided"
	//Com1001
	StreamControlAPI = "The maximum request receiving rate is exceeded"
	//Com1002
	InvalidInput1 = "Invalid input provided for argument"
	InvalidInput2 = "Options type is not a struct"
	//Com1003
	ResourceNotFound1 = "Unable to find * with name *"
	ResourceNotFound2 = "Found * matching *"
	//Com1004
	NoClientProvided = "A service client must be provided to find a resource ID by name"
)

// BaseError is an error type that all other error types embed.
type BaseError struct {
	DefaultErrString string
	Info             string
}

func (e BaseError) Error() string {
	e.DefaultErrString = "An error occurred while executing a Gophercloud request."
	return e.choseErrString()
}

func (e BaseError) choseErrString() string {

	return ParseClientError(e)
}

// ErrMissingInput is the error when input is required in a particular
// situation but not provided by the user
type ErrMissingInput struct {
	BaseError
	Argument string
}

func (e ErrMissingInput) Error() string {
	e.DefaultErrString = fmt.Sprintf("Missing input for argument [%s]", e.Argument)
	return e.choseErrString()
}

// ErrInvalidInput is an error type used for most non-HTTP Gophercloud errors.
type ErrInvalidInput struct {
	ErrMissingInput
	Value interface{}
}

func (e ErrInvalidInput) Error() string {
	e.DefaultErrString = fmt.Sprintf("Invalid input provided for argument [%s]: [%+v]", e.Argument, e.Value)
	return e.choseErrString()
}

// ErrUnexpectedResponseCode is returned by the Request method when a response code other than
// those listed in OkCodes is encountered.
type ErrUnexpectedResponseCode struct {
	BaseError
	URL      string
	Method   string
	Expected []int
	Actual   int
	Body     []byte
}

func (e ErrUnexpectedResponseCode) Error() string {
	e.DefaultErrString = fmt.Sprintf(
		"Expected HTTP response code %v when accessing [%s %s], but got %d instead\n%s",
		e.Expected, e.Method, e.URL, e.Actual, e.Body,
	)
	return e.choseErrString()
}

// ErrDefault400 is the default error type returned on a 400 HTTP response code.
type ErrDefault400 struct {
	ErrUnexpectedResponseCode
}

// ErrDefault401 is the default error type returned on a 401 HTTP response code.
type ErrDefault401 struct {
	ErrUnexpectedResponseCode
}

type ErrDefault403 struct {
	ErrUnexpectedResponseCode
}

// ErrDefault404 is the default error type returned on a 404 HTTP response code.
type ErrDefault404 struct {
	ErrUnexpectedResponseCode
}

// ErrDefault405 is the default error type returned on a 405 HTTP response code.
type ErrDefault405 struct {
	ErrUnexpectedResponseCode
}

// ErrDefault408 is the default error type returned on a 408 HTTP response code.
type ErrDefault408 struct {
	ErrUnexpectedResponseCode
}

type ErrDefault409 struct {
	ErrUnexpectedResponseCode
}

// ErrDefault429 is the default error type returned on a 429 HTTP response code.
type ErrDefault429 struct {
	ErrUnexpectedResponseCode
}

// ErrDefault500 is the default error type returned on a 500 HTTP response code.
type ErrDefault500 struct {
	ErrUnexpectedResponseCode
}

// ErrDefault503 is the default error type returned on a 503 HTTP response code.
type ErrDefault503 struct {
	ErrUnexpectedResponseCode
}

type APIError struct {
	Message    string
	Request_id interface{}
}

type ErrMessage struct {
	Message string
	Code    interface{}
}

type ErrorMessage struct {
	ErrCode string
	Message string
}

func ParseClientError(e BaseError) string {
	var errorMessage ErrorMessage
	var errCode string
	var message string
	errCode = DefaultCommonCode
	if e.Info != "" {
		message = e.Info
	} else {
		message = e.DefaultErrString
	}
	if ok, _ := regexp.MatchString(MissingInput1, message); ok {
		errCode = Com1000
	}
	if ok, _ := regexp.MatchString(MissingInput2, message); ok {
		errCode = Com1000
	}
	if ok, _ := regexp.MatchString(MissingInput3, message); ok {
		errCode = Com1000
	}
	if ok, _ := regexp.MatchString(MissingInput4, message); ok {
		errCode = Com1000
	}
	if ok, _ := regexp.MatchString(FlavorParamsError, message); ok {
		errCode = Com1000
	}
	if ok, _ := regexp.MatchString(InvalidInput1, message); ok {
		errCode = Com1002
	}
	if ok, _ := regexp.MatchString(InvalidInput2, message); ok {
		errCode = Com1002
	}
	if ok, _ := regexp.MatchString(ResourceNotFound1, message); ok {
		errCode = Com1003
	}
	if ok, _ := regexp.MatchString(ResourceNotFound2, message); ok {
		errCode = Com1003
	}

	if ok, _ := regexp.MatchString(NoClientProvided, message); ok {
		errCode = Com1004
	}

	errorMessage.Message = message
	errorMessage.ErrCode = errCode
	msg, err2 := json.Marshal(errorMessage)
	if err2 != nil {
		return fmt.Sprintf("{'Message': %s, 'StatusCode': %s}", message, errCode)
	}
	return string(msg)
}

func Parse(e ErrUnexpectedResponseCode) string {
	//e.Body "{'badRequest':{'message':'Network could not be found.','code':'400'}}"
	var errorMessage ErrorMessage
	var apiError APIError
	var errCode string
	var message string
	var errMsgBody = make(map[string]ErrMessage)

	err := json.Unmarshal(e.Body, &errMsgBody)
	if err != nil {
		err1 := json.Unmarshal(e.Body, &apiError)
		if err1 != nil {
			message = string(e.Body)
		} else {
			message = apiError.Message
		}
	} else {
		for _, em := range errMsgBody {
			message = em.Message
		}
	}
	errCode = "Ecs." + strconv.Itoa(e.Actual)
	switch e.Actual {
	case 503:
		if ok, _ := regexp.MatchString(StreamControlAPI, string(e.Body)); ok {
			errCode = Com1001
			break
		}
	case 400:
		if ok, _ := regexp.MatchString(VolumeNumberOver, message); ok {
			errCode = Ecs1503
			break
		}

		if ok, _ := regexp.MatchString(BlockImageNotFound, message); ok {
			errCode = Ecs1511
			break
		}

		if ok, _ := regexp.MatchString(ImageNotFound, message); ok {
			errCode = Ecs1511
			break
		}

		if ok, _ := regexp.MatchString(FlavorNotFound, message); ok {
			errCode = Ecs1512
			break
		}
		if ok, _ := regexp.MatchString(InvalidFlavorRef, message); ok {
			errCode = Ecs1512
			break
		}

		if ok, _ := regexp.MatchString(NetworkNotFound, message); ok {
			errCode = Ecs1513
			break
		}

		if ok, _ := regexp.MatchString(BlockDevInvalid, message); ok {
			errCode = Ecs1514
			break
		}

		if ok, _ := regexp.MatchString(AZUnavailable, message); ok {
			errCode = Ecs1515
			break
		}

		if ok, _ := regexp.MatchString(SecurityGroupNotFound, message); ok {
			errCode = Ecs1516
			break
		}

		if ok, _ := regexp.MatchString(KeyPairNotFound, message); ok {
			errCode = Ecs1517
			break
		}

		if ok, _ := regexp.MatchString(InvalidKeyName, message); ok {
			errCode = Ecs1517
			break
		}

		if ok, _ := regexp.MatchString(InstanceGroupNotFound, message); ok {
			errCode = Ecs1518
			break
		}

		if ok, _ := regexp.MatchString(InvalidMetadata, message); ok {
			errCode = Ecs1519
			break
		}

		if ok, _ := regexp.MatchString(InvalidInputField, message); ok {
			errCode = Ecs1519
			break
		}

		if ok, _ := regexp.MatchString(UserDataBase64, message); ok {
			errCode = Ecs1520
			break
		}

		if ok, _ := regexp.MatchString(UserDataTooLarge, message); ok {
			errCode = Ecs1521
			break
		}

		if ok, _ := regexp.MatchString(InstanceDiskExceed, message); ok {
			errCode = Ecs1522
			break
		}

		if ok, _ := regexp.MatchString(FlavorMemoryNotEnough, message); ok {
			errCode = Ecs1523
			break
		}

		if ok, _ := regexp.MatchString(ServerCantComply, message); ok {
			errCode = Ecs1599
			break
		}

		if ok, _ := regexp.MatchString(UnexpectedApiERROR, message); ok {
			errCode = Ecs1599
			break
		}
	case 401:
		if ok, _ := regexp.MatchString(AuthRequired, message); ok {
			errCode = Ecs1499
			break
		}
	case 403:
		if ok, _ := regexp.MatchString(PoilcyNotAllow, message); ok {
			errCode = Ecs1500
			break
		}

		if ok, _ := regexp.MatchString(TokenRoleEmpty, message); ok {
			errCode = Ecs1500
			break
		}

		if ok, _ := regexp.MatchString(TokenRoleForbidden, message); ok {
			errCode = Ecs1500
			break
		}

		if ok, _ := regexp.MatchString(ErrorRoleToPerform, message); ok {
			errCode = Ecs1500
			break
		}

		if ok, _ := regexp.MatchString(QuotaExceeded, message); ok {
			errCode = Ecs1501
			break
		}
		if ok, _ := regexp.MatchString(PortNumberExceed, message); ok {
			errCode = Ecs1502
			break
		}
	case 404:
		if ok, _ := regexp.MatchString(InstanceNotFound, message); ok {
			errCode = Ecs1544
			break
		}

		if ok, _ := regexp.MatchString(APINotFound, message); ok {
			errCode = Com1005
			break
		}
	case 409:
		if ok, _ := regexp.MatchString(InstanceIsLocked, message); ok {
			errCode = Ecs1545
			break
		}
		if ok, _ := regexp.MatchString(InstCantBeOperated, message); ok {
			errCode = Ecs1546
			break
		}
	default:
		message = string(e.Body)
	}
	errorMessage.Message = message
	errorMessage.ErrCode = errCode
	msg, err2 := json.Marshal(errorMessage)
	if err2 != nil {
		return fmt.Sprintf("{'Message': %s, 'StatusCode': %s}", message, errCode)
	}
	return string(msg)
}

func (e ErrDefault400) Error() string {
	return Parse(e.ErrUnexpectedResponseCode)
}

func (e ErrDefault401) Error() string {
	//return "Authentication failed"
	return Parse(e.ErrUnexpectedResponseCode)
}

func (e ErrDefault403) Error() string {
	return Parse(e.ErrUnexpectedResponseCode)
}

func (e ErrDefault404) Error() string {
	return Parse(e.ErrUnexpectedResponseCode)
}
func (e ErrDefault405) Error() string {
	return Parse(e.ErrUnexpectedResponseCode)
}
func (e ErrDefault408) Error() string {
	//return "The server timed out waiting for the request"
	return Parse(e.ErrUnexpectedResponseCode)
}

func (e ErrDefault409) Error() string {
	return Parse(e.ErrUnexpectedResponseCode)
}

func (e ErrDefault429) Error() string {
	return "Too many requests have been sent in a given amount of time. Pause" +
		" requests, wait up to one minute, and try again."
}
func (e ErrDefault500) Error() string {
	//return "Internal Server Error"
	return Parse(e.ErrUnexpectedResponseCode)
}
func (e ErrDefault503) Error() string {
	//return "The service is currently unable to handle the request due to a temporary" +
	//	" overloading or maintenance. This is a temporary condition. Try again later."

	return Parse(e.ErrUnexpectedResponseCode)
}

// Err400er is the interface resource error types implement to override the error message
// from a 400 error.
type Err400er interface {
	Error400(ErrUnexpectedResponseCode) error
}

// Err401er is the interface resource error types implement to override the error message
// from a 401 error.
type Err401er interface {
	Error401(ErrUnexpectedResponseCode) error
}

// Err404er is the interface resource error types implement to override the error message
// from a 404 error.
type Err404er interface {
	Error404(ErrUnexpectedResponseCode) error
}

// Err405er is the interface resource error types implement to override the error message
// from a 405 error.
type Err405er interface {
	Error405(ErrUnexpectedResponseCode) error
}

// Err408er is the interface resource error types implement to override the error message
// from a 408 error.
type Err408er interface {
	Error408(ErrUnexpectedResponseCode) error
}

// Err429er is the interface resource error types implement to override the error message
// from a 429 error.
type Err429er interface {
	Error429(ErrUnexpectedResponseCode) error
}

// Err500er is the interface resource error types implement to override the error message
// from a 500 error.
type Err500er interface {
	Error500(ErrUnexpectedResponseCode) error
}

// Err503er is the interface resource error types implement to override the error message
// from a 503 error.
type Err503er interface {
	Error503(ErrUnexpectedResponseCode) error
}

// ErrTimeOut is the error type returned when an operations times out.
type ErrTimeOut struct {
	BaseError
}

func (e ErrTimeOut) Error() string {
	e.DefaultErrString = "A time out occurred"
	return e.choseErrString()
}

// ErrUnableToReauthenticate is the error type returned when reauthentication fails.
type ErrUnableToReauthenticate struct {
	BaseError
	ErrOriginal error
}

func (e ErrUnableToReauthenticate) Error() string {
	e.DefaultErrString = fmt.Sprintf("Unable to re-authenticate: %s", e.ErrOriginal)
	return e.choseErrString()
}

// ErrErrorAfterReauthentication is the error type returned when reauthentication
// succeeds, but an error occurs afterword (usually an HTTP error).
type ErrErrorAfterReauthentication struct {
	BaseError
	ErrOriginal error
}

func (e ErrErrorAfterReauthentication) Error() string {
	e.DefaultErrString = fmt.Sprintf("Successfully re-authenticated, but got error executing request: %s", e.ErrOriginal)
	return e.choseErrString()
}

// ErrServiceNotFound is returned when no service in a service catalog matches
// the provided EndpointOpts. This is generally returned by provider service
// factory methods like "NewComputeV2()" and can mean that a service is not
// enabled for your account.
type ErrServiceNotFound struct {
	BaseError
}

func (e ErrServiceNotFound) Error() string {
	e.DefaultErrString = "No suitable service could be found in the service catalog."
	return e.choseErrString()
}

// ErrEndpointNotFound is returned when no available endpoints match the
// provided EndpointOpts. This is also generally returned by provider service
// factory methods, and usually indicates that a region was specified
// incorrectly.
type ErrEndpointNotFound struct {
	BaseError
}

func (e ErrEndpointNotFound) Error() string {
	e.DefaultErrString = "No suitable endpoint could be found in the service catalog."
	return e.choseErrString()
}

// ErrResourceNotFound is the error when trying to retrieve a resource's
// ID by name and the resource doesn't exist.
type ErrResourceNotFound struct {
	BaseError
	Name         string
	ResourceType string
}

func (e ErrResourceNotFound) Error() string {
	e.DefaultErrString = fmt.Sprintf("Unable to find %s with name %s", e.ResourceType, e.Name)
	return e.choseErrString()
}

// ErrMultipleResourcesFound is the error when trying to retrieve a resource's
// ID by name and multiple resources have the user-provided name.
type ErrMultipleResourcesFound struct {
	BaseError
	Name         string
	Count        int
	ResourceType string
}

func (e ErrMultipleResourcesFound) Error() string {
	e.DefaultErrString = fmt.Sprintf("Found %d %ss matching %s", e.Count, e.ResourceType, e.Name)
	return e.choseErrString()
}

// ErrUnexpectedType is the error when an unexpected type is encountered
type ErrUnexpectedType struct {
	BaseError
	Expected string
	Actual   string
}

func (e ErrUnexpectedType) Error() string {
	e.DefaultErrString = fmt.Sprintf("Expected %s but got %s", e.Expected, e.Actual)
	return e.choseErrString()
}

func unacceptedAttributeErr(attribute string) string {
	return fmt.Sprintf("The base Identity V3 API does not accept authentication by %s", attribute)
}

func redundantWithTokenErr(attribute string) string {
	return fmt.Sprintf("%s may not be provided when authenticating with a TokenID", attribute)
}

func redundantWithUserID(attribute string) string {
	return fmt.Sprintf("%s may not be provided when authenticating with a UserID", attribute)
}

// ErrAPIKeyProvided indicates that an APIKey was provided but can't be used.
type ErrAPIKeyProvided struct{ BaseError }

func (e ErrAPIKeyProvided) Error() string {
	return unacceptedAttributeErr("APIKey")
}

// ErrTenantIDProvided indicates that a TenantID was provided but can't be used.
type ErrTenantIDProvided struct{ BaseError }

func (e ErrTenantIDProvided) Error() string {
	return unacceptedAttributeErr("TenantID")
}

// ErrTenantNameProvided indicates that a TenantName was provided but can't be used.
type ErrTenantNameProvided struct{ BaseError }

func (e ErrTenantNameProvided) Error() string {
	return unacceptedAttributeErr("TenantName")
}

// ErrUsernameWithToken indicates that a Username was provided, but token authentication is being used instead.
type ErrUsernameWithToken struct{ BaseError }

func (e ErrUsernameWithToken) Error() string {
	return redundantWithTokenErr("Username")
}

// ErrUserIDWithToken indicates that a UserID was provided, but token authentication is being used instead.
type ErrUserIDWithToken struct{ BaseError }

func (e ErrUserIDWithToken) Error() string {
	return redundantWithTokenErr("UserID")
}

// ErrDomainIDWithToken indicates that a DomainID was provided, but token authentication is being used instead.
type ErrDomainIDWithToken struct{ BaseError }

func (e ErrDomainIDWithToken) Error() string {
	return redundantWithTokenErr("DomainID")
}

// ErrDomainNameWithToken indicates that a DomainName was provided, but token authentication is being used instead.s
type ErrDomainNameWithToken struct{ BaseError }

func (e ErrDomainNameWithToken) Error() string {
	return redundantWithTokenErr("DomainName")
}

// ErrUsernameOrUserID indicates that neither username nor userID are specified, or both are at once.
type ErrUsernameOrUserID struct{ BaseError }

func (e ErrUsernameOrUserID) Error() string {
	return "Exactly one of Username and UserID must be provided for password authentication"
}

// ErrDomainIDWithUserID indicates that a DomainID was provided, but unnecessary because a UserID is being used.
type ErrDomainIDWithUserID struct{ BaseError }

func (e ErrDomainIDWithUserID) Error() string {
	return redundantWithUserID("DomainID")
}

// ErrDomainNameWithUserID indicates that a DomainName was provided, but unnecessary because a UserID is being used.
type ErrDomainNameWithUserID struct{ BaseError }

func (e ErrDomainNameWithUserID) Error() string {
	return redundantWithUserID("DomainName")
}

// ErrDomainIDOrDomainName indicates that a username was provided, but no domain to scope it.
// It may also indicate that both a DomainID and a DomainName were provided at once.
type ErrDomainIDOrDomainName struct{ BaseError }

func (e ErrDomainIDOrDomainName) Error() string {
	return "You must provide exactly one of DomainID or DomainName to authenticate by Username"
}

// ErrMissingPassword indicates that no password was provided and no token is available.
type ErrMissingPassword struct{ BaseError }

func (e ErrMissingPassword) Error() string {
	return "You must provide a password to authenticate"
}

// ErrScopeDomainIDOrDomainName indicates that a domain ID or Name was required in a Scope, but not present.
type ErrScopeDomainIDOrDomainName struct{ BaseError }

func (e ErrScopeDomainIDOrDomainName) Error() string {
	return "You must provide exactly one of DomainID or DomainName in a Scope with ProjectName"
}

// ErrScopeProjectIDOrProjectName indicates that both a ProjectID and a ProjectName were provided in a Scope.
type ErrScopeProjectIDOrProjectName struct{ BaseError }

func (e ErrScopeProjectIDOrProjectName) Error() string {
	return "You must provide at most one of ProjectID or ProjectName in a Scope"
}

// ErrScopeProjectIDAlone indicates that a ProjectID was provided with other constraints in a Scope.
type ErrScopeProjectIDAlone struct{ BaseError }

func (e ErrScopeProjectIDAlone) Error() string {
	return "ProjectID must be supplied alone in a Scope"
}

// ErrScopeEmpty indicates that no credentials were provided in a Scope.
type ErrScopeEmpty struct{ BaseError }

func (e ErrScopeEmpty) Error() string {
	return "You must provide either a Project or Domain in a Scope"
}
