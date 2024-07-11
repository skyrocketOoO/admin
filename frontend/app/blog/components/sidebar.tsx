import { useState } from 'react';
import { Story } from '../loadMarkdownFiles';

interface SidebarProps {
  stories: Story[];
  onSelectPost: (title: string) => void;
}

const Sidebar: React.FC<SidebarProps> = ({ stories, onSelectPost }) => {
  const [selectedType, setSelectedType] = useState('All');

  // Get unique story types
  const articleTypes = ['All', ...new Set(stories.map(story => story.topic))];

  // Filter stories by selected type
  const filteredstories = selectedType === 'All'
    ? stories
    : stories.filter(story => story.topic === selectedType);

  return (
    <aside className="w-80 bg-amber-50 p-4 min-h-full border">
      <h2 className="mb-8 lg:mb-3 font-semibold text-slate-900 text-xl">Stories</h2>
      
      <label htmlFor="typeFilter" className="block mb-4 text-gray-700">
        Topic:
      </label>
      <select
        id="typeFilter"
        value={selectedType}
        onChange={(e) => setSelectedType(e.target.value)}
        className="mb-6 p-2 border rounded w-full"
      >
        {articleTypes.map(type => (
          <option key={type} value={type}>{type}</option>
        ))}
      </select>

      <ul className='leading-6'>
        {filteredstories.map((story) => (
          <li key={story.title} className="text-lg block border-l pl-4 -ml-px border-transparent hover:border-slate-400 dark:hover:border-slate-500 text-slate-700 hover:text-slate-900 dark:text-slate-400 dark:hover:text-slate-300">
            <button
              className="text-gray-800 hover:underline"
              onClick={() => onSelectPost(story.title)}
            >
              {story.title}
            </button>
          </li>
        ))}
      </ul>
    </aside>
  );
};

export default Sidebar;
