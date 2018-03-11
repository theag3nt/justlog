pipeline {
  agent any
  stages {
    stage('Build') {
      steps {
        sh '''docker build .
docker tag gempbotgo_gempbotgo gempir/gempbotgo
docker push gempir/gempbotgo'''
      }
    }
    stage('Deploy') {
      steps {
        sh 'cp ./prod.yml /home/gempir/gempbotgo'
        sh 'cd /home/gempir/gempbotgo'
        sh 'docker-compose -f prod.yml pull'
        sh 'docker-compose -f prod.yml up -d'
      }
    }
  }
}