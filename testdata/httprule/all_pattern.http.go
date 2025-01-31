// Code generated by protoc-gen-gohttp. DO NOT EDIT.
// source: httprule/all_pattern.proto

package httprulepb

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AllPatternHTTPConverter has a function to convert AllPatternServer interface to http.HandlerFunc.
type AllPatternHTTPConverter struct {
	srv AllPatternServer
}

// NewAllPatternHTTPConverter returns AllPatternHTTPConverter.
func NewAllPatternHTTPConverter(srv AllPatternServer) *AllPatternHTTPConverter {
	return &AllPatternHTTPConverter{
		srv: srv,
	}
}

// AllPattern returns AllPatternServer interface's AllPattern converted to http.HandlerFunc.
func (h *AllPatternHTTPConverter) AllPattern(cb func(ctx context.Context, w http.ResponseWriter, r *http.Request, arg, ret proto.Message, err error)) http.HandlerFunc {
	if cb == nil {
		cb = func(ctx context.Context, w http.ResponseWriter, r *http.Request, arg, ret proto.Message, err error) {
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				p := status.New(codes.Unknown, err.Error()).Proto()
				switch r.Header.Get("Content-Type") {
				case "application/protobuf", "application/x-protobuf":
					buf, err := proto.Marshal(p)
					if err != nil {
						return
					}
					if _, err := io.Copy(w, bytes.NewBuffer(buf)); err != nil {
						return
					}
				case "application/json":
					if err := json.NewEncoder(w).Encode(p); err != nil {
						return
					}
				default:
				}
			}
		}
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		arg := &AllPatternRequest{}
		contentType := r.Header.Get("Content-Type")
		if r.Method != http.MethodGet {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				cb(ctx, w, r, nil, nil, err)
				return
			}

			switch contentType {
			case "application/protobuf", "application/x-protobuf":
				if err := proto.Unmarshal(body, arg); err != nil {
					cb(ctx, w, r, nil, nil, err)
					return
				}
			case "application/json":
				if err := jsonpb.Unmarshal(bytes.NewBuffer(body), arg); err != nil {
					cb(ctx, w, r, nil, nil, err)
					return
				}
			default:
				w.WriteHeader(http.StatusUnsupportedMediaType)
				_, err := fmt.Fprintf(w, "Unsupported Content-Type: %s", contentType)
				cb(ctx, w, r, nil, nil, err)
				return
			}
		}

		ret, err := h.srv.AllPattern(ctx, arg)
		if err != nil {
			cb(ctx, w, r, arg, nil, err)
			return
		}

		accepts := strings.Split(r.Header.Get("Accept"), ",")
		accept := accepts[0]
		if accept == "*/*" || accept == "" {
			if contentType != "" {
				accept = contentType
			} else {
				accept = "application/json"
			}
		}

		w.Header().Set("Content-Type", accept)

		switch accept {
		case "application/protobuf", "application/x-protobuf":
			buf, err := proto.Marshal(ret)
			if err != nil {
				cb(ctx, w, r, arg, ret, err)
				return
			}
			if _, err := io.Copy(w, bytes.NewBuffer(buf)); err != nil {
				cb(ctx, w, r, arg, ret, err)
				return
			}
		case "application/json":
			m := jsonpb.Marshaler{
				EnumsAsInts:  true,
				EmitDefaults: true,
			}
			if err := m.Marshal(w, ret); err != nil {
				cb(ctx, w, r, arg, ret, err)
				return
			}
		default:
			w.WriteHeader(http.StatusUnsupportedMediaType)
			_, err := fmt.Fprintf(w, "Unsupported Accept: %s", accept)
			cb(ctx, w, r, arg, ret, err)
			return
		}
		cb(ctx, w, r, arg, ret, nil)
	})
}

// AllPatternWithName returns Service name, Method name and AllPatternServer interface's AllPattern converted to http.HandlerFunc.
func (h *AllPatternHTTPConverter) AllPatternWithName(cb func(ctx context.Context, w http.ResponseWriter, r *http.Request, arg, ret proto.Message, err error)) (string, string, http.HandlerFunc) {
	return "AllPattern", "AllPattern", h.AllPattern(cb)
}

