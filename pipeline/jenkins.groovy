// Declarative Jenkins Pipeline for a multi-platform, parameterized build
// of the telebot service. The developer can pick the target OS/ARCH and
// toggle tests/lint, or just run with the defaults.
//
// Docs: https://www.jenkins.io/doc/book/pipeline/

pipeline {
    agent any

    parameters {
        choice(
            name: 'OS',
            choices: ['linux', 'darwin', 'windows'],
            description: 'Target operating system'
        )
        choice(
            name: 'ARCH',
            choices: ['amd64', 'arm64'],
            description: 'Target architecture'
        )
        booleanParam(
            name: 'SKIP_TESTS',
            defaultValue: false,
            description: 'Skip running tests'
        )
        booleanParam(
            name: 'SKIP_LINT',
            defaultValue: false,
            description: 'Skip running linter'
        )
    }

    environment {
        REPO_URL   = 'https://github.com/vladyslav-ops/telebot.git'
        BRANCH     = 'develop'
        TARGETOS   = "${params.OS}"
        TARGETARCH = "${params.ARCH}"
    }

    stages {
        stage('Clone') {
            steps {
                git branch: "${BRANCH}", url: "${REPO_URL}"
            }
        }

        stage('Lint') {
            when { expression { return !params.SKIP_LINT } }
            steps {
                sh 'make lint'
            }
        }

        stage('Test') {
            when { expression { return !params.SKIP_TESTS } }
            steps {
                sh 'make test'
            }
        }

        stage('Build') {
            steps {
                sh 'make build TARGETOS=${TARGETOS} TARGETARCH=${TARGETARCH}'
            }
        }

        stage('Image') {
            steps {
                sh 'make image TARGETOS=${TARGETOS} TARGETARCH=${TARGETARCH}'
            }
        }

        stage('Push') {
            steps {
                sh 'make push TARGETOS=${TARGETOS} TARGETARCH=${TARGETARCH}'
            }
        }
    }

    post {
        always {
            // Remove the locally built image to keep the agent clean.
            sh 'make clean TARGETOS=${TARGETOS} TARGETARCH=${TARGETARCH} || true'
        }
    }
}
