{
  description = "valorant prometheus exporter";

  inputs = { flake-utils.url = "github:numtide/flake-utils"; };

  outputs = { self, nixpkgs, flake-utils, ... }:

    {
      nixosModules.default = self.nixosModules.valorant-exporter;
      nixosModules.valorant-exporter = { lib, pkgs, config, ... }:
        with lib;

        let cfg = config.services.valorant-exporter;
        in
        {

          options.services.valorant-exporter = {

            enable = mkEnableOption "valorant-exporter";

            configure-prometheus = mkEnableOption "enable valorant-exporter in prometheus";

            port = mkOption {
              type = types.str;
              default = "1091";
              description = "Port under which valorant-exporter is accessible.";
            };

            listen = mkOption {
              type = types.str;
              default = "localhost";
              example = "127.0.0.1";
              description = "Address under which valorant-exporter is accessible.";
            };

            targets = mkOption {
              type = types.listOf types.str;
              default = [ "mayniklas/niki" ];
              description = "valorant players to monitor";
            };

            user = mkOption {
              type = types.str;
              default = "valorant-exporter";
              description = "User account under which valorant-exporter runs.";
            };

            group = mkOption {
              type = types.str;
              default = "valorant-exporter";
              description = "Group under which valorant-exporter runs.";
            };

          };

          config = mkIf cfg.enable {

            systemd.services.valorant-exporter = {
              description = "A valorant metrics exporter";
              wantedBy = [ "multi-user.target" ];
              serviceConfig = mkMerge [{
                User = cfg.user;
                Group = cfg.group;
                ExecStart = "${self.packages."${pkgs.system}".valorant-exporter}/bin/valorant-exporter -port ${cfg.port} -listen ${cfg.listen}";
                Restart = "on-failure";
              }];
            };

            users.users = mkIf (cfg.user == "valorant-exporter") {
              valorant-exporter = {
                isSystemUser = true;
                group = cfg.group;
                description = "valorant-exporter system user";
              };
            };

            users.groups =
              mkIf (cfg.group == "valorant-exporter") { valorant-exporter = { }; };

            services.prometheus = mkIf cfg.configure-prometheus {
              scrapeConfigs = [{
                job_name = "valorant";
                scrape_interval = "15m";
                metrics_path = "/probe";
                static_configs = [{ targets = cfg.targets; }];
                relabel_configs = [
                  {
                    source_labels = [ "__address__" ];
                    target_label = "__param_target";
                  }
                  {
                    source_labels = [ "__param_target" ];
                    target_label = "instance";
                  }
                  {
                    target_label = "__address__";
                    replacement =
                      "127.0.0.1:${cfg.port}";
                  }
                ];
              }];
            };

          };
          meta = { maintainers = with lib.maintainers; [ mayniklas ]; };
        };
    }

    //

    flake-utils.lib.eachDefaultSystem (system:
      let pkgs = nixpkgs.legacyPackages.${system};

      in
      rec {

        formatter = pkgs.nixpkgs-fmt;
        packages = flake-utils.lib.flattenTree rec {

          default = valorant-exporter;

          valorant-exporter = pkgs.buildGoModule rec {
            pname = "valorant-exporter";
            version = "1.0.0";
            src = self;
            vendorSha256 =
              "sha256-WtO+3uH6H2um6pcdqhU/Yaw6HDNkz1XGjslGQphyMiA=";
            installCheckPhase = ''
              runHook preCheck
              $out/bin/valorant-exporter -h
              runHook postCheck
            '';
            doCheck = false;
            meta = with pkgs.lib; {
              description = "valorant exporter";
              homepage =
                "https://github.com/MayNiklas/valorant-exporter";
              platforms = platforms.unix;
              maintainers = with maintainers; [ mayniklas ];
            };
          };

        };
      });
}
