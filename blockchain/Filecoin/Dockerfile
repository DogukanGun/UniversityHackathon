# Use an official Python runtime as the base image with the latest versions
FROM python:latest
RUN apt-get update
RUN apt-get install ca-certificates -y
RUN update-ca-certificates
# Set the working directory inside the container
WORKDIR /

# Copy your Python script and requirements file to the container
ADD requirements.txt ./requirements.txt
ADD worker.py ./worker.py

# Install dependencies
RUN pip install --no-cache-dir -r requirements.txt

CMD ["python", "worker.py", "${WALLET_ADDRESS}"]
