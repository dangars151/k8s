pipeline {

  environment {
    DOCKERHUB_CREDENTIALS = credentials('docker-account')
  }

  agent {
    kubernetes {
      yaml '''
      apiVersion: v1
      kind: Pod
      spec:
        serviceAccountName: jenkins-admin
        dnsConfig:
          nameservers:
            - 8.8.8.8
        containers:
        - name: docker
          image: docker:latest
          command:
          - cat
          tty: true
          volumeMounts:
          - mountPath: /var/run/docker.sock
            name: docker-sock
        - name: kubectl
          image: bitnami/kubectl:latest
          command:
          - cat
          tty: true
        securityContext:
          runAsUser: 0
          runAsGroup: 0
        imagePullSecrets:
          - name: regcred
        volumes:
          - name: docker-sock
            hostPath:
              path: /var/run/docker.sock
            '''
    }
  }

  stages {
    
    stage('Unit Test') {
      steps {
        sh 'echo Unit Test' 
      }
    }

    stage('Build image') {
      steps{
        container('docker') {
          script {
            sh 'docker build --network=host -t dangars151/go-todo .'
          }
        }
      }
    }

    stage('Pushing Image') {
      steps{
        container('docker') {
          script {
            sh 'echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin'
            sh 'docker tag dangars151/go-todo 12345688/go-todo'
            sh 'docker push 12345688/go-todo:latest'
          }
        }
      }
    }

    stage('Deploying App to Kubernetes') {
      steps {
        container('kubectl') {
          withCredentials([file(credentialsId: 'kube-config', variable: 'TMPKUBECONFIG')]) {
            sh "cp \$TMPKUBECONFIG /.kube/config"
            sh 'kubectl apply -f deployment.yaml'
            sh 'kubectl apply -f service.yaml'
          }
        }
      }
    }
  }
}