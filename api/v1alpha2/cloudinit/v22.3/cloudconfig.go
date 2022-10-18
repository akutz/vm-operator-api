// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    cloudConfig, err := UnmarshalCloudConfig(bytes)
//    bytes, err = cloudConfig.Marshal()

package cloudinit

import (
	"bytes"
	"encoding/json"
	"errors"
)

func UnmarshalCloudConfig(data []byte) (CloudConfig, error) {
	var r CloudConfig
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CloudConfig) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CloudConfig struct {
	// +optional
	Ansible *Ansible `json:"ansible,omitempty"`
	// +optional
	ApkRepos *ApkRepos `json:"apk_repos,omitempty"`
	// +optional
	Apt *Apt `json:"apt,omitempty"`
	// +optional
	AptPipelining *AptPipeliningUnion `json:"apt_pipelining"`
	// +optional
	Autoinstall *Autoinstall `json:"autoinstall,omitempty"` // Opaque autoinstall schema definition for Ubuntu autoinstall. Full schema processed by; live-installer. See: https://ubuntu.com/server/docs/install/autoinstall-reference
	// +optional
	Bootcmd []Cmd `json:"bootcmd,omitempty"`
	// +optional
	ByobuByDefault *ByobuByDefault `json:"byobu_by_default,omitempty"`
	// +optional
	CACerts *CACerts `json:"ca-certs,omitempty"`
	// +optional
	CloudConfigCACerts *CACerts `json:"ca_certs,omitempty"`
	// +optional
	Chef *Chef `json:"chef,omitempty"`
	// +optional
	DisableEc2Metadata *bool `json:"disable_ec2_metadata,omitempty"` // Set true to disable IPv4 routes to EC2 metadata. Default: false.
	// +optional
	DeviceAliases *DeviceAliases `json:"device_aliases,omitempty"`
	// +optional
	DiskSetup *DiskSetup `json:"disk_setup,omitempty"`
	// +optional
	FSSetup []FSSetup `json:"fs_setup,omitempty"`
	// +optional
	Fan *Fan `json:"fan,omitempty"`
	// +optional
	FinalMessage *string `json:"final_message,omitempty"` // The message to display at the end of the run
	// +optional
	Growpart *GrowpartClass `json:"growpart,omitempty"`
	// +optional
	GrubDpkg map[string]interface{} `json:"grub-dpkg,omitempty"` // Use ``grub_dpkg`` instead
	// +optional
	CloudConfigGrubDpkg *GrubDpkg `json:"grub_dpkg,omitempty"`
	// +optional
	Updates *Updates `json:"updates,omitempty"`
	// +optional
	Keyboard *Keyboard `json:"keyboard,omitempty"`
	// +optional
	SSH *SSH `json:"ssh,omitempty"`
	// +optional
	SSHFPConsoleBlacklist []string `json:"ssh_fp_console_blacklist,omitempty"` // Avoid printing matching SSH fingerprints to the system console.
	// +optional
	SSHKeyConsoleBlacklist []string `json:"ssh_key_console_blacklist,omitempty"` // Avoid printing matching SSH key types to the system console.
	// +optional
	Landscape *Landscape `json:"landscape,omitempty"`
	// +optional
	Locale *string `json:"locale,omitempty"` // The locale to set as the system's locale (e.g. ar_PS)
	// +optional
	LocaleConfigfile *string `json:"locale_configfile,omitempty"` // The file in which to write the locale configuration (defaults to the distro's default; location)
	// +optional
	Lxd *Lxd `json:"lxd,omitempty"`
	// +optional
	Mcollective *Mcollective `json:"mcollective,omitempty"`
	// +optional
	Migrate *bool `json:"migrate,omitempty"` // Whether to migrate legacy cloud-init semaphores to new format. Default: ``true``
	// +optional
	MountDefaultFields []*string `json:"mount_default_fields,omitempty"` // Default mount configuration for any mount entry with less than 6 options provided. When; specified, 6 items are required and represent ``/etc/fstab`` entries. Default:; ``defaults,nofail,x-systemd.requires=cloud-init.service,_netdev``
	// +optional
	Mounts [][]string `json:"mounts,omitempty"` // List of lists. Each inner list entry is a list of ``/etc/fstab`` mount declarations of; the format: [ fs_spec, fs_file, fs_vfstype, fs_mntops, fs-freq, fs_passno ]. A mount; declaration with less than 6 items will get remaining values from; ``mount_default_fields``. A mount declaration with only `fs_spec` and no `fs_file`; mountpoint will be skipped.
	// +optional
	Swap *Swap `json:"swap,omitempty"`
	// +optional
	NTP *NTP `json:"ntp"`
	// +optional
	AptRebootIfRequired *bool `json:"apt_reboot_if_required,omitempty"` // Dropped after April 2027. Use ``package_reboot_if_required``. Default: ``false``
	// +optional
	AptUpdate *bool `json:"apt_update,omitempty"` // Dropped after April 2027. Use ``package_update``. Default: ``false``
	// +optional
	AptUpgrade *bool `json:"apt_upgrade,omitempty"` // Dropped after April 2027. Use ``package_upgrade``. Default: ``false``
	// +optional
	PackageRebootIfRequired *bool `json:"package_reboot_if_required,omitempty"` // Set ``true`` to reboot the system if required by presence of `/var/run/reboot-required`.; Default: ``false``
	// +optional
	PackageUpdate *bool `json:"package_update,omitempty"` // Set ``true`` to update packages. Happens before upgrade or install. Default: ``false``
	// +optional
	PackageUpgrade *bool `json:"package_upgrade,omitempty"` // Set ``true`` to upgrade packages. Happens before install. Default: ``false``
	// +optional
	Packages []Cmd `json:"packages,omitempty"` // A list of packages to install. Each entry in the list can be either a package name or a; list with two entries, the first being the package name and the second being the specific; package version to install.
	// +optional
	PhoneHome *PhoneHome `json:"phone_home,omitempty"`
	// +optional
	PowerState *PowerState `json:"power_state,omitempty"`
	// +optional
	Puppet *Puppet `json:"puppet,omitempty"`
	// +optional
	ResizeRootfs *ResizeRootfsUnion `json:"resize_rootfs"` // Whether to resize the root partition. ``noblock`` will resize in the background. Default:; ``true``
	// +optional
	ManageResolvConf *bool `json:"manage_resolv_conf,omitempty"` // Whether to manage the resolv.conf file. ``resolv_conf`` block will be ignored unless this; is set to ``true``. Default: ``false``
	// +optional
	ResolvConf *ResolvConf `json:"resolv_conf,omitempty"`
	// +optional
	RhSubscription *RhSubscription `json:"rh_subscription,omitempty"`
	// +optional
	Rsyslog *Rsyslog `json:"rsyslog,omitempty"`
	// +optional
	Runcmd []Runcmd `json:"runcmd,omitempty"`
	// +optional
	SaltMinion *SaltMinion `json:"salt_minion,omitempty"`
	// +optional
	VendorData *VendorData `json:"vendor_data,omitempty"`
	// +optional
	RandomSeed *RandomSeed `json:"random_seed,omitempty"`
	// +optional
	FQDN *string `json:"fqdn,omitempty"` // The fully qualified domain name to set; ; Optional fully qualified domain name to use when updating ``/etc/hosts``. Preferred over; ``hostname`` if both are provided. In absence of ``hostname`` and ``fqdn`` in; cloud-config, the ``local-hostname`` value will be used from datasource metadata.
	// +optional
	Hostname *string `json:"hostname,omitempty"` // The hostname to set; ; Hostname to set when rendering ``/etc/hosts``. If ``fqdn`` is set, the hostname extracted; from ``fqdn`` overrides ``hostname``.
	// +optional
	PreferFQDNOverHostname *bool `json:"prefer_fqdn_over_hostname,omitempty"` // If true, the fqdn will be used if it is set. If false, the hostname will be used. If; unset, the result is distro-dependent; ; By default, it is distro-dependent whether cloud-init uses the short hostname or fully; qualified domain name when both ``local-hostname` and ``fqdn`` are both present in; instance metadata. When set ``true``, use fully qualified domain name if present as; hostname instead of short hostname. When set ``false``, use ``hostname`` config value if; present, otherwise fallback to ``fqdn``.
	// +optional
	PreserveHostname *bool `json:"preserve_hostname,omitempty"` // If true, the hostname will not be changed. Default: ``false``; ; Do not update system hostname when ``true``. Default: ``false``.
	// +optional
	Chpasswd *Chpasswd `json:"chpasswd,omitempty"`
	// +optional
	Password *CloudConfigPassword `json:"password"` // Set the default user's password. Ignored if ``chpasswd`` ``list`` is used
	// +optional
	SSHPwauth *GrubPCInstallDevicesEmpty `json:"ssh_pwauth"` // Sets whether or not to accept password authentication. ``true`` will enable password; auth. ``false`` will disable. Default is to leave the value unchanged.
	// +optional
	Snap *Snap `json:"snap,omitempty"`
	// +optional
	Spacewalk *Spacewalk `json:"spacewalk,omitempty"`
	// +optional
	AuthkeyHash *string `json:"authkey_hash,omitempty"` // The hash type to use when generating SSH fingerprints. Default: ``sha256``
	// +optional
	NoSSHFingerprints *bool `json:"no_ssh_fingerprints,omitempty"` // If true, SSH fingerprints will not be written. Default: ``false``
	// +optional
	SSHImportID []string `json:"ssh_import_id,omitempty"`
	// +optional
	AllowPublicSSHKeys *bool `json:"allow_public_ssh_keys,omitempty"` // If ``true``, will import the public SSH keys from the datasource's metadata to the user's; ``.ssh/authorized_keys`` file. Default: ``true``
	// +optional
	DisableRoot *bool `json:"disable_root,omitempty"` // Disable root login. Default: ``true``
	// +optional
	DisableRootOpts *string `json:"disable_root_opts,omitempty"` // Disable root login options.  If ``disable_root_opts`` is specified and contains the; string ``$USER``, it will be replaced with the username of the default user. Default:; ``no-port-forwarding,no-agent-forwarding,no-X11-forwarding,command="echo 'Please login as; the user \"$USER\" rather than the user \"$DISABLE_USER\".';echo;sleep 10;exit 142"``
	// +optional
	SSHAuthorizedKeys []string `json:"ssh_authorized_keys,omitempty"` // The SSH public keys to add ``.ssh/authorized_keys`` in the default user's home directory
	// +optional
	SSHDeletekeys *bool `json:"ssh_deletekeys,omitempty"` // Remove host SSH keys. This prevents re-use of a private host key from an image with; default host SSH keys. Default: ``true``
	// +optional
	SSHGenkeytypes []SSHGenkeytype `json:"ssh_genkeytypes,omitempty"` // The SSH key types to generate. Default: ``[rsa, dsa, ecdsa, ed25519]``
	// +optional
	SSHKeys *SSHKeys `json:"ssh_keys,omitempty"` // A dictionary entries for the public and private host keys of each desired key type.; Entries in the ``ssh_keys`` config dict should have keys in the format ``<key; type>_private``, ``<key type>_public``, and, optionally, ``<key type>_certificate``, e.g.; ``rsa_private: <key>``, ``rsa_public: <key>``, and ``rsa_certificate: <key>``. Not all; key types have to be specified, ones left unspecified will not be used. If this config; option is used, then separate keys will not be automatically generated. In order to; specify multiline private host keys and certificates, use yaml multiline syntax.
	// +optional
	SSHPublishHostkeys *SSHPublishHostkeys `json:"ssh_publish_hostkeys,omitempty"`
	// +optional
	SSHQuietKeygen *bool `json:"ssh_quiet_keygen,omitempty"` // If ``true``, will suppress the output of key generation to the console. Default: ``false``
	// +optional
	Timezone *string `json:"timezone,omitempty"` // The timezone to use as represented in /usr/share/zoneinfo
	// +optional
	UbuntuAdvantage *UbuntuAdvantage `json:"ubuntu_advantage,omitempty"`
	// +optional
	Drivers *Drivers `json:"drivers,omitempty"`
	// +optional
	ManageEtcHosts *ManageEtcHostsUnion `json:"manage_etc_hosts"` // Whether to manage ``/etc/hosts`` on the system. If ``true``, render the hosts file using; ``/etc/cloud/templates/hosts.tmpl`` replacing ``$hostname`` and ``$fdqn``. If; ``localhost``, append a ``127.0.1.1`` entry that resolves from FQDN and hostname every; boot. Default: ``false``.
	// +optional
	Groups *CloudConfigGroups `json:"groups"`
	// +optional
	User *CloudConfigUser `json:"user"` // The ``user`` dictionary values override the ``default_user`` configuration from; ``/etc/cloud/cloud.cfg``. The `user` dictionary keys supported for the default_user are; the same as the ``users`` schema.
	// +optional
	Users *Users `json:"users"`
	// +optional
	Wireguard *Wireguard `json:"wireguard"`
	// +optional
	WriteFiles []WriteFile `json:"write_files,omitempty"`
	// +optional
	YumRepoDir *string `json:"yum_repo_dir,omitempty"` // The repo parts directory where individual yum repo config files will be written. Default:; ``/etc/yum.repos.d``
	// +optional
	YumRepos *YumRepos `json:"yum_repos,omitempty"`
	// +optional
	Zypper *Zypper `json:"zypper,omitempty"`
	// +optional
	Reporting *Reporting `json:"reporting,omitempty"`
}

