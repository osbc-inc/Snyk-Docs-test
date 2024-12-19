# JetBrains 플러그인을 사용하여 분석 실행

{% hint style="info" %}
현재 프로젝트에서 Snyk 확장이 구성되어 있고 인증되어 있으며 신뢰할 수 있도록 설정되어 있는지 확인하십시오. 이에 대한 자세한 내용은 구성 및 인증 페이지에서 설명됩니다.
{% endhint %}

다음 중 하나의 방법으로 `snyk test`를 트리거할 수 있습니다:

- 자동 (기본)
- 수동

프로젝트가 열리거나 지원되는 파일이 저장될 때 Snyk 스캔이 자동으로 트리거됩니다. 이 동작은 [기존 구성](configuration-for-the-snyk-jetbrains-plugin-and-ide-proxy.md#user-experience)을 사용하여 끌 수 있습니다.

{% hint style="info" %}
수동으로 분석을 실행하기 전에 파일을 저장해야 합니다.
{% endhint %}

`snyk test`를 수동으로 트리거하는 방법은 다음과 같습니다(아래 이미지 참조):

1. 사이드바에서 Snyk 아이콘을 클릭하여 Snyk 패널을 엽니다.
2. 플러그인 사이드바 상단의 실행(재생) 버튼을 클릭합니다.

<figure><img src="../../../.gitbook/assets/SCR-20241024-lqcw.png" alt="Snyk 분석을 수동으로 트리거하는 방법"><figcaption><p>Snyk 분석을 수동으로 트리거하는 방법</p></figcaption></figure>

재생 버튼이 회색으로 표시되면 스캔이 진행 중입니다. 다른 작업을 시작하기 전에 완료될 때까지 기다리십시오.

## 스캔 구성

회사 보안 정책을 반영하거나 특정 영역에 중점을 둘 수 있도록 스캔 동작을 사용자 정의할 수 있습니다.

### 심각도 필터

Snyk는 심각, 높음, 중간 및 낮음의 심각성을 보고합니다. 심각도를 제어하는 두 가지 방법이 있습니다:

- 플러그인 설정은 [스캔 구성](configuration-for-the-snyk-jetbrains-plugin-and-ide-proxy.md#scan-configuration)
- 다음 화면 이미지에서 볼 수 있는 Snyk 패널의 문제 상단의 작은 버튼

기본적으로 모든 수준이 선택되어 있습니다. 하나 이상을 선택해야 합니다.

<figure><img src="../../../.gitbook/assets/SCR-20241024-mfpi.png" alt="심각도 레벨 선택" width="375"><figcaption><p>심각도 레벨 선택</p></figcaption></figure>

Snyk 심각성 아이콘의 의미는 다음과 같습니다:

| ![](../../../.gitbook/assets/image (201) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1).png) | 크리티컬 심각도                                                                                                   | 공격자가 민감한 데이터에 액세스하고 애플리케이션에서 코드를 실행할 수 있을 수 있습니다.                                                               |
| --------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| ![](../../../.gitbook/assets/image (10) (1) (1) (2) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (5) (3).png)    | 높은 심각도 | 공격자가 애플리케이션에서 민감한 데이터에 액세스할 수 있을 수 있습니다.                                                                            |
| ![](../../../.gitbook/assets/image (116) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (5) (6).png)                                          | 중간 심각도                                                                                                                                                                           | 일부 조건 하에서 공격자가 애플리케이션에서 민감한 데이터에 액세스할 수 있을 수 있습니다.                                                      |
| ![](../../../.gitbook/assets/image (114) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1).png) | 낮은 심각도                                                                                                         | 응용프로그램이 취약성 매핑을 허용하고 다른 취약성과 함께 사용될 수 있는 일부 데이터를 노출할 수 있습니다. |

### 문제 유형별로 필터링

Snyk는 다음 유형의 문제를 보고합니다:

