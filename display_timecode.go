package hyperdeck

import (
	"github.com/pkg/errors"
	"github.com/trimmer-io/go-timecode/timecode"
)

type DisplayTimecodeListener func(DisplayTimecode)

type DisplayTimecode struct {
	Timecode timecode.Timecode `header:"display timecode"`
}

func ParseDisplayTimecode(payload []byte) (DisplayTimecode, error) {
	var res DisplayTimecode
	err := Parse(payload, &res)
	if err != nil {
		return res, errors.Wrap(err, "fail to parse display timecode")
	}
	return res, nil

}
