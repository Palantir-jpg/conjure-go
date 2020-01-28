// This file was generated by Conjure and should not be manually edited.

package api

import (
	"context"
	"io"
	"net/http"
	"strconv"

	"github.com/palantir/conjure-go-runtime/conjure-go-contract/codecs"
	"github.com/palantir/pkg/bearertoken"
	"github.com/palantir/pkg/safejson"
	"github.com/palantir/pkg/safelong"
	"github.com/palantir/pkg/uuid"
	werror "github.com/palantir/witchcraft-go-error"
	"github.com/palantir/witchcraft-go-server/rest"
	"github.com/palantir/witchcraft-go-server/witchcraft/wresource"
	"github.com/palantir/witchcraft-go-server/wrouter"
)

type TestService interface {
	Echo(ctx context.Context, cookieToken bearertoken.Token) error
	GetPathParam(ctx context.Context, authHeader bearertoken.Token, myPathParamArg string) error
	GetPathParamAlias(ctx context.Context, authHeader bearertoken.Token, myPathParamArg StringAlias) error
	QueryParamList(ctx context.Context, authHeader bearertoken.Token, myQueryParam1Arg []string) error
	PostPathParam(ctx context.Context, authHeader bearertoken.Token, myPathParam1Arg string, myPathParam2Arg bool, myBodyParamArg CustomObject, myQueryParam1Arg string, myQueryParam2Arg string, myQueryParam3Arg float64, myQueryParam4Arg *safelong.SafeLong, myQueryParam5Arg *string, myHeaderParam1Arg safelong.SafeLong, myHeaderParam2Arg *uuid.UUID) (CustomObject, error)
	Bytes(ctx context.Context) (CustomObject, error)
	GetBinary(ctx context.Context) (io.ReadCloser, error)
	PostBinary(ctx context.Context, myBytesArg io.ReadCloser) (io.ReadCloser, error)
	PutBinary(ctx context.Context, myBytesArg io.ReadCloser) error
	// An endpoint that uses go keywords
	Chan(ctx context.Context, varArg string, importArg map[string]string, typeArg string, returnArg safelong.SafeLong) error
}

// RegisterRoutesTestService registers handlers for the TestService endpoints with a witchcraft wrouter.
// This should typically be called in a witchcraft server's InitFunc.
// impl provides an implementation of each endpoint, which can assume the request parameters have been parsed
// in accordance with the Conjure specification.
func RegisterRoutesTestService(router wrouter.Router, impl TestService) error {
	handler := testServiceHandler{impl: impl}
	resource := wresource.New("testservice", router)
	if err := resource.Get("Echo", "/echo", rest.NewJSONHandler(handler.HandleEcho, rest.StatusCodeMapper, rest.ErrHandler)); err != nil {
		return werror.Wrap(err, "failed to add route", werror.SafeParam("routeName", "Echo"))
	}
	if err := resource.Get("GetPathParam", "/path/{myPathParam}", rest.NewJSONHandler(handler.HandleGetPathParam, rest.StatusCodeMapper, rest.ErrHandler)); err != nil {
		return werror.Wrap(err, "failed to add route", werror.SafeParam("routeName", "GetPathParam"))
	}
	if err := resource.Get("GetPathParamAlias", "/path/alias/{myPathParam}", rest.NewJSONHandler(handler.HandleGetPathParamAlias, rest.StatusCodeMapper, rest.ErrHandler)); err != nil {
		return werror.Wrap(err, "failed to add route", werror.SafeParam("routeName", "GetPathParamAlias"))
	}
	if err := resource.Get("QueryParamList", "/pathNew", rest.NewJSONHandler(handler.HandleQueryParamList, rest.StatusCodeMapper, rest.ErrHandler)); err != nil {
		return werror.Wrap(err, "failed to add route", werror.SafeParam("routeName", "QueryParamList"))
	}
	if err := resource.Post("PostPathParam", "/path/{myPathParam1}/{myPathParam2}", rest.NewJSONHandler(handler.HandlePostPathParam, rest.StatusCodeMapper, rest.ErrHandler)); err != nil {
		return werror.Wrap(err, "failed to add route", werror.SafeParam("routeName", "PostPathParam"))
	}
	if err := resource.Get("Bytes", "/bytes", rest.NewJSONHandler(handler.HandleBytes, rest.StatusCodeMapper, rest.ErrHandler)); err != nil {
		return werror.Wrap(err, "failed to add route", werror.SafeParam("routeName", "Bytes"))
	}
	if err := resource.Get("GetBinary", "/binary", rest.NewJSONHandler(handler.HandleGetBinary, rest.StatusCodeMapper, rest.ErrHandler)); err != nil {
		return werror.Wrap(err, "failed to add route", werror.SafeParam("routeName", "GetBinary"))
	}
	if err := resource.Post("PostBinary", "/binary", rest.NewJSONHandler(handler.HandlePostBinary, rest.StatusCodeMapper, rest.ErrHandler)); err != nil {
		return werror.Wrap(err, "failed to add route", werror.SafeParam("routeName", "PostBinary"))
	}
	if err := resource.Put("PutBinary", "/binary", rest.NewJSONHandler(handler.HandlePutBinary, rest.StatusCodeMapper, rest.ErrHandler)); err != nil {
		return werror.Wrap(err, "failed to add route", werror.SafeParam("routeName", "PutBinary"))
	}
	if err := resource.Post("Chan", "/chan/{var}", rest.NewJSONHandler(handler.HandleChan, rest.StatusCodeMapper, rest.ErrHandler)); err != nil {
		return werror.Wrap(err, "failed to add route", werror.SafeParam("routeName", "Chan"))
	}
	return nil
}

