// Code generated by "protoc-gen-orm-go", DO NOT EDIT.

package library

import (
	bytes "bytes"
)

func (x *Author) Pick() *AuthorGetRequest {
	return AuthorById(x.GetId())
}

func (x *AuthorGetRequest) Picks(v *Author) bool {
	switch x.WhichKey() {
	case AuthorGetRequest_Id_case:
		return bytes.Equal(x.GetId(), v.GetId())

	case AuthorGetRequest_Alias_case:
		return x.GetAlias() == v.GetAlias()

	default:
		return false
	}
}
func (x *AuthorGetRequest) Select(v *AuthorSelect) *AuthorGetRequest {
	x.SetSelect(v)
	return x
}

func AuthorById(v []byte) *AuthorGetRequest {
	if v == nil {
		return nil
	}

	return AuthorGetRequest_builder{Id: v}.Build()
}

func AuthorByAlias(v string) *AuthorGetRequest {
	return AuthorGetRequest_builder{Alias: &v}.Build()
}
func (x *Book) Pick() *BookGetRequest {
	return BookById(x.GetId())
}

func (x *BookGetRequest) Picks(v *Book) bool {
	switch x.WhichKey() {
	case BookGetRequest_Id_case:
		return bytes.Equal(x.GetId(), v.GetId())

	default:
		return false
	}
}
func (x *BookGetRequest) Select(v *BookSelect) *BookGetRequest {
	x.SetSelect(v)
	return x
}

func BookById(v []byte) *BookGetRequest {
	if v == nil {
		return nil
	}

	return BookGetRequest_builder{Id: v}.Build()
}
func (x *FooEf) Pick() *FooEfGetRequest {
	return FooEfById(x.GetId())
}

func (x *FooEfGetRequest) Picks(v *FooEf) bool {
	switch x.WhichKey() {
	case FooEfGetRequest_Id_case:
		return x.GetId() == v.GetId()

	default:
		return false
	}
}
func (x *FooEfGetRequest) Select(v *FooEfSelect) *FooEfGetRequest {
	x.SetSelect(v)
	return x
}

func FooEfById(v int64) *FooEfGetRequest {
	return FooEfGetRequest_builder{Id: &v}.Build()
}
func (x *FooK) Pick() *FooKGetRequest {
	return FooKById(x.GetId())
}

func (x *FooKGetRequest) Picks(v *FooK) bool {
	switch x.WhichKey() {
	case FooKGetRequest_Id_case:
		return x.GetId() == v.GetId()

	default:
		return false
	}
}
func (x *FooKGetRequest) Select(v *FooKSelect) *FooKGetRequest {
	x.SetSelect(v)
	return x
}

func FooKById(v int64) *FooKGetRequest {
	return FooKGetRequest_builder{Id: &v}.Build()
}
func (x *FooMi) Pick() *FooMiGetRequest {
	return FooMiById(x.GetId())
}

func (x *FooMiGetRequest) Picks(v *FooMi) bool {
	switch x.WhichKey() {
	case FooMiGetRequest_Id_case:
		return x.GetId() == v.GetId()

	default:
		return false
	}
}
func (x *FooMiGetRequest) Select(v *FooMiSelect) *FooMiGetRequest {
	x.SetSelect(v)
	return x
}

func FooMiById(v int64) *FooMiGetRequest {
	return FooMiGetRequest_builder{Id: &v}.Build()
}
func (x *FooMs) Pick() *FooMsGetRequest {
	return FooMsById(x.GetId())
}

func (x *FooMsGetRequest) Picks(v *FooMs) bool {
	switch x.WhichKey() {
	case FooMsGetRequest_Id_case:
		return x.GetId() == v.GetId()

	default:
		return false
	}
}
func (x *FooMsGetRequest) Select(v *FooMsSelect) *FooMsGetRequest {
	x.SetSelect(v)
	return x
}

func FooMsById(v int64) *FooMsGetRequest {
	return FooMsGetRequest_builder{Id: &v}.Build()
}
func (x *FooMsd) Pick() *FooMsdGetRequest {
	return FooMsdById(x.GetId())
}

func (x *FooMsdGetRequest) Picks(v *FooMsd) bool {
	switch x.WhichKey() {
	case FooMsdGetRequest_Id_case:
		return x.GetId() == v.GetId()

	default:
		return false
	}
}
func (x *FooMsdGetRequest) Select(v *FooMsdSelect) *FooMsdGetRequest {
	x.SetSelect(v)
	return x
}

