
import React from 'react';
import Markdown  from 'react-markdown';
import remarkGfm from 'remark-gfm'
import './white.css';

interface Story {
  title: string;
  content: string;
  author: string;
  date: string;
}

const BlogStory: React.FC<Story> = ({ title, content, author, date }) => {
  let markdown = `A paragraph with *emphasis* and **strong importance**.

> A block quote with ~strikethrough~ and a URL: https://reactjs.org.

* Lists
* [ ] todo
* [x] done

A table:

| a | b |
| - | - |
12332
`;

  return (
    <article className="mb-8 markdown-body">
      <h2 className="text-2xl font-bold mb-2">{title}</h2>
      <p className="text-gray-600 mb-2">{author} - {date}</p>
      
      <Markdown remarkPlugins={[[remarkGfm]]}>
        {markdown}
      </Markdown>
    </article>
  );
};

export default BlogStory;
