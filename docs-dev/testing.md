# Testing

In the `realm-export.json` are already client applications included to test the API with

-   [API Specification](#testing-with-api-specification) or
-   [Swagger Editor](#testing-with-swagger-editor) or
-   [Postman](#testing-with-postman).

## Request Token JWT Generation

To create a sufficient Token Request JWT, you can go to [JWT.io (external URL)](https://jwt.io) and create one.
You can use the ES256 private key from the [key examples (external repository)](https://github.com/JonasPrimbs/draft-ietf-mla-oidc/tree/main/examples/key-examples.md#private-key) to sign your key or use the [code example (external repository)](https://github.com/JonasPrimbs/draft-ietf-mla-oidc/tree/main/examples/code-examples.md#key-pair-generation-and-export) to generate a new one.
You can also use the Token Request JWT from the [communication examples (external repository)](https://github.com/JonasPrimbs/draft-ietf-mla-oidc/tree/main/examples/communication-examples.md#token-request) as template.

The following JWT can be used as a template:

```jwt
eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCIsImp3ayI6eyJrdHkiOiJFQyIsImNydiI6IlAtMjU2IiwieCI6ImNYUThiZGVOZWVTd2ZMa0h6TWZBVUZySGxMWFpXdkpybW9NMnNDUEdVbmciLCJ5IjoiN0Rwd21Pb0hJbmQwUWNSRVJUS1pBQ2k5YndzYTVnR0tER3hGeG00OEdSQSJ9fQ.eyJpc3MiOiJwb3N0bWFuIiwic3ViIjoiOWJiYWEyZjctNjlhOS00ZWFlLWI2YjgtOTRmYzY2MDExMmZjIiwiYXVkIjoiaHR0cDovL29wLmxvY2FsaG9zdC9yZWFsbXMvcmlkdCIsImlhdCI6MTY1OTM1NTIwNSwibmJmIjoxNjU5MzU1MjA1LCJleHAiOjE2NjkzNTUyMzUsIm5vbmNlIjoiVmpmVTQ2WjV5a0lobjdqSnpxWm9XSytwYXE2M0VLdUgiLCJ0b2tlbl9jbGFpbXMiOiJuYW1lIGVtYWlsIGVtYWlsX3ZlcmlmaWVkIiwidG9rZW5fbGlmZXRpbWUiOjM2MDAsInRva2VuX25vbmNlIjoiQmp4cTI3RlVsQjBYQVcyaWIrWnM2czU3UlFyY21VeEEifQ.BrfJYyrU1bZVWRawXO3Jowic3H84RaIzZDp_e8obviBlLLaq09tAnSUuVGLJ2hw4EIw1enALLtk_F5ZwEMqLlQ
```

It has the following payload (without comments):

```json
{
    "alg": "ES256",
    "typ": "JWT",
    "jwk": {
        // The client's public key:
        "kty": "EC",
        "crv": "P-256",
        "x": "cXQ8bdeNeeSwfLkHzMfAUFrHlLXZWvJrmoM2sCPGUng",
        "y": "7DpwmOoHInd0QcRERTKZACi9bwsa5gGKDGxFxm48GRA"
    }
}
```

and the following payload (without comments):

```json
{
    "iss": "postman", // The client ID.
    "sub": "9bbaa2f7-69a9-4eae-b6b8-94fc660112fc", // The user's unique identifier. In Keycloak, this is a UUID which is displayed in the Users menu.
    "aud": "http://op.localhost/realms/ict", // The OpenID Provider's URL = issuer of the ID Certifcation Token.
    "iat": 1659355205, // Unix timestamp when the token was issued.
    "nbf": 1659355205, // Unix timestamp when the token becomes valid.
    "exp": 1669355235, // Unix timestamp when the token expires.
    "nonce": "VjfU46Z5ykIhn7jJzqZoWK+paq63EKuH", // A random nonce.
    "token_claims": "name email email_verified", // The requested identity claims for the ID Certifcation Token.
    "token_lifetime": 3600, // The requested lifetime of the ID Certifcation Token.
    "token_nonce": "Bjxq27FUlB0XAW2ib+Zs6s57RQrcmUxA" // A random nonce to set into the ID Certifcation Token.
}
```

## Testing with API Specification

You can test the infrastructure with our API documentation.
This is recommended if you want to play with the API.

Therefore, you must authorize the API documentation as follows:

<details>
  <summary><b>For Public Authorization Server</b></summary>

1. Open your browser and navigate to the [API documentation (external URL)](https://api.oidc-e2e.primbs.dev/).
2. Click _Authorize_.
3. Scroll down to the authorization **oauth2_public**.
4. Enter the _client_id_ `api` and _Select all_ scopes.
5. Click _Authorize_ and _Sign In_ with your test user.
6. Click _Close_.

</details>
<details>
  <summary><b>For Local Authorization Server</b></summary>

1. Open your browser and navigate to the [API documentation (external URL)](https://api.oidc-e2e.primbs.dev/).
2. Click _Authorize_.
3. Scroll down to the authorization **oauth2_local**.
4. Enter the _client_id_ `api` and _Select all_ scopes.
5. Click _Authorize_ and _Sign In_ with your test user.
6. Click _Close_.

</details>

Now you can perform requests to the server as follows:

<details>
  <summary><b>For Public Authorization Server</b></summary>

1. Make sure that the server starting with URL `https://op.oidc-e2e.primbs.dev/...` is selected.
2. Open the _POST /_ Endpoint.
3. Click _Try it out_.
4. Paste a sufficient Token Request JWT to the _Request Body_.
5. Click _Execute_ to send the request.

</details>
<details>
  <summary><b>For Local Authorization Server</b></summary>

1. Make sure that the server starting with URL `http://op.localhost/...` is selected.
2. Open the _POST /_ Endpoint.
3. Click _Try it out_.
4. Paste a sufficient Token Request JWT to the _Request Body_.
5. Click _Execute_ to send the request.

</details>

## Testing with Swagger Editor

You can test the infrastructure with Swagger Editor.
This is recommended while editing the API Specification.

Therefore, you must authorize Swagger Editor as follows:

<details>
  <summary><b>For Public Authorization Server</b></summary>

1. Open your browser and navigate to the [Swagger Editor (external URL)](https://editor.swagger.io/).
2. Click _Authorize_.
3. Scroll down to the authorization **oauth2_public**.
4. Enter the _client_id_ `swagger` and _Select all_ scopes.
5. Click _Authorize_ and _Sign In_ with your test user.
6. Click _Close_.

</details>
<details>
  <summary><b>For Local Authorization Server</b></summary>

1. Open your browser and navigate to the [Swagger Editor (external URL)](https://editor.swagger.io/).
2. Click _Authorize_.
3. Scroll down to the authorization **oauth2_local**.
4. Enter the _client_id_ `swagger` and _Select all_ scopes.
5. Click _Authorize_ and _Sign In_ with your test user.
6. Click _Close_.

</details>

Now you can perform requests to the server as follows:

<details>
  <summary><b>For Public Authorization Server</b></summary>

1. Make sure that the server starting with URL `https://op.oidc-e2e.primbs.dev/...` is selected.
2. Open the _POST /_ Endpoint.
3. Click _Try it out_.
4. Paste a sufficient Token Request JWT to the _Request Body_.
5. Click _Execute_ to send the request.

</details>
<details>
  <summary><b>For Local Authorization Server</b></summary>

1. Make sure that the server starting with URL `http://op.localhost/...` is selected.
2. Open the _POST /_ Endpoint.
3. Click _Try it out_.
4. Paste a sufficient Token Request JWT to the _Request Body_.
5. Click _Execute_ to send the request.

</details>

## Testing with Postman

You can test the infrastructure with Postman.

Therefore, you must authorize Postman as follows:

<details>
  <summary><b>For Public Authorization Server</b></summary>

1. Open a new Tab and go to the _Authorization_ tab.
2. As _Type_, select `OAuth 2.0`.
3. In _Configure New Token_ > _Configuration Options_ insert the following values:
    - _Grant Type_: `Authorization Code (With PKCE)`
    - _Callback URL_: `https://oauth.pstmn.io/v1/callback` and tick _Authorize using browser_.
    - _Auth URL_: `https://op.oidc-e2e.primbs.dev/realms/ict/protocol/openid-connect/auth`
    - _Access Token URL_: `https://op.oidc-e2e.primbs.dev/realms/ict/protocol/openid-connect/token`
    - _Client ID_: `postman`
4. Click _Get New Access Token_.
5. _Sign In_ to your test user account, if requested.
6. Click _Use Token_.

</details>

<details>
  <summary><b>For Local Authorization Server</b></summary>

1. Open a new Tab and go to the _Authorization_ tab.
2. As _Type_, select `OAuth 2.0`.
3. In _Configure New Token_ > _Configuration Options_ insert the following values:
    - _Grant Type_: `Authorization Code (With PKCE)`
    - _Callback URL_: `https://oauth.pstmn.io/v1/callback` and tick _Authorize using browser_.
    - _Auth URL_: `http://op.localhost/realms/ict/protocol/openid-connect/auth`
    - _Access Token URL_: `http://op.localhost/realms/ict/protocol/openid-connect/token`
    - _Client ID_: `postman`
4. Click _Get New Access Token_.
5. _Sign In_ to your test user account, if requested.
6. Click _Use Token_.

</details>

Now you can perform requests to the server as follows:

<details>
  <summary><b>For Public Authorization Server</b></summary>

1. Select the HTTP Method _POST_.
2. Insert the URL `https://op.oidc-e2e.primbs.dev/realms/ict/protocol/openid-connect/userinfo/ict`.
3. Go to the _Body_ tab and insert the Token Request JWT as _raw_.
4. Click _Send_.

</details>

<details>
  <summary><b>For Local Authorization Server</b></summary>

1. Select the HTTP Method _POST_.
2. Insert the URL `http://op.localhost/realms/ict/protocol/openid-connect/userinfo/ict`.
3. Go to the _Body_ tab and insert the Token Request JWT as _raw_.
4. Click _Send_.

</details>
