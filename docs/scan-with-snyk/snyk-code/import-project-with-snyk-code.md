# Snyk 코드로 프로젝트 가져오기

가져온 프로젝트는 프로젝트 페이지에서 대상 폴더 아래에 구성되며, Git 저장소 계정 및 특정 저장소를 기준으로 명명됩니다. [스캔 및 문제 식별을 위한 프로젝트 가져오기](../../getting-started/#import-a-project-to-scan-and-identify-issues)를 참조하세요.

**코드 분석 프로젝트** > **설정**으로 이동하여 테스트 빈도를 구성할 수 있습니다. [분석 테스트 빈도 설정](../../snyk-admin/snyk-projects/#test-frequency-settings)를 참조하세요.

저장소에 수동 테스트를 수행하려면 [지금 다시 테스트 옵션](manage-code-vulnerabilities/#retesting-code-repository)을 사용할 수 있습니다.

## Snyk에 저장소 다시 가져오기

이미 Snyk에 가져온 저장소를 [옵션이 활성화되기 전에](configure-snyk-code.md#enable-snyk-code-in-snyk-web-ui) 테스트하려면, 이를 다시 가져와야 합니다([ 조건](configure-snyk-code.md#conditions) 참조).

1. Snyk 웹 UI에 로그인하고 [그룹 및 조직](../../snyk-admin/groups-and-organizations/)을 선택합니다.
2. **프로젝트** > **프로젝트 추가**로 이동합니다.
3. **Snyk 코드 활성화** 섹션에서 설정을 **활성화**로 변경합니다.
4. 다시 가져올 저장소를 포함하는 Git 저장소를 선택합니다.
5. **개인 및 조직 저장소** 페이지에서 Snyk에 다시 가져올 저장소를 선택합니다. 이미 가져온 저장소는 체크 표시로 표시됩니다.
6. **선택한 저장소 추가**를 선택하여 선택한 저장소를 다시 가져옵니다.

선택한 저장소가 Snyk에 다시 가져와집니다. 다시 가져오기 프로세스 중에 는 이러한 저장소를 분석하고, 각 저장소에는 테스트 결과가 포함된 코드 분석 프로젝트가 추가됩니다.

<figure><img src="../../.gitbook/assets/Re-test repository.png" alt="저장소 다시 테스트"><figcaption><p>저장소 다시 테스트</p></figcaption></figure>

## 다음 단계

* [저장소 가져오기 작업 방식](../import-project-repository/#how-importing-repositories-works)
* [가져오기 프로세스에서 디렉터리 및 파일 제외](../import-project-repository/exclude-directories-and-files-from-project-import.md)
* [가져온 저장소 제거](../import-project-repository/remove-imported-repository-from-a-project.md)
