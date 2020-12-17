pipeline {
  agent none
  stages {
    stage('Build') {
      agent {
        docker {
          image 'golang:1.15-alpine'
        }

      }
      steps {
        dir(path: 'cidr_convert_api/go/') {
          sh '''ls
pwd
apk add --update git
apk add build-base
go get github.com/karmakaze/goop \\
    && go get github.com/gorilla/mux \\
    && go get github.com/stretchr/testify/assert \\
    && goop install

go build'''
          warnError(message: 'Failed to build :( ') {
            sh 'go build'
          }

        }

      }
    }

    stage('Sonar Qube') {
      agent {
        docker {
          image 'sonarsource/sonar-scanner-cli'
          args '--network host -e SONAR_HOST_URL="http://3.22.117.110/" -e SONAR_LOGIN="99a7536d3c88fc79e7f1dd189f99b4cf59926cc6" -v "/var/jenkins_home/workspace/DOTT_master/cidr_convert_api/go/"'
        }

      }
      steps {
        sh '''#sonar-scanner \\
#  -Dsonar.projectKey=test-key1 \\
#  -Dsonar.sources=. \\
#  -Dsonar.host.url=http://3.22.117.110 \\
#  -Dsonar.login=test'''
      }
    }

    stage('Tests') {
      agent {
        docker {
          image 'golang:1.15-alpine'
        }

      }
      steps {
        dir(path: 'cidr_convert_api/go/') {
          sh '''apk add --update git
apk add build-base

go get github.com/karmakaze/goop \\
    && go get github.com/gorilla/mux \\
    && go get github.com/stretchr/testify/assert \\
    && go get golang.org/x/lint/golint \\
    && goop install
'''
          warnError(message: 'Failed :( ') {
            sh 'go fmt'
            sh '''golint api.go convert.go convert_test.go
'''
            sh 'goop go test'
          }

        }

      }
    }

    stage('Deploy') {
      agent {
        docker {
          image 'golang:1.15-alpine'
        }

      }
      steps {
        dir(path: 'cidr_convert_api/go/') {
          sh '''apk add --update git
apk add build-base
go get github.com/karmakaze/goop \\
    && go get github.com/gorilla/mux \\
    && go get github.com/stretchr/testify/assert \\
    && goop install
'''
          sh 'echo "hello world!"'
        }

      }
    }

  }
}