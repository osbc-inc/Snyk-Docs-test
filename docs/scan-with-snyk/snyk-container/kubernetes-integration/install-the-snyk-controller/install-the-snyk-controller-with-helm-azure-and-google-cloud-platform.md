# Helm을 사용하여 Snyk Controller 설치하기 (Azure 및 Google Cloud Platform)

{% hint style="info" %}
[Snyk Controller를 설치하는 데 필요한 전제 조건](./#prerequisites-for-installing-the-snyk-controller)을 검토했는지 확인하십시오.
{% endhint %}

Kubernetes 워크로드에 대한 취약성 세부 정보를 받으려면 Snyk 관리자가 먼저 클러스터에 Snyk Controller를 설치해야 합니다. Snyk Controller는 [Helm Hub](https://hub.helm.sh/charts/snyk/snyk-monitor)에서 발행됩니다.

설치 단계는 다음을 다룹니다:

* 대부분의 Kubernetes 플랫폼에 대한 Snyk 통합
* Amazon Elastic Kubernetes Service (EKS) 클러스터에서 Amazon Elastic Container Registry (ECR)를 사용할 때 통합을 위한 추가 구성 단계

## Helm을 사용한 Snyk Controller의 필수 설치 단계

다음과 같이 Helm을 사용하여 Snyk Controller를 설치합니다:

1. Kubernetes 환경에 접속하십시오. 다음 명령을 실행하여 Snyk 차트 리포지토리를 Helm에 추가하십시오:

   <pre><code><strong>helm repo add snyk-charts https://snyk.github.io/kubernetes-monitor --force-update
   </strong></code></pre>

2. 리포지토리가 추가된 후, Snyk Controller를 위한 고유한 이름 공간을 생성하십시오:

   ```
   kubectl create namespace snyk-monitor
   ```

{% hint style="info" %}
Kubernetes 애플리케이션에 대한 좋은 관행으로, 컨트롤러 리소스를 쉽게 분리하기 위해 고유한 이름 공간을 사용하십시오.&#x20;

`Snyk-monitor` 이름 공간을 기억해야 합니다. 다른 리소스를 구성할 때 사용합니다.
{% endhint %}

Snyk 모니터에는 다음이 필요합니다:

* Snyk 통합 ID
* 서비스 계정 토큰
* dockercfg.json 파일

### 공용 컨테이너 레지스트리 설치

공용 컨테이너 레지스트리를 설치하려면 `snyk-monitor`라는 Kubernetes 시크릿을 생성해야 합니다. 이 시크릿에는 Snyk 통합 ID와 서비스 계정 토큰이 포함되어 있어야 합니다.

이를 위해 다음 명령을 실행하십시오:

```
kubectl create secret generic snyk-monitor -n snyk-monitor \
        --from-literal=dockercfg.json={} \
        --from-literal=integrationId=abcd1234-abcd-1234-abcd-1234abcd1234 \
        --from-literal=serviceAccountApiToken=bdca4123-dbca-4343-bbaa-1313cbad4231
```

{% hint style="info" %}
성공적인 통합을 위해서는 시크릿을 `snyk-monitor`로 명명해야 합니다.
{% endhint %}

### Snyk Helm 차트 설치

```
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
             --namespace snyk-monitor \
             --set clusterName="Production cluster"
```

자체 Snyk 인스턴스를 실행 중인 경우 컨트롤러를 설치할 때 API 엔드포인트를 지정해야 합니다.

다음 명령에서는 Snyk 인스턴스의 전체 호스트 이름을 제공하십시오.

```
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
             --namespace snyk-monitor \
             --set clusterName="Production cluster" \
             --set integrationApi=https://<server>/kubernetes-upstream
```

* `"Production cluster"`를 모니터링 중인 클러스터에 기반하여 이름으로 바꿔야 합니다. 나중에 Snyk에서 워크로드를 찾을 때 이 레이블을 사용할 수 있습니다.
* 클러스터 이름에 `/` (슬래시)를 사용하는 것은 허용되지 않습니다. 클러스터 이름에 포함된 모든 `/`은 제거됩니다.
* 업데이트 시 클러스터의 이름을 변경하지 않으려면 Helm의 `--reuse-values` 옵션을 사용할 수 있습니다. 업그레이드할 때 Helm은 마지막 릴리스에서 값들을 재사용하고 `--set` 및 `-f`를 사용하여 명령 줄에서 재정의된 값을 병합합니다.

### 관리 ID를 사용하여 AKS와 ACR 통합

다음을 수행하려면:

1. AKS를 사용하여 ACR에 액세스 권한 부여를 위해 사용자가 관리하는 ID를 사용하고 있고 VM 스케일 세트에 `AcrPull` 역할을 할당하는 여러 ID가 있는 경우, 사용할 원하는 사용자가 관리 ID의 클라이언트 ID도 지정해야 합니다. 이 값을 `.Values.azureEnvVars`에서 오버라이드로 설정해야 합니다:

   ```yaml
   azureEnvVars:
     - name: AZURE_CLIENT_ID
       value: "abcd1234-abcd-1234-abcd-1234abcd1234"
   ```

2. 위의 YAML을 `override.yaml`에 저장한 후, 다음 명령을 실행하십시오:

   ```bash
   helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
     --namespace snyk-monitor \
     -f override.yaml
   ```

기본적으로 이 값은 빈 문자열로 설정되어 있으며 사용되지 않습니다.

{% hint style="info" %}
시스템 관리 ID를 사용하여 `AcrPull` 역할을 할당하는 경우, 이 변수를 설정할 필요가 없습니다.
{% endhint %}

## 기존 설치 업데이트

기존 고객이며 Snyk Controller를 업데이트하는 경우:

1. 서비스 계정 토큰을 생성하십시오. 자세한 내용은 [Snyk Controller 설치 전제 조건](./#prerequisites-for-installing-the-snyk-controller)을 참조하십시오. 이 토큰은 `snyk-monitor` 시크릿에 저장됩니다.
2. 기존 `snyk-monitor` 시크릿을 삭제하십시오:

   ```shell
   kubectl delete secret snyk-monitor -n snyk-monitor
   ```

3. [필수 설치 단계](install-the-snyk-controller-with-helm-azure-and-google-cloud-platform.md#mandatory-installation-steps-for-the-snyk-controller-with-helm)를 따르십시오. 최신 Helm 차트 버전을 얻으려면 다음 명령을 실행하십시오:

   <pre><code><strong>helm repo add snyk-charts https://snyk.github.io/kubernetes-monitor --force-update
   </strong></code></pre>