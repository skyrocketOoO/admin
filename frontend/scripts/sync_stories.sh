VM_USER="ubuntu"
VM_HOST="ec2-3-106-200-208.ap-southeast-2.compute.amazonaws.com"

rsync -avz -e "ssh -i '~/.ssh/home.pem'" app/blog/stories $VM_USER@$VM_HOST:frontend/app/blog/