type Ansible struct {
	// +optional
	InstallMethod *InstallMethod `json:"install-method,omitempty"` // The type of installation for ansible. It can be one of the following values:; ; - ``distro``; - ``pip``
	// +optional
	PackageName *string `json:"package-name,omitempty"`
	// +optional
	Pull *Pull `json:"pull,omitempty"`
}

type Pull struct {
	// +optional
	AcceptHostKey *bool `json:"accept-host-key,omitempty"`
	// +optional
	Checkout *string `json:"checkout,omitempty"`
	// +optional
	Clean *bool `json:"clean,omitempty"`
	// +optional
	Connection *string `json:"connection,omitempty"`
	// +optional
	Diff *bool `json:"diff,omitempty"`
	// +optional
	Full *bool `json:"full,omitempty"`
	// +optional
	ModuleName *string `json:"module-name,omitempty"`
	// +optional
	ModulePath *string `json:"module-path,omitempty"`
	// +optional
	PlaybookName string `json:"playbook-name"`
	// +optional
	PrivateKey *string `json:"private-key,omitempty"`
	// +optional
	SCPExtraArgs *string `json:"scp-extra-args,omitempty"`
	// +optional
	SFTPExtraArgs *string `json:"sftp-extra-args,omitempty"`
	// +optional
	SkipTags *string `json:"skip-tags,omitempty"`
	// +optional
	Sleep *string `json:"sleep,omitempty"`
	// +optional
	SSHCommonArgs *string `json:"ssh-common-args,omitempty"`
	// +optional
	Tags *string `json:"tags,omitempty"`
	// +optional
	Timeout *string `json:"timeout,omitempty"`
	// +optional
	URL string `json:"url"`
	// +optional
	VaultID *string `json:"vault-id,omitempty"`
	// +optional
	VaultPasswordFile *string `json:"vault-password-file,omitempty"`
}

type ApkRepos struct {
	// +optional
	AlpineRepo *AlpineRepo `json:"alpine_repo"`
	// +optional
	LocalRepoBaseURL *string `json:"local_repo_base_url,omitempty"` // The base URL of an Alpine repository containing unofficial packages
	// +optional
	PreserveRepositories *bool `json:"preserve_repositories,omitempty"` // By default, cloud-init will generate a new repositories file ``/etc/apk/repositories``; based on any valid configuration settings specified within a apk_repos section of cloud; config. To disable this behavior and preserve the repositories file from the pristine; image, set ``preserve_repositories`` to ``true``.; ; The ``preserve_repositories`` option overrides all other config keys that would alter; ``/etc/apk/repositories``.
}

type AlpineRepo struct {
	// +optional
	BaseURL *string `json:"base_url,omitempty"` // The base URL of an Alpine repository, or mirror, to download official packages from. If; not specified then it defaults to ``https://alpine.global.ssl.fastly.net/alpine``
	// +optional
	CommunityEnabled *bool `json:"community_enabled,omitempty"` // Whether to add the Community repo to the repositories file. By default the Community repo; is not included.
	// +optional
	TestingEnabled *bool `json:"testing_enabled,omitempty"` // Whether to add the Testing repo to the repositories file. By default the Testing repo is; not included. It is only recommended to use the Testing repo on a machine running the; ``Edge`` version of Alpine as packages installed from Testing may have dependencies that; conflict with those in non-Edge Main or Community repos.
	// +optional
	Version string `json:"version"` // The Alpine version to use (e.g. ``v3.12`` or ``edge``)
}

type Apt struct {
	// +optional
	AddAptRepoMatch *string `json:"add_apt_repo_match,omitempty"` // All source entries in ``apt-sources`` that match regex in ``add_apt_repo_match`` will be; added to the system using ``add-apt-repository``. If ``add_apt_repo_match`` is not; specified, it defaults to ``^[\w-]+:\w``
	// +optional
	Conf *string `json:"conf,omitempty"` // Specify configuration for apt, such as proxy configuration. This configuration is; specified as a string. For multiline APT configuration, make sure to follow yaml syntax.
	// +optional
	DebconfSelections *DebconfSelections `json:"debconf_selections,omitempty"` // Debconf additional configurations can be specified as a dictionary under the; ``debconf_selections`` config key, with each key in the dict representing a different set; of configurations. The value of each key must be a string containing all the debconf; configurations that must be applied. We will bundle all of the values and pass them to; ``debconf-set-selections``. Therefore, each value line must be a valid entry for; ``debconf-set-selections``, meaning that they must possess for distinct fields:; ; ``pkgname question type answer``; ; Where:; ; - ``pkgname`` is the name of the package.; - ``question`` the name of the questions.; - ``type`` is the type of question.; - ``answer`` is the value used to answer the question.; ; For example: ``ippackage ippackage/ip string 127.0.01``
	// +optional
	DisableSuites []string `json:"disable_suites,omitempty"` // Entries in the sources list can be disabled using ``disable_suites``, which takes a list; of suites to be disabled. If the string ``$RELEASE`` is present in a suite in the; ``disable_suites`` list, it will be replaced with the release name. If a suite specified; in ``disable_suites`` is not present in ``sources.list`` it will be ignored. For; convenience, several aliases are provided for`` disable_suites``:; ; - ``updates`` => ``$RELEASE-updates``; - ``backports`` => ``$RELEASE-backports``; - ``security`` => ``$RELEASE-security``; - ``proposed`` => ``$RELEASE-proposed``; - ``release`` => ``$RELEASE``.; ; When a suite is disabled using ``disable_suites``, its entry in ``sources.list`` is not; deleted; it is just commented out.
	// +optional
	FTPProxy *string `json:"ftp_proxy,omitempty"` // More convenient way to specify ftp APT proxy. ftp proxy url is specified in the format; ``ftp://[[user][:pass]@]host[:port]/``.
	// +optional
	HTTPProxy *string `json:"http_proxy,omitempty"` // More convenient way to specify http APT proxy. http proxy url is specified in the format; ``http://[[user][:pass]@]host[:port]/``.
	// +optional
	HTTPSProxy *string `json:"https_proxy,omitempty"` // More convenient way to specify https APT proxy. https proxy url is specified in the; format ``https://[[user][:pass]@]host[:port]/``.
	// +optional
	PreserveSourcesList *bool `json:"preserve_sources_list,omitempty"` // By default, cloud-init will generate a new sources list in ``/etc/apt/sources.list.d``; based on any changes specified in cloud config. To disable this behavior and preserve the; sources list from the pristine image, set ``preserve_sources_list`` to ``true``.; ; The ``preserve_sources_list`` option overrides all other config keys that would alter; ``sources.list`` or ``sources.list.d``, **except** for additional sources to be added to; ``sources.list.d``.
	// +optional
	Primary []PrimaryElement `json:"primary,omitempty"` // The primary and security archive mirrors can be specified using the ``primary`` and; ``security`` keys, respectively. Both the ``primary`` and ``security`` keys take a list; of configs, allowing mirrors to be specified on a per-architecture basis. Each config is; a dictionary which must have an entry for ``arches``, specifying which architectures that; config entry is for. The keyword ``default`` applies to any architecture not explicitly; listed. The mirror url can be specified with the ``uri`` key, or a list of mirrors to; check can be provided in order, with the first mirror that can be resolved being; selected. This allows the same configuration to be used in different environment, with; different hosts used for a local APT mirror. If no mirror is provided by ``uri`` or; ``search``, ``search_dns`` may be used to search for dns names in the format; ``<distro>-mirror`` in each of the following:; ; - fqdn of this host per cloud metadata,; - localdomain,; - domains listed in ``/etc/resolv.conf``.; ; If there is a dns entry for ``<distro>-mirror``, then it is assumed that there is a; distro mirror at ``http://<distro>-mirror.<domain>/<distro>``. If the ``primary`` key is; defined, but not the ``security`` key, then then configuration for ``primary`` is also; used for ``security``. If ``search_dns`` is used for the ``security`` key, the search; pattern will be ``<distro>-security-mirror``.; ; Each mirror may also specify a key to import via any of the following optional keys:; ; - ``keyid``: a key to import via shortid or fingerprint.; - ``key``: a raw PGP key.; - ``keyserver``: alternate keyserver to pull ``keyid`` key from.; ; If no mirrors are specified, or all lookups fail, then default mirrors defined in the; datasource are used. If none are present in the datasource either the following defaults; are used:; ; - ``primary`` => ``http://archive.ubuntu.com/ubuntu``.; - ``security`` => ``http://security.ubuntu.com/ubuntu``
	// +optional
	Proxy *string `json:"proxy,omitempty"` // Alias for defining a http APT proxy.
	// +optional
	Security []PrimaryElement `json:"security,omitempty"` // Please refer to the primary config documentation
	// +optional
	Sources *Sources `json:"sources,omitempty"` // Source list entries can be specified as a dictionary under the ``sources`` config key,; with each key in the dict representing a different source file. The key of each source; entry will be used as an id that can be referenced in other config entries, as well as; the filename for the source's configuration under ``/etc/apt/sources.list.d``. If the; name does not end with ``.list``, it will be appended. If there is no configuration for a; key in ``sources``, no file will be written, but the key may still be referred to as an; id in other ``sources`` entries.; ; Each entry under ``sources`` is a dictionary which may contain any of the following; optional keys:; - ``source``: a sources.list entry (some variable replacements apply).; - ``keyid``: a key to import via shortid or fingerprint.; - ``key``: a raw PGP key.; - ``keyserver``: alternate keyserver to pull ``keyid`` key from.; - ``filename``: specify the name of the list file; ; The ``source`` key supports variable replacements for the following strings:; ; - ``$MIRROR``; - ``$PRIMARY``; - ``$SECURITY``; - ``$RELEASE``; - ``$KEY_FILE``
	// +optional
	SourcesList *string `json:"sources_list,omitempty"` // Specifies a custom template for rendering ``sources.list`` . If no ``sources_list``; template is given, cloud-init will use sane default. Within this template, the following; strings will be replaced with the appropriate values:; ; - ``$MIRROR``; - ``$RELEASE``; - ``$PRIMARY``; - ``$SECURITY``; - ``$KEY_FILE``
}

// Debconf additional configurations can be specified as a dictionary under the
// ``debconf_selections`` config key, with each key in the dict representing a different set
// of configurations. The value of each key must be a string containing all the debconf
// configurations that must be applied. We will bundle all of the values and pass them to
// ``debconf-set-selections``. Therefore, each value line must be a valid entry for
// ``debconf-set-selections``, meaning that they must possess for distinct fields:
//
// ``pkgname question type answer``
//
// Where:
//
// - ``pkgname`` is the name of the package.
// - ``question`` the name of the questions.
// - ``type`` is the type of question.
// - ``answer`` is the value used to answer the question.
//
// For example: ``ippackage ippackage/ip string 127.0.01``
type DebconfSelections struct {
}

// The primary and security archive mirrors can be specified using the ``primary`` and
// ``security`` keys, respectively. Both the ``primary`` and ``security`` keys take a list
// of configs, allowing mirrors to be specified on a per-architecture basis. Each config is
// a dictionary which must have an entry for ``arches``, specifying which architectures that
// config entry is for. The keyword ``default`` applies to any architecture not explicitly
// listed. The mirror url can be specified with the ``uri`` key, or a list of mirrors to
// check can be provided in order, with the first mirror that can be resolved being
// selected. This allows the same configuration to be used in different environment, with
// different hosts used for a local APT mirror. If no mirror is provided by ``uri`` or
// ``search``, ``search_dns`` may be used to search for dns names in the format
// ``<distro>-mirror`` in each of the following:
//
// - fqdn of this host per cloud metadata,
// - localdomain,
// - domains listed in ``/etc/resolv.conf``.
//
// If there is a dns entry for ``<distro>-mirror``, then it is assumed that there is a
// distro mirror at ``http://<distro>-mirror.<domain>/<distro>``. If the ``primary`` key is
// defined, but not the ``security`` key, then then configuration for ``primary`` is also
// used for ``security``. If ``search_dns`` is used for the ``security`` key, the search
// pattern will be ``<distro>-security-mirror``.
//
// Each mirror may also specify a key to import via any of the following optional keys:
//
// - ``keyid``: a key to import via shortid or fingerprint.
// - ``key``: a raw PGP key.
// - ``keyserver``: alternate keyserver to pull ``keyid`` key from.
//
// If no mirrors are specified, or all lookups fail, then default mirrors defined in the
// datasource are used. If none are present in the datasource either the following defaults
// are used:
//
// - ``primary`` => ``http://archive.ubuntu.com/ubuntu``.
// - ``security`` => ``http://security.ubuntu.com/ubuntu``
type PrimaryElement struct {
	// +optional
	Arches []string `json:"arches"`
	// +optional
	Key *string `json:"key,omitempty"`
	// +optional
	Keyid *string `json:"keyid,omitempty"`
	// +optional
	Keyserver *string `json:"keyserver,omitempty"`
	// +optional
	Search []string `json:"search,omitempty"`
	// +optional
	SearchDNS *bool `json:"search_dns,omitempty"`
	// +optional
	URI *string `json:"uri,omitempty"`
}

