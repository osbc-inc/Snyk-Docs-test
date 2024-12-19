# Helm Charts의 보안 문제를 스캔하고 수정하세요(현재 IaC)

{% hint style="info" %}
이 페이지는 현재 IaC에만 적용됩니다.
{% endhint %}

Kubernetes 구성 파일을 미스 구성 및 보안 문제에 대해 스캔하는 것 외에도 Snyk은 Helm 차트를 템플릿화하고 결과 매니페스트를 스캔하는 기능을 제공합니다. 이 템플릿 기능은 Snyk UI를 사용하여 리포지토리를 가져올 때만 사용할 수 있습니다. Snyk CLI를 사용하여 템플릿화된 Helm 차트를 스캔하는 방법에 대한 전제 조건과 안내에 대해 다음 섹션을 참조하세요. Helm 차트가 스캔된 후, Snyk은 각 템플릿 및 의존성 템플릿에 대해 프로젝트를 생성하고 구성 오류를 보고하며 이를 수정하기 위한 권장 사항을 제공합니다.

## Helm Charts에서 문제를 스캔하고 수정하기 위한 전제 조건

* 관리자는 귀하의 조직을 선호하는 Git 리포지토리와 연결하고 구성 파일을 감지할 수 있도록 설정해야 합니다. 이에 대한 설명은 [Terraform 파일에 대해 설명된 대로](../scan-terraform-files/configure-your-integration-to-find-security-issues-in-your-terraform-files-current-iac.md#configure-snyk-to-scan-your-terraform-configuration-files) 수행되어야 합니다.
* 리포지토리는 [표준 Chart 디렉터리 구조](https://helm.sh/docs/topics/charts/#the-chart-file-structure)를 따라야 합니다. 그러나 `charts/` 하위 디렉터리는 지원되지 않습니다.
* Snyk은 현재 기본 값 파일인 `values.yaml`을 사용하여 Helm 차트의 템플릿화만 지원합니다.
  * Helm 값의 특정 구성을 스캔하려면 지원되는 워크플로는 Snyk 외부에서 차트를 템플릿화하고 매니페스트를 일반 Kubernetes 파일로 스캔하는 것입니다.
  * 기본 값 파일에서 템플릿화할 수 없는 Helm 차트는 현재 지원되지 않습니다.
* 모든 차트 종속성은 구성된 Helm 저장소에서 공개로 다운로드 가능해야 합니다. 서브차트나 공개로 다운로드할 수 없는 종속성은 현재 지원되지 않습니다. 이러한 경우에는 차트를 Snyk 외부에서 템플릿화하고 매니페스트를 일반 Kubernetes 파일로 스캔하는 것이 지원되는 워크플로입니다.

## Helm Charts를 스캔하고 수정하세요

* 계정에 로그인하고 관리하려는 해당 그룹과 조직으로 이동하세요.
* 관리자에 의해 구성 파일 감지가 활성화되기 전에 테스트를 위해 리포지토리를 가져온 경우, Helm 차트를 가져오기 위해 해당 리포지토리를 다시 가져와야 합니다:

리포지토리가 스캔될 때마다, Snyk은 Helm 차트의 각 템플릿에 대해 리포지토리 별로 그룹화된 프로젝트를 생성합니다.\
구성 파일을 가져오기 위해 리포지토리를 다시 가져왔다면, Snyk은 구성 파일을 가져오고 테스트하며 이미 가져온 응용 프로그램 매니페스트 파일도 다시 테스트하여 테스트 시간을 "지금"으로 표시합니다.

* 관련 프로젝트 링크를 클릭하여 스캔 결과를 보고 구성 파일을 수정하세요.\
  외부 종속성에서 생성된 프로젝트도 스캔되고 문제가 표시됩니다.

<figure><img src="../../../../.gitbook/assets/image (208) (1) (2) (2).png" alt="헬름 차트 프로젝트 세부 정보"><figcaption><p>헬름 차트 프로젝트 세부 정보</p></figcaption></figure>

## 차트를 템플릿화하고 Kubernetes 매니페스트 스캔하기

가끔은 기본 값만 사용하여 차트를 테스트하거나 의존성이 서브차트이거나 공개로 다운로드할 수 없을 경우입니다. Snyk은 현재 이러한 시나리오를 가져오하는 동안 지원하지 않습니다. 이 섹션은 Snyk 외부에서 사용자 지정 구성을 템플릿화하고 결과 Kubernetes 매니페스트를 스캔하는 방법에 대한 가이던스를 제공합니다.

Snyk CLI와 Helm을 함께 사용할 수 있습니다:

```bash
cd /path/to/helm/chart
helm dependency update
helm template . --output-dir out
snyk iac test out/
```

`helm template`에 표준 Helm 값 플래그(예: `--set`, `--value`)를 전달하여 기본 구성이 아닌 구성을 테스트할 수 있습니다.

이 프로세스를 스크립팅하고 CLI 파이프라인에서 실행하거나 Snyk으로 가져올 수 있는 리포지토리로 helm-템플릿 파일을 만들 수 있습니다.

CLI 결과를 Snyk Web UI와 공유하려면 `--report` CLI 옵션을 사용하세요. 예를 들면:

```
snyk iac test out/ --report
```