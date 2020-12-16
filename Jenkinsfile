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
        sh '''docker run --rm \\
-e SONAR_HOST_URL="http://albertoefg1c.mylabserver.com:9090" \\
-e SONAR_LOGIN="99a7536d3c88fc79e7f1dd189f99b4cf59926cc6" \\
-v "/var/jenkins_home/workspace/DOTT_master/cidr_convert_api/go/" \\
sonarsource/sonar-scanner-cli'''
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