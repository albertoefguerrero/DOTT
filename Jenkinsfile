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
                sh '''cd cidr_convert_api/go/
                ls
                goop go test'''               
            }
        }
    }
}
