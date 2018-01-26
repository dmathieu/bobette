# Bobette [![CircleCI](https://circleci.com/gh/dmathieu/bobette.svg?style=svg)](https://circleci.com/gh/dmathieu/bobette)

Bobette, as in Bob the Builder's wife is a kubernetes-aware service used to perform build actions inside a pod.

I need to build docker images to run on my [Raspberry PI kubernetes cluster](https://github.com/dmathieu/kubepi). Unfortunately, the PI is the only ARM machine I have around and I don't want to have to ssh into it every time I need to build a new image.  
To solve this, bobette will run a pod, executing any commands configured for a repository.

## Installation

You should first install and setup Go on your machine.

```
go get -u github.com/dmathieu/bobette
```

You should then have a `bobette` binary available in your PATH

## Configuration

In order to configure a repository to use bobette, create a `bobette.yml` file in the repository's root folder.

```
repository: https://github.com/kubernetes/kubernetes
commands:
  - docker build -t kubernetes .
  - docker push kubernetes
```

Push and pull that code to your repository's master branch, and run the `bobette` binary while you're at the root folder of your repository.  
Doing so will start a new pod that will:

* Pull your repository
* Run the commands specified

And that's all!

### Private repositories

You can specify authentication for your GIT repository. All you need for that is a `user:password` combination with access to your repository.  
Set the authentication like this:

```
bobette env set repo_auth=<username:password>
```

See the [Config Vars](#config-vars) section below for more details on this.

## Config Vars

You can set any config vars you wish (for example, to log into the docker registry):

```
bobette env set docker_login=<something> foo=bar hello=world
```

Doing this will set a single secret for your repository containing 3 values.  
All those values will be available as environment variables on your app.
