# Git 리포지토리 및 JavaScript

Snyk가 지원하는 Git 서비스를 통해 JavaScript 저장소를 가져올 수 있습니다. 가져온 후에 Snyk는 지원하는 manifest 파일을 기반으로 프로젝트를 분석합니다.

## JavaScript 언어 설정을 위한 Snyk

오픈 소스 및 라이선스에 대한 언어 설정은 조직 수준에서 구성할 수 있습니다. 구성 설정은 해당 조직의 모든 프로젝트에 적용됩니다.

1. Snyk 계정에서 **설정 >** **언어**로 이동합니다.
2. **언어** 아래에서 **JavaScript**을 찾아 **설정 편집**을 선택합니다.
3. 패키지 관리자, **npm** 또는 **Yarn**에 기반하여 설정을 구성합니다.
   * **Dev dependencies 스캔 및 수정**: 이 옵션을 선택하면 Snyk이 `package.json`의 `devDependencies` 속성을 읽고 해당 취약점을 보고하고 수정합니다.
   * **package.json 및 package-lock.json/yarn.lock 파일 동기화 필요**: 이 옵션을 선택하면 `package.json`과 `package-lock.json`/`yarn.lock` 파일이 동기화되지 않은 경우 Snyk가 가져오기를 실패합니다.
   * **취약점 수정 시 package-lock.json 생성 제외**: 개인 미러 또는 레지스트리를 사용하는 경우, Snyk이락 파일을 업데이트하기 위해 npm 레지스트리를 사용하기 때문에 Snyk에서 생성된 락 파일이 적합하지 않을 수 있습니다. 엔터프라이즈 고객은 [패키지 저장소 통합](../../scan-with-snyk/snyk-open-source/package-repository-integrations/)을 사용하여 락 파일이 올바르게 업데이트되도록 할 수 있습니다. 또는 이 설정을 통해 Snyk 수정 풀 리퀘스트 및 머지 리퀘스트에 대한 락 파일 생성을 건너뛸 수 있습니다.
4. 변경 사항을 저장하기 위해 **설정 업데이트**를 클릭합니다.

## JavaScript에서의 Snyk 워크스페이스

{% hint style="info" %}
Snyk Git 저장소 통합 스캔에서 Yarn 및 npm 워크스페이스는 명시적으로 지원되지 않습니다.
{% endhint %}

인접한 락 파일이 있는 루트 수준의 `package.json` manifest 파일은 일반적으로 스캔됩니다.

락 파일이 없는 중첩된 manifest 파일의 경우, Snyk는 루트 락 파일을 사용하지 않고 빌드 시점의 종속성 트리가 어떻게 보일지 대략적으로 계산합니다.

또한, 중첩된 manifest 파일에 대한 수정 PR은 지원되지 않습니다.

## 수정 PR 및 npm save-prefix

npm v7+ 프로젝트를 사용하여 취약점을 수정하는 경우, Snyk은 프로젝트에서 유추하는 대신 기본 npm `save-prefix`를 사용합니다.

즉, caret 범위(`^`) 이외의 범위 형식을 사용하는 종속성이 있는 경우, `package-lock.json` 파일의 `version` 필드에 추가적인 변경 사항이 있을 수 있습니다.

이 변경 사항은 일상적인 기능에 영향을 주지 않으며, 범위는 `package.json`에서 읽힐 것입니다.

## Yarn zero-installs 사용자를 위한 수정 PR

Yarn v2에서 [zero-installs](https://yarnpkg.com/features/zero-installs) 기능이 출시되었는데, 이는 Yarn 개발자가 로컬 환경에서 종속성을 설치하지 않고도 프로젝트에 작업할 수 있게 해주었습니다.

Zero-installs는 프로젝트의 모든 종속성을 `.yarn/cache` 디렉토리 내에 설치하고, 다음 개발자가 새로운 종속성을 직접 리포지토리에서 가져올 수 있도록 이를 버전 관리 시스템에 커밋하도록 합니다.

{% hint style="info" %}
**zero-installs** 기능을 사용하는 경우, Snyk 수정 PR은 **.yarn/cache** 디렉토리를 업데이트하지 않습니다. 이 디렉토리를 업데이트하려면 `yarn`을 실행해야 합니다.
{% endhint %}
