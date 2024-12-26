# Broker 이미지 서명 검증

버전 4.169.1부터 모든 Broker 컨테이너 이미지는 Cosign을 사용하여 서명됩니다.

Cosign은 컨테이너를 서명하고 검증하는 도구입니다. Cosign은 Open Container Initiative (OCI) 레지스트리에 저장을 포함합니다.

Cosign은 컨테이너 이미지의 보안을 강화하기 위해 서명 및 검증을 간단하고 효율적인 방법으로 제공합니다. 이 도구는 디지털 서명의 개념을 활용하여 컨테이너 이미지를 개인 키로 서명하고 수령 측이 해당 공개 키를 사용하여 서명을 확인할 수 있도록 합니다.

## Broker 이미지 서명을 검증하기 위한 전제 조건

[Cosign을 설치](https://docs.sigstore.dev/system\_config/installation/)해야 합니다. 버전 2.2.0 이상이어야 합니다.

## 서명된 컨테이너 이미지 검증

서명된 이미지를 확인하려면 내장된 `cosign verify` 명령을 사용해야 합니다.

검증 단계에서 Broker 컨테이너 이미지를 가져올 필요가 없습니다.

```
$ cosign verify --key cosign.pub snyk/broker:4.169.1-github-com

index.docker.io/snyk/broker:4.169.1-github-com에 대한 검증 --
다음의 체크가 각 서명에 대해 수행되었습니다:
  - Cosign 클레임이 유효성이 검사됨
  - 투명성 로그에 대한 체크의 존재가 오프라인에서 확인됨
  - 서명이 지정된 공개 키에 대해 검증됨

[{"critical":{"identity":{"docker-reference":"index.docker.io/snyk/broker"},"image":{"docker-manifest-digest":"sha256:a2ff856b180a532c3e31a90b9788cad567fa05d78c84bccc637de54c6f46ebf2"},"type":"cosign container image signature"},"optional":{"Bundle":{"SignedEntryTimestamp":"MEUCIEOmElvKK0eC/hvpM9SE66RAekcV6DpF6NSO4Gz5aftrAiEAlQY8lKe1RUqYtCK1WRwWhWT/f/PvyBTJC8qBjgU20kU=","Payload":{"body":"eyJhcGlWZXJzaW9uIjoiMC4wLjEiLCJraW5kIjoiaGFzaGVkcmVrb3JkIiwic3BlYyI6eyJkYXRhIjp7Imhhc2giOnsiYWxnb3JpdGhtIjoic2hhMjU2IiwidmFsdWUiOiJjOTE2NDQ1N2YxMDA0NTQxNWNlMjBlN2I3YjNmYjg4YjZmMmNhNzI4MDNkODY4NTk0ZDhlY2UzMGJkYTFiZjQ4In19LCJzaWduYXR1cmUiOnsiY29udGVudCI6Ik1FWUNJUUNRQVp6VVdqbkNFai9GZkpxTGU4YVdoYXhacWJzZnZTc21JNXRiRzZuRmdBSWhBTFgyMWVLRHl6OXNqWHQrVStVZUZNTUFyN1oyV09Gd0k1b1oxclc0dlJBUCIsInB1YmxpY0tleSI6eyJjb250ZW50IjoiTFMwdExTMUNSVWRKVGlCUVZVSk1TVU1nUzBWWkxTMHRMUzBLVFVacmQwVjNXVWhMYjFwSmVtb3dRMEZSV1VsTGIxcEplbW93UkVGUlkwUlJaMEZGWkNzeWJVVlhlVVJyT0VOdGJUQkRSREZhT0dwamMxaEhZVkV5YVFwelREaHdlRWh5ZDI5SlNEUkVlRzFrZVVveWJucDNWMkY0V1daelpscE5OazV2UTFKV2MyZFpRVlpsTlVkQ2FFWmljalpvZW1OcU5XZDNQVDBLTFMwdExTMUZUa1FnVUZWQ1RFbERJRXRGV1MwdExTMHRDZz09In19fX0=","integratedTime":1698788303,"logIndex":46732614,"logID":"c0d23d6ad406973f9559f3ba2d1ca01f84147d8ffc5b8445c224f98b9591801d","tag":"4.169.1-github-com"]
```

## Broker Cosign 공개 키

```
-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEd+2mEWyDk8Cmm0CD1Z8jcsXGaQ2i
sL8pxHrwoIH4DxmdyJ2nzwWaxYfsfZM6NoCRVsgYAVe5GBhFbr6hzcj5gw==
-----END PUBLIC KEY-----
```