type testServiceHandler struct {
	impl TestService
}

func (t *testServiceHandler) HandleEcho(rw http.ResponseWriter, req *http.Request) error {
	authCookie, err := req.Cookie("PALANTIR_TOKEN")
	if err != nil {
		return rest.NewError(err, rest.StatusCode(http.StatusForbidden))
	}
	cookieToken := bearertoken.Token(authCookie.Value)
	return t.impl.Echo(req.Context(), cookieToken)
}

func (t *testServiceHandler) HandleGetPathParam(rw http.ResponseWriter, req *http.Request) error {
	authHeader, err := rest.ParseBearerTokenHeader(req)
	if err != nil {
		return rest.NewError(err, rest.StatusCode(http.StatusForbidden))
	}
	pathParams := wrouter.PathParams(req)
	if pathParams == nil {
		return werror.Error("path params not found on request: ensure this endpoint is registered with wrouter")
	}
	myPathParam, ok := pathParams["myPathParam"]
	if !ok {
		err := werror.Error("path param not present", werror.SafeParam("pathParamName", "myPathParam"))
		return rest.NewError(err, rest.StatusCode(http.StatusBadRequest))
	}
	return t.impl.GetPathParam(req.Context(), bearertoken.Token(authHeader), myPathParam)
}

func (t *testServiceHandler) HandleGetPathParamAlias(rw http.ResponseWriter, req *http.Request) error {
	authHeader, err := rest.ParseBearerTokenHeader(req)
	if err != nil {
		return rest.NewError(err, rest.StatusCode(http.StatusForbidden))
	}
	pathParams := wrouter.PathParams(req)
	if pathParams == nil {
		return werror.Error("path params not found on request: ensure this endpoint is registered with wrouter")
	}
	myPathParamStr, ok := pathParams["myPathParam"]
	if !ok {
		err := werror.Error("path param not present", werror.SafeParam("pathParamName", "myPathParam"))
		return rest.NewError(err, rest.StatusCode(http.StatusBadRequest))
	}
	var myPathParam StringAlias
	myPathParamQuote := strconv.Quote(myPathParamStr)
	if err := safejson.Unmarshal([]byte(myPathParamQuote), &myPathParam); err != nil {
		return werror.Wrap(err, "failed to unmarshal argument", werror.SafeParam("argName", "myPathParam"), werror.SafeParam("argType", "StringAlias"))
	}
	return t.impl.GetPathParamAlias(req.Context(), bearertoken.Token(authHeader), myPathParam)
}

func (t *testServiceHandler) HandleQueryParamList(rw http.ResponseWriter, req *http.Request) error {
	authHeader, err := rest.ParseBearerTokenHeader(req)
	if err != nil {
		return rest.NewError(err, rest.StatusCode(http.StatusForbidden))
	}
	myQueryParam1 := req.URL.Query()["myQueryParam1"]
	return t.impl.QueryParamList(req.Context(), bearertoken.Token(authHeader), myQueryParam1)
}

