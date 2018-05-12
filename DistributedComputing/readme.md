# A workspace for understanding Go with distributed computing.


## Tools
1. [Glide](https://glide.sh/) - A more advanced package management tool for Go - [Github: Glide](https://github.com/Masterminds/glide)
    - Installation (and sample project):
        - `curl https://glide.sh/get | sh`
        - `mkdir new-project && cd new-project`
        - `glide create`
        - `glide get github.com/last-ent/skelgor`  
            - a skeleton go helper project
        - `glide install`
            - in case any dependencies or configuration were manually added
        - `glide up`
            - update dependencies to latest versions
        - `tree`
2. Docker
    1. Working with images
        1. `docker --version`
        2. `docker info` - more in depth docker info
        3. `docker run docker/whalesay cowsay` - run a docker image
            - `docker pull docker/whalesay & docker run docker/whalesay cowsay` - Alternatively, this is the same command as (3)
        4. `docker container ls --all` - list of all docker containers
        5. Removing docker images:
            - `docker rmi --force 'docker images -q -f dangling=true'` - remove dangling, "unused" docker image(s)
            - `docker images prune` - remove unused docker images
            - `docker rm -v imagename` - remove a specific docker image
    2. Dockerfile(s)
        1. A basic Docker example, with a Go Program:
            1. [Example of a basic GoLang Dockerfile](./Docker/Basic/Dockerfile)
            2. [A basic Go Program](./Docker/Basic/main.go)
            3. Commands to install, run our docker image:
                - `docker build . -t imagefriendlyname` - build the docker image
                - `docker run imagefriendlyname` - run the docker image program
                - `docker run -e NAME=Bill imagefriendlyname` - edit the value in the NAME environment variable

## References and Resources
1. Distributed Computing with Go - Practical Concurrency and Parallelism for Go Applications [Book](https://www.amazon.com/Distributed-Computing-concurrency-parallelism-applications/dp/1787125386)
    1. [Color Images for Go Concepts](https://www.packtpub.com/sites/default/files/downloads/DistributedComputingwithGo_ColorImages.pdf)
    2. Source Code Repositories:
        - [PacktPublishing](https://github.com/PacktPublishing/Distributed-Computing-with-Go)
        - [Book Last-End](http://github.com/last-end/distributed-go)
2. Designing Distributed Systems: Patterns and Paradigms for Scalable, Reliable Services [Book](https://www.amazon.com/Designing-Distributed-Systems-Patterns-Paradigms-ebook/dp/B07CQ9GZ8R/ref=mt_kindle?_encoding=UTF8&me=)
