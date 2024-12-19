# 고정 GitHub 계정에서 고침 및 업그레이드 프로젝트 요청 열기

GitHub 사용자 계정을 무작위로 선택하는 대신에 특정 GitHub 계정을 설정하여 PR(Pull Request)을 열고 수정 및 업그레이드할 수 있습니다.

{% hint style="info" %}
UI에서 수행되지 않는 다른 작업(예: 매일 또는 매주 테스트)은 여전히 GItHub 계정을 Snyk에 연결한 Snyk 조직 구성원이 무작위로 선택하여 수행합니다.
{% endhint %}

이 기능을 사용하려면:

1. Snyk 웹 UI에서 **Settings, then** **Integrations, Source control,** 그리고 **GitHub**로 이동하여 **Edit Settings**를 클릭합니다.
2. **Open fix and upgrade pull requests from a fixed GitHub account** 설정 아래의 토글 버튼을 활성화합니다.
3. GitHub에서 개인 액세스 토큰을 생성하는 방법에 대한 페이지 안내에 따릅니다.
4. 새로 생성된 토큰을 Snyk에 제공하여 GitHub에서 Fix PR을 열 수 있도록 합니다.

{% hint style="info" %}
토큰이 제공된 GitHub 계정이 Snyk와 모니터링하려는 리포지토리에 대해 **쓰기** 권한 이상을 갖고 있는지 확인하세요.
{% endhint %}

GitHub에서 리포지토리 권한 수준에 대한 자세한 정보는 [GitHub integration](../../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/github.md)을 참조하세요.