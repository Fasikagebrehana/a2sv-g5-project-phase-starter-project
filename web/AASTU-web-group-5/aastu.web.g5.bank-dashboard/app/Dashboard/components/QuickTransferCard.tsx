import React from 'react';

export const QuickTransferCard = () => {
  return (
    <div className='card flex flex-col items-center gap-4 p-4 rounded-lg'>
      <div className='profile-picture flex flex-col items-center gap-2'>
        <svg className='w-20 h-20' xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 48 48">
          <g fill="black" fill-rule="evenodd" clip-rule="evenodd">
            <path d="M24 27a8 8 0 1 0 0-16a8 8 0 0 0 0 16m0-2a6 6 0 1 0 0-12a6 6 0 0 0 0 12" />
            <path d="M44 24c0 11.046-8.954 20-20 20S4 35.046 4 24S12.954 4 24 4s20 8.954 20 20M33.63 39.21A17.9 17.9 0 0 1 24 42a17.9 17.9 0 0 1-9.831-2.92q-.36-.45-.73-.93A2.14 2.14 0 0 1 13 36.845c0-1.077.774-1.98 1.809-2.131c6.845-1 11.558-.914 18.412.035A2.08 2.08 0 0 1 35 36.818c0 .48-.165.946-.463 1.31q-.461.561-.907 1.082m3.355-2.744c-.16-1.872-1.581-3.434-3.49-3.698c-7.016-.971-11.92-1.064-18.975-.033c-1.92.28-3.335 1.856-3.503 3.733A17.94 17.94 0 0 1 6 24c0-9.941 8.059-18 18-18s18 8.059 18 18a17.94 17.94 0 0 1-5.015 12.466" />
          </g>
        </svg>
        <p className='text-[#232323] font-Inter text-sm mt-2'>Livia Beltor</p>
        <p className='text-[#718EBF] font-Inter text-sm mt-1'>CEO</p>
      </div>
    </div>
  );
};
