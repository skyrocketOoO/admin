import AcmeLogo from '@/app/ui/acme-logo';
import LoginForm from '@/app/ui/login-form';
 
export default function LoginPage() {
  return (
    <
      main className="flex items-center justify-center md:h-screen"
      style={{ backgroundImage: "url('/login_background.png')" }}
    >
      <div className="absolute right-20 mx-auto flex w-full max-w-[400px] flex-col space-y-2.5 p-4 md:-mt-32">
        <div className="flex h-5 w-full items-end rounded-lg bg-neutral-900 p-3 md:h-20">
          <div className="w-32 text-white md:w-36 ">
            <AcmeLogo />
          </div>
        </div>
        <LoginForm />
      </div>
    </main>
  );
}