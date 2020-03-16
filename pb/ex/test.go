package main


// Redact clears every sensitive field in pb.
func Redact(pb proto.Message) {
	// ...
}

type ProtoMessage interface{
	ProtoReflect() Message
}

