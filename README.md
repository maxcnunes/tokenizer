# tokenizer

Generates oauth2 tokens

Working with micro services and oauth2 authentications can cause such hard work generating new access tokens for different clients and generating new tokens when the previous one has expired. Due that I created this tool to help generating new tokens.

## Install

### From binary

[Releases](https://github.com/maxcnunes/tokenizer/releases)

### With Go

```bash
go get -u github.com/maxcnunes/tokenizer
```

## Configuration

Configuration file **.tokenizer.json** must be located in HOME directory:

```json
{
  "OAuth2Services": [
    {
      "name": "client1",
      "url": "http://oauth2-service.example/authentication-route",
      "grantType": "client_credentials",
      "clientId": "1",
      "clientSecret": "xxxxxxxxxxxx"
    },
    {
      "name": "client2",
      "url": "http://oauth2-service.example/authentication-route",
      "grantType": "client_credentials",
      "clientId": "2",
      "clientSecret": "xxxxxxxxxxxx"
    }
  ]
}
```

## Usage

```bash
$GOPATH/bin/tokenizer
```

## Example

### Select options

```bash
» $GOPATH/bin/tokenizer
Select the oauth2 service:
    > client1
    > client2
>>> client1
{
    "access_token": "947115cd41ae4e9597a4f1d93b9415d5",
    "expires_in": 3600,
    "refresh_token": "x7d00ce9b2874c11903da50cf6cabdf1",
    "token_type": "bearer"
}%
```

### Single command

```bash
» $GOPATH/bin/tokenizer -name client1
{
    "access_token": "947115cd41ae4e9597a4f1d93b9415d5",
    "expires_in": 3600,
    "refresh_token": "x7d00ce9b2874c11903da50cf6cabdf1",
    "token_type": "bearer"
}%
```

## Build

Using [goxc](https://github.com/laher/goxc).

```bash
goxc
```
