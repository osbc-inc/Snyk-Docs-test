# 파일 제외 및 문제 무시 FAQ

여러 가지 고려 사항이 있으며 몇 가지 요소에 따라 파일을 제외하고 문제를 무시하는 방법이 작동하는 경우가 있습니다:

* 프로젝트가 가져온 방법: SCM 통합을 통해 가져온 것이거나 CLI 또는 IDE를 통해 가져온 것인지
* 사용 중인 스캐닝 방법: 오픈 소스, 코드, 컨테이너, 현재 IaC, IaC+
* 테스트가 어떻게 수행되는지: UI, CLI 또는 IDE를 통해
* 제외하거나 무시한 방법: 정책, UI 또는 API를 통해 설정되었는지 또는 `.snyk` 파일에

이 문서는 지원팀이 자주 받는 질문을 모아서 답변을 제공합니다.

## 스캐닝 방법과 관련된 질문

### Q: 코드(SAST) 스캔에서 문제와 취약점을 무시하는 방법은 무엇인가요?

* 코드 취약점을 무시하려면 프로젝트를 Snyk UI로 가져와서 무시 버튼을 사용하세요.&#x20;
* 코드 스캔에서 문제를 무시하려면 `.snyk` 파일을 사용할 수 없습니다.
* 무시된 문제로 코드 스캔된 결과는 현재 `snyk-to-html` 도구에서 표시되지 않습니다. 이러한 문제는 여전히 문제로 표시됩니다.

### Q: 오픈 소스 스캔에서 특정 파일을 스캔하지 않게 하는 방법은 무엇인가요?

* CLI를 사용하여 스캔할 때 --`exclude` 옵션을 사용하여 특정 디렉터리나 파일을 스캔에서 제외합니다. 이 옵션은 모든 디렉터리 또는 지정한 이름의 모든 파일을 제외합니다. 자세한 내용은 CLI 테스트 명령어 도움말의 [--exclude 옵션](../../../snyk-cli/commands/test.md#exclude-less-than-name-greater-than-less-than-name-greater-than-...greater-than)을 참조하십시오.
* SCM 통합을 통해 프로젝트를 가져오는 경우, 가져오기 창의 맨 아래에 제외할 폴더를 추가하십시오. Git 리포지토리 배포 권장사항에서 [Stage 2: 프로젝트 가져오기](../../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/introduction-to-git-repository-integrations/deployment-recommendations-for-scm-integrations.md#stage-2-import-projects)을 참조하십시오.
* API 가져오기에서도 `snyk-api-import` 도구를 통해 제외용 글로브를 사용할 수 있습니다. 이 제외는 SCM 통합 제외와 동일하게 작동합니다. 이는 코드 및 컨테이너 스캔에만 작동합니다.
* 오픈 소스 스캔에서 `--exclude` 옵션을 `.snyk` 파일에 사용할 수 없습니다(제어되지 않는 스캔을 제외하고). 자세한 내용은 [글로브 표현식을 사용하여 파일 또는 폴더 제외 -  및 `unmanaged`only](../../../snyk-cli/commands/ignore.md#ignore-files-or-folders-using-glob-expression-snyk-code-and-unmanaged-only)을 참조하십시오.

### Q: 코드 스캔에서 특정 파일을 스캔하지 않게 하는 방법은 무엇인가요?

* 특정 파일이나 폴더의 모든 스캔을 제외하려면 `.snyk` 파일에서 `--exclude` 옵션을 사용하십시오. 자세한 내용은 [글로브 표현식을 사용하여 파일 또는 폴더 제외 -  및 `unmanaged`only](../../../snyk-cli/commands/ignore.md#ignore-files-or-folders-using-glob-expression-snyk-code-and-unmanaged-only)을 참조하십시오.
*  스캔에서 특정 파일이나 폴더의 스캔을 제외하려면 `snyk ignore --file-path` 명령을 사용하십시오. 자세한 내용은 [{Snyk Code CLI 테스트에서 디렉터리 및 파일 제외](../../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/exclude-directories-and-files-from-snyk-code-cli-tests.md)을 참조하십시오.
* 를 사용하여 리포지토리를 가져올 때, 가져오기 중에 `.snyk` 파일에 `exclude:` 문을 사용하여 가져오기에서 특정 디렉터리와 파일을 제외합니다. 자세한 내용은 [{Snyk Code CLI 테스트에서 디렉터리 및 파일 제외](../../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/exclude-directories-and-files-from-snyk-code-cli-tests.md)을 참조하십시오.
* 리포지토리나 SCM의 루트 디렉터리에 있는 파일이나 폴더 제외사항을 담은 `.snyk` 파일은 SCM을 사용하여 가져올 때 해당 파일과 폴더가 스캔에서 제외됩니다.
* CLI `--exclude` 옵션은 Code 스캔에 대해 작동하지 않습니다.
* 가져오기 창의 "폴더 제외" 옵션은 Code 스캔에 대해 작동하지 않습니다.
* IDE 스캔에서 코드의 파일 및 디렉터리 제외에는 `.snyk` 파일이 작동하지 않습니다.

### Q: 컨테이너 스캔에서 특정 파일을 스캔하지 않게 하는 방법은 무엇인가요?

`snyk container test` 도움말에서 [`--exclude-app-vulns`](../../../snyk-cli/commands/container-test.md#exclude-app-vulns)와 [`--exclude-base-image-vulns`](../../../snyk-cli/commands/container-test.md#exclude-base-image-vulns) 옵션을 참조하십시오.

### Q: IaC 스캔에서 특정 파일을 스캔하지 않게 하는 방법은 무엇인가요?

[명령줄을 사용한 IaC 제외사항](../../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-iac/iac-exclusions-using-the-command-line.md)을 참조하십시오.

## 제외 및 무시 설정 방법과 관련된 질문

### Q: CLI로 스캔할 때 루트 디렉터리의 snyk 파일에서 무시하지만 SCM을 사용하여 가져올 때 작동하지 않는 이유는 무엇인가요?

SCM 스캔에서 `.snyk` 파일은 각 관련 있는 하위 디렉터리에 있어야 합니다. `.snyk` 파일과 관련된 [모노리포 및 복잡한 프로젝트 고려 사항](../../policies/the-.snyk-file.md#monorepos-and-complex-project-considerations)을 참조하십시오.