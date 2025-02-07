# 분석 결과: Snyk 코드

분석은 매 스캔 시 코드 내의 보안 취약점과 품질 문제를 보여줍니다.

{% hint style="info" %}
2025년 6월 24일부터, 품질 문제는 더 이상 제공되지 않습니다.
{% endhint %}

Visual Studio Code 결과 화면의 **문제** 탭에서 프로젝트 내에서 발견된 모든 코드 문제를 확인할 수 있습니다.

<figure><img src="../../../../.gitbook/assets/Screenshot 2023-03-17 at 13.41.55.png" alt="Visual Studio Code Problems tab"><figcaption><p>Visual Studio Code Problems tab</p></figcaption></figure>

## 편집기 창

편집기 내에서 취약점이 보이며, hover를 통해 자세한 정보를 확인할 수 있습니다.

<figure><img src="../../../../.gitbook/assets/Screenshot 2023-03-17 at 12.31.45.png" alt="편집기 창"><figcaption><p>편집기 창</p></figcaption></figure>

코드 액션을 통해 문제의 세부 정보 패널을 열기 위해 **Quick Fix**를 선택합니다.

현재 파일에서 특정 또는 반복되는 제안을 무시하려면 **Quick Fix**를 통해 선택할 수도 있습니다.

<figure><img src="../../../../.gitbook/assets/Screenshot 2023-03-17 at 16.34.21.png" alt="Quick Fix 메뉴"><figcaption><p>Quick Fix 메뉴</p></figcaption></figure>

<figure><img src="../../../../.gitbook/assets/Screenshot 2023-03-17 at 12.32.22.png" alt="문제 상세 내용을 무시하는 옵션"><figcaption><p>문제 상세 내용을 무시하는 옵션</p></figcaption></figure>

## 취약점 창

<figure><img src="../../../../.gitbook/assets/Screenshot 2023-03-17 at 12.25.28 (1).png" alt="Snuk 코드 취약점 창"><figcaption><p>Snuk 코드 취약점 창</p></figcaption></figure>

결과 화면 우측에 있는 Snyk 제안 패널에는 Snyk 코드 취약점의 이름, 발견된 라인, 수정 제안 및 무시 옵션이 표시됩니다.

Snyk은 잘못된 양성 결과를 이유로 보고하기 위한 피드백 메카니즘도 포함하고 있습니다(왼쪽 하단).
