// automatically generated by stateify.

//go:build linux && ((linux && amd64) || (linux && arm64)) && linux && linux
// +build linux
// +build linux,amd64 linux,arm64
// +build linux
// +build linux

package fdbased

import (
	"context"

	"gvisor.dev/gvisor/pkg/state"
)

func (f *fdInfo) StateTypeName() string {
	return "pkg/tcpip/link/fdbased.fdInfo"
}

func (f *fdInfo) StateFields() []string {
	return []string{
		"fd",
		"isSocket",
	}
}

func (f *fdInfo) beforeSave() {}

// +checklocksignore
func (f *fdInfo) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.fd)
	stateSinkObject.Save(1, &f.isSocket)
}

func (f *fdInfo) afterLoad(context.Context) {}

// +checklocksignore
func (f *fdInfo) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.fd)
	stateSourceObject.Load(1, &f.isSocket)
}

func (e *endpoint) StateTypeName() string {
	return "pkg/tcpip/link/fdbased.endpoint"
}

func (e *endpoint) StateFields() []string {
	return []string{
		"fds",
		"hdrSize",
		"caps",
		"inboundDispatchers",
		"dispatcher",
		"packetDispatchMode",
		"gsoMaxSize",
		"gsoKind",
		"maxSyscallHeaderBytes",
		"writevMaxIovs",
		"addr",
		"mtu",
	}
}

func (e *endpoint) beforeSave() {}

// +checklocksignore
func (e *endpoint) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.fds)
	stateSinkObject.Save(1, &e.hdrSize)
	stateSinkObject.Save(2, &e.caps)
	stateSinkObject.Save(3, &e.inboundDispatchers)
	stateSinkObject.Save(4, &e.dispatcher)
	stateSinkObject.Save(5, &e.packetDispatchMode)
	stateSinkObject.Save(6, &e.gsoMaxSize)
	stateSinkObject.Save(7, &e.gsoKind)
	stateSinkObject.Save(8, &e.maxSyscallHeaderBytes)
	stateSinkObject.Save(9, &e.writevMaxIovs)
	stateSinkObject.Save(10, &e.addr)
	stateSinkObject.Save(11, &e.mtu)
}

func (e *endpoint) afterLoad(context.Context) {}

// +checklocksignore
func (e *endpoint) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.fds)
	stateSourceObject.Load(1, &e.hdrSize)
	stateSourceObject.Load(2, &e.caps)
	stateSourceObject.Load(3, &e.inboundDispatchers)
	stateSourceObject.Load(4, &e.dispatcher)
	stateSourceObject.Load(5, &e.packetDispatchMode)
	stateSourceObject.Load(6, &e.gsoMaxSize)
	stateSourceObject.Load(7, &e.gsoKind)
	stateSourceObject.Load(8, &e.maxSyscallHeaderBytes)
	stateSourceObject.Load(9, &e.writevMaxIovs)
	stateSourceObject.Load(10, &e.addr)
	stateSourceObject.Load(11, &e.mtu)
}

func (o *Options) StateTypeName() string {
	return "pkg/tcpip/link/fdbased.Options"
}

func (o *Options) StateFields() []string {
	return []string{
		"FDs",
		"MTU",
		"EthernetHeader",
		"ClosedFunc",
		"Address",
		"SaveRestore",
		"DisconnectOk",
		"GSOMaxSize",
		"GVisorGSOEnabled",
		"PacketDispatchMode",
		"TXChecksumOffload",
		"RXChecksumOffload",
		"MaxSyscallHeaderBytes",
		"InterfaceIndex",
		"GRO",
		"ProcessorsPerChannel",
	}
}

func (o *Options) beforeSave() {}

// +checklocksignore
func (o *Options) StateSave(stateSinkObject state.Sink) {
	o.beforeSave()
	stateSinkObject.Save(0, &o.FDs)
	stateSinkObject.Save(1, &o.MTU)
	stateSinkObject.Save(2, &o.EthernetHeader)
	stateSinkObject.Save(3, &o.ClosedFunc)
	stateSinkObject.Save(4, &o.Address)
	stateSinkObject.Save(5, &o.SaveRestore)
	stateSinkObject.Save(6, &o.DisconnectOk)
	stateSinkObject.Save(7, &o.GSOMaxSize)
	stateSinkObject.Save(8, &o.GVisorGSOEnabled)
	stateSinkObject.Save(9, &o.PacketDispatchMode)
	stateSinkObject.Save(10, &o.TXChecksumOffload)
	stateSinkObject.Save(11, &o.RXChecksumOffload)
	stateSinkObject.Save(12, &o.MaxSyscallHeaderBytes)
	stateSinkObject.Save(13, &o.InterfaceIndex)
	stateSinkObject.Save(14, &o.GRO)
	stateSinkObject.Save(15, &o.ProcessorsPerChannel)
}

