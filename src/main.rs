use structopt::StructOpt;

#[derive(StructOpt)]
struct Cli {
    command: String,
    #[structopt(parse(from_os_str))]
    args: std::path::PathBuf,
}

fn main() {
    let args = Cli::from_args();
}
