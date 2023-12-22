# Visionline Integration Example Go

To run the steps you just need have Go 1.21.4 or later installed and run `go run .` inside of any step folder.

## Step One: Create Webviews (One Time Only)

To connect to visionline, you must create a 3 webviews, each for one of the services you're connecting:
* Seam Bridge (to connect you to the Visionline Network)
* Assa Abloy Credential Service
* On-Prem Visionline

Each webview has a URL, you go to this URL to log in!

Here are the commands, after executing each one, open the webview and fill
in the sample credentials:

- `seam connect-webviews create --accepted-providers seam_bridge`
  - Pairing Code: `1234`
  - Name: `My Network`
- `seam connect-webviews create --accepted-providers assa_abloy_credential_service`
  - Username: `jane`
  - Password: `1234`
- `seam connect-webviews create --accepted-providers seam_bridge`
  - Username: `jane`
  - Password: `1234`
  - Lan IP: `192.168.1.100`

---

> All steps beyond this point you repeat for each incoming Guest

## Step Two: Create a User Identity to Represent the Guest

- `seam user-identities create --email-address jane@example.com`

## Step Three: Enroll the User Identity to allow Phones to Sync

- `seam user-identities enrollment-automations launch`
  - For `user_identity_id` select `jane@example.com`

TODO

## Step Four: Create the Visionline User

Now we create a Visionline User and connect it to our User Identity.

## Step Five: Create the Visionline Credential

TODO
