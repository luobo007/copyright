
版本：
go get github.com/btcsuite/btcd@v0.22
 
go get github.com/ethereum/go-ethereum@v1.9.12
 
go get github.com/btcsuite/btcutil@v1.0.1
  
 geth部署： 
geth init --datadir ./data genesis.json

开发模式：
geth --datadir ./data --networkid 18 --port 30303 --http --http.addr 0.0.0.0 --http.vhosts "*"  --http.port 8545 --http.api 'db,net,eth,web3,personal' --http.corsdomain "*" --rpc.enabledeprecatedpersonal --dev --dev.period 1 console 2> 1.log

personal.newAccount('666')
  
eth.getBalance('0xCCAEE0d151a7269D22d66f3E123292081C36DB22')

eth.sendTransaction({from:'0xbe9ab3661f83d3385b896513ead7ec3ef78d9bbc',to:'0xEa0d55DF0d31709866bd947Ad65c22b7be803477',value:web3.toWei(100)})
