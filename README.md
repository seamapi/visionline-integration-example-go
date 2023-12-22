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
  - For `credential_manager_acs_system_id` select "Assa Abloy Credential Service"
  - For `create_credential_manager_user` select `true`


## Step Four: Create the Visionline User

- `seam acs users create`
  - For `acs_system_id` select `Visionline`
  - For `user_identity_id` select `jane@example.com`
  - For `full_name` enter `Jane Doe` (or anything, but cannot be blank)
  - Our command line doesn't support editing the `access_schedule` at the
    moment, but normally you would want to set this to the duration of the
    Guest's stay

Now we create a Visionline User and connect it to our User Identity.

## Step Five: Create the Visionline Credential

- `seam acs credentials create`
  - For `acs_user_id` select `Visionline` then `jane@example.com`
  - For `access_method` select `mobile_key`
  - For `is_multi_phone_sync_credential` select `true`


