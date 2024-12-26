# CLI를 사용하여 로 소스 코드 스캔

를 사용하면 Snyk CLI를 통해 저장소의 소스 코드를 테스트할 수 있습니다.

CLI를 통해 저장소 코드를 테스트할 때 다음과 같은 작업을 수행할 수 있습니다:

* [루트 폴더에서 저장소 직접 테스트하기](scan-source-code-with-snyk-code-using-the-cli.md#testing-a-repository-from-its-root-folder).
* [다른 위치에서 저장소 테스트하기](scan-source-code-with-snyk-code-using-the-cli.md#testing-a-repository-from-a-different-location).

폴더를 테스트할 때 해당 폴더의 하위 폴더와 파일도 모두 테스트됩니다. 현재 폴더의 단일 파일 또는 다른 폴더의 단일 파일을 테스트하려면 파일의 절대 경로를 지정하세요.

상대 경로 참조를 사용하여 경로 앞에 `$PWD`를 접두어로 붙인 파일도 테스트할 수 있습니다. 예를 들어, `snyk code test $PWD/path/to/file`와 같이 사용합니다. 이는 bash와 함께 작동합니다.

 CLI 테스트에서 특정 디렉토리 또는 파일을 제외하려면 다음 방법을 사용할 수 있습니다:

* `snyk ignore --file-path` 명령어. [디렉토리 및 파일을  테스트에서 제외하기](exclude-directories-and-files-from-snyk-code-cli-tests.md)를 참조하세요.
* 테스트된 폴더에 `.snyk` 파일을 수동으로 만드는 것. [프로젝트 가져오기 과정에서 디렉토리 및 파일을 제외하기](../../../scan-with-snyk/import-project-repository/exclude-directories-and-files-from-project-import.md)를 참조하세요.

## **루트 폴더에서 저장소 테스트**

현재 저장소 폴더를 테스트하려면 터미널에 다음을 입력하세요:

```
snyk code test
```

루트 폴더에서 저장소를 테스트하기 위해 `snyk code test` 명령에 추가 옵션이 필요하지 않습니다.

는 현재 폴더를 테스트하고 터미널에 [테스트 결과](view-snyk-code-cli-results.md)를 표시합니다.

예를 들어, `snyk-goof` 저장소를 루트 폴더에서 테스트하려면 먼저 해당 저장소의 루트 폴더로 디렉토리를 변경한 다음 다음을 입력하세요:

```
snyk code test
```

가 `snyk-goof` 저장소를 테스트하고 발견된 취약점 문제를 표시합니다:

<figure><img src="../../../.gitbook/assets/ - CLI - snyk code test - Results - 1 (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (2) (5).png" alt=" CLI 테스트 결과 예제"><figcaption><p> CLI 테스트 결과 예제</p></figcaption></figure>

## **다른 위치에서 저장소 테스트**

다른 폴더에서 저장소를 테스트하려면 터미널에 다음을 입력하세요:

```
snyk code test <path/to/folder>
```

`path/to/folder`는 CLI를 통해 Snyk Code를 사용하여 테스트하려는 저장소의 전체 경로입니다.

예를 들어, 다른 디렉토리에서 **snyk-goof** 저장소를 테스트하려면 다음과 같이 입력하세요:

```
snyk code test /Users/username/Documents/Repositories/snyk-goof
```

<figure><img src="../../../.gitbook/assets/snyk Code - CLI - snyk code test - Any folder - 2 (1).png" alt=" CLI 테스트 결과 예제"><figcaption><p> CLI 테스트 결과 예제</p></figcaption></figure>

* 테스트 결과를 살펴보려면 [Snyk Code CLI 결과보기](view-snyk-code-cli-results.md)를 참조하세요.
* 테스트 결과를 활용하려면 [Snyk-to-HTML 기능을 사용하여 CLI 결과를 HTML 형식으로 표시](../cli-tools/snyk-to-html.md)를 참조하세요.