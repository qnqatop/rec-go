package db

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestJSONNumber_UnmarshalJSON(t *testing.T) {
	t.Run("empty inline", func(t *testing.T) {
		var (
			args = `""`

			wantTarget JSONNumber
			target     JSONNumber
		)
		const wantErr = false
		if err := json.Unmarshal([]byte(args), &target); (err != nil) != wantErr {
			t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, wantErr)
		}
		if !reflect.DeepEqual(target, wantTarget) {
			t.Errorf("UnmarshalJSON() got = %v, want %v", target, wantTarget)
		}
	})
	t.Run("empty number inline", func(t *testing.T) {
		var (
			args = `123`

			wantTarget = JSONNumber("123")
			target     JSONNumber
		)
		const wantErr = false
		if err := json.Unmarshal([]byte(args), &target); (err != nil) != wantErr {
			t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, wantErr)
		}
		if !reflect.DeepEqual(target, wantTarget) {
			t.Errorf("UnmarshalJSON() got = %v, want %v", target, wantTarget)
		}
	})
	t.Run("empty text inline", func(t *testing.T) {
		var (
			args = `"123"`

			wantTarget = JSONNumber("123")
			target     JSONNumber
		)
		const wantErr = false
		if err := json.Unmarshal([]byte(args), &target); (err != nil) != wantErr {
			t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, wantErr)
		}
		if !reflect.DeepEqual(target, wantTarget) {
			t.Errorf("UnmarshalJSON() got = %v, want %v", target, wantTarget)
		}
	})
	t.Run("empty not valid text inline", func(t *testing.T) {
		var (
			args = `"abc123"`

			target JSONNumber
		)
		const wantErr = true
		if err := json.Unmarshal([]byte(args), &target); (err != nil) != wantErr {
			t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, wantErr)
		}
	})
	t.Run("null inline", func(t *testing.T) {
		var (
			args = `null`

			wantTarget = JSONNumber("")
			target     JSONNumber
		)
		const wantErr = true
		if err := json.Unmarshal([]byte(args), &target); (err != nil) != wantErr {
			t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, wantErr)
		}
		if !reflect.DeepEqual(target, wantTarget) {
			t.Errorf("UnmarshalJSON() got = %v, want %v", target, wantTarget)
		}
	})
	t.Run("struct field empty", func(t *testing.T) {
		type T struct {
			X JSONNumber
		}
		var (
			args = `{"X":""}`

			wantTarget = T{X: ""}
			target     T
		)
		const wantErr = false
		if err := json.Unmarshal([]byte(args), &target); (err != nil) != wantErr {
			t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, wantErr)
		}
		if !reflect.DeepEqual(target, wantTarget) {
			t.Errorf("UnmarshalJSON() got = %v, want %v", target, wantTarget)
		}
	})
	t.Run("struct field null", func(t *testing.T) {
		type T struct {
			X JSONNumber
		}
		var (
			args = `{"X":null}`

			wantTarget = T{X: ""}
			target     T
		)
		const wantErr = true
		if err := json.Unmarshal([]byte(args), &target); (err != nil) != wantErr {
			t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, wantErr)
		}
		if !reflect.DeepEqual(target, wantTarget) {
			t.Errorf("UnmarshalJSON() got = %v, want %v", target, wantTarget)
		}
	})
	t.Run("struct field text", func(t *testing.T) {
		type T struct {
			X JSONNumber
		}
		var (
			args = `{"X":"123"}`

			wantTarget = T{X: "123"}
			target     T
		)
		const wantErr = false
		if err := json.Unmarshal([]byte(args), &target); (err != nil) != wantErr {
			t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, wantErr)
		}
		if !reflect.DeepEqual(target, wantTarget) {
			t.Errorf("UnmarshalJSON() got = %v, want %v", target, wantTarget)
		}
	})
	t.Run("struct field number", func(t *testing.T) {
		type T struct {
			X JSONNumber
		}
		var (
			args = `{"X":123}`

			wantTarget = T{X: "123"}
			target     T
		)
		const wantErr = false
		if err := json.Unmarshal([]byte(args), &target); (err != nil) != wantErr {
			t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, wantErr)
		}
		if !reflect.DeepEqual(target, wantTarget) {
			t.Errorf("UnmarshalJSON() got = %v, want %v", target, wantTarget)
		}
	})
	t.Run("struct field invalid number as text", func(t *testing.T) {
		type T struct {
			X JSONNumber
		}
		var (
			args = `{"X":"abc123"}`

			target T
		)
		const wantErr = true
		if err := json.Unmarshal([]byte(args), &target); (err != nil) != wantErr {
			t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, wantErr)
		}
	})
	t.Run("struct field pointer null", func(t *testing.T) {
		type T struct {
			X *JSONNumber
		}
		var (
			args = `{"X":null}`

			wantTarget = T{X: nil}
			target     T
		)
		const wantErr = false
		if err := json.Unmarshal([]byte(args), &target); (err != nil) != wantErr {
			t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, wantErr)
		}
		if !reflect.DeepEqual(target, wantTarget) {
			t.Errorf("UnmarshalJSON() got = %v, want %v", target, wantTarget)
		}
	})
	t.Run("struct field pointer empty text", func(t *testing.T) {
		type T struct {
			X *JSONNumber
		}
		var (
			args = `{"X":""}`

			wantTarget = T{X: np("")}
			target     T
		)
		const wantErr = false
		if err := json.Unmarshal([]byte(args), &target); (err != nil) != wantErr {
			t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, wantErr)
		}
		if !reflect.DeepEqual(target, wantTarget) {
			t.Errorf("UnmarshalJSON() got = %v, want %v", target, wantTarget)
		}
	})
	t.Run("struct field pointer text", func(t *testing.T) {
		type T struct {
			X *JSONNumber
		}
		var (
			args = `{"X":"123"}`

			wantTarget = T{X: np("123")}
			target     T
		)
		const wantErr = false
		if err := json.Unmarshal([]byte(args), &target); (err != nil) != wantErr {
			t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, wantErr)
		}
		if !reflect.DeepEqual(target, wantTarget) {
			t.Errorf("UnmarshalJSON() got = %v, want %v", target, wantTarget)
		}
	})
	t.Run("struct field pointer invalid text", func(t *testing.T) {
		type T struct {
			X *JSONNumber
		}
		var (
			args = `{"X":"abc123"}`

			target T
		)
		const wantErr = true
		if err := json.Unmarshal([]byte(args), &target); (err != nil) != wantErr {
			t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, wantErr)
		}
	})
	t.Run("struct field pointer number", func(t *testing.T) {
		type T struct {
			X *JSONNumber
		}
		var (
			args = `{"X":123}`

			wantTarget = T{X: np("123")}
			target     T
		)
		const wantErr = false
		if err := json.Unmarshal([]byte(args), &target); (err != nil) != wantErr {
			t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, wantErr)
		}
		if !reflect.DeepEqual(target, wantTarget) {
			t.Errorf("UnmarshalJSON() got = %v, want %v", target, wantTarget)
		}
	})
	t.Run("struct field array", func(t *testing.T) {
		type T struct {
			X []JSONNumber
		}
		var (
			args = `{"X":[123, "", "789"]}`

			wantTarget = T{X: []JSONNumber{"123", "", "789"}}
			target     T
		)
		const wantErr = false
		if err := json.Unmarshal([]byte(args), &target); (err != nil) != wantErr {
			t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, wantErr)
		}
		if !reflect.DeepEqual(target, wantTarget) {
			t.Errorf("UnmarshalJSON() got = %v, want %v", target, wantTarget)
		}
	})
	t.Run("struct field array with null", func(t *testing.T) {
		type T struct {
			X []JSONNumber
		}
		var (
			args = `{"X":[123, null, "789"]}`

			target T
		)
		const wantErr = true
		if err := json.Unmarshal([]byte(args), &target); (err != nil) != wantErr {
			t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, wantErr)
		}
	})
}

func np(v string) *JSONNumber {
	x := JSONNumber(v)
	return &x
}

func TestJSONNumber_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		n       JSONNumber
		want    []byte
		wantErr bool
	}{
		{
			name:    "test empty string",
			n:       JSONNumber(""),
			want:    []byte("null"),
			wantErr: false,
		},
		{
			name:    "test not empty string",
			n:       JSONNumber("123"),
			want:    []byte("123"),
			wantErr: false,
		},
		{
			name:    "test json string",
			n:       JSONNumber(`{"test": 123}`),
			want:    []byte(`{"test":123}`),
			wantErr: false,
		},
		{
			name:    "test invalid json string",
			n:       JSONNumber(`"test": 123}`),
			want:    nil,
			wantErr: true,
		},
		{
			name:    "test valid json struct",
			n:       JSONNumber(`{"X":[123, null, "789"]}`),
			want:    []byte(`{"X":[123,null,"789"]}`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.Marshal(tt.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}
