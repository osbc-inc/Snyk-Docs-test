# Snyk 프로젝트

Snyk 프로젝트 정보는 **Projects** 목록에서 나타납니다. Snyk 대시보드의 메뉴에서 표시할 수 있습니다. 추가할 수 있는 필터는 오른쪽의 드롭다운에서 선택한 **그룹화 기준**에 따라 다릅니다. Origin 또는 소스로 필터링하려면 통합 필터를 사용하세요.

저장소 및 컨테이너 레지스트리 이미지 가져오기 상태에 대한 진행 및 오류 정보는 [Import Log](import-log.md)에서 표시됩니다.

{% hint style="info" %}
프로젝트 목록 페이지에 필터가 적용된 후 해당 URL을 즐겨찾기에 추가하고 조직의 다른 사용자와 공유할 수 있습니다. 이를 통해 모든 사용자가 페이지의 동일한 보기를 볼 수 있습니다.
{% endhint %}

<figure><img src="../../.gitbook/assets/Screenshot 2023-01-24 at 09.09.25.png" alt="Snyk 프로젝트 목록을 Target별로 그룹화한 것"><figcaption><p>Snyk 프로젝트 목록을 Target별로 그룹화한 것</p></figcaption></figure>

Snyk 프로젝트 개념에는 다음이 포함됩니다:

