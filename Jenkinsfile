pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                sh 'go build -o JenkinsGHTesting'
            }
        }
        stage('Test'){
            environment {
              CGO_ENABLED = 0
            }
            steps { 
                sh 'go test'
            }
        }
    }
}