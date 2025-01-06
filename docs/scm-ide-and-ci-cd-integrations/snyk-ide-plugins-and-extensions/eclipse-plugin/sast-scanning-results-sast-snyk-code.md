# SAST scanning results (SAST, Snyk Code)

이클립스 플러그인 버전 2.0.0 이후, Snyk은 이클립스의 네이티브 플로우와 더 깊은 통합을 도입하고 있습니다. (인라인 하이라이트, 문제 통합, 툴팁에 문제에 대한 정보) 다음은 `js` 파일에서 발견된 심각한 보안 취약점에 대한 이러한 모든 내용을 보여줍니다:

1. 보안 취약점이 하이라이트되어 코드에서 심각한 보안 취약점이 있음을 나타냅니다. 툴팁에서 해당 취약점의 ID 및 문제 내용을 볼 수 있습니다.
2. **Problems** 뷰(화면 하단)와의 통합을 볼 수 있습니다. **Problems** 뷰를 사용하여 문제를 필터링하고 그룹화하는 경우 유용합니다. Snyk은 또한 문제가 있는 줄을 표시하며, 문제 뷰에서 문제를 클릭하면 해당 위치로 이동합니다.
3. 왼쪽에는 거터 아이콘을 볼 수 있으며, 오른쪽에는 파일맵 하이라이트(우선순위와 일치하는 색상)를 볼 수 있습니다.

{% hint style="info" %}
현재 툴팁 정보는 JavaEditor 및 GenericEditor에 제한되어 있으며, 후자는 Wild Web Developer와 같은 플러그인의 기본 편집기입니다.
{% endhint %}

<figure><img src="https://docs.snyk.io/~gitbook/image?url=https%3A%2F%2F2533899886-files.gitbook.io%2F%7E%2Ffiles%2Fv0%2Fb%2Fgitbook-x-prod.appspot.com%2Fo%2Fspaces%252F-MdwVZ6HOZriajCf5nXH%252Fuploads%252F6lEpnzBNfOzm9WgNvoSu%252Fimage.png%3Falt%3Dmedia%26token%3D9fc312a7-63e6-43f3-8a7b-18f7cccb05d6&#x26;width=768&#x26;dpr=1&#x26;quality=100&#x26;sign=ddabdbd0&#x26;sv=2" alt=""><figcaption><p>Snyk Code 결과가 표시된 Eclipse</p></figcaption></figure>
