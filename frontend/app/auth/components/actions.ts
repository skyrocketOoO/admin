import { signIn } from '@/app/auth/auth';

export function authenticate(
  formData: FormData,
) {
  try {
    signIn(formData);
  } catch (error) {
    switch (error) {
      case 'CredentialsSignin':
        return 'Invalid credentials.';
      default:
        throw error;
    }
  }
}