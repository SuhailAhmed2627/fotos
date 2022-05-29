const imgStyle = {
	transition:
		"transform .135s cubic-bezier(0.0,0.0,0.2,1),opacity linear .15s",
};

const cont = {
	backgroundColor: "#eee",
	cursor: "pointer",
	overflow: "hidden",
	position: "relative",
};

const SelectedImage = ({ photo, margin, direction, top, left }: any) => {
	if (direction === "column") {
		cont.position = "absolute";
		// @ts-ignore
		cont.left = left;
		// @ts-ignore
		cont.top = top;
	}

	// @ts-ignore
	const handleOnClick = (e) => {
		window.location.href = `/image/${photo.id}`;
	};

	return (
		<div
			// @ts-ignore
			style={{ margin, height: photo.height, width: photo.width, ...cont }}
			className={"hover:transition-all hover:brightness-75"}
		>
			<img
				alt={photo.title}
				style={{ ...imgStyle }}
				{...photo}
				onClick={handleOnClick}
			/>
		</div>
	);
};

export default SelectedImage;
