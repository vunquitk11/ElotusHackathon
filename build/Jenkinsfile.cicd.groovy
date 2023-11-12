pipeline {
  agent any
  stages {
    stage("ci:init") {
      steps {
        echo 'Hello init!'
      }
    }
    stage("ci:check images exist") {
        steps {
            echo 'Hello check images exist!'
        }
    }
    stage("ci:test & build") {
      stages {
        stage("test & Build") {
          parallel {
            stage("api") { steps { sh """
              make api-gen-mocks
              make api-test
              make api-build-binaries
              mkdir api/binaries
              ${DOCKER_BIN} cp ${PROJECT_NAME}-alpine-${CONTAINER_SUFFIX}:api/binaries api/.
              make -f build/Makefile api-report-coverage || exit 0
            """ } }
          }
        }
      }
      post {
        always { sh "make teardown" }
      }
    }
  }
}

def buildImage(String docker_file_name, String tag) {
  sh "${DOCKER_BIN} build -f ./build/${docker_file_name}.Dockerfile -t ${tag} ."
}

def pushImage(String tag) {
  sh "${DOCKER_BIN} push ${tag}"
}
