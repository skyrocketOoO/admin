import Link from 'next/link';

const Sidebar = ({ articles, onSelectPost }) => {
  return (
    <aside className="w-64 bg-gray-100 p-4">
      <h2 className="text-xl font-bold mb-4">Articles</h2>
      <ul>
        {articles.map((article) => (
          <li key={article.slug} className="mb-2">
            <button
              className="text-blue-500 hover:underline"
              onClick={() => onSelectPost(article.slug)}
            >
              {article.title}
            </button>
          </li>
        ))}
      </ul>
    </aside>
  );
};
export default Sidebar;