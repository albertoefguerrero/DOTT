pipeline {
  agent {
    dockerfile {
      filename 'Dockerfile'
      dir 'cidr_convert_api/go/'
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
}
