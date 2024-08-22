"use client";
import React, { useState } from "react";
import { useRouter } from "next/navigation";
import BankDashlogo from "../../assets/Sidebar/BankDashlogo.svg";
import Image from "next/image";
import Link from "next/link";
import { usePathname } from "next/navigation";

const Sidebar = () => {
  const [isOpen, setIsOpen] = useState(false);
  const pathname = usePathname();
  const isActive = (path : string) => pathname === path;

  const toggleSidebar = () => {
    setIsOpen(!isOpen);
  };

  return (
    <>
      {/* Burger Menu */}
      <div className="mobile:hidden p-4 text-Dark-Slate-Blue h-[60px] fixed top-2">
        <button onClick={toggleSidebar} className="text-xl">
          ☰
        </button>
      </div>

      {/* Sidebar */}
      <div
        className={`flex-grow w-[250px] pr-1 border border-l-Very-Pale-Blue  h-full max-mobile:fixed transition-transform transform ${
          isOpen ? "translate-x-0" : "-translate-x-full"
        } mobile:translate-x-0`}
      >
        <ul className="">
          <li className="flex items-center w-[189px] h-[101px] ml-[38px]">
            <Image className="mr-2" src={BankDashlogo} alt="BankDash Logo" />{" "}
            <p className="text-2xl font-black text-[#343C6A]">BankDash.</p>
          </li>

          {/* Dashboard */}
          <li className=" flex items-centermt-3 w-[189px] h-[60px]">
            {/* active indicator */}
            <svg
              // width="5"
              // height="inherit"
              viewBox="0 0 5 50"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                d="M0 0C2.76142 0 5 2.23858 5 5V45C5 47.7614 2.76142 50 0 50V0Z"
                fill={isActive("/") ? "#2D60FF" : "white"}
              />
            </svg>

            <Link href="/" className=" ml-9 flex items-center w-full ">
              <svg
                width="25"
                height="25"
                viewBox="0 0 20 20"
                fill=""
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  d="M19.4602 8.69904C19.4598 8.69858 19.4593 8.69812 19.4588 8.69766L11.3004 0.539551C10.9527 0.19165 10.4903 0 9.99855 0C9.50676 0 9.04442 0.191498 8.69652 0.539398L0.542379 8.69339C0.539632 8.69614 0.536886 8.69904 0.534139 8.70178C-0.179972 9.42001 -0.178751 10.5853 0.537649 11.3017C0.86495 11.6292 1.29723 11.8188 1.75942 11.8387C1.77819 11.8405 1.79711 11.8414 1.81618 11.8414H2.14135V17.8453C2.14135 19.0334 3.10799 20 4.29635 20H7.48818C7.81166 20 8.07412 19.7377 8.07412 19.4141V14.707C8.07412 14.1649 8.51509 13.7239 9.05724 13.7239H10.9399C11.482 13.7239 11.923 14.1649 11.923 14.707V19.4141C11.923 19.7377 12.1853 20 12.5089 20H15.7008C16.8891 20 17.8558 19.0334 17.8558 17.8453V11.8414H18.1573C18.6489 11.8414 19.1113 11.6499 19.4593 11.302C20.1765 10.5844 20.1768 9.41711 19.4602 8.69904Z"
                  fill={isActive("/") ? "#2D60FF" : "#B1B1B1"}
                />
              </svg>
              <p
                className={`ml-6 text-lg font-medium leading-[21.7px] ${
                  isActive("/") ? "text-[#2D60FF]" : "text-[#B1B1B1]"
                }`}
              >
                Dashboard
              </p>
            </Link>
          </li>

          {/* Transactions */}
          <li className=" flex items-centermt-3 w-[189px] h-[60px]">
            {/* active indicator */}
            <svg
              // width="5"
              // height="inherit"
              viewBox="0 0 5 50"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                d="M0 0C2.76142 0 5 2.23858 5 5V45C5 47.7614 2.76142 50 0 50V0Z"
                fill={isActive("/transactions") ? "#2D60FF" : "white"}
              />
            </svg>

            <Link href="/" className=" ml-9 flex items-center w-full ">
              <svg
                width="20"
                height="20"
                viewBox="0 0 20 20"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  d="M4.16663 18.3334C4.16715 18.7753 4.34291 19.1989 4.65536 19.5113C4.96781 19.8238 5.39143 19.9996 5.83329 20.0001H14.1666C14.6085 19.9996 15.0321 19.8238 15.3446 19.5113C15.657 19.1989 15.8328 18.7753 15.8333 18.3334V17.6042H4.16663V18.3334Z"
                  fill={isActive("/transactions") ? "#2D60FF" : "#B1B1B1"}
                />
                <path
                  d="M15.8333 1.66667C15.8328 1.2248 15.657 0.801181 15.3446 0.488734C15.0321 0.176287 14.6085 0.000523784 14.1666 0L5.83329 0C5.39143 0.000523784 4.96781 0.176287 4.65536 0.488734C4.34291 0.801181 4.16715 1.2248 4.16663 1.66667V2.5H15.8333V1.66667Z"
                  fill={isActive("/transactions") ? "#2D60FF" : "#B1B1B1"}
                />
                <path
                  d="M19.7678 5.36052L17.2678 2.75635L16.0655 3.91051L17.1114 5.0001H15.8333V6.66677H17.2013L16.0901 7.73177L17.2432 8.9351L19.7432 6.53927C19.8223 6.46351 19.8856 6.37291 19.9296 6.27266C19.9737 6.1724 19.9975 6.06445 19.9998 5.95498C20.0021 5.84551 19.9827 5.73667 19.9429 5.63467C19.9031 5.53266 19.8436 5.4395 19.7678 5.36052Z"
                  fill={isActive("/transactions") ? "#2D60FF" : "#B1B1B1"}
                />
                <path
                  d="M13.3333 4.99992H15.8333V3.33325H4.16663V13.3333H6.66663V14.9999H4.16663V16.6666H15.8333V6.66659H13.3333V4.99992ZM12.5 8.33325H9.58329C9.47279 8.33325 9.36681 8.37715 9.28866 8.45529C9.21052 8.53343 9.16663 8.63941 9.16663 8.74992C9.16663 8.86043 9.21052 8.96641 9.28866 9.04455C9.36681 9.12269 9.47279 9.16659 9.58329 9.16659H10.4166C10.9333 9.16607 11.4317 9.35754 11.8151 9.70383C12.1985 10.0501 12.4396 10.5265 12.4915 11.0405C12.5434 11.5546 12.4025 12.0696 12.0961 12.4855C11.7897 12.9015 11.3396 13.1888 10.8333 13.2916V14.1666H9.16663V13.3333H7.49996V11.6666H10.4166C10.5271 11.6666 10.6331 11.6227 10.7113 11.5445C10.7894 11.4664 10.8333 11.3604 10.8333 11.2499C10.8333 11.1394 10.7894 11.0334 10.7113 10.9553C10.6331 10.8772 10.5271 10.8333 10.4166 10.8333H9.58329C9.06665 10.8338 8.56824 10.6423 8.18482 10.296C7.8014 9.94973 7.56034 9.47332 7.50841 8.95929C7.45649 8.44526 7.59742 7.93027 7.90384 7.51431C8.21026 7.09834 8.66031 6.81106 9.16663 6.70825V5.83325H10.8333V6.66659H12.5V8.33325Z"
                  fill={isActive("/transactions") ? "#2D60FF" : "#B1B1B1"}
                />
                <path
                  d="M2.79862 13.3333L3.90987 12.2683L2.75674 11.0649L0.256745 13.4608C0.177692 13.5365 0.114344 13.6271 0.0703209 13.7274C0.0262982 13.8276 0.00246423 13.9356 0.000181182 14.0451C-0.00210186 14.1545 0.0172108 14.2634 0.0570153 14.3654C0.0968199 14.4674 0.156336 14.5605 0.232162 14.6395L2.73216 17.2437L3.93445 16.0895L2.88852 14.9999H4.16664V13.3333H2.79862Z"
                  fill={isActive("/transactions") ? "#2D60FF" : "#B1B1B1"}
                />
              </svg>

              <p
                className={`ml-6 text-lg font-medium leading-[21.7px] ${
                  isActive("/transactions")
                    ? "text-[#2D60FF]"
                    : "text-[#B1B1B1]"
                }`}
              >
                Transactions
              </p>
            </Link>
          </li>

          {/* Accounts */}

          <li className=" flex items-centermt-3 w-[189px] h-[60px]">
            {/* active indicator */}
            <svg
              // width="5"
              // height="inherit"
              viewBox="0 0 5 50"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                d="M0 0C2.76142 0 5 2.23858 5 5V45C5 47.7614 2.76142 50 0 50V0Z"
                fill={isActive("/accounts") ? "#2D60FF" : "white"}
              />
            </svg>

            <Link href="/" className=" ml-9 flex items-center w-full ">
              <svg
                width="20"
                height="20"
                viewBox="0 0 20 20"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
              >
                <g clipPath="url(#clip0_147_69)">
                  <path
                    d="M9.85748 9.63406C11.181 9.63406 12.3271 9.15936 13.2635 8.22278C14.2 7.28635 14.6747 6.14057 14.6747 4.81688C14.6747 3.49364 14.2 2.34771 13.2634 1.41097C12.3268 0.474699 11.1809 0 9.85748 0C8.53378 0 7.388 0.474699 6.45157 1.41113C5.51514 2.34756 5.04028 3.49349 5.04028 4.81688C5.04028 6.14057 5.51514 7.2865 6.45172 8.22293C7.38831 9.1592 8.53424 9.63406 9.85748 9.63406Z"
                    fill={isActive("/transactions") ? "#2D60FF" : "#B1B1B1"}
                  />
                  <path
                    d="M18.2863 15.3789C18.2593 14.9891 18.2047 14.564 18.1242 14.1151C18.0431 13.6629 17.9385 13.2353 17.8134 12.8445C17.6842 12.4406 17.5084 12.0418 17.2911 11.6595C17.0656 11.2628 16.8007 10.9174 16.5034 10.6331C16.1926 10.3357 15.8121 10.0966 15.372 9.92218C14.9335 9.74869 14.4475 9.6608 13.9276 9.6608C13.7234 9.6608 13.526 9.74457 13.1447 9.99283C12.91 10.1459 12.6355 10.3229 12.3291 10.5186C12.0671 10.6856 11.7122 10.842 11.2738 10.9836C10.8461 11.122 10.4118 11.1922 9.98322 11.1922C9.5546 11.1922 9.12048 11.122 8.69232 10.9836C8.25439 10.8421 7.89948 10.6857 7.63779 10.5188C7.33429 10.3249 7.05963 10.1479 6.82144 9.99267C6.44058 9.74441 6.24298 9.66064 6.03882 9.66064C5.5188 9.66064 5.03296 9.74869 4.59457 9.92233C4.15482 10.0964 3.77411 10.3355 3.46298 10.6332C3.16589 10.9177 2.90085 11.263 2.67563 11.6595C2.4585 12.0418 2.28271 12.4405 2.15332 12.8447C2.02835 13.2355 1.92383 13.6629 1.84265 14.1151C1.76224 14.5634 1.70761 14.9887 1.6806 15.3793C1.65405 15.762 1.64062 16.1592 1.64062 16.5603C1.64062 17.6043 1.9725 18.4495 2.62695 19.0728C3.27332 19.6879 4.12857 20 5.16861 20H14.7987C15.8388 20 16.6937 19.6881 17.3402 19.0728C17.9948 18.45 18.3267 17.6046 18.3267 16.5602C18.3266 16.1572 18.313 15.7597 18.2863 15.3789Z"
                    fill={isActive("/transactions") ? "#2D60FF" : "#B1B1B1"}
                  />
                </g>
                <defs>
                  <clipPath id="clip0_147_69">
                    <rect width="20" height="20" fill="white" />
                  </clipPath>
                </defs>
              </svg>
              <p
                className={`ml-6 text-lg font-medium leading-[21.7px] ${
                  isActive("/accounts") ? "text-[#2D60FF]" : "text-[#B1B1B1]"
                }`}
              >
                Accounts
              </p>
            </Link>
          </li>

          {/* Investements */}

          <li className=" flex items-centermt-3 w-full h-[60px]">
            {/* active indicator */}
            <svg
              // width="5"
              // height="inherit"
              viewBox="0 0 5 50"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                d="M0 0C2.76142 0 5 2.23858 5 5V45C5 47.7614 2.76142 50 0 50V0Z"
                fill={isActive("/investements") ? "#2D60FF" : "white"}
              />
            </svg>

            <Link href="/" className=" ml-9 flex items-center w-full ">
              <svg
                width="20"
                height="20"
                viewBox="0 0 20 20"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  d="M3.1366 8.18164H0.909362C0.407234 8.18164 0 8.58888 0 9.091V19.091C0 19.5927 0.407234 19.9999 0.909362 19.9999H3.1366C3.63872 19.9999 4.04553 19.5927 4.04553 19.091V9.091C4.04553 8.58888 3.63872 8.18164 3.1366 8.18164Z"
                  fill={isActive("/investements") ? "#2D60FF" : "#B1B1B1"}
                />
                <path
                  d="M8.45489 10.9084H6.22766C5.72511 10.9084 5.3183 11.3157 5.3183 11.8178V19.0901C5.3183 19.5927 5.72511 19.9995 6.22766 19.9995H8.45489C8.95702 19.9995 9.36383 19.5923 9.36383 19.0901V11.8178C9.36383 11.3157 8.95702 10.9084 8.45489 10.9084Z"
                  fill={isActive("/investements") ? "#2D60FF" : "#B1B1B1"}
                />
                <path
                  d="M13.7723 10.9084H11.5451C11.043 10.9084 10.6362 11.3157 10.6362 11.8178V19.0901C10.6362 19.5927 11.043 19.9995 11.5451 19.9995H13.7723C14.2749 19.9995 14.6817 19.5923 14.6817 19.0901V11.8178C14.6817 11.3157 14.2749 10.9084 13.7723 10.9084Z"
                  fill={isActive("/investements") ? "#2D60FF" : "#B1B1B1"}
                />
                <path
                  d="M19.0906 8.18164H16.8634C16.3613 8.18164 15.9545 8.58888 15.9545 9.091V19.091C15.9545 19.5931 16.3613 19.9999 16.8634 19.9999H19.0906C19.5928 19.9999 20 19.5923 20 19.091V9.091C20 8.58888 19.5928 8.18164 19.0906 8.18164Z"
                  fill={isActive("/investements") ? "#2D60FF" : "#B1B1B1"}
                />
                <path
                  d="M10.2651 5.47656V7.11316C10.7605 7.08039 11.2826 6.84805 11.2826 6.30295C11.2826 5.74039 10.7102 5.57571 10.2651 5.47656Z"
                  fill={isActive("/investements") ? "#2D60FF" : "#B1B1B1"}
                />
                <path
                  d="M8.84167 3.51786C8.84167 3.93147 9.14933 4.17062 9.76933 4.29488V2.81445C9.20593 2.83105 8.84167 3.16211 8.84167 3.51786Z"
                  fill={isActive("/investements") ? "#2D60FF" : "#B1B1B1"}
                />
                <path
                  d="M10 0C7.24298 0 5 2.2434 5 5C5 7.75575 7.24298 9.99915 10 9.99915C12.757 9.99915 15 7.75575 15 5C15 2.2434 12.757 0 10 0ZM10.2651 7.95702V8.46894C10.2651 8.60979 10.157 8.75021 10.0157 8.75021C9.87617 8.75021 9.76936 8.60979 9.76936 8.46894V7.95702C8.37234 7.92298 7.67745 7.08809 7.67745 6.43489C7.67745 6.10511 7.87702 5.91447 8.18936 5.91447C9.11532 5.91447 8.39532 7.05532 9.76936 7.11277V5.38511C8.54383 5.16255 7.8017 4.62511 7.8017 3.70766C7.8017 2.58383 8.73617 2.00426 9.76936 1.97192V1.53106C9.76936 1.39021 9.87617 1.24979 10.0157 1.24979C10.157 1.24979 10.2651 1.39021 10.2651 1.53106V1.97192C10.9094 1.98894 12.2323 2.39319 12.2323 3.20383C12.2323 3.52596 11.9915 3.71575 11.7102 3.71575C11.1723 3.71575 11.18 2.83192 10.2651 2.81489V4.38553C11.3562 4.61745 12.3226 4.93958 12.3226 6.21277C12.3226 7.32 11.4962 7.8817 10.2651 7.95702Z"
                  fill={isActive("/investements") ? "#2D60FF" : "#B1B1B1"}
                />
              </svg>
              <p
                className={`ml-6 text-lg font-medium leading-[21.7px] ${
                  isActive("/investements")
                    ? "text-[#2D60FF]"
                    : "text-[#B1B1B1]"
                }`}
              >
                Investements
              </p>
            </Link>
          </li>

          {/* credit cards */}

          <li className=" flex items-centermt-3 w-full h-[60px]">
            {/* active indicator */}
            <svg
              // width="5"
              // height="inherit"
              viewBox="0 0 5 50"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                d="M0 0C2.76142 0 5 2.23858 5 5V45C5 47.7614 2.76142 50 0 50V0Z"
                fill={isActive("/creditcards") ? "#2D60FF" : "white"}
              />
            </svg>

            <Link href="/" className=" ml-9 flex items-center w-full ">
              <svg
                width="20"
                height="16"
                viewBox="0 0 20 16"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  d="M17.6621 3.51034V3.23675C17.6621 1.97905 16.6389 0.955811 15.3812 0.955811H2.28094C1.0232 0.95585 0 1.97905 0 3.23675V3.51034H17.6621Z"
                  fill={isActive("/creditcards") ? "#2D60FF" : "#B1B1B1"}
                />
                <path
                  d="M10.4347 10.8476C10.4347 9.6644 10.8152 8.53983 11.5178 7.61304H0V10.7948C0 12.0525 1.0232 13.0757 2.28094 13.0757H10.9162C10.601 12.3852 10.4347 11.6292 10.4347 10.8476ZM8.83109 10.1562H6.8625V8.98437H8.83109V10.1562ZM2.62727 8.98437H5.69062V10.1562H2.62727V8.98437Z"
                  fill={isActive("/creditcards") ? "#2D60FF" : "#B1B1B1"}
                />
                <path
                  d="M12.7344 6.44124C13.6281 5.81597 14.6898 5.47909 15.8033 5.47909C16.4478 5.47909 17.0748 5.59226 17.6621 5.80855V4.68237H0V6.44124H12.7344Z"
                  fill={isActive("/creditcards") ? "#2D60FF" : "#B1B1B1"}
                />
                <path
                  d="M20 10.8476C20 8.52979 18.1211 6.65088 15.8033 6.65088C13.4855 6.65088 11.6066 8.52979 11.6066 10.8476C11.6066 13.1653 13.4855 15.0442 15.8033 15.0442C18.1211 15.0442 20 13.1653 20 10.8476ZM16.369 13.0896V13.5665H15.7831V13.5665V13.5665H15.1971V13.0931C14.8429 12.9731 14.553 12.7621 14.27 12.5551L14.9618 11.6092C15.342 11.8873 15.5366 12.021 15.8033 12.021C15.9541 12.021 16.0761 11.9492 16.1219 11.8337C16.1773 11.6938 16.099 11.565 15.9124 11.489C15.9124 11.489 15.0748 11.2097 14.6876 10.8149C14.3627 10.4837 14.2594 10.0169 14.3628 9.57291C14.467 9.12584 14.7688 8.77627 15.1971 8.6017V8.12858H16.369V8.58209C16.6666 8.66475 16.9185 8.7869 17.0634 8.86553L16.5043 9.89541C16.1336 9.6942 15.7925 9.63381 15.6598 9.6792C15.531 9.72318 15.5114 9.80717 15.5041 9.83881C15.4936 9.88361 15.4881 9.95237 15.5597 10.0322C15.6286 10.109 16.3542 10.4037 16.3542 10.4037C17.1371 10.7224 17.5055 11.5228 17.2113 12.2654C17.0585 12.6514 16.7508 12.9437 16.369 13.0896Z"
                  fill={isActive("/creditcards") ? "#2D60FF" : "#B1B1B1"}
                />
              </svg>

              <p
                className={`ml-6 text-lg font-medium leading-[21.7px] ${
                  isActive("/creditcards") ? "text-[#2D60FF]" : "text-[#B1B1B1]"
                }`}
              >
                Credit Cards
              </p>
            </Link>
          </li>

          {/* loans */}

          <li className=" flex items-centermt-3 w-[189px] h-[60px]">
            {/* active indicator */}
            <svg
              // width="5"
              // height="inherit"
              viewBox="0 0 5 50"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                d="M0 0C2.76142 0 5 2.23858 5 5V45C5 47.7614 2.76142 50 0 50V0Z"
                fill={isActive("/loans") ? "#2D60FF" : "white"}
              />
            </svg>

            <Link href="/loans" className=" ml-9 flex items-center w-full ">
              <svg
                width="20"
                height="20"
                viewBox="0 0 20 20"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  d="M11.5939 12.2698C14.9499 12.2698 17.6801 9.51765 17.6801 6.1349C17.6801 2.75215 14.9499 0 11.5939 0C8.23792 0 5.50763 2.75211 5.50763 6.13486C5.50763 9.51761 8.23792 12.2698 11.5939 12.2698ZM9.63818 8.04927C9.81522 7.77865 10.1782 7.70268 10.4488 7.8798C10.8432 8.13773 10.9919 8.16151 11.5123 8.15788C12.0202 8.15452 12.3149 7.77595 12.3737 7.42562C12.4024 7.25521 12.4134 6.83909 11.8977 6.65678C11.2927 6.44292 10.6737 6.1985 10.243 5.86069C9.81229 5.52288 9.61507 4.93971 9.72832 4.33886C9.8511 3.68749 10.3055 3.16897 10.9142 2.98564C10.9197 2.984 10.9251 2.98268 10.9306 2.98104V2.75906C10.9306 2.43566 11.1928 2.17347 11.5162 2.17347C11.8396 2.17347 12.1017 2.43566 12.1017 2.75906V2.94414C12.4995 3.03909 12.7772 3.22105 12.89 3.30534C13.1491 3.49901 13.2021 3.86594 13.0084 4.125C12.8148 4.38407 12.4479 4.43708 12.1888 4.24337C12.0688 4.15366 11.7372 3.9608 11.2519 4.10701C10.9685 4.19242 10.895 4.4721 10.8792 4.55576C10.8482 4.72016 10.883 4.87425 10.9657 4.93913C11.2643 5.17329 11.8017 5.38062 12.288 5.55251C13.1847 5.86947 13.6833 6.70023 13.5287 7.61976C13.4529 8.07094 13.2259 8.48951 12.8896 8.79851C12.6605 9.00901 12.3938 9.15908 12.1017 9.24477V9.51066C12.1017 9.83406 11.8396 10.0963 11.5162 10.0963C11.1928 10.0963 10.9306 9.83406 10.9306 9.51066V9.30317C10.5519 9.25726 10.2341 9.13885 9.80761 8.85988C9.53703 8.68284 9.46113 8.31989 9.63818 8.04927Z"
                  fill={isActive("/loans") ? "#2D60FF" : "#B1B1B1"}
                />
                <path
                  d="M2.21958 14.2373H0.884356C0.560955 14.2373 0.298767 14.4995 0.298767 14.8229V19.4135C0.298767 19.7369 0.560955 19.9991 0.884356 19.9991H2.21961V14.2373H2.21958Z"
                  fill={isActive("/loans") ? "#2D60FF" : "#B1B1B1"}
                />
                <path
                  d="M19.5295 14.1965C18.432 13.0989 16.646 13.0989 15.5485 14.1965L13.7943 15.9507L13.0754 16.6697C12.7848 16.9602 12.3906 17.1235 11.9797 17.1235H8.48359C8.16784 17.1235 7.89613 16.8808 7.88126 16.5654C7.86541 16.2287 8.13372 15.9507 8.467 15.9507H12.0206C12.7351 15.9507 13.3548 15.442 13.4776 14.7382C13.5058 14.5765 13.5206 14.4104 13.5206 14.2408C13.5206 13.9168 13.2581 13.6539 12.9341 13.6539H10.987C10.3506 13.6539 9.73956 13.3652 9.09257 13.0596C8.41395 12.739 7.71225 12.4075 6.89177 12.3529C6.17415 12.3051 5.45489 12.3837 4.75386 12.5861C4.00325 12.8029 3.46369 13.4697 3.39826 14.2397C3.39576 14.2395 3.39322 14.2394 3.39069 14.2393V19.9972L13.4795 20C14.1731 20 14.8253 19.7298 15.3159 19.2393L19.5294 15.0258C19.7585 14.7969 19.7585 14.4255 19.5295 14.1965Z"
                  fill={isActive("/loans") ? "#2D60FF" : "#B1B1B1"}
                />
              </svg>

              <p
                className={`ml-6 text-lg font-medium leading-[21.7px] ${
                  isActive("/loans") ? "text-[#2D60FF]" : "text-[#B1B1B1]"
                }`}
              >
                Loans
              </p>
            </Link>
          </li>

          {/* services */}

          <li className=" flex items-centermt-3 w-[189px] h-[60px]">
            {/* active indicator */}
            <svg
              // width="5"
              // height="inherit"
              viewBox="0 0 5 50"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                d="M0 0C2.76142 0 5 2.23858 5 5V45C5 47.7614 2.76142 50 0 50V0Z"
                fill={isActive("/services") ? "#2D60FF" : "white"}
              />
            </svg>

            <Link href="/" className=" ml-9 flex items-center w-full ">
              <svg
                width="20"
                height="20"
                viewBox="0 0 20 20"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
              >
                <g clipPath="url(#clip0_147_10)">
                  <path
                    d="M19.8283 1.00036L18.9997 0.17173C18.8269 -0.00108164 18.5651 -0.0484253 18.343 0.0523557L13.8924 2.01708C13.7172 2.09677 13.5926 2.25704 13.5582 2.44594C13.5242 2.63536 13.5845 2.82911 13.7204 2.96508L17.0349 6.27957C17.1709 6.41555 17.3647 6.47582 17.5541 6.4418C17.743 6.40746 17.9033 6.28281 17.9829 6.10762L19.9476 1.657C20.0484 1.43485 20.0011 1.1731 19.8283 1.00036Z"
                    fill={isActive("/services") ? "#2D60FF" : "#B1B1B1"}
                  />
                  <path
                    d="M3.78516 12.9004L0.514071 16.1715C-0.171357 16.8569 -0.171357 17.9719 0.514071 18.6573L1.34266 19.4859C2.02809 20.1714 3.14313 20.1714 3.82852 19.4859L7.09961 16.2148L3.78516 12.9004ZM3.41996 17.4169C3.19094 17.6459 2.82028 17.6459 2.59129 17.4169C2.36235 17.1878 2.36235 16.8172 2.59129 16.5883L4.20516 14.9744C4.43418 14.7453 4.8048 14.7453 5.03375 14.9744C5.26277 15.2034 5.26277 15.574 5.03375 15.803L3.41996 17.4169Z"
                    fill={isActive("/services") ? "#2D60FF" : "#B1B1B1"}
                  />
                  <path
                    d="M9.58545 13.7289L6.27097 10.4144C5.81413 9.95757 5.07054 9.95757 4.61374 10.4144C4.15691 10.8712 4.15691 11.6148 4.61374 12.0716L7.92823 15.3861C8.38499 15.843 9.12862 15.843 9.58545 15.3861C10.0423 14.9294 10.0423 14.1857 9.58545 13.7289Z"
                    fill={isActive("/services") ? "#2D60FF" : "#B1B1B1"}
                  />
                  <path
                    d="M15.3778 6.27958L13.7205 4.62231L9.17117 9.1716C8.71355 8.71399 7.97164 8.71399 7.51394 9.1716L7.09961 9.58594L10.4141 12.9004L10.8284 12.4861C11.286 12.0284 11.286 11.2865 10.8284 10.8289L15.3778 6.27958Z"
                    fill={isActive("/services") ? "#2D60FF" : "#B1B1B1"}
                  />
                  <path
                    d="M18.9718 13.9983C18.0206 13.0471 16.8768 12.8426 15.5643 13.0771L12.9005 10.4136L12.2644 11.0497C12.4738 11.8368 12.2702 12.701 11.6567 13.3145L11.2432 13.7279L13.0765 15.561C12.8916 16.5954 12.909 17.4727 13.5514 18.4243C14.2927 19.5227 15.5902 20.1471 16.9495 19.96C17.1742 19.9291 17.3619 19.7732 17.4341 19.5582C17.5063 19.3432 17.4506 19.1059 17.2903 18.9454L16.4845 18.141V16.484H18.1432L18.9446 17.2854C19.1053 17.4461 19.3433 17.5014 19.5585 17.4286C19.7737 17.3557 19.9291 17.1672 19.9593 16.942C20.1024 15.8762 19.7537 14.7799 18.9718 13.9983Z"
                    fill={isActive("/services") ? "#2D60FF" : "#B1B1B1"}
                  />
                  <path
                    d="M6.91557 4.42589C7.14811 3.12718 6.95982 1.98468 5.99436 1.01851C5.33744 0.36164 4.45854 0 3.51956 0C3.36147 0 3.20507 0.0103125 3.05069 0.0307421C2.82526 0.0606639 2.63655 0.216211 2.56366 0.43164C2.49077 0.646991 2.54643 0.885116 2.70722 1.0459L3.51643 1.84707V3.51558H1.85108L1.04788 2.70113C0.887453 2.5407 0.649836 2.48519 0.434759 2.55757C0.219759 2.62996 0.0642515 2.81785 0.0333532 3.04265C-0.145162 4.34253 0.418079 5.66311 1.56726 6.43889C2.52241 7.08241 3.40499 7.09835 4.43163 6.91327L6.27229 8.75694L6.68592 8.34331C7.29943 7.72979 8.16353 7.52624 8.95064 7.73561L9.5865 7.09975L6.91557 4.42589Z"
                    fill={isActive("/services") ? "#2D60FF" : "#B1B1B1"}
                  />
                </g>
                <defs>
                  <clipPath id="clip0_147_10">
                    <rect width="20" height="20" fill="white" />
                  </clipPath>
                </defs>
              </svg>

              <p
                className={`ml-6 text-lg font-medium leading-[21.7px] ${
                  isActive("/services") ? "text-[#2D60FF]" : "text-[#B1B1B1]"
                }`}
              >
                Services
              </p>
            </Link>
          </li>

          {/* setting */}

          <li className=" flex items-centermt-3 w-[189px] h-[60px]">
            {/* active indicator */}
            <svg
              // width="5"
              // height="inherit"
              viewBox="0 0 5 50"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                d="M0 0C2.76142 0 5 2.23858 5 5V45C5 47.7614 2.76142 50 0 50V0Z"
                fill={isActive("/settings") ? "#2D60FF" : "white"}
              />
            </svg>

            <Link href="/" className=" ml-9 flex items-center w-full ">
              <svg
                width="20"
                height="20"
                viewBox="0 0 20 20"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  d="M18.1588 7.53125H17.7342C17.5963 7.09961 17.4223 6.68031 17.2138 6.27746L17.5145 5.97676C18.2436 5.24844 18.2241 4.08125 17.5147 3.37266L16.6276 2.48555C15.9194 1.77645 14.752 1.7559 14.0235 2.48527L13.7225 2.78621C13.3197 2.57777 12.9004 2.40371 12.4688 2.26578V1.84113C12.4688 0.825937 11.6428 0 10.6276 0H9.37238C8.35719 0 7.53125 0.825937 7.53125 1.84113V2.26578C7.09965 2.40367 6.68031 2.57773 6.27746 2.78621L5.9768 2.48555C5.24973 1.75762 4.08234 1.77469 3.37273 2.48531L2.48551 3.37246C1.77645 4.08074 1.75594 5.24801 2.48527 5.97656L2.78621 6.2775C2.57773 6.68035 2.40371 7.09961 2.26578 7.53129H1.84117C0.825976 7.53125 0 8.35719 0 9.37238V10.6276C0 11.6428 0.825976 12.4688 1.84117 12.4688H2.26578C2.40371 12.9004 2.57773 13.3197 2.78621 13.7225L2.48551 14.0232C1.75641 14.7516 1.77594 15.9188 2.48527 16.6273L3.37242 17.5145C4.08059 18.2236 5.24801 18.2441 5.97652 17.5147L6.27746 17.2138C6.68031 17.4222 7.09965 17.5963 7.53125 17.7342V18.1589C7.53125 19.1741 8.35723 20 9.37242 20H10.6276C11.6428 20 12.4688 19.1741 12.4688 18.1589V17.7342C12.9004 17.5963 13.3197 17.4223 13.7226 17.2138L14.0232 17.5145C14.7503 18.2424 15.9177 18.2253 16.6273 17.5147L17.5145 16.6275C18.2236 15.9192 18.2441 14.752 17.5148 14.0234L17.2138 13.7225C17.4223 13.3196 17.5963 12.9004 17.7343 12.4687H18.1589C19.1741 12.4687 20 11.6427 20 10.6275V9.3723C20 8.35719 19.174 7.53125 18.1588 7.53125ZM10 14.3516C7.60051 14.3516 5.64844 12.3995 5.64844 10C5.64844 7.60055 7.60051 5.64844 10 5.64844C12.3995 5.64844 14.3516 7.60055 14.3516 10C14.3516 12.3995 12.3995 14.3516 10 14.3516Z"
                  fill={isActive("/settings") ? "#2D60FF" : "#B1B1B1"}
                />
              </svg>

              <p
                className={`ml-6 text-lg font-medium leading-[21.7px] ${
                  isActive("/settings") ? "text-[#2D60FF]" : "text-[#B1B1B1]"
                }`}
              >
                Settings
              </p>
            </Link>
          </li>

          {/* Add more sidebar items here */}
        </ul>
      </div>

      {/* Overlay to close the sidebar when clicking outside */}
      {isOpen && (
        <div
          onClick={toggleSidebar}
          className="fixed inset-0 bg-black opacity-50 mobile:hidden"
        ></div>
      )}
    </>
  );
};

export default Sidebar;