"use client";
import React from "react";
import { sidebarLinks } from "@/constants";
import Link from "next/link";
import Image from "next/image";
import { usePathname } from "next/navigation";
import { cn } from "@/lib/utils";
import { useUser } from "@/contexts/UserContext";

const Sidebar = () => {
  const pathname = usePathname();
  const { isDarkMode } = useUser();
  return (
    <div
      className={cn(
        "sticky left-0 top-0 h-screen border-r pt-4 text-white max-md:hidden sm:p-2 xl:p-4 2xl:w-[300px]",
        {
          " bg-white": !isDarkMode,
          "border-gray-700 bg-gray-800": isDarkMode,
        }
      )}
    >
      <div className="flex items-center gap-2 p-3 pb-8">
        <Image src="/icons/logo.png" width={25} height={25} alt="logo" />
        <h1
          className={cn("font-[900] text-[1.5rem]", {
            "text-primaryBlack": !isDarkMode,
            "text-white": isDarkMode,
          })}
        >
          BankDash.
        </h1>
      </div>
      <div className="flex flex-col gap-2">
        {sidebarLinks.map((link, index) => {
          const isActive =
            pathname === link.route ||
            pathname.startsWith(`dashboard/${link.route}/`);
          return (
            <Link
              href={link.route}
              key={link.title}
              className={cn(
                "flex gap-6 items-center py-1 md:p-3 2xl:px-4 pl-0 justify-center xl:justify-start",
                {
                  "border-l-4 bg-nav-focus border-orange-1 border-primaryBlue":
                    isActive && !isDarkMode,
                  "border-l-4 bg-gray-800 border-primaryBlue":
                    isActive && isDarkMode,
                }
              )}
            >
              <Image
                src={link.icon}
                alt={link.title}
                width={20}
                height={20}
                className={cn({
                  "filter-custom-blue": isActive,
                })}
              />
              <p
                className={cn("text-sm font-semibold", {
                  "text-[#B1B1B1]": !isDarkMode && !isActive,
                  "text-gray-400": isDarkMode && !isActive,
                  "text-primaryBlue": isActive,
                })}
              >
                {link.title}
              </p>
            </Link>
          );
        })}
      </div>
    </div>
  );
};

export default Sidebar;