// Source list entries can be specified as a dictionary under the ``sources`` config key,
// with each key in the dict representing a different source file. The key of each source
// entry will be used as an id that can be referenced in other config entries, as well as
// the filename for the source's configuration under ``/etc/apt/sources.list.d``. If the
// name does not end with ``.list``, it will be appended. If there is no configuration for a
// key in ``sources``, no file will be written, but the key may still be referred to as an
// id in other ``sources`` entries.
//
// Each entry under ``sources`` is a dictionary which may contain any of the following
// optional keys:
// - ``source``: a sources.list entry (some variable replacements apply).
// - ``keyid``: a key to import via shortid or fingerprint.
// - ``key``: a raw PGP key.
// - ``keyserver``: alternate keyserver to pull ``keyid`` key from.
// - ``filename``: specify the name of the list file
//
// The ``source`` key supports variable replacements for the following strings:
//
// - ``$MIRROR``
// - ``$PRIMARY``
// - ``$SECURITY``
// - ``$RELEASE``
// - ``$KEY_FILE``
type Sources struct {
}

// Opaque autoinstall schema definition for Ubuntu autoinstall. Full schema processed by
// live-installer. See: https://ubuntu.com/server/docs/install/autoinstall-reference
type Autoinstall struct {
	// +optional
	Version int64 `json:"version"`
}

// Dropped after April 2027. Use ``ca_certs``.
type CACerts struct {
	// +optional
	RemoveDefaults *bool `json:"remove-defaults,omitempty"` // Dropped after April 2027. Use ``remove_defaults``.
	// +optional
	CACertsRemoveDefaults *bool `json:"remove_defaults,omitempty"` // Remove default CA certificates if true. Default: false
	// +optional
	Trusted []string `json:"trusted,omitempty"` // List of trusted CA certificates to add.
}

type Chef struct {
	// +optional
	ChefLicense *string `json:"chef_license,omitempty"` // string that indicates if user accepts or not license related to some of chef products
	// +optional
	ClientKey *string `json:"client_key,omitempty"` // Optional path for client_cert. Default to ``/etc/chef/client.pem``.
	// +optional
	Directories []string `json:"directories,omitempty"` // Create the necessary directories for chef to run. By default, it creates the following; directories:; ; - ``/etc/chef``; - ``/var/log/chef``; - ``/var/lib/chef``; - ``/var/cache/chef``; - ``/var/backups/chef``; - ``/var/run/chef``
	// +optional
	EncryptedDataBagSecret *string `json:"encrypted_data_bag_secret,omitempty"` // Specifies the location of the secret key used by chef to encrypt data items. By default,; this path is set to null, meaning that chef will have to look at the path; ``/etc/chef/encrypted_data_bag_secret`` for it.
	// +optional
	Environment *string `json:"environment,omitempty"` // Specifies which environment chef will use. By default, it will use the ``_default``; configuration.
	// +optional
	Exec *bool `json:"exec,omitempty"` // Set true if we should run or not run chef (defaults to false, unless a gem installed is; requested where this will then default to true).
	// +optional
	FileBackupPath *string `json:"file_backup_path,omitempty"` // Specifies the location in which backup files are stored. By default, it uses the; ``/var/backups/chef`` location.
	// +optional
	FileCachePath *string `json:"file_cache_path,omitempty"` // Specifies the location in which chef cache files will be saved. By default, it uses the; ``/var/cache/chef`` location.
	// +optional
	FirstbootPath *string `json:"firstboot_path,omitempty"` // Path to write run_list and initial_attributes keys that should also be present in this; configuration, defaults to ``/etc/chef/firstboot.json``
	// +optional
	ForceInstall *bool `json:"force_install,omitempty"` // If set to ``true``, forces chef installation, even if it is already installed.
	// +optional
	InitialAttributes map[string]interface{} `json:"initial_attributes,omitempty"` // Specify a list of initial attributes used by the cookbooks.
	// +optional
	InstallType *ChefInstallType `json:"install_type,omitempty"` // The type of installation for chef. It can be one of the following values:; ; - ``packages``; - ``gems``; - ``omnibus``
	// +optional
	JSONAttribs *string `json:"json_attribs,omitempty"` // Specifies the location in which some chef json data is stored. By default, it uses the; ``/etc/chef/firstboot.json`` location.
	// +optional
	LogLevel *string `json:"log_level,omitempty"` // Defines the level of logging to be stored in the log file. By default this value is set; to ``:info``.
	// +optional
	LogLocation *string `json:"log_location,omitempty"` // Specifies the location of the chef lof file. By default, the location is specified at; ``/var/log/chef/client.log``.
	// +optional
	NodeName *string `json:"node_name,omitempty"` // The name of the node to run. By default, we will use th instance id as the node name.
	// +optional
	OmnibusURL *string `json:"omnibus_url,omitempty"` // Omnibus URL if chef should be installed through Omnibus. By default, it uses the; ``https://www.chef.io/chef/install.sh``.
	// +optional
	OmnibusURLRetries *int64 `json:"omnibus_url_retries,omitempty"` // The number of retries that will be attempted to reach the Omnibus URL. Default is 5.
	// +optional
	OmnibusVersion *string `json:"omnibus_version,omitempty"` // Optional version string to require for omnibus install.
	// +optional
	PIDFile *string `json:"pid_file,omitempty"` // The location in which a process identification number (pid) is saved. By default, it; saves in the ``/var/run/chef/client.pid`` location.
	// +optional
	RunList []string `json:"run_list,omitempty"` // A run list for a first boot json.
	// +optional
	ServerURL *string `json:"server_url,omitempty"` // The URL for the chef server
	// +optional
	ShowTime *bool `json:"show_time,omitempty"` // Show time in chef logs
	// +optional
	SSLVerifyMode *string `json:"ssl_verify_mode,omitempty"` // Set the verify mode for HTTPS requests. We can have two possible values for this; parameter:; ; - ``:verify_none``: No validation of SSL certificates.; - ``:verify_peer``: Validate all SSL certificates.; ; By default, the parameter is set as ``:verify_none``.
	// +optional
	ValidationCERT *string `json:"validation_cert,omitempty"` // Optional string to be written to file validation_key. Special value ``system`` means set; use existing file.
	// +optional
	ValidationKey *string `json:"validation_key,omitempty"` // Optional path for validation_cert. default to ``/etc/chef/validation.pem``
	// +optional
	ValidationName *string `json:"validation_name,omitempty"` // The name of the chef-validator key that Chef Infra Client uses to access the Chef Infra; Server during the initial Chef Infra Client run.
}

type Chpasswd struct {
	// +optional
	Expire *bool `json:"expire,omitempty"` // Whether to expire all user passwords such that a password will need to be reset on the; user's next login. Default: ``true``
	// +optional
	List *List `json:"list"` // List of ``username:password`` pairs. Each user will have the corresponding password set.; A password can be randomly generated by specifying ``RANDOM`` or ``R`` as a user's; password. A hashed password, created by a tool like ``mkpasswd``, can be specified. A; regex (``r'\$(1|2a|2y|5|6)(\$.+){2}'``) is used to determine if a password value should; be treated as a hash.; ; Use of a multiline string for this field is DEPRECATED and will result in an error in a; future version of cloud-init.
	// +optional
	Users []UserClass `json:"users,omitempty"` // Replaces the deprecated ``list`` key. This key represents a list of existing users to set; passwords for. Each item under users contains the following required keys: ``name`` and; ``password`` or in the case of a randomly generated password, ``name`` and ``type``. The; ``type`` key has a default value of ``hash``, and may alternatively be set to ``text`` or; ``RANDOM``.
}

type UserClass struct {
	// +optional
	Name string `json:"name"`
	// +optional
	Type *Type `json:"type,omitempty"`
	// +optional
	Password *UserPassword `json:"password"`
}

type PurplePassword struct {
	// +optional
	Key string `json:"key"` // The key of the secret to select from.  Must be a valid secret key.
	// +optional
	Name *string `json:"name,omitempty"` // The name of the secret in the namespace to select from.
	// +optional
	Optional *bool `json:"optional,omitempty"` // Specify whether the Secret or its key must be defined
}

type GrubDpkg struct {
	// +optional
	Enabled *bool `json:"enabled,omitempty"` // Whether to configure which device is used as the target for grub installation. Default:; ``true``
	// +optional
	GrubPCInstallDevices *string `json:"grub-pc/install_devices,omitempty"` // Device to use as target for grub installation. If unspecified, ``grub-probe`` of; ``/boot`` will be used to find the device
	// +optional
	GrubPCInstallDevicesEmpty *GrubPCInstallDevicesEmpty `json:"grub-pc/install_devices_empty"` // Sets values for ``grub-pc/install_devices_empty``. If unspecified, will be set to; ``true`` if ``grub-pc/install_devices`` is empty, otherwise ``false``.
}

type DeviceAliases struct {
}

type DiskSetup struct {
}

type Drivers struct {
	// +optional
	Nvidia *Nvidia `json:"nvidia,omitempty"`
}

type Nvidia struct {
	// +optional
	LicenseAccepted bool `json:"license-accepted"` // Do you accept the NVIDIA driver license?
	// +optional
	Version *string `json:"version,omitempty"` // The version of the driver to install (e.g. "390", "410"). Defaults to the latest version.
}

type FSSetup struct {
	// +optional
	Cmd *Cmd `json:"cmd"` // Optional command to run to create the filesystem. Can include string substitutions of the; other ``fs_setup`` config keys. This is only necessary if you need to override the; default command.
	// +optional
	Device *string `json:"device,omitempty"` // Specified either as a path or as an alias in the format ``<alias name>.<y>`` where; ``<y>`` denotes the partition number on the device. If specifying device using the; ``<device name>.<partition number>`` format, the value of ``partition`` will be; overwritten.
	// +optional
	ExtraOpts *Cmd `json:"extra_opts"` // Optional options to pass to the filesystem creation command. Ignored if you using ``cmd``; directly.
	// +optional
	Filesystem *string `json:"filesystem,omitempty"` // Filesystem type to create. E.g., ``ext4`` or ``btrfs``
	// +optional
	Label *string `json:"label,omitempty"` // Label for the filesystem.
	// +optional
	Overwrite *bool `json:"overwrite,omitempty"` // If ``true``, overwrite any existing filesystem. Using ``overwrite: true`` for filesystems; is **dangerous** and can lead to data loss, so double check the entry in ``fs_setup``.; Default: ``false``
	// +optional
	Partition *Partition `json:"partition,omitempty"` // The partition can be specified by setting ``partition`` to the desired partition number.; The ``partition`` option may also be set to ``auto``, in which this module will search; for the existence of a filesystem matching the ``label``, ``type`` and ``device`` of the; ``fs_setup`` entry and will skip creating the filesystem if one is found. The; ``partition`` option may also be set to ``any``, in which case any file system that; matches ``type`` and ``device`` will cause this module to skip filesystem creation for; the ``fs_setup`` entry, regardless of ``label`` matching or not. To write a filesystem; directly to a device, use ``partition: none``. ``partition: none`` will **always** write; the filesystem, even when the ``label`` and ``filesystem`` are matched, and ``overwrite``; is ``false``.
	// +optional
	ReplaceFS *string `json:"replace_fs,omitempty"` // Ignored unless ``partition`` is ``auto`` or ``any``. Default ``false``.
}

type Fan struct {
	// +optional
	Config string `json:"config"` // The fan configuration to use as a single multi-line string
	// +optional
	ConfigPath *string `json:"config_path,omitempty"` // The path to write the fan configuration to. Default: ``/etc/network/fan``
}

type GroupsClass struct {
}

type GrowpartClass struct {
	// +optional
	Devices []string `json:"devices,omitempty"` // The devices to resize. Each entry can either be the path to the device's mountpoint in; the filesystem or a path to the block device in '/dev'. Default: ``[/]``
	// +optional
	IgnoreGrowrootDisabled *bool `json:"ignore_growroot_disabled,omitempty"` // If ``true``, ignore the presence of ``/etc/growroot-disabled``. If ``false`` and the file; exists, then don't resize. Default: ``false``
	// +optional
	Mode *ModeUnion `json:"mode"` // The utility to use for resizing. Default: ``auto``; ; Possible options:; ; * ``auto`` - Use any available utility; ; * ``growpart`` - Use growpart utility; ; * ``gpart`` - Use BSD gpart utility; ; * ``off`` - Take no action.
}

type Keyboard struct {
	// +optional
	Layout string `json:"layout"` // Required. Keyboard layout. Corresponds to XKBLAYOUT.
	// +optional
	Model *string `json:"model,omitempty"` // Optional. Keyboard model. Corresponds to XKBMODEL. Default: ``pc105``.
	// +optional
	Options *string `json:"options,omitempty"` // Optional. Keyboard options. Corresponds to XKBOPTIONS.
	// +optional
	Variant *string `json:"variant,omitempty"` // Optional. Keyboard variant. Corresponds to XKBVARIANT.
}

type Landscape struct {
	// +optional
	Client Client `json:"client"`
}

