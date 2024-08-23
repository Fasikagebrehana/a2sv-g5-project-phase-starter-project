import React from "react";
import logo from "../../../public/images/Logo.svg";
import bell from "../../../public/images/bell.svg";
import settings from "../../../public/images/settings.svg";
import hamburger from "../../../public/images/ham.jpg";
import Image from "next/image";
import search from "../../../public/images/search.svg";
import profile from "../../../public/images/Mask Group.svg";

interface Props {
  title: string;
  profilepic: string;
}

const Header = () => {
  const title = "Overview"
  const profilepic = profile


  return (
    <div className="flex bg-white md:pr-10 pb-1 md:px-0 items-center justify-between">
      <div className="w-1/6 md:block hidden">
        <Image src={logo} className="ml-1" alt="LOGO" />
      </div>

      <div className="md:w-5/6 w-full md:pl-4 flex">
        <div className="flex flex-row flex-wrap justify-center md:items-center align-middle w-full space-y-4 md:space-y-0 ">
          <div className="md:hidden  flex justify-start align-middle pt-3 w-1/3">
            <Image
              className="w-8 h-8 order-1 bg-white"
              src={hamburger}
              alt="hamburger"
            />
          </div>

          <div className="md:text-2xl md:order-none text-xl order-2 md:mr-96 md:w-fit w-1/3 md:mb-0 text-[#343C6A] font-bold">
            <h1 className="ml-2 font-bold">{title}</h1>
          </div>

          <div className="flex md:ml-32 order-4 md:order-none bg-[#F5F7FA] h-10 md:w-48 w-full items-center rounded-full  ">
            <Image src={search} className="w-3 h-3 ml-2" alt="search" />
            <input
              className="bg-[#F5F7FA] md:w-38 w-10/12 px-4 outline-none placeholder:bg-[#F5F7FA] placeholder:text-xs"
              type="text"
              placeholder={`search for something`}
            />
          </div>

          <Image
            className="w-10 ml-7 h-20 md:block hidden"
            src={settings}
            alt="settings"
          />
          <Image
            className="w-10 h-20 ml-7 self-end md:block hidden"
            src={bell}
            alt="bell"
          />
          <div className="flex justify-end md:w-fit w-1/3 md:mb-5 order-3 md:order-3 md:ml-7">
            <Image
              className="w-10 h-10"
              src={profilepic}
              alt="profile_picture"
            />
          </div>
        </div>
      </div>
    </div>
  );
};

export default Header;
