pipeline {
    agent any 
    stages {
        stage('Build') {
            steps {
                echo "Running build in branch ${env.BRANCH_NAME} ...."
                echo "Adding new piece of code"
                echo "update on GitHub"
            }
        }
        stage('Test') {
            steps {
                echo "Running tests in branch ${env.BRANCH_NAME} ...."
                echo "update number 2 on GitHub"
            }
        }
    }
}
