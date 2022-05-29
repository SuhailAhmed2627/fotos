import { Link } from "react-router-dom";
import { useState } from "react";
import { AppShell, Navbar, Header } from "@mantine/core";
import { getUser } from "../../utils/helperFuntions";

const links = [
	{
		to: "#yourface",
		title: "Your Face",
	},
	{
		to: "#secure",
		title: "Security",
	},
	{
		to: "#features",
		title: "Features",
	},
];

const NavLink = (to: string, title: string): JSX.Element => {
	return (
		<a
			href={to}
			className="block mt-4 text-center md:text-left md:inline-block md:mt-0 text-lg md:text-neutral-200 transition ease-in-out md:hover:text-white md:mr-10 md:hover:scale-110 duration-150"
		>
			{title}
		</a>
	);
};

const NavBar = () => {
	const [viewMenu, setViewMenu] = useState<boolean>(false);
	const user = getUser();

	if (user) {
		return <></>;
	}

	return (
		<Header
			className={
				"fixed border-none top-0 left-0 z-100 w-full flex items-center justify-between flex-wrap bg-gradient-to-r from-primary to-secondary text-white pt-5 pb-5 pr-6 pl-6"
			}
			height={80}
			p="xs"
		>
			<div className="flex items-center flex-shrink-0 text-white mr-6">
				<svg
					height="30"
					viewBox="0 0 324 70"
					fill="none"
					xmlns="http://www.w3.org/2000/svg"
				>
					<path
						d="M0.12 1.824V68H13.56V41.888H43V28.32H13.56V15.392H43.896V1.824H13.944H0.12ZM88.808 0.159996C69.608 0.159996 54.12 15.776 54.12 34.976C54.12 54.048 69.608 69.664 88.808 69.664C108.008 69.664 123.624 54.048 123.624 34.976C123.624 15.776 108.008 0.159996 88.808 0.159996ZM88.808 6.432C104.552 6.432 117.224 19.104 117.224 34.976C117.224 50.72 104.552 63.392 88.808 63.392C73.064 63.392 60.392 50.72 60.392 34.976C60.392 19.104 73.064 6.432 88.808 6.432ZM133.742 1.696V15.392H153.454V68H167.278V15.392H186.99V1.696H133.742ZM231.933 0.159996C212.733 0.159996 197.245 15.776 197.245 34.976C197.245 54.048 212.733 69.664 231.933 69.664C251.133 69.664 266.749 54.048 266.749 34.976C266.749 15.776 251.133 0.159996 231.933 0.159996ZM231.933 6.432C247.677 6.432 260.349 19.104 260.349 34.976C260.349 50.72 247.677 63.392 231.933 63.392C216.189 63.392 203.517 50.72 203.517 34.976C203.517 19.104 216.189 6.432 231.933 6.432ZM300.803 0.159996C289.027 0.159996 279.171 8.864 279.171 19.744C279.171 31.264 288.771 37.024 298.115 39.584C302.723 40.992 310.019 43.168 310.019 48.16C310.019 53.536 304.899 55.84 300.291 55.84C293.891 55.84 288.515 52.896 283.907 49.44L276.995 61.088C283.523 66.208 290.947 69.664 300.419 69.664C314.755 69.664 323.843 59.04 323.843 48.16V48.032C323.843 36.768 314.627 29.984 303.619 27.04C300.035 25.888 292.995 23.712 292.995 19.744C292.995 15.776 297.219 13.6 300.803 13.6C305.539 13.6 309.123 15.392 312.195 17.312L319.875 6.304C315.011 3.36 309.763 0.159996 300.803 0.159996Z"
						fill="white"
					/>
				</svg>
			</div>
			<div className="block md:hidden">
				<button
					onClick={() => setViewMenu(!viewMenu)}
					className="flex group hover:bg-white items-center border rounded px-3 py-2 text-teal-200 border-teal-400 hover:text-white hover:border-white"
				>
					<svg
						className="fill-white h-4 w-4 group-hover:fill-black"
						viewBox="0 0 20 20"
						xmlns="http://www.w3.org/2000/svg"
					>
						<title>Menu</title>
						<path d="M0 3h20v2H0V3zm0 6h20v2H0V9zm0 6h20v2H0v-2z" />
					</svg>
				</button>
			</div>
			<div
				className={`${
					viewMenu ? "" : "hidden"
				} transition-height duration-500 ease-in-out w-full block flex-grow md:flex md:items-center md:w-auto`}
			>
				<div className="w-full text-sm md:flex md:flex-row md:justify-end md:pr-4">
					{links.map((link) => NavLink(link.to, link.title))}
				</div>
				<div className="flex items-center justify-center">
					<Link
						to="/login"
						className="inline-block text-sm px-4 py-2 leading-none border rounded text-white border-white hover:border-transparent hover:text-teal-500 hover:bg-white mt-4 md:mt-0"
					>
						Login
					</Link>
				</div>
			</div>
		</Header>
	);
};

export default NavBar;
