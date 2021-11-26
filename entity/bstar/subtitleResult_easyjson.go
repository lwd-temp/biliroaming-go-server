// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package bstar

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

func easyjson18ab02cbDecodeGithubComJasonKhew96BiliroamingGoServerEntityBstar(in *jlexer.Lexer, out *Subtitles) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = int64(in.Int64())
		case "key":
			out.Key = string(in.String())
		case "title":
			out.Title = string(in.String())
		case "url":
			out.URL = string(in.String())
		case "is_machine":
			out.IsMachine = bool(in.Bool())
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
func easyjson18ab02cbEncodeGithubComJasonKhew96BiliroamingGoServerEntityBstar(out *jwriter.Writer, in Subtitles) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.ID))
	}
	{
		const prefix string = ",\"key\":"
		out.RawString(prefix)
		out.String(string(in.Key))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"url\":"
		out.RawString(prefix)
		out.String(string(in.URL))
	}
	{
		const prefix string = ",\"is_machine\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsMachine))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Subtitles) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson18ab02cbEncodeGithubComJasonKhew96BiliroamingGoServerEntityBstar(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Subtitles) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson18ab02cbEncodeGithubComJasonKhew96BiliroamingGoServerEntityBstar(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Subtitles) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson18ab02cbDecodeGithubComJasonKhew96BiliroamingGoServerEntityBstar(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Subtitles) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson18ab02cbDecodeGithubComJasonKhew96BiliroamingGoServerEntityBstar(l, v)
}
func easyjson18ab02cbDecodeGithubComJasonKhew96BiliroamingGoServerEntityBstar1(in *jlexer.Lexer, out *SubtitleResultData) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "suggest_key":
			out.SuggestKey = string(in.String())
		case "subtitles":
			if in.IsNull() {
				in.Skip()
				out.Subtitles = nil
			} else {
				in.Delim('[')
				if out.Subtitles == nil {
					if !in.IsDelim(']') {
						out.Subtitles = make([]Subtitles, 0, 1)
					} else {
						out.Subtitles = []Subtitles{}
					}
				} else {
					out.Subtitles = (out.Subtitles)[:0]
				}
				for !in.IsDelim(']') {
					var v1 Subtitles
					(v1).UnmarshalEasyJSON(in)
					out.Subtitles = append(out.Subtitles, v1)
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
func easyjson18ab02cbEncodeGithubComJasonKhew96BiliroamingGoServerEntityBstar1(out *jwriter.Writer, in SubtitleResultData) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"suggest_key\":"
		out.RawString(prefix[1:])
		out.String(string(in.SuggestKey))
	}
	{
		const prefix string = ",\"subtitles\":"
		out.RawString(prefix)
		if in.Subtitles == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Subtitles {
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
func (v SubtitleResultData) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson18ab02cbEncodeGithubComJasonKhew96BiliroamingGoServerEntityBstar1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SubtitleResultData) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson18ab02cbEncodeGithubComJasonKhew96BiliroamingGoServerEntityBstar1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SubtitleResultData) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson18ab02cbDecodeGithubComJasonKhew96BiliroamingGoServerEntityBstar1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SubtitleResultData) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson18ab02cbDecodeGithubComJasonKhew96BiliroamingGoServerEntityBstar1(l, v)
}
func easyjson18ab02cbDecodeGithubComJasonKhew96BiliroamingGoServerEntityBstar2(in *jlexer.Lexer, out *SubtitleResult) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "code":
			out.Code = int(in.Int())
		case "message":
			out.Message = string(in.String())
		case "ttl":
			out.TTL = int(in.Int())
		case "data":
			(out.Data).UnmarshalEasyJSON(in)
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
func easyjson18ab02cbEncodeGithubComJasonKhew96BiliroamingGoServerEntityBstar2(out *jwriter.Writer, in SubtitleResult) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"code\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Code))
	}
	{
		const prefix string = ",\"message\":"
		out.RawString(prefix)
		out.String(string(in.Message))
	}
	{
		const prefix string = ",\"ttl\":"
		out.RawString(prefix)
		out.Int(int(in.TTL))
	}
	{
		const prefix string = ",\"data\":"
		out.RawString(prefix)
		(in.Data).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SubtitleResult) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson18ab02cbEncodeGithubComJasonKhew96BiliroamingGoServerEntityBstar2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SubtitleResult) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson18ab02cbEncodeGithubComJasonKhew96BiliroamingGoServerEntityBstar2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SubtitleResult) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson18ab02cbDecodeGithubComJasonKhew96BiliroamingGoServerEntityBstar2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SubtitleResult) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson18ab02cbDecodeGithubComJasonKhew96BiliroamingGoServerEntityBstar2(l, v)
}
