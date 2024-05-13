# InstanceCountHUAWEI(3) Manual Page

## Name

InstanceCountHUAWEI - cluster culling shader output variable



## <a href="#_description" class="anchor"></a>Description

`InstanceCountHUAWEI`  
The `InstanceCountHUAWEI` decoration can be used to decorate a cluster
culling shader output variable,this variable will contain an integer
value that specifies the number of instance to draw in a cluster.

Valid Usage

- <a href="#VUID-InstanceCountHUAWEI-InstanceCountHUAWEI-07807"
  id="VUID-InstanceCountHUAWEI-InstanceCountHUAWEI-07807"></a>
  VUID-InstanceCountHUAWEI-InstanceCountHUAWEI-07807  
  The `InstanceCountHUAWEI` decoration **must** be used only within the
  `ClusterCullingHUAWEI` `Execution` `Model`

- <a href="#VUID-InstanceCountHUAWEI-InstanceCountHUAWEI-07808"
  id="VUID-InstanceCountHUAWEI-InstanceCountHUAWEI-07808"></a>
  VUID-InstanceCountHUAWEI-InstanceCountHUAWEI-07808  
  The variable decorated with `InstanceCountHUAWEI` **must** be declared
  as a scalar 32-bit integer value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#InstanceCountHUAWEI"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