func FooMsdById(v int64) *FooMsdGetRequest {
	return FooMsdGetRequest_builder{Id: &v}.Build()
}
func (x *FooV) Pick() *FooVGetRequest {
	return FooVById(x.GetId())
}

func (x *FooVGetRequest) Picks(v *FooV) bool {
	switch x.WhichKey() {
	case FooVGetRequest_Id_case:
		return x.GetId() == v.GetId()

	default:
		return false
	}
}
func (x *FooVGetRequest) Select(v *FooVSelect) *FooVGetRequest {
	x.SetSelect(v)
	return x
}

func FooVById(v int64) *FooVGetRequest {
	return FooVGetRequest_builder{Id: &v}.Build()
}
func (x *FooVd) Pick() *FooVdGetRequest {
	return FooVdById(x.GetId())
}

func (x *FooVdGetRequest) Picks(v *FooVd) bool {
	switch x.WhichKey() {
	case FooVdGetRequest_Id_case:
		return x.GetId() == v.GetId()

	default:
		return false
	}
}
func (x *FooVdGetRequest) Select(v *FooVdSelect) *FooVdGetRequest {
	x.SetSelect(v)
	return x
}

func FooVdById(v int64) *FooVdGetRequest {
	return FooVdGetRequest_builder{Id: &v}.Build()
}
func (x *FooVo) Pick() *FooVoGetRequest {
	return FooVoById(x.GetId())
}

func (x *FooVoGetRequest) Picks(v *FooVo) bool {
	switch x.WhichKey() {
	case FooVoGetRequest_Id_case:
		return x.GetId() == v.GetId()

	default:
		return false
	}
}
func (x *FooVoGetRequest) Select(v *FooVoSelect) *FooVoGetRequest {
	x.SetSelect(v)
	return x
}

func FooVoById(v int64) *FooVoGetRequest {
	return FooVoGetRequest_builder{Id: &v}.Build()
}
func (x *FooVon) Pick() *FooVonGetRequest {
	return FooVonById(x.GetId())
}

func (x *FooVonGetRequest) Picks(v *FooVon) bool {
	switch x.WhichKey() {
	case FooVonGetRequest_Id_case:
		return x.GetId() == v.GetId()

	default:
		return false
	}
}
func (x *FooVonGetRequest) Select(v *FooVonSelect) *FooVonGetRequest {
	x.SetSelect(v)
	return x
}

func FooVonById(v int64) *FooVonGetRequest {
	return FooVonGetRequest_builder{Id: &v}.Build()
}
func (x *FooVr) Pick() *FooVrGetRequest {
	return FooVrById(x.GetId())
}

func (x *FooVrGetRequest) Picks(v *FooVr) bool {
	switch x.WhichKey() {
	case FooVrGetRequest_Id_case:
		return x.GetId() == v.GetId()

	default:
		return false
	}
}
func (x *FooVrGetRequest) Select(v *FooVrSelect) *FooVrGetRequest {
	x.SetSelect(v)
	return x
}

func FooVrById(v int64) *FooVrGetRequest {
	return FooVrGetRequest_builder{Id: &v}.Build()
}
func (x *Like) Pick() *LikeGetRequest {
	return LikeById(x.GetId())
}

func (x *LikeGetRequest) Picks(v *Like) bool {
	switch x.WhichKey() {
	case LikeGetRequest_Id_case:
		return bytes.Equal(x.GetId(), v.GetId())

	case LikeGetRequest_Holders_case:
		return true &&
			bytes.Equal(x.GetHolders().GetSubjectId(), v.GetSubjectId()) &&
			x.GetHolders().GetActor().Picks(v.GetActor())
	default:
		return false
	}
}
func (x *LikeGetRequest) Select(v *LikeSelect) *LikeGetRequest {
	x.SetSelect(v)
	return x
}

func LikeById(v []byte) *LikeGetRequest {
	if v == nil {
		return nil
	}

	return LikeGetRequest_builder{Id: v}.Build()
}

func LikeByHolders(subject_id []byte, actor *MemberGetRequest) *LikeGetRequest {
	v := &LikeGetByHolders{}
	v.SetSubjectId(subject_id)
	v.SetActor(actor)
	return LikeGetRequest_builder{Holders: v}.Build()
}
func (x *Loan) Pick() *LoanGetRequest {
	return LoanById(x.GetId())
}

