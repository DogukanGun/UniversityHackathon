"use client"
import Link from "next/link";
import { EventType, MetaMaskSDK, SDKProvider, ServiceStatus } from '@metamask/sdk';
import { useEffect, useState } from "react";

const CustomNavbar = () =>{
    const [hasProvider, setHasProvider] = useState<boolean | null>(null)
    const initialState = { accounts: [""], balance: "", chainId: "" }
    const [wallet, setWallet] = useState(initialState)
    const [loading, setLoading] = useState(false)
    const [sdk, setSDK] = useState<MetaMaskSDK>();
    const [chain, setChain] = useState('');
    const [account, setAccount] = useState<string>("");
    const [serviceStatus, setServiceStatus] = useState<ServiceStatus>();
    const [connected, setConnected] = useState(false);
    const [activeProvider, setActiveProvider] = useState<SDKProvider>();
    const formatBalance = (rawBalance: string) => {
        const balance = (parseInt(rawBalance) / 1000000000000000000).toFixed(2)
        return balance
    }
    const updateWallet = async (_accounts: string[]) => {
        if (!window.ethereum) {
            throw new Error(`invalid ethereum provider`);
        }
        window.ethereum.request({
            method: "eth_getBalance",
            params: [_accounts[0], "latest"],
        }).then((value)=>{
            const _balance = formatBalance(value as string)
            window.ethereum!.request({
                method: "eth_chainId",
            }).then((_chainID) =>{
                setWallet({ accounts:_accounts, balance:_balance, chainId:(_chainID as string) });
            })
        })
    } 
    const connect = () => {
        if (!window.ethereum) {
          throw new Error(`invalid ethereum provider`);
        }
        window.ethereum
          .request({
            method: 'eth_requestAccounts',
            params: [],
          })
          .then((accounts) => {
            if (accounts) {
              console.debug(`connect:: accounts result`, accounts);
              updateWallet(accounts as string[])
              sessionStorage.setItem("account", wallet.accounts[0] as string)
              setConnected(true);
            }
          })
          .catch((e) => console.log('request accounts ERR', e));
    };
    useEffect(() => {
        const doAsync = async () => {
            const clientSDK = new MetaMaskSDK({
                useDeeplink: false,
                communicationServerUrl: process.env.NEXT_PUBLIC_COMM_SERVER_URL,
                checkInstallationImmediately: false,
                dappMetadata: {
                    name: 'NEXTJS demo',
                    url: window.location.host,
                },
                logging: {
                    developerMode: false,
                },
                storage: {
                    enabled: true,
                },
            });
            await clientSDK.init();
            setSDK(clientSDK);
        };
        doAsync();
    }, []);

    useEffect(() => {
        if (!sdk || !activeProvider) {
            return;
        }

        // activeProvider is mapped to window.ethereum.
        console.debug(`App::useEffect setup active provider listeners`);
        if (window.ethereum?.selectedAddress) {
            console.debug(`App::useEffect setting account from window.ethereum `);
            setAccount(window.ethereum?.selectedAddress);
            setConnected(true);
        } else {
            setConnected(false);
        }

        const onChainChanged = (chain: unknown) => {
            console.log(`App::useEfect on 'chainChanged'`, chain);
            setChain(chain as string);
            setConnected(true);
        };

        const onInitialized = () => {
            console.debug(`App::useEffect on _initialized`);
            setConnected(true);
            if (window.ethereum?.selectedAddress) {
                setAccount(window.ethereum?.selectedAddress);
            }

            if (window.ethereum?.chainId) {
                setChain(window.ethereum.chainId);
            }
        };

        const onAccountsChanged = (accounts: unknown) => {
            console.log(`App::useEfect on 'accountsChanged'`, accounts);
            setAccount((accounts as string[])?.[0]);
            setConnected(true);
        };

        const onConnect = (_connectInfo: unknown) => {
            console.log(`App::useEfect on 'connect'`, _connectInfo);
            setConnected(true);
        };

        const onDisconnect = (error: unknown) => {
            console.log(`App::useEfect on 'disconnect'`, error);
            setConnected(false);
            setChain('');
        };

        const onServiceStatus = (_serviceStatus: ServiceStatus) => {
            console.debug(`sdk connection_status`, _serviceStatus);
            setServiceStatus(_serviceStatus);
        };

        window.ethereum?.on('chainChanged', onChainChanged);

        window.ethereum?.on('_initialized', onInitialized);

        window.ethereum?.on('accountsChanged', onAccountsChanged);

        window.ethereum?.on('connect', onConnect);

        window.ethereum?.on('disconnect', onDisconnect);

        sdk.on(EventType.SERVICE_STATUS, onServiceStatus);

        return () => {
            console.debug(`App::useEffect cleanup activeprovider events`);
            window.ethereum?.removeListener('chainChanged', onChainChanged);
            window.ethereum?.removeListener('_initialized', onInitialized);
            window.ethereum?.removeListener('accountsChanged', onAccountsChanged);
            window.ethereum?.removeListener('connect', onConnect);
            window.ethereum?.removeListener('disconnect', onDisconnect);
            sdk.removeListener(EventType.SERVICE_STATUS, onServiceStatus);
        }
    }, [activeProvider])

    useEffect(() => {
        if (!sdk?.isInitialized()) {
          return;
        }
    
        const onProviderEvent = (accounts?: string[]) => {
          if (accounts?.[0]?.startsWith('0x')) {
            setConnected(true);
            setAccount(accounts?.[0]);
          } else {
            setConnected(false);
            setAccount("");
          }
          setActiveProvider(sdk.getProvider());
        };
        // listen for provider change events
        sdk.on(EventType.PROVIDER_UPDATE, onProviderEvent);
        return () => {
          sdk.removeListener(EventType.PROVIDER_UPDATE, onProviderEvent);
        };
      }, [sdk]);

    const handleConnect = async (event: React.MouseEvent<HTMLButtonElement>) => {
        if(account === ""){
            event.preventDefault()
            connect()
        }
    }


    return (
        <nav className="bg-gray-900 border-gray-200 dark:bg-gray-900">
            <div className="max-w-screen-xl flex flex-wrap items-center justify-between mx-auto p-4">
                <a href="https://flowbite.com/" className="flex items-center">
                    <span className="self-center text-2xl font-semibold whitespace-nowrap dark:text-white">SolTeam</span>
                </a>
                <div className="flex md:order-2">
                    <button onClick={handleConnect} type="button" className="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-4 py-2 text-center mr-3 md:mr-0 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">{account === "" ? "Connect Wallet": account}</button>
                    <button data-collapse-toggle="navbar-cta" type="button" className="inline-flex items-center p-2 w-10 h-10 justify-center text-sm text-gray-500 rounded-lg md:hidden hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 dark:text-gray-400 dark:hover:bg-gray-700 dark:focus:ring-gray-600" aria-controls="navbar-cta" aria-expanded="false">
                        <span className="sr-only">{account === "" ? "Connect Wallet": account}</span>
                        <svg className="w-5 h-5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 17 14">
                            <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M1 1h15M1 7h15M1 13h15"/>
                        </svg>
                    </button>
                </div>
                { account !== "" &&
                    <div className="bg-gray-900 items-center justify-between hidden w-full md:flex md:w-auto md:order-1" id="navbar-cta">
                        <ul className="bg-gray-900 flex flex-col font-medium p-4 md:p-0 mt-4 border border-gray-100 rounded-lg md:flex-row md:space-x-8 md:mt-0 md:border-0 dark:bg-gray-800 md:dark:bg-gray-900 dark:border-gray-700">
                            <li>
                                <Link href="/"><p className="block py-2 pl-3 pr-4 text-white rounded md:bg-transparent" aria-current="page">Home</p></Link>
                            </li>
                            <li>
                                <Link href="/donation/request"><p className="block py-2 pl-3 pr-4 text-white rounded md:bg-transparent" aria-current="page">Start Donation</p></Link>
                            </li>
                            <li>
                                <Link href="/donation/list"><p className="block py-2 pl-3 pr-4 text-white rounded md:bg-transparent" aria-current="page">Donate</p></Link>
                            </li>
                        </ul>
                    </div>
                }
            </div>
        </nav>
    )
}

export default CustomNavbar;