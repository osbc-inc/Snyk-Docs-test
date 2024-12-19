# [자산 인벤토리 내 백스테이지 파일 - 사용 사례](./#backstage-file-for-scm-integrations)

[백스테이지 카탈로그](./#backstage-file-for-scm-integrations)를 구성한 후, Snyk AppRisk가 저장소 자산을 [인벤토리 레이아웃의 모든 자산](../../../manage-assets/#inventory-overview) 탭에서 찾은 데이터로 보강합니다. 이때 백스테이지의 `catalog-info.yaml` 파일에서 찾은 데이터를 사용합니다.

백스테이지 카탈로그를 사용하여 저장소 자산을 보강하고 구성 요소 엔티티를 정의하세요. 이런 유형의 상황에서 구성 요소는 서비스, 저장소, 웹사이트, 라이브러리 등과 같은 소프트웨어 구성 요소로 정의됩니다.&#x20;

구성 요소는 여러 속성을 가지며 그 중 대부분은 선택 사항입니다:

- `spec.type` (필수) - 저장소의 분류를 나타냅니다.&#x20;
- `spec.owner` (필수) - 저장소를 소유한 팀을 나타냅니다.
- `spec.lifecycle` - 구성 요소의 라이프사이클 상태를 나타냅니다. 예를 들어 `production`, `experimental`, `deprecated` 등이 있습니다.
- `spec.system` (선택 사항) - 동일한 목적을 제공하는 구성 요소 그룹을 나타냅니다. 이 개념은 "애플리케이션"이라고 합니다.
- `Metadata.name` (필수) - 구성 요소의 이름을 나타냅니다.
- `Metadata.title` (선택 사항) - 구성 요소의 이름을 나타냅니다.

백스테이지 데이터는 동적이며 시간이 흐름에 따라 변할 수 있습니다:

- `catalog-info.yaml` 파일에 새로운 커밋이나 업데이트가 이루어지면 Snyk AppRisk가 해당 특정 저장소 자산의 속성을 업데이트합니다.
- 저장소에서 `catalog-info.yaml` 파일이 제거되면 Snyk AppRisk가 해당 특정 저장소 자산의 속성을 삭제합니다.

{% hint style="info" %}
점(`.`)을 포함하는 키를 이스케이프 하기 위해 쌍따옴표(`""`)를 사용할 수 있습니다. 예를 들어`"`[`example.com`](http://example.com/)`".owner`와 같이 사용할 수 있습니다.
{% endhint %}

## 인벤토리 메뉴와 백스테이지 파일&#x20;

통합 허브 구성 메뉴에서 선택한 내용에 따라 인벤토리 메뉴의 필터에만 해당 선택 사항이 표시됩니다. 예를 들어, 카테고리 속성을 선택한 경우에는 필터 목록에도 표시됩니다.

## 자산 요약 탭과 백스테이지 파일&#x20;

자산 요약 탭에는 백스테이지와 통합하도록 선택한 여섯 가지 속성만 표시됩니다.

## 자산 속성 탭과 백스테이지 파일&#x20;

자산 속성 탭에서는 선택한 속성만 저장소 자산에 메타데이터로 추가해야 합니다.

```
{
    name:"spring.goof",
    repositoryURL:"https://github.com/snyk/spring.goof.git",
    context:[
             {
              name: "super-duper-component",
              title: "Super Duper Component",
	      application: "super-duper-app",
	      lifecycle: "production",
	      owner: "super-duper-team",
	      category: "service",
              source: "Backstage"
              }]
}
```

## 정책 필터와 백스테이지 파일&#x20;

정책 빌더에서는 이전에 백스테이지 카탈로그 파일을 구성할 때 선택한 속성만 찾을 수 있습니다.&#x20;

다음 목록은 백스테이지 카탈로그 파일을 구성할 때 선택할 수 있는 모든 가능한 백스테이지 속성을 설명합니다.&#x20;

- **애플리케이션** - 동일한 목적을 제공하는 구성 요소 그룹을 나타냅니다.&#x20;
- **소유자** - 저장소를 소유한 팀을 지정합니다.
- **카탈로그 이름** - 메타데이터 이름입니다.
- **제목** - 개체에 표시할 이름으로, 카탈로그 이름이 읽기 어려울 때 대체로 사용됩니다.
- **카테고리** - 저장소의 분류를 나타냅니다. 조직은 임의의 이름 또는 텍스트를 선택할 수 있습니다.
- **라이프사이클** - 구성 요소의 라이프 사이클 상태를 나타냅니다. 예를 들어 production, experimental, deprecated 등이 있습니다.

다음 비디오는 통합 허브의 백스테이지 파일 옵션을 개요로 제공하며 사용 가능한 속성에 대한 간단한 설명을 제공합니다:

{% embed url="https://www.youtube.com/watch?v=gfOUSE0UhHA" %}
비디오가 마음에 드셨나요? [Snyk Learn](https://learn.snyk.io/catalog/?type=product-training\&topics=AppRisk)에서 수강할 수 있는 나머지 코스를 확인해보세요!
{% endembed %}