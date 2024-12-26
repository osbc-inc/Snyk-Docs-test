# Snyk 브로커 -  탐지

## **Snyk 브로커 내  (IaC)**

기본적으로, 인프라스트럭처 코드 (IaC)에서 사용되는 일부 파일 유형은 활성화되어 있지 않습니다. 브로커가 귀하의 저장소에 있는 Terraform과 같은 IaC 파일에 액세스하도록 허용하려면 `ACCEPT_IAC` 환경 변수를 추가할 수 있습니다. 이 변수에는 `tf, yaml, yml, json, tpl` 중 임의의 조합을 사용합니다.

예시:

```
docker run --restart=always \
           -p 8000:8000 \
           -e BROKER_TOKEN=secret-broker-token \
           -e GITHUB_TOKEN=secret-github-token \
           -e PORT=8000 \
           -e BROKER_CLIENT_URL=http://my.broker.client:8000 \
           -e ACCEPT_IAC=tf,yaml,yml,json,tpl
       snyk/broker:github-com
```

그렇지 않을 경우, `accept.json`을 편집하고 관련 IaC 특정 규칙을 추가하여 사용자 정의 accept 파일을 컨테이너에 로드할 수 있습니다. 사용자 정의 accept 파일(다른 폴더에서 사용하는 경우)이 사용된다면(`ACCEPT` 환경 변수를 사용), `ACCEPT_IAC` 메커니즘을 사용할 수 없다는 점에 유의하십시오.

## `ACCEPT`를 통한 사용자 정의 구성

`ACCEPT` 환경 변수를 사용하여 Snyk 브로커를 사용하여  파일을 감지하는 방법에 대한 지침은 다음 페이지를 참조하십시오:

* [브로커를 사용하여 Terraform 구성 파일 감지](detecting-terraform-configuration-files-using-snyk-broker-custom.md)
* [브로커를 사용하여 CloudFormation 구성 파일 감지](detecting-cloudformation-configuration-files-using-snyk-broker-custom.md)
* [브로커를 사용하여 Kubernetes 파일 감지](detecting-kubernetes-configuration-files-using-a-broker.md)