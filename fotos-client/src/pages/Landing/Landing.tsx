import { Button } from "@mantine/core";

const Landing = (): JSX.Element => {
	return (
		<div className="slide-bg w-full h-full flex items-center justify-center flex-col gap-14 md:gap-10">
			<div className=" md:w-[80%] lg:w-[70%] text-center leading-none cursor-default text-white font-display font-light text-[3.5rem] md:text-[4rem] lg:text-display">
				<div className="transition-all tracking-normal hover:tracking-wide">
					YOUR <span className="font-bold">PHOTOS</span>
				</div>
				<div className="transition-all tracking-normal hover:tracking-wide">
					YOUR <span className="font-bold">STORIES</span>
				</div>
				<div className="transition-all tracking-[.12em] hover:tracking-[.15em] ">
					YOUR{" "}
					<span className="font-bold bg-gradient-to-r text-transparent bg-clip-text from-primary via-primary-500 to-secondary-800 background-animate">
						FACE
					</span>
				</div>
			</div>
			<div className="md:w-[80%] lg:w-[70%]">
				<div className="flex flex-row justify-center gap-5">
					<Button
						classNames={{
							label: "font-medium font-title group-hover:font-semibold",
						}}
						className={
							" transition-all group text-white bg-secondary mt-5 hover:bg-secondary-700"
						}
						variant="filled"
						size={"lg"}
					>
						Join Now
					</Button>
					<Button
						classNames={{
							label: "font-medium font-title group-hover:font-semibold",
						}}
						className={
							" transition-all group text-white bg-secondary mt-5 hover:bg-secondary-700"
						}
						color="white"
						variant="filled"
						size={"lg"}
					>
						Explore
					</Button>
				</div>
			</div>
		</div>
	);
};

export default Landing;
