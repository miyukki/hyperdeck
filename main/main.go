package main

import (
	"context"
	"fmt"
	"time"

	"github.com/a-contre-plongee/hyperdeck"
)

func main() {
	c := hyperdeck.New("192.168.20.243",
		hyperdeck.WithRepeater("0.0.0.0"),
		hyperdeck.WithTransportListener(OnTransportNotification),
		hyperdeck.WithSlotListener(OnSlotNotification),
	)
	err := c.Start(context.Background())
	if err != nil {
		panic(err)
	}
	defer c.Stop()
	fmt.Println("Success")

	list, err := c.DiskList(hyperdeck.SlotCurrent)
	if err != nil {
		if _, ok := err.(hyperdeck.HyperdeckError); !ok {
			panic(err)
		}
	}
	for _, res := range list {
		fmt.Printf("- %s: %s: %v\n", res.Name, res.Duration, res.Duration.Duration())
	}

	clips, err := c.ClipsGet()
	if err != nil {
		fmt.Println(err)
		if _, ok := err.(hyperdeck.HyperdeckError); !ok {
			panic(err)
		}
	}
	for _, res := range clips {
		fmt.Printf("- %s: %s: %v\n", res.Name, res.StartAt, res.Duration)
	}

	transport, err := c.TransportInfo()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", transport)

	slot, err := c.SlotInfo(hyperdeck.Slot1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", slot)

	for {
		time.Sleep(1 * time.Second)
		if !c.Running() {
			return
		}
	}
}

func OnTransportNotification(transport hyperdeck.Transport) {
	fmt.Printf("\n NEW NOTIFICATION: TRANSPORT\n\n%#v\n\n", transport)
}

func OnSlotNotification(slot hyperdeck.Slot) {
	fmt.Printf("\n NEW NOTIFICATION: SLOT\n\n%#v\n\n", slot)
}
