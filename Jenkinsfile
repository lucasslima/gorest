pipeline {
    agent {
        docker { image 'golang:alpine' }
    }
    stages {
        stage('Build') {
            steps {
                sh 'go build my_simple_http.go'
            }
        }
    }
}