# scalable-chat-service
A scalable chat BE service implemented in golang

## Setup Redis
- Install docker 
- `docker pull redis`
- `docker run --name <some_container_name > -d redis` // use -d to run in detached mode
- `docker ps -a` // check the status of the container
- `docker exec -it <CONTAINER_ID> bash` // ssh into the container 
- `redis-cli` // start cli shell for redis
- `ping` // if success, it will respond with PONG
- To find the host
  - `docker inspect <CONTAINER_ID>`
  - Under Network section find the IP address of your container
  - `telnet <continer_ID_addr> 6379` // If no port forwarding is done then redis will run on default port
  - A successful connection will tell us that redis is running successfully
