## How to Configure a Spotify App on developer.spotify.com

To use this Spotify MCP server, you need to create and configure a Spotify App to obtain your credentials (Client ID and Client Secret). Follow these steps:

### 1. Log in to Spotify Developer Dashboard
- Go to [https://developer.spotify.com/dashboard](https://developer.spotify.com/dashboard)
- Log in with your Spotify account.

### 2. Create a New App
- Click on **"Create an App"**.
- Enter an **App name** and **App description** (e.g., "Spotify MCP Server").
- Click **"Create"**.

### 3. View Your Client ID and Client Secret
- After creating the app, click on your app in the dashboard.
- You will see your **Client ID** on the app page.
- Click **"Show Client Secret"** to reveal your **Client Secret**.

### 4. Set Redirect URIs
- In your app page, click **"Edit Settings"**.
- Under **Redirect URIs**, add the URI `http://localhost:8821/callback`.
- Click **"Add"** and then **"Save"**.

### 5. Use Your Credentials
- Use the **Client ID**, **Client Secret**, and **Redirect URI** in your MCP server configuration or environment variables as required by the project.

---

For more details, see the [Spotify Developer Documentation](https://developer.spotify.com/documentation/web-api/).

