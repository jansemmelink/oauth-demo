package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func main() {
	codeVerifierStr := "34c8a34770dc9c10c2bcd6e432d00ef9d7fd84f16d89f45939af0bb3"
	h := sha256.New()
	h.Write([]byte(codeVerifierStr))
	hashedCodeVerifier := h.Sum(nil)
	fmt.Printf("##=%X\n", hashedCodeVerifier)
	fmt.Printf("ch=%s\n", base64.URLEncoding.EncodeToString(hashedCodeVerifier))
}

// curl -XGET 'https://dev-n8598cw4.us.auth0.com/authorize?response_type=code&client_id=llanHIAyvyoTh0VLGpYIEQcXHxeZcUxx&state=testing123&redirect_uri=https://example-app.com/redirect&code_challenge=4hHvOBn_FGgIYzA_dixxR4CdzWor4cbIMZ-rbAnH8Dg&code_challenge_method=S256'
// Found. Redirecting to /u/login?state=hKFo2SBvZVM0a0I4NjlETmhNM0xsY3lnTE1kYmJDd05qODRUcaFur3VuaXZlcnNhbC1sb2dpbqN0aWTZIFZLYjZqTGE2d3RxMlJjWE8xYVBsZ1lUa0NXdjZYc052o2NpZNkgbGxhbkhJQXl2eW9UaDBWTEdwWUlFUWNYSHhlWmNVeHg%

// Congrats!
// The authorization server redirected you back to the app and issued an authorization code!
// You can exchange this authorization code for an access token now!
// Your app can read the authorization code and state from the URL, and they are printed below for your convenience as well.
// code=
// 1UhokHP88pnDF238nX4YDpy9QWiq6A2Vl-Etx_HLLuiuK
// state=Testing123
// You should verify that the state parameter here matches the one you set at the beginning. Otherwise it's possible someone is trying to trick your app!

//curl -X POST https://dev-n8598cw4.us.auth0.com/oauth/token \
// -d grant_type=authorization_code \
// -d redirect_uri=https://example-app.com/redirect \
// -d client_id={YOUR_CLIENT_ID} \ 
// -d client_secret={YOUR_CLIENT_SECRET} \
// -d code_verifier={YOUR_CODE_VERIFIER} \
// -d code={YOUR_AUTHORIZATION_CODE}

curl -X POST https://dev-n8598cw4.us.auth0.com/oauth/token -d grant_type=authorization_code -d redirect_uri=https://example-app.com/redirect -d client_id=34c8a34770dc9c10c2bcd6e432d00ef9d7fd84f16d89f45939af0bb3 -d client_secret=4qLvILzVaPMnGAJPbYEdCxgdK9Igu_h74l4AF9X4RavtjB-NdWKBu-pradXWlHZG -d code_verifier=34c8a34770dc9c10c2bcd6e432d00ef9d7fd84f16d89f45939af0bb3 -d code=