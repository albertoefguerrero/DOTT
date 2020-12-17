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
        sh '''apk add --update git
apk add build-base
go get github.com/karmakaze/goop \\
    && go get github.com/gorilla/mux \\
    && go get github.com/stretchr/testify/assert \\
    && goop install

ls
pwd'''
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