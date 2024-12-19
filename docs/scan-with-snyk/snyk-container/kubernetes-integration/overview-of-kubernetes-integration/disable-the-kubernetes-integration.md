# Kubernetes 통합 비활성화

Kubernetes 통합을 비활성화하려면 다음 단계를 따르십시오:

1. 클러스터에서 Snyk Controller를 제거합니다.

이를 위해 Helm을 사용할 수 있습니다. Helm을 사용하면 릴리스 이름을 결정해야 합니다. 예를 들어:

`$ helm ls --short`

`snyk-monitor`

그런 다음 릴리스를 삭제할 수 있습니다: `$ helm delete snyk-monitor`

2. 조직 통합 설정 페이지에서 **연결 끊기**를 클릭합니다. 이렇게 하면 새로운 작업 부하를 가져오고 동기화하는 데 사용된 자격 증명이 무효화됩니다.

&#x20;Helm을 사용하지 않는 경우 `snyk-monitor` 네임스페이스를 제거하십시오.&#x20;