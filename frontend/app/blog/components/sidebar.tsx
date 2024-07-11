import Link from 'next/link';

const Sidebar = ({ articles, onSelectPost }) => {
  return (
    <aside className="w-60 bg-amber-50 p-4 min-h-full border">
      <h2 className="mb-8 lg:mb-3 font-semibold text-slate-900 dark:text-slate-200 text-xl">Stories</h2>
      <ul className='leading-6'>
        {articles.map((article) => (
          <li key={article.title} className="text-lg block border-l pl-4 -ml-px border-transparent hover:border-slate-400 dark:hover:border-slate-500 text-slate-700 hover:text-slate-900 dark:text-slate-400 dark:hover:text-slate-300">
            <button
              className="text-gray-800 hover:underline"
              onClick={() => onSelectPost(article.title)}
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