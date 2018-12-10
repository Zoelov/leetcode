package queue

import "errors"

type Queue []interface{}

func (q Queue) Len() int {
	return len(q)
}

func (q Queue) IsEmpty() bool {
	return q.Len() == 0
}

func (q Queue) Front() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("empty")
	}

	v := q[q.Len()-1]

	return v, nil
}

func (q *Queue) EnQueue(v interface{}) {
	*q = append(*q, v)
}

func (q *Queue) DeQueue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("empty")
	}

	v := (*q)[0]
	*q = (*q)[1:]

	return v, nil
}

func (q *Queue) DeQueueBack() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("empty")
	}

	v := (*q)[q.Len()-1]
	*q = (*q)[0 : q.Len()-1]

	return v, nil
}
