"use client";
import React, { useState, useEffect } from 'react';
import Sidebar from './components/sidebar';
import Header from './components/header';
import BlogStory from './components/story';
import './scroll.css';
import { getMarkdownFiles, PostData } from './loadMarkdownFiles';

const BlogPage = () => {
  const [articles, setArticles] = useState<PostData[]>([]);
  const [selectedPost, setSelectedPost] = useState<PostData | null>(null);

  useEffect(() => {
    const fetchArticles = async () => {
      const loadedArticles = await getMarkdownFiles();
      setArticles(loadedArticles);
      if (loadedArticles.length > 0) {
        setSelectedPost(loadedArticles[0]);
      }
    };

    fetchArticles();
  }, []);

  const handleSelectPost = (title: string) => {
    const post = articles.find((article) => article.title === title);
    if (post) {
      setSelectedPost(post);
    }
  };

  return (
    <div className="flex flex-col min-h-screen">
      <Header />
      <div className="flex flex-grow">
        <Sidebar articles={articles} onSelectPost={handleSelectPost} />
        {/* <div className="flex-grow  mx-0 bg-gray-700 "> */}
          {selectedPost ? (
            <BlogStory
              title={selectedPost.title}
              content={selectedPost.content}
            />
          ) : (
            <p>No story selected.</p>
          )}
        {/* </div> */}
      </div>
    </div>
  );
};

export default BlogPage;