type Client struct {
	// +optional
	AccountName *string `json:"account_name,omitempty"` // The account this computer belongs to.
	// +optional
	ComputerTitle *string `json:"computer_title,omitempty"` // The title of this computer.
	// +optional
	DataPath *string `json:"data_path,omitempty"` // The directory to store data files in. Default: ``/var/lib/land‐scape/client/``.
	// +optional
	HTTPProxy *string `json:"http_proxy,omitempty"` // The URL of the HTTP proxy, if one is needed.
	// +optional
	HTTPSProxy *string `json:"https_proxy,omitempty"` // The URL of the HTTPS proxy, if one is needed.
	// +optional
	LogLevel *LogLevel `json:"log_level,omitempty"` // The log level for the client. Default: ``info``.
	// +optional
	PingURL *string `json:"ping_url,omitempty"` // The URL to perform lightweight exchange initiation with. Default:; ``https://landscape.canonical.com/ping``.
	// +optional
	RegistrationKey *string `json:"registration_key,omitempty"` // The account-wide key used for registering clients.
	// +optional
	Tags *string `json:"tags,omitempty"` // Comma separated list of tag names to be sent to the server.
	// +optional
	URL *string `json:"url,omitempty"` // The Landscape server URL to connect to. Default:; ``https://landscape.canonical.com/message-system``.
}

type Lxd struct {
	// +optional
	Bridge *Bridge `json:"bridge,omitempty"`
	// +optional
	Init *Init `json:"init,omitempty"`
}

type Bridge struct {
	// +optional
	Domain *string `json:"domain,omitempty"` // Domain to advertise to DHCP clients and use for DNS resolution.
	// +optional
	Ipv4Address *string `json:"ipv4_address,omitempty"` // IPv4 address for the bridge. If set, ``ipv4_netmask`` key required.
	// +optional
	Ipv4DHCPFirst *string `json:"ipv4_dhcp_first,omitempty"` // First IPv4 address of the DHCP range for the network created. This value will combined; with ``ipv4_dhcp_last`` key to set LXC ``ipv4.dhcp.ranges``.
	// +optional
	Ipv4DHCPLast *string `json:"ipv4_dhcp_last,omitempty"` // Last IPv4 address of the DHCP range for the network created. This value will combined; with ``ipv4_dhcp_first`` key to set LXC ``ipv4.dhcp.ranges``.
	// +optional
	Ipv4DHCPLeases *int64 `json:"ipv4_dhcp_leases,omitempty"` // Number of DHCP leases to allocate within the range. Automatically calculated based on; `ipv4_dhcp_first` and `ipv4_dchp_last` when unset.
	// +optional
	Ipv4Nat *bool `json:"ipv4_nat,omitempty"` // Set ``true`` to NAT the IPv4 traffic allowing for a routed IPv4 network. Default:; ``false``.
	// +optional
	Ipv4Netmask *int64 `json:"ipv4_netmask,omitempty"` // Prefix length for the ``ipv4_address`` key. Required when ``ipv4_address`` is set.
	// +optional
	Ipv6Address *string `json:"ipv6_address,omitempty"` // IPv6 address for the bridge (CIDR notation). When set, ``ipv6_netmask`` key is required.; When absent, no IPv6 will be configured.
	// +optional
	Ipv6Nat *bool `json:"ipv6_nat,omitempty"` // Whether to NAT. Default: ``false``.
	// +optional
	Ipv6Netmask *int64 `json:"ipv6_netmask,omitempty"` // Prefix length for ``ipv6_address`` provided. Required when ``ipv6_address`` is set.
	// +optional
	Mode BridgeMode `json:"mode"` // Whether to setup LXD bridge, use an existing bridge by ``name`` or create a new bridge.; `none` will avoid bridge setup, `existing` will configure lxd to use the bring matching; ``name`` and `new` will create a new bridge.
	// +optional
	MTU *int64 `json:"mtu,omitempty"` // Bridge MTU, defaults to LXD's default value
	// +optional
	Name *string `json:"name,omitempty"` // Name of the LXD network bridge to attach or create. Default: ``lxdbr0``.
}

type Init struct {
	// +optional
	NetworkAddress *string `json:"network_address,omitempty"` // IP address for LXD to listen on
	// +optional
	NetworkPort *int64 `json:"network_port,omitempty"` // Network port to bind LXD to.
	// +optional
	StorageBackend *StorageBackend `json:"storage_backend,omitempty"` // Storage backend to use. Default: ``dir``.
	// +optional
	StorageCreateDevice *string `json:"storage_create_device,omitempty"` // Setup device based storage using DEVICE
	// +optional
	StorageCreateLoop *int64 `json:"storage_create_loop,omitempty"` // Setup loop based storage with SIZE in GB
	// +optional
	StoragePool *string `json:"storage_pool,omitempty"` // Name of storage pool to use or create
	// +optional
	TrustPassword *TrustPasswordUnion `json:"trust_password"` // The password required to add new clients
}

type TrustPasswordClass struct {
	// +optional
	Key string `json:"key"` // The key of the secret to select from.  Must be a valid secret key.
	// +optional
	Name *string `json:"name,omitempty"` // The name of the secret in the namespace to select from.
	// +optional
	Optional *bool `json:"optional,omitempty"` // Specify whether the Secret or its key must be defined
}

type Mcollective struct {
	// +optional
	Conf *McollectiveConf `json:"conf,omitempty"`
}

type McollectiveConf struct {
	// +optional
	PrivateCERT *string `json:"private-cert,omitempty"` // Optional value of server private certificate which will be written to; ``/etc/mcollective/ssl/server-private.pem``
	// +optional
	PublicCERT *string `json:"public-cert,omitempty"` // Optional value of server public certificate which will be written to; ``/etc/mcollective/ssl/server-public.pem``
}

type NTP struct {
	// +optional
	Config *NTPConfig `json:"config,omitempty"` // Configuration settings or overrides for the; ``ntp_client`` specified.
	// +optional
	Enabled *bool `json:"enabled,omitempty"` // Attempt to enable ntp clients if set to True.  If set; to False, ntp client will not be configured or; installed
	// +optional
	NTPClient *string `json:"ntp_client,omitempty"` // Name of an NTP client to use to configure system NTP.; When unprovided or 'auto' the default client preferred; by the distribution will be used. The following; built-in client names can be used to override existing; configuration defaults: chrony, ntp, ntpdate,; systemd-timesyncd.
	// +optional
	Pools []string `json:"pools,omitempty"` // List of ntp pools. If both pools and servers are; empty, 4 default pool servers will be provided of; the format ``{0-3}.{distro}.pool.ntp.org``. NOTE:; for Alpine Linux when using the Busybox NTP client; this setting will be ignored due to the limited; functionality of Busybox's ntpd.
	// +optional
	Servers []string `json:"servers,omitempty"` // List of ntp servers. If both pools and servers are; empty, 4 default pool servers will be provided with; the format ``{0-3}.{distro}.pool.ntp.org``.
}

// Configuration settings or overrides for the
// ``ntp_client`` specified.
type NTPConfig struct {
	// +optional
	CheckExe *string `json:"check_exe,omitempty"` // The executable name for the ``ntp_client``.; For example, ntp service ``check_exe`` is; 'ntpd' because it runs the ntpd binary.
	// +optional
	Confpath *string `json:"confpath,omitempty"` // The path to where the ``ntp_client``; configuration is written.
	// +optional
	Packages []string `json:"packages,omitempty"` // List of packages needed to be installed for the; selected ``ntp_client``.
	// +optional
	ServiceName *string `json:"service_name,omitempty"` // The systemd or sysvinit service name used to; start and stop the ``ntp_client``; service.
	// +optional
	Template *string `json:"template,omitempty"` // Inline template allowing users to define their; own ``ntp_client`` configuration template.; The value must start with '## template:jinja'; to enable use of templating support.
}

type FluffyPassword struct {
	// +optional
	Key string `json:"key"` // The key of the secret to select from.  Must be a valid secret key.
	// +optional
	Name *string `json:"name,omitempty"` // The name of the secret in the namespace to select from.
	// +optional
	Optional *bool `json:"optional,omitempty"` // Specify whether the Secret or its key must be defined
}

type PhoneHome struct {
	// +optional
	Post *PostUnion `json:"post"` // A list of keys to post or ``all``. Default: ``all``
	// +optional
	Tries *int64 `json:"tries,omitempty"` // The number of times to try sending the phone home data. Default: ``10``
	// +optional
	URL string `json:"url"` // The URL to send the phone home data to.
}

type PowerState struct {
	// +optional
	Condition *Condition `json:"condition"` // Apply state change only if condition is met. May be boolean true (always met), false; (never met), or a command string or list to be executed. For command formatting, see the; documentation for ``cc_runcmd``. If exit code is 0, condition is met, otherwise not.; Default: ``true``
	// +optional
	Delay *Delay `json:"delay"` // Time in minutes to delay after cloud-init has finished. Can be ``now`` or an integer; specifying the number of minutes to delay. Default: ``now``
	// +optional
	Message *string `json:"message,omitempty"` // Optional message to display to the user when the system is powering off or rebooting.
	// +optional
	Mode PowerStateMode `json:"mode"` // Must be one of ``poweroff``, ``halt``, or ``reboot``.
	// +optional
	Timeout *int64 `json:"timeout,omitempty"` // Time in seconds to wait for the cloud-init process to finish before executing shutdown.; Default: ``30``
}

type Puppet struct {
	// +optional
	AioInstallURL *string `json:"aio_install_url,omitempty"` // If ``install_type`` is ``aio``, change the url of the install script.
	// +optional
	Cleanup *bool `json:"cleanup,omitempty"` // Whether to remove the puppetlabs repo after installation if ``install_type`` is ``aio``; Default: ``true``
	// +optional
	Collection *string `json:"collection,omitempty"` // Puppet collection to install if ``install_type`` is ``aio``. This can be set to one of; ``puppet`` (rolling release), ``puppet6``, ``puppet7`` (or their nightly counterparts) in; order to install specific release streams.
	// +optional
	Conf *PuppetConf `json:"conf,omitempty"` // Every key present in the conf object will be added to puppet.conf. As such, section names; should be one of: ``main``, ``server``, ``agent`` or ``user`` and keys should be valid; puppet configuration options. The configuration is specified as a dictionary containing; high-level ``<section>`` keys and lists of ``<key>=<value>`` pairs within each section.; The ``certname`` key supports string substitutions for ``%i`` and ``%f``, corresponding; to the instance id and fqdn of the machine respectively.; ; ``ca_cert`` is a special case. It won't be added to puppet.conf. It holds the; puppetserver certificate in pem format. It should be a multi-line string (using the |; yaml notation for multi-line strings).
	// +optional
	ConfFile *string `json:"conf_file,omitempty"` // The path to the puppet config file. Default depends on ``install_type``
	// +optional
	CsrAttributes *CsrAttributes `json:"csr_attributes,omitempty"` // create a ``csr_attributes.yaml`` file for CSR attributes and certificate extension; requests. See https://puppet.com/docs/puppet/latest/config_file_csr_attributes.html
	// +optional
	CsrAttributesPath *string `json:"csr_attributes_path,omitempty"` // The path to the puppet csr attributes file. Default depends on ``install_type``
	// +optional
	Exec *bool `json:"exec,omitempty"` // Whether or not to run puppet after configuration finishes. A single manual run can be; triggered by setting ``exec`` to ``true``, and additional arguments can be passed to; ``puppet agent`` via the ``exec_args`` key (by default the agent will execute with the; ``--test`` flag). Default: ``false``
	// +optional
	ExecArgs []string `json:"exec_args,omitempty"` // A list of arguments to pass to 'puppet agent' if 'exec' is true Default: ``['--test']``
	// +optional
	Install *bool `json:"install,omitempty"` // Whether or not to install puppet. Setting to ``false`` will result in an error if puppet; is not already present on the system. Default: ``true``
	// +optional
	InstallType *PuppetInstallType `json:"install_type,omitempty"` // Valid values are ``packages`` and ``aio``. Agent packages from the puppetlabs; repositories can be installed by setting ``aio``. Based on this setting, the default; config/SSL/CSR paths will be adjusted accordingly. Default: ``packages``
	// +optional
	PackageName *string `json:"package_name,omitempty"` // Name of the package to install if ``install_type`` is ``packages``. Default: ``puppet``
	// +optional
	SSLDir *string `json:"ssl_dir,omitempty"` // The path to the puppet SSL directory. Default depends on ``install_type``
	// +optional
	StartService *bool `json:"start_service,omitempty"` // By default, the puppet service will be automatically enabled after installation and set; to automatically start on boot. To override this in favor of manual puppet execution set; ``start_service`` to ``false``
	// +optional
	Version *string `json:"version,omitempty"` // Optional version to pass to the installer script or package manager. If unset, the latest; version from the repos will be installed.
}

// Every key present in the conf object will be added to puppet.conf. As such, section names
// should be one of: ``main``, ``server``, ``agent`` or ``user`` and keys should be valid
// puppet configuration options. The configuration is specified as a dictionary containing
// high-level ``<section>`` keys and lists of ``<key>=<value>`` pairs within each section.
// The ``certname`` key supports string substitutions for ``%i`` and ``%f``, corresponding
// to the instance id and fqdn of the machine respectively.
//
// ``ca_cert`` is a special case. It won't be added to puppet.conf. It holds the
// puppetserver certificate in pem format. It should be a multi-line string (using the |
// yaml notation for multi-line strings).
type PuppetConf struct {
	// +optional
	Agent map[string]interface{} `json:"agent,omitempty"`
	// +optional
	CACERT *string `json:"ca_cert,omitempty"`
	// +optional
	Main map[string]interface{} `json:"main,omitempty"`
	// +optional
	Server map[string]interface{} `json:"server,omitempty"`
	// +optional
	User map[string]interface{} `json:"user,omitempty"`
}

