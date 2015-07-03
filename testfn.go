package spec2test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// need to get some kind of template test for each CRUD action
// from the user somehow

// let's start with the entire test func in a text/template

/* Endpoint test func TEMPLATE design

specific inputs:
	- url path
	- http verb
	- go type
	- one|many
	- content-type
	- additional headers

  generic inputs:
	- net/http handler/mux
*/

var (
	server          *httptest.Server
	urlPath         string
	action          Action
	expectedRespObj interface{} // unmarshalled response data
	typ             reflect.Type
	resourceType    ResourceType
	header          http.Header

	contentType = "application/json"
)

func EndpointTestTmp(t *testing.T) {
	urlStr := server.URL + urlPath
	action := actionHTTPMethods[action]

	var body bytes.Buffer

	if action.hasBody {
		// TODO: use real content-type
		err := Encode(contentType, expectedRespObj, &body)
		if err != nil {
			t.Fatal(err)
		}
	}

	httpClient := http.DefaultClient
	req, err := http.NewRequest(string(action.httpMethod), urlStr, &body)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected 200 response code, got %d\n", resp.StatusCode)
	}

	respContentType := resp.Header.Get("Content-Type")
	respObj := NewEmptyClone(expectedRespObj)

	err = Decode(respContentType, resp.Body, respObj)
	defer resp.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	if !ValueDeepEqual(respObj, expectedRespObj) {
		t.Fatalf("Expected response obj == %#v, got %#v\n", expectedRespObj, respObj)
	}
}

func NewEmptyClone(obj interface{}) interface{} {
	val := reflect.ValueOf(obj)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return reflect.New(val.Type()).Interface()
}

func ValueDeepEqual(v1, v2 interface{}) bool {
	v1val := dereference(v1)
	v2val := dereference(v2)

	return reflect.DeepEqual(v1val, v2val)
}

func dereference(v interface{}) reflect.Value {
	reflectv := reflect.ValueOf(v)
	if reflectv.Kind() == reflect.Ptr {
		reflectv = reflectv.Elem()
	}

	return reflectv
}
