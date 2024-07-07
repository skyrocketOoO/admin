import React from 'react';
import { useRouter } from 'next/router';
import BlogPage from '../page';
import BlogPost from '@/app/components/blog-post';

const articles = [
  { slug: 'first-post', title: 'First Blog Post', content: 'This is the content of the first blog post.', author: 'John Doe', date: 'July 7, 2024' },
  { slug: 'second-post', title: 'Second Blog Post', content: 'This is the content of the second blog post.', author: 'Jane Smith', date: 'July 8, 2024' },
];

const ArticlePage = () => {
  const router = useRouter();
  const { slug } = router.query;
  const article = articles.find((article) => article.slug === slug);

  if (!article) {
    return <p>Article not found</p>;
  }

  return (
    <BlogPage>
      <BlogPost 
        title={article.title}
        content={article.content}
        author={article.author}
        date={article.date}
      />
    </BlogPage>
  );
};

export default ArticlePage;
