use std::{
    collections::HashSet,
    path::{Path, PathBuf},
    process::Command,
};

use clap::Parser;
use colored::*;
use regex::Regex;

const VK_SPEC_REGISTRY_URL: &str =
    "https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/";
const HTML_PAGES_DIR: &str = "vulkan_man_html_pages";
const MD_PAGES_DIR: &str = "vulkan_man_md_pages";

#[derive(Parser)]
#[command(version, about, long_about = None)]
struct Args {
    #[command(subcommand)]
    create: Commands,
}

#[derive(Parser)]
enum Commands {
    #[command()]
    Create {
        /// Skip downloading Vulkan man pages
        #[arg(long)]
        skip_download: bool,

        /// Skip generating Markdown files
        #[arg(long)]
        skip_md_gen: bool,

        /// Skip post-processing Markdown files
        #[arg(long)]
        skip_postprocess: bool,
    },
}

fn main() {
    let args = Args::parse();

    match args.create {
        Commands::Create {
            skip_download,
            skip_md_gen,
            skip_postprocess,
        } => {
            if !skip_download {
                download_vk_man_pages(HTML_PAGES_DIR);
            }

            if !skip_md_gen {
                generate_md_files(HTML_PAGES_DIR, MD_PAGES_DIR);
            }

            if !skip_postprocess {
                postprocess_md_files(MD_PAGES_DIR);
            }
        }
    }
}

fn download_vk_man_pages(target_dir: &str) {
    let message = format!(
        "Downloading Vulkan man pages from {}:",
        VK_SPEC_REGISTRY_URL
    )
    .bright_blue()
    .bold();
    println!("{}", message);

    let mut cmd = Command::new("wget");
    cmd.arg("--execute=robots=off")
        .arg("-m") // Turn on mirroring
        .arg("-nd") // No directories
        .arg("-np") // No parent
        .arg("-nH") // Disable generation of host-prefixed directories
        .arg("-P") // Directory prefix
        .arg(target_dir) // Directory to save files
        .arg(VK_SPEC_REGISTRY_URL);

    spawn_command(cmd);

    let message = "Download complete!".bright_green().bold();
    println!("{}", message);

    let message = "Removing unnecessary files:".bright_blue().bold();
    println!("{}", message);

    let html_pages_dir = std::fs::read_dir(target_dir);
    let html_pages = match html_pages_dir {
        Ok(dir) => dir,
        Err(e) => {
            let e = format!("Error: {}", e).red().bold();
            eprintln!("{}", e);

            std::process::exit(1);
        }
    };

    // delete files with index.html in their name
    for entry in html_pages {
        if let Ok(page) = entry {
            if page.file_name().to_str().unwrap().contains("index.html") {
                let message =
                    format!("Removing {}", page.file_name().to_str().unwrap()).bright_blue();
                println!("{}", message);

                let result = std::fs::remove_file(page.path());
                match result {
                    Ok(_) => {}
                    Err(e) => {
                        let e = format!("Error: {}", e).red().bold();
                        eprintln!("{}", e);

                        std::process::exit(1);
                    }
                }
            }
        }
    }
}

fn generate_md_files(input_dir: &str, output_dir: &str) {
    let message = "Generating Markdown files:".bright_blue().bold();
    println!("{}", message);

    if !Path::new(output_dir).exists() {
        let message = format!("Creating output directory {}", output_dir).bright_blue();
        println!("{}", message);

        let result = std::fs::create_dir(output_dir);

        match result {
            Ok(_) => {}
            Err(e) => {
                let e = format!("Error: {}", e).red().bold();
                eprintln!("{}", e);

                std::process::exit(1);
            }
        }
    }

    let html_pages_dir = std::fs::read_dir(input_dir);
    let html_pages = match html_pages_dir {
        Ok(dir) => dir,
        Err(e) => {
            let e = format!("Error: {}", e).red().bold();
            eprintln!("{}", e);

            std::process::exit(1);
        }
    };

    for entry in html_pages {
        if let Ok(page) = entry {
            println!("Processing {}", page.file_name().to_str().unwrap());

            let input_html_file_path = page.path();
            let output_md_file_path = PathBuf::from(output_dir).join(
                page.file_name()
                    .to_str()
                    .expect("Failed to convert file name to string")
                    .replace(".html", ".md"),
            );

            let mut cmd = Command::new("pandoc");
            cmd.arg("--from=html-native_divs-native_spans")
                .arg("--to=gfm")
                .arg(input_html_file_path)
                .arg("--output")
                .arg(output_md_file_path);

            spawn_command(cmd);
        }
    }

    let message = "Markdown generation complete!".bright_green().bold();
    println!("{}", message);
}

/// Removes "Loading... please wait."  
/// Changes "highlight" to "c" for proper syntax highlighting  
/// Prepends URL to the references in "See Also" section
fn postprocess_md_files(target_dir: &str) {
    let message = "Post-processing Markdown files:".bright_blue().bold();
    println!("{}", message);

    let md_pages_dir = std::fs::read_dir(target_dir);
    let md_pages = match md_pages_dir {
        Ok(dir) => dir,
        Err(e) => {
            let e = format!("Error: {}", e).red().bold();
            eprintln!("{}", e);

            std::process::exit(1);
        }
    };

    for entry in md_pages {
        if let Ok(page) = entry {
            println!("Post-processing {}", page.file_name().to_str().unwrap());

            let md_file_path = page.path();
            let md_file_content = std::fs::read_to_string(&md_file_path).unwrap();

            let mut new_content = md_file_content
                .replace("Loading… please wait.", "")
                .replace("highlight", "c");

            let regex = Regex::new(r"\[\w+\]\((\w+.html)\)").expect("Failed to create regex");

            let local_links: HashSet<String> = regex
                .captures_iter(&new_content)
                .map(|c| c[1].to_string())
                .collect();

            for link in local_links {
                new_content = new_content
                    .replace(&link, format!("{}{}", VK_SPEC_REGISTRY_URL, link).as_str());
            }

            std::fs::write(&md_file_path, new_content).unwrap();
        }
    }

    let message = "Post-processing complete!".bright_green().bold();
    println!("{}", message);
}

fn spawn_command(mut cmd: Command) {
    let result = cmd.spawn();

    match result {
        Ok(mut child) => {
            child.wait().expect("Failed to wait on command");
        }
        Err(e) => {
            if e.kind() == std::io::ErrorKind::NotFound {
                let e = format!(
                    "Couldn't find `{}`, make sure it's on your PATH",
                    cmd.get_program().to_str().unwrap()
                )
                .red()
                .bold();

                eprintln!("{}", e);
            } else {
                let e = format!("Error: {}", e).red().bold();

                eprintln!("{}", e);
            }

            std::process::exit(1);
        }
    }
}
