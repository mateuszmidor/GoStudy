# Social poll
Filter twitter tweets against keywords and periodically display updates on a web page in form of a pie chart  

# Design
Container diagram (C4 arch model): [design.txt](design.txt)  

# Environment variables
The following strings are required as env variables (see: envconfig.sh):  
> SP_TWITTER_KEY=  
> SP_TWITTER_SECRET=  
> SP_TWITTER_ACCESSTOKEN=  
> SP_TWITTER_ACCESSSECRET=  
> SP_MONGODB_ADDR=  
> SP_NSQD_ADDR=
> SP_NSQLOOKUP_ADDR=

# Setup MongoDB sample poll keywords (need install mongo first) 
> sudo pacman -S mongodb  
> mongo  
> use ballots  
> db.polls.insert({"title":"Mood Poll", "options":["love", "kiss", "hate", "fight"]})  

# Track NSQ messages (need install nsq first)
> sudo pacman -S yaourt  
> yaourt nsq  
> nsq_tail --topic="votes" --lookupd-http-address=localhost:4161  

# Run the app locally
The app is made of 3 services and 4 components running inside 4 containers under docker-compose:
> mongodb  
> nsqlookupd  
> nsqd  
> socialpoll: twittervotes + counter + api + web  

So we need to install docker & docker-compose, later we run with one command :)
1. Install docker, add yourself to docker group so you dont need be root:  
> sudo pacman -S docker  
> sudo usermod -aG docker $(whoami)  
> logout/login to shell again
2. Enable docker service autostart and run the service:  
> sudo systemctl enable docker.service  
> sudo systemctl start docker.service  
3. Install docker-compose:  
> sudo pacman -S python-pip  
> sudo pip install docker-compose 
4. Run the app:
> ./docker-compose_up.sh
5. See: localhost:8081

# Run the app on AWS instance
0. Create network security key-pair, save the key-pair as ~/.ssh/mm_keypair.pem  
1. Create Amazon Linux instance with docker package, set security key-pair  
2. Setup instance->security groups->inbound:  
    allow SSH 22 for your computer IP
3. SSH to the instance (replace the instance dns with yours):  
    ssh -v -i ~/.ssh/mm_keypair.pem ec2-user@ec2-54-93-96-97.eu-central-1.compute.amazonaws.com 
4. Install docker, add yourself to docker group so you dont need be root:  
    sudo yum install docker  
    sudo usermod -aG docker $(whoami)  
    logout/login again
5. Enable docker service autostart and run the service:  
    sudo systemctl enable docker.service  
    sudo systemctl start docker.service  
6. Install docker-compose:  
    sudo yum install -y python-pip  
    sudo pip install docker-compose  
7. sftp entire project to AWS instance  
8. ./docker-compose_up.sh
9. See: AWS_INSTANCE_IP:8081