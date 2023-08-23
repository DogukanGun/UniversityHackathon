sudo docker build -t  dogukangun/donationforuniversity:go .
sudo docker push dogukangun/donationforuniversity:go
bacalhau docker run --wait dogukangun/donationforuniversity:go
