# 종속성 및 라이선스

귀하의 그룹 또는 조직의 모든 프로젝트에 대한 종속성 및 라이선스 정보를 볼 수 있습니다. **그룹** 또는 **조직** 메뉴의 **Dependencies** 옵션을 사용하십시오:

<figure><img src="../../../.gitbook/assets/Screenshot 2023-05-11 at 12.45.48.png" alt="조직을 위한 종속성 탭"><figcaption><p>조직을 위한 종속성 탭</p></figcaption></figure>

{% hint style="info" %}
프로젝트를 가져오거나 다시 테스트하면 10초 후에 **Dependencies** UI에 변경 사항이 반영됩니다.
{% endhint %}

종속성과 라이선스 모두 프로젝트 또는 다른 필터 기준으로 필터링할 수 있습니다:

<div align="left">

<figure><img src="../../../.gitbook/assets/Screenshot 2023-05-11 at 13.11.22.png" alt="프로젝트 및 필터 선택"><figcaption><p>프로젝트 및 필터 선택</p></figcaption></figure>

</div>

* **프로젝트** 드롭다운에서 특정 프로젝트를 선택합니다.
* **필터** 드롭다운에서 [심각도 수준](../../prioritize-issues-for-fixing/severity-levels.md) 또는 프로젝트 유형으로 필터링하려면 해당 상자를 확인합니다.

{% hint style="info" %}
Dockerfile 프로젝트 유형의 결과는 이미지 빌드 후 스캔 결과의 중복으로 이어질 수 있기 때문에 필터 기준에서 기본적으로 Dockerfile 프로젝트 유형은 제외됩니다. [AP](../../../snyk-api/reference/reporting-api-v1.md)I 호출 결과와 일치하려면 API 결과에서 Dockerfile을 제외하거나 필터의 프로젝트 유형 열에서 Dockfiles를 켜십시오.
{% endhint %}  