pipeline {
    agent any
    stages {
        stage('build') {
            steps {
                echo 'Build Project:'
                sh 'bash build-linux.sh'
            }
        }
    }
}
