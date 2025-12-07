import type { NextConfig } from "next";
import path from "path";

const nextConfig: NextConfig = {
  /* config options here */
  reactCompiler: true,
  cacheComponents: true,
  // turbopack:  {
  //   root: path.join(__dirname,'frontend'),
  // }
};

export default nextConfig;
