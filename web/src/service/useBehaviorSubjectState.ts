import { useEffect, useMemo, useRef, useState } from 'react';
import { BehaviorSubject, type Observable, type OperatorFunction } from 'rxjs';

export const useBehaviorSubjectState = <T, R = T>(
	initValue: R,
	pipe?: OperatorFunction<T, R>,
) => {
	const sub$Ref = useRef(new BehaviorSubject<R>(initValue));
	const [state, setState] = useState(initValue);
	const pipeRef = useRef(pipe);

	const unsubscribe = useMemo(() => {
		let sub$: Observable<any> = sub$Ref.current;

		if (pipeRef.current) {
			sub$ = sub$.pipe(pipeRef.current);
		}

		return sub$.subscribe((value: R) => {
			setState(value);
		}).unsubscribe;
	}, []);

	useEffect(() => () => unsubscribe(), [unsubscribe]);

	return [
		state,
		{
			sub$Ref,
			unsubscribe,
		},
	] as const;
};
