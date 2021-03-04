#!/usr/bin/env bash

num_of_miners=$1

# delete all host folders
# delete all logs
rm -rf ~/BBB/hosts/* ~/BBB/data/*

# make num_of_miners # of folders
cur_miner=1
while [ $cur_miner -le $num_of_miners ]; do
  mkdir -p ~/BBB/hosts/ethData$cur_miner/keystore
  source ~/.bashrc

  # test and set env var
  dirName=DDR${cur_miner}
  if [[ -z ${!dirName} || ${!dirName} != ~/BBB/hosts/ethData$cur_miner ]]; then
    echo export DDR$cur_miner=~/BBB/hosts/ethData$cur_miner >>~/.bashrc
  fi
  source ~/.bashrc

  # geth init
  geth init --datadir ${!dirName} ~/BBB/genesis.json

  # copy keys
  cp ~/BBB/keys/UTC--2020-02-25T03-34-02.408074800Z--64fe25df48a9f21887d3ee886435e58e501e7623 ~/BBB/hosts/ethData$cur_miner/keystore
  cp ~/BBB/keys/UTC--2020-05-18T14-38-18.472973295Z--c838c31639157d40e029c18fb54226212c060c0b ~/BBB/hosts/ethData$cur_miner/keystore
  cp ~/BBB/keys/UTC--2020-05-18T14-38-29.783530609Z--2292d6fbd29694edb25f68a13df023bfe32133f3 ~/BBB/hosts/ethData$cur_miner/keystore
  cp ~/BBB/keys/UTC--2020-05-18T14-38-37.867695000Z--2be146309b4b01b0a0ed51c26cb14bc82917540f ~/BBB/hosts/ethData$cur_miner/keystore
  ((cur_miner++))
done
