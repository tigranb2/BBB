while true
sleep 3
do 
    geth attach ws://10.0.0.1:8101 --exec 'eth.sendTransaction({from:eth.accounts[0],to:eth.accounts[1],value:1,data:web3.toHex(new Date().getTime().toString() + "00")})'
done &