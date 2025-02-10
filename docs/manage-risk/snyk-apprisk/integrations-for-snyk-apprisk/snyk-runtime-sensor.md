# Snyk 런타임 센서

{% hint style="info" %}
**릴리스 상태**

Snyk 애플리케이션 리스크 프로를 위한 Snyk 런타임 센서는 초안 버전입니다. 현재는 Snyk 엔터프라이즈 플랜과 Snyk 앱리스크 프로를 이용할 수 있습니다. 그룹에서 설정하려면 Snyk 계정 팀에 문의하십시오.
{% endhint %}

Snyk 런타임 센서는 쿠버네티스 클러스터 상의 배포를 감시하고 수집된 데이터를 Snyk로 전송하는 [쿠버네티스 데몬셋](https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/)입니다.

다음 [위험 요소](https://docs.snyk.io/manage-risk/prioritize-issues-for-fixing/assets-and-risk-factors-for-snyk-apprisk#risk-factors)가 Snyk 런타임 센서에서 보고됩니다: [배포됨](https://docs.snyk.io/manage-risk/prioritize-issues-for-fixing/assets-and-risk-factors-for-snyk-apprisk/risk-factor-deployed) 및 [로드된 패키지](https://docs.snyk.io/manage-risk/prioritize-issues-for-fixing/assets-and-risk-factors-for-snyk-apprisk/risk-factor-loaded-package).

{% hint style="info" %}
* Snyk 런타임 센서는 로드된 패키지 위험 요소를 애플리케이션 패키지에 대해서만 보고합니다. 다음 생태계가 지원됩니다:
  * Node.js
  * Java
  * Python
  * Go
  * .NET
* 대부분의 패키지는 애플리케이션 시작 시에 메모리에 로드되므로 Pod가 초기화될 때만 보고됩니다.
{% endhint %}

이 페이지에서 다음 정보를 찾을 수 있습니다:

* [Snyk 런타임 센서를 위한 전제 조건](snyk-runtime-sensor.md#prerequisites-for-snyk-runtime-sensor)
* [Snyk 런타임 센서 설치](snyk-runtime-sensor.md#install-snyk-runtime-sensor)
  * [Helm 차트 사용](snyk-runtime-sensor.md#using-a-helm-chart)
  * [Helm 차트 및 AWS Secrets Manager 사용](snyk-runtime-sensor.md#using-a-helm-chart-and-the-aws-secrets-manager)
  * [OpenShift에 설치](snyk-runtime-sensor.md#on-openshift)
  * [AWS Marketplace를 통한 EKS 애드온으로 설치](snyk-runtime-sensor.md#aws-eks-deployment)
* [Snyk 런타임 센서에 대한 문제 해결](snyk-runtime-sensor.md#troubleshooting-for-snyk-runtime-sensor)

## Snyk 런타임 센서를 위한 전제 조건

Snyk 런타임 센서를 올바르게 사용하려면 환경이 다음 기술적 요구 사항을 충족하는지 확인하십시오:

* 지원되는 Kubernetes 버전 - Kubernetes v.1.19 이상 사용.

{% hint style="info" %}
EKS Fargate 또는 GKE Autopilot과 같은 관리형 쿠버네티스 서비스는 지원되지 않습니다. 왜냐하면 클러스터 노드가 클라우드 공급업체에 의해 관리됩니다.
{% endhint %}

* CPU 아키텍처 - AMD64 또는 ARM64.
* Linux 커널 - 버전 5.8 이상.
* 권한 있는 액세스 - `BPF`, `PERFMON`, `SYS_RESOURCES`, `DAC_READ_SEARCH`, `SYS_PTRACE`, `NET_ADMIN` Linux 기능 중 하나를 가져야 합니다.
* 클러스터 노드는 BTF 지원해야 합니다.
* 언어 지원 - Go, Java v8 이상, .NET v2.0.9 이상, Node.js v10 이상, Python 3.6 이상.
* 네트워크 정책 - 클러스터가 모든 외부 트래픽을 허용하지 않으면, 다음 호스트로의 443 포트 발신 트래픽을 활성화하는 정책을 설정하십시오:
  * `api.snyk.io` 또는 `api.<<REGION>>.snyk.io` (다른 지역에 호스팅된 경우).
  * `api.sentry.io`
  * `kubernetes.default.svc.cluster.local`

{% hint style="warning" %}
네트워크 제한 사항이 발생하는 경우, 포트 443가 활성화되어 있는지 확인하고 정책이 상태를 유지하는지 확인하십시오.
{% endhint %}

또한 [서비스 계정](https://docs.snyk.io/snyk-admin/service-accounts)용 토큰이 필요합니다. 서비스 계정은 다음 중 하나의 역할을 가져야합니다:

* 그룹 관리자
* `AppRisk edit`권한이 활성화된 사용자 지정 그룹 레벨 역할.

## Snyk 런타임 센서 설치

{% hint style="info" %}
* Snyk 런타임 센서는 [Kubernetes 데몬셋](https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/)으로 배포되며, 이는 클러스터의 각 노드당 센서 당 하나의 [Pod](https://kubernetes.io/docs/concepts/workloads/pods/)가 있음을 의미합니다.
* Snyk 런타임 센서는 영구 저장소를 사용하지 않으며, 디스크 사용량은 무시할 수준입니다.
{% endhint %}

* Snyk 런타임 센서 데몬셋은 다음 최소 요구 사항을 충족해야합니다:
  * `CPU: 100m` (Helm을 사용하여 증가 가능)
  * `Memory: 512Mi` (Helm을 사용하여 증가 가능)
* 다음 중 하나의 방법을 선택하여 Snyk 런타임 센서를 배포할 수 있습니다:
  * [Helm 차트를 사용하여 Snyk 런타임 센서 설치](snyk-runtime-sensor.md#using-a-helm-chart)
  * [Helm 차트 및 AWS Secrets Manager를 사용하여 Snyk 런타임 센서 설치](snyk-runtime-sensor.md#using-a-helm-chart-and-the-aws-secrets-manager)
  * [OpenShift에 Snyk 런타임 센서 설치](snyk-runtime-sensor.md#on-openshift)
  * [AWS Marketplace를 통해 EKS 애드온으로 Snyk 런타임 센서 설치](snyk-runtime-sensor.md#through-the-aws-marketplace-as-an-eks-add-on)

### Helm 차트 사용

이 리포지토리 내 [helm/runtime-sensor](https://github.com/snyk/runtime-sensor)에 [Helm 차트](https://helm.sh)가 있으며, 이는 `https://snyk.github.io/runtime-sensor`의 GitHub 페이지를 통해 호스팅됩니다.

Helm 차트를 사용하여 Snyk 런타임 센서를 설치하려면 다음 단계를 따를 수 있습니다:

1. Helm이 설치되었는지 확인하십시오.
2.  `snyk-runtime-sensor` 네임스페이스를 만드십시오:

    <pre><code><strong>kubectl create namespace snyk-runtime-sensor
    </strong></code></pre>
3. AppRisk 전제 사항 섹션에 지시된 대로 해당 네임스페이스에 적절한 권한이 있는 서비스 계정 토큰을 포함하는 시크릿을 생성하십시오:

{% code overflow="wrap" %}
```
kubectl create secret generic <<YOUR_SECRET_NAME>> --from-literal=snykToken=<<YOUR_TOKEN>> -n snyk-runtime-sensor
```
{% endcode %}

4.  Helm 리포지토리를 추가하십시오:

    ```
    helm repo add runtime-sensor https://snyk.github.io/runtime-sensor
    ```
5. 데이터가 기본 지역(미국)과 다른 지역에 호스팅될 경우, Helm 차트 설치 시 `snykAPIBaseURL`을 다음 형식으로 설정해야합니다: `api.<<REGION>>.snyk.io:443`, 예: `api.eu.snyk.io:443`
6.  (선택 사항) 런타임 센서 이미지의 사용자 정의 리소스(CPU/메모리)를 구성하려면 다음 값을 지정하십시오(여기서는 기본값 사용):

    ```
    ...
    --set sensor.resources.requests.memory=512Mi
    --set sensor.resources.requests.cpu=100m
    --set sensor.resources.limits.memory=1024Mi
    --set sensor.resources.limits.cpu=500m
    ...
    ```
7.  Helm 차트를 설치하십시오:

    ```
    helm install my-runtime-sensor \
    --set secretName=<<YOUR_SECRET_NAME>> \
    --set clusterName=<<CLUSTER_NAME>> \
    --set snykGroupId=<<YOUR_GROUP_ID>> \
    --set snykAPIBaseURL=api.<<REGION>>.snyk.io:443 \ # 선택 사항
    -n snyk-runtime-sensor \
    runtime-sensor/runtime-sensor
    ```

#### 최신 버전으로 업그레이드

1. 센서에 주어진 이름을 확인하십시오.

{% code overflow="wrap" %}
```
helm repo list
```
{% endcode %}

2. 리포지토리를 업데이트하십시오((1)에서 받은 이름으로):

{% code overflow="wrap" %}
```
helm repo update <<SENSOR_REPO_NAME>>
```
{% endcode %}

3.  설치를 업그레이드하십시오:

    ```
    helm upgrade --install <<SENSOR_REPO_NAME>> \
    --set secretName=<<YOUR_SECRET_NAME>> \
    --set clusterName=<<CLUSTER_NAME>> \
    --set snykGroupId=<<YOUR_GROUP_ID>> \
    -n snyk-runtime-sensor \
    runtime-sensor/runtime-sensor
    ```

### Helm 차트 및 AWS Secrets Manager 사용

이 리포지토리 내 [helm/runtime-sensor](https://github.com/snyk/runtime-sensor)에 [Helm 차트](https://helm.sh)가 있으며, 이는 `https://snyk.github.io/runtime-sensor`의 GitHub 페이지를 통해 호스팅됩니다.

Helm 차트 및 AWS Secrets Manager를 사용하여 Snyk 런타임 센서를 설치하려면 다음 단계를 따를 수 있습니다:

전제 조건: 권장 [여기](https://github.com/aws/secrets-store-csi-driver-provider-aws)를 따라 클러스터에 AWS Provider 및 CSI Secrets Store를 설치하십시오.

1. Helm이 설치되었는지 확인하십시오.
2.  `snyk-runtime-sensor` 네임스페이스를 만드십시오:

    <pre><code><strong>kubectl create namespace snyk-runtime-sensor
    </strong></code></pre>
3.  AWS 계정에 `snykToken` 키 아래 컨텐츠를 포함하는 Snyk 런타임 센서 시크릿을 만들고 얻으십시오:

    ```
    aws secretsmanager create-secret \
    --name snyk-runtime-sensor-secret \
    --secret-string '{"snykToken":"<<YOUR_SERVICE_ACCOUNT_TOKEN>>"}'
    ```
4.  새로 생성된 시크릿에 대한 액세스 정책을 생성:

    ```
    POLICY_ARN=$(aws --query Policy.Arn --output text iam create-policy --policy-name snyk-runtime-sensor-secret-policy --policy-document '{
        "Version": "2012-10-17",
        "Statement": [ {
            "Effect": "Allow",
            "Action": ["secretsmanager:GetSecretValue", "secretsmanager:DescribeSecret"],
            "Resource": ["<<YOUR_SECRET'S_ARN>>"]
        } ]
    }')
    ```
5.  아직 수행하지 않은 경우, 클러스터에 IAM OIDC 프로바이더를 만드십시오 (한 번만 실행):

    ```
    eksctl utils associate-iam-oidc-provider \
    --cluster="<<CLUSTER_NAME>>" \
    --approve
    ```
6.  Helm 리포지토리를 추가하십시오:

    ```
    helm repo add runtime-sensor https://snyk.github.io/runtime-sensor
    ```
7. 데이터가 기본 지역(미국)과 다른 지역에 호스팅될 경우, Helm 차트 설치 시 `snykAPIBaseURL`을 다음 형식으로 설정해야합니다: `api.<<REGION>>.snyk.io:443`, 예: `api.eu.snyk.io:443`
8.  (선택 사항) 런타임 센서 이미지의 사용자 정의 리소스(CPU/메모리)를 구성하려면 다음 값을 지정하십시오(여기서는 기본값 사용):

    ```
    ...
    --set sensor.resources.requests.memory=512Mi
    --set sensor.resources.requests.cpu=100m
    --set sensor.resources.limits.memory=1024Mi
    --set sensor.resources.limits.cpu=500m
    ...
    ```
9.  Helm 차트를 설치하십시오:

    ```
    helm install my-runtime-sensor \
    --set secretProvider=aws \
    --set secretName=snyk-runtime-sensor-secret \
    --set clusterName=<<CLUSTER_NAME>> \
    --set snykGroupId=<<YOUR_GROUP_ID>> \
    --set snykAPIBaseURL=api.<<REGION>>.snyk.io:443 \ # 선택 사항
    -n snyk-runtime-sensor \
    runtime-sensor/runtime-sensor
    ```
10. 다음과 같이 새롭게 생성된 서비스 계정에 단계 4에서 생성된 정책의 ARN을 첨부하여 새롭게 생성된 서비스 계정에 대한 새 역할을 생성하십시오:

{% code overflow="wrap" %}
````
```
eksctl create iamserviceaccount \
--name runtime-sensor \
--region=<<REGION>> \
--cluster "<<CLUSTER_NAME>>" \
--attach-policy-arn "$POLICY_ARN" \
--approve \
--override-existing-serviceaccounts \
--namespace=snyk-runtime-sensor
```
````
{% endcode %}

11. `snyk-runtime-sensor` 네임스페이스로 시크릿이 성공적으로 마운트되었는지 (`kubectl get secrets -n snyk-runtime-sensor`) 및 센서 포드가 성공적으로 실행 중인지 확인하십시오 (`kubectl get pods -n snyk-runtime-sensor`).

#### 최신 버전으로 업그레이드

1. 센서에 주어진 이름을 확인하십시오.

{% code overflow="wrap" %}
```
helm repo list
```
{% endcode %}

2. 리포지토리를 업데이트하십시오((1)에서 받은 이름으로):

{% code overflow="wrap" %}
```
helm repo update <<SENSOR_REPO_NAME>>
```
{% endcode %}

3.  설치를 업그레이드하십시오:

    ```
    helm upgrade --install <<SENSOR_REPO_NAME>> \
    --set secretProvider=aws**AWS 콘솔**
    ```

Snyk 런타임 센서를 AWS Marketplace에서 성공적으로 구독하였고 화면 안내에 따라 진행한 후에는 Amazon EKS 콘솔로 리디렉션됩니다.

Amazon EKS 클러스터에 Snyk 런타임 센서를 활성화하려면 Amazon EKS 콘솔에서 클러스터를 선택합니다. 그런 다음 추가 기능 탭으로 이동하고 "Get more add-ons"을 선택합니다. 검색 창을 사용하여 "runtime"을 찾은 후 설치할 클러스터에 추가 기능을 활성화하기 위해 화면 안내에 따릅니다. 이 프로세스는 센서를 통합하려는 각 클러스터에 대해 수행되어야 합니다(즉, 데이터를 수집하려는 각 클러스터에 대해).

<figure><img src="../../../.gitbook/assets/Screenshot 2024-05-26 at 15.53.23.png" alt="Snyk 런타임 센서 추가 기능 선택"><figcaption><p>Snyk 런타임 센서 추가 기능 선택</p></figcaption></figure>

다음 화면에서 최신 버전을 선택한 후(이미 선택한 경우도) "Optional configuration settings"를 엽니다.

<figure><img src="../../../.gitbook/assets/Screenshot 2024-05-26 at 15.55.19.png" alt=""><figcaption><p>가용한 최신 버전을 선택하고 'Optional configuration settings'를 엽니다</p></figcaption></figure>

"구성 값" 아래 YAML 또는 JSON 형식으로 다음 속성을 설정합니다:

* `secretName` - 나중에 생성될 시크릿 이름. 기본값은 `snyk-secret` 입니다.
* `clusterName` - 추가 기능이 설치된 클러스터의 이름.
* `snykGroupId` - 사용된 서비스 계정과 관련된 그룹 ID.
* `snykAPIBaseURL` - 기본값(미국이 아닌 다른 지역에 데이터가 호스팅된 경우를 제외하고)은 `api.snyk.io:443`로 설정해야 합니다.

다음은 복사할 기본 구성입니다:

```
secretName: snyk-secret
clusterName: <<MY_CLUSTER>>
snykGroupId: <<MY_SNYK_GROUP_ID>>
snykAPIBaseURL: api.snyk.io:443
```

<figure><img src="../../../.gitbook/assets/Screenshot 2024-05-26 at 15.58.12.png" alt="&#x27;Optional configuraiton settings&#x27; 아래 적절한 구성 값 설정"><figcaption><p>'Optional configuraiton settings' 아래 적절한 구성 값 설정</p></figcaption></figure>

**다음** 및 **생성** 옵션을 선택하면 페이지 상단에 `클러스터 <<YOUR_CLUSTER>>에 Add-on snyk-runtimesensor가 성공적으로 추가되었음` 알림이 표시됩니다.

<figure><img src="../../../.gitbook/assets/Screenshot 2024-05-26 at 16.26.13.png" alt="성공 메시지."><figcaption><p>성공 메시지.</p></figcaption></figure>

#### **AWS CLI를 사용하여 Snyk 런타임 센서 추가 기능 활성화**

작업 공간에서 다음 명령을 실행하여 Amazon EKS 클러스터에 Snyk 런타임 센서 추가 기능을 활성화합니다. 다음 환경 변수를 설정하여 Snyk 계정 및 대상 EKS 클러스터와 일치시켜야 합니다:

* $CLUSTER\_NAME
* $AWS\_REGION
* $SNYK\_GROUP\_ID
* $SNYK\_API\_BASE\_URL(미국을 제외한 다른 지역에서 호스팅되는 경우 `api.snyk.io:443`로 설정해야 합니다).

```
aws eks create-addon \
--cluster-name $CLUSTER_NAME \
--region $AWS_REGION \
--addon-name snyk_runtime-sensor \
--configuration-values '{"secretName":"snyk-secret","clusterName":"$CLUSTER_NAME","snykGroupId":"$SNYK_GROUP_ID","snykAPIBaseURL": "$SNYK_API_BASE_URL"}' \
--resolve-conflicts OVERWRITE
```

[아래](snyk-runtime-sensor.md#add-your-snyk-service-account-token-to-the-eks-cluster)에 설명된대로 Snyk 서비스 계정 토큰을 추가한 후 다음 명령을 실행하여 설치가 성공적으로 완료되었는지 확인합니다:

```
aws eks describe-addon --addon-name snyk_runtime-sensor --cluster-name $CLUSTER_NAME --region $AWS_REGION
```

받은 응답이 이와 유사하고 상태가 ACTIVE로 표시되는지 확인하십시오. 몇 분이 걸려 이 상태에 도달할 수 있습니다.

```
{
    "addon": {
        "addonName": "snyk_runtime-sensor",
        "clusterName": "<<YOUR_CLUSTER>>",
        "status": "ACTIVE",
        "addonVersion": "v1.17.2-eksbuild.1",
        "health": {
            "issues": []
        },
        "addonArn": "arn:aws:eks:us-east-1:XXXX:addon/<<YOUR_CLUSTER>>/snyk-runtimesensor/ffffffff-ffff-ffff-ffff-ffffffffffff",
        "createdAt": "2024-05-26T16:43:15.551000+03:00",
        "modifiedAt": "2024-05-26T16:43:28.714000+03:00",
        "tags": {},
        "configurationValues": "{\"secretName\":\"snyk-secret\",\"clusterName\":\"YOUR_CLUSTER\",\"snykGroupId\":\"YOUR_GROUP_ID\",\"snykAPIBaseURL\": \"api.snyk.io:443\"}"
    }
}
```

#### **EKS 클러스터에 Snyk 서비스 계정 토큰 추가하기**

* `aws eks`를 사용하여 클러스터를 제어하기 위해 `kubectl` 컨텍스트를 설정합니다:

```
aws eks update-kubeconfig --name $CLUSTER_NAME --region $AWS_REGION
```

* `snykToken`을 포함하는 `snyk-secret` 이름의 시크릿을 생성합니다. `snykToken`은 서비스 계정 토큰입니다:

```
kubectl create secret generic snyk-secret \
--from-literal=snykToken=$SNYK_TOKEN \
-n snyk-runtime-sensor
```

* AWS EKS 클러스터에서 데이터는 Snyk 런타임 센서를 통해 보고됩니다.

#### **Snyk 런타임 센서 추가 기능 비활성화**

다음 명령을 실행하여 Snyk 런타임 센서 추가 기능을 비활성화할 수 있습니다:

```
aws eks delete-addon --addon-name snyk_runtime-sensor --cluster-name $CLUSTER_NAME --region $AWS_REGION
```

## Snyk 런타임 센서의 문제 해결

Snyk 런타임 센서가 `is_loaded` 리스크 팩터를 적절하게 보고하지 않는 경우, Linux 커널 `perf_event_paranoid` 구성의 기본값이 아닌 값을 사용하는 것이 원인일 수 있습니다.\
이러한 경우, `--set securityContext.privileged=true` 또는 `--set "securityContext.capabilities={SYS_ADMIN}"`을 사용하여 헬름 차트를 설치합니다.

{% hint style="info" %}
로드된 패키지 리스크 팩터는 Snyk에서 운영 체제 패키지(예: Debian 패키지)에 대해 지원되지 않으며 npm, Maven 또는 PyPI와 같은 패키지 관리자에 호스팅된 패키지에만 지원됩니다.
{% endhint %}

릴리스 버전은 [GitHub](https://github.com/snyk/runtime-sensor/releases)에서 확인할 수 있습니다.
