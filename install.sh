#!/bin/sh
# Install script for Binance Chain Tools
#   - bairdrop
#   - bdumper

# Note: this is based on current structure of `node-binary` repo, which is not optimal
# - The installer script is a hack to simplify the installation process
# - Our binaries should eventually be refactor into a `apt` or `npm` repo, which features upgradability
# - We should not rely on folders for addressing (instead use git branches for versions)

# Detect operating system
# Future Improvement: Refactor into helper function
if [[ "$OSTYPE" == "linux-gnu" ]]; then
  DETECTED_OS="linux"
elif [[ "$OSTYPE" == "darwin"* ]]; then
  DETECTED_OS="mac"
elif [[ "$OSTYPE" == "cygwin" ]]; then
  DETECTED_OS="linux"
elif [[ "$OSTYPE" == "msys" ]]; then
  DETECTED_OS="windows"
elif [[ "$OSTYPE" == "win32" ]]; then
  DETECTED_OS="windows" # TODO(Dan): can you run shell on windows?
elif [[ "$OSTYPE" == "freebsd"* ]]; then
  DETECTED_OS="linux"
else
  FULLNODE_echo "Error: unable to detect operating system. Please install manually by referring to $DOCS_WEB_LINK"
  LIGHTNODE_DOCS_WEB_LINK=""
  exit 1
fi

# Check for existence of wget
if [ ! -x /usr/bin/wget ]; then
  # some extra check if wget is not installed at the usual place
  command -v wget >/dev/null 2>&1 || {
    echo >&2 "Error: you need to have wget installed and in your path. Use brew (mac) or apt (unix) to install wget"
    exit 1
  }
fi

echo "@@@@@@@@@@@@@@@@@@@ @@@@@@@@@@@@@@@@@@@"
echo "@@@@@@@@@@@@@@@@@     @@@@@@@@@@@@@@@@@"
echo "@@@@@@@@@@@@@@@         @@@@@@@@@@@@@@@"
echo "@@@@@@@@@@@@@             @@@@@@@@@@@@@"
echo "@@@@@@@@@@@                 @@@@@@@@@@@"
echo "@@@@@@@@@         @@@         @@@@@@@@@"
echo "@@@@@@@@        @@@@@@@        @@@@@@@@"
echo "@@@@@@@@@@    @@@@@@@@@@@    @@@@@@@@@@"
echo "@@@   @@@@@@@@@@@@   @@@@@@@@@@@@   @@@"
echo "@       @@@@@@@@       @@@@@@@@       @"
echo "@       @@@@@@@@       @@@@@@@@       @"
echo "@@@   @@@@@@@@@@@@   @@@@@@@@@@@@   @@@"
echo "@@@@@@@@@@    @@@@@@@@@@@    @@@@@@@@@@"
echo "@@@@@@@@        @@@@@@@        @@@@@@@@"
echo "@@@@@@@@@         @@@         @@@@@@@@@"
echo "@@@@@@@@@@@                 @@@@@@@@@@@"
echo "@@@@@@@@@@@@@             @@@@@@@@@@@@@"
echo "@@@@@@@@@@@@@@@         @@@@@@@@@@@@@@@"
echo "@@@@@@@@@@@@@@@@@     @@@@@@@@@@@@@@@@@"
echo ""

echo "========== Binance Chain Node Installation =========="
echo "Installer Version: 0.1.beta"
echo "Detected OS: $DETECTED_OS"
echo "====================================================="

# Install location
USR_LOCAL_BIN="/usr/local/bin"
# Note: /usr/local/bin choice from https://unix.stackexchange.com/questions/259231/difference-between-usr-bin-and-usr-local-bin
# Future improvement: needs uninstall script (brew uninstall) that removes executable from bin

# Choose airdrop Directory
read -e -p "Choose home directory for airdrop [default: ~/.bairdrop]:" AIRDROP_DIR
AIRDROP_DIR=${AIRDROP_DIR:-"$HOME/.bairdrop"}

# Choose dumper directory
read -e -p "Choose home directory for balance dumper [default: ~/.bdumper]:" DUMPER_DIR
DUMPER_DIR=${DUMPER_DIR:-"$HOME/.bdumper"}

# Detect previous installation and create .bairdrop folder,
echo "... creating $AIRDROP_DIR"
if [ -d "$AIRDROP_DIR" ]; then
  echo "... Error: Binance Chain Airdrop has already been installed"
  echo "... Error: Please remove contents of ${AIRDROP_DIR} before reinstalling."
  exit 1
else
  mkdir -p $AIRDROP_DIR
fi

if [ -f "$USR_LOCAL_BIN/bairdrop" ]; then
  echo "... Error: Binance Chain Mainnet Airdrop  has already been installed"
  echo "... Error: Please remove bairdrop from /usr/local/bin before reinstalling."
  exit 1
fi
if [ -f "$USR_LOCAL_BIN/tbairdrop" ]; then
  echo "... Error: Binance Chain Testnet Airdrop has already been installed"
  echo "... Error: Please remove tbairdropfrom /usr/local/bin before reinstalling."
  exit 1
fi
if [ -f "$USR_LOCAL_BIN/bdumper" ]; then
  echo "... Error: Binance Chain Mainnet Balance Dumper has already been installed"
  echo "... Error: Please remove bdumper from /usr/local/bin before reinstalling."
  exit 1
fi
if [ -f "$USR_LOCAL_BIN/tbdumper" ]; then
  echo "... Error: Binance Chain Testnet Balance Dumper has already been installed"
  echo "... Error: Please remove tbdumper from /usr/local/bin before reinstalling."
  exit 1
fi

# File Download URLs
GH_REPO_URL="https://github.com/binance-chain/chain-tooling/raw/airdrop"

# Download both Testnet and Mainnet Airdrop
for NETWORK in "prod" "testnet"; do
  if [ "$NETWORK" = "prod" ]; then
    FILENAME="bairdrop"
  else
    FILENAME="tbairdrop"
  fi
  AD_VERSION_PATH="airdrop/$NETWORK/$DETECTED_OS/$FILENAME"
  AD_BINARY_URL="$GH_REPO_URL/$AD_VERSION_PATH"
  cd $USR_LOCAL_BIN
  echo "... Downloading $FILENAME executable:" $AD_BINARY_URL
  wget -q --show-progress "$AD_BINARY_URL"
  chmod 755 "./$FILENAME"
done

# Download both Testnet and Mainnet Dumper
for NETWORK in "prod" "testnet"; do
  if [ "$NETWORK" = "prod" ]; then
    FILENAME="bdumper"
  else
    FILENAME="tbdumper"
  fi
  BD_VERSION_PATH="balance-dumper/$NETWORK/$DETECTED_OS/$FILENAME"
  BD_BINARY_URL="$GH_REPO_URL/$BD_VERSION_PATH"
  cd $USR_LOCAL_BIN
  echo "... Downloading $FILENAME executable:" $BD_BINARY_URL
  wget -q --show-progress "$BD_BINARY_URL"
  chmod 755 "./$FILENAME"
done

# Add installed version of Binance Chain to path
echo "... Installation successful!"
echo "... \`bairdrop\`, \`tbairdrop\`, \`bdumper\`, \`tbdumper\` added to $USR_LOCAL_BIN"