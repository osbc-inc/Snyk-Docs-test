# 주석이 달린 가져오기(Annotated import)

{% hint style="warning" %}
**폐기됨**

주석이 달린 가져오기(Annotated import)는 더 이상 유지되지 않으며 폐기되었습니다.
{% endhint %}

Snyk 관리자가 Kubernetes 클러스터에 Snyk 컨트롤러를 설치한 후에는 스캔을 위한 워크로드를 추가할 수 있습니다. Kubernetes 협력자는 클러스터에서 워크로드를 표시하여 Snyk에 자동으로 추가할 수 있습니다.

## 자동으로 추가, 업데이트 및 제거되는 워크로드

Snyk와 클러스터 간의 통합을 구성한 후에는 작업로드에 주석을 달아서 해당 프로젝트를 스캔할 때 Snyk에 자동으로 추가할 수 있습니다.

{% hint style="info" %}
주석이 달린 가져오기는 이미지 자체가 변경될 때(이미지 변경으로 인한 작업로드 재스캔) 또는 작업로드 세부 정보가 변경될 때(작업로드의 새 리비전을 만드는 경우) 발생합니다. 작업로드의 주석을 변경해도 작업로드 변경을 일으키지는 않습니다.

작업로드가 `snyk monitor`로 스캔된 후에만 주석이 달리면, 주석은 극적인 변경이 발생하여 완전한 재스캔을 일으킬 때까지 인식되지 않습니다. 재스캔을 강제하는 한 가지 방법은 `snyk monitor` 팟을 종료하는 것입니다.
{% endhint %}

다음과 같은 작업로드 유형에 어노테이션을 추가할 수 있습니다:

* 배포 (Deployments)
* 레플리카셋 (ReplicaSets)
* 데몬셋 (DaemonSets)
* 스테이트풀셋 (StatefulSets)
* 작업 (Jobs)
* 크론작업 (CronJobs)
* 복제 컨트롤러 (ReplicationControllers)
* 파드 (Pods)

작업을 수행하려면:

1. Snyk에서 관리하려는 관련 그룹 및 조직으로 이동합니다.
2. **설정** > **일반**으로 이동합니다.
3. **조직 ID**를 복사합니다.
4. 작업로드에 `orgs.k8s.snyk.io/v1` 키로 주석을 추가하고 값으로 조직 ID를 쉼표로 구분된 목록으로 추가합니다.

여러 조직에 추가하려면 단일 작업로드에 주석을 달 수도 있습니다.

1. Snyk 컨트롤러는 작업로드의 변경 사항을 자동으로 감지하여 해당 작업로드가 Snyk 프로젝트로 자동으로 가져오도록 보장합니다.

예시:

조직으로 자동으로 가져오도록 주석이 달린 배포 YAML 파일:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-app-deployment
  annotations:
    orgs.k8s.snyk.io/v1: cacb791e-07cc-4b10-b4be-64de19f532f1
spec:
  template:
    spec:
      containers:
      …
```

여러 조직에 주석을 달려면 쉼표로 구분된 목록을 사용합니다.

2. 가져온 후, 프로젝트는 Snyk 조직에 유지되며 주석을 제거해도 유지됩니다. Snyk에서 프로젝트를 제거하려면 주석을 삭제하고 Snyk UI를 사용하여 프로젝트를 삭제하거나 API 엔드포인트 [프로젝트 삭제](../../../snyk-api/reference/projects-v1.md#org-orgid-project-projectid-2)를 사용해야 합니다.
