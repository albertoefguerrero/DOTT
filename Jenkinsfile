pipeline {
  agent none
  stages {
    stage('Build') {
      agent {
        docker {
          image 'golang:1.15-alpine'
          args '-v $HOME/jenkins:/app'
        }

      }
      steps {
        dir(path: 'cidr_convert_api/go/') {
          sh '''ls ~/
pwd
apk add --update git
apk add build-base
go get github.com/karmakaze/goop \\
    && go get github.com/gorilla/mux \\
    && go get github.com/stretchr/testify/assert \\
    && goop install

go build'''
          warnError(message: 'Failed to build :( ') {
            sh '''ls -lah
go env GOOS GOARCH
env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o api
ls -lah
 '''
          }

          sh '''tar -cvzf artifact.tar.gz api
cp artifact.tar.gz /app
'''
        }

      }
    }

    stage('Tests') {
      agent {
        docker {
          image 'golang:1.15-alpine'
          args '-v $HOME/jenkins:/app'
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
    && go get github.com/jstemmer/go-junit-report \\
    && goop install
'''
          warnError(message: 'Failed :( ') {
            sh 'go fmt'
            sh '''golint api.go convert.go convert_test.go
'''
            sh '''go test
go test -v 2>&1 | go-junit-report > report.xml
go test -v -bench . -count 5 2>&1 | go-junit-report > report-benchmark.xml'''
            sh '''go test -coverprofile=coverage.out
cp coverage.out /app'''
          }

          junit(testResults: '*.xml', allowEmptyResults: true)
        }

      }
    }

    stage('Sonar Qube') {
      agent {
        docker {
          image 'sonarsource/sonar-scanner-cli'
          args '''--network host -e SONAR_HOST_URL="http://albertoefg1c.mylabserver.com" -e SONAR_LOGIN="${env.TOKEN}"
-v $HOME/jenkins:/app'''
        }

      }
      steps {
        sh '''pwd
ls cidr_convert_api/go/'''
        dir(path: 'cidr_convert_api/go/') {
          withCredentials(bindings: [string(credentialsId: 'sonarlogin', variable: 'TOKEN')]) {
            sh '''pwd
ls
cp /app/coverage.out .'''
            sh '''sonar-scanner \\
-Dsonar.projectKey=newjenkins \\
-Dsonar.sources=.   \\
-Dsonar.host.url=http://albertoefg1c.mylabserver.com \\
-Dsonar.login=$TOKEN \\
-Dsonar.go.coverage.reportPaths=/app/coverage.out
'''
          }

        }

      }
    }

    stage('Deploy') {
      agent {
        docker {
          image 'golang:1.15-alpine'
          args '-v $HOME/jenkins:/app'
        }

      }
      steps {
        dir(path: 'cidr_convert_api/go/') {
          withCredentials(bindings: [string(credentialsId: 'github-token', variable: 'GTOKEN')]) {
            sh '''apk add --update git
go get github.com/github-release/github-release'''
            sh '''cp /app/artifact.tar.gz .
ls
'''
            sh '''echo ${BUILD_NUMBER} and $BUILD_NUMBER 
echo "Exporting token to enable github-release tool"
export GITHUB_TOKEN=$GTOKEN

#echo "Deleting release from github before creating new one"
#github-release delete --user ${GITHUB_USER} --repo ${GITHUB_REPO} --tag ${BUILD_NUMBER}

echo "Creating a new release in github"
github-release release --user ${GITHUB_USER} --repo ${GITHUB_REPO} --tag ${BUILD_NUMBER} --name "api-${BUILD_NUMBER}"

echo "Uploading the artifacts into github"
github-release upload --user ${GITHUB_USER} --repo ${GITHUB_REPO} --tag ${BUILD_NUMBER} --name "${PROJECT_NAME}-${BUILD_NUMBER}.tar.gz" --file artifact.tar.gz'''
          }

        }

      }
    }

  }
  environment {
    TOKEN = credentials('sonarlogin')
    GITHUB_USER = 'albertoefguerrero'
    GITHUB_REPO = 'DOTT'
    PROJECT_NAME = 'go-api'
  }
}