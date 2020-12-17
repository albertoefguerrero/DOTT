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
'''
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
          args '--network host -e SONAR_HOST_URL="http://3.22.117.110/" -e SONAR_LOGIN="f902cca8ced0b0c3e740b60954b5979b3f117967" -v "/var/jenkins_home/workspace/DOTT_master/cidr_convert_api/go/"'
        }

      }
      steps {
        sh 'sonar-scanner'
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
    && goop install
'''
          warnError(message: 'Failed :( ') {
            sh 'goop go test'
          }

        }

      }
    }

    stage('Deploy') {
      steps {
        sh 'echo "hello deploy!"'
      }
    }

  }
}