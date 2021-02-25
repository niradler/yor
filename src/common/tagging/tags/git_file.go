package tags

import (
	"bridgecrewio/yor/src/common/gitservice"
	"fmt"
	"reflect"
)

type GitFileTag struct {
	Tag
}

func (t *GitFileTag) Init() {
	t.Key = "git_file"
}

func (t *GitFileTag) CalculateValue(data interface{}) (ITag, error) {
	gitBlame, ok := data.(*gitservice.GitBlame)
	if !ok {
		return nil, fmt.Errorf("failed to convert data to *GitBlame, which is required to calculte tag value. Type of data: %s", reflect.TypeOf(data))
	}
	return &Tag{Key: t.Key, Value: gitBlame.FilePath}, nil
}