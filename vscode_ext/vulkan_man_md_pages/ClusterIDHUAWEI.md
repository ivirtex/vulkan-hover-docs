# ClusterIDHUAWEI(3) Manual Page

## Name

ClusterIDHUAWEI - cluster culling shader output variable



## <a href="#_description" class="anchor"></a>Description

`ClusterIDHUAWEI`  
The `ClusterIDHUAWEI` decoration can be used to decorate a cluster
culling shader output variable,this variable will contain an integer
value that specifies the id of cluster being rendered by this drawing
command. When Cluster Culling Shader enable, `ClusterIDHUAWEI` will
replace gl_DrawID pass to vertex shader for cluster-related information
fetching.

Valid Usage

- <a href="#VUID-ClusterIDHUAWEI-ClusterIDHUAWEI-07797"
  id="VUID-ClusterIDHUAWEI-ClusterIDHUAWEI-07797"></a>
  VUID-ClusterIDHUAWEI-ClusterIDHUAWEI-07797  
  The `ClusterIDHUAWEI` decoration **must** be used only within the
  `ClusterCullingHUAWEI` `Execution` `Model`

- <a href="#VUID-ClusterIDHUAWEI-ClusterIDHUAWEI-07798"
  id="VUID-ClusterIDHUAWEI-ClusterIDHUAWEI-07798"></a>
  VUID-ClusterIDHUAWEI-ClusterIDHUAWEI-07798  
  The variable decorated with `ClusterIDHUAWEI` **must** be declared as
  a scalar 32-bit integer value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#ClusterIDHUAWEI"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
