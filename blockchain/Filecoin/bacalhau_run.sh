sudo docker build -t  dogukangun/donationforuniversity:go_latest6 .
sudo docker push dogukangun/donationforuniversity:go_latest6
bacalhau docker run --wait --network=http --domain=eth.public-rpc.com -e WALLET_ADDRESS=0x6097f22127E2EF98C2cF31335Bede16D742f6890 dogukangun/donationforuniversity:go_latest6
