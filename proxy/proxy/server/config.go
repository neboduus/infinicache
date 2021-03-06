package server

import (
	"time"

	"github.com/neboduus/infinicache/proxy/proxy/lambdastore"
)

const AWSRegion = "us-east-2"
const LambdaMaxDeployments = 20
const NumLambdaClusters = 20
// fixed size array with lambdas addresses
var LambdaAddresses = [...]string {
/*	"http://infinicache-node-0.default.34.91.116.154.xip.io",
	"http://infinicache-node-1.default.34.91.116.154.xip.io",
	"http://infinicache-node-2.default.34.91.116.154.xip.io",
	"http://infinicache-node-3.default.34.91.116.154.xip.io",
	"http://infinicache-node-4.default.34.91.116.154.xip.io",
	"http://infinicache-node-5.default.34.91.116.154.xip.io",
	"http://infinicache-node-6.default.34.91.116.154.xip.io",
	"http://infinicache-node-7.default.34.91.116.154.xip.io",
	"http://infinicache-node-8.default.34.91.116.154.xip.io",
	"http://infinicache-node-9.default.34.91.116.154.xip.io",
	"http://infinicache-node-10.default.34.91.116.154.xip.io",
	"http://infinicache-node-12.default.34.91.116.154.xip.io",
	"http://infinicache-node-13.default.34.91.116.154.xip.io",
	"http://infinicache-node-14.default.34.91.116.154.xip.io",
	"http://infinicache-node-15.default.34.91.116.154.xip.io",
	"http://infinicache-node-16.default.34.91.116.154.xip.io",
	"http://infinicache-node-17.default.34.91.116.154.xip.io",
	"http://infinicache-node-18.default.34.91.116.154.xip.io",
	"http://infinicache-node-19.default.34.91.116.154.xip.io",*/
	"http://infinicache-node-0.default.svc.cluster.local",
	"http://infinicache-node-1.default.svc.cluster.local",
	"http://infinicache-node-2.default.svc.cluster.local",
	"http://infinicache-node-3.default.svc.cluster.local",
	"http://infinicache-node-4.default.svc.cluster.local",
	"http://infinicache-node-5.default.svc.cluster.local",
	"http://infinicache-node-6.default.svc.cluster.local",
	"http://infinicache-node-7.default.svc.cluster.local",
	"http://infinicache-node-8.default.svc.cluster.local",
	"http://infinicache-node-9.default.svc.cluster.local",
	"http://infinicache-node-10.default.svc.cluster.local",
	"http://infinicache-node-11.default.svc.cluster.local",
	"http://infinicache-node-12.default.svc.cluster.local",
	"http://infinicache-node-13.default.svc.cluster.local",
	"http://infinicache-node-14.default.svc.cluster.local",
	"http://infinicache-node-15.default.svc.cluster.local",
	"http://infinicache-node-16.default.svc.cluster.local",
	"http://infinicache-node-17.default.svc.cluster.local",
	"http://infinicache-node-18.default.svc.cluster.local",
	"http://infinicache-node-19.default.svc.cluster.local",
}
const LambdaStoreName = "LambdaStore" // replica version (no use)
const LambdaPrefix = "CacheNode"
const InstanceWarmTimout = 1 * time.Minute
const InstanceCapacity = 1536 * 1000000 // MB
const InstanceOverhead = 100 * 1000000  // MB
const ServerPublicIp = "10.4.0.100" // Leave it empty if using VPC.

func init() {
	lambdastore.WarmTimout = InstanceWarmTimout
}
