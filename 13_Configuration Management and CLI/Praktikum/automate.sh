username_facebook="$aisyah"
username_linkedin="$aisyahrahmadanipohontu"

current_time=$(date "+%a %b %d %T %Z %Y")
main_dir="aisyah at $current_time"

mkdir -p "$main_dir"

mkdir -p "$main_dir/about_me/personal"
mkdir -p "$main_dir/about_me/professional"

echo "Facebook: $username_facebook" > "$main_dir/about_me/personal/facebook.txt"
echo "LinkedIn: $username_linkedin" > "$main_dir/about_me/professional/linkedin.txt"

mkdir -p "$main_dir/my_friends"
curl -o "$main_dir/my_friends/list_of_my_friends.txt" https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt

mkdir -p "$main_dir/my_system_info"
about_this_laptop_info=$(uname -a)
echo "User: $USER" > "$main_dir/my_system_info/about_this_laptop.txt"
echo "System Info: $about_this_laptop_info" >> "$main_dir/my_system_info/about_this_laptop.txt"

ping -c 3 google.com > "$main_dir/my_system_info/internet_connection.txt"

tree "$main_dir"