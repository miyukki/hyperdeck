package main

import (
	"context"
	"fmt"
	"time"

	"github.com/a-contre-plongee/hyperdeck"
)

var clips hyperdeck.Clips
var c *hyperdeck.Client
var target *hyperdeck.Clip

func main() {
	c = hyperdeck.New("192.168.10.26",
		hyperdeck.WithTransportListener(OnTransportNotification),
		hyperdeck.WithDisplayTimecodeListener(OnDisplayTimecodeNotification),
	)
	err := c.Start(context.Background())
	if err != nil {
		panic(err)
	}
	defer c.Stop()

	clips, err = c.ClipsGet()
	if err != nil {
		panic(err)
	}

	transport, err := c.TransportInfo()
	if err != nil {
		panic(err)
	}
	OnTransportNotification(transport)

	// fmt.Println("Success")

	// list, err := c.DiskList(hyperdeck.SlotCurrent)
	// if err != nil {
	// 	if _, ok := err.(hyperdeck.HyperdeckError); !ok {
	// 		panic(err)
	// 	}
	// }
	// for _, res := range list {
	// 	fmt.Printf("- %s: %s: %v\n", res.Name, res.Duration, res.Duration.Duration())
	// }

	// clips, err := c.ClipsGet()
	// if err != nil {
	// 	fmt.Println(err)
	// 	if _, ok := err.(hyperdeck.HyperdeckError); !ok {
	// 		panic(err)
	// 	}
	// }
	// for _, res := range clips {
	// 	fmt.Printf("- %s: %s: %v\n", res.Name, res.StartAt, res.Duration)
	// }

	// transport, err := c.TransportInfo()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%#v\n", transport)

	// slot, err := c.SlotInfo(hyperdeck.Slot1)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%#v\n", slot)

	for {
		time.Sleep(1 * time.Second)
		if !c.Running() {
			return
		}
	}
}

func OnTransportNotification(transport hyperdeck.Transport) {
	// fmt.Printf("\n NEW NOTIFICATION: TRANSPORT\n\n%#v\n\n", transport)
	if transport.ClipID == 0 {
		return
	}

	// go func() {
	// 	time.Sleep(10 * time.Millisecond)

	// var target *hyperdeck.Clip
	for i := range clips {
		if transport.ClipID == clips[i].ID {
			target = &clips[i]
			break
		}
	}
	if target == nil {
		panic("not found clip")
	}
	// fmt.Printf("\n TARGET CLIP\n\n%#v\n\n", target)
	// }()
}

func OnDisplayTimecodeNotification(timecode hyperdeck.DisplayTimecode) {
	fmt.Printf("\n NEW NOTIFICATION: DISPLAY TIMECODE\n\n%#v\n\n", timecode)
	if target == nil {
		return
	}

	remain := target.Duration.Duration() + target.StartAt.Duration() - timecode.Timecode.Duration()
	fmt.Println(target.Duration.Duration(), target.StartAt.Duration(), timecode.Timecode.Duration(), "\n")
	fmt.Printf("\n REMAIN: %s \n", remain)

	// timecode.Timecode.Duration()
}
