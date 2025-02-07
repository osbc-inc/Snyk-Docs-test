# Amazon Elastic Kubernetes Service(Amazon EKS)에 Snyk 컨트롤러 설치하기

{% hint style="info" %}
**Snyk** Controller를 설치하기 전에 [**Snyk** Controller 설치 전 필수 요구 사항](./#prerequisites-for-installing-the-snyk-controller)을 검토했는지 확인하십시오.
{% endhint %}

{% hint style="info" %}
이 설치 단계는 동일한 AWS 계정의 EKS 및 ECR에서 잘 작동합니다. 다른 설정을 사용하시는 경우 [**Snyk** 지원팀](https://support.snyk.io)에 문의하십시오.
{% endhint %}

**Snyk** Controller를 설치하면 실행 중인 EKS 워크로드를 가져와 테스트하고 해당 이미지와 구성에 있는 취약점을 식별할 수 있으며, 이는 워크로드를 덜 안전하게 만들 수 있습니다. 워크로드가 가져오 된 후에는 **Snyk**이 추가적인 보안 문제를 식별하여 지속적으로 워크로드를 모니터링하며 새 이미지가 배포되고 워크로드 구성이 변경될 때 보안 문제를 식별합니다.

아래 설명된 단계는 **Snyk** Controller를 구성하여 ECR에서 프라이빗 이미지를 가져와 스캔하는 방법에 대한 지침을 제공합니다.

Amazon EKS를 설치하려면:

1\. Kubernetes 환경에 액세스하십시오. 다음 명령어를 실행하여 **Snyk** Charts 리포지토리를 Helm에 추가하십시오:

```
helm repo add snyk-charts https://snyk.github.io/kubernetes-monitor --force-update
```

2\. 리포지토리를 추가한 후, **Snyk** Controller를 위한 고유한 네임스페이스를 생성하십시오:

```
kubectl create namespace snyk-monitor
```

{% hint style="info" %}
Kubernetes 응용 프로그램의 좋은 관행으로 컨트롤러 리소스를 쉽게 격리하기 위해 고유한 네임스페이스를 사용하십시오.

`"snyk-monitor"` 네임스페이스를 기억하십시오. 다른 리소스를 구성할 때 사용합니다.
{% endhint %}

3\. **dockercfg.json** 파일을 만들고 다음 예제와 일치하는지 확인하십시오:

```
{
  "credsStore": "ecr-login"
}
```

프라이빗 레지스트리에 대한 추가 설정은 [프라이빗 컨테이너 레지스트리 인증](authenticate-to-private-container-registries.md)를 참조하십시오.

4\. 통합 ID, 서비스 계정 토큰 및 **dockercfg.json** 파일을 포함하는 Kubernetes 시크릿 생성:

```
kubectl create secret generic snyk-monitor \
        -n snyk-monitor --from-file=dockercfg.json \
        --from-literal=integrationId=abcd1234-abcd-1234-abcd-1234abcd1234 \
        --from-literal=serviceAccountApiToken=bdca4123-dbca-4343-bbaa-1313cbad4231
```

5\. 노드에 정책 또는 역할을 연결하십시오. 다음 중 하나를 사용하여 수행할 수 있습니다:

#### Worker 노드에 정책 첨부

1. `NodeInstanceRole` 정책 첨부. [Amazon ECR 이미지와 Amazon EKS 사용](https://docs.aws.amazon.com/AmazonECR/latest/userguide/ECR_on_EKS.html)을 참조하십시오.
2. EKS worker 노드에 `AmazonEC2ContainerRegistryReadOnly` 정책을 첨부하십시오.\
   **Snyk** Controller는 이러한 worker 노드에서 실행 중일 때 프라이빗 이미지를 가져올 수 있습니다.

#### **노드 그룹용 EKS 노드 역할 생성**

1. [Amazon EKS 노드 IAM 역할](https://docs.aws.amazon.com/eks/latest/userguide/create-node-role.html)에서 지침을 따르고 기존 노드 역할을 확인하십시오. `AmazonEC2ContainerRegistryReadOnly` 정책을 첨부했는지 확인하십시오.
2. EKS 노드 그룹 페이지의 **세부 정보** 탭으로 이동하여 `노드 IAM 역할 ARN`을 확인할 수 있습니다.

```
arn:aws:iam::<role-id>:role/<role-name>
```

3. 다음 내용과 같은 \<newFile>.yaml 파일을 만드십시오:

```
volumes:
  projected:
    serviceAccountToken: true
    
securityContext:
  fsGroup: 65534

rbac:
  serviceAccount:
    annotations:
      eks.amazonaws.com/role-arn: <Node IAM 역할 ARN>
```

4. **Snyk** Controller를 설치하십시오.

서비스 계정을 위해 IAM 역할을 생성한 후, 새로 생성된 YAML 파일을 사용하여 Helm 차트의 값들을 덮어쓰기하여 **Snyk** Controller를 설치할 수 있습니다.

```
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
             --namespace snyk-monitor \
             --set clusterName=<클러스터 이름 입력> \
             -f <newFile>.yaml
```
