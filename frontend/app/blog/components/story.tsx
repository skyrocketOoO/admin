import React from 'react';

interface Story {
  title: string;
  content: string;
  author: string;
  date: string;
}

const BlogStory: React.FC<Story> = ({ title, content, author, date }) => {
  return (
    <article className="mb-8">
      <h2 className="text-2xl font-bold mb-2">{title}</h2>
      <p className="text-gray-600 mb-2">{author} - {date}</p>
      <p>{content}</p>
    </article>
  );
};

export default BlogStory;
