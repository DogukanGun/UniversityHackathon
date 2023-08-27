interface DonationItemProps{
  onClick:()=>void
  donationName:string
}

const DonationItem = (props:DonationItemProps) =>{
    return(
      <div>
        <a onClick={props.onClick} className="group relative mb-2 block h-80 overflow-hidden rounded-lg bg-gray-100 lg:mb-3">
          <img src="/donation.png" loading="lazy" alt="Photo by Rachit Tank" className="h-full w-full object-cover object-center transition duration-200 group-hover:scale-110" />
        </a>

        <div>
          <a href="#" className="hover:gray-800 mb-1 text-gray-500 transition duration-100 lg:text-lg">Donetion for {props.donationName}</a>
        </div>
      </div>
    )
}   

export default DonationItem;