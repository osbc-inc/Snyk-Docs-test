# Snyk 설정 작업

[Snyk Setup Action](https://github.com/snyk/actions/tree/master/setup)은 취약점을 확인하기 위해 Snyk CLI를 설치하는 방법을 제공합니다. 이 작업을 사용해야 하는 시점에 대한 자세한 정보는 GitHub Actions 통합 페이지의 [자체 개발 환경 사용](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration#use-your-own-development-environment)을 참조하십시오.

Snyk Setup Action을 사용하여 취약점을 확인하기 위해 [Snyk](https://snyk.co/SnykGH)을 다음과 같이 설치할 수 있습니다:

```yaml
name: Snyk 예시 
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@master
    - uses: snyk/actions/setup@master
    - uses: actions/setup-go@v1
      with:
        go-version: "1.13"
    - name: Snyk 테스트
      run: snyk test
      env:
        SNYK_TOKEN: $
```

설정 액션을 사용할 때 Snyk을 실행하는데 필요한 개발 환경을 설정해야 합니다. 이 경우 Go 프로젝트이므로 `actions/setup-go`가 사용되었지만 프로젝트에 따라 다를 수 있습니다. [GitHub 언어 및 프레임워크 가이드](https://docs.github.com/en/actions/language-and-framework-guides)가 좋은 시작 점이 될 것입니다.

Snyk Setup Action은 `with`를 사용하여 하부 이미지로 전달되는 속성을 가지고 있습니다.

| 속성           | 기본값    | 설명                 |
| ------------ | ------ | ------------------ |
| snyk-version | latest | 특정 버전의 Snyk를 설치합니다 |

이 작업에는 출력도 있습니다:

| 속성      | 기본값 | 설명                  |
| ------- | --- | ------------------- |
| version |     | 설치된 Snyk CLI의 전체 버전 |

예를 들어, 특정 버전의 Snyk을 설치하고 출력에서 설치된 버전을 가져올 수 있습니다:

```yaml
name: Snyk 예시
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@master
    - uses: snyk/actions/setup@master
      id: snyk
      with:
        snyk-version: v1.391.0
    - uses: actions/setup-go@v1
      with:
        go-version: "1.13"
    - name: Snyk 버전
      run: echo "$"
    - name: Snyk 모니터링 
      run: snyk monitor
      env:
        SNYK_TOKEN: $
```
