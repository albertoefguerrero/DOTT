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
goop go test'''
        }

        sh '''cd cidr_convert_api/go/
                ls
                which go
                #goop go test'''
      }
    }

    stage('Sonar Qube') {
      steps {
        sh '''echo "hello world!"
'''
      }
    }

    stage('Run') {
      steps {
        sh '''cd cidr_convert_api/go/
goop go run api.go convert.go'''
      }
    }

  }
}