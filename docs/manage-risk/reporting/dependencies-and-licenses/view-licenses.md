# 라이선스 보기

**라이선스** 탭은 프로젝트에서 검색된 모든 라이선스를 표시하며, 프로젝트 내의 모든 종속 항목 및 해당 라이선스를 사용하는 모든 프로젝트의 개요가 제공됩니다. 이를 통해 라이선스를 가진 프로젝트 및 종속 항목을 확인할 수 있습니다. 다음은 예시입니다:

<div align="left">

<figure><img src="../../../.gitbook/assets/Screenshot 2023-05-11 at 15.38.44.png" alt="라이선스 탭"><figcaption><p>라이선스 탭</p></figcaption></figure>

</div>

## **필드 세부정보**

<table><thead><tr><th width="162">필드</th><th>설명</th></tr></thead><tbody><tr><td><strong>심각도</strong></td><td>평가된 심각도 수준입니다. 라이선스 정책에서 심각도 수준을 설정할 수 있습니다. 라이선스 정책은 <a href="../../../manage-risk/policies/license-policies/create-a-license-policy-and-rules.md">라이선스 정책 및 규칙 생성</a>을 참조하십시오. 기본 Snyk 라이선스 정책은 일부 라이선스에 대해 심각도를 정의합니다.</td></tr><tr><td><strong>라이선스</strong></td><td>라이선스의 전체 이름입니다. SPDX 라이선스는 <a href="https://spdx.org">SPDX</a> 사이트로 연결되며, 비-SPDX 라이선스는 해당 사이트로 연결됩니다.<br>라이선스가 <strong>Unknown</strong>으로 표시된 경우, Snyk가이 패키지에 대한 라이선스를 찾을 수 없었습니다.</td></tr><tr><td><strong>종속 항목</strong></td><td>프로젝트 내 해당 라이선스를 가진 종속 패키지의 총 수이며, 이는 측면 패널에 연결되어 동일한 레이아웃으로 영향을 받는 종속 항목의 전체 목록을 표시합니다.</td></tr><tr><td><strong>프로젝트</strong></td><td>이 라이선스를 가진 프로젝트의 수이며, 이는 측면 패널에 연결되어 동일한 레이아웃으로 영향을 받는 프로젝트의 전체 목록을 표시합니다.</td></tr></tbody></table>

## **라이선스 탭 작업**

작업은 탭 상단에 나타납니다.

<div align="left">

<figure><img src="../../../.gitbook/assets/Screenshot 2023-05-11 at 15.50.22 (1) (1).png" alt="라이선스 탭 작업"><figcaption><p>라이선스 탭 작업</p></figcaption></figure>

</div>

* **라이선스 검색**: 자유 텍스트를 입력하고 입력하는 첫 문자로 검색을 시작합니다. 필드를 클릭하면 열리는 드롭다운 목록에서 여러 패키지를 선택할 수도 있습니다. 또한 드롭다운 목록에서 동적으로 나타나는 **모두 선택** 또는 **모두 선택 해제** 링크를 클릭할 수도 있습니다.
* **CSV로 내보내기**: 문제 데이터를 CSV 파일 형식으로 내보냅니다.