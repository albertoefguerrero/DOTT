pipeline {
    agent { 
      dockerfile {
          filename 'Dockerfile'
          dir 'cidr_convert_api/go/'
        }
    stages {
        stage('Test') {
            steps {
                sh 'ls'
                sh 'go version'
            }
        }
    }
}
