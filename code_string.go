// Code generated by "stringer -linecomment -type=Code"; DO NOT EDIT.

package sysexit

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Ok-0]
	_ = x[ErrBase-64]
	_ = x[ErrUsage-64]
	_ = x[ErrData-65]
	_ = x[ErrNoInput-66]
	_ = x[ErrNoUser-67]
	_ = x[ErrNoHost-68]
	_ = x[ErrUnavailable-69]
	_ = x[ErrSoftware-70]
	_ = x[ErrOS-71]
	_ = x[ErrOSFile-72]
	_ = x[ErrCantCreat-73]
	_ = x[ErrIO-74]
	_ = x[ErrTempFail-75]
	_ = x[ErrProtocol-76]
	_ = x[ErrNoPerm-77]
	_ = x[ErrConfig-78]
}

const (
	_Code_name_0 = "successful termination"
	_Code_name_1 = "base value for error messagesdata format errorcannot open inputaddressee unknownhost name unknownservice unavailableinternal software errorsystem error (e.g., can't fork)critical OS file missingcan't create (user) output fileinput/output errortemp failure; user is invited to retryremote error in protocolpermission deniedconfiguration error"
)

var (
	_Code_index_1 = [...]uint16{0, 29, 46, 63, 80, 97, 116, 139, 170, 194, 225, 243, 281, 305, 322, 341}
)

func (i Code) String() string {
	switch {
	case i == 0:
		return _Code_name_0
	case 64 <= i && i <= 78:
		i -= 64
		return _Code_name_1[_Code_index_1[i]:_Code_index_1[i+1]]
	default:
		return "Code(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}