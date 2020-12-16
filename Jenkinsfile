pipeline {
  agent {
    dockerfile {
      filename 'Dockerfile'
    }

  }
  stages {
    stage('') {
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