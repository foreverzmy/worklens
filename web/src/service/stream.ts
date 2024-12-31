import type { StreamMessage } from '@typings/message';
import { Observable, switchMap } from 'rxjs';
import { fromFetch } from 'rxjs/fetch';

export const streamFetch = (repo: string, path: string) => {
	return fromFetch(path, {
		method: 'POST',
		headers: {
			'content-type': 'application/json',
		},
		body: JSON.stringify({ repo }),
	}).pipe(
		switchMap((resp) => {
			if (!resp.ok || !resp.body) {
				throw new Error(`Error fetching stream: ${resp.statusText}`);
			}

			const reader = resp.body.getReader();
			const decoder = new TextDecoder();

			return new Observable<StreamMessage>((subscriber) => {
				const read = async () => {
					const { done, value } = await reader.read();

					if (done) {
						subscriber.complete();
						return;
					}

					const text = decoder.decode(value, { stream: true });
					const lines = text.split('\n').filter(Boolean);

					for (const line of lines) {
						try {
							const message = JSON.parse(line);
							subscriber.next(message);
							read(); // 继续读取
						} catch (error) {
							console.info(lines, line);
							console.error(error);
						}
					}
				};

				read();

				return () => {
					reader.cancel().catch((err) => subscriber.error(err));
				};
			});
		}),
	);
};
