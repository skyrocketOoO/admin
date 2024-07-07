"use client"
import React from 'react';
import Link from 'next/link';
import { useRouter } from 'next/navigation';
import { useParams } from 'next/navigation';

const Header = () => {
  return (
    <header className="bg-blue-500 text-white p-4 flex justify-between items-center">
      <h1 className="text-3xl">My Blog</h1>
      <nav>
        <Link href="/"><a className="mr-4">Home</a></Link>
        <Link href="/admin"><a className="mr-4">Admin</a></Link>
        <Link href="/blog"><a>Blog</a></Link>
      </nav>
    </header>
  );
};

const Sidebar = ({ articles }) => {
  return (
    <aside className="w-64 bg-gray-100 p-4">
      <h2 className="text-xl font-bold mb-4">Articles</h2>
      <ul>
        {articles.map((article) => (
          <li key={article.slug} className="mb-2">
            <Link href={`/blog/${article.slug}`}>
              <a className="text-blue-500 hover:underline">{article.title}</a>
            </Link>
          </li>
        ))}
      </ul>
    </aside>
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


const articles = [
  { slug: 'first-post', title: 'First Blog Post', content: 'This is the content of the first blog post.', author: 'John Doe', date: 'July 7, 2024' },
  { slug: 'second-post', title: 'Second Blog Post', content: 'This is the content of the second blog post.', author: 'Jane Smith', date: 'July 8, 2024' },
];

const BlogPage = () => {
  const params = useParams();
  const { slug } = params;
  const article = articles.find((article) => article.slug === slug);

  if (!article) return <p>Article not found</p>;

  return (
    <div className="flex flex-col min-h-screen">
      <Header />
      <div className="flex flex-grow">
        <Sidebar articles={articles} />
        <main className="flex-grow container mx-auto p-4">
          <BlogPost
            title={article.title}
            content={article.content}
            author={article.author}
            date={article.date}
          />
        </main>
      </div>
    </div>
  );
};

export default BlogPage;