// create a ``csr_attributes.yaml`` file for CSR attributes and certificate extension
// requests. See https://puppet.com/docs/puppet/latest/config_file_csr_attributes.html
type CsrAttributes struct {
	// +optional
	CustomAttributes map[string]interface{} `json:"custom_attributes,omitempty"`
	// +optional
	ExtensionRequests map[string]interface{} `json:"extension_requests,omitempty"`
}

type RandomSeed struct {
	// +optional
	Command []string `json:"command,omitempty"` // Execute this command to seed random. The command will have RANDOM_SEED_FILE in its; environment set to the value of ``file`` above.
	// +optional
	CommandRequired *bool `json:"command_required,omitempty"` // If true, and ``command`` is not available to be run then an exception is raised and; cloud-init will record failure. Otherwise, only debug error is mentioned. Default:; ``false``
	// +optional
	Data *string `json:"data,omitempty"` // This data will be written to ``file`` before data from the datasource. When using a; multiline value or specifying binary data, be sure to follow yaml syntax and use the; ``|`` and ``!binary`` yaml format specifiers when appropriate
	// +optional
	Encoding *RandomSeedEncoding `json:"encoding,omitempty"` // Used to decode ``data`` provided. Allowed values are ``raw``, ``base64``, ``b64``,; ``gzip``, or ``gz``.  Default: ``raw``
	// +optional
	File *string `json:"file,omitempty"` // File to write random data to. Default: ``/dev/urandom``
}

type Reporting struct {
}

type ResolvConf struct {
	// +optional
	Domain *string `json:"domain,omitempty"` // The domain to be added as ``domain`` line
	// +optional
	Nameservers []string `json:"nameservers,omitempty"` // A list of nameservers to use to be added as ``nameserver`` lines
	// +optional
	Options map[string]interface{} `json:"options,omitempty"` // Key/value pairs of options to go under ``options`` heading. A unary option should be; specified as ``true``
	// +optional
	Searchdomains []string `json:"searchdomains,omitempty"` // A list of domains to be added ``search`` line
	// +optional
	Sortlist []string `json:"sortlist,omitempty"` // A list of IP addresses to be added to ``sortlist`` line
}

type RhSubscription struct {
	// +optional
	ActivationKey *ActivationKeyUnion `json:"activation-key"` // The activation key to use. Must be used with ``org``. Should not be used with; ``username`` or ``password``
	// +optional
	AddPool []string `json:"add-pool,omitempty"` // A list of pools ids add to the subscription
	// +optional
	AutoAttach *bool `json:"auto-attach,omitempty"` // Whether to attach subscriptions automatically
	// +optional
	DisableRepo []string `json:"disable-repo,omitempty"` // A list of repositories to disable
	// +optional
	EnableRepo []string `json:"enable-repo,omitempty"` // A list of repositories to enable
	// +optional
	Org *int64 `json:"org,omitempty"` // The organization number to use. Must be used with ``activation-key``. Should not be used; with ``username`` or ``password``
	// +optional
	Password *RhSubscriptionPassword `json:"password"` // The password to use. Must be used with username. Should not be used with; ``activation-key`` or ``org``
	// +optional
	RhsmBaseurl *string `json:"rhsm-baseurl,omitempty"` // Sets the baseurl in ``/etc/rhsm/rhsm.conf``
	// +optional
	ServerHostname *string `json:"server-hostname,omitempty"` // Sets the serverurl in ``/etc/rhsm/rhsm.conf``
	// +optional
	ServiceLevel *string `json:"service-level,omitempty"` // The service level to use when subscribing to RH repositories. ``auto-attach`` must be; true for this to be used
	// +optional
	Username *string `json:"username,omitempty"` // The username to use. Must be used with password. Should not be used with; ``activation-key`` or ``org``
}

type ActivationKey struct {
	// +optional
	Key string `json:"key"` // The key of the secret to select from.  Must be a valid secret key.
	// +optional
	Name *string `json:"name,omitempty"` // The name of the secret in the namespace to select from.
	// +optional
	Optional *bool `json:"optional,omitempty"` // Specify whether the Secret or its key must be defined
}

type TentacledPassword struct {
	// +optional
	Key string `json:"key"` // The key of the secret to select from.  Must be a valid secret key.
	// +optional
	Name *string `json:"name,omitempty"` // The name of the secret in the namespace to select from.
	// +optional
	Optional *bool `json:"optional,omitempty"` // Specify whether the Secret or its key must be defined
}

type Rsyslog struct {
	// +optional
	ConfigDir *string `json:"config_dir,omitempty"` // The directory where rsyslog configuration files will be written. Default:; ``/etc/rsyslog.d``
	// +optional
	ConfigFilename *string `json:"config_filename,omitempty"` // The name of the rsyslog configuration file. Default: ``20-cloud-config.conf``
	// +optional
	Configs []ConfigElement `json:"configs,omitempty"` // Each entry in ``configs`` is either a string or an object. Each config entry contains a; configuration string and a file to write it to. For config entries that are an object,; ``filename`` sets the target filename and ``content`` specifies the config string to; write. For config entries that are only a string, the string is used as the config string; to write. If the filename to write the config to is not specified, the value of the; ``config_filename`` key is used. A file with the selected filename will be written inside; the directory specified by ``config_dir``.
	// +optional
	Remotes map[string]interface{} `json:"remotes,omitempty"` // Each key is the name for an rsyslog remote entry. Each value holds the contents of the; remote config for rsyslog. The config consists of the following parts:; ; - filter for log messages (defaults to ``*.*``); ; - optional leading ``@`` or ``@@``, indicating udp and tcp respectively (defaults to; ``@``, for udp); ; - ipv4 or ipv6 hostname or address. ipv6 addresses must be in ``[::1]`` format, (e.g.; ``@[fd00::1]:514``); ; - optional port number (defaults to ``514``); ; This module will provide sane defaults for any part of the remote entry that is not; specified, so in most cases remote hosts can be specified just using ``<name>:; <address>``.
	// +optional
	ServiceReloadCommand *ServiceReloadCommandUnion `json:"service_reload_command"` // The command to use to reload the rsyslog service after the config has been updated. If; this is set to ``auto``, then an appropriate command for the distro will be used. This is; the default behavior. To manually set the command, use a list of command args (e.g.; ``[systemctl, restart, rsyslog]``).
}

type ConfigConfig struct {
	// +optional
	Content string `json:"content"`
	// +optional
	Filename *string `json:"filename,omitempty"`
}

type SSH struct {
	// +optional
	EmitKeysToConsole bool `json:"emit_keys_to_console"` // Set false to avoid printing SSH keys to system console. Default: ``true``.
}

// A dictionary entries for the public and private host keys of each desired key type.
// Entries in the ``ssh_keys`` config dict should have keys in the format ``<key
// type>_private``, ``<key type>_public``, and, optionally, ``<key type>_certificate``, e.g.
// ``rsa_private: <key>``, ``rsa_public: <key>``, and ``rsa_certificate: <key>``. Not all
// key types have to be specified, ones left unspecified will not be used. If this config
// option is used, then separate keys will not be automatically generated. In order to
// specify multiline private host keys and certificates, use yaml multiline syntax.
type SSHKeys struct {
}

type SSHPublishHostkeys struct {
	// +optional
	Blacklist []string `json:"blacklist,omitempty"` // The SSH key types to ignore when publishing. Default: ``[dsa]``
	// +optional
	Enabled *bool `json:"enabled,omitempty"` // If true, will read host keys from ``/etc/ssh/*.pub`` and publish them to the datasource; (if supported). Default: ``true``
}

type SaltMinion struct {
	// +optional
	Conf map[string]interface{} `json:"conf,omitempty"` // Configuration to be written to `config_dir`/minion
	// +optional
	ConfigDir *string `json:"config_dir,omitempty"` // Directory to write config files to. Default: ``/etc/salt``
	// +optional
	Grains map[string]interface{} `json:"grains,omitempty"` // Configuration to be written to `config_dir`/grains
	// +optional
	PkgName *string `json:"pkg_name,omitempty"` // Package name to install. Default: ``salt-minion``
	// +optional
	PKIDir *string `json:"pki_dir,omitempty"` // Directory to write key files. Default: `config_dir`/pki/minion
	// +optional
	PrivateKey *string `json:"private_key,omitempty"` // Private key to be used by salt minion
	// +optional
	PublicKey *string `json:"public_key,omitempty"` // Public key to be used by the salt minion
	// +optional
	ServiceName *string `json:"service_name,omitempty"` // Service name to enable. Default: ``salt-minion``
}

type Snap struct {
	// +optional
	Assertions *Assertions `json:"assertions"` // Properly-signed snap assertions which will run before and snap ``commands``.
	// +optional
	Commands *Commands `json:"commands"` // Snap commands to run on the target system
}

type Spacewalk struct {
	// +optional
	ActivationKey *string `json:"activation_key,omitempty"` // The activation key to use when registering with Spacewalk
	// +optional
	Proxy *string `json:"proxy,omitempty"` // The proxy to use when connecting to Spacewalk
	// +optional
	Server *string `json:"server,omitempty"` // The Spacewalk server to use
}

type Swap struct {
	// +optional
	Filename *string `json:"filename,omitempty"` // Path to the swap file to create
	// +optional
	Maxsize *Size `json:"maxsize"` // The maxsize in bytes of the swap file
	// +optional
	Size *Size `json:"size"` // The size in bytes of the swap file, 'auto' or a human-readable size abbreviation of the; format <float_size><units> where units are one of B, K, M, G or T.
}

type UbuntuAdvantage struct {
	// +optional
	Config *UbuntuAdvantageConfig `json:"config,omitempty"` // Configuration settings or override Ubuntu Advantage config
	// +optional
	Enable []string `json:"enable,omitempty"` // Optional list of ubuntu-advantage services to enable. Any of: cc-eal, cis, esm-infra,; fips, fips-updates, livepatch. By default, a given contract token will automatically; enable a number of services, use this list to supplement which services should; additionally be enabled. Any service unavailable on a given Ubuntu release or unentitled; in a given contract will remain disabled.
	// +optional
	Token string `json:"token"` // Required contract token obtained from https://ubuntu.com/advantage to attach.
}

// Configuration settings or override Ubuntu Advantage config
type UbuntuAdvantageConfig struct {
	// +optional
	GlobalAptHTTPProxy *string `json:"global_apt_http_proxy,omitempty"` // HTTP Proxy URL used for all APT repositories on a system. Stored at; ``/etc/apt/apt.conf.d/90ubuntu-advantage-aptproxy``
	// +optional
	GlobalAptHTTPSProxy *string `json:"global_apt_https_proxy,omitempty"` // HTTPS Proxy URL used for all APT repositories on a system. Stored at; ``/etc/apt/apt.conf.d/90ubuntu-advantage-aptproxy``
	// +optional
	HTTPProxy *string `json:"http_proxy,omitempty"` // Ubuntu Advantage HTTP Proxy URL
	// +optional
	HTTPSProxy *string `json:"https_proxy,omitempty"` // Ubuntu Advantage HTTPS Proxy URL
	// +optional
	UaAptHTTPProxy *string `json:"ua_apt_http_proxy,omitempty"` // HTTP Proxy URL used only for Ubuntu Advantage APT repositories. Stored at; ``/etc/apt/apt.conf.d/90ubuntu-advantage-aptproxy``
	// +optional
	UaAptHTTPSProxy *string `json:"ua_apt_https_proxy,omitempty"` // HTTPS Proxy URL used only for Ubuntu Advantage APT repositories. Stored at; ``/etc/apt/apt.conf.d/90ubuntu-advantage-aptproxy``
}

type Updates struct {
	// +optional
	Network *Network `json:"network,omitempty"`
}

type Network struct {
	// +optional
	When []When `json:"when"`
}

