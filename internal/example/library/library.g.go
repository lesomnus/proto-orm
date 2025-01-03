// Code generated by "protoc-gen-orm-go". DO NOT EDIT

package library

func AuthorById(v []byte) *AuthorGetRequest {
	return AuthorGetRequest_builder{Id: v}.Build()
}

func AuthorByAlias(v string) *AuthorGetRequest {
	return AuthorGetRequest_builder{Alias: &v}.Build()
}

func BookById(v []byte) *BookGetRequest {
	return BookGetRequest_builder{Id: v}.Build()
}

func FooEfById(v int64) *FooEfGetRequest {
	return FooEfGetRequest_builder{Id: &v}.Build()
}

func FooKById(v int64) *FooKGetRequest {
	return FooKGetRequest_builder{Id: &v}.Build()
}

func FooMiById(v int64) *FooMiGetRequest {
	return FooMiGetRequest_builder{Id: &v}.Build()
}

func FooMsById(v int64) *FooMsGetRequest {
	return FooMsGetRequest_builder{Id: &v}.Build()
}

func FooMsdById(v int64) *FooMsdGetRequest {
	return FooMsdGetRequest_builder{Id: &v}.Build()
}

func FooVById(v int64) *FooVGetRequest {
	return FooVGetRequest_builder{Id: &v}.Build()
}

func FooVdById(v int64) *FooVdGetRequest {
	return FooVdGetRequest_builder{Id: &v}.Build()
}

func FooVoById(v int64) *FooVoGetRequest {
	return FooVoGetRequest_builder{Id: &v}.Build()
}

func FooVonById(v int64) *FooVonGetRequest {
	return FooVonGetRequest_builder{Id: &v}.Build()
}

func FooVrById(v int64) *FooVrGetRequest {
	return FooVrGetRequest_builder{Id: &v}.Build()
}

func LikeById(v []byte) *LikeGetRequest {
	return LikeGetRequest_builder{Id: v}.Build()
}

func LikeByHolders(subject_id []byte, actor *MemberGetRequest) *LikeGetRequest {
	v := &LikeGetByHolders{}
	v.SetSubjectId(subject_id)
	v.SetActor(actor)
	return LikeGetRequest_builder{Holders: v}.Build()
}

func LoanById(v []byte) *LoanGetRequest {
	return LoanGetRequest_builder{Id: v}.Build()
}

func LockerById(v []byte) *LockerGetRequest {
	return LockerGetRequest_builder{Id: v}.Build()
}

func MemberById(v []byte) *MemberGetRequest {
	return MemberGetRequest_builder{Id: v}.Build()
}

func MemberByName(parent *MemberGetRequest, name string) *MemberGetRequest {
	v := &MemberGetByName{}
	v.SetParent(parent)
	v.SetName(name)
	return MemberGetRequest_builder{Name: v}.Build()
}

func MembershipById(v []byte) *MembershipGetRequest {
	return MembershipGetRequest_builder{Id: v}.Build()
}

func PressById(v []byte) *PressGetRequest {
	return PressGetRequest_builder{Id: v}.Build()
}

func PressBySerialNumber(book *BookGetRequest, serial_number string) *PressGetRequest {
	v := &PressGetBySerialNumber{}
	v.SetBook(book)
	v.SetSerialNumber(serial_number)
	return PressGetRequest_builder{SerialNumber: v}.Build()
}
