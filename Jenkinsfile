pipeline {
    agent { 
      dockerfile {
          filename 'Dockerfile'
          dir 'cidr_convert_api/go/'
          label 'dott-go'
        }
    }
    stages {
        stage('Test') {
            steps {
                sh '''pwd
                ls'''
                sh 'go version'
            }
        }
    }
}
