pipeline {
    agent {
        docker { image 'golang:alpine' }
    }

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
        stage('Deploy to Artifactory'){
            steps{
                script{
                    def server = Artifactory.newServer url: 'http://10.138.100.36:8081', username: 'admin', password: 'Password1'
                    def uploadSpec = """{
                        "files": [
                            {
                                "pattern": "./JenkinsGHTesting",
                                "target": "DEV/"
                            }
                        ]
                    }"""
                    server.upload(uploadSpec)
                }
            }
        }
    }
}