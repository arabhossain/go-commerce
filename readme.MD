
# Go Commerce

Test project while learning GoLang



## Jenkins pipeline


```bash
  pipeline {
    agent any
    tools {
        go '1.16.4'
    }
    stages {
        stage("Pull Latest Code"){
            steps{
                 echo 'LATEST CODE DOWNLOADING STARTED'
                 git branch: 'main', url: 'https://github.com/arabhossain/go-commerce.git'
            }
        }
        stage("Install Dependency") {
            steps {
                echo 'BUILD EXECUTION STARTED'
                sh 'go mod tidy'
            }
        }
        stage("Build") {
            steps {
                echo 'BUILD EXECUTION STARTED'
                sh 'go build main.go'
            }
        }
        stage("Test") {
            steps {
                echo 'TEST EXECUTION STARTED'
                sh 'go test -v ./tests/...'
            }
        }
        stage("Run") {
            steps {
                echo 'BUILD EXECUTION STARTED'
                sh 'nohup go run main.go 2>&1 &'
            }
        }
    }
}
```
