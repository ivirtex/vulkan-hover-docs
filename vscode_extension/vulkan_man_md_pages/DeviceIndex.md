# DeviceIndex(3) Manual Page

## Name

DeviceIndex - Index of the device executing the shader



## <a href="#_description" class="anchor"></a>Description

`DeviceIndex`  
The `DeviceIndex` decoration **can** be applied to a shader input which
will be filled with the device index of the physical device that is
executing the current shader invocation. This value will be in the range
\[0,max(1,physicalDeviceCount)), where physicalDeviceCount is the
`physicalDeviceCount` member of
[VkDeviceGroupDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupDeviceCreateInfo.html).

Valid Usage

- <a href="#VUID-DeviceIndex-DeviceIndex-04205"
  id="VUID-DeviceIndex-DeviceIndex-04205"></a>
  VUID-DeviceIndex-DeviceIndex-04205  
  The variable decorated with `DeviceIndex` **must** be declared using
  the `Input` `Storage` `Class`

- <a href="#VUID-DeviceIndex-DeviceIndex-04206"
  id="VUID-DeviceIndex-DeviceIndex-04206"></a>
  VUID-DeviceIndex-DeviceIndex-04206  
  The variable decorated with `DeviceIndex` **must** be declared as a
  scalar 32-bit integer value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#DeviceIndex"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
