"use client"
import React, { useState } from 'react';
import Sidebar from './components/sidebar';
import Header from './components/header';
import BlogStory from './components/story';

const articles = [
  { slug: 'first-post', title: 'First Blog Post', content: 'This is the content of the first blog post.', author: 'John Doe', date: 'July 7, 2024' },
  { slug: 'second-post', title: 'Second Blog Post', content: 'This is the content of the second blog post.', author: 'Jane Smith', date: 'July 8, 2024' },
];

const BlogPage = () => {
  const [selectedPost, setSelectedPost] = useState(articles[0]);

  const handleSelectPost = (slug: string) => {
    const post = articles.find((article) => article.slug === slug);
    if (post) {
      post.content = `
        # Markdown Example
        - List item 1
        - List item 2
      `;
      setSelectedPost(post);
    }
  };

  return (
    <div className="flex flex-col min-h-screen">
      <Header />
      <div className="flex flex-grow">
        <Sidebar articles={articles} onSelectPost={handleSelectPost} />
        <main className="flex-grow container mx-auto p-4 bg-gray-700">
          {selectedPost ? (
            <BlogStory
              title={selectedPost.title}
              content={selectedPost.content}
              author={selectedPost.author}
              date={selectedPost.date}
            />
          ) : (
            <p>No story selected.</p>
          )}
        </main>
      </div>
    </div>
  );
};

export default BlogPage;