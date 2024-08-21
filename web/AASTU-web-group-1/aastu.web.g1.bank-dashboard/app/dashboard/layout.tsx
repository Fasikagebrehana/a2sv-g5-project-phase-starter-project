"use client";
import Image from "next/image";
import Header from "./_components/Header";
import Sidebar from "./_components/Sidebar";
import MobileSidebar from "./_components/MobileSidebar";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { usePathname } from "next/navigation";
import { sidebarLinks } from "@/constants";
import { SessionProvider } from "next-auth/react";
import { useUser } from "@/contexts/UserContext";

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const pathname = usePathname();
  const title: string = sidebarLinks.filter(
    (item) => item.route === pathname
  )[0].title;
  const { isDarkMode } = useUser();
  return (
    <main
      className={`flex h-screen font-inter ${
        isDarkMode ? "bg-gray-700" : "bg-[#F5F7FA]"
      } `}
    >
      <Sidebar />
      <div className={`w-full flex-grow flex flex-col overflow-x-hidden `}>
        <Header title={title} />

        <div className="flex size-full flex-col">
          <div
            className={`${
              isDarkMode ? "bg-gray-800" : "bg-[#F5F7FA]"
            } flex h-16 items-center justify-between p-5 sm:p-8 md:hidden`}
          >
            <div>
              <MobileSidebar />
            </div>
            <h1 className="text-[#343C6A] text-xl font-bold">{title}</h1>
            <Avatar>
              <AvatarImage src="https://github.com/shadcn.png" />
              <AvatarFallback>CN</AvatarFallback>
            </Avatar>
          </div>
          <div
            className={`justify-center items-center flex ${
              isDarkMode ? "bg-gray-800" : "bg-[#F5F7FA]"
            }`}
          >
            <div
              className={`justify-center items-center flex gap-3${
                isDarkMode ? "bg-gray-700" : "bg-[#F5F7FA]"
              } p-3 rounded-full md:hidden mb-4 md:mb-0`}
            >
              <Image
                src="/icons/Search.svg"
                width={20}
                height={20}
                alt="Search"
              />
              <input
                className="outline-none bg-[#F5F7FA]"
                type="text"
                placeholder="Search for something"
              />
            </div>
          </div>

          {children}
        </div>
      </div>
    </main>
  );
}
