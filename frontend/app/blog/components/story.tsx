
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
A paragraph with *emphasis* and **strong importance**.

> A block quote with ~strikethrough~ and a URL: https://reactjs.org.

* Lists
* [ ] todo
* [x] done

A table:

| a | b |
| - | - |
12332
A paragraph with *emphasis* and **strong importance**.

> A block quote with ~strikethrough~ and a URL: https://reactjs.org.

* Lists
* [ ] todo
* [x] done

A table:

| a | b |
| - | - |
12332
A paragraph with *emphasis* and **strong importance**.

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
    <article className="markdown-body w-full h-full">
      <p className="font-medium" style={{ fontSize: '4rem' }}>{title}</p>
      <p className="text-gray-900 font-medium">{date}</p>
      
      <Markdown remarkPlugins={[[remarkGfm]]}>
        {markdown}
      </Markdown>
    </article>
  );
};

export default BlogStory;
