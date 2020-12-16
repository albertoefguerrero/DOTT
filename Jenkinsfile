pipeline {
  agent {
    dockerfile {
      filename 'Dockerfile'
      dir 'cidr_convert_api/go/'
    }

  }
  stages {
    stage('Test') {
      steps {
        warnError(message: 'Failing :( ') {
          sh '''cd cidr_convert_api/go/
go goop test'''
        }

        sh '''cd cidr_convert_api/go/
                ls
                which go
                #goop go test'''
      }
    }

  }
}