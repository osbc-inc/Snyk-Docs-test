# Kubernetes 파일

코드형 Snyk 인프라를 사용하면 CLI로 구성 파일을 테스트할 수 있습니다. 쿠버네티스용 Snyk의 Infrastructure는 다음을 지원합니다:

* 배포(Deployments), 파드(Pods) 및 서비스(Services)
* 크론잡(CronJobs), 잡(Jobs), 스테이트풀세트(StatefulSet), 레플리카세트(ReplicaSet), 데몬세트(DaemonSet) 및 레플리케이션 컨트롤러(ReplicationController)

지정된 파일에서 문제를 테스트하기 위해 다음 Snyk CLI 명령을 사용합니다:

```
snyk iac test
```

예시:

```
snyk iac test deploy.yaml
```

또한 다음과 같이 파일 이름들을 이어서 추가하여 여러 파일을 지정할 수도 있습니다:

```
snyk iac test file-1.yaml file-2.yaml
```

Snyk CLI를 사용하여 헬름 차트를 스캔하는 단계는 [헬름 차트](helm-charts.md)를 참조하십시오.

Snyk CLI를 사용하여 Kustomize 템플릿을 스캔하는 단계는 [Kustomize 파일](kustomize-files.md)을 참조하십시오.