- [Target](./#target)
- [Origin](./#origin)
- [Project](./#project)
- [Targetfile](./#targetfile)
- [Type](./#type)

## Target

프로젝트는 Target에 보관됩니다. Target은 Snyk가 스캔한 외부 리소스를 나타냅니다: 코드 저장소, Kubernetes 작업로드 또는 Snyk 외부에서 스캔 가능한 다른 리소스 등.

Snyk는 프로젝트 가져오기를 요청하거나 CLI를 사용하여 스캔할 때 Target을 생성합니다. 가져오기가 실패하거나 아무것도 찾지 못한 경우 Target은 비어 있을 것입니다.

**Target으로 그룹화**를 선택하면 **Projects** 목록에 Snyk Targets이 나타납니다. 또한 [조직별 Target 가져오기](../../snyk-api/reference/targets.md#orgs-org_id-targets)와 [주어진 Org ID에 대한 모든 프로젝트 나열](../../snyk-api/reference/projects.md#orgs-org_id-projects) 엔드포인트를 사용하여 Targets을 찾을 수 있습니다.

목록에서 각 Target 우측의 **점**을 클릭하여 **사용 가능한 동작**을 볼 수 있습니다; 이는 Target을 **삭제**하는 등의 작업을 포함합니다.

<figure><img src="../../.gitbook/assets/Screenshot 2023-01-24 at 08.59.20.png" alt="Snyk Target 및 해당 Target의 프로젝트"><figcaption><p>Snyk Target 및 해당 Target의 프로젝트</p></figcaption></figure>

각 Snyk 프로젝트는 부모 Target과 연관되어 있습니다. 하나의 Target에는 여러 프로젝트가 포함될 수 있습니다. Target의 구조는 Origin에 따라 달라집니다.

그룹화 옵션은 필터링 속성이 Target 또는 프로젝트 수준에서 적용되는지를 제어합니다. **그룹화 없음**은 개별 프로젝트에 대해 [태그](../introduction-to-snyk-projects/project-tags.md) 및 [프로젝트 수준의 필터링 속성](project-attributes.md)을 적용할 수 있도록 합니다.

Snyk는 페이지 로딩 시간을 개선하기 위해 페이지네이션 및 수백만 개의 프로젝트를 스캔해야 하는 경우 특히 유용한 필터링을 제공합니다.

**정렬 기준** (**우측의 드롭다운**)을 사용하여 **Projects** 목록을 심각도, 프로젝트 가져오기 시기 또는 알파벳순으로 정렬할 수 있습니다.

<figure><img src="../../.gitbook/assets/image (2) (5).png" alt="Target으로 그룹화할 때 사용 가능한 정렬 속성"><figcaption><p>Target으로 그룹화할 때 사용 가능한 정렬 속성</p></figcaption></figure>

## Origin 또는 소스

Origin은 CLI, GitHub 또는 Kubernetes와 같은 Target 생태계를 정의합니다. Origin은 Targets의 속성이며 (**이전 섹션 참조**), Target 이름 옆에 아이콘으로 나타납니다.

<figure><img src="../../.gitbook/assets/Screenshot 2023-01-24 at 08.59.07.png" alt="Target 이름 옆에 Origin 아이콘"><figcaption><p>Target 이름 옆에 Origin 아이콘</p></figcaption></figure>

가능한 Origin 값은 다음과 같습니다:

- acr
- api
- artifactory-cr
- aws-config
- aws-lambda
- azure-functions
- azure-repos
- bitbucket-cloud
- bitbucket-server
- cli
- cloud-foundry
- digitalocean-cr
- docker-hub
- ecr
- gcr
- github
- github-cr
- github-enterprise
- gitlab
- gitlab-cr
- google-artifact-cr
- harbor-cr
- heroku
- ibm-cloud
- kubernetes
- nexus-cr
- pivotal
- quay-cr
- terraform-cloud

## Project

Snyk 프로젝트는 특정 Target에서 스캔할 매니페스트 파일과 같은 항목을 정의합니다. 스캔을 실행하는 방법을 정의하는 구성 정보가 포함되어 있습니다.

프로젝트는 **Projects** 목록에 나타납니다. [조직별 주어진 Org ID의 모든 프로젝트 나열](../../snyk-api/reference/projects.md#orgs-org_id-projects) 엔드포인트를 사용하여 프로젝트를 찾을 수도 있습니다.

개선된 프로젝트 가시성 및 [프로젝트 수준의 필터링 속성을 적용하려면 **그룹화 없음**](project-attributes.md)을 사용하세요.

<figure><img src="../../.gitbook/assets/Screenshot 2023-01-23 at 18.07.46 (1) (1) (1) (1) (1) (1).png" alt="프로젝트 수준에서 적용된 필터링 속성"><figcaption><p>프로젝트 수준에서 적용된 필터링 속성</p></figcaption></figure>

## Targetfile

Targetfile은 Target에서 스캔할 특정 항목을 나타냅니다. 예를 들어 GitHub 리포지토리의 `pom.xml` 파일입니다.

{% hint style="info" %}
[{{Snyk Code}}](../../scan-with-snyk/snyk-code/) 스캔은 Targetfile을 사용하지 않습니다.
{% endhint %}

## Type

Type은 특정 프로젝트에 사용할 스캔 방법을 나타냅니다. 예를 들어, Snyk Code를 사용하는 SAST (정적 응용 프로그램 보안 테스트) 또는 Snyk 오픈 소스를 사용하는 Maven 프로젝트에 대한 Maven과 같은 것입니다. 이는 스캔 구성의 일부입니다.

## 프로젝트 목록 페이지에서 프로젝트 작업

### 삭제, 활성화 또는 비활성화

프로젝트에 대해 일괄 작업을 수행하려면 먼저 프로젝트를 선택한 다음 **프로젝트를 삭제, 활성화 또는 비활성화**를 선택하세요.&#x20;

<figure><img src="../../.gitbook/assets/Screenshot 2023-11-15 at 10.42.37.png" alt=""><figcaption><p>일괄적으로 프로젝트 삭제하기</p></figcaption></figure>

### 테스트 빈도 설정

각 프로젝트의 테스트 빈도를 설정할 수 있습니다.&#x20;

각 항목에서 해당 프로젝트의 테스트 빈도 (`never`, `daily`, 또는 `weekly`)를 선택할 수 있습니다. Open Source, Code Analysis, Container 또는 IaC 유형에 적용되는 것에 따라 해당 프로젝트의 테스트 빈도를 선택하세요.

<div align="left"><figure><img src="../../.gitbook/assets/Screenshot 2023-11-15 at 10.42.46.png" alt="테스트 빈도 선택"><figcaption><p>테스트 빈도 선택</p></figcaption></figure></div>

{% hint style="info" %}
기본 테스트 빈도 및 제한 사항은 다음과 같습니다:

- Open Source: 기본값은 매일입니다.
- Code Analysis 프로젝트: 기본값은 매주이며, 매일은 사용할 수 없습니다. 코드를 매일 테스트하려면 [Snyk 지원팀에 요청](https://support.snyk.io)하세요.
- Container: 기본값은 매일입니다.
- IaC: 기본값은 매주입니다.
{% endhint %}

**비활성화**를 클릭하여 테스트를 하지 않고, 웹훅을 제거하고 프로젝트 결과를 보고서에 표시하지 않도록 설정할 수 있습니다.