type PurpleSchemaCloudConfigV1 struct {
	// +optional
	CreateGroups *bool `json:"create_groups,omitempty"` // Boolean set ``false`` to disable creation of specified user ``groups``. Default: ``true``.
	// +optional
	Expiredate *string `json:"expiredate,omitempty"` // Optional. Date on which the user's account will be disabled. Default: ``null``
	// +optional
	Gecos *string `json:"gecos,omitempty"` // Optional comment about the user, usually a comma-separated string of real name and; contact information
	// +optional
	Groups *UserGroups `json:"groups"` // Optional comma-separated string of groups to add the user to.
	// +optional
	HashedPasswd *HashedPasswdUnion `json:"hashed_passwd"` // Hash of user password to be applied. This will be applied even if the user is; pre-existing. To generate this hash, run: mkpasswd --method=SHA-512 --rounds=4096.; **Note:** While ``hashed_password`` is better than ``plain_text_passwd``, using; ``passwd`` in user-data represents a security risk as user-data could be accessible by; third-parties depending on your cloud platform.
	// +optional
	Homedir *string `json:"homedir,omitempty"` // Optional home dir for user. Default: ``/home/<username>``
	// +optional
	Inactive *string `json:"inactive,omitempty"` // Optional string representing the number of days until the user is disabled.
	// +optional
	LockPasswd *bool `json:"lock-passwd,omitempty"` // Dropped after April 2027. Use ``lock_passwd``. Default: ``true``
	// +optional
	SchemaCloudConfigV1LockPasswd *bool `json:"lock_passwd,omitempty"` // Disable password login. Default: ``true``
	// +optional
	Name *string `json:"name,omitempty"` // The user's login name. Required otherwise user creation will be skipped for this user.
	// +optional
	NoCreateHome *bool `json:"no_create_home,omitempty"` // Do not create home directory. Default: ``false``
	// +optional
	NoLogInit *bool `json:"no_log_init,omitempty"` // Do not initialize lastlog and faillog for user. Default: ``false``
	// +optional
	NoUserGroup *bool `json:"no_user_group,omitempty"` // Do not create group named after user. Default: ``false``
	// +optional
	Passwd *PasswdUnion `json:"passwd"` // Hash of user password applied when user does not exist. This will NOT be applied if the; user already exists. To generate this hash, run: mkpasswd --method=SHA-512 --rounds=4096.; **Note:** While hashed password is better than plain text, using ``passwd`` in user-data; represents a security risk as user-data could be accessible by third-parties depending on; your cloud platform.
	// +optional
	PlainTextPasswd *PlainTextPasswdUnion `json:"plain_text_passwd"` // Clear text of user password to be applied. This will be applied even if the user is; pre-existing. There are many more secure options than using plain text passwords, such as; ``ssh_import_id`` or ``hashed_passwd``. Do not use this in production as user-data and; your password can be exposed.
	// +optional
	PrimaryGroup *string `json:"primary_group,omitempty"` // Primary group for user. Default: ``<username>``
	// +optional
	SelinuxUser *string `json:"selinux_user,omitempty"` // SELinux user for user's login. Default to default SELinux user.
	// +optional
	Shell *string `json:"shell,omitempty"` // Path to the user's login shell. The default is to set no shell, which results in a; system-specific default being used.
	// +optional
	Snapuser *string `json:"snapuser,omitempty"` // Specify an email address to create the user as a Snappy user through ``snap; create-user``. If an Ubuntu SSO account is associated with the address, username and SSH; keys will be requested from there.
	// +optional
	SSHAuthorizedKeys []string `json:"ssh_authorized_keys,omitempty"` // List of SSH keys to add to user's authkeys file. Can not be combined with; ``ssh_redirect_user``
	// +optional
	SSHImportID []string `json:"ssh_import_id,omitempty"` // List of SSH IDs to import for user. Can not be combined with ``ssh_redirect_user``.
	// +optional
	SSHRedirectUser *bool `json:"ssh_redirect_user,omitempty"` // Boolean set to true to disable SSH logins for this user. When specified, all cloud; meta-data public SSH keys will be set up in a disabled state for this username. Any SSH; login as this username will timeout and prompt with a message to login instead as the; ``default_username`` for this instance. Default: ``false``. This key can not be combined; with ``ssh_import_id`` or ``ssh_authorized_keys``.
	// +optional
	Sudo *Sudo `json:"sudo"`
	// +optional
	System *bool `json:"system,omitempty"` // Optional. Create user as system user with no home directory. Default: ``false``.
	// +optional
	Uid *Uid `json:"uid"` // The user's ID. Default is next available value.
}

type HashedPasswdClass struct {
	// +optional
	Key string `json:"key"` // The key of the secret to select from.  Must be a valid secret key.
	// +optional
	Name *string `json:"name,omitempty"` // The name of the secret in the namespace to select from.
	// +optional
	Optional *bool `json:"optional,omitempty"` // Specify whether the Secret or its key must be defined
}

type PasswdClass struct {
	// +optional
	Key string `json:"key"` // The key of the secret to select from.  Must be a valid secret key.
	// +optional
	Name *string `json:"name,omitempty"` // The name of the secret in the namespace to select from.
	// +optional
	Optional *bool `json:"optional,omitempty"` // Specify whether the Secret or its key must be defined
}

type PlainTextPasswdClass struct {
	// +optional
	Key string `json:"key"` // The key of the secret to select from.  Must be a valid secret key.
	// +optional
	Name *string `json:"name,omitempty"` // The name of the secret in the namespace to select from.
	// +optional
	Optional *bool `json:"optional,omitempty"` // Specify whether the Secret or its key must be defined
}

type FluffySchemaCloudConfigV1 struct {
	// +optional
	CreateGroups *bool `json:"create_groups,omitempty"` // Boolean set ``false`` to disable creation of specified user ``groups``. Default: ``true``.
	// +optional
	Expiredate *string `json:"expiredate,omitempty"` // Optional. Date on which the user's account will be disabled. Default: ``null``
	// +optional
	Gecos *string `json:"gecos,omitempty"` // Optional comment about the user, usually a comma-separated string of real name and; contact information
	// +optional
	Groups *UserGroups `json:"groups"` // Optional comma-separated string of groups to add the user to.
	// +optional
	HashedPasswd *HashedPasswdUnion `json:"hashed_passwd"` // Hash of user password to be applied. This will be applied even if the user is; pre-existing. To generate this hash, run: mkpasswd --method=SHA-512 --rounds=4096.; **Note:** While ``hashed_password`` is better than ``plain_text_passwd``, using; ``passwd`` in user-data represents a security risk as user-data could be accessible by; third-parties depending on your cloud platform.
	// +optional
	Homedir *string `json:"homedir,omitempty"` // Optional home dir for user. Default: ``/home/<username>``
	// +optional
	Inactive *string `json:"inactive,omitempty"` // Optional string representing the number of days until the user is disabled.
	// +optional
	LockPasswd *bool `json:"lock-passwd,omitempty"` // Dropped after April 2027. Use ``lock_passwd``. Default: ``true``
	// +optional
	SchemaCloudConfigV1LockPasswd *bool `json:"lock_passwd,omitempty"` // Disable password login. Default: ``true``
	// +optional
	Name *string `json:"name,omitempty"` // The user's login name. Required otherwise user creation will be skipped for this user.
	// +optional
	NoCreateHome *bool `json:"no_create_home,omitempty"` // Do not create home directory. Default: ``false``
	// +optional
	NoLogInit *bool `json:"no_log_init,omitempty"` // Do not initialize lastlog and faillog for user. Default: ``false``
	// +optional
	NoUserGroup *bool `json:"no_user_group,omitempty"` // Do not create group named after user. Default: ``false``
	// +optional
	Passwd *PasswdUnion `json:"passwd"` // Hash of user password applied when user does not exist. This will NOT be applied if the; user already exists. To generate this hash, run: mkpasswd --method=SHA-512 --rounds=4096.; **Note:** While hashed password is better than plain text, using ``passwd`` in user-data; represents a security risk as user-data could be accessible by third-parties depending on; your cloud platform.
	// +optional
	PlainTextPasswd *PlainTextPasswdUnion `json:"plain_text_passwd"` // Clear text of user password to be applied. This will be applied even if the user is; pre-existing. There are many more secure options than using plain text passwords, such as; ``ssh_import_id`` or ``hashed_passwd``. Do not use this in production as user-data and; your password can be exposed.
	// +optional
	PrimaryGroup *string `json:"primary_group,omitempty"` // Primary group for user. Default: ``<username>``
	// +optional
	SelinuxUser *string `json:"selinux_user,omitempty"` // SELinux user for user's login. Default to default SELinux user.
	// +optional
	Shell *string `json:"shell,omitempty"` // Path to the user's login shell. The default is to set no shell, which results in a; system-specific default being used.
	// +optional
	Snapuser *string `json:"snapuser,omitempty"` // Specify an email address to create the user as a Snappy user through ``snap; create-user``. If an Ubuntu SSO account is associated with the address, username and SSH; keys will be requested from there.
	// +optional
	SSHAuthorizedKeys []string `json:"ssh_authorized_keys,omitempty"` // List of SSH keys to add to user's authkeys file. Can not be combined with; ``ssh_redirect_user``
	// +optional
	SSHImportID []string `json:"ssh_import_id,omitempty"` // List of SSH IDs to import for user. Can not be combined with ``ssh_redirect_user``.
	// +optional
	SSHRedirectUser *bool `json:"ssh_redirect_user,omitempty"` // Boolean set to true to disable SSH logins for this user. When specified, all cloud; meta-data public SSH keys will be set up in a disabled state for this username. Any SSH; login as this username will timeout and prompt with a message to login instead as the; ``default_username`` for this instance. Default: ``false``. This key can not be combined; with ``ssh_import_id`` or ``ssh_authorized_keys``.
	// +optional
	Sudo *Sudo `json:"sudo"`
	// +optional
	System *bool `json:"system,omitempty"` // Optional. Create user as system user with no home directory. Default: ``false``.
	// +optional
	Uid *Uid `json:"uid"` // The user's ID. Default is next available value.
}

type VendorData struct {
	// +optional
	Enabled *GrubPCInstallDevicesEmpty `json:"enabled"` // Whether vendor data is enabled or not. Default: ``true``
	// +optional
	Prefix *Prefix `json:"prefix"` // The command to run before any vendor scripts. Its primary use case is for profiling a; script, not to prevent its run
}

type Wireguard struct {
	// +optional
	Interfaces []Interface `json:"interfaces"`
	// +optional
	Readinessprobe []string `json:"readinessprobe,omitempty"` // List of shell commands to be executed as probes.
}

type Interface struct {
	// +optional
	ConfigPath *string `json:"config_path,omitempty"` // Path to configuration file of Wireguard interface
	// +optional
	Content *string `json:"content,omitempty"` // Wireguard interface configuration. Contains key, peer, ...
	// +optional
	Name *string `json:"name,omitempty"` // Name of the interface. Typically wgx (example: wg0)
}

type WriteFile struct {
	// +optional
	Append *bool `json:"append,omitempty"` // Whether to append ``content`` to existing file if ``path`` exists. Default: ``false``.
	// +optional
	Content *ContentUnion `json:"content"` // Optional content to write to the provided ``path``. When content is present and encoding; is not 'text/plain', decode the content prior to writing. Default: ``''``
	// +optional
	Defer *bool `json:"defer,omitempty"` // Defer writing the file until 'final' stage, after users were created, and packages were; installed. Default: ``false``.
	// +optional
	Encoding *WriteFileEncoding `json:"encoding,omitempty"` // Optional encoding type of the content. Default is ``text/plain`` and no content decoding; is performed. Supported encoding types are: gz, gzip, gz+base64, gzip+base64, gz+b64,; gzip+b64, b64, base64
	// +optional
	Owner *string `json:"owner,omitempty"` // Optional owner:group to chown on the file. Default: ``root:root``
	// +optional
	Path string `json:"path"` // Path of the file to which ``content`` is decoded and written
	// +optional
	Permissions *string `json:"permissions,omitempty"` // Optional file permissions to set on ``path`` represented as an octal string '0###'.; Default: ``0o644``
}

type ContentClass struct {
	// +optional
	Key string `json:"key"` // The key of the secret to select from.  Must be a valid secret key.
	// +optional
	Name *string `json:"name,omitempty"` // The name of the secret in the namespace to select from.
	// +optional
	Optional *bool `json:"optional,omitempty"` // Specify whether the Secret or its key must be defined
}

type YumRepos struct {
}

type Zypper struct {
	// +optional
	Config map[string]interface{} `json:"config,omitempty"` // Any supported zypo.conf key is written to ``/etc/zypp/zypp.conf``
	// +optional
	Repos []Repo `json:"repos,omitempty"`
}

type Repo struct {
	// +optional
	Baseurl string `json:"baseurl"` // The base repositoy URL
	// +optional
	ID string `json:"id"` // The unique id of the repo, used when writing /etc/zypp/repos.d/<id>.repo.
}

// The type of installation for ansible. It can be one of the following values:
//
// - ``distro``
// - ``pip``
type InstallMethod string

const (
	Distro InstallMethod = "distro"
	Pip    InstallMethod = "pip"
)

type AptPipeliningEnum string

const (
	AptPipeliningNone AptPipeliningEnum = "none"
	OS                AptPipeliningEnum = "os"
	Unchanged         AptPipeliningEnum = "unchanged"
)

type ByobuByDefault string

const (
	Disable       ByobuByDefault = "disable"
	DisableSystem ByobuByDefault = "disable-system"
	DisableUser   ByobuByDefault = "disable-user"
	Enable        ByobuByDefault = "enable"
	EnableSystem  ByobuByDefault = "enable-system"
	EnableUser    ByobuByDefault = "enable-user"
	System        ByobuByDefault = "system"
	User          ByobuByDefault = "user"
)

// The type of installation for chef. It can be one of the following values:
//
// - ``packages``
// - ``gems``
// - ``omnibus``
type ChefInstallType string

const (
	Gems           ChefInstallType = "gems"
	Omnibus        ChefInstallType = "omnibus"
	PurplePackages ChefInstallType = "packages"
)

type Type string

const (
	Hash   Type = "hash"
	Random Type = "RANDOM"
	Text   Type = "text"
)

