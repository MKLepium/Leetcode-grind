import os

# Path to the folder containing your LeetCode solutions
folder_path = './'

# Path to the README file
readme_path = 'README.md'

def count_solved_problems(folder_path):
    # Count folders that start with numbers (representing problem IDs)
    return len([name for name in os.listdir(folder_path)
                if os.path.isdir(os.path.join(folder_path, name)) and name[:5].isdigit()])

def update_readme(readme_path, count):
    # Content to update in the README
    new_content = f'# LeetCode Solutions\n\nSolved **{count}** problems so far!'

    # Write to the README file
    with open(readme_path, 'w') as readme_file:
        readme_file.write(new_content)

if __name__ == '__main__':
    # Count the number of solved problems
    solved_count = count_solved_problems(folder_path)

    # Update the README file with the new count
    update_readme(readme_path, solved_count)

    print(f'Readme updated with {solved_count} solved problems.')

