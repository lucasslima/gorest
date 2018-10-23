pipeline {
    agent {
        docker { image 'golang:alpine' }
        
    }
    environment {
        GO_ENABLED = 0
    }
    stages {
        stage('Build') {
            steps {
                sh 'go build'
            }
        }
        stage('Test'){ 
            steps { 
                sh 'go test'
            }
        }
        stage('Deploy to Artifactory'){
            steps{
                script{
                    def server = Artifactory.newServer url: 'http://10.138.100.36:8081', username: 'admin', password: 'Password1'
                    def uploadSpec = """{
                        "files": [
                            {
                                "pattern": "my_simple_http",
                                "target": "DEV/my_simple_http/"
                            }
                        ]
                    }"""
                    server.upload(uploadSpec)
                }
            }
        }
    }
}