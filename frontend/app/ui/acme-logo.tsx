import { GlobeAltIcon } from '@heroicons/react/24/outline';
import { lusitana } from '@/app/ui/fonts';

export default function AcmeLogo() {
  return (
    <div
      className={`${lusitana.className} flex flex-row items-center leading-none text-sky-300`}
    >
      <img src="github-mark-white.png" className="h-12 w-12 mr-2" />
      <p className="text-[24px] whitespace-nowrap">skyrocketOoO Playground</p>
    </div>
  );
}
