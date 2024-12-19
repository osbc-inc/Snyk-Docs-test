# 미구성 스캐닝 결과 (Snyk Infrastructure as Code)

Eclipse 플러그인 버전 2.0.0 이후, Snyk은 Eclipse의 네이티브 플로에 대한 심층 통합을 도입하고 있습니다: 인라인 강조, 문제 통합, 그리고 호버 시 문제에 대한 정보를 표시합니다. 다음은 Terraform 파일에서 발견된 심각한 중요도의 구성 오류에 대한 모든 정보를 보여줍니다:

1. 구성 오류가 강조 표시됩니다 (빨간색 squiggly line) 이 파일에 심각한 보안 취약점이 있음을 나타내고 그 행 번호가 표시됩니다. 호버하면 모든 정보를 확인할 수 있으며 링크를 클릭하면 더 많은 정보를 얻을 수 있습니다. 구성 오류를 해결하는 방법에 대한 조언은 구성 오류가 있는 위치에 바로 표시됩니다.
2. **Problems** 뷰와의 통합을 볼 수 있습니다. 이는 문제를 필터링하고 그룹화하는 데 **Problems** 뷰를 사용하는 경우 유용합니다. Snyk은 또한 문제가 있는 행을 표시하며 문제 보기에서 문제를 클릭하면 해당 위치로 이동합니다.
3. 왼쪽에 있는 gutter 아이콘과 오른쪽에 우선순위와 일치하는 색상의 파일 맵 하이라이트를 볼 수 있습니다.

{% hint style="info" %}
현재 호버 정보는 JavaEditor 및 GenericEditor로 제한되어 있습니다. 후자는 Wild Web Developer와 같은 플러그인의 기본 에디터입니다.
{% endhint %}

<figure><img src="../../../.gitbook/assets/image (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (2) (1).png" alt=""><figcaption><p>Snyk Infrastructure as Code가 Eclipse에 표시된 결과</p></figcaption></figure>