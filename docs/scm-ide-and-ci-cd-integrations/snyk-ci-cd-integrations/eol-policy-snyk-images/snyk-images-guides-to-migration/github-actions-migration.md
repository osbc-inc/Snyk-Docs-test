# GitHub 작업 이관

## 영향을 받는 GitHub 작업을 전환하는 방법

영향을 받는 워크플로를 **삭제 예정이 아닌 최신 작업**을 사용하도록 업데이트하는 것을 권장합니다.

### A. 대안 지원 소프트웨어로 전환할 수 있습니다

#### Python-3.6/Python-3.7에 대한 다음 단계를 따르십시오

1. **작업 결정 및 찾기**:
   * 업데이트해야 할 워크플로에 어떤 작업이 필요한지 결정합니다. 이 경우 `python-3.6` 작업을 snyk 빌드 도구 체인에서 사용할 수 있는 새 작업인 [python-3.10](https://github.com/snyk/actions/tree/master/python-3.10)로 대체하려고 합니다.
2. **워크플로 파일 업데이트**:
   * 현재 작업이 정의된 워크플로 파일을 엽니다.
   * 현재 작업을 지정하는 섹션을 찾습니다. 예를 들어 `python:3.6`과 같은 섹션을 찾습니다.
   * 현재 작업을 새 작업으로 대체합니다.
3. **변경 내용 저장**: 새로운 작업 버전으로 업데이트된 워크플로 파일을 저장합니다.
4. **워크플로 테스트**: 새 작업이 예상대로 작동하는지 확인하도록 업데이트된 워크플로를 테스트합니다.

<mark style="color:red;">예시 - 이전:</mark>

```yaml
name: Example workflow for Python-3.6 using Snyk
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: Run Snyk to check for vulnerabilities
        uses: snyk/actions/python-3.6@master // <- python 3.6 사용
        env:
          SNYK_TOKEN: $
```

<mark style="color:green;">예시 - 이후:</mark>

```yaml
name: Example workflow for Python using Snyk
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: Run Snyk to check for vulnerabilities
        uses: snyk/actions/python-3.10@master // <- python 3.10 사용
        env:
          SNYK_TOKEN: $
```

#### scala/sbt에 대한 다음 단계를 따르십시오 <a href="#a.2-please-follow-these-steps-for-scala-sbt" id="a.2-please-follow-these-steps-for-scala-sbt"></a>

1. **작업 결정 및 찾기**:
   * 업데이트해야 할 워크플로에 어떤 작업이 필요한지 결정합니다. 이 경우 `scala` 작업을 snyk 빌드 도구 체인에서 사용할 수 있는 새 작업인 [https://github.com/snyk/actions/tree/master/sbt1.10.0-scala3.4.2](https://github.com/snyk/actions/tree/master/sbt1.10.0-scala3.4.2)로 대체하려고 합니다.
2. **워크플로 파일 업데이트**:
   * 현재 작업이 정의된 워크플로 파일을 엽니다.
   * 현재 작업을 지정하는 섹션을 찾습니다. 예를 들어 `scala`와 같은 섹션을 찾습니다.
   * 현재 작업을 `sbt1.10.0-scala3.4.2@master`와 같은 새 작업으로 대체합니다.
3. **변경 내용 저장**: 새로운 작업 버전으로 업데이트된 워크플로 파일을 저장합니다.
4. **워크플로 테스트**: 새 작업이 예상대로 작동하는지 확인하도록 업데이트된 워크플로를 테스트합니다.

<mark style="color:red;">예시 - 이전:</mark>

```yaml
name: Example workflow for Scala using Snyk
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: Run Snyk to check for vulnerabilities
        uses: snyk/actions/scala@master // <- 이전 scala 작업 사용
        env:
          SNYK_TOKEN: $
```

<mark style="color:green;">예시 - 이후:</mark>

```yaml
name: Example workflow for Scala using Snyk
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: Run Snyk to check for vulnerabilities
        uses: snyk/actions/sbt1.10.0-scala3.4.2@master // <- 새 scala 작업 사용
        env:
          SNYK_TOKEN: $
      
```



### B. 전용 사용자 정의 작업 <a href="#b.-you-can-roll-your-own-custom-actions" id="b.-you-can-roll-your-own-custom-actions"></a>

\
Snyk 고객은 Snyk에서 제공하는 미리 제작된 작업에서 이동하길 원하는 경우 자체 사용 사례에 맞는 사용자 정의 작업을 만들 수 있습니다. 이 접근 방식을 통해 워크플로에 사용되는 작업에 대해 더 많은 사용자 정의와 제어가 가능해집니다.

사용자가 직접 작업을 생성함으로써 고객은 이미지/작업이 벤더에서 지원을 중단할 때의 영향을 피할 수 있습니다.

#### B.1 [Snyk 설정 작업](https://github.com/snyk/actions/tree/master/setup) 활용 <a href="#b.1-leveraging-the-snyk-setup-action" id="b.1-leveraging-the-snyk-setup-action"></a>

이 [작업](https://docs.snyk.io/scm-ide-and-ci-cd-workflow-and-integrations/snyk-ci-cd-integrations/github-actions-for-snyk-setup-and-checking-for-vulnerabilities/snyk-setup-action)은 Snyk를 효과적으로 워크플로에 통합하는 다양한 방법을 제공합니다. 자세한 문서는 공식 문서에서 찾을 수 있습니다.

이 작업을 사용해야 하는 경우:

* 이미 개발 도구가 설치된 워크플로가 있는 경우
* 특정 환경을 위해 미리 정의된 Snyk 작업에 의존하고 싶지 않지만 워크플로에서 snyk-cli를 효율적으로 설정하고 싶은 경우
* 특정 환경용으로 구축된 작업을 찾을 수 없는 경우

#### B.2 직접 CLI 설치 <a href="#b.2-direct-cli-installation" id="b.2-direct-cli-installation"></a>

다른 옵션은 Snyk CLI를 직접 설치하고 GitHub Actions 워크플로에서 활용하는 것입니다. 이 방법을 사용하면 전용 GitHub Actions 통합 요구 사항을 건너뛸 수 있습니다.



```yaml
name: Example workflow using Snyk
on: push
jobs:
  snyk_scan:
    runs-on: ubuntu-latest
    steps:
    - name: Checkout code
      uses: actions/checkout@v2
    - name: Install Snyk CLI
      run: |
        curl https://downloads.snyk.io/cli/stable/snyk-linux -o snyk-linux
        curl https://downloads.snyk.io/cli/stable/snyk-linux.sha256 -o snyk.sha256
        sha256sum -c snyk.sha256
        chmod +x snyk-linux
        sudo mv snyk-linux /usr/local/bin/snyk
    - name: Run Snyk to test project dependencies
      env:
        SNYK_TOKEN: $
      run: |
        snyk test
```  