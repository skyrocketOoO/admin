"use client";
import React, { useState, useEffect } from 'react';
import Sidebar from './components/sidebar';
import Header from './components/header';
import BlogStory from './components/story';
import './scroll.css';
import { getMarkdownFiles, Story } from './loadMarkdownFiles';

const mainpage: Story = {
  title: 'Welcome to my playground!!',
  date: '2024/7/12',
  topic: 'Mainpage',
  content: 
  `
  This is a place where I write some stories about technology, personal... or lifestyle
  
   This is also a place I'll be experimenting with new technologies and learning about the world around me.
  `
}

const BlogPage = () => {
  const [stories, setArticles] = useState<Story[]>([]);
  const [selectedPost, setSelectedPost] = useState<Story | null>(null);

  useEffect(() => {
    const fetchArticles = async () => {
      const loadedArticles = await getMarkdownFiles();
      setArticles(loadedArticles);
      // if (loadedArticles.length > 0) {
      //   setSelectedPost(loadedArticles[0]);
      // }
    };

    fetchArticles();
  }, []);

  const handleSelectPost = (title: string) => {
    const post = stories.find((story) => story.title === title);
    if (post) {
      setSelectedPost(post);
    }
  };

  return (
    <div className="flex flex-col min-h-screen">
      <Header />
      <div className="flex flex-grow">
        <Sidebar stories={stories} onSelectPost={handleSelectPost} />
          {selectedPost ? (
            <BlogStory
              title={selectedPost.title}
              content={selectedPost.content}
              date={selectedPost.date}
              topic={selectedPost.topic}
            />
          ) : (
            <BlogStory
            title={mainpage.title}
            content={mainpage.content}
            date={mainpage.date}
            topic={mainpage.topic}
          />
          )}
      </div>
    </div>
  );
};

export default BlogPage;
