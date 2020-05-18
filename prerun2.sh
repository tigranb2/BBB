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
  cp ~/BBB/keys/UTC--2020-03-07T23-25-18.301273092Z--67e37abe6fb7bb2b0d61b9c6f53c71623ae65551 ~/BBB/hosts/ethData$cur_miner/keystore
  cp ~/BBB/keys/UTC--2020-03-07T23-27-12.170957487Z--2dec65f7f6fecef9088afed7ab41ad0f1173ddb4 ~/BBB/hosts/ethData$cur_miner/keystore
  cp ~/BBB/keys/UTC--2020-03-07T23-27-53.050585459Z--0213af577d12cf11a5baf5a869e0b1305684ca0a ~/BBB/hosts/ethData$cur_miner/keystore
  cp ~/BBB/keys/UTC--2020-03-16T21-29-01.688713168Z--7d8466475a66c4363da52494af4a3c20298f5f73 ~/BBB/hosts/ethData$cur_miner/keystore
  ((cur_miner++))
done



