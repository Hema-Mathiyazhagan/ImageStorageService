# Image store service using golang and react
Image Storage Service is based on microservice architecture that runs with REST APIS. I automated swagger doc generation for the apis using swaggo framework.The apis perform the following functions:
1)	Create/Delete Image Album 
2)	Create/Delete Image in an Album
3)	Get a single Image in an Album
4)	Get All Images in an Album
# Usage
clone the repository using Http or ssh
## To run the backend
```
$ cd backend
$ mkdir image-db
$ go build
$ ./image-storage-service.exe
```
backend uses port 3333 for the operations
The code is made concurrent on the backend with the use of locks
The service is docker ready and deployed on kubernetes.
Kubernetes config files are in k8s folder.
## To run the frontend
```
$ cd my-app
$ npm start
```
front end uses port 3000 for the operations
### Docker
There are seperate Dockerfiles for frontend and backend
Dockerfile for backend is stored inside backend folder
Dockerfile for frontend is stored inside frontend folder
### k8s
k8s files are stored inside k8s folder.
There are seperate yaml files for running front end and backend services
### Testing
The service is unittested on the backend with GINKGO framework
```
$ cd backend
$ ginkgo
```
