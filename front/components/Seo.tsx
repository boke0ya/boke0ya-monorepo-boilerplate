import Head from "next/head";

interface SeoProps {
  title?: string;
  description?: string;
  imageUrl?: string;
}

const Seo = ({
  title = "",
  description = "",
  imageUrl = process.env.NEXT_PUBLIC_BASE_URL + "/ogp.png",
}: SeoProps) => {
  if (title === "") {
    title = "";
  }
  return (
    <Head>
      <title>{title}</title>
      <meta name="description" content={description} />
      <meta property="og:description" content={description} />
      <meta property="og:url" content={process.env.NEXT_PUBLIC_BASE_URL} />
      <meta property="og:type" content="website" />
      <meta property="og:image" content={imageUrl} />
      <meta property="og:title" content={title} />
      <meta name="twitter:card" content="summary_large_image" />
      <meta name="twitter:site" content="@" />
      <meta name="twitter:image" content={imageUrl} />
      <meta name="twitter:title" content={title} />
      <meta name="twitter:description" content={description} />
    </Head>
  );
};

export default Seo
