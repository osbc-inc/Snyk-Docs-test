# Helm이 포함된 Snyk Controller의 선택적 설치 단계

설치 단계는 Snyk Controller를 환경에 맞게 구성하는 방법에 따라 다릅니다. 해당 상황에 맞는 단계를 따르십시오.

## **자체 서명 또는 추가 인증서를 사용하는 레지스트리 사용**

레지스트리가 자체 서명 또는 기타 추가 인증서를 사용하는 경우, 해당 인증서를 Snyk monitor에 사용할 수 있도록 해야 합니다. 먼저 `.crt`, `.cert`, 그리고 `.key` 파일을 디렉터리에 넣고 ConfigMap을 생성하십시오:

```
kubectl create configmap snyk-monitor-certs \
        -n snyk-monitor --from-file=<path_to_certs_folder>
```

## **보안되지 않은 컨테이너 레지스트리 또는 미지정 이미지를 사용하는 레지스트리 사용**

보안되지 않은 레지스트리를 사용하거나 레지스트리가 미지정 이미지를 사용하는 경우, `registries.conf` 파일을 제공할 수 있습니다.

```
[[registry]]
location = "internal-registry-for-example.net/bar"
insecure = true
```

형식 및 예시에 대한 자세한 정보는 [GitHub 컨테이너 이미지 문서](https://github.com/containers/image/blob/master/docs/containers-registries.conf.5.md)를 참조하십시오. 파일을 생성한 후, 다음과 같은 ConfigMap을 만들 수 있습니다:

```
kubectl create configmap snyk-monitor-registries-conf \
        -n snyk-monitor \
        --from-file=<path_to_registries_conf_file>
```

## **Snyk로의 아웃바운드 연결에 프록시 사용하기**

Snyk의 아웃바운드 연결에 프록시를 사용 중이면 통합을 해당 프록시를 사용하도록 구성해야 합니다. 프록시를 구성하려면 Helm Chart에서 제공하는 다음 값들을 설정하십시오:

* `http_proxy`
* `https_proxy`
* `no_proxy`

예를 들어:

```
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
             --namespace snyk-monitor \
             --set clusterName="운영 클러스터" \
             --set https_proxy=http://192.168.99.100:8080
```

{% hint style="info" %}
통합은 `no_proxy` 값에서 CIDR 주소 범위나 와일드카드를 지원하지 않습니다. 정확히 일치하는 것만이 지원됩니다.
{% endhint %}

## **로깅 세심도 변경하기**

로깅 세심도를 변경할 수 있습니다. 유효한 레벨은 `INFO`, `WARN`, `ERROR`, 그리고 `DEBUG`입니다. Snyk는 기본적으로 `INFO`로 설정됩니다. 로깅 세심도를 변경하려면 다음 명령을 사용하십시오:

```
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
             --namespace snyk-monitor \
             --set clusterName="운영 클러스터" \
             --set log_level="WARN"
```

## **Pod 보안 정책 활성화**

기본적으로 컨트롤러는 [Pod 보안 정책](https://kubernetes.io/docs/concepts/policy/pod-security-policy/)을 사용하지 않고 실행됩니다. 설정을 전달하여 이를 활성화할 수 있습니다:

```
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
             --namespace snyk-monitor \
             --set clusterName="운영 클러스터" \
             --set psp.enabled=true
```

기존 Pod 보안 정책을 지정하여 재사용할 수 있습니다. 이름을 지정하지 않으면 새 정책이 자동으로 생성됩니다.

```
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
             --namespace snyk-monitor \
             --set clusterName="운영 클러스터" \
             --set psp.enabled=true \
             --set psp.name=something
```

## **Snyk Controller를 PersistentVolumeClaim (PVC)로 구성**

기본 emptyDir 저장 매체 대신 PersistentVolumeClaim (PVC)를 사용하도록 Snyk Controller를 구성할 수 있습니다.

PVC를 생성하려면 Snyk 차트에서 제공하는 Helm 템플릿이나 이미 프로비저닝된 PVC를 사용할 수 있습니다.

PVC를 제어하기 위해 다음 플래그를 사용하십시오:

* `pvc.enabled` - 빈 디렉터리 대신 PVC를 사용하도록 Helm Chart에 지시합니다.
* `pvc.create` - PVC를 생성합니다. 처음으로 프로비저닝할 때 유용합니다.
* `pvc.storageClassName` - PVC의 StorageClass를 제어합니다.
* `pvc.name` - Kubernetes에서 사용할 PVC의 이름입니다.

예를 들어, 다음 명령을 실행하여 PVC를 프로비저닝하고 생성할 수 있습니다:

```
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
             --namespace snyk-monitor \
             --set pvc.enabled=true \
             --set pvc.create=true \
             --set pvc.name="snyk-monitor-pvc"
```

이미 PVC가 존재하기 때문에 이후 업그레이드에서 `pvc.create` 플래그를 제거할 수 있습니다:

```
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \             
             --namespace snyk-monitor \
             --set pvc.enabled=true \
             --set pvc.name="snyk-monitor-pvc"
```

## **특정 네임스페이스 제외하도록 Snyk Controller 구성하기**

기본적으로 Snyk은 쿠버네티스 내부적으로 간주되는 특정 네임스페이스의 스캔을 무시합니다. 전체 목록은 [제외할 네임스페이스 구성](https://github.com/snyk/kubernetes-monitor/tree/master/snyk-monitor#configuring-excluded-namespaces)을 참조하십시오.

기본 설정을 변경할 수 있으며 `excludedNamespaces` 설정으로 제외할 네임스페이스의 목록을 추가할 수 있습니다. Snyk은 기본 설정을 무시하고 제공하는 네임스페이스 목록을 사용합니다.

```
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \ 
             --namespace snyk-monitor \
             --set excludedNamespaces="{kube-node-lease,local-path-storage,some_namespace}"
```

## **Snyk Controller 리소스 구성**

컨트롤러를 배포하는 데 필요한 추가 리소스가 필요한 경우, 요청 및 제한에 대한 Helm Chart의 기본 값 설정을 `--set` 플래그로 구성할 수 있습니다.

```
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
             --namespace snyk-monitor \
             --set requests."ephemeral-storage"="50Gi"
             --set limits."ephemeral-storage"="50Gi"
```

## Snyk Controller 작업자 수 설정

작업자 수를 (XX)로 조정할 수 있습니다. `--set` 플래그를 사용하여 설정할 수 있습니다.

```
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
             --namespace snyk-monitor \
             --set workers.count="XX"
```

## Snyk Controller CPU 설정

컨트롤러를 배포하는 데 CPU가 더 많거나 적게 필요한 경우, 요청 (X) 및 제한 (Y)에 대한 HelmChart의 기본 값 설정을 `--set` 플래그로 구성할 수 있습니다.

```
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
             --namespace snyk-monitor \
             --set requests.cpu="X"
             --set limits.cpu="Y"
```

## Snyk Controller initContainers 설정 구성

기본 **emptyDir** 저장 대신 **PersistentVolumeClaim** (PVC)를 사용하는 경우 볼륨 권한을 프로비저닝해야 합니다. 이를 위해 Helm Chart에서 `--set` 플래그를 사용하여 initContainer를 비활성화할 수 있습니다:

```
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
             --namespace snyk-monitor \
             --set initContainers.enabled=false
```

{% hint style="info" %}
Openshift 플랫폼의 경우 `initContainers` 설정이 필수입니다.
{% endhint %}
