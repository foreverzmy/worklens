fetch("/blame")
	.then((response) => {
		const reader = response.body.getReader();
		const decoder = new TextDecoder();
		let result = "";

		// 定义一个函数用于读取流数据
		function read() {
			reader.read().then(({ done, value }) => {
				if (done) {
					console.log("Stream complete");
					console.log(result);
					return;
				}

				const text = decoder.decode(value, { stream: true });
				// 将读取到的流数据解码并拼接到 result 中
				result += text;

				console.log(text);
				read();
			});
		}

		// 开始读取流数据
		read();
	})
	.catch((error) => {
		console.error("Error:", error);
	});
