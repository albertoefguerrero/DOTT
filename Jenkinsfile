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
                sh '''pwd
                ls'''
                sh 'goop go test'
            }
        }
    }
}
