import { Box, Center, Group, Image, Stack, Text } from "@mantine/core";
import { DropzoneStatus } from "@mantine/dropzone";
import { showNotification } from "@mantine/notifications";
import { AnyAction, Dispatch } from "@reduxjs/toolkit";
import { MdClose, MdFileUpload, MdInsertPhoto } from "react-icons/md";
import { useDispatch } from "react-redux";
import { useNavigate } from "react-router-dom";
import { loginSuccess } from "../../actions/user";
import DropZone from "../../components/DropZone/DropZone";
import { User } from "../../types";
import { dataFetch, getUser } from "../../utils/helperFuntions";

function ImageUploadIcon({
	status,
	size,
}: {
	status: DropzoneStatus;
	size: number;
}) {
	if (status.accepted) {
		return <MdFileUpload size={size} />;
	}

	if (status.rejected) {
		return <MdClose size={size} />;
	}

	return <MdInsertPhoto size={size} />;
}

const userFirstImageDropzoneChildren = (
	status: DropzoneStatus,
	files: File[]
) => {
	if (files.length > 0) {
		return (
			<>
				<Box
					key={files[0].name}
					className="w-full flex justify-center flex-row gap-5"
				>
					<Image width={200} src={URL.createObjectURL(files[0])}></Image>
				</Box>
				<Box className="mb-4">Drop/Select Again to Change Image</Box>
			</>
		);
	}
	return (
		<Group
			position="center"
			spacing="xl"
			style={{ minHeight: 220, pointerEvents: "none" }}
		>
			<ImageUploadIcon status={status} size={80} />

			<div>
				<Text size="xl" inline>
					Drag images here or click to select files
				</Text>
				<Text size="sm" color="dimmed" inline mt={7}>
					Attach a photo of yourself
				</Text>
			</div>
		</Group>
	);
};

const generateFunc = (dispatch: Dispatch<AnyAction>) => {
	const userFirstImageHandleUpload = async (
		user: User,
		files: File[],
		setIsDropZoneLoading: React.Dispatch<React.SetStateAction<boolean>>,
		setDropzoneFiles: React.Dispatch<React.SetStateAction<File[]>>,
		eventId?: number
	) => {
		const formData = new FormData();
		files.forEach((file) => {
			formData.append("files", file);
		});
		const response = await dataFetch(
			"/api/user/face",
			user,
			{},
			"POST",
			formData
		);

		const responseData = await response.json();
		setIsDropZoneLoading(false);
		setDropzoneFiles([]);
		if (response.ok) {
			showNotification({
				title: "Upload Successful",
				message: `${responseData.message}`,
				color: "green",
			});
			const updatedUser: User = {
				...user,
				firstLogin: false,
			};
			dispatch(loginSuccess(updatedUser));
			window.location.href = "/home";
		} else {
			showNotification({
				title: "Upload Failed",
				message: `Error Handling File: ${responseData.message}`,
				color: "red",
			});
		}
	};
	return userFirstImageHandleUpload;
};

const UserImageUpload = () => {
	const user = getUser();
	const navigate = useNavigate();
	const dispatch = useDispatch();

	if (!user) {
		navigate("/login");
	}
	if (user && !user.firstLogin) {
		navigate("/home");
	}
	return (
		<Stack className="items-center justify-center h-full w-full p-5 text-gray-600">
			<Text className="w-full text-center font-title font-bold text-h2 md:text-h1">
				<span>Welcome to</span>
				<Image
					className="flex justify-center"
					width={300}
					src="assets/images/logo-color.svg"
				></Image>
			</Text>
			<Text className="font-body text-center">
				Please Upload a High Quality Image of your Face to get Started
			</Text>
			<Center className="flex-col w-full md:w-[700px]">
				<DropZone
					handleUpload={generateFunc(dispatch)}
					dropzoneChildren={userFirstImageDropzoneChildren}
				/>
			</Center>
		</Stack>
	);
};

export default UserImageUpload;
