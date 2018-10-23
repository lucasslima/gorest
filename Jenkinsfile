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
                    def server = Artifactory.newServer url: 'http://10.138.100.36:8081/artifactory', username: 'admin', password: 'Password1'
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
        stage('Order Infra'){
            steps{
                script{
                def orderDefinition = """{
                    "group": "/api/v2/groups/2/",
                    "items": {
                        "deploy-items": [
                            {
                                "blueprint": "/api/v2/blueprints/17/",
                                "blueprint-items-arguments": {
                                    "build-item-Install OpenJDK 8": {},
                                    "build-item-TMPL_CB_CENTOS7": {
                                        "attributes": {
                                            "hostname": "java-8-lsousa-test",
                                            "quantity": 1
                                        },
                                        "environment": "/api/v2/environments/2/",
                                        "os-build": "/api/v2/os-builds/4",
                                        "parameters": {
                                            "application": "lsousa-java-8-bp-testing",
                                            "application-team": "lsousa",
                                            "compliance": "False",
                                            "sc-nic-0": "VLAN300-DEV.APP.1",
                                            "vm-size": "t3.small",
                                            "vmware-datastore": "CML-CMLSC4K-IAD-DEVQC",
                                            "vmware-disk-type": "VMware Template Default"
                                        }
                                    }
                                },
                                "resource-name": "Java 8 OpenJDK base",
                                "resource-parameters": {}
                            }
                        ]
                    },
                    "submit-now": "true"
                }"""
            }
            }
            
        }
    }
}