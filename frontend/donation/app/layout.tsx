"use client"
import CustomNavbar from '@/components/navbar/CustomNavbar'
import './globals.css'
import type { Metadata } from 'next'
import { WagmiConfig, configureChains, createConfig } from 'wagmi'
import { goerli, mainnet } from 'wagmi/chains'
import { publicProvider } from 'wagmi/providers/public'


const chains = [mainnet, goerli]
const { publicClient } = configureChains(chains, [publicProvider()])

const wagmiConfig = createConfig({
  publicClient,
})

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <WagmiConfig config={wagmiConfig}>
        <body className="bg-white">
          <CustomNavbar/>
          {children}
        </body>
      </WagmiConfig>
    </html>
  )
}
