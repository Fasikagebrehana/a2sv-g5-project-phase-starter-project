import EditProfileForm from "@/app/components/Forms/EditProfileForm";
import Navigation from "@/app/components/Settings/Navigation";
import React from "react";

const ProfilePage: React.FC = () => {
  return (
    <div className="w-full mt-2 flex flex-col justify-start items-center ">
      <div className="bg-white  p-4 dark:bg-[#232328] rounded-xl min-h-screen w-[80%]">
        <Navigation />
        <EditProfileForm />
      </div>
    </div>
  );
};

export default ProfilePage;
