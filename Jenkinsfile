pipeline {
  agent any
  stages {
    stage('Build') {
      steps {
        sh '''echo "Starting Pipeline"
cd cidr_convert_api/go/
pwd
ls
goop install
#goop go run api.go convert.go
go build api.go convert.go'''
      }
    }

    stage('Code Analysis') {
      steps {
        sh '''echo "Analyzing"
'''
      }
    }

    stage('Testing') {
      steps {
        sh 'echo "Testing"'
      }
    }

    stage('Deploy') {
      steps {
        sh 'echo "Deploying"'
      }
    }

  }
}