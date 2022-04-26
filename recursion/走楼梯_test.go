package recursion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalWaysNum(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantNum int
	}{
		{name: "n < 0", args: args{n: -1}, wantNum: 0},
		{name: "n normal", args: args{n: 5}, wantNum: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalWaysNum(tt.args.n)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantNum, got)
		})
	}
}
