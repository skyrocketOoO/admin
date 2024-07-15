"use server";
import fs from 'fs/promises';
import path from 'path';

const postsDirectory = path.join(process.cwd(), 'app/blog/stories');

export interface Story {
  title: string;
  content: string;
  date: string;
  topic: string;
}

export async function getMarkdownFiles(): Promise<Story[]> {
  const fileNames = await fs.readdir(postsDirectory);
  const allPostsData: Story[] = await Promise.all(
    fileNames.map(async (fileName) => {
      const fullPath = path.join(postsDirectory, fileName);
      const fileContents = await fs.readFile(fullPath, 'utf8');
      
      // Split the file contents by lines
      const lines = fileContents.split('\n');
      const date = lines[0] || '';
      const topic = lines[1] || '';
      const content = lines.slice(2).join('\n');

      return {
        title: fileName.replace(/\.md$/, '').replace('_', ' '),
        content: content,
        date: date,
        topic: topic,
      };
    })
  );
  return allPostsData;
}