func (x *LoanGetRequest) Picks(v *Loan) bool {
	switch x.WhichKey() {
	case LoanGetRequest_Id_case:
		return bytes.Equal(x.GetId(), v.GetId())

	default:
		return false
	}
}
func (x *LoanGetRequest) Select(v *LoanSelect) *LoanGetRequest {
	x.SetSelect(v)
	return x
}

func LoanById(v []byte) *LoanGetRequest {
	if v == nil {
		return nil
	}

	return LoanGetRequest_builder{Id: v}.Build()
}
func (x *Locker) Pick() *LockerGetRequest {
	return LockerById(x.GetId())
}

func (x *LockerGetRequest) Picks(v *Locker) bool {
	switch x.WhichKey() {
	case LockerGetRequest_Id_case:
		return bytes.Equal(x.GetId(), v.GetId())

	case LockerGetRequest_Alias_case:
		return x.GetAlias() == v.GetAlias()

	default:
		return false
	}
}
func (x *LockerGetRequest) Select(v *LockerSelect) *LockerGetRequest {
	x.SetSelect(v)
	return x
}

func LockerById(v []byte) *LockerGetRequest {
	if v == nil {
		return nil
	}

	return LockerGetRequest_builder{Id: v}.Build()
}

func LockerByAlias(v string) *LockerGetRequest {
	return LockerGetRequest_builder{Alias: &v}.Build()
}
func (x *Member) Pick() *MemberGetRequest {
	return MemberById(x.GetId())
}

func (x *MemberGetRequest) Picks(v *Member) bool {
	switch x.WhichKey() {
	case MemberGetRequest_Id_case:
		return bytes.Equal(x.GetId(), v.GetId())

	case MemberGetRequest_Name_case:
		return true &&
			x.GetName().GetParent().Picks(v.GetParent()) &&
			x.GetName().GetName() == v.GetName()
	default:
		return false
	}
}
func (x *MemberGetRequest) Select(v *MemberSelect) *MemberGetRequest {
	x.SetSelect(v)
	return x
}

func MemberById(v []byte) *MemberGetRequest {
	if v == nil {
		return nil
	}

	return MemberGetRequest_builder{Id: v}.Build()
}

func MemberByName(parent *MemberGetRequest, name string) *MemberGetRequest {
	v := &MemberGetByName{}
	v.SetParent(parent)
	v.SetName(name)
	return MemberGetRequest_builder{Name: v}.Build()
}
func (x *Membership) Pick() *MembershipGetRequest {
	return MembershipById(x.GetId())
}

func (x *MembershipGetRequest) Picks(v *Membership) bool {
	switch x.WhichKey() {
	case MembershipGetRequest_Id_case:
		return bytes.Equal(x.GetId(), v.GetId())

	default:
		return false
	}
}
func (x *MembershipGetRequest) Select(v *MembershipSelect) *MembershipGetRequest {
	x.SetSelect(v)
	return x
}

func MembershipById(v []byte) *MembershipGetRequest {
	if v == nil {
		return nil
	}

	return MembershipGetRequest_builder{Id: v}.Build()
}
func (x *Press) Pick() *PressGetRequest {
	return PressById(x.GetId())
}

func (x *PressGetRequest) Picks(v *Press) bool {
	switch x.WhichKey() {
	case PressGetRequest_Id_case:
		return bytes.Equal(x.GetId(), v.GetId())

	case PressGetRequest_SerialNumber_case:
		return true &&
			x.GetSerialNumber().GetBook().Picks(v.GetBook()) &&
			x.GetSerialNumber().GetSerialNumber() == v.GetSerialNumber()
	default:
		return false
	}
}
func (x *PressGetRequest) Select(v *PressSelect) *PressGetRequest {
	x.SetSelect(v)
	return x
}

func PressById(v []byte) *PressGetRequest {
	if v == nil {
		return nil
	}

	return PressGetRequest_builder{Id: v}.Build()
}

func PressBySerialNumber(book *BookGetRequest, serial_number string) *PressGetRequest {
	v := &PressGetBySerialNumber{}
	v.SetBook(book)
	v.SetSerialNumber(serial_number)
	return PressGetRequest_builder{SerialNumber: v}.Build()
}
