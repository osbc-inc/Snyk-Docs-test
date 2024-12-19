# CircleCI 마이그레이션

이 문서는 영향을 받는 작업에서 전환하는 방법에 대해 설명합니다.

## snyk orb 업데이트 및 `iac test` 사용

`scan-iac` 작업을 사용하는 고객은 `snyk/scan`과 `iac test` 명령을 사용하도록 전환해야 합니다. 예를 보려면 snyk-orb 저장소의 [IaC 스캔 예시](https://github.com/snyk/snyk-orb/blob/v2.0.0/src/examples/quickstart-iac-scanning.yml)를 참조하십시오.

circleci 구성 파일을 업데이트하여 사용 중인 snyk orb 버전을 최신 버전으로 업데이트하는 것이 중요합니다. 현재 최신 Snyk orb는 `snyk/snyk@2.1.0`입니다.

## Snyk orb를 사용하여 Snyk CLI만 설치하고 자체 단계에서 Snyk CLI 명령 실행<a href="#use-the-snyk-orb-to-only-install-the-snyk-cli-and-then-run-the-snyk-cli-commands-in-your-own-steps" id="use-the-snyk-orb-to-only-install-the-snyk-cli-and-then-run-the-snyk-cli-commands-in-your-own-steps"></a>

사전 정의된 작업에 의존하는 대신 고객은 Snyk orb를 사용하여 Snyk CLI를 설치한 후 명령을 자체 단계로 실행할 수 있습니다. snyk-orb 저장소의 [설치 예시](https://github.com/snyk/snyk-orb/blob/v2.0.0/src/examples/only-install.yml)를 참조하십시오.

`scan-iac` 작업을 대체하는 경우 예시 구성은 다음과 같을 수 있습니다:

```yaml
version: 2.1
orbs:
  node: circleci/node@5
  snyk: snyk/snyk@2.1.0
jobs:
  snyk_scan:
    docker:
      - image: cimg/node:lts
    steps:
      - checkout
      - run: npm ci
      - snyk/install
      - run:
          command: snyk iac test
          name: Run iac test 
```

## Snyk orb를 사용하지 않고 직접 CLI 설치<a href="#direct-cli-installation-without-using-the-snyk-orb" id="direct-cli-installation-without-using-the-snyk-orb"></a>

대신 Snyk CircleCI orb에 의존하길 원치 않거나 파이프라인을 완전히 제어하려는 고객은 Snyk CLI를 직접 설치할 수 있습니다. 다음은 예시입니다:

```yaml
version: 2.1
jobs:
  snyk_scan:
    docker:
      - image: cimg/node:lts
    steps:
      - checkout
      - run: npm ci
      - run:
          name: Snyk CLI 다운로드
          command: |
            curl -Lo snyk-linux https://downloads.snyk.io/cli/stable/snyk-linux
      - run:
          name: Snyk CLI SHA256 체크섬 다운로드
          command: |
            curl -Lo snyk-linux.sha256 https://downloads.snyk.io/cli/stable/snyk-linux.sha256
      - run:
          name: SHA256 체크섬 확인
          command: |
            sha256sum -c snyk-linux.sha256
      - run:
          name: Snyk CLI 설치
          command: |
            chmod +x snyk-linux
            ./snyk-linux --version
      - run:
          name: Snyk iac test 실행
          command: |
            ./snyk-linux iac test
workflows:
  version: 2
  build_and_scan:
    jobs:
      - snyk_scan
```