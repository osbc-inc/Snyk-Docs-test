# 커버리지 제어 정책 - 사용 사례

**정책 보기**에서 **커버리지 제어 설정** 작업을 사용하여 애플리케이션 자산을 식별하고 모니터링하며 위험을 우선 순위로 지정할 수 있습니다. 하나 이상의 보안 제품을 선택하고 마지막 스캔이 이루어져야 하는 시점을 지정할 수 있습니다.

자산을 식별하고 커버리지 정책을 설정하면 특정 보안 제어가 있는 곳을 정의할 수 있습니다.

다음 예는 필요로 하는 자산과 그의 커버리지 정책을 설정하는 예시입니다.

<figure><img src="../../../../.gitbook/assets/image (2) (10).png" alt="AppRisk - 커버리지 제어 정책 설정"><figcaption><p>Snyk AppRisk - 커버리지 제어 정책 설정</p></figcaption></figure>

예시를 따르기 위해서 다음과 같은 필터를 적용해야 합니다:

* **필터 1**: 리포지토리인 자산만 포함합니다.

<figure><img src="../../../../.gitbook/assets/image (3) (5).png" alt="자산 유형을 위한 필터 구성" width="350"><figcaption><p>자산 유형을 위한 필터 구성</p></figcaption></figure>

* **필터 2**: 애플리케이션에 관련된 코딩 언어 태그가 있는 자산을 포함합니다 (Apex, ASP, C, C#, C++, CMake, Go, HTML, Java, JavaScript, Kotlin, PHP, Python, Ruby, Scala, Swift, TypeScript, VisualBasic, Handlebars, Makefile, Objective-C).

<figure><img src="../../../../.gitbook/assets/image (4) (7).png" alt="태그를 위한 필터 구성" width="354"><figcaption><p>태그를 위한 필터 구성</p></figcaption></figure>

이후에는 예시의 두 보안 제어에 초점을 맞추기 때문에 Snyk Open Source와 Snyk Code에 각각 두 개의 작업을 정의해야 합니다.

* **작업 1**: Snyk Open Source가 선택된 **커버리지 제어 정책 설정** 작업 추가.

    기본적으로 Snyk에서 SCM 통합을 통해 가져온 오픈 소스 manifest 파일은 매일 스캔됩니다. 커버리지 제어 정책은 해당 리포지토리에 대해 최근 2일 이내에 Snyk Open Source 스캔이 이루어졌는지 확인해야 합니다.

<figure><img src="../../../../.gitbook/assets/image (5) (3).png" alt="커버리지 제어를 위한 필터 구성" width="352"><figcaption><p>커버리지 제어를 위한 필터 구성</p></figcaption></figure>

* **작업 2**: Snyk Code가 선택된 **커버리지 제어 정책 설정** 작업 추가.

    Snyk Code의 경우, 기본적으로 주 1회 또는 모니터링 중인 브랜치에 변경이 발생할 때 스캔이 수행됩니다. 커버리지 제어 정책은 해당 리포지토리에 대해 최근 1주일 이내에 Snyk Code 스캔이 이루어졌는지 확인해야 합니다.

<figure><img src="../../../../.gitbook/assets/image (6) (6) (1).png" alt="커버리지 제어를 위한 필터 구성" width="350"><figcaption><p>커버리지 제어를 위한 필터 구성</p></figcaption></figure>

인벤토리 보기에서 커버리지 갭이 있는 경우 제어 커버리지 아이콘에 취소선이 표시됩니다. 각 아이콘에 대한 자세한 내용은 [인벤토리 기능](../../../../manage-assets/assets-inventory-components.md) 페이지에서 확인할 수 있습니다.

다음 비디오는 커버리지 정책을 구성하는 방법에 대해 설명합니다:

{% embed url="https://youtu.be/Tmoc_qQ8qA8" %}
비디오가 도움이 되었습니까? [Snyk Learn](https://learn.snyk.io/lesson/snyk-apprisk-essentials/)에서 나머지 코스를 확인해보세요!{% endembed %}  