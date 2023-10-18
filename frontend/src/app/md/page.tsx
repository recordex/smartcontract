import Layout from "@/app/layout";
import matter from "gray-matter";
import {remark} from "remark";
import html from "remark-html";

async function getPostData() {
  const id = 1;
  const fileContents = "# こんにちは";

  // Use gray-matter to parse the post metadata section
  const matterResult = matter(fileContents);

  // Use remark to convert markdown into HTML string
  const processedContent = await remark()
    .use(html)
    .process(matterResult.content);
  const contentHtml = processedContent.toString();

  // Combine the data with the id and contentHtml
  return {
    title: id,
    id: contentHtml,
    ...matterResult.data,
  };
}


export default async function Page() {
  const postData = await getPostData();

  return (
    <Layout>
      {postData.title}
      <br/>
      {postData.id}
      {/*<br/>*/}
      {/*{postData.date}*/}
      {/*<br/>*/}
      {/*<div dangerouslySetInnerHTML={{__html: postData.contentHtml}}/>*/}
    </Layout>
  );
}
