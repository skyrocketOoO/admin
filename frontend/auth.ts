// import NextAuth from 'next-auth';
// import CredentialsProvider  from 'next-auth/providers/credentials';
// import { authConfig } from './auth.config';
// import { z } from 'zod';
// import { MainClient, LoginRequest, LoginResponse } from '@/proto/main';
// import * as grpc from '@grpc/grpc-js';
 
// type User = {
//   name: string;
//   password: string;
//   sessionID: string;
// };

// async function getUser(username: string, password: string): Promise<User | undefined> {
//   const client = new MainClient("localhost:50051", grpc.credentials.createInsecure());
//   const loginReq: LoginRequest = {
//     UserName: username,
//     Password: password,
//   };
//   try {
//     return new Promise((resolve, reject) => {
//       client.login(loginReq, (error: grpc.ServiceError | null, resp: LoginResponse) => {
//         if (error) {
//           console.error('Error:', error.message);
//           reject(new Error('Failed to login.'));
//         } else {
//           const user: User = {
//             sessionID: resp.SessionID,
//             name: username,
//             password: password
//           };
//           resolve(user);
//         }
//       });
//     });
//   }catch (err) {
//     console.error('Failed to fetch user:', err);
//     throw new Error('Failed to fetch user.');
//   }
// }
 
// export const { auth, signIn, signOut } = NextAuth({
//   ...authConfig,
//   providers: [
//     CredentialsProvider({
//       name: 'Credentials',
//       credentials: {
//         username: { label: "Username", type: "text" },
//         password: {  label: "Password", type: "password" }
//       },
//       authorize: async (credentials) => {
//         const parsedCredentials = z
//           .object({ username: z.string(), password: z.string() })
//           .safeParse(credentials);
 
//         if (parsedCredentials.success) {
//           const { username, password } = parsedCredentials.data;
//           const user = await getUser(username, password);
//           if (!user) return null;
//           return user;
//         }
 
//         return null;
//       },
//     }),
//   ],
//   pages: {
//     signIn: '/login',  // Displays the login page
//   }
// });