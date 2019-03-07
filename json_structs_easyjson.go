// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package main

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson6a93d021DecodeTest(in *jlexer.Lexer, out *UsrRequest) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "login":
			out.Login = string(in.String())
		case "password":
			out.Password = string(in.String())
		case "email":
			out.Email = string(in.String())
		case "name":
			out.Name = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6a93d021EncodeTest(out *jwriter.Writer, in UsrRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"login\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Login))
	}
	{
		const prefix string = ",\"password\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Password))
	}
	{
		const prefix string = ",\"email\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UsrRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6a93d021EncodeTest(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UsrRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6a93d021EncodeTest(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UsrRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6a93d021DecodeTest(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UsrRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6a93d021DecodeTest(l, v)
}
func easyjson6a93d021DecodeTest1(in *jlexer.Lexer, out *UsersPayload) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "users":
			if in.IsNull() {
				in.Skip()
				out.Users = nil
			} else {
				in.Delim('[')
				if out.Users == nil {
					if !in.IsDelim(']') {
						out.Users = make([]User, 0, 1)
					} else {
						out.Users = []User{}
					}
				} else {
					out.Users = (out.Users)[:0]
				}
				for !in.IsDelim(']') {
					var v1 User
					(v1).UnmarshalEasyJSON(in)
					out.Users = append(out.Users, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6a93d021EncodeTest1(out *jwriter.Writer, in UsersPayload) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"users\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Users == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Users {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UsersPayload) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6a93d021EncodeTest1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UsersPayload) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6a93d021EncodeTest1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UsersPayload) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6a93d021DecodeTest1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UsersPayload) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6a93d021DecodeTest1(l, v)
}
func easyjson6a93d021DecodeTest2(in *jlexer.Lexer, out *UserDataPayload) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "login":
			out.Login = string(in.String())
		case "email":
			out.Email = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "avatar":
			out.AvatarPath = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6a93d021EncodeTest2(out *jwriter.Writer, in UserDataPayload) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Login != "" {
		const prefix string = ",\"login\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Login))
	}
	if in.Email != "" {
		const prefix string = ",\"email\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Email))
	}
	if in.Name != "" {
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	if in.AvatarPath != "" {
		const prefix string = ",\"avatar\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.AvatarPath))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserDataPayload) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6a93d021EncodeTest2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserDataPayload) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6a93d021EncodeTest2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserDataPayload) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6a93d021DecodeTest2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserDataPayload) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6a93d021DecodeTest2(l, v)
}
func easyjson6a93d021DecodeTest3(in *jlexer.Lexer, out *User) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6a93d021EncodeTest3(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v User) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6a93d021EncodeTest3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v User) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6a93d021EncodeTest3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6a93d021DecodeTest3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6a93d021DecodeTest3(l, v)
}
func easyjson6a93d021DecodeTest4(in *jlexer.Lexer, out *Response) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "type":
			out.Type = string(in.String())
		case "status":
			out.Status = string(in.String())
		case "payload":
			if m, ok := out.Payload.(easyjson.Unmarshaler); ok {
				m.UnmarshalEasyJSON(in)
			} else if m, ok := out.Payload.(json.Unmarshaler); ok {
				_ = m.UnmarshalJSON(in.Raw())
			} else {
				out.Payload = in.Interface()
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6a93d021EncodeTest4(out *jwriter.Writer, in Response) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"status\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Status))
	}
	if in.Payload != nil {
		const prefix string = ",\"payload\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if m, ok := in.Payload.(easyjson.Marshaler); ok {
			m.MarshalEasyJSON(out)
		} else if m, ok := in.Payload.(json.Marshaler); ok {
			out.Raw(m.MarshalJSON())
		} else {
			out.Raw(json.Marshal(in.Payload))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Response) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6a93d021EncodeTest4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Response) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6a93d021EncodeTest4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Response) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6a93d021DecodeTest4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Response) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6a93d021DecodeTest4(l, v)
}
func easyjson6a93d021DecodeTest5(in *jlexer.Lexer, out *LeaderboardRequest) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "count":
			out.Count = int(in.Int())
		case "page":
			out.Page = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6a93d021EncodeTest5(out *jwriter.Writer, in LeaderboardRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"count\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Count))
	}
	{
		const prefix string = ",\"page\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Page))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v LeaderboardRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6a93d021EncodeTest5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v LeaderboardRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6a93d021EncodeTest5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *LeaderboardRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6a93d021DecodeTest5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *LeaderboardRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6a93d021DecodeTest5(l, v)
}
func easyjson6a93d021DecodeTest6(in *jlexer.Lexer, out *ErrorPayload) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "message":
			out.Message = string(in.String())
		case "field":
			out.Field = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6a93d021EncodeTest6(out *jwriter.Writer, in ErrorPayload) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Message != "" {
		const prefix string = ",\"message\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Message))
	}
	if in.Field != "" {
		const prefix string = ",\"field\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Field))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ErrorPayload) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6a93d021EncodeTest6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ErrorPayload) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6a93d021EncodeTest6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ErrorPayload) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6a93d021DecodeTest6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ErrorPayload) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6a93d021DecodeTest6(l, v)
}
