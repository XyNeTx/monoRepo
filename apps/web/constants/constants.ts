export const PUBLIC_API_LINK = process.env.NEXT_PUBLIC_API_LINK;
export const SSR_API_LINK = process.env.SSR_API_LINK;
export const IsProd = process.env.NODE_ENV == "production"
export const API_LINK = IsProd ? SSR_API_LINK : PUBLIC_API_LINK