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
        sh '''cd cidr_convert_api/go/
sonar-scanner \\
  -Dsonar.projectKey=test-key1 \\
  -Dsonar.sources=. \\
  -Dsonar.host.url=http://3.138.247.153:9090 \\
  -Dsonar.login=f902cca8ced0b0c3e740b60954b5979b3f117967
'''
      }
    }

    stage('Run') {
      steps {
        sh '''cd cidr_convert_api/go/
#goop go run api.go convert.go'''
      }
    }

  }
}