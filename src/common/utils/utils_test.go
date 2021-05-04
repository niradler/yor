package utils

import (
	"os"
	"testing"
)

var currentDir, _ = os.Getwd()

func TestGetFileFormat(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "yaml",
			args: args{filePath: "dir/file.yaml"},
			want: "yaml",
		},
		{
			name: "yml",
			args: args{filePath: "dir/file.yml"},
			want: "yml",
		},
		{
			name: "json",
			args: args{filePath: "dir/file.json"},
			want: "json",
		},
		{
			name: "no file type",
			args: args{filePath: "dir/file"},
			want: "",
		},
		{
			name: "empty string",
			args: args{filePath: ""},
			want: "",
		},
		{
			name: "template-yaml",
			args: args{filePath: currentDir + "/../../../tests/cloudformation/resources/extensions/ebs.template"},
			want: "yaml",
		},
		{
			name: "template-yaml",
			args: args{filePath: currentDir + "/../../../tests/cloudformation/resources/extensions/ebs2.template"},
			want: "json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFileFormat(tt.args.filePath); got != tt.want {
				t.Errorf("GetFileFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractIndentationOfLine(t *testing.T) {
	type args struct {
		textLine string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "on indent",
			args: args{textLine: "some text line"},
			want: "",
		},
		{
			name: "3 indents",
			args: args{textLine: "   some text line"},
			want: "   ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtractIndentationOfLine(tt.args.textLine); got != tt.want {
				t.Errorf("ExtractIndentationOfLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInSlice(t *testing.T) {
	type args struct {
		slice interface{}
		elem  interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "in slice string",
			args: args{slice: []string{"a", "b", "c", "e"}, elem: "a"},
			want: true,
		},
		{
			name: "not in slice string",
			args: args{slice: []string{"a", "b", "c", "e"}, elem: "d"},
			want: false,
		},
		{
			name: "in slice int",
			args: args{slice: []int{1, 2, 3, 4}, elem: 1},
			want: true,
		},
		{
			name: "not in slice int",
			args: args{slice: []int{1, 2, 3, 4}, elem: 5},
			want: false,
		},
		{
			name: "slice in slice ",
			args: args{slice: [][]int{{1, 2, 3, 4}, {5, 6}, {7}}, elem: []int{5, 6}},
			want: true,
		},
		{
			name: "not slice in slice ",
			args: args{slice: [][]int{{1, 2, 3, 4}, {5, 6}, {7}}, elem: []int{5, 7}},
			want: false,
		},
		{
			name: "different kinds",
			args: args{slice: []int{1, 2, 3, 4}, elem: "bana"},
			want: false,
		},
		{
			name: "nil slice",
			args: args{slice: nil, elem: "bana"},
			want: false,
		},
		{
			name: "empty slice",
			args: args{slice: []int{}, elem: "bana"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InSlice(tt.args.slice, tt.args.elem); got != tt.want {
				t.Errorf("InSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}