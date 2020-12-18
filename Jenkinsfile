pipeline {
  agent none
  environment {
  TOKEN = credentials("sonarlogin")  
}
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
            sh 'go build *.go'
          }

        }

      }
    }

    stage('Sonar Qube') {
      agent {
        docker { image 'sonarsource/sonar-scanner-cli'        
                args '--network host -e SONAR_HOST_URL="http://albertoefg1c.mylabserver.com" -e SONAR_LOGIN="${env.TOKEN}"'        
       }
      }
      steps {
        sh '''pwd
ls cidr_convert_api/go/'''
        dir(path: 'cidr_convert_api/go/') {        
          sh '''pwd
                ls'''
          sh '''sonar-scanner \
  -Dsonar.projectKey=newjenkins \
  -Dsonar.sources=. \
  -Dsonar.host.url=http://18.223.182.245 \
  -Dsonar.login="${env.TOKEN}"'''        
       }
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
