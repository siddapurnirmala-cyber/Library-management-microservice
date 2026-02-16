# Setting up Google OAuth 2.0 Credentials

To enable Google Login, you need to create a project in the Google Cloud Console and obtain credentials.

## Step 1: Create a Google Cloud Project
1. Go to the [Google Cloud Console](https://console.cloud.google.com/).
2. Click on the project dropdown at the top left and select **"New Project"**.
3. Enter a project name (e.g., "Library System") and click **"Create"**.
4. Once created, select the project from the notification or project dropdown.

## Step 2: Configure OAuth Consent Screen
1. In the left sidebar, navigate to **APIs & Services** > **OAuth consent screen**.
2. Select **External** (unless you are a Google Workspace user and want it internal only) and click **Create**.
3. Fill in the required fields:
   - **App Name**: Library System
   - **User Support Email**: Select your email
   - **Developer Contact Information**: Enter your email
4. Click **Save and Continue**.
5. You can skip "Scopes" for now (we use default email/profile). Click **Save and Continue**.
6. **Test Users**: Add your own Google email address under "Test Users" so you can log in while the app is in "Testing" mode.
7. Click **Save and Continue** and then **Back to Dashboard**.

## Step 3: Create Credentials
1. In the left sidebar, go to **APIs & Services** > **Credentials**.
2. Click **+ CREATE CREDENTIALS** at the top and select **OAuth client ID**.
3. **Application type**: Select **Web application**.
4. **Name**: Web Client 1 (default is fine).
5. **Authorized redirect URIs**: This is critical.
   - **Authorized JavaScript origins**: Enter `http://localhost:8082` (Do NOT include a path or trailing slash).
   - **Authorized redirect URIs**: Here you enter the full callback path.
     - Click **+ ADD URI**.
     - Enter: `http://localhost:8082/auth/google/callback`
   (Note: Use port `8080` or `8082` depending on what your server actually runs on so check `.env`. If `PORT=8082` in `.env`, use `http://localhost:8082/auth/google/callback`).
6. Click **Create**.

## Step 4: Get Your Credentials
1. A popup will show your **Client ID** and **Client Secret**.
2. Copy these strings.

## Step 5: Update `.env`
Open your `.env` file and paste the values:

```env
GOOGLE_CLIENT_ID=your_copied_client_id_ending_in_googleusercontent.com
GOOGLE_CLIENT_SECRET=your_copied_client_secret
JWT_SECRET=any_random_long_string_for_security
```

### Note on Port
Your `.env` seems to have `PORT=8082`.
- If your server runs on port **8082**, your Redirect URI in Google Cloud MUST be `http://localhost:8082/auth/google/callback`.
- If you change the port, you must update the Redirect URI in Google Cloud Console.
