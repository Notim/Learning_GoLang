package string

func (str String) IsNullOrEmpty() bool{
    return len(str) == 0
}