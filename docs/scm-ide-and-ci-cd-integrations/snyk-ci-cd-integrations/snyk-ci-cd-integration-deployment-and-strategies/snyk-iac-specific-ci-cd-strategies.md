# Snyk IaC 특정 CI/CD 전략

`Snyk` `Infrastructure as Code`를 파이프라인에서 구현하는 가장 좋은 방법은 스테이지의 일부로, 그리고 SCA 및 컨테이너 스테이지 이후에 포함시키는 것입니다.

`Snyk` `Infrastructure as Code`는 다음을 지원합니다:

* 배포, Pod 및 서비스
* CronJobs, Jobs, StatefulSet, ReplicaSet, DaemonSet 및 ReplicationController

자세한 내용은 [Snyk CLI를 사용하여 Kubernetes 파일을 테스트하십시오](https://docs.snyk.io/snyk-infrastructure-as-code/snyk-cli-for-infrastructure-as-code/test-your-kubernetes-files-with-our-cli-tool) 를 참조하십시오.