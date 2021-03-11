# Notes for The API
Designed with JWT, Casbin, CockroachDB and CI tested using Jenkins and docker with a
script to start up CockroachDB


Introduction
------------
Currently the most up to date version of the API is in the 'failedDesign' branch. It was thought that the design was really bad, but in fact it is actually better than the design in the master branch. Soon I will fix this so that the best work will be on master.


Building
------------
This is not yet tested on Windows and macOS   
The ReadMe assumes that you already have GO installed and your Environment PATHS properly setup  
You must also get [CockroachDB](https://www.cockroachlabs.com/docs/stable/install-cockroachdb-linux.html)  
For BSD systems, GO can be installed from the respective ports  
For Linux, consult your respective Distribution docs  

[Otherwise you can follow instructions from the GO site](https://golang.org/doc/install)  
   
  Clone the API repository  
  Execute go build main.go It should automatically retrieve the necessary libraries. If not then execute the command below to obtain the following packages
  ```
  go get github.com/dgrijalva/jwt-go github.com/fsnotify/fsnotify github.com/gorilla/mux github.com/jinzhu/gorm github.com/joho/godotenv github.com/lib/pq github.com/magiconair/properties github.com/mitchellh/mapstructure github.com/pelletier/go-toml github.com/spf13/afero github.com/spf13/cobra golang.org/x/crypto golang.org/x/sys golang.org/x/text gopkg.in/ini.v1  
  ```  

   Execute go build main.go


Running
-------------
You can modify the port of the API in the .env file 
 - Execute the bash script start.sh
 - Execute ```go run main.go```

Anatomy
-------------
The API follows the [MVC Architecture](https://en.wikipedia.org/wiki/Model%E2%80%93view%E2%80%93controller) using an [Active Record](https://en.wikipedia.org/wiki/Active_record_pattern) Persistence Pattern (Managed by the Model)   
The code is divided into isolated components where each component performs a certain set of tasks. The Model manages the Data and interacts with the Database   
The Controller interacts with the View and Model   
The View is the front end   
<p align="center">
  <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/a/a0/MVC-Process.svg/218px-MVC-Process.svg.png">
</p>


   
This design was chosen to follow the microservices design for scalability and potential reuse in other projects such as GANDALF
   
API Files
-------------
   
### Folder Structure   
```
├─app   
├─controllers   
├─models    
├─resources   
├───diagrams   
└─utils   
```
    

The 'app' dir contains error and authentication files  
The 'controllers' dir contains controller files  
The 'models' dir contains model files  
The 'resources' dir contains SQL files and DB Diagram information   
The 'utils' dir contains useful functions for JSON messaging

### Files of Interest in the root directory  
```
- start.sh   
- main.go  
- .env   
- doc.go   
- swagger.json   
- Jenkinsfile   
- Dockerfile   
```

   
The start.sh is a script to start the CockroachDB   
main.go is the entry of the API  
.env contains variables useful for the API  
doc.go & swagger.json are OpenAPI documentation files   
Jenkinsfile will be used to implement a CI Pipeline later  
Dockerfile is a file for docker to create a container of the API  





### API UML   
![image info](./resources/diagrams/UML.png)




Setting Up Jenkins
--------------------------
https://www.jenkins.io/doc/book/installing/docker/

Jenkins Username: admin

Jenkins Password: 
```
0e33d7d079fa4848b3dee52494674c33
```



Establish the network bridge
```
docker network create jenkins
```

Obtain the dind image to allow jenkins to execute docker commands
```
docker run --name jenkins-docker --rm --detach \
  --privileged --network jenkins --network-alias docker \
  --env DOCKER_TLS_CERTDIR=/certs \
  --volume jenkins-docker-certs:/certs/client \
  --volume jenkins-data:/var/jenkins_home \
  --publish 2376:2376 docker:dind
```

Build the custom Jenkins image called 'blueocean' using the Dockerfile
in jenkins/
```
docker build -t myjenkins-blueocean:1.1 ./jenkins
```

Finally create a container of blueocean
```
docker run --name jenkins-blueocean --rm --detach \
  --network jenkins --env DOCKER_HOST=tcp://docker:2376 \
  --env DOCKER_CERT_PATH=/certs/client --env DOCKER_TLS_VERIFY=1 \
  --publish 8080:8080 --publish 50000:50000 \
  --volume jenkins-data:/var/jenkins_home \
  --volume jenkins-docker-certs:/certs/client:ro \
  myjenkins-blueocean:1.1
```

Interesting Note
------------------
Apparently the Jenkins test setup can be mobilised using this docker compose file
below
```
version: '3'services:
  jenkins:
    image: jenkins/jenkins:alpine
    container_name: jenkins
    ports:
      - 9090:8080
      - 50000:50000
    volumes:
      - ./docker_volumes/jenkins:/var/jenkins_home
```

### STUB
STUB Paragraph
- **STUB**: Stub
- **STUB**: 

**Note STUB**: Useless text here, will be updated later 

Jenkins
--------------------------

### Jenkins & Docker-in-Docker
https://medium.com/@davidessienshare/how-to-run-jenkins-in-a-docker-container-e782647b3259

https://tutorials.releaseworksacademy.com/learn/the-simple-way-to-run-docker-in-docker-for-ci

This runs Jenkins in a container which will then create docker containers inside of the jenkins docker. It can optionally also create docker containers outside by binding the docker socket (-v /var/run/docker.sock:/var/run/docker.sock), however, running Jenkins in a container proved to be a bit challenging so this unfinished.

Jenkins Username: admin

Jenkins Password: 
```
d00e252af1f943878b88444e8395fe0c
```

You might end up with issues regarding the permissions with volume directories
 I solved this by:
```
mkdir jenkins-home; chmod -R 777 jenkins-home
```

Then I launched the jenkins container using:
```
docker run -dit --name jenkins -p 3002:8080 -p 3003:50000 -v /home/ziad/jenkins-home:/var/jenkins_home c5eca72556f6
```


### Jenkins Standalone

### External URL
```
api.chibois.net
```

This is easier but less portable 
Execute the following:
```
wget -q -O - https://pkg.jenkins.io/debian/jenkins.io.key | sudo apt-key add -
sudo sh -c 'echo deb https://pkg.jenkins.io/debian binary/ > \
    /etc/apt/sources.list.d/jenkins.list'
sudo apt-get update
sudo apt-get install jenkins
```

I then got an error with Failed to start LSB: Start Jenkins at boot time.
This was fixed by installing default-jre default-jdk and ensuring that
Java is installed. Strangely, it requires Java 8 while insisting that support
for Java 11 is limited but works with Java 11

Now to select the HTTP Port edit the file: /etc/default/jenkins

Jenkins username: 
```
admin
``` 
Jenkins password: 
```
578c964c35e9457a991293f54b0f34a2
```

Encountered a permission denied error while trying to build a docker container in the pipeline which was solved by using these commands:
```
sudo usermod -aG docker jenkins
sudo systemctl restart jenkins
```
The problem with this is that the docker group has root permissions so this is risky, I will look into a better solution later

Running a CockroachDB container
--------------------------

Create a network:   
```
docker network create -d bridge roachnet
```  
    
Start a container (assuming that you've got cockroachdb/cockroach image):  
```
docker run -d \
--name=roach3 \
--hostname=roach3 \
--net=roachnet \
-p 26257:26257 -p 8080:8080 \
-v "roach3:/cockroach/cockroach-data" \
cockroachdb/cockroach start \
--insecure \
--join=roach1,roach2,roach3
```

Initialise the cluster (we aren't using multiple databases yet but keep this command in mind):  
```
docker exec -it roach3 ./cockroach init --insecure
```

You can access the cockroach SQL Shell:
```
docker exec -it roach3 ./cockroach sql --insecure
```

When execution is done be sure to remove any docker volumes (this is for testing purposes and will be used in the pipeline):  
```
docker volume rm roach3
```

Swaggerio Docs
--------------------------
```
swagger generate spec -o ./swagger.json
```
```
swagger serve ./swagger.json
```
