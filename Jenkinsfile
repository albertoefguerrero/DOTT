pipeline {
  agent none
  stage ('Build') {
    agent {
      dockerfile {
        filename 'Dockerfile'
        dir 'cidr_convert_api/go/'
      }
      steps {
        sh '''ls
              pwd'''
      }
    }
  }
}
