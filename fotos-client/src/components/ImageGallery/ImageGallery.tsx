import { Center, Loader } from "@mantine/core";
import { useCallback, useEffect, useState } from "react";
import Gallery from "react-photo-gallery";
import { useDispatch, useSelector } from "react-redux";
import { useParams } from "react-router-dom";
import { putImages } from "../../actions/images";
import { ImageType } from "../../types";
import { dataFetch, getImages, getUser } from "../../utils/helperFuntions";
import SelectedImage from "./components/SelectedImage";

const ImageGallery = () => {
	const dispatch = useDispatch();
	const [isLoading, setIsLoading] = useState(true);
	const [isError, setIsError] = useState(false);
	const { eventId } = useParams();
	const user = getUser();
	const images = getImages(Number(eventId));

	const [selectAll, setSelectAll] = useState(false);
	const [selectedImages, setSelectedImages] = useState<ImageType[]>([]);

	const toggleSelectAll = () => {
		setSelectAll(!selectAll);
	};

	const imageRenderer = useCallback(
		({ index, left, top, key, photo }: any) => (
			<SelectedImage
				margin={"2px"}
				index={index}
				photo={photo}
				left={left}
				top={top}
				setSelectedImages={setSelectedImages}
			/>
		),
		[selectAll]
	);

	useEffect(() => {
		if (images && images?.length != 0) {
			setIsLoading(false);
			return;
		}
		const fetchImages = async () => {
			const response = await dataFetch(
				"/api/image/get_all",
				user,
				{
					"Content-Type": "application/json",
				},
				"POST",
				JSON.stringify({
					eventId,
				})
			);
			const responseData = await response.json();
			if (response.ok) {
				dispatch(putImages(responseData));
				setIsLoading(false);
			} else {
				setIsError(true);
			}
		};
		fetchImages().catch(() => {
			setIsError(true);
		});
	}, [eventId]);

	if (isLoading) {
		return (
			<Center className="w-full h-full">
				<Loader size={"lg"} />
			</Center>
		);
	}

	if (isError) {
		return <Center className="w-full h-full">Some Error has Occurred</Center>;
	}

	if (images === null) {
		return (
			<Center className="w-full h-full">
				No Images found for this event
			</Center>
		);
	}

	if (images) {
		return (
			<div>
				<Gallery photos={images} renderImage={imageRenderer} />
			</div>
		);
	}

	return <> </>;
};

export default ImageGallery;
