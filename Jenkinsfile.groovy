pipeline {
    agent any
    stages {
        stage('build') {
            steps {
                echo 'Build Project:'
                sh 'ls -l'
                sh 'bash build-linux.sh'
            }
        }
    }
}
