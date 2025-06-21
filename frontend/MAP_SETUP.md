# Map Provider Setup Guide

This application supports **three different map providers**: Google Maps, Apple MapKit, and Mapbox. You can switch between them using the dropdown in the app.

## 🗺️ Mapbox (Recommended)

**Best Choice**: 50,000 free map views per month, excellent performance, beautiful styling.

### Setup Steps:

1. **Sign up for Mapbox** at [mapbox.com](https://mapbox.com)
2. **Get your PUBLIC Access Token** from [mapbox.com/account/access-tokens](https://account.mapbox.com/access-tokens/)
   
   ⚠️ **IMPORTANT**: Use the **Default public token** that starts with `pk.` (NOT the secret token that starts with `sk.`)
   
3. **Add to .env.local**:
   ```
   VUE_APP_MAPBOX_TOKEN=pk.eyJ1IjoieW91cl91c2VybmFtZSIsImEiOiJjbGFbc29tZXRoaW5nIn0.something
   ```

### Token Types:
- ✅ **Public Token** (`pk.*`) - Use this for frontend/browser
- ❌ **Secret Token** (`sk.*`) - Only for backend/server (will cause errors in browser)

### Features:
- ✅ **50,000 free requests/month**
- ✅ **No credit card required**
- ✅ **Beautiful custom styling**
- ✅ **Excellent performance**
- ✅ **Real-time directions**

---

## 🍎 Apple MapKit

**Good Choice**: Free with Apple Developer account, native iOS look.

### Setup Steps:

1. **Apple Developer Account** ($99/year required)
2. **Generate MapKit JWT Token** (server-side)
3. **Add to .env.local**:
   ```
   VUE_APP_MAPKIT_TOKEN=your_jwt_token_here
   ```

### Features:
- ✅ **Free with Apple dev account**
- ✅ **Native iOS styling**
- ✅ **Good performance**
- ❌ **Requires $99/year Apple dev account**
- ❌ **Complex JWT token setup**

---

## 🌍 Google Maps

**Classic Choice**: Most familiar, comprehensive features.

### Setup Steps:

1. **Google Cloud Console** at [console.cloud.google.com](https://console.cloud.google.com)
2. **Enable Maps JavaScript API**
3. **Get API Key**
4. **Add to .env.local**:
   ```
   VUE_APP_GOOGLE_MAPS_API_KEY=AIzaSyC...
   ```

### Features:
- ✅ **Most comprehensive**
- ✅ **Familiar interface**
- ✅ **Excellent directions**
- ❌ **Requires credit card**
- ❌ **Limited free tier**

---

## 🚀 Quick Start

1. **Choose your provider** (we recommend Mapbox)
2. **Get your API key/token**
3. **Create `.env.local` file**:
   ```bash
   cp .env.example .env.local
   ```
4. **Add your token**:
   ```
   VUE_APP_MAPBOX_TOKEN=your_token_here
   ```
5. **Restart the dev server**:
   ```bash
   npm run serve
   ```

## 🔧 Environment File Structure

Your `.env.local` should look like:

```bash
# Choose ONE or ALL providers
VUE_APP_GOOGLE_MAPS_API_KEY=your_google_key_here
VUE_APP_MAPKIT_TOKEN=your_mapkit_token_here
VUE_APP_MAPBOX_TOKEN=your_mapbox_token_here

# Backend API
VUE_APP_API_BASE_URL=http://localhost:8080
```

## 💡 Tips

- **Start with Mapbox** - it's the easiest and most generous
- **All providers work independently** - you can set up just one
- **Switch providers** using the dropdown in the app
- **Free tiers** are usually sufficient for development

## 🛠️ Troubleshooting

### Mapbox Issues:
- **Wrong token type**: Use PUBLIC token (`pk.*`), NOT secret token (`sk.*`)
- **Invalid token**: Verify token starts with `pk.` and is copied correctly
- **Permissions**: Check token permissions in Mapbox dashboard
- **Account**: Ensure token is for the correct account
- **Environment**: Restart dev server after updating .env.local

### MapKit Issues:
- Verify JWT token is valid
- Check Apple Developer account status
- Ensure MapKit is enabled in Apple Developer console

### Google Maps Issues:
- Verify API key is valid
- Check billing account is set up
- Ensure Maps JavaScript API is enabled

---

🎉 **Ready to Go!** Once you've set up your preferred provider, you'll have a fully functional map with:

- Color-coded facility markers
- User location detection
- Real-time directions
- Interactive popups
- Distance calculations
- Beautiful styling 