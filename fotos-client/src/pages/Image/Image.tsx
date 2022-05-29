import { Box, Center, Grid, Loader, Text } from "@mantine/core";
import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { ImageType } from "../../types";
import { dataFetch, getUser } from "../../utils/helperFuntions";

const Image = () => {
	const { imageId } = useParams();

	const [isLoading, setIsLoading] = useState(true);
	const [image, setImage] = useState<ImageType | null>(null);
	const [isError, setIsError] = useState(false);
	const navigate = useNavigate();
	const user = getUser();

	if (!user) {
		navigate("/login");
	}

	useEffect(() => {
		const fetchImage = async () => {
			const response = await dataFetch(
				"/api/image/get",
				user,
				{
					"Content-Type": "application/json",
				},
				"POST",
				JSON.stringify({
					imageId: imageId,
				})
			);
			const responseData = await response.json();
			if (response.ok) {
				setImage(responseData);
			} else {
				setIsError(true);
			}
			setIsLoading(false);
		};
		fetchImage().catch(console.error);
	}, []);

	if (isLoading) {
		return (
			<Center className="w-full h-full">
				<Loader size={"lg"} />
			</Center>
		);
	}

	if (isError || !image) {
		return <Center className="w-full h-full">Some Error has Occurred</Center>;
	}

	return (
		<Grid className="overflow-hidden w-full h-full items-center text-gray-700">
			<Center className=" w-[70%] h-[90%] object-scale-down">
				<img className="" src={image.src} />
			</Center>
			<Box className="w-[30%] pl-5 pt-20 h-full border-l-2">
				<Text className="text-h5 font-semibold">{image.name}</Text>
				<Text>By {image.uploaderName}</Text>
				<Text className=" text-small">
					{new Date(image.clickedAt as string).toLocaleString()}
				</Text>
				{!image.faceProcessed && (
					<Box className="flex flex-row items-center gap-2 mt-2 mb-2">
						<span>
							<Loader size={"xs"} />
						</span>
						Processing Faces...
					</Box>
				)}
				{image.faceProcessed && image.hasYou && (
					<Box className="flex flex-row items-center gap-2 mt-2 mb-2">
						<Text className="font-extrabold font-title text-transparent text-h3 bg-clip-text bg-gradient-to-r from-primary to-secondary">
							Has You
						</Text>
					</Box>
				)}
			</Box>
		</Grid>
	);
};

export default Image;
