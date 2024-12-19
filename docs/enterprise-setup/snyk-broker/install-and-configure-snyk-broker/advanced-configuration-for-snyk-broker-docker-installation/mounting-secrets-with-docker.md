# 도커로 시크릿 마운트하기

가끔은 민감한 구성 요소, GitHub 또는 Snyk 토큰과 같은 것을 환경 변수가 아닌 파일에서로드해야 할 수 있습니다. 브로커는 [dotenv](https://www.npmjs.com/package/dotenv)를 사용하여 구성을로드하므로 프로세스는 비교적 간단합니다:

* `.env`라는 파일을 만들고 민감한 구성을 거기에 넣습니다.
* 이 파일을 마운트합니다. 예를 들어 [Kubernetes 시크릿](https://kubernetes.io/docs/tasks/inject-data-application/distribute-credentials-secure/#create-a-pod-that-has-access-to-the-secret-data-through-a-volume)을 사용하여 마운트합니다. 파일을 `/broker`와 같은 위치로 마운트합니다.
* 도커 이미지의 작업 디렉토리를 `/broker/`로 변경합니다. 이와 관련된 파일 예는 브로커 컨테이너의 $HOME/.env에 있습니다.