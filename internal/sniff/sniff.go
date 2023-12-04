package sniff

/*
TODO: ebpf, xdp, io_uring
*/

import (
	"capture/internal/app"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func OpenInterface(iface string, opts *Options) *gopacket.PacketSource {
	app.Log.Trace("sniff: open interface " + iface)
	myIface, err := net.InterfaceByName(iface)
	if err == nil {
		app.Log.Debugf("mtu=%d", myIface.MTU)
		app.Log.Debugf("address=%s", myIface.HardwareAddr.String())
	}

	inactive, err := pcap.NewInactiveHandle(iface)
	app.DieOnErr(err)
	defer inactive.CleanUp()

	err = inactive.SetTimeout(pcap.BlockForever)
	app.DieOnErr(err)

	err = inactive.SetBufferSize(opts.BufSize)
	app.DieOnErr(err)

	err = inactive.SetSnapLen(opts.SnapLen)
	app.DieOnErr(err)

	// Active handle
	handle, err := inactive.Activate() // after this, inactive is no longer valid
	app.DieOnErr(err)

	err = handle.SetDirection(opts.Direction)
	app.DieOnErr(err)

	err = handle.SetBPFFilter(opts.Filter)
	app.DieOnErr(err)

	return gopacket.NewPacketSource(handle, handle.LinkType())
}

func OpenFilename(filename string, bpfFilter string) *gopacket.PacketSource {
	app.Log.Trace("sniff: open file " + filename)
	handle, err := pcap.OpenOffline(filename)
	app.DieOnErr(err)

	if bpfFilter != "" {
		err = handle.SetBPFFilter(bpfFilter)
		app.DieOnErr(err)
	}

	return gopacket.NewPacketSource(handle, handle.LinkType())
}
