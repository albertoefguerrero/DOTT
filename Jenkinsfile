pipeline {
  agent none
  stages {
    stage('Build') {
      steps {
        sh '''echo "Starting Pipeline"
cd cidr_convert_api/go/
pwd
ls
goop install
#goop go run api.go convert.go
'''
      }
    }

    stage('Code Analysis') {
      agent { 
        dockerfile { 
          dir 'cidr_convert_api/go/' 
        }
        steps {
                sh 'go version'
            }
      }
    }

    stage('Testing') {
      steps {
        sh '''echo "Testing"
cd cidr_convert_api/go/
goop go test'''
      }
    }

    stage('Deploy') {
      steps {
        sh 'echo "Deploying"'
      }
    }

  }
}
