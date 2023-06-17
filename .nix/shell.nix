# { pkgs ? import <nixpkgs> {} }:

# let gotk4-nix = pkgs.fetchFromGitHub {
		# owner = "diamondburned";
		# repo  = "gotk4-nix";
		# rev   = "b5bb50b31ffd7202decfdb8e84a35cbe88d42c64";
		# hash  = "sha256:18wxf24shsra5y5hdbxqcwaf3abhvx1xla8k0vnnkrwz0r9n4iqq";
	# };

# in import "${gotk4-nix}/shell.nix" {
	# base = {
		# # pname = "gotk4-adwaita";
		# pname = "gtk-layer-shell";
		# version = "dev";

		# buildInputs = pkgs: with pkgs; [
            # gtk-layer-shell
            # gtk4-layer-shell
			# gobject-introspection
			# glib
			# graphene
			# gdk-pixbuf
			# gtk4
			# gtk3
			# vulkan-headers
			# libadwaita
		# ];
	# };
# }
let systemPkgs = import <nixpkgs> {};

in {
	gotk4 ? (systemPkgs.fetchFromGitHub {
		owner = "brotholo";
		repo  = "gotk4";
        rev = "09b88ab1da1224390123ee22758410d93c9694f4";
        hash = "sha256:IC/XtqnFhsHi3biIOe+77rMdjMsYKn6lvN7upcTdn9M=";
        # rev = "297248f1acafec14cbd494e145ce10b032161cdf";
        # hash = "sha256:owo8Xp9yVqPIs1eU+2lQVmMv0DSZi3vSValo1+TG+NA=";
        # rev   = "fd960d20b525a07580938d10a214336bafb47d12";
		# hash  = "sha256:0zijivbyjfbb2vda05vpvq268i7vx9bhzlbzzsa4zfzzr9427w66";
	}),

	pkgs ? (import "${gotk4}/.nix/pkgs.nix" {
		src = systemPkgs.fetchFromGitHub {
			# owner = "NixOS";
			# repo  = "nixpkgs";
			# rev   = "3fdd780";
			# hash  = "sha256:0df9v2snlk9ag7jnmxiv31pzhd0rqx2h3kzpsxpj07xns8k8dghz";
            owner = "brotholo";
            repo  = "nixpkgs";
            # rev  = "5c905089d8a127df647791f7b4b4985462b1fe47";
            # rev = "cf7ab33478d7a8e53d25b34ecb462fd2b3c3d52c";
            rev  = "feb5f65c13014db0ada62b8b0dc8e8a433fefed1";
            hash  = "sha256:NSSOx/wuslUCXcuQRpeaLuRD5S3mMAxsIv6x2OCXZBk=";
            # hash = "sha256:U0wMJdcDIfSrPKbGxilVEsEZQ5pFvEVtEtGGr8JUNMs=";
            # rev = "462116c86a32bc8274733f542497d82eda53b445";
            # hash = "sha256:/Jb1hoOoEsaTe1ZO2dQrpAhHPSfHURPAyno0jqpOMSU=";
            # rev = "bb00692596b96447458b0d57d2d73c4057a68295";
            # hash = "sha256:ms6fITTzKyxfrjGMjO8MFqGgnZaNqy3uH0IZ6/3KttY=";
            # hash = "sha256:MEZmb3sX/eSoVNPYVBnoRaX1Yvd9z98EpUYz2fwdW18=";
            # rev   = "147b03fa8ebf9d5d5f6784f87dc61f0e7beee911"; # release-21.11
            # rev    = "5ff511f2f83d0d4ec475943ec18f66dbba541130";
            # hash   = "sha256:Kg9FscbReso3YFQxsfSQx1CmOQ2ARAlkSeHuSkVDAn8=";
		};
	}),
}:

in import "${gotk4}/.nix/shell.nix".pkgs {
	buildInputs = old.buildInputs ++ (with pkgs; [
		gtk4-layer-shell
	]);
}
# let shell = import "${gotk4}/.nix/shell.nix" {
	# inherit pkgs;
# };

# in shell.overrideAttrs (old: {
	# buildInputs = old.buildInputs ++ (with pkgs; [
		# gtk4-layer-shell
	# ]);
# })

