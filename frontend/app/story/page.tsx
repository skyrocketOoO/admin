import React from 'react';

const Header = () => {
  return (
    <header className="bg-blue-500 text-white p-4">
      <h1 className="text-3xl">My Blog</h1>
    </header>
  );
};



const Footer = () => {
  return (
    <footer className="bg-gray-800 text-white p-4 mt-8">
      <p className="text-center">&copy; 2024 My Blog. All rights reserved.</p>
    </footer>
  );
};

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
      <Footer />
    </div>
  );
};

export default BlogPage;