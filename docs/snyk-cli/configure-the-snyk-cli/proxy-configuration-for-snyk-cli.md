# Snyk CLI를 위한 프록시 구성

## 일반 프록시 구성

Snyk CLI를 프록시 뒤에서 사용할 때 다음 환경 변수를 사용하여 프록시 구성을 제공해야 합니다:

`HTTP_PROXY` 또는 `http_proxy`

`HTTPS_PROXY` 또는 `https_proxy`

`NO_PROXY` 또는 `no_proxy`

`HTTPS_PROXY`의 `https`는 `https` 프로토콜을 사용하는 요청이이 프록시를 사용한다는 것을 의미합니다. 프록시 자체는 `https`를 사용할 필요가 없습니다.

자세한 정보는 [Snyk CLI를 Snyk API에 연결하도록 구성](configure-snyk-cli-to-connect-to-snyk-api.md) 및 [프록시 뒤에서 Snyk를 어떻게 사용할 수 있나요?](https://support.snyk.io/s/article/How-can-I-use-Snyk-behind-a-proxy)을(를) 참조하십시오.

## 프록시 인증

기본적으로 Snyk CLI는 프록시 인증을 감지하고 적용하려고 시도합니다.

프록시 서버가 프록시 인증을 요청하고(`PROXY-AUTHENTICATE` 응답 헤더에 표시됨) 서버와 CLI가 동일한 인증 기구를 지원하는 경우, CLI는 현재 운영 체제에 로그인된 사용자로 인증(단일 로그인)합니다.

다음 인증 기구가 지원됩니다:

* 협상
  * 전체 OS에서 Kerberos
  * NTLM(Windows NT LAN Manager)

## Windows 운영 체제에서의 구성

Windows 운영 체제(OS)에서는 Kerberos 및 NTLM 인증 기구가 OS 자체에서 제공되며 모든 도메인 사용자에게 사용 가능합니다.

Snyk CLI는 특정 구성을 필요로 하지 않습니다.

## Windows 이외의 운영 체제(Linux, macOS)에서의 구성

Windows 이외의 운영 체제에서도 Snyk CLI는 SSO를 지원하지만, 추가로 다음 환경 변수로 구성되어야 합니다.

```bash
KRB5_CONFIG # 기본값 "/etc/krb5.conf"
KRB5CCNAME # 기본값 "FILE:/tmp/krb5cc_<사용자UID>"
```

이러한 변수의 사용은 MIT Kerberos 구현을 따릅니다:

* [KRB5\_CONFIG](https://web.mit.edu/kerberos/krb5-devel/doc/admin/conf\_files/krb5\_conf.html) - Kerberos 구성 파일
* [KRB5CCNAME](https://web.mit.edu/kerberos/krb5-1.12/doc/basic/ccache\_def.html) - Kerberos 자격 캐시
   * 현재 지원되는 자격 캐시 형식(`ccache` 형식)은 **FILE**뿐입니다.
   * 중요한 점은 CLI로 캐시 파일을 업데이트할 수 없다는 것입니다. 이는 캐시 파일이 외부적으로 업데이트되어야 한다는 것을 의미합니다. 예를 들어, [kinit](https://web.mit.edu/kerberos/krb5-1.12/doc/user/%EC%82%AC%EC%9A%A9%EC%9E%90/%EC%82%AC%EC%9A%A9%EC%9E%90_%EC%BB%AE%EB%A6%AC%EB%94%96/kinit.html)를 실행함으로써 업데이트될 수 있습니다.

## 프록시 인증 해제

인증을 비활성화하려면 다음 명령줄 매개변수를 지정하십시오:

```jsx
--proxy-noauth
```

## 문제 해결

연결 문제가 발생하면 유용한 통찰력을 얻으려면 디버그 출력 `-d`를 활성화하십시오.