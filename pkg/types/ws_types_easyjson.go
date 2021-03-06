// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package types

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

func easyjson8df87204DecodeGithubComWebhookrelayRelayGoPkgTypes(in *jlexer.Lexer, out *EventStatus) {
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
		case "id":
			out.ID = string(in.String())
		case "status_code":
			out.StatusCode = int(in.Int())
		case "retries":
			out.Retries = int(in.Int())
		case "message":
			out.Message = string(in.String())
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
func easyjson8df87204EncodeGithubComWebhookrelayRelayGoPkgTypes(out *jwriter.Writer, in EventStatus) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"status_code\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.StatusCode))
	}
	{
		const prefix string = ",\"retries\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Retries))
	}
	{
		const prefix string = ",\"message\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Message))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EventStatus) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8df87204EncodeGithubComWebhookrelayRelayGoPkgTypes(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventStatus) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8df87204EncodeGithubComWebhookrelayRelayGoPkgTypes(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EventStatus) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8df87204DecodeGithubComWebhookrelayRelayGoPkgTypes(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventStatus) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8df87204DecodeGithubComWebhookrelayRelayGoPkgTypes(l, v)
}
func easyjson8df87204DecodeGithubComWebhookrelayRelayGoPkgTypes1(in *jlexer.Lexer, out *EventMeta) {
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
		case "id":
			out.ID = string(in.String())
		case "bucked_id":
			out.BucketID = string(in.String())
		case "bucket_name":
			out.BucketName = string(in.String())
		case "input_id":
			out.InputID = string(in.String())
		case "input_name":
			out.InputName = string(in.String())
		case "output_name":
			out.OutputName = string(in.String())
		case "output_destination":
			out.OutputDestination = string(in.String())
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
func easyjson8df87204EncodeGithubComWebhookrelayRelayGoPkgTypes1(out *jwriter.Writer, in EventMeta) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"bucked_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.BucketID))
	}
	{
		const prefix string = ",\"bucket_name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.BucketName))
	}
	{
		const prefix string = ",\"input_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.InputID))
	}
	{
		const prefix string = ",\"input_name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.InputName))
	}
	{
		const prefix string = ",\"output_name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.OutputName))
	}
	{
		const prefix string = ",\"output_destination\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.OutputDestination))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EventMeta) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8df87204EncodeGithubComWebhookrelayRelayGoPkgTypes1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventMeta) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8df87204EncodeGithubComWebhookrelayRelayGoPkgTypes1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EventMeta) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8df87204DecodeGithubComWebhookrelayRelayGoPkgTypes1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventMeta) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8df87204DecodeGithubComWebhookrelayRelayGoPkgTypes1(l, v)
}
func easyjson8df87204DecodeGithubComWebhookrelayRelayGoPkgTypes2(in *jlexer.Lexer, out *Event) {
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
		case "meta":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Meta).UnmarshalJSON(data))
			}
		case "headers":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Headers = make(map[string][]string)
				} else {
					out.Headers = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v1 []string
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						in.Delim('[')
						if v1 == nil {
							if !in.IsDelim(']') {
								v1 = make([]string, 0, 4)
							} else {
								v1 = []string{}
							}
						} else {
							v1 = (v1)[:0]
						}
						for !in.IsDelim(']') {
							var v2 string
							v2 = string(in.String())
							v1 = append(v1, v2)
							in.WantComma()
						}
						in.Delim(']')
					}
					(out.Headers)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
			}
		case "query":
			out.RawQuery = string(in.String())
		case "body":
			out.Body = string(in.String())
		case "method":
			out.Method = string(in.String())
		case "status":
			out.Status = string(in.String())
		case "message":
			out.Message = string(in.String())
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
func easyjson8df87204EncodeGithubComWebhookrelayRelayGoPkgTypes2(out *jwriter.Writer, in Event) {
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
		const prefix string = ",\"meta\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Meta).MarshalJSON())
	}
	{
		const prefix string = ",\"headers\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Headers == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v3First := true
			for v3Name, v3Value := range in.Headers {
				if v3First {
					v3First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v3Name))
				out.RawByte(':')
				if v3Value == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
					out.RawString("null")
				} else {
					out.RawByte('[')
					for v4, v5 := range v3Value {
						if v4 > 0 {
							out.RawByte(',')
						}
						out.String(string(v5))
					}
					out.RawByte(']')
				}
			}
			out.RawByte('}')
		}
	}
	{
		const prefix string = ",\"query\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.RawQuery))
	}
	{
		const prefix string = ",\"body\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Body))
	}
	{
		const prefix string = ",\"method\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Method))
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
	{
		const prefix string = ",\"message\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Message))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Event) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8df87204EncodeGithubComWebhookrelayRelayGoPkgTypes2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Event) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8df87204EncodeGithubComWebhookrelayRelayGoPkgTypes2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Event) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8df87204DecodeGithubComWebhookrelayRelayGoPkgTypes2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Event) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8df87204DecodeGithubComWebhookrelayRelayGoPkgTypes2(l, v)
}
func easyjson8df87204DecodeGithubComWebhookrelayRelayGoPkgTypes3(in *jlexer.Lexer, out *ActionRequest) {
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
		case "action":
			out.Action = string(in.String())
		case "key":
			out.Key = string(in.String())
		case "secret":
			out.Secret = string(in.String())
		case "buckets":
			if in.IsNull() {
				in.Skip()
				out.Buckets = nil
			} else {
				in.Delim('[')
				if out.Buckets == nil {
					if !in.IsDelim(']') {
						out.Buckets = make([]string, 0, 4)
					} else {
						out.Buckets = []string{}
					}
				} else {
					out.Buckets = (out.Buckets)[:0]
				}
				for !in.IsDelim(']') {
					var v6 string
					v6 = string(in.String())
					out.Buckets = append(out.Buckets, v6)
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
func easyjson8df87204EncodeGithubComWebhookrelayRelayGoPkgTypes3(out *jwriter.Writer, in ActionRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"action\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Action))
	}
	{
		const prefix string = ",\"key\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Key))
	}
	{
		const prefix string = ",\"secret\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Secret))
	}
	{
		const prefix string = ",\"buckets\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Buckets == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v7, v8 := range in.Buckets {
				if v7 > 0 {
					out.RawByte(',')
				}
				out.String(string(v8))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ActionRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8df87204EncodeGithubComWebhookrelayRelayGoPkgTypes3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ActionRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8df87204EncodeGithubComWebhookrelayRelayGoPkgTypes3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ActionRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8df87204DecodeGithubComWebhookrelayRelayGoPkgTypes3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ActionRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8df87204DecodeGithubComWebhookrelayRelayGoPkgTypes3(l, v)
}
