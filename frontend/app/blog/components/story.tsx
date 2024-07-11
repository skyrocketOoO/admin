
import React from 'react';
import Markdown  from 'react-markdown';
import remarkGfm from 'remark-gfm'
import './white.css';

interface Story {
  title: string;
  content: string;
  // author: string;
  // date: string;
}

const BlogStory: React.FC<Story> = ({ title, content }) => {

  return (
    <article className="markdown-body w-full h-full p-12">
      <div className="font-medium text-5xl" >{title}</div>
      <div className="p-2"></div>
      {/* <p className="text-gray-900 font-medium">{date}</p> */}
      <Markdown remarkPlugins={[[remarkGfm]]}>
        {content}
      </Markdown>
    </article>
  );
};

export default BlogStory;
