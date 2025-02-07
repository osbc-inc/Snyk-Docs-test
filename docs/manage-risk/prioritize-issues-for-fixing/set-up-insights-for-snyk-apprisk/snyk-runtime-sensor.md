# Snyk 런타임 센서

{% hint style="info" %}
**릴리스 상태**

Snyk AppRisk용 Snyk 런타임 센서는 현재 초기 액세스(Early Access) 단계에 있으며, Snyk AppRisk가 포함된 Snyk Enterprise 플랜에서만 사용할 수 있습니다. 그룹에서 설정하려면 Snyk 계정 팀에 문의하세요.
{% endhint %}

Snyk 런타임 센서는 [Kubernetes DaemonSet](https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/)으로, Kubernetes 클러스터에서 배포를 감시하고 수집된 데이터를 Snyk으로 전송합니다.

Snyk 런타임 센서에서 보고하는 [위험 요소](https://docs.snyk.io/manage-risk/prioritize-issues-for-fixing/assets-and-risk-factors-for-snyk-apprisk#risk-factors)는 다음과 같습니다: [배포됨(Deployed)](https://docs.snyk.io/manage-risk/prioritize-issues-for-fixing/assets-and-risk-factors-for-snyk-apprisk/risk-factor-deployed) 및 [로드된 패키지(Loaded package)](https://docs.snyk.io/manage-risk/prioritize-issues-for-fixing/assets-and-risk-factors-for-snyk-apprisk/risk-factor-loaded-package).

{% hint style="info" %}
* Snyk 런타임 센서는 애플리케이션 패키지에 대해서만 로드된 패키지(Loaded package) 위험 요소를 보고합니다. 지원되는 생태계는 다음과 같습니다:
  * Node.js
  * Java
  * Python
  * Go
  * .NET
* 대부분의 패키지는 애플리케이션 시작 시 메모리에 로드되므로, Pod가 초기화될 때만 보고됩니다.
{% endhint %}

이 페이지에서는 다음 정보를 확인할 수 있습니다:

* [Snyk 런타임 센서의 필수 조건](snyk-runtime-sensor.md#prerequisites-for-snyk-runtime-sensor)
* [Snyk 런타임 센서 설치](snyk-runtime-sensor.md#install-snyk-runtime-sensor)
  * [Kubernetes DaemonSet으로 Helm 차트를 사용하여 설치](snyk-runtime-sensor.md#as-a-kubernetes-daemonset-using-a-helm-chart)
  * [Kubernetes Deployment로 Helm 차트를 사용하여 설치](snyk-runtime-sensor.md#as-a-kubernetes-deployment-using-a-helm-chart)
  * [Helm 차트와 AWS Secrets Manager를 사용하여 설치](snyk-runtime-sensor.md#using-a-helm-chart-and-the-aws-secrets-manager)
  * [OpenShift에서 설치](snyk-runtime-sensor.md#on-openshift)
  * [AWS Marketplace를 통해 EKS 애드온으로 설치](snyk-runtime-sensor.md#aws-eks-deployment)
* [Snyk 런타임 센서 문제 해결](snyk-runtime-sensor.md#troubleshooting-for-snyk-runtime-sensor)

## Snyk 런타임 센서의 필수 조건

Snyk 런타임 센서를 올바르게 사용하려면 환경이 다음 기술적 필수 조건을 충족해야 합니다:

* 지원되는 Kubernetes 버전 - Kubernetes v1.19 이상 사용.

{% hint style="info" %}
EKS Fargate 또는 GKE Autopilot과 같은 관리형 Kubernetes 서비스는 지원되지 않습니다. 클러스터 노드는 클라우드 공급자가 관리하기 때문입니다.
{% endhint %}

* CPU 아키텍처 - AMD64 또는 ARM64.
* Linux 커널 - 버전 5.8 이상.
* 권한이 있는 액세스 - 루트(root) 권한 또는 다음 Linux 기능(capabilities)이 필요함: `BPF`, `PERFMON`, `SYS_RESOURCES`, `DAC_READ_SEARCH`, `SYS_PTRACE`, `NET_ADMIN`
* 클러스터 노드는 BTF를 지원해야 함.
* 지원되는 프로그래밍 언어 - Go, Java v8 이상, .NET v2.0.9 이상, Node.js v10 이상, Python 3.6 이상.
* 네트워크 정책 - 클러스터에서 모든 아웃바운드 트래픽을 허용하지 않는 경우, 다음 호스트에 대한 443번 포트의 아웃바운드 트래픽을 허용하도록 정책을 설정해야 합니다:
  * `api.snyk.io` 또는 다른 지역에 호스팅된 경우 `api.<<REGION>>.snyk.io`
  * `api.sentry.io`
  * `kubernetes.default.svc.cluster.local`

{% hint style="warning" %}
네트워크 제한이 있는 경우, 443번 포트가 활성화되었는지 확인하고 정책이 상태 저장(stateful)인지 확인하세요.
{% endhint %}

또한 [서비스 계정](https://docs.snyk.io/snyk-admin/service-accounts)용 토큰이 필요합니다. 서비스 계정은 다음 중 하나의 역할을 가져야 합니다:

* 그룹 관리자(Group Admin)
* `AppRisk edit` 권한이 활성화된 사용자 정의 그룹 수준 역할(Custom Group Level Role)

## Snyk 런타임 센서 설치

* Snyk 런타임 센서 DaemonSet은 다음 최소 요구 사항을 충족해야 합니다:
  * `CPU: 100m` (Helm을 사용하여 증가 가능)
  * `메모리: 512Mi` (Helm을 사용하여 증가 가능)
* Snyk 런타임 센서를 배포하는 방법 중 하나를 선택하세요:
  * [Helm 차트를 사용하여 Snyk 런타임 센서 설치](snyk-runtime-sensor.md#using-a-helm-chart)
  * [Helm 차트와 AWS Secrets Manager를 사용하여 Snyk 런타임 센서 설치](snyk-runtime-sensor.md#using-a-helm-chart-and-the-aws-secrets-manager)
  * [OpenShift에서 Snyk 런타임 센서 설치](snyk-runtime-sensor.md#on-openshift)
  * [AWS Marketplace를 통해 EKS 애드온으로 Snyk 런타임 센서 설치](snyk-runtime-sensor.md#through-the-aws-marketplace-as-an-eks-add-on)


### Kubernetes DaemonSet으로서 Helm 차트 사용하기  

{% hint style="info" %}
* Snyk 런타임 센서는 [Kubernetes DaemonSet](https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/)으로 설치되며, 클러스터 내 사용 가능한 각 [노드](https://kubernetes.io/docs/concepts/architecture/nodes/)마다 하나의 [Pod](https://kubernetes.io/docs/concepts/workloads/pods/)이 생성됩니다.
* Snyk 런타임 센서는 영구 저장소를 사용하지 않으며, 디스크 사용량이 최소화되어 있습니다.
* Kubernetes DaemonSet으로 설치될 경우 eBPF 기능이 활성화되며, 배포된(Deployed) 패키지 및 로드된(Loaded) 패키지의 위험 요소를 보고할 수 있습니다.
{% endhint %}

[helm/runtime-sensor](https://github.com/snyk/runtime-sensor) 저장소에서 제공되는 [Helm 차트](https://helm.sh)를 참고하여 설치를 진행할 수 있으며, 해당 차트는 GitHub Pages를 통해 `https://snyk.github.io/runtime-sensor`에서 호스팅됩니다.

Kubernetes DaemonSet으로 Helm 차트를 사용하여 Snyk 런타임 센서를 설치하려면 다음 단계를 따르세요:

1. Helm이 설치되어 있는지 확인합니다.
2. `snyk-runtime-sensor` 네임스페이스를 생성합니다:

    <pre><code><strong>kubectl create namespace snyk-runtime-sensor
    </strong></code></pre>
3. [필수 조건](snyk-runtime-sensor.md#prerequisites-for-snyk-runtime-sensor) 섹션에서 안내하는 적절한 권한이 있는 서비스 계정 토큰을 사용하여 시크릿을 생성하고, 해당 네임스페이스에 저장합니다:

    {% code overflow="wrap" %}
    ```
    kubectl create secret generic <<YOUR_SECRET_NAME>> --from-literal=snykToken=<<YOUR_TOKEN>> -n snyk-runtime-sensor
    ```
    {% endcode %}
4. Helm 저장소를 추가합니다:

    ```
    helm repo add runtime-sensor https://snyk.github.io/runtime-sensor
    ```
5. 기본 지역(미국) 외의 [다른 지역](../../../working-with-snyk/regional-hosting-and-data-residency.md)에 데이터가 호스팅되는 경우, Helm 차트를 설치할 때 `snykAPIBaseURL`을 다음 형식으로 설정해야 합니다:
   `api.<<REGION>>.snyk.io:443`, 예: `api.eu.snyk.io:443`
6.  (선택 사항) 센서가 모니터링할 Pod을 필터링하려면 특정 워크로드 유형, 네임스페이스 및 Pod 레이블을 허용 목록에 추가할 수 있습니다. 기본적으로 모든 Pod이 모니터링됩니다.

    ```
    ...
    --set 'sensor.filters.workloadTypes={deployment,cronjob}'
    --set 'sensor.filters.namespaces={ns1,ns2}'
    --set sensor.filters.podLabels.label_key1='label_value1'
    --set sensor.filters.podLabels.label_key2='label_value2'
    ...
    ```

    사용 가능한 워크로드 유형:

    ```json
    deployment
    daemonset
    statefulset
    replicaset
    job
    cronjob
    ```
7.  (선택 사항) 런타임 센서 이미지의 리소스(CPU/메모리)를 사용자 정의하려면 다음 값을 설정해야 합니다 (기본값 사용 예시):

    ```
    ...
    --set sensor.resources.requests.memory=512Mi
    --set sensor.resources.requests.cpu=100m
    --set sensor.resources.limits.memory=1024Mi
    --set sensor.resources.limits.cpu=500m
    ...
    ```
8.  Helm 차트를 설치합니다:

    ```
    helm install my-runtime-sensor \
    --set workloadType=daemonset \ # 기본값이 'daemonset'이므로 생략 가능
    --set secretName=<<YOUR_SECRET_NAME>> \
    --set clusterName=<<CLUSTER_NAME>> \
    --set snykGroupId=<<YOUR_GROUP_ID>> \
    --set snykAPIBaseURL=api.<<REGION>>.snyk.io:443 \ # 선택 사항
    -n snyk-runtime-sensor \
    runtime-sensor/runtime-sensor
    ```

### Kubernetes Deployment으로서 Helm 차트 사용하기  

{% hint style="info" %}  
* Snyk 런타임 센서는 [Kubernetes Deployment](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/)로 설치되며, 이는 클러스터 전체에서 단일 [Pod](https://kubernetes.io/docs/concepts/workloads/pods/)만 실행됨을 의미합니다.  
* Snyk 런타임 센서는 지속적인 스토리지를 사용하지 않으며, 디스크 사용량이 최소화됩니다.  
{% endhint %}  

{% hint style="warning" %}  
Kubernetes Deployment에서는 eBPF 기능이 비활성화되므로, Snyk 런타임 센서는 배포된 위험 요소(Deployed risk factor)만 보고합니다.  
{% endhint %}  

GitHub Pages를 통해 `https://snyk.github.io/runtime-sensor`에서 호스팅되는 [helm/runtime-sensor](https://github.com/snyk/runtime-sensor) 저장소의 [Helm 차트](https://helm.sh)를 참고하여 설치할 수 있습니다.  

Helm 차트를 사용하여 Kubernetes Deployment로 Snyk 런타임 센서를 설치하려면 다음 단계를 따르세요:  

1. Helm이 설치되어 있는지 확인합니다.  
2. `snyk-runtime-sensor` 네임스페이스를 생성합니다:  

    <pre><code><strong>kubectl create namespace snyk-runtime-sensor
    </strong></code></pre>  

3. 생성한 네임스페이스에서, 적절한 권한이 있는 서비스 계정 토큰(자세한 내용은 [사전 요구 사항](snyk-runtime-sensor.md#prerequisites-for-snyk-runtime-sensor) 섹션 참고)을 포함하는 시크릿을 생성합니다:  

    {% code overflow="wrap" %}  
    ```  
    kubectl create secret generic <<YOUR_SECRET_NAME>> --from-literal=snykToken=<<YOUR_TOKEN>> -n snyk-runtime-sensor  
    ```  
    {% endcode %}  

4. Helm 저장소를 추가합니다:  

    ```  
    helm repo add runtime-sensor https://snyk.github.io/runtime-sensor  
    ```  

5. 데이터가 기본 지역(미국) 외의 [다른 지역](../../../working-with-snyk/regional-hosting-and-data-residency.md)에 호스팅되는 경우, Helm 차트를 설치할 때 `snykAPIBaseURL`을 다음 형식으로 설정해야 합니다:  
   `api.<<REGION>>.snyk.io:443` (예: `api.eu.snyk.io:443`)  

6. (선택 사항) 센서가 모니터링할 Pod을 필터링하려면 특정 워크로드 유형, 네임스페이스, Pod 레이블을 허용 목록(allow list)에 추가할 수 있습니다. 센서는 제공된 모든 필터 조건을 충족하는 Pod만 모니터링하며, 기본적으로 모든 Pod이 모니터링됩니다.  

    ```  
    ...  
    --set 'sensor.filters.workloadTypes={deployment,cronjob}'  
    --set 'sensor.filters.namespaces={ns1,ns2}'  
    --set sensor.filters.podLabels.label_key1='label_value1'  
    --set sensor.filters.podLabels.label_key2='label_value2'  
    ...  
    ```  

    사용 가능한 워크로드 유형:  

    ```json  
    deployment  
    daemonset  
    statefulset  
    replicaset  
    job  
    cronjob  
    ```  

7. (선택 사항) 런타임 센서 이미지에 대해 사용자 지정 리소스(CPU/메모리)를 설정하려면 다음 값을 지정해야 합니다 (기본값 사용 예시):  

    ```  
    ...  
    --set sensor.resources.requests.memory=512Mi  
    --set sensor.resources.requests.cpu=100m  
    --set sensor.resources.limits.memory=1024Mi  
    --set sensor.resources.limits.cpu=500m  
    ...  
    ```  

8. Helm 차트를 설치합니다:  

    ```  
    helm install my-runtime-sensor \  
    --set workloadType=deployment \  
    --set secretName=<<YOUR_SECRET_NAME>> \  
    --set clusterName=<<CLUSTER_NAME>> \  
    --set snykGroupId=<<YOUR_GROUP_ID>> \  
    --set snykAPIBaseURL=api.<<REGION>>.snyk.io:443 \ # 선택 사항  
    -n snyk-runtime-sensor \  
    runtime-sensor/runtime-sensor  
    ```  

### 최신 버전으로 업그레이드하기  

1. 센서에 할당된 이름을 확인합니다.  

    {% code overflow="wrap" %}  
    ```  
    helm repo list  
    ```  
    {% endcode %}  

2. (1)에서 확인한 이름을 사용하여 저장소를 업데이트합니다.  

    {% code overflow="wrap" %}  
    ```  
    helm repo update <<SENSOR_REPO_NAME>>  
    ```  
    {% endcode %}  

3. 설치를 업그레이드합니다.  

    ```  
    helm upgrade --install <<SENSOR_REPO_NAME>> \  
    --set secretName=<<YOUR_SECRET_NAME>> \  
    --set clusterName=<<CLUSTER_NAME>> \  
    --set snykGroupId=<<YOUR_GROUP_ID>> \  
    -n snyk-runtime-sensor \  
    runtime-sensor/runtime-sensor  
    ```  

### Helm 차트 및 AWS Secrets Manager 사용하기  

이 저장소에는 [Helm 차트](https://helm.sh)가 있으며, [helm/runtime-sensor](https://github.com/snyk/runtime-sensor)에서 관리되며 GitHub Pages(`https://snyk.github.io/runtime-sensor`)를 통해 호스팅됩니다.  

Helm 차트와 AWS Secrets Manager를 사용하여 Snyk 런타임 센서를 설치하려면 다음 단계를 따르세요:  

사전 요구 사항: [여기](https://github.com/aws/secrets-store-csi-driver-provider-aws)의 가이드에 따라 AWS Provider 및 CSI Secrets Store를 클러스터에 설치합니다.  

1. Helm이 설치되어 있는지 확인합니다.  
2. `snyk-runtime-sensor` 네임스페이스를 생성합니다.  

    <pre><code><strong>kubectl create namespace snyk-runtime-sensor
    </strong></code></pre>  

3. AWS 계정에서 서비스 계정 토큰을 포함하는 Snyk 런타임 센서 시크릿을 생성하고, 생성된 ARN을 확인합니다.  

    ```  
    aws secretsmanager create-secret \  
    --name snyk-runtime-sensor-secret \  
    --secret-string '{"snykToken":"<<YOUR_SERVICE_ACCOUNT_TOKEN>>"}'  
    ```  

4. 새로 생성된 시크릿에 대한 액세스 정책을 생성합니다.  

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

5. 클러스터에 대한 IAM OIDC 공급자를 생성합니다 (이미 생성된 경우 건너뛰어도 됩니다).  

    ```  
    eksctl utils associate-iam-oidc-provider \  
    --cluster="<<CLUSTER_NAME>>" \  
    --approve  
    ```  

6. Helm 저장소를 추가합니다.  

    ```  
    helm repo add runtime-sensor https://snyk.github.io/runtime-sensor  
    ```  

7. 데이터가 기본 지역(미국) 외의 [다른 지역](../../../working-with-snyk/regional-hosting-and-data-residency.md)에 호스팅되는 경우, Helm 차트를 설치할 때 `snykAPIBaseURL`을 다음 형식으로 설정해야 합니다:  
   `api.<<REGION>>.snyk.io:443` (예: `api.eu.snyk.io:443`)  

8. (선택 사항) 런타임 센서 이미지에 대해 사용자 지정 리소스(CPU/메모리)를 설정하려면 다음 값을 지정합니다 (기본값 사용 예시):  

    ```  
    ...  
    --set sensor.resources.requests.memory=512Mi  
    --set sensor.resources.requests.cpu=100m  
    --set sensor.resources.limits.memory=1024Mi  
    --set sensor.resources.limits.cpu=500m  
    ...  
    ```  

9. Helm 차트를 설치합니다.  

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

10. 4단계에서 생성한 정책의 ARN을 새 서비스 계정에 연결하려면 다음 명령을 실행하여 새 역할을 생성합니다.  

    {% code overflow="wrap" %}  
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
    {% endcode %}  

11. 시크릿이 `snyk-runtime-sensor` 네임스페이스에 성공적으로 마운트되었는지 확인합니다. (`kubectl get secrets -n snyk-runtime-sensor`)  
    또한, 센서 Pod이 정상적으로 실행되고 있는지 확인합니다. (`kubectl get pods -n snyk-runtime-sensor`)  

### 최신 버전으로 업그레이드하기  

1. 센서에 할당된 이름을 확인합니다.  

    {% code overflow="wrap" %}  
    ```  
    helm repo list  
    ```  
    {% endcode %}  

2. (1)에서 확인한 이름을 사용하여 저장소를 업데이트합니다.  

    {% code overflow="wrap" %}  
    ```  
    helm repo update <<SENSOR_REPO_NAME>>  
    ```  
    {% endcode %}  

3. 설치를 업그레이드합니다.  

    ```  
    helm upgrade --install <<SENSOR_REPO_NAME>> \  
    --set secretProvider=aws \  
    --set secretName=snyk-runtime-sensor-secret \  
    --set clusterName=<<CLUSTER_NAME>> \  
    --set snykGroupId=<<YOUR_GROUP_ID>> \  
    -n snyk-runtime-sensor \  
    runtime-sensor/runtime-sensor  
    ```  

### OpenShift에서 실행하기  

OpenShift에서 Kubernetes 클러스터를 실행하는 경우, Snyk 런타임 센서의 서비스 계정에 `privileged` 보안 컨텍스트 제약(SCC)을 적용해야 합니다. 다음 명령어를 실행하세요:  

```  
oc adm policy add-scc-to-user privileged \  
system:serviceaccount:<<YOUR_NAMESPACE>>:runtime-sensor  
```  

이 명령은 센서가 설치된 후 실행해야 합니다. 서비스 계정은 설치가 완료될 때까지 사용 가능하지 않기 때문입니다.  

### AWS Marketplace를 통한 EKS 애드온 설치  <a href="#aws-eks-deployment" id="aws-eks-deployment"></a>  

Snyk은 AWS EKS 클러스터에서 Snyk 런타임 센서를 쉽게 설치할 수 있도록 지원합니다. 아래 단계에 따라 보안 기능을 환경에 통합하고 클러스터의 보안을 강화할 수 있습니다.  

Amazon EKS에서 EKS 애드온으로 Snyk 런타임 센서를 배포하려면 다음과 같은 사전 요구 사항을 충족해야 합니다.  

1. AWS Marketplace에서 Snyk 런타임 센서를 구독합니다. [여기](https://aws.amazon.com/marketplace/pp/prodview-i23vvrxuamcya)에서 구독할 수 있으며, 센서를 통합할 클러스터가 있는 모든 계정에서 수행해야 합니다.  
2. 다음 도구를 설치합니다: `kubectl`, `AWS CLI`, 선택적으로 `eksctl`.  
3. 센서를 설치하려는 Amazon EKS 클러스터에 대한 액세스 권한이 있는지 확인합니다.  
4. 올바른 권한을 가진 Snyk 서비스 계정 토큰이 있는지 확인합니다. 자세한 내용은 [사전 요구 사항](snyk-runtime-sensor.md#prerequisites)을 참조하세요.  

#### **AWS 콘솔에서 Snyk 런타임 센서 애드온 활성화하기**  

AWS Marketplace에서 Snyk 런타임 센서 구독을 설정하고 화면의 지침을 따른 후, Amazon EKS 콘솔로 리디렉션됩니다.  

Amazon EKS 콘솔에서 Snyk 런타임 센서를 활성화하려면, EKS 클러스터를 선택한 후 **Add-ons** 탭으로 이동하여 **Get more add-ons**을 선택하세요. 검색창에서 `"runtime"`을 입력하여 애드온을 찾고, 화면의 지침을 따라 클러스터에 애드온을 활성화합니다. 센서를 통합할 각 클러스터(즉, 데이터를 수집하려는 모든 클러스터)에 대해 이 과정을 수행해야 합니다.  

<figure><img src="../../../.gitbook/assets/Screenshot 2024-05-26 at 15.53.23.png" alt="Snyk 런타임 센서 애드온 선택"><figcaption><p>Snyk 런타임 센서 애드온 선택</p></figcaption></figure>  

다음 화면에서 최신 버전을 선택하고(이미 선택되어 있어도 다시 선택), **Optional configuration settings**을 엽니다.  

<figure><img src="../../../.gitbook/assets/Screenshot 2024-05-26 at 15.55.19.png" alt=""><figcaption><p>최신 버전을 선택하고 'Optional configuration settings' 열기</p></figcaption></figure>  

**Configuration values** 섹션에서 YAML 또는 JSON 형식으로 다음 속성을 설정합니다:  

* `secretName` - 이후 과정에서 생성될 시크릿 이름(기본값: `snyk-secret`)  
* `clusterName` - 애드온이 설치될 클러스터 이름  
* `snykGroupId` - 사용 중인 서비스 계정과 연결된 Group ID  
* `snykAPIBaseURL` - 기본적으로 `api.snyk.io:443`로 설정하되, 데이터가 [다른 지역](../../../working-with-snyk/regional-hosting-and-data-residency.md#what-regions-are-available)에 호스팅되는 경우 해당 지역의 URL을 사용해야 합니다.  

아래는 기본 설정 예제입니다:  

```  
secretName: snyk-secret  
clusterName: <<MY_CLUSTER>>  
snykGroupId: <<MY_SNYK_GROUP_ID>>  
snykAPIBaseURL: api.snyk.io:443  
```  

<figure><img src="../../../.gitbook/assets/Screenshot 2024-05-26 at 15.58.12.png" alt="&#x22;Optional configuration settings&#x22;에서 적절한 구성 값을 설정"><figcaption><p>"Optional configuration settings"에서 적절한 구성 값을 설정</p></figcaption></figure>  

**Next** 및 **Create**를 선택한 후, 페이지 상단에서 `Add-on snyk-runtimesensor successfully added to cluster <<YOUR_CLUSTER>>`라는 성공 메시지를 확인할 수 있습니다.  

<figure><img src="../../../.gitbook/assets/Screenshot 2024-05-26 at 16.26.13.png" alt="성공 메시지"><figcaption><p>성공 메시지</p></figcaption></figure>  

### **AWS CLI를 사용하여 Snyk Runtime Sensor 애드온 활성화하기**  

다음 명령어를 실행하여 Amazon EKS 클러스터에서 Snyk Runtime Sensor 애드온을 활성화하세요. Snyk 계정 및 대상 EKS 클러스터에 맞게 다음 환경 변수를 설정해야 합니다.  

* `$CLUSTER_NAME`  
* `$AWS_REGION`  
* `$SNYK_GROUP_ID`  
* `$SNYK_API_BASE_URL` (기본적으로 `api.snyk.io:443`로 설정해야 하며, 계정이 미국 이외의 지역에 호스팅된 경우 해당 지역에 맞게 변경해야 합니다.)  

```
aws eks create-addon \
--cluster-name $CLUSTER_NAME \
--region $AWS_REGION \
--addon-name snyk_runtime-sensor \
--configuration-values '{"secretName":"snyk-secret","clusterName":"$CLUSTER_NAME","snykGroupId":"$SNYK_GROUP_ID","snykAPIBaseURL": "$SNYK_API_BASE_URL"}' \
--resolve-conflicts OVERWRITE
```  

아래 [설명](snyk-runtime-sensor.md#add-your-snyk-service-account-token-to-the-eks-cluster)에 따라 Snyk 서비스 계정 토큰을 추가한 후, 다음 명령어를 실행하여 설치가 성공적으로 완료되었는지 확인하세요.  

```
aws eks describe-addon --addon-name snyk_runtime-sensor --cluster-name $CLUSTER_NAME --region $AWS_REGION
```  

응답이 아래와 유사하며, `status`가 `ACTIVE` 상태인지 확인하세요. 이 상태로 변경되는 데 몇 분이 걸릴 수 있습니다.  

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

### **Snyk 서비스 계정 토큰을 EKS 클러스터에 추가하기**  

* `aws eks`를 사용하여 `kubectl` 컨텍스트를 설정하고 클러스터를 제어할 수 있도록 합니다.  

```
aws eks update-kubeconfig --name $CLUSTER_NAME --region $AWS_REGION
```  

* `snyk-runtime-sensor` 네임스페이스 내에 `snyk-secret`이라는 이름의 시크릿을 생성하고, `snykToken`을 포함시킵니다. `snykToken`은 서비스 계정 토큰입니다.  

```
kubectl create secret generic snyk-secret \
--from-literal=snykToken=$SNYK_TOKEN \
-n snyk-runtime-sensor
```  

* AWS EKS 클러스터의 데이터는 Snyk Runtime Sensor를 통해 Snyk에 보고됩니다.  

### **Snyk Runtime Sensor 애드온 비활성화하기**  

다음 명령어를 실행하여 Snyk Runtime Sensor 애드온을 비활성화할 수 있습니다.  

```
aws eks delete-addon --addon-name snyk_runtime-sensor --cluster-name $CLUSTER_NAME --region $AWS_REGION
```  

## **Snyk Runtime Sensor 문제 해결**  

Snyk Runtime Sensor가 `is_loaded` 위험 요소를 올바르게 보고하지 않는 경우, 이는 Linux 커널의 `perf_event_paranoid` 설정이 기본값이 아닐 때 발생할 수 있습니다.  
이런 경우, Helm 차트를 설치할 때 다음 옵션을 사용하여 해결할 수 있습니다.  

* `--set securityContext.privileged=true` 옵션을 추가  
* `SYS_ADMIN`을 필수 Linux 기능(capability)으로 추가:  
  ```
  --set "securityContext.capabilities={SYS_ADMIN}"
  ```  

{% hint style="info" %}  
Loaded 패키지 위험 요소는 Debian 패키지와 같은 운영 체제 패키지에는 적용되지 않으며, npm, Maven, PyPI와 같은 패키지 관리자에서 호스팅되는 패키지에만 적용됩니다.  
{% endhint %}  

릴리스 버전은 [GitHub](https://github.com/snyk/runtime-sensor/releases)에서 확인할 수 있습니다.