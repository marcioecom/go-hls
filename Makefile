hls:
	ffmpeg -v verbose \
		-re \
		-i videos/video.mp4 \
		-c:v libx264 \
		-b:v 5000k \
		-f hls \
		-hls_time 6 \
		-hls_list_size 4 \
		-hls_wrap 40 \
		-hls_delete_threshold 1 \
		-hls_flags delete_segments \
		-preset superfast \
		-start_number 1 \
		./videos/stream.m3u8