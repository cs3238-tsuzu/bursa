#
# Cookbook Name:: bursa-gpm
# Recipe:: default
#
# Copyright (c) 2014 The Authors, All Rights Reserved.

remote_file "/usr/local/bin/gpm" do
  source "https://raw.githubusercontent.com/pote/gpm/v1.3.1/bin/gpm"
  mode 0755
end

remote_file "/usr/local/bin/gvp" do
  source "https://raw.githubusercontent.com/pote/gvp/master/bin/gvp"
  mode 0755
end