/* Requires the Docker Pipeline plugin */
pipeline {
    agent any

    stages {
        stage('build') {
            steps {
                sh "docker build -t go-echo:latest --cache-from go-echo:latest ."
            }
        }

        stage('Deploy') {
            steps {
                sh "docker run \ 
                        -d \
                        -e SECRET=zhopa \
                        -p 8080:8080 \
                        go-echo:latest"
            }
        }
    }
}