func (o *Options) afterLoad(context.Context) {}

// +checklocksignore
func (o *Options) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &o.FDs)
	stateSourceObject.Load(1, &o.MTU)
	stateSourceObject.Load(2, &o.EthernetHeader)
	stateSourceObject.Load(3, &o.ClosedFunc)
	stateSourceObject.Load(4, &o.Address)
	stateSourceObject.Load(5, &o.SaveRestore)
	stateSourceObject.Load(6, &o.DisconnectOk)
	stateSourceObject.Load(7, &o.GSOMaxSize)
	stateSourceObject.Load(8, &o.GVisorGSOEnabled)
	stateSourceObject.Load(9, &o.PacketDispatchMode)
	stateSourceObject.Load(10, &o.TXChecksumOffload)
	stateSourceObject.Load(11, &o.RXChecksumOffload)
	stateSourceObject.Load(12, &o.MaxSyscallHeaderBytes)
	stateSourceObject.Load(13, &o.InterfaceIndex)
	stateSourceObject.Load(14, &o.GRO)
	stateSourceObject.Load(15, &o.ProcessorsPerChannel)
}

func (e *InjectableEndpoint) StateTypeName() string {
	return "pkg/tcpip/link/fdbased.InjectableEndpoint"
}

func (e *InjectableEndpoint) StateFields() []string {
	return []string{
		"endpoint",
		"dispatcher",
	}
}

func (e *InjectableEndpoint) beforeSave() {}

// +checklocksignore
func (e *InjectableEndpoint) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.endpoint)
	stateSinkObject.Save(1, &e.dispatcher)
}

func (e *InjectableEndpoint) afterLoad(context.Context) {}

// +checklocksignore
func (e *InjectableEndpoint) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.endpoint)
	stateSourceObject.Load(1, &e.dispatcher)
}

func (d *packetMMapDispatcher) StateTypeName() string {
	return "pkg/tcpip/link/fdbased.packetMMapDispatcher"
}

func (d *packetMMapDispatcher) StateFields() []string {
	return []string{
		"StopFD",
		"fd",
		"e",
		"ringBuffer",
		"ringOffset",
		"mgr",
	}
}

func (d *packetMMapDispatcher) beforeSave() {}

// +checklocksignore
func (d *packetMMapDispatcher) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.StopFD)
	stateSinkObject.Save(1, &d.fd)
	stateSinkObject.Save(2, &d.e)
	stateSinkObject.Save(3, &d.ringBuffer)
	stateSinkObject.Save(4, &d.ringOffset)
	stateSinkObject.Save(5, &d.mgr)
}

func (d *packetMMapDispatcher) afterLoad(context.Context) {}

// +checklocksignore
func (d *packetMMapDispatcher) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.StopFD)
	stateSourceObject.Load(1, &d.fd)
	stateSourceObject.Load(2, &d.e)
	stateSourceObject.Load(3, &d.ringBuffer)
	stateSourceObject.Load(4, &d.ringOffset)
	stateSourceObject.Load(5, &d.mgr)
}

func (b *iovecBuffer) StateTypeName() string {
	return "pkg/tcpip/link/fdbased.iovecBuffer"
}

func (b *iovecBuffer) StateFields() []string {
	return []string{
		"views",
		"sizes",
		"skipsVnetHdr",
		"pulledIndex",
	}
}

func (b *iovecBuffer) beforeSave() {}

// +checklocksignore
func (b *iovecBuffer) StateSave(stateSinkObject state.Sink) {
	b.beforeSave()
	stateSinkObject.Save(0, &b.views)
	stateSinkObject.Save(1, &b.sizes)
	stateSinkObject.Save(2, &b.skipsVnetHdr)
	stateSinkObject.Save(3, &b.pulledIndex)
}

func (b *iovecBuffer) afterLoad(context.Context) {}

// +checklocksignore
func (b *iovecBuffer) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &b.views)
	stateSourceObject.Load(1, &b.sizes)
	stateSourceObject.Load(2, &b.skipsVnetHdr)
	stateSourceObject.Load(3, &b.pulledIndex)
}

func (d *readVDispatcher) StateTypeName() string {
	return "pkg/tcpip/link/fdbased.readVDispatcher"
}

func (d *readVDispatcher) StateFields() []string {
	return []string{
		"StopFD",
		"fd",
		"e",
		"buf",
		"mgr",
	}
}

func (d *readVDispatcher) beforeSave() {}

// +checklocksignore
func (d *readVDispatcher) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.StopFD)
	stateSinkObject.Save(1, &d.fd)
	stateSinkObject.Save(2, &d.e)
	stateSinkObject.Save(3, &d.buf)
	stateSinkObject.Save(4, &d.mgr)
}

func (d *readVDispatcher) afterLoad(context.Context) {}

// +checklocksignore
func (d *readVDispatcher) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.StopFD)
	stateSourceObject.Load(1, &d.fd)
	stateSourceObject.Load(2, &d.e)
	stateSourceObject.Load(3, &d.buf)
	stateSourceObject.Load(4, &d.mgr)
}

func (r *recvMMsgDispatcher) StateTypeName() string {
	return "pkg/tcpip/link/fdbased.recvMMsgDispatcher"
}

func (r *recvMMsgDispatcher) StateFields() []string {
	return []string{
		"StopFD",
		"fd",
		"e",
		"bufs",
		"pkts",
		"gro",
		"mgr",
	}
}

func (r *recvMMsgDispatcher) beforeSave() {}

// +checklocksignore
func (r *recvMMsgDispatcher) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.StopFD)
	stateSinkObject.Save(1, &r.fd)
	stateSinkObject.Save(2, &r.e)
	stateSinkObject.Save(3, &r.bufs)
	stateSinkObject.Save(4, &r.pkts)
	stateSinkObject.Save(5, &r.gro)
	stateSinkObject.Save(6, &r.mgr)
}

// +checklocksignore
func (r *recvMMsgDispatcher) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.StopFD)
	stateSourceObject.Load(1, &r.fd)
	stateSourceObject.Load(2, &r.e)
	stateSourceObject.Load(3, &r.bufs)
	stateSourceObject.Load(4, &r.pkts)
	stateSourceObject.Load(5, &r.gro)
	stateSourceObject.Load(6, &r.mgr)
	stateSourceObject.AfterLoad(func() { r.afterLoad(ctx) })
}

func (p *processor) StateTypeName() string {
	return "pkg/tcpip/link/fdbased.processor"
}

func (p *processor) StateFields() []string {
	return []string{
		"pkts",
		"e",
		"gro",
		"sleeper",
		"packetWaker",
		"closeWaker",
	}
}

func (p *processor) beforeSave() {}

// +checklocksignore
func (p *processor) StateSave(stateSinkObject state.Sink) {
	p.beforeSave()
	stateSinkObject.Save(0, &p.pkts)
	stateSinkObject.Save(1, &p.e)
	stateSinkObject.Save(2, &p.gro)
	stateSinkObject.Save(3, &p.sleeper)
	stateSinkObject.Save(4, &p.packetWaker)
	stateSinkObject.Save(5, &p.closeWaker)
}

func (p *processor) afterLoad(context.Context) {}

// +checklocksignore
func (p *processor) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &p.pkts)
	stateSourceObject.Load(1, &p.e)
	stateSourceObject.Load(2, &p.gro)
	stateSourceObject.Load(3, &p.sleeper)
	stateSourceObject.Load(4, &p.packetWaker)
	stateSourceObject.Load(5, &p.closeWaker)
}

func (m *processorManager) StateTypeName() string {
	return "pkg/tcpip/link/fdbased.processorManager"
}

func (m *processorManager) StateFields() []string {
	return []string{
		"processors",
		"seed",
		"e",
		"ready",
	}
}

func (m *processorManager) beforeSave() {}

// +checklocksignore
func (m *processorManager) StateSave(stateSinkObject state.Sink) {
	m.beforeSave()
	stateSinkObject.Save(0, &m.processors)
	stateSinkObject.Save(1, &m.seed)
	stateSinkObject.Save(2, &m.e)
	stateSinkObject.Save(3, &m.ready)
}

func (m *processorManager) afterLoad(context.Context) {}

// +checklocksignore
func (m *processorManager) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &m.processors)
	stateSourceObject.Load(1, &m.seed)
	stateSourceObject.Load(2, &m.e)
	stateSourceObject.Load(3, &m.ready)
}

func init() {
	state.Register((*fdInfo)(nil))
	state.Register((*endpoint)(nil))
	state.Register((*Options)(nil))
	state.Register((*InjectableEndpoint)(nil))
	state.Register((*packetMMapDispatcher)(nil))
	state.Register((*iovecBuffer)(nil))
	state.Register((*readVDispatcher)(nil))
	state.Register((*recvMMsgDispatcher)(nil))
	state.Register((*processor)(nil))
	state.Register((*processorManager)(nil))
}
