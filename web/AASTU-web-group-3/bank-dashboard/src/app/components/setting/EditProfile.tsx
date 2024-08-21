import React, { useRef } from "react";
import { useForm } from "react-hook-form";
import Image from "next/image";
import {
  getStorage,
  ref,
  uploadBytesResumable,
  getDownloadURL,
} from "firebase/storage";
import app from "@/app/firebase";
import { circleWithPen, profilepic } from "@/../../public/Icons";
import { useDispatch, useSelector } from "react-redux";
import { usePutSettingMutation } from "@/lib/redux/api/settingApi";
import { RootState } from "@/lib/redux/store";

interface FormInput{
  name: string;
  email: string;
  dateOfBirth: string;
  permanentAddress: string;
  postalCode: string;
  username: string;
  presentAddress: string;
  city: string;
  country: string;
  profilePicture: string; // URL of the uploaded image
}


const EditProfile = () => {

  const { register, handleSubmit, setValue } = useForm<FormInput>({
    defaultValues: {
      name: "",
      email: "",
      dateOfBirth:"",
      permanentAddress: "",
      postalCode: "",
      username: "",
      presentAddress: "",
      city: "",
      country: "",
      profilePicture: "",
    },
  });

  const dispatch = useDispatch();
  const { loading, error } = useSelector((state: RootState) => state.service);

  const [putSetting] = usePutSettingMutation();

  const fileInputRef = useRef<HTMLInputElement>(null);

  const onSubmit = async (data: FormInput) => {
    try {
  
      
      await putSetting(data ).unwrap();
    } catch (err) {
      console.error(err);
    }
  };
  

  const handleImageClick = () => {
    if (fileInputRef.current) {
      fileInputRef.current.click();
    }
  };

  const handleFileChange = async (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.[0];
    if (file) {
      const imageUrl = await uploadImageToCloud(file);
      console.log(imageUrl)
      setValue("profilePicture", imageUrl);
    } else {
      setValue("profilePicture", "");
    }
  };

  // Mock upload function (replace this with actual implementation)
  const uploadImageToCloud = async (file: File): Promise<string> => {
    const storage = getStorage(app);
    const storageRef = ref(storage, file.name);
  
    const uploadTask = uploadBytesResumable(storageRef, file);
  
    return new Promise((resolve, reject) => {
      uploadTask.on(
        'state_changed',
        (snapshot) => {
          const progress = (snapshot.bytesTransferred / snapshot.totalBytes) * 100;
          console.log(`Upload is ${progress}% done`);
        },
        (error) => {
          console.error("Error during upload:", error);
          reject(error);
        },
        async () => {
          try {
            const url = await getDownloadURL(uploadTask.snapshot.ref);
            console.log("File available at", url);
            resolve(url);
          } catch (err) {
            console.error("Error getting download URL:", err);
            reject(err);
          }
        }
      );
    });
  };

  const formFields = [
    { id: "name", label: "Your Name", placeholder: "Charlene Reed", type: "text" },
    { id: "username", label: "User Name", placeholder: "Charlene Reed", type: "text" },
    { id: "email", label: "Email", placeholder: "charlenereed@gmail.com", type: "email" },
    { id: "dateOfBirth", label: "Date of Birth", placeholder: "Enter Date of Birth", type: "date" },
    { id: "presentAddress", label: "Present Address", placeholder: "San Jose, California, USA", type: "text" },
    { id: "permanentAddress", label: "Permanent Address", placeholder: "San Jose, California, USA", type: "text" },
    { id: "city", label: "City", placeholder: "San Jose", type: "text" },
    { id: "postalCode", label: "Postal Code", placeholder: "45962", type: "text" },
    { id: "country", label: "Country", placeholder: "USA", type: "text" },
  ];

  return (
    <div className="p-4 flex flex-col md:flex-row gap-8">
      <div className="relative rounded-full w-64 h-64 mb-5 md:w-40 md:h-40">
        <Image
          src={profilepic}
          width={256}
          height={256}
          alt="profilepic"
          className="w-64 h-64 md:w-40 md:h-40 object-cover rounded-full"
        />
        <Image
          src={circleWithPen}
          alt="edit icon"
          width={64}
          height={64}
          className="absolute z-30 right-1 bottom-10 object-cover md:w-10 md:h-10 md:bottom-5 cursor-pointer"
          onClick={handleImageClick}
        />
        <input
          type="file"
          ref={fileInputRef}
          className="hidden"
          onChange={handleFileChange}
        />
      </div>

      <form
        onSubmit={handleSubmit(onSubmit)}
        className="flex flex-wrap items-end justify-between w-full md:w-4/5"
      >
        {formFields.map((field) => (
          <div key={field.id} className="mb-3 w-full md:w-[45%]">
            <label className="block text-black text-sm mb-2" htmlFor={field.id}>
              {field.label}
            </label>
            <input
              className="w-full p-3 md:p-2 text-[#718EBF] border-2 text-sm border-[#DFEAF2] rounded-lg focus:outline-none"
              type={field.type}
              id={field.id}
              placeholder={field.placeholder}
              {...register(field.id as keyof FormInput, {
                required: `${field.label} is required`,
              })}
            />
          </div>
        ))}
        <div className="flex justify-end w-full">
          <button
            className="w-full md:w-1/5 bg-[#1814F3] text-white font-semibold py-2 px-4 rounded-lg focus:outline-none"
            type="submit"
          >
            Save
          </button>
        </div>
      </form>
    </div>
  );
};

export default EditProfile;
 