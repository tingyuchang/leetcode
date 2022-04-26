package stack

import (
	"reflect"
	"testing"
)

func TestFakeQueue_Add(t *testing.T) {
	type fields struct {
		stackA *Stack
		StackB *Stack
	}
	type args struct {
		record interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"add 1 elements, should be executed normally.",
			fields{
			&Stack{[]interface{}{}},
			&Stack{[]interface{}{}},
		}, args{1}},
		{"add 1 elements, should be executed normally.",
			fields{
				&Stack{[]interface{}{}},
				&Stack{[]interface{}{}},
			}, args{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fq := &FakeQueue{
				stackA: tt.fields.stackA,
				StackB: tt.fields.StackB,
			}
			fq.Add(tt.args.record)
		})
	}
}

func TestFakeQueue_Remove(t *testing.T) {
	type fields struct {
		stackA *Stack
		StackB *Stack
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fq := &FakeQueue{
				stackA: tt.fields.stackA,
				StackB: tt.fields.StackB,
			}
			if got := fq.Remove(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	type fields struct {
		data []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				data: tt.fields.data,
			}
			if got := s.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Push(t *testing.T) {
	type fields struct {
		data []interface{}
	}
	type args struct {
		record interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"test stack", fields{[]interface{}{}}, args{1}},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				data: tt.fields.data,
			}
			s.Push(tt.fields.data)
		})
	}
}

func TestStack_length(t *testing.T) {
	type fields struct {
		data []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"test length", fields{
			[]interface{}{
				Stack{data: []interface{}{1, 2, 3}},
			},
		}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				data: tt.fields.data,
			}
			if got := s.length(); got != tt.want {
				t.Errorf("length() = %v, want %v", got, tt.want)
			}
		})
	}
}
