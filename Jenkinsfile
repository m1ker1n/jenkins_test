/* Requires the Docker Pipeline plugin */
pipeline {
    agent any

    stages {
        stage('build') {
            steps {
                sh "docker build -t 127.0.0.1:5001/go-echo:latest --cache-from 127.0.0.1:5001/go-echo:latest ."
            }
        }

        stage('Deploy') {
            steps {
                sh "docker run \
                        -e SECRET=zhopa \
                        -p 8080:8080 \
                        127.0.0.1:5001/go-echo:latest"
            }
        }
    }
}