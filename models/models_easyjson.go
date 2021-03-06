// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

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

func easyjsonD2b7633eDecodeHlModels(in *jlexer.Lexer, out *UserVisitsSl) {
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
		case "visits":
			if in.IsNull() {
				in.Skip()
				out.Visits = nil
			} else {
				in.Delim('[')
				if out.Visits == nil {
					if !in.IsDelim(']') {
						out.Visits = make([]UserVisit, 0, 2)
					} else {
						out.Visits = []UserVisit{}
					}
				} else {
					out.Visits = (out.Visits)[:0]
				}
				for !in.IsDelim(']') {
					var v1 UserVisit
					(v1).UnmarshalEasyJSON(in)
					out.Visits = append(out.Visits, v1)
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
func easyjsonD2b7633eEncodeHlModels(out *jwriter.Writer, in UserVisitsSl) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"visits\":")
	if in.Visits == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in.Visits {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserVisitsSl) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeHlModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserVisitsSl) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeHlModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserVisitsSl) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeHlModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserVisitsSl) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeHlModels(l, v)
}
func easyjsonD2b7633eDecodeHlModels1(in *jlexer.Lexer, out *UserVisit) {
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
		case "mark":
			out.Mark = int(in.Int())
		case "visited_at":
			out.VisitedAt = int(in.Int())
		case "place":
			out.Place = string(in.String())
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
func easyjsonD2b7633eEncodeHlModels1(out *jwriter.Writer, in UserVisit) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"mark\":")
	out.Int(int(in.Mark))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"visited_at\":")
	out.Int(int(in.VisitedAt))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"place\":")
	out.String(string(in.Place))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserVisit) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeHlModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserVisit) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeHlModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserVisit) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeHlModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserVisit) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeHlModels1(l, v)
}
