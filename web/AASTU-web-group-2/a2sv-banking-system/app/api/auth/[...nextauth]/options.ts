import CredentialsProvider from "next-auth/providers/credentials";
import { NextAuthOptions } from 'next-auth';
import { login, refreshToken } from "@/lib/api/authenticationController";
import { jwtDecode, JwtPayload } from "jwt-decode";
export const options: NextAuthOptions = {
  session: {
    strategy: "jwt", // Use JWT for session strategy
  },


  providers: [
    CredentialsProvider({
      name: "credentials",
      credentials: {
        userName: { label: "Username", type: "string" },
        password: { label: "Password", type: "string" }
      },
    
      async authorize(credentials, req) {
        try {
          const response = await login(credentials as { userName: string; password: string });
          if (response.success) {
            return response.data; // Return the user data object
          } else {
            return null;
          }
        } catch (error) {
          console.error('Authorization error:', error);
          return null;
        }
      },
    })
    
  ],

  pages:{
    signIn: '/api/auth/signin'
    // signUp: '/api/auth/signup',
  },
  callbacks: {
    // Store the user information in the JWT token
    async jwt({ token, user }: any) {
      if (user) {
        token.access_token = user.access_token;
        token.data = user.data;     // Assuming the user object has a 'name' field
        token.refresh_token = user.refresh_token
      }
      
      // Decode the access token to check expiry
      const decodedToken = jwtDecode<JwtPayload>(token.access_token);
      const currentTime = Date.now() / 1000;

      if (decodedToken && decodedToken.exp !== undefined && decodedToken.exp < currentTime) {
        // If the token is expired, refresh it
        try {
          const newTokens = await refreshToken();
          token.access_token = newTokens.data;
        } catch (error) {
          console.error("Failed to refresh access token:", error);
        }
      }

      return token;
    },
    // Make custom user data available in the session
    async session({ session, token }: any) {
      session.user.access_token = token.access_token;
      session.user.data = token.data;
      session.user.refresh_token = token.refresh_token;
      return session;
    },
  },
};
