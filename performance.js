// Transactions should be sent with the following sytax in order to properly measure performance:
// eth.sendTransaction({from:<from>,to:<to>,value:1,data:web3.toHex(new Date().getTime().toString() + "00")})
// + "00" is important in order to avoid the following error: "cannot unmarshal hex string of odd length"

var args = require('yargs').argv; //nodejs module for command line inputs
var connect = (args.connect === undefined) ? 'ws://localhost:8101' : args.connect; //if --connect is undefined, set it node1 addr:port

var Web3 = require('web3'); //nodejs module for web3
var web3 = new Web3();
web3.setProvider(new web3.providers.WebsocketProvider(connect)) //connect to specified node via WebSocket

fs = require('fs'); //nodejs module for writing to file system
fs.unlink('./data/performance.txt', (_) => {}); //comment this line to keep file between uses

async function checkPerformance() {
    var latestBlock = await web3.eth.getBlockNumber();
    var block = await web3.eth.getBlock(latestBlock, true);
    if (block.transactions != null) {
        var minedTxTime = new Date().getTime();
        block.transactions.forEach(async function (tx) {
            try{
                var writtenTxTime = web3.utils.hexToNumber(tx.input)/100 //divide by 100 to get rid of extra 0s
                var latency = minedTxTime - writtenTxTime;
                fs.appendFileSync('./data/performance.txt', (latency.toString() + "\n"), (err) => { //synchronously writes to ./data/performance.txt
                    if (err) {return console.log(err);}
                });
            } catch (error) {
                console.log("Tx input not correctly formatted");
            }
        })   
    }
}

checkPerformance();

web3.eth.subscribe('newBlockHeaders', function(error, _){
    if (!error) {checkPerformance()} 
})