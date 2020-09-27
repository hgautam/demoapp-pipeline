pipeline {
    agent any 
    stages {
        stage('Build') {
            steps {
                echo "Running build in branch ${env.BRANCH_NAME} ...."
                echo "GitHub webhook test"
            }
        }
        stage('Test') {
            steps {
                echo "Running tests in branch ${env.BRANCH_NAME} ...."               
            }
        }
    }
}