- **오픈 소스** 문제: 오픈 소스 종속성에서 발견됩니다. [아래 섹션](run-an-analysis-with-the-jetbrains-plugins.md#snyk-open-source-issues)에서 자세한 내용을 확인하십시오.
- **코드 보안** 문제: 응용 프로그램의 소스 코드에서 발견됩니다. [아래 섹션](run-an-analysis-with-the-jetbrains-plugins.md#snyk-code-security-vulnerabilities-and-quality-issues)에서 자세한 내용을 확인하십시오.
- **코드 품질** 문제: 응용 프로그램 소스 코드에서 발견됩니다. [아래 섹션](run-an-analysis-with-the-jetbrains-plugins.md#snyk-code-security-vulnerabilities-and-quality-issues)에서 자세한 내용을 확인하십시오.
- **{{인프라스트럭처 애스 코드}}** 문제: IaC 파일에서 발견됩니다. [아래 섹션](run-an-analysis-with-the-jetbrains-plugins.md#snyk-infrastructure-as-code-issues)에서 자세한 내용을 확인하십시오.
- **컨테이너** 문제: Kubernetes 워크로드 파일에서 가져온 이미지에서 발견됩니다. [아래 섹션](run-an-analysis-with-the-jetbrains-plugins.md#snyk-container-issues)에서 자세한 내용을 확인하십시오.

{% hint style="info" %}
정확한 기능 및 사용 가능한 스캐너는 요금제에 따라 다릅니다. IDE 플러그인에서 이들 중 하나를 구성하기 전에 귀하의 조직 관리자가 모든 Snyk 제품을 활성화했는지 확인하십시오.
{% endhint %}

특정 문제 유형을 표시하거나 숨기는 두 가지 방법이 있습니다:

- 플러그인 설정은 [스캔 구성](configuration-for-the-snyk-jetbrains-plugin-and-ide-proxy.md#scan-configuration)
- 다음 화면 이미지에서 볼 수 있는 패널 사이드바의 필터 버튼

기본적으로 모든 문제 유형이 선택되어 있습니다.

<figure><img src="../../../.gitbook/assets/SCR-20241024-miah.png" alt="특정 문제 유형 표시 또는 숨기기"><figcaption><p>특정 문제 유형 표시 또는 숨기기</p></figcaption></figure>

### 새 문제 대 전체 문제

[2.10.0 버전](https://plugins.jetbrains.com/plugin/10972-snyk-security/versions/stable/623034)부터는 **신규 도입된 문제만** 볼 수 있습니다.

이 기능은 노이즈를 _**줄여주고**_ 개발자가 현재 변경 사항에만 _**집중할 수 있도록**_ 합니다. 개발자들은 이 조치로 초기에 문제를 방지하여 CI/CD 파이프라인을 해제하고 제공 속도를 높일 수 있습니다.

이 논리는 지역 Git 저장소를 사용하고 현재 발견된 결과와 베이스 브랜치에서 제외된 결과를 보여줍니다. 

[스캔 구성 설정을 사용하여 구성](configuration-for-the-snyk-jetbrains-plugin-and-ide-proxy.md#scan-configuration)할 수 있으며 Net New Issues는 기본적으로 꺼져 있으므로 켜려면 수동 조치를 취해야 합니다.

이 기능이 활성화된 후 Snyk은 증분 결과만 보고합니다.

새로 만든 기능 브랜치의 경우 보고된 문제가 없습니다. 이것은 개발자들이 목표로 하는 의도된 상태이며, 다음 화면 이미지를 참조하십시오:

<figure><img src="../../../.gitbook/assets/SCR-20241024-ngbm.png" alt="성공적인 상태. 새로운 문제 없음"><figcaption><p>성공적인 상태. 새로운 문제가 없습니다.</p></figcaption></figure>

기본 브랜치는 일반적으로 각 Git 저장소에 대해 자동으로 결정됩니다. 

고급 경우에는 개발자가 다음 단계를 따라 기본 브랜치를 변경할 수 있습니다(다음 화면 이미지 참조): 

1. 이슈 트리의 최상위 노드를 클릭합니다.
2. 드롭다운 선택을 사용하여 모든 브랜치를 선택합니다.
3. 선택을 저장하려면 확인을 클릭합니다.

<figure><img src="../../../.gitbook/assets/SCR-20241024-nfhj.png" alt="Net New 문제를 계산하기 위해 기본 브랜치 변경"><figcaption><p>Net New 문제를 계산하기 위해 기본 브랜치 변경</p></figcaption></figure>

## 사용 가능한 Snyk 문제 유형

### Snyk 코드 보안 취약점 및 품질 문제

{{Snyk 코드}} 분석은 응용 프로그램 코드에서 발견된 보안 취약점과 코드 품질 문제 목록을 보여줍니다.

{% hint style="info" %}
2025년 6월 24일부터는 {{Snyk 코드}} 품질 문제가 더 이상 제공되지 않습니다.
{% endhint %}

자세한 정보 및 문제를 해결하기 위해 다른 사람이 사용한 예제에 대해 알아보려면 보안 취약점 또는 코드 보안 문제를 선택하십시오.

<figure><img src="../../../.gitbook/assets/SCR-20241024-npba.png" alt=""><figcaption><p>{{Snyk 코드}} 문제 세부 정보</p></figcaption></figure>

### Snyk 오픈 소스 문제

{{Snyk 오픈 소스}} 분석은 모든 매니페스트 파일에서 발견된 취약점 및 라이선스 문제 목록을 보여줍니다. 취약점 또는 라이선스 문제를 선택하여 자세한 정보를 확인할 수 있습니다.

<figure><img src="../../../.gitbook/assets/SCR-20241024-nrsk.png" alt="{{Snyk 오픈 소스}} 문제 세부 정보"><figcaption><p>{{Snyk 오픈 소스}} 문제 세부 정보</p></figcaption></figure>

### Snyk 인프라스트럭처 애스 코드 문제

각 스캔에서 Snyk IaC 분석은 Terraform, Kubernetes, AWS CloudFormation 및 Azure Resource Manager (ARM) 코드에서 발견된 문제를 보여줍니다. 이 스캔은 Snyk CLI에 기반하며 로컬 개발에 빠르고 편리합니다. 자세한 정보를 보려면 문제를 선택하십시오.

<figure><img src="../../../.gitbook/assets/SCR-20241024-ntcr.png" alt="Snyk IaC 문제 세부 정보"><figcaption><p>Snyk IaC 문제 세부 정보</p></figcaption></figure>

### Snyk 컨테이너 문제

{% hint style="info" %}
2025년 6월 24일 이후 릴리스된 버전에서는 Snyk JetBrains IDE 플러그인이 컨테이너 이미지를 탐지하지 않습니다.
{% endhint %}

JetBrains 플러그인은 Kubernetes 구성 파일을 스캔하고 컨테이너 이미지를 찾습니다. 추출된 컨테이너 이미지를 사용하여 빠르게 취약성을 찾아내며 [Snyk 취약성 데이터베이스](https://security.snyk.io)의 최신 정보와 비교 분석합니다.

{{Snyk 컨테이너}} 분석은 이미지에 영향을 줄 수 있는 각 보안 취약점을 보여줍니다. 보다 자세한 정보를 보려면 취약성을 선택하십시오.

비교 테이블은 심각도 수준(크리티컬 또는 높음)과 함께 표시됩니다. 이는 현재 이미지와 Snyk에서 추천하는 이미지 간의 취약성 차이를 보여주며, 동일한 특성으로 분류된 심각성에 따라 이미지를 업그레이드하고 프로덕션에서 실행 중인 이미지에 대한 신뢰 수준을 높일지 결정하는 데 도움이 됩니다.