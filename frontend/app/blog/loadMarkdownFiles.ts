"use server";
import fs from 'fs/promises';
import path from 'path';

const postsDirectory = path.join(process.cwd(), 'app/blog/stories');

export interface PostData {
  title: string;
  content: string;
}

export async function getMarkdownFiles(): Promise<PostData[]> {
  const fileNames = await fs.readdir(postsDirectory);
  const allPostsData: PostData[] = await Promise.all(
    fileNames.map(async (fileName) => {
      const fullPath = path.join(postsDirectory, fileName);
      const fileContents = await fs.readFile(fullPath, 'utf8');

      return {
        title: fileName.replace(/\.md$/, ''),
        content: fileContents,
      };
    })
  );
  return allPostsData;
}
