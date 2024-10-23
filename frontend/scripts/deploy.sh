#!/bin/bash

# Variables
VM_USER="ubuntu"
VM_HOST="ec2-3-106-200-208.ap-southeast-2.compute.amazonaws.com"
PROJECT_DIR_NAME="fontend" # The directory name after cloning

# Build the project locally
rm -rf .next
pnpm install
if ! pnpm build; then
  echo "Build failed. Exiting..."
  exit 1
fi

# Create a tarball of the build and necessary files
mkdir -p frontend
mkdir -p frontend/app/blog/stories
cp app/blog/stories/*  frontend/app/blog/stories/
cp -r .next frontend/
cp -r public frontend/
cp next.config.mjs frontend/
cp .env frontend/
cp pnpm-lock.yaml frontend/
cp package.json frontend/
tar -czvf frontend.tar.gz frontend
rm -rf frontend

ssh -i "~/.ssh/home.pem" $VM_USER@$VM_HOST << 'ENDSSH'
sudo service frontend stop
sudo rm -rf frontend
ENDSSH

# Transfer the tarball to the VM
scp -i "~/.ssh/home.pem" frontend.tar.gz $VM_USER@$VM_HOST:

# Connect to the VM and set up the project
ssh -i "~/.ssh/home.pem" $VM_USER@$VM_HOST << 'ENDSSH'

tar -xzvf frontend.tar.gz
rm frontend.tar.gz

# Install dependencies
cd frontend
pnpm install --prod

# Start the application (optional: configure to run as a service)
sudo service frontend restart
ENDSSH

# Clean up local tarball
rm frontend.tar.gz

echo "deployment complete!"
