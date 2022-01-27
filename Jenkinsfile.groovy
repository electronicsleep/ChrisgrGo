pipeline {
    agent any
    stages {
        stage('build') {
            steps {
                echo 'Step: Build'
                sh 'bash build-linux.sh'
                sh 'ls -ltra'
            }
        }
    }
}
