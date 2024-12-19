# Universal Broker

{% hint style="info" %}
**릴리스 상태**  
유니버설 브로커는 초기 액세스 중이며 엔터프라이즈 플랜에서만 사용할 수 있습니다. 설정하려면 Snyk 계정팀에 문의하십시오.
{% endhint %}

유니버설 브로커는 Snyk 브로커 배포 및 연결 관리를 용이하게 하는 개선 사항을 제공합니다.

이 페이지에서는 클라이언트를 위한 구성 흐름을 설명합니다. 다음 페이지에서는 유니버설 브로커를 사용하여 브로커 배포 설정을 안내합니다:

- [유니버설 브로커의 초기 구성](initial-configuration-of-the-universal-broker.md)
- [API를 사용하여 GitHub 연결 설정](set-up-a-github-connection-using-the-api.md)
- [필수 환경 변수로 브로커 다시 시작 및 연결](restart-your-broker-with-the-required-environment-variable-and-connect.md)

## 유니버설 브로커 배포 및 연결 모델 <a href="#universal-broker-deployment-and-connection-model" id="universal-broker-deployment-and-connection-model"></a>

유니버설 브로커는 배포와 컨테이너 관련 사항을 연결 관련 사항과 분리합니다. 이를 통해 작은 배포나 한 가지 배포가 다양한 유형의 연결을 지원할 수 있습니다.

<figure><img src="../../../.gitbook/assets/image 5.png" alt="Universal Broker Model"><figcaption><p>유니버설 브로커 모델</p></figcaption></figure>

각 배포와 컨테이너가 하나의 연결 유형만 지원하는 기존 브로커 모델과 대조적으로, Snyk는 많은 종류의 연결을 지원하는 브로커 배포를 설정하고 잊는 방식으로 제공합니다.

단일 배포는 여러 유형의 연결을 지원할 수 있습니다. 새로운 기능을 제공할 수 있지만 Snyk은 컨테이너 리소스가 무한정 수직 확장 불가능하기 때문에 25개 연결 미만을 유지하는 것을 권장합니다.

이 모델은 자격 증명을 저장하지 않으며 자격 증명 참조 또는 참조를 사용합니다. 이는 특정 이름이 지정된 환경 변수에 찾아야 하는 연결을 지원하는 자격 증명을 브로커 클라이언트에 알려줍니다.

예를 들어, 배포는 연결(`github` 유형)을 실행하고 연관된 자격 증명 참조(`$MY_CRED_REF`)에서 정의된 환경 변수 이름을 사용합니다.

## 유니버설 브로커 구현 단계

{% hint style="warning" %}
사전 요구 사항: 배포, 자격 증명 참조 및 연결을 생성하려면 테넌트 관리자여야 합니다.
{% endhint %}

유니버설 브로커를 구현하는 데 필요한 고수준 단계는 다음과 같습니다.

1. **일회성:** Snyk 브로커 앱을 조직에 설치합니다. 이렇게 하면 플랫폼과 상호 작용하기 위해 설치 ID, 클라이언트 ID 및 클라이언트 비밀번호가 반환됩니다. 배포를 생성하려면 조직 ID가 필요합니다.
2. **일회성:** 테넌트 ID 및 설치 ID에 대한 배포를 정의합니다.
3. **일회성:** 연결에 필요한 자격 증명 참조를 정의합니다.
4. **일회성:** 원하는 연결 또는 연결을 정의합니다.
5. **각 조직 통합에 대해 한 번:** 해당 연결을 사용해야 하는 조직과 통합을 구성합니다.

<figure><img src="../../../.gitbook/assets/image 6.png" alt="Deployment, connections, and organizations"><figcaption><p>배포, 연결 및 조직</p></figcaption></figure>

원하는 연결을 정의한 후 브로커 클라이언트를 실행합니다. 연결과 조직 및 통합 사이의 통합은 유니버설 브로커 배포와 독립적으로 수행되어 실행 중인 컨테이너나 Kubernetes 배포에서 필요한 활동을 줄입니다.

`$snyk-broker-config` 명령을 사용하거나 관련 API 엔드포인트에 대한 HTTP 요청을 보내면 조직과의 통합을 통합하거나 연결을 해제할 수 있습니다. 브로커 클라이언트는 브로커 관리자가 이 모델에 따라 원하는 상태로 선언한 연결을 관리하며 배포 내에서 변경 사항을 반영하는 데 최대 2분이 걸릴 수 있습니다.