import CustomNavbar from '@/components/navbar/CustomNavbar'
import './globals.css'
import type { Metadata } from 'next'
import { Inter } from 'next/font/google'


export const metadata: Metadata = {
  title: 'Create Next App',
  description: 'Generated by create next app',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <body className="bg-white">
        <CustomNavbar/>
        {children}
      </body>
    </html>
  )
}
