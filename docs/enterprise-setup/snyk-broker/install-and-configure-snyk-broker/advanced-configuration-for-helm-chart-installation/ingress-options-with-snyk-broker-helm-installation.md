# Snyk Broker Helm 설치 시 Ingress 옵션

Helm을 사용하여 브로커를 설정할 때 `brokerClientUrl` 매개변수를 구성해야 할 수 있습니다. 이 매개변수는 SCM에 연결하거나 컨테이너 레지스트리에 연결할 경우 PR 확인을 활성화합니다.

이 연결이 발생하려면 SCM 또는 컨테이너 레지스트리에서 브로커로의 수신 연결이 있어야 합니다. Kubernetes는 이러한 수신 연결을 Ingress를 통해 관리합니다.

Ingress는 Kubernetes 클러스터의 특정 서비스로 들어오는 네트워크 트래픽을 경로 지정하는 방법입니다. Ingress 설정에 대한 자세한 정보는 [Kubernetes Ingress 가이드](https://kubernetes.io/docs/concepts/services-networking/ingress/)를 참조하십시오.

Ingress 트래픽에 대해 사용 가능한 두 가지 옵션이 있습니다. 기본적으로 팟은 클러스터 외부에서 접근할 수 없습니다.

**로드 밸런서**를 활성화하려면, `--set service.<service-type>=LoadBalancer`를 추가하십시오. 허용되는 값은 `brokerType`, `crType`, `caType`입니다. 서비스 유형은 실행 중인 브로커 유형을 나타냅니다.

Github의 예시:

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=github-com \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set scmToken=<ENTER_REPO_TOKEN> \
             --set brokerClientUrl=<ENTER_BROKER_CLIENT_URL>:<ENTER_BROKER_CLIENT_PORT> \
             --set service.brokerType=LoadBalancer \
             -n snyk-broker --create-namespace
```

Ingress를 추가하려면 `values.yaml` 파일에서 활성화하고, 값 파일의 예시 값을 따르는 관련 구성 매개변수를 추가하십시오. 이 작업을 위해 클러스터에는 Ingress 컨트롤러가 올바르게 구성되어 있어야 합니다.
