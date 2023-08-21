sudo docker build -t  dogukangun/donationforuniversity:python3 .
sudo docker push dogukangun/donationforuniversity:python3
bacalhau docker run -e WALLET_ADDRESS=12313dasdas --wait dogukangun/donationforuniversity:python3
