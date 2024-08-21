'use client';
import React, { useEffect } from 'react';
import Head from 'next/head';
import Navbar from '../components/navbar/Navbar';
import Sidebar from '../components/sidebar/Sidebar';
import { Inter } from 'next/font/google';
import { useSession } from 'next-auth/react';
import { useRouter } from 'next/navigation';
import { useGetCurrentUserQuery } from '@/lib/service/UserService';
import { useDispatch, useSelector } from 'react-redux';
import { RootState, AppDispatch } from '@/lib/store'; // Adjust path as necessary
import { setUser } from '@/lib/features/userSlice/userSlice'; // Adjust path as necessary

const inter = Inter({ subsets: ['latin'] });

const Layout = ({ children, title = 'My Next.js App' }: { children: React.ReactNode; title?: string }) => {
  const { data: session, status } = useSession();
  const router = useRouter();
  const dispatch = useDispatch<AppDispatch>();
  const user = useSelector((state: RootState) => state.user.user);

  const { data: userData, isLoading } = useGetCurrentUserQuery(session?.user?.accessToken ?? '', {
    skip: !session?.user?.accessToken,
  });

  useEffect(() => {
    if (status === 'unauthenticated') {
      router.push('/login');
    }
  }, [status, router]);

  useEffect(() => {
    if (userData?.data) {
      dispatch(setUser(userData.data));
    }
  }, [userData, dispatch]);

  if (status === 'loading' || isLoading) {
    return <p>Loading...</p>;
  }

  return (
    <>
      <Head>
        <title>{title}</title>
        <meta name="viewport" content="width=device-width, initial-scale=1" />
      </Head>
      <div className={`${inter.className} flex flex-col min-h-screen`}>
        <Navbar />
        <div className="flex flex-1">
          <Sidebar />
          <main className="max-md:pt-[100px] flex-1 p-4 mt-[60px] lg:ml-[240px] sm:ml-[240px] ml-0 bg-[#F5F7FA]">
            {children}
          </main>
        </div>
      </div>
    </>
  );
};

export default Layout;
