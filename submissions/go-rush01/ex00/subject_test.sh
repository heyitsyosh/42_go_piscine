set -exu
go run . "...2." "..6.4" "5...6" "7.6.." ".3..." | cat -e
go run . "....." "8.8.." ".7.7." "..8.5" "....." | cat -e
go run . "37..." "..8.." "....." "..8.." "...69" | cat -e
go run . "5..." "...." "...." "...." | cat -e
go run . "...." "5..." "...." "...." | cat -e
go run . "...." "...." "5..." "...." | cat -e
go run . "...." "...." "...." "5..." | cat -e
go run . "...." "...." "...." "...." | cat -e

echo "Error cases:"
echo "Check for capital letter"
go run . "...A." "..6.4" "5....6" "7.6.." ".3..." | cat -e
echo "欠けた正方形"
go run . "....." "..6.4" "5...6" "7.6.." ".3.." | cat -e
echo "abcを含む"
go run . "37..." "..8.." "....." "..8.." "..abc" | cat -e
echo "１を含む"
go run . "17..." "..8.." "....." "..8.." "...69" | cat -e
echo "高さ5 幅4"
go run . "5..." "...." "...." "...." "...." | cat -e
echo "Not solvable"
go run . "27..." "..8.." "....." "..8.." "...69" | cat -e
echo "No args"
go run . | cat -e