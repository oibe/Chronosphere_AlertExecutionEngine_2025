package main

type Silence struct {
	id       int64
	alertIds []*int32
	teams    []string
	//todo: expire the silence?
}
