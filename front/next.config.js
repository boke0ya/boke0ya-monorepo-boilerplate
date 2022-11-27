/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  swcMinify: true,
  rewrites() {
    return [
      {
        source: `/api/:path*`,
        destination: `${process.env.API_BASE_URL}/v1/:path*`,
      }
    ]
  }
}

module.exports = nextConfig
