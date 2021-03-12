from sys import modules
from os import system
from functools import partial
from time import time, sleep

from mininet.net import Mininet
from mininet.node import CPULimitedHost
from mininet.link import TCLink
from mininet.util import dumpNodeConnections
from mininet.log import setLogLevel
from mininet.cli import CLI

from v2_topos import *
from v2_config import conf

def get_topology():
    privateDirs = []
    host = partial(CPULimitedHost, privateDirs=privateDirs)
    try:
        topo_cls = getattr(modules[__name__], conf["topo"]["class"])
        topo_obj = topo_cls(*conf['topo']["args"], **conf['topo']["kwargs"])
        net = Mininet(topo=topo_obj, host=host, link=TCLink)
        return topo_obj, net
    except Exception as e:
        print("Specified topology not found: ", e)
        exit(0)


def test_topology(topo: Topo, net: Mininet):
    print("Dumping host connections")
    dumpNodeConnections(net.hosts)
    print("Waiting switch connections")
    net.waitConnected()

    print("Testing network connectivity - (i: switches are learning)")
    net.pingAll()
    print("Testing network connectivity - (ii: after learning)")
    net.pingAll()

    print("Get all hosts")
    print(topo.hosts(sort=True))

    # print("Get all links")
    # for link in topo.links(sort=True, withKeys=True, withInfo=True):
    #     pprint(link)
    # print()

    if conf['test']['iperf'] == -1:
        return
    else:
        hosts = [net.get(i) for i in topo.hosts(sort=True)]
        if conf['test']['iperf'] == 0:
            net.iperf((hosts[0], hosts[-1]))
        else:
            [net.iperf((i, j)) for i in hosts for j in hosts if i != j]


def main():
    def delay_command(host, cmd, print=True):
        sleep(0.5)
        if print:
            hs[host - 1].cmdPrint(cmd)
        else:
            hs[host - 1].cmd(cmd)
        sleep(0.5)

    system('sudo mn --clean')
    setLogLevel('info')

    # reads YAML configs and creates the network
    topo, net = get_topology()
    net.start()

    # tests connections (include iperf)
    # test_topology(topo, net)

    hs = topo.hosts(sort=True)
    hs = [net.getNodeByName(h) for h in hs]
    for h in hs[::-1]:
        h.cmdPrint("cd ~")
        h.cmdPrint("ls")
    
    for i in range(len(hs)):
        delay_command(i+1, "~/cassandra/bin/cassandra -R")
        sleep(30)

    for i in range(len(hs)):
        delay_command(i+1, "~/cassandra/bin/nodetool status")
    
    CLI(net) 
    # puts user in CLI

if __name__ == '__main__':
	main()