func (h *AllPatternHTTPConverter) AllPatternHTTPRule(cb func(ctx context.Context, w http.ResponseWriter, r *http.Request, arg, ret proto.Message, err error)) (string, string, http.HandlerFunc) {
	if cb == nil {
		cb = func(ctx context.Context, w http.ResponseWriter, r *http.Request, arg, ret proto.Message, err error) {
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				p := status.New(codes.Unknown, err.Error()).Proto()
				switch r.Header.Get("Content-Type") {
				case "application/protobuf", "application/x-protobuf":
					buf, err := proto.Marshal(p)
					if err != nil {
						return
					}
					if _, err := io.Copy(w, bytes.NewBuffer(buf)); err != nil {
						return
					}
				case "application/json":
					if err := json.NewEncoder(w).Encode(p); err != nil {
						return
					}
				default:
				}
			}
		}
	}
	return http.MethodGet, "/all/pattern", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		arg := &AllPatternRequest{}
		contentType := r.Header.Get("Content-Type")
		if r.Method == http.MethodGet {
			if v := r.URL.Query().Get("double"); v != "" {
				d, err := strconv.ParseFloat(v, 64)
				if err != nil {
					cb(ctx, w, r, nil, nil, err)
					return
				}
				arg.Double = d
			}
			if v := r.URL.Query().Get("float"); v != "" {
				f, err := strconv.ParseFloat(v, 32)
				if err != nil {
					cb(ctx, w, r, nil, nil, err)
					return
				}
				arg.Float = float32(f)
			}
			if v := r.URL.Query().Get("int32"); v != "" {
				i32, err := strconv.ParseInt(v, 10, 32)
				if err != nil {
					cb(ctx, w, r, nil, nil, err)
					return
				}
				arg.Int32 = int32(i32)
			}
			if v := r.URL.Query().Get("int64"); v != "" {
				i64, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					cb(ctx, w, r, nil, nil, err)
					return
				}
				arg.Int64 = i64
			}
			if v := r.URL.Query().Get("uint32"); v != "" {
				ui32, err := strconv.ParseUint(v, 10, 32)
				if err != nil {
					cb(ctx, w, r, nil, nil, err)
					return
				}
				arg.Uint32 = uint32(ui32)
			}
			if v := r.URL.Query().Get("uint64"); v != "" {
				ui64, err := strconv.ParseUint(v, 10, 64)
				if err != nil {
					cb(ctx, w, r, nil, nil, err)
					return
				}
				arg.Uint64 = uint64(ui64)
			}
			if v := r.URL.Query().Get("fixed32"); v != "" {
				f32, err := strconv.ParseUint(v, 10, 32)
				if err != nil {
					cb(ctx, w, r, nil, nil, err)
					return
				}
				arg.Fixed32 = uint32(f32)
			}
			if v := r.URL.Query().Get("fixed64"); v != "" {
				f64, err := strconv.ParseUint(v, 10, 64)
				if err != nil {
					cb(ctx, w, r, nil, nil, err)
					return
				}
				arg.Fixed64 = uint64(f64)
			}
			if v := r.URL.Query().Get("sfixed32"); v != "" {
				sf32, err := strconv.ParseInt(v, 10, 32)
				if err != nil {
					cb(ctx, w, r, nil, nil, err)
					return
				}
				arg.Sfixed32 = int32(sf32)
			}
			if v := r.URL.Query().Get("sfixed64"); v != "" {
				sf64, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					cb(ctx, w, r, nil, nil, err)
					return
				}
				arg.Sfixed64 = int64(sf64)
			}
			if v := r.URL.Query().Get("bool"); v != "" {
				b, err := strconv.ParseBool(v)
				if err != nil {
					cb(ctx, w, r, nil, nil, err)
					return
				}
				arg.Bool = b
			}
			if v := r.URL.Query().Get("string"); v != "" {
				arg.String_ = v
			}
			if v := r.URL.Query().Get("bytes"); v != "" {
				b, err := base64.StdEncoding.DecodeString(v)
				if err != nil {
					cb(ctx, w, r, nil, nil, err)
					return
				}
				arg.Bytes = b
			}
			if repeated := r.URL.Query()["repeated_double"]; len(repeated) != 0 {
				arr := make([]float64, 0, len(repeated))
				for _, v := range repeated {
					d, err := strconv.ParseFloat(v, 64)
					if err != nil {
						cb(ctx, w, r, nil, nil, err)
						return
					}
					arr = append(arr, d)
				}
				arg.RepeatedDouble = arr
			}
			if repeated := r.URL.Query()["repeated_float"]; len(repeated) != 0 {
				arr := make([]float32, 0, len(repeated))
				for _, v := range repeated {
					f, err := strconv.ParseFloat(v, 32)
					if err != nil {
						cb(ctx, w, r, nil, nil, err)
						return
					}
					arr = append(arr, float32(f))
				}
				arg.RepeatedFloat = arr
			}
			if repeated := r.URL.Query()["repeated_int32"]; len(repeated) != 0 {
				arr := make([]int32, 0, len(repeated))
				for _, v := range repeated {
					i32, err := strconv.ParseFloat(v, 32)
					if err != nil {
						cb(ctx, w, r, nil, nil, err)
						return
					}
					arr = append(arr, int32(i32))
				}
				arg.RepeatedInt32 = arr
			}
			if repeated := r.URL.Query()["repeated_int64"]; len(repeated) != 0 {
				arr := make([]int64, 0, len(repeated))
				for _, v := range repeated {
					i64, err := strconv.ParseFloat(v, 64)
					if err != nil {
						cb(ctx, w, r, nil, nil, err)
						return
					}
					arr = append(arr, int64(i64))
				}
				arg.RepeatedInt64 = arr
			}
			if repeated := r.URL.Query()["repeated_uint32"]; len(repeated) != 0 {
				arr := make([]uint32, 0, len(repeated))
				for _, v := range repeated {
					ui32, err := strconv.ParseFloat(v, 32)
					if err != nil {
						cb(ctx, w, r, nil, nil, err)
						return
					}
					arr = append(arr, uint32(ui32))
				}
				arg.RepeatedUint32 = arr
			}
			if repeated := r.URL.Query()["repeated_uint64"]; len(repeated) != 0 {
				arr := make([]uint64, 0, len(repeated))
				for _, v := range repeated {
					ui64, err := strconv.ParseFloat(v, 64)
					if err != nil {
						cb(ctx, w, r, nil, nil, err)
						return
					}
					arr = append(arr, uint64(ui64))
				}
				arg.RepeatedUint64 = arr
			}
			if repeated := r.URL.Query()["repeated_fixed32"]; len(repeated) != 0 {
				arr := make([]uint32, 0, len(repeated))
				for _, v := range repeated {
					f32, err := strconv.ParseFloat(v, 32)
					if err != nil {
						cb(ctx, w, r, nil, nil, err)
						return
					}
					arr = append(arr, uint32(f32))
				}
				arg.RepeatedFixed32 = arr
			}
			if repeated := r.URL.Query()["repeated_fixed64"]; len(repeated) != 0 {
				arr := make([]uint64, 0, len(repeated))
				for _, v := range repeated {
					f64, err := strconv.ParseFloat(v, 64)
					if err != nil {
						cb(ctx, w, r, nil, nil, err)
						return
					}
					arr = append(arr, uint64(f64))
				}
				arg.RepeatedFixed64 = arr
			}
			if repeated := r.URL.Query()["repeated_sfixed32"]; len(repeated) != 0 {
				arr := make([]int32, 0, len(repeated))
				for _, v := range repeated {
					sf32, err := strconv.ParseFloat(v, 32)
					if err != nil {
						cb(ctx, w, r, nil, nil, err)
						return
					}
					arr = append(arr, int32(sf32))
				}
				arg.RepeatedSfixed32 = arr
			}
			if repeated := r.URL.Query()["repeated_sfixed64"]; len(repeated) != 0 {
				arr := make([]int64, 0, len(repeated))
				for _, v := range repeated {
					sf64, err := strconv.ParseFloat(v, 64)
					if err != nil {
						cb(ctx, w, r, nil, nil, err)
						return
					}
					arr = append(arr, int64(sf64))
				}
				arg.RepeatedSfixed64 = arr
			}
			if repeated := r.URL.Query()["repeated_bool"]; len(repeated) != 0 {
				arr := make([]bool, 0, len(repeated))
				for _, v := range repeated {
					b, err := strconv.ParseBool(v)
					if err != nil {
						cb(ctx, w, r, nil, nil, err)
						return
					}
					arr = append(arr, b)
				}
				arg.RepeatedBool = arr
			}
			if repeated := r.URL.Query()["repeated_string"]; len(repeated) != 0 {
				arr := make([]string, 0, len(repeated))
				for _, v := range repeated {
					arr = append(arr, v)
				}
				arg.RepeatedString = arr
			}
			if repeated := r.URL.Query()["repeated_bytes"]; len(repeated) != 0 {
				arr := make([][]byte, 0, len(repeated))
				for _, v := range repeated {
					b, err := base64.StdEncoding.DecodeString(v)
					if err != nil {
						cb(ctx, w, r, nil, nil, err)
						return
					}
					arr = append(arr, b)
				}
				arg.RepeatedBytes = arr
			}
		}

		ret, err := h.srv.AllPattern(ctx, arg)
		if err != nil {
			cb(ctx, w, r, arg, nil, err)
			return
		}

		accepts := strings.Split(r.Header.Get("Accept"), ",")
		accept := accepts[0]
		if accept == "*/*" || accept == "" {
			if contentType != "" {
				accept = contentType
			} else {
				accept = "application/json"
			}
		}

		w.Header().Set("Content-Type", accept)

		switch accept {
		case "application/protobuf", "application/x-protobuf":
			buf, err := proto.Marshal(ret)
			if err != nil {
				cb(ctx, w, r, arg, ret, err)
				return
			}
			if _, err := io.Copy(w, bytes.NewBuffer(buf)); err != nil {
				cb(ctx, w, r, arg, ret, err)
				return
			}
		case "application/json":
			m := jsonpb.Marshaler{
				EnumsAsInts:  true,
				EmitDefaults: true,
			}
			if err := m.Marshal(w, ret); err != nil {
				cb(ctx, w, r, arg, ret, err)
				return
			}
		default:
			w.WriteHeader(http.StatusUnsupportedMediaType)
			_, err := fmt.Fprintf(w, "Unsupported Accept: %s", accept)
			cb(ctx, w, r, arg, ret, err)
			return
		}
		cb(ctx, w, r, arg, ret, nil)
	})
}
