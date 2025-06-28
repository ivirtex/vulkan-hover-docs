# DeviceIndex(3) Manual Page

## Name

DeviceIndex - Index of the device executing the shader



## [](#_description)Description

`DeviceIndex`

The `DeviceIndex` decoration **can** be applied to a shader input which will be filled with the device index of the physical device that is executing the current shader invocation. This value will be in the range \[0,max(1,physicalDeviceCount)), where physicalDeviceCount is the `physicalDeviceCount` member of [VkDeviceGroupDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupDeviceCreateInfo.html).

Valid Usage

- [](#VUID-DeviceIndex-DeviceIndex-04205)VUID-DeviceIndex-DeviceIndex-04205  
  The variable decorated with `DeviceIndex` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-DeviceIndex-DeviceIndex-04206)VUID-DeviceIndex-DeviceIndex-04206  
  The variable decorated with `DeviceIndex` **must** be declared as a scalar 32-bit integer value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#DeviceIndex)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0