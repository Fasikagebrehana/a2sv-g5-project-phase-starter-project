import React from "react";

const Loading = () => {
  return (
    <div className="flex items-center justify-center min-h-[100vh]  bg-gray-100">
      <div className="flex flex-col items-center">
        <div className="animate-spin rounded-full h-16 w-16 border-t-4 border-black border-solid"></div>
        <p className="mt-4 text-lg text-gray-700">Loading...</p>
      </div>
    </div>
  );
};

export default Loading;
