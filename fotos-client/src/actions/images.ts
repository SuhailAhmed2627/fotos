import { ImageType } from "../types";
import { GET_IMAGES } from "./types";

export const putImages = (data: { images: ImageType[]; eventId: number }) => {
	const payload: Map<number, ImageType[]> = new Map();
	payload.set(data.eventId, data.images);
	return {
		type: GET_IMAGES,
		payload: payload,
	};
};
