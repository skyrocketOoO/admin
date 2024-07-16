#!/bin/bash

# Variables
REPO_SSH_URL="git@github.com:yourusername/your-repo.git"
VM_USER="ubuntu"
VM_HOST="your-vm-domain-or-ip"
PROJECT_DIR_NAME="your-repo" # The directory name after cloning

# Build the project locally
pnpm install
pnpm build

# Create a tarball of the build and necessary files
mkdir -p deploy
cp -r .next deploy/
cp -r public deploy/
cp next.config.js deploy/
cp .env deploy/
cp pnpm-lock.yaml deploy/
cp package.json deploy/
tar -czvf deploy.tar.gz deploy
rm -rf deploy

# Transfer the tarball to the VM
scp deploy.tar.gz $VM_USER@$VM_HOST:

# Connect to the VM and set up the project
ssh $VM_USER@$VM_HOST << 'ENDSSH'
cd $DEST_DIR
tar -xzvf deploy.tar.gz
rm deploy.tar.gz

# Install dependencies
cd deploy
pnpm install --prod

# Start the application (optional: configure to run as a service)
sudo service frontend restart
ENDSSH

# Clean up local tarball
rm deploy.tar.gz

echo "Deployment complete!"
