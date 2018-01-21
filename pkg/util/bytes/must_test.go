package bytes

import (
	"bytes"
	"errors"
	"reflect"
	"testing"
)

func TestMust(t *testing.T) {
	data := []byte("Arbitrary data")
	var err error

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Must(%x, %v) panicked with %v", data, err, r)
		}
	}()

	got := Must(data, err)

	switch {
	case reflect.ValueOf(got).Pointer() != reflect.ValueOf(data).Pointer():
		t.Errorf("Must(%x, %v) points to different array than parameter", data, err)
	case !bytes.Equal(got, data):
		t.Errorf("Must(%x, %v) == %x, want %x", data, err, got, data)
	}
}

func TestMust_Panic(t *testing.T) {
	data := []byte("More arbitrary data")
	err := errors.New("Arbitrary error")

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Must(%x, %v) did not panic", data, err)
		}
	}()

	var _ = Must(data, err)
}
