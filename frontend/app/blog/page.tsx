import React from 'react';
import Sidebar from './components/sidebar';
import Header from './components/header';



const articles = [
  { slug: 'first-post', title: 'First Blog Post', content: 'This is the content of the first blog post.', author: 'John Doe', date: 'July 7, 2024' },
  { slug: 'second-post', title: 'Second Blog Post', content: 'This is the content of the second blog post.', author: 'Jane Smith', date: 'July 8, 2024' },
];

interface BlogPostProps {
  title: string;
  content: string;
  author: string;
  date: string;
}

const BlogPost: React.FC<BlogPostProps> = ({ title, content, author, date }) => {
  return (
    <article className="mb-8">
      <h2 className="text-2xl font-bold mb-2">{title}</h2>
      <p className="text-gray-600 mb-2">{author} - {date}</p>
      <p>{content}</p>
    </article>
  );
};


const BlogPage = () => {
  const posts = [
    {
      title: 'First Blog Post',
      content: 'This is the content of the first blog post.',
      author: 'John Doe',
      date: 'July 7, 2024',
    },
    {
      title: 'Second Blog Post',
      content: 'This is the content of the second blog post.',
      author: 'Jane Smith',
      date: 'July 8, 2024',
    },
  ];

  return (
    <div className="flex flex-col min-h-screen">
      <Header />
      <div className="flex flex-grow">
        <Sidebar articles={articles} />
        <main className="flex-grow container mx-auto p-4">
          {posts.map((post, index) => (
            <BlogPost 
              key={index}
              title={post.title}
              content={post.content}
              author={post.author}
              date={post.date}
            />
          ))}
        </main>
      </div>
    </div>
  );
};

export default BlogPage;