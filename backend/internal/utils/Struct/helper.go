package Struct

func MakeWith[Out any](assign func(it *Out)) (out *Out) {
	out = new(Out)
	assign(out)
	return
}
