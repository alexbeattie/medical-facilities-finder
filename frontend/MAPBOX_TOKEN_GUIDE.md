# 🗝️ Mapbox Token Quick Guide

## ✅ **What You Need: PUBLIC Token**

When you visit [mapbox.com/account/access-tokens](https://account.mapbox.com/access-tokens/), you'll see two types of tokens:

### **DEFAULT PUBLIC TOKEN** ✅
```
pk.eyJ1IjoidXNlcm5hbWUiLCJhIjoiY2xhMGR...
^^^ Starts with "pk." - USE THIS ONE!
```

### **SECRET TOKEN** ❌
```
sk.eyJ1IjoidXNlcm5hbWUiLCJhIjoiY2xhMGR...
^^^ Starts with "sk." - DON'T USE THIS!
```

## 🎯 **Quick Steps**

1. **Go to**: [mapbox.com/account/access-tokens](https://account.mapbox.com/access-tokens/)

2. **Copy the Default public token** (the long string starting with `pk.`)

3. **Add to `.env.local`**:
   ```bash
   VUE_APP_MAPBOX_TOKEN=pk.eyJ1IjoidXNlcm5hbWUiLCJhIjoiY2xhMGR...
   ```

4. **Restart your dev server**:
   ```bash
   npm run serve
   ```

## 🚨 **Common Errors**

### **"Use a public access token (pk.*)"**
- **Problem**: You copied the secret token (`sk.*`)
- **Solution**: Copy the PUBLIC token (`pk.*`) instead

### **"Invalid token format"**
- **Problem**: Token doesn't start with `pk.`
- **Solution**: Make sure you copied the complete token

### **"Failed to initialize"**
- **Problem**: Token might be expired or invalid
- **Solution**: Generate a new token at mapbox.com

## 💡 **Token Permissions**

Your public token needs these scopes:
- ✅ **Maps** (for displaying maps)
- ✅ **Directions** (for turn-by-turn directions)
- ✅ **Geocoding** (for address search)

These are included by default in the public token.

## 🔄 **Need a New Token?**

1. Go to [mapbox.com/account/access-tokens](https://account.mapbox.com/access-tokens/)
2. Click **"Create a token"**
3. Give it a name (e.g., "Medical Facilities App")
4. Make sure it's a **Public token**
5. Copy the token that starts with `pk.`

## ✨ **You'll Know It's Working When:**

- The map loads without errors
- You see beautiful Mapbox styling
- Facility markers appear correctly
- Directions work smoothly

**Happy mapping!** 🗺️✨ 