func (t *testServiceHandler) HandlePostPathParam(rw http.ResponseWriter, req *http.Request) error {
	authHeader, err := rest.ParseBearerTokenHeader(req)
	if err != nil {
		return rest.NewError(err, rest.StatusCode(http.StatusForbidden))
	}
	pathParams := wrouter.PathParams(req)
	if pathParams == nil {
		return werror.Error("path params not found on request: ensure this endpoint is registered with wrouter")
	}
	myPathParam1, ok := pathParams["myPathParam1"]
	if !ok {
		err := werror.Error("path param not present", werror.SafeParam("pathParamName", "myPathParam1"))
		return rest.NewError(err, rest.StatusCode(http.StatusBadRequest))
	}
	myPathParam2Str, ok := pathParams["myPathParam2"]
	if !ok {
		err := werror.Error("path param not present", werror.SafeParam("pathParamName", "myPathParam2"))
		return rest.NewError(err, rest.StatusCode(http.StatusBadRequest))
	}
	myPathParam2, err := strconv.ParseBool(myPathParam2Str)
	if err != nil {
		return err
	}
	myQueryParam1 := req.URL.Query().Get("query1")
	myQueryParam2 := req.URL.Query().Get("myQueryParam2")
	myQueryParam3, err := strconv.ParseFloat(req.URL.Query().Get("myQueryParam3"), 64)
	if err != nil {
		return err
	}
	var myQueryParam4 *safelong.SafeLong
	if myQueryParam4Str := req.URL.Query().Get("myQueryParam4"); myQueryParam4Str != "" {
		myQueryParam4Internal, err := safelong.ParseSafeLong(myQueryParam4Str)
		if err != nil {
			return err
		}
		myQueryParam4 = &myQueryParam4Internal
	}
	var myQueryParam5 *string
	if myQueryParam5Str := req.URL.Query().Get("myQueryParam5"); myQueryParam5Str != "" {
		myQueryParam5Internal := myQueryParam5Str
		myQueryParam5 = &myQueryParam5Internal
	}
	myHeaderParam1, err := safelong.ParseSafeLong(req.Header.Get("X-My-Header1-Abc"))
	if err != nil {
		return err
	}
	var myHeaderParam2 *uuid.UUID
	if myHeaderParam2Str := req.Header.Get("X-My-Header2"); myHeaderParam2Str != "" {
		myHeaderParam2Internal, err := uuid.ParseUUID(myHeaderParam2Str)
		if err != nil {
			return err
		}
		myHeaderParam2 = &myHeaderParam2Internal
	}
	var myBodyParam CustomObject
	if err := codecs.JSON.Decode(req.Body, &myBodyParam); err != nil {
		return rest.NewError(err, rest.StatusCode(http.StatusBadRequest))
	}
	respArg, err := t.impl.PostPathParam(req.Context(), bearertoken.Token(authHeader), myPathParam1, myPathParam2, myBodyParam, myQueryParam1, myQueryParam2, myQueryParam3, myQueryParam4, myQueryParam5, myHeaderParam1, myHeaderParam2)
	if err != nil {
		return err
	}
	rw.Header().Add("Content-Type", codecs.JSON.ContentType())
	return codecs.JSON.Encode(rw, respArg)
}

func (t *testServiceHandler) HandleBytes(rw http.ResponseWriter, req *http.Request) error {
	respArg, err := t.impl.Bytes(req.Context())
	if err != nil {
		return err
	}
	rw.Header().Add("Content-Type", codecs.JSON.ContentType())
	return codecs.JSON.Encode(rw, respArg)
}

func (t *testServiceHandler) HandleGetBinary(rw http.ResponseWriter, req *http.Request) error {
	respArg, err := t.impl.GetBinary(req.Context())
	if err != nil {
		return err
	}
	rw.Header().Add("Content-Type", codecs.Binary.ContentType())
	return codecs.Binary.Encode(rw, respArg)
}

func (t *testServiceHandler) HandlePostBinary(rw http.ResponseWriter, req *http.Request) error {
	myBytes := req.Body
	respArg, err := t.impl.PostBinary(req.Context(), myBytes)
	if err != nil {
		return err
	}
	rw.Header().Add("Content-Type", codecs.Binary.ContentType())
	return codecs.Binary.Encode(rw, respArg)
}

func (t *testServiceHandler) HandlePutBinary(rw http.ResponseWriter, req *http.Request) error {
	myBytes := req.Body
	return t.impl.PutBinary(req.Context(), myBytes)
}

func (t *testServiceHandler) HandleChan(rw http.ResponseWriter, req *http.Request) error {
	pathParams := wrouter.PathParams(req)
	if pathParams == nil {
		return werror.Error("path params not found on request: ensure this endpoint is registered with wrouter")
	}
	var_, ok := pathParams["var"]
	if !ok {
		err := werror.Error("path param not present", werror.SafeParam("pathParamName", "var"))
		return rest.NewError(err, rest.StatusCode(http.StatusBadRequest))
	}
	type_ := req.URL.Query().Get("type")
	return_, err := safelong.ParseSafeLong(req.Header.Get("X-My-Header2"))
	if err != nil {
		return err
	}
	var import_ map[string]string
	if err := codecs.JSON.Decode(req.Body, &import_); err != nil {
		return rest.NewError(err, rest.StatusCode(http.StatusBadRequest))
	}
	return t.impl.Chan(req.Context(), var_, import_, type_, return_)
}
