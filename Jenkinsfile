pipeline {
  agent {
    dockerfile {
      filename 'cidr_convert_api/go/Dockerfile'
    }

  }
  stages {
    stage('error') {
      steps {
        sh '''ls
pwd'''
      }
    }

  }
  environment {
    dir = 'cidr_convert_api/go/'
  }
}