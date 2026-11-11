import { GetCurrentOS } from "$lib/bindings/App";

// This file is used to determine if the app is running on the website(vercel client only) or not.
const onwebsite = import.meta.env?.VITE_ON_WEBSITE === 'true';

export const IS_CLIENT_ONLY = async () => {
  if (onwebsite) return true
  
  try {    
    return undefined === await GetCurrentOS()
  } catch {
    return true
  }
}

export default onwebsite;