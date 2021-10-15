package vipd

import "time"

type VirtualIP struct {
	// Interface is the interface to add the IP alias
	Interface string
	// IPAddress is the IP alias to add
	IPAddress string
}

type Config struct {
	// NodeName is the name of the node
	NodeName string
	// VirtualIPs is a list of virtual IP addresses for the node
	VirtualIPs []*VirtualIP
	// ClusterAddress is the cluster bind address
	ClusterAddress string
	// AdvertiseAddress is the cluster address to advertise
	AdvertiseAddress string
	// Peers is a list of peers to join
	Peers []string
	// LeaderPromotionTimeout is the timeout to self-elect as leader
	LeaderPromotionTimeout duration
}

type duration struct {
	time.Duration
}

func (d *duration) UnmarshalText(v []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(v))
	return err
}