// The partition can be specified by setting ``partition`` to the desired partition number.
// The ``partition`` option may also be set to ``auto``, in which this module will search
// for the existence of a filesystem matching the ``label``, ``type`` and ``device`` of the
// ``fs_setup`` entry and will skip creating the filesystem if one is found. The
// ``partition`` option may also be set to ``any``, in which case any file system that
// matches ``type`` and ``device`` will cause this module to skip filesystem creation for
// the ``fs_setup`` entry, regardless of ``label`` matching or not. To write a filesystem
// directly to a device, use ``partition: none``. ``partition: none`` will **always** write
// the filesystem, even when the ``label`` and ``filesystem`` are matched, and ``overwrite``
// is ``false``.
//
// Optional command to run to create the filesystem. Can include string substitutions of the
// other ``fs_setup`` config keys. This is only necessary if you need to override the
// default command.
//
// Optional options to pass to the filesystem creation command. Ignored if you using ``cmd``
// directly.
//
// Use a boolean value instead.
//
// Use of string for this value is DEPRECATED. Use a boolean value instead.
//
// Use of non-boolean values for this field is DEPRECATED and will result in an error in a
// future version of cloud-init.
//
// Properly-signed snap assertions which will run before and snap ``commands``.
//
// The SSH public key to import
//
// The use of ``string`` type will be dropped after April 2027. Use an ``integer`` instead.
type Partition string

const (
	Any           Partition = "any"
	PartitionAuto Partition = "auto"
	PartitionNone Partition = "none"
)

type ModeMode string

const (
	Gpart    ModeMode = "gpart"
	Growpart ModeMode = "growpart"
	ModeAuto ModeMode = "auto"
	Off      ModeMode = "off"
)

// The log level for the client. Default: ``info``.
type LogLevel string

const (
	Critical LogLevel = "critical"
	Debug    LogLevel = "debug"
	Error    LogLevel = "error"
	Info     LogLevel = "info"
	Warning  LogLevel = "warning"
)

// Whether to setup LXD bridge, use an existing bridge by ``name`` or create a new bridge.
// `none` will avoid bridge setup, `existing` will configure lxd to use the bring matching
// ``name`` and `new` will create a new bridge.
type BridgeMode string

const (
	Existing BridgeMode = "existing"
	ModeNone BridgeMode = "none"
	New      BridgeMode = "new"
)

// Storage backend to use. Default: ``dir``.
type StorageBackend string

const (
	Btrfs StorageBackend = "btrfs"
	Dir   StorageBackend = "dir"
	LVM   StorageBackend = "lvm"
	Zfs   StorageBackend = "zfs"
)

// Value ``template`` will be dropped after April 2027. Use ``true`` instead.
type ManageEtcHostsEnum string

const (
	Localhost ManageEtcHostsEnum = "localhost"
	Template  ManageEtcHostsEnum = "template"
)

type PostElement string

const (
	FQDN          PostElement = "fqdn"
	Hostname      PostElement = "hostname"
	InstanceID    PostElement = "instance_id"
	PubKeyDSA     PostElement = "pub_key_dsa"
	PubKeyEcdsa   PostElement = "pub_key_ecdsa"
	PubKeyEd25519 PostElement = "pub_key_ed25519"
	PubKeyRSA     PostElement = "pub_key_rsa"
)

type PurplePost string

const (
	All PurplePost = "all"
)

// Must be one of ``poweroff``, ``halt``, or ``reboot``.
type PowerStateMode string

const (
	Halt     PowerStateMode = "halt"
	Poweroff PowerStateMode = "poweroff"
	Reboot   PowerStateMode = "reboot"
)

// Valid values are ``packages`` and ``aio``. Agent packages from the puppetlabs
// repositories can be installed by setting ``aio``. Based on this setting, the default
// config/SSL/CSR paths will be adjusted accordingly. Default: ``packages``
type PuppetInstallType string

const (
	Aio            PuppetInstallType = "aio"
	FluffyPackages PuppetInstallType = "packages"
)

// Used to decode ``data`` provided. Allowed values are ``raw``, ``base64``, ``b64``,
// ``gzip``, or ``gz``.  Default: ``raw``
type RandomSeedEncoding string

const (
	PurpleB64    RandomSeedEncoding = "b64"
	PurpleBase64 RandomSeedEncoding = "base64"
	PurpleGz     RandomSeedEncoding = "gz"
	PurpleGzip   RandomSeedEncoding = "gzip"
	Raw          RandomSeedEncoding = "raw"
)

type ResizeRootfsEnum string

const (
	Noblock ResizeRootfsEnum = "noblock"
)

type ServiceReloadCommandEnum string

const (
	ServiceReloadCommandAuto ServiceReloadCommandEnum = "auto"
)

type SSHGenkeytype string

const (
	DSA     SSHGenkeytype = "dsa"
	Ecdsa   SSHGenkeytype = "ecdsa"
	Ed25519 SSHGenkeytype = "ed25519"
	RSA     SSHGenkeytype = "rsa"
)

type When string

const (
	Boot            When = "boot"
	BootLegacy      When = "boot-legacy"
	BootNewInstance When = "boot-new-instance"
	Hotplug         When = "hotplug"
)

// Optional encoding type of the content. Default is ``text/plain`` and no content decoding
// is performed. Supported encoding types are: gz, gzip, gz+base64, gzip+base64, gz+b64,
// gzip+b64, b64, base64
type WriteFileEncoding string

const (
	FluffyB64    WriteFileEncoding = "b64"
	FluffyBase64 WriteFileEncoding = "base64"
	FluffyGz     WriteFileEncoding = "gz"
	FluffyGzip   WriteFileEncoding = "gzip"
	GzB64        WriteFileEncoding = "gz+b64"
	GzBase64     WriteFileEncoding = "gz+base64"
	GzipB64      WriteFileEncoding = "gzip+b64"
	GzipBase64   WriteFileEncoding = "gzip+base64"
	TextPlain    WriteFileEncoding = "text/plain"
)

type AptPipeliningUnion struct {
	Bool    *bool
	Enum    *AptPipeliningEnum
	Integer *int64
}

func (x *AptPipeliningUnion) UnmarshalJSON(data []byte) error {
	x.Enum = nil
	object, err := unmarshalUnion(data, &x.Integer, nil, &x.Bool, nil, false, nil, false, nil, false, nil, true, &x.Enum, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *AptPipeliningUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(x.Integer, nil, x.Bool, nil, false, nil, false, nil, false, nil, x.Enum != nil, x.Enum, false)
}

// Optional command to run to create the filesystem. Can include string substitutions of the
// other ``fs_setup`` config keys. This is only necessary if you need to override the
// default command.
//
// Optional options to pass to the filesystem creation command. Ignored if you using ``cmd``
// directly.
//
// Snap commands to run on the target system
type Cmd struct {
	String      *string
	StringArray []string
}

func (x *Cmd) UnmarshalJSON(data []byte) error {
	x.StringArray = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.StringArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Cmd) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.StringArray != nil, x.StringArray, false, nil, false, nil, false, nil, false)
}

// List of ``username:password`` pairs. Each user will have the corresponding password set.
// A password can be randomly generated by specifying ``RANDOM`` or ``R`` as a user's
// password. A hashed password, created by a tool like ``mkpasswd``, can be specified. A
// regex (``r'\$(1|2a|2y|5|6)(\$.+){2}'``) is used to determine if a password value should
// be treated as a hash.
//
// Use of a multiline string for this field is DEPRECATED and will result in an error in a
// future version of cloud-init.
type List struct {
	String      *string
	StringArray []string
}

func (x *List) UnmarshalJSON(data []byte) error {
	x.StringArray = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.StringArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *List) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.StringArray != nil, x.StringArray, false, nil, false, nil, false, nil, false)
}

type UserPassword struct {
	PurplePassword *PurplePassword
	String         *string
}

func (x *UserPassword) UnmarshalJSON(data []byte) error {
	x.PurplePassword = nil
	var c PurplePassword
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.PurplePassword = &c
	}
	return nil
}

func (x *UserPassword) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.PurplePassword != nil, x.PurplePassword, false, nil, false, nil, false)
}

type GrubPCInstallDevicesEmpty struct {
	Bool   *bool
	String *string
}

func (x *GrubPCInstallDevicesEmpty) UnmarshalJSON(data []byte) error {
	object, err := unmarshalUnion(data, nil, nil, &x.Bool, &x.String, false, nil, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *GrubPCInstallDevicesEmpty) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, x.Bool, x.String, false, nil, false, nil, false, nil, false, nil, false)
}

type CloudConfigGroups struct {
	AnythingArray []interface{}
	GroupsClass   *GroupsClass
	String        *string
}

func (x *CloudConfigGroups) UnmarshalJSON(data []byte) error {
	x.AnythingArray = nil
	x.GroupsClass = nil
	var c GroupsClass
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.AnythingArray, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.GroupsClass = &c
	}
	return nil
}

func (x *CloudConfigGroups) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.AnythingArray != nil, x.AnythingArray, x.GroupsClass != nil, x.GroupsClass, false, nil, false, nil, false)
}

// The utility to use for resizing. Default: ``auto``
//
// Possible options:
//
// * ``auto`` - Use any available utility
//
// * ``growpart`` - Use growpart utility
//
// * ``gpart`` - Use BSD gpart utility
//
// * ``off`` - Take no action.
type ModeUnion struct {
	Bool *bool
	Enum *ModeMode
}

func (x *ModeUnion) UnmarshalJSON(data []byte) error {
	x.Enum = nil
	object, err := unmarshalUnion(data, nil, nil, &x.Bool, nil, false, nil, false, nil, false, nil, true, &x.Enum, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *ModeUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, x.Bool, nil, false, nil, false, nil, false, nil, x.Enum != nil, x.Enum, false)
}

// The password required to add new clients
type TrustPasswordUnion struct {
	String             *string
	TrustPasswordClass *TrustPasswordClass
}

func (x *TrustPasswordUnion) UnmarshalJSON(data []byte) error {
	x.TrustPasswordClass = nil
	var c TrustPasswordClass
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.TrustPasswordClass = &c
	}
	return nil
}

func (x *TrustPasswordUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.TrustPasswordClass != nil, x.TrustPasswordClass, false, nil, false, nil, false)
}

// Whether to manage ``/etc/hosts`` on the system. If ``true``, render the hosts file using
// ``/etc/cloud/templates/hosts.tmpl`` replacing ``$hostname`` and ``$fdqn``. If
// ``localhost``, append a ``127.0.1.1`` entry that resolves from FQDN and hostname every
// boot. Default: ``false``.
type ManageEtcHostsUnion struct {
	Bool *bool
	Enum *ManageEtcHostsEnum
}

func (x *ManageEtcHostsUnion) UnmarshalJSON(data []byte) error {
	x.Enum = nil
	object, err := unmarshalUnion(data, nil, nil, &x.Bool, nil, false, nil, false, nil, false, nil, true, &x.Enum, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *ManageEtcHostsUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, x.Bool, nil, false, nil, false, nil, false, nil, x.Enum != nil, x.Enum, false)
}

// Set the default user's password. Ignored if ``chpasswd`` ``list`` is used
type CloudConfigPassword struct {
	FluffyPassword *FluffyPassword
	String         *string
}

func (x *CloudConfigPassword) UnmarshalJSON(data []byte) error {
	x.FluffyPassword = nil
	var c FluffyPassword
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.FluffyPassword = &c
	}
	return nil
}

func (x *CloudConfigPassword) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.FluffyPassword != nil, x.FluffyPassword, false, nil, false, nil, false)
}

// A list of keys to post or ``all``. Default: ``all``
type PostUnion struct {
	Enum      *PurplePost
	EnumArray []PostElement
}

func (x *PostUnion) UnmarshalJSON(data []byte) error {
	x.EnumArray = nil
	x.Enum = nil
	object, err := unmarshalUnion(data, nil, nil, nil, nil, true, &x.EnumArray, false, nil, false, nil, true, &x.Enum, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *PostUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, nil, x.EnumArray != nil, x.EnumArray, false, nil, false, nil, x.Enum != nil, x.Enum, false)
}

// Apply state change only if condition is met. May be boolean true (always met), false
// (never met), or a command string or list to be executed. For command formatting, see the
// documentation for ``cc_runcmd``. If exit code is 0, condition is met, otherwise not.
// Default: ``true``
type Condition struct {
	AnythingArray []interface{}
	Bool          *bool
	String        *string
}

func (x *Condition) UnmarshalJSON(data []byte) error {
	x.AnythingArray = nil
	object, err := unmarshalUnion(data, nil, nil, &x.Bool, &x.String, true, &x.AnythingArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Condition) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, x.Bool, x.String, x.AnythingArray != nil, x.AnythingArray, false, nil, false, nil, false, nil, false)
}

// Time in minutes to delay after cloud-init has finished. Can be ``now`` or an integer
// specifying the number of minutes to delay. Default: ``now``
type Delay struct {
	Integer *int64
	String  *string
}

func (x *Delay) UnmarshalJSON(data []byte) error {
	object, err := unmarshalUnion(data, &x.Integer, nil, nil, &x.String, false, nil, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Delay) MarshalJSON() ([]byte, error) {
	return marshalUnion(x.Integer, nil, nil, x.String, false, nil, false, nil, false, nil, false, nil, false)
}

// Whether to resize the root partition. ``noblock`` will resize in the background. Default:
// ``true``
type ResizeRootfsUnion struct {
	Bool *bool
	Enum *ResizeRootfsEnum
}

func (x *ResizeRootfsUnion) UnmarshalJSON(data []byte) error {
	x.Enum = nil
	object, err := unmarshalUnion(data, nil, nil, &x.Bool, nil, false, nil, false, nil, false, nil, true, &x.Enum, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *ResizeRootfsUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, x.Bool, nil, false, nil, false, nil, false, nil, x.Enum != nil, x.Enum, false)
}

// The activation key to use. Must be used with ``org``. Should not be used with
// ``username`` or ``password``
type ActivationKeyUnion struct {
	ActivationKey *ActivationKey
	String        *string
}

func (x *ActivationKeyUnion) UnmarshalJSON(data []byte) error {
	x.ActivationKey = nil
	var c ActivationKey
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.ActivationKey = &c
	}
	return nil
}

func (x *ActivationKeyUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.ActivationKey != nil, x.ActivationKey, false, nil, false, nil, false)
}

// The password to use. Must be used with username. Should not be used with
// ``activation-key`` or ``org``
type RhSubscriptionPassword struct {
	String            *string
	TentacledPassword *TentacledPassword
}

func (x *RhSubscriptionPassword) UnmarshalJSON(data []byte) error {
	x.TentacledPassword = nil
	var c TentacledPassword
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.TentacledPassword = &c
	}
	return nil
}

func (x *RhSubscriptionPassword) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.TentacledPassword != nil, x.TentacledPassword, false, nil, false, nil, false)
}

type ConfigElement struct {
	ConfigConfig *ConfigConfig
	String       *string
}

func (x *ConfigElement) UnmarshalJSON(data []byte) error {
	x.ConfigConfig = nil
	var c ConfigConfig
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.ConfigConfig = &c
	}
	return nil
}

func (x *ConfigElement) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.ConfigConfig != nil, x.ConfigConfig, false, nil, false, nil, false)
}

// The command to use to reload the rsyslog service after the config has been updated. If
// this is set to ``auto``, then an appropriate command for the distro will be used. This is
// the default behavior. To manually set the command, use a list of command args (e.g.
// ``[systemctl, restart, rsyslog]``).
type ServiceReloadCommandUnion struct {
	Enum        *ServiceReloadCommandEnum
	StringArray []string
}

func (x *ServiceReloadCommandUnion) UnmarshalJSON(data []byte) error {
	x.StringArray = nil
	x.Enum = nil
	object, err := unmarshalUnion(data, nil, nil, nil, nil, true, &x.StringArray, false, nil, false, nil, true, &x.Enum, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *ServiceReloadCommandUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, nil, x.StringArray != nil, x.StringArray, false, nil, false, nil, x.Enum != nil, x.Enum, false)
}

type Runcmd struct {
	String      *string
	StringArray []string
}

func (x *Runcmd) UnmarshalJSON(data []byte) error {
	x.StringArray = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.StringArray, false, nil, false, nil, false, nil, true)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Runcmd) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.StringArray != nil, x.StringArray, false, nil, false, nil, false, nil, true)
}

// Properly-signed snap assertions which will run before and snap ``commands``.
type Assertions struct {
	StringArray []string
	StringMap   map[string]string
}

func (x *Assertions) UnmarshalJSON(data []byte) error {
	x.StringArray = nil
	x.StringMap = nil
	object, err := unmarshalUnion(data, nil, nil, nil, nil, true, &x.StringArray, false, nil, true, &x.StringMap, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Assertions) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, nil, x.StringArray != nil, x.StringArray, false, nil, x.StringMap != nil, x.StringMap, false, nil, false)
}

// Snap commands to run on the target system
type Commands struct {
	UnionArray []Cmd
	UnionMap   map[string]*Cmd
}

func (x *Commands) UnmarshalJSON(data []byte) error {
	x.UnionArray = nil
	x.UnionMap = nil
	object, err := unmarshalUnion(data, nil, nil, nil, nil, true, &x.UnionArray, false, nil, true, &x.UnionMap, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Commands) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, nil, x.UnionArray != nil, x.UnionArray, false, nil, x.UnionMap != nil, x.UnionMap, false, nil, false)
}

// The maxsize in bytes of the swap file
//
// The size in bytes of the swap file, 'auto' or a human-readable size abbreviation of the
// format <float_size><units> where units are one of B, K, M, G or T.
type Size struct {
	Integer *int64
	String  *string
}

func (x *Size) UnmarshalJSON(data []byte) error {
	object, err := unmarshalUnion(data, &x.Integer, nil, nil, &x.String, false, nil, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Size) MarshalJSON() ([]byte, error) {
	return marshalUnion(x.Integer, nil, nil, x.String, false, nil, false, nil, false, nil, false, nil, false)
}

// The ``user`` dictionary values override the ``default_user`` configuration from
// ``/etc/cloud/cloud.cfg``. The `user` dictionary keys supported for the default_user are
// the same as the ``users`` schema.
type CloudConfigUser struct {
	PurpleSchemaCloudConfigV1 *PurpleSchemaCloudConfigV1
	String                    *string
}

func (x *CloudConfigUser) UnmarshalJSON(data []byte) error {
	x.PurpleSchemaCloudConfigV1 = nil
	var c PurpleSchemaCloudConfigV1
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.PurpleSchemaCloudConfigV1 = &c
	}
	return nil
}

func (x *CloudConfigUser) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.PurpleSchemaCloudConfigV1 != nil, x.PurpleSchemaCloudConfigV1, false, nil, false, nil, false)
}

// Optional comma-separated string of groups to add the user to.
type UserGroups struct {
	AnythingMap map[string]interface{}
	String      *string
	StringArray []string
}

func (x *UserGroups) UnmarshalJSON(data []byte) error {
	x.StringArray = nil
	x.AnythingMap = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.StringArray, false, nil, true, &x.AnythingMap, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *UserGroups) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.StringArray != nil, x.StringArray, false, nil, x.AnythingMap != nil, x.AnythingMap, false, nil, false)
}

// Hash of user password to be applied. This will be applied even if the user is
// pre-existing. To generate this hash, run: mkpasswd --method=SHA-512 --rounds=4096.
// **Note:** While ``hashed_password`` is better than ``plain_text_passwd``, using
// ``passwd`` in user-data represents a security risk as user-data could be accessible by
// third-parties depending on your cloud platform.
type HashedPasswdUnion struct {
	HashedPasswdClass *HashedPasswdClass
	String            *string
}

func (x *HashedPasswdUnion) UnmarshalJSON(data []byte) error {
	x.HashedPasswdClass = nil
	var c HashedPasswdClass
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.HashedPasswdClass = &c
	}
	return nil
}

func (x *HashedPasswdUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.HashedPasswdClass != nil, x.HashedPasswdClass, false, nil, false, nil, false)
}

// Hash of user password applied when user does not exist. This will NOT be applied if the
// user already exists. To generate this hash, run: mkpasswd --method=SHA-512 --rounds=4096.
// **Note:** While hashed password is better than plain text, using ``passwd`` in user-data
// represents a security risk as user-data could be accessible by third-parties depending on
// your cloud platform.
type PasswdUnion struct {
	PasswdClass *PasswdClass
	String      *string
}

func (x *PasswdUnion) UnmarshalJSON(data []byte) error {
	x.PasswdClass = nil
	var c PasswdClass
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.PasswdClass = &c
	}
	return nil
}

func (x *PasswdUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.PasswdClass != nil, x.PasswdClass, false, nil, false, nil, false)
}

// Clear text of user password to be applied. This will be applied even if the user is
// pre-existing. There are many more secure options than using plain text passwords, such as
// ``ssh_import_id`` or ``hashed_passwd``. Do not use this in production as user-data and
// your password can be exposed.
type PlainTextPasswdUnion struct {
	PlainTextPasswdClass *PlainTextPasswdClass
	String               *string
}

func (x *PlainTextPasswdUnion) UnmarshalJSON(data []byte) error {
	x.PlainTextPasswdClass = nil
	var c PlainTextPasswdClass
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.PlainTextPasswdClass = &c
	}
	return nil
}

func (x *PlainTextPasswdUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.PlainTextPasswdClass != nil, x.PlainTextPasswdClass, false, nil, false, nil, false)
}

type Sudo struct {
	Bool   *bool
	String *string
}

func (x *Sudo) UnmarshalJSON(data []byte) error {
	object, err := unmarshalUnion(data, nil, nil, &x.Bool, &x.String, false, nil, false, nil, false, nil, false, nil, true)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Sudo) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, x.Bool, x.String, false, nil, false, nil, false, nil, false, nil, true)
}

// The command to run before any vendor scripts. Its primary use case is for profiling a
// script, not to prevent its run
//
// The user's ID. Default is next available value.
type Uid struct {
	Integer *int64
	String  *string
}

func (x *Uid) UnmarshalJSON(data []byte) error {
	object, err := unmarshalUnion(data, &x.Integer, nil, nil, &x.String, false, nil, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Uid) MarshalJSON() ([]byte, error) {
	return marshalUnion(x.Integer, nil, nil, x.String, false, nil, false, nil, false, nil, false, nil, false)
}

type Users struct {
	AnythingMap map[string]interface{}
	String      *string
	UnionArray  []UsersUser
}

func (x *Users) UnmarshalJSON(data []byte) error {
	x.UnionArray = nil
	x.AnythingMap = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.UnionArray, false, nil, true, &x.AnythingMap, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Users) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.UnionArray != nil, x.UnionArray, false, nil, x.AnythingMap != nil, x.AnythingMap, false, nil, false)
}

type UsersUser struct {
	FluffySchemaCloudConfigV1 *FluffySchemaCloudConfigV1
	String                    *string
	StringArray               []string
}

func (x *UsersUser) UnmarshalJSON(data []byte) error {
	x.StringArray = nil
	x.FluffySchemaCloudConfigV1 = nil
	var c FluffySchemaCloudConfigV1
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.StringArray, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.FluffySchemaCloudConfigV1 = &c
	}
	return nil
}

func (x *UsersUser) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.StringArray != nil, x.StringArray, x.FluffySchemaCloudConfigV1 != nil, x.FluffySchemaCloudConfigV1, false, nil, false, nil, false)
}

// The command to run before any vendor scripts. Its primary use case is for profiling a
// script, not to prevent its run
type Prefix struct {
	String     *string
	UnionArray []Uid
}

func (x *Prefix) UnmarshalJSON(data []byte) error {
	x.UnionArray = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.UnionArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Prefix) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.UnionArray != nil, x.UnionArray, false, nil, false, nil, false, nil, false)
}

// Optional content to write to the provided ``path``. When content is present and encoding
// is not 'text/plain', decode the content prior to writing. Default: ``''``
type ContentUnion struct {
	ContentClass *ContentClass
	String       *string
}

func (x *ContentUnion) UnmarshalJSON(data []byte) error {
	x.ContentClass = nil
	var c ContentClass
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.ContentClass = &c
	}
	return nil
}

func (x *ContentUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.ContentClass != nil, x.ContentClass, false, nil, false, nil, false)
}

func unmarshalUnion(data []byte, pi **int64, pf **float64, pb **bool, ps **string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) (bool, error) {
	if pi != nil {
		*pi = nil
	}
	if pf != nil {
		*pf = nil
	}
	if pb != nil {
		*pb = nil
	}
	if ps != nil {
		*ps = nil
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	tok, err := dec.Token()
	if err != nil {
		return false, err
	}

	switch v := tok.(type) {
	case json.Number:
		if pi != nil {
			i, err := v.Int64()
			if err == nil {
				*pi = &i
				return false, nil
			}
		}
		if pf != nil {
			f, err := v.Float64()
			if err == nil {
				*pf = &f
				return false, nil
			}
			return false, errors.New("Unparsable number")
		}
		return false, errors.New("Union does not contain number")
	case float64:
		return false, errors.New("Decoder should not return float64")
	case bool:
		if pb != nil {
			*pb = &v
			return false, nil
		}
		return false, errors.New("Union does not contain bool")
	case string:
		if haveEnum {
			return false, json.Unmarshal(data, pe)
		}
		if ps != nil {
			*ps = &v
			return false, nil
		}
		return false, errors.New("Union does not contain string")
	case nil:
		if nullable {
			return false, nil
		}
		return false, errors.New("Union does not contain null")
	case json.Delim:
		if v == '{' {
			if haveObject {
				return true, json.Unmarshal(data, pc)
			}
			if haveMap {
				return false, json.Unmarshal(data, pm)
			}
			return false, errors.New("Union does not contain object")
		}
		if v == '[' {
			if haveArray {
				return false, json.Unmarshal(data, pa)
			}
			return false, errors.New("Union does not contain array")
		}
		return false, errors.New("Cannot handle delimiter")
	}
	return false, errors.New("Cannot unmarshal union")

}

func marshalUnion(pi *int64, pf *float64, pb *bool, ps *string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) ([]byte, error) {
	if pi != nil {
		return json.Marshal(*pi)
	}
	if pf != nil {
		return json.Marshal(*pf)
	}
	if pb != nil {
		return json.Marshal(*pb)
	}
	if ps != nil {
		return json.Marshal(*ps)
	}
	if haveArray {
		return json.Marshal(pa)
	}
	if haveObject {
		return json.Marshal(pc)
	}
	if haveMap {
		return json.Marshal(pm)
	}
	if haveEnum {
		return json.Marshal(pe)
	}
	if nullable {
		return json.Marshal(nil)
	}
	return nil, errors.New("Union must not be null")
}
