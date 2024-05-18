# GitHub Contributions Visualization

![CleanShot 2024-05-18 at 15 27 19](https://github.com/marcusziade/contributions/assets/47460844/cbd473e7-e507-49c9-83ab-6fdd453da307)



## How to Generate Your Own Visualizations


### Prerequisites

1. **GitHub Personal Access Token:**
   - You need a GitHub Personal Access Token with `repo` and `user` permissions.
   - Set it as an environment variable `GITHUB_TOKEN`.

2. **Python Environment:**
   - Ensure you have Python installed.
   - You will create a virtual environment and install necessary libraries.

3. **Go Environment:**
   - Ensure you have Go installed to run the Go script.

### Step-by-Step Instructions

1. **Clone the Repository:**
   ```sh
   git clone https://github.com/marcusziade/github-contributions-visualization.git
   cd github-contributions-visualization
   ```

2. **Set Up Python Virtual Environment:**
   ```sh
   python3 -m venv env
   source env/bin/activate  # On Windows use `env\Scripts\activate`
   ```

3. **Install Python Dependencies:**
   ```sh
   pip install pandas matplotlib
   ```

4. **Fetch GitHub Contributions:**
   - Update the `username` variable in the Go script (`fetch_contributions.go`) to your GitHub username.
   - Run the Go script to fetch your GitHub contributions and save them to a CSV file:
     ```sh
     go run fetch_contributions.go
     ```

5. **Generate Visualizations:**
   - Run the Python script to generate visualizations from the CSV file:
     ```sh
     python3 visualize_contributions.py
     ```

### Result

After running the above steps, you will have the following files generated in your repository:

- `contributions.csv`: CSV file containing your GitHub contributions data.
- `contributions_overview.png`: Image containing the visualizations of your GitHub contributions.

Include the generated image in your GitHub repository README or any documentation to showcase your contributions visually.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
```
