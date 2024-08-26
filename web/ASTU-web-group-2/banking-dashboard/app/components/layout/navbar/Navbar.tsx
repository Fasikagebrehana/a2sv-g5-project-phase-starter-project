import React, { useEffect, useState, useRef } from "react";
import Link from "next/link";
import { usePathname } from "next/navigation";
import { useSelector } from "react-redux";
import { RootState } from "@/lib/store";
import { signOut, useSession } from "next-auth/react";
import Image from "next/image";
import Notification from "./Notification";
import { useGetAllTransactionQuery } from "@/lib/service/TransactionService";
import { Item } from "../../lastTransaction/lastTransactionItems";

interface NavbarProps {
  setter: React.Dispatch<React.SetStateAction<boolean>>;
}

const Navbar: React.FC<NavbarProps> = ({ setter }) => {
  // Access the user data from Redux store
  const user = useSelector((state: RootState) => state.user.user);

  // State for sidebar visibility
  const [isSidebarOpen, setIsSidebarOpen] = useState(false);

  // State for dropdown visibility
  const [isDropdownOpen, setIsDropdownOpen] = useState(false);

  // state for notification visibility
  const [isNotificationOpen, setIsNotificationOpen] = useState(false);

  // Reference for the dropdown
  const dropdownRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (
        dropdownRef.current &&
        !dropdownRef.current.contains(event.target as Node)
      ) {
        setIsDropdownOpen(false);
        setIsNotificationOpen(false);
      }
    };

    document.addEventListener("mousedown", handleClickOutside);

    // Cleanup the event listener
    return () => {
      document.removeEventListener("mousedown", handleClickOutside);
    };
  }, [dropdownRef]);

  // Toggle dropdown visibility
  const toggleDropdown = () => {
    setIsDropdownOpen((prev) => !prev);
  };

  // toggle notification dropwon visibility
  const toggleNotification = () => {
    setIsNotificationOpen((prev) => !prev);
  };

  // Get the current pathname from next/navigation
  const pathname = usePathname();

  // Function to capitalize the first letter of the page title
  const capitalizeFirstLetter = (text: string) => {
    if (!text || text === "/") return "Dashboard";
    text = text.replace("-", " ");
    return text.charAt(1).toUpperCase() + text.slice(2).toLowerCase();
  };
  // const notifications = [
  //   { type: "New", description: "Transaction" },
  //   { type: "New", description: "Transaction" },
  //   { type: "New", description: "Transaction" },
  //   { type: "New", description: "Transaction" },
  // ];

  let access: string = "";
  const { data: session, status } = useSession();
  useEffect(() => {}, [status, session]);
  if (session) {
    access = session?.user?.accessToken;
  }

  const { data, isError, isLoading } = useGetAllTransactionQuery(access);

  if (isLoading) {
    return <p>loading...</p>;
  }
  if (isError) {
    return <p>Error</p>;
  }

  let transactions: { content: Item[]; totalPages: number } = data?.data || "";

  const notifications = transactions.content;

  const title = capitalizeFirstLetter(pathname);
  return (
    <nav className="fixed top-0 md:left-[240px] max-md:left-0 right-0 bg-white shadow-md z-10 max-md:h-[140px]">
      <div className="flex w-full items-center justify-between px-4 py-2">
        <div className="flex items-center">
          <Image
            src="/assets/navbar/hamburger.svg"
            width={25}
            height={25}
            alt="hamburger"
            className="lg:hidden md:hidden flex  cursor-pointer"
            onClick={() => {
              setter((prev) => !prev);
            }}
          />
        </div>
        <div className="flex gap-2 w-full items-center justify-between px-4 py-2">
          <div className="flex-grow flex justify-start max-md:justify-center">
            <p
              data-id="title"
              className="font-semibold text-[20px] sm:text-[25px] text-[#343C6A]"
            >
              {title.slice(0)}
            </p>
          </div>

          <div className="sm:flex gap-5 items-center relative">
            <div className="hidden sm:flex search-div bg-[#F5F7FA]  items-center rounded-full h-[50px] px-4 py-2 w-full max-w-[400px] mx-auto">
              <Image
                src="/assets/navbar/magnifying-glass.svg"
                width={20}
                height={20}
                alt="magnifying-glass"
                className="mr-3 "
              />
              <input
                type="text"
                placeholder="Search for something"
                className="text-[15px] bg-[#F5F7FA] border border-transparent rounded-full p-2 w-full placeholder:text-[#B1B1B1] focus:outline-none focus:ring-0 focus:border-transparent"
              />
            </div>

            <Link
              href="/settings"
              className="sm:flex hidden bg-[#F5F7FA] rounded-full justify-center items-center"
            >
              <Image
                src="/assets/navbar/settings.svg"
                width={50}
                height={50}
                alt="settings"
                className="flex-shrink-0 min-w-fit"
              />
            </Link>
            <div
              className=" sm:flex hidden bg-[#F5F7FA] rounded-full  justify-center items-center"
              ref={dropdownRef}
            >
              <div className="relative">
                <Image
                  src="/assets/navbar/notification.svg"
                  width={50}
                  height={50}
                  alt="notification"
                  className="flex-shrink-0 min-w-fit"
                  onClick={toggleNotification}
                />
                {notifications.length > 0 && (
                  <span className="absolute top-0 right-0 flex items-center justify-center w-4 h-4 text-white bg-red-600 rounded-full text-xs">
                    {notifications.length}
                  </span>
                )}
                {isNotificationOpen && (
                  <Notification notifications={notifications} />
                )}
              </div>
            </div>

            <div className="relative" ref={dropdownRef}>
              <div className="flex justify-center items-center w-12 h-12 rounded-full object-scale-down">
                <Image
                  src={
                    user?.profilePicture === "assets/navbar/default-image.svg"
                      ? ""
                      : user?.profilePicture ||
                        "assets/navbar/default-image.svg"
                  }
                  width={50}
                  height={50}
                  alt="profile-picture"
                  className="block mx-auto h-12 rounded-full cursor-pointer"
                  onClick={toggleDropdown}
                  data-id="Profile-image"
                />
              </div>

              {isDropdownOpen && (
                <div className="absolute right-0 mt-2 w-48 rounded-md shadow-lg py-2 z-50 bg-white">
                  <div className="flex flex-col gap-1">
                    <Link
                      href="/settings"
                      className="md:hidden  flex bg-[#F5F7FA] rounded-full gap-2 items-center"
                    >
                      <Image
                        src="/assets/navbar/settings.svg"
                        width={30}
                        height={30}
                        alt="settings"
                        className="flex-shrink-0"
                      />
                      <span className="ml-2 text-sm">Setting</span>
                    </Link>

                    <Link
                      href="/"
                      className="md:hidden flex bg-[#F5F7FA] rounded-full gap-2 items-center"
                    >
                      <Image
                        src="/assets/navbar/notification.svg"
                        width={30}
                        height={30}
                        alt="notification"
                        className="flex-shrink-0"
                      />
                      <span className="ml-2 text-sm">Notification</span>
                    </Link>
                    <button
                      onClick={() => signOut({ callbackUrl: "/" })}
                      className="flex bg-[#F5F7FA] rounded-full gap-4 items-center pl-2"
                    >
                      <Image
                        src="/assets/transaction/deposit.svg"
                        width={16}
                        height={16}
                        alt="Logout"
                        className="flex-shrink-0 rotate-90"
                      />
                      <span className="ml-2 text-sm">Logout</span>
                    </button>
                  </div>
                </div>
              )}
            </div>
          </div>
        </div>
      </div>

      <div className="search-div  flex w-11/12 sm:hidden bg-[#F5F7FA] justify-between items-center rounded-full pl-5 pr-5 mt-2 mx-auto h-10">
        <Image
          src="/assets/navbar/magnifying-glass.svg"
          width={20}
          height={20}
          alt="magnifying-glass"
          className="mr-5"
        />
        <input
          type="text"
          placeholder="Search for something"
          className="text-md bg-[#F5F7FA] w-full outline-none"
        />
      </div>
    </nav>
  );
};

export default Navbar;
