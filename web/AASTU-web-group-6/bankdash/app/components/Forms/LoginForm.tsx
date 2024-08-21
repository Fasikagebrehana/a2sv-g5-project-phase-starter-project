"use client";
import { useForm } from 'react-hook-form';
import Link from "next/link";
import React, { useState } from "react";
import { useRouter } from "next/navigation";
import LoginValue from '@/types/LoginValue';
import { signIn } from 'next-auth/react';

const LoginForm = () => {
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);
  const router = useRouter();
  const { register, handleSubmit, formState: { errors } } = useForm<LoginValue>();

  const onSubmit = async (data: LoginValue) => {
    setLoading(true);
    setError("");
    console.log("login")

    const result = await signIn("credentials", {
      redirect: false,
      userName: data.userName,
      password: data.password,
    });

    if (result?.error) {
      console.log(result.error)
      setError(result.error);
      setLoading(false);
    } else {
      console.log("Login Successful:", result);
      router.push("/");
    }
  };

  return (
    <div className="max-w-[408px] min-w-72 p-4 bg-[#fafafa] rounded-lg">
      <div className='flex justify-center'>
        <h1 className='min-w-[408px] p-1 pb-1 mb-4 mt-1 text-[#25324B] font-black text-4xl text-center'>
          Welcome Back,
        </h1>
      </div>
      <form className="p-3" onSubmit={handleSubmit(onSubmit)}>
        <div className="mt-3 flex flex-col">
          <label className='mb-1' htmlFor="userName">
            UserName
          </label>
          <input
            {...register("userName", { required: "UserName is required" })}
            placeholder='Enter UserName'
            id="userName"
            className='h-8 py-1 px-2 border-[1px] border-gray-400 rounded-lg'
            type="text"
            disabled={loading}
          />
          {errors.userName && <p className="text-red-500">{errors.userName.message}</p>}
        </div>

        <div className="mt-3 flex flex-col">
          <label className='mb-1' htmlFor="password">
            Password
          </label>
          <input
            {...register("password", { required: "Password is required" })}
            placeholder='Enter password'
            id="password"
            className='h-8 py-1 px-2 border-[1px] border-gray-400 rounded-lg'
            type="password"
            disabled={loading}
          />
          {errors.password && <p className="text-red-500">{errors.password.message}</p>}
        </div>

        <div className="px-6 py-3 mt-3 flex flex-col bg-[#4640DE] rounded-xl">
          <button type='submit' disabled={loading} className="text-white">
            {loading ? 'Loading...' : 'Login'}
          </button>
          {error && <p className="text-red-500 mt-2">{error}</p>}
        </div>
      </form>

      <div className='mt-3 py-1'>
        <p className='text-base font-serif font-thin text-[#202430]'>
          Don&apos;t have an account?
          <Link href={`/signup`} passHref>
            <span className='ml-2 font-bold font-serif text-[#4640DE]'>
              SignUp
            </span>
          </Link>
        </p>
      </div>
    </div>
  );
};

export default LoginForm;
