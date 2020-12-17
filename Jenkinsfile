pipeline {
  agent none
  stages {
    stage('Back-end') {
      agent {
        docker {
          image 'golang:1.15-alpine'
        }

      }
      steps {
        sh '''ls
pwd
cd cidr_convert_api/go/
apk add --update git
apk add build-base
go get github.com/karmakaze/goop \\
    && go get github.com/gorilla/mux \\
    && go get github.com/stretchr/testify/assert \\
    && goop install
'''
      }
    }

    stage('Front-end') {
      agent {
        docker {
          image 'node:14-alpine'
        }

      }
      steps {
        sh 'node --version'
      }
    }

  }
}