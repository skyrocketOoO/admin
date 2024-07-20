#!/bin/bash

# Variables
VM_USER="ubuntu"
VM_HOST="ec2-3-106-200-208.ap-southeast-2.compute.amazonaws.com"

mkdir -p backend
go build -o backend/server

# Transfer the tarball to the VM

ssh -i "~/.ssh/home.pem" $VM_USER@$VM_HOST << 'ENDSSH'
sudo service backend stop
rm -rf backend
ENDSSH

cp -r manifest backend/
scp -r -i "~/.ssh/home.pem" backend $VM_USER@$VM_HOST:


# Connect to the VM and set up the project
ssh -i "~/.ssh/home.pem" $VM_USER@$VM_HOST << 'ENDSSH'
sudo service backend restart
ENDSSH

# Clean up local tarball
rm -rf backend

echo "deployment complete!"
