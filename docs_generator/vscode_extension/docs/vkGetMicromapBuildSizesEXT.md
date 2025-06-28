# vkGetMicromapBuildSizesEXT(3) Manual Page

## Name

vkGetMicromapBuildSizesEXT - Retrieve the required size for a micromap



## [](#_c_specification)C Specification

To get the build sizes for a micromap, call:

```c++
// Provided by VK_EXT_opacity_micromap
void vkGetMicromapBuildSizesEXT(
    VkDevice                                    device,
    VkAccelerationStructureBuildTypeKHR         buildType,
    const VkMicromapBuildInfoEXT*               pBuildInfo,
    VkMicromapBuildSizesInfoEXT*                pSizeInfo);
```

## [](#_parameters)Parameters

- `device` is the logical device that will be used for creating the micromap.
- `buildType` defines whether host or device operations (or both) are being queried for.
- `pBuildInfo` is a pointer to a [VkMicromapBuildInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapBuildInfoEXT.html) structure describing parameters of a build operation.
- `pSizeInfo` is a pointer to a [VkMicromapBuildSizesInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapBuildSizesInfoEXT.html) structure which returns the size required for a micromap and the sizes required for the scratch buffers, given the build parameters.

## [](#_description)Description

The `dstMicromap` and `mode` members of `pBuildInfo` are ignored. Any [VkDeviceOrHostAddressKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceOrHostAddressKHR.html) members of `pBuildInfo` are ignored by this command.

A micromap created with the `micromapSize` returned by this command supports any build with a [VkMicromapBuildInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapBuildInfoEXT.html) structure subject to the following properties:

- The build command is a host build command, and `buildType` is `VK_ACCELERATION_STRUCTURE_BUILD_TYPE_HOST_KHR` or `VK_ACCELERATION_STRUCTURE_BUILD_TYPE_HOST_OR_DEVICE_KHR`
- The build command is a device build command, and `buildType` is `VK_ACCELERATION_STRUCTURE_BUILD_TYPE_DEVICE_KHR` or `VK_ACCELERATION_STRUCTURE_BUILD_TYPE_HOST_OR_DEVICE_KHR`
- For [VkMicromapBuildInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapBuildInfoEXT.html):
  
  - Its `type`, and `flags` members are equal to `pBuildInfo->type` and `pBuildInfo->flags`, respectively.
  - The sum of usage information in either `pUsageCounts` or `ppUsageCounts` is equal to the sum of usage information in either `pBuildInfo->pUsageCounts` or `pBuildInfo->ppUsageCounts`.

Similarly, the `buildScratchSize` value will support any build command specifying the `VK_BUILD_MICROMAP_MODE_BUILD_EXT` `mode` under the above conditions.

Valid Usage

- [](#VUID-vkGetMicromapBuildSizesEXT-dstMicromap-09180)VUID-vkGetMicromapBuildSizesEXT-dstMicromap-09180  
  [VkMicromapBuildInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapBuildInfoEXT.html)::`dstMicromap` **must** have been created from `device`
- [](#VUID-vkGetMicromapBuildSizesEXT-micromap-07439)VUID-vkGetMicromapBuildSizesEXT-micromap-07439  
  The [`micromap`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-micromap) feature **must** be enabled
- [](#VUID-vkGetMicromapBuildSizesEXT-device-07440)VUID-vkGetMicromapBuildSizesEXT-device-07440  
  If `device` was created with multiple physical devices, then the [`bufferDeviceAddressMultiDevice`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-bufferDeviceAddressMultiDevice) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-vkGetMicromapBuildSizesEXT-device-parameter)VUID-vkGetMicromapBuildSizesEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetMicromapBuildSizesEXT-buildType-parameter)VUID-vkGetMicromapBuildSizesEXT-buildType-parameter  
  `buildType` **must** be a valid [VkAccelerationStructureBuildTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildTypeKHR.html) value
- [](#VUID-vkGetMicromapBuildSizesEXT-pBuildInfo-parameter)VUID-vkGetMicromapBuildSizesEXT-pBuildInfo-parameter  
  `pBuildInfo` **must** be a valid pointer to a valid [VkMicromapBuildInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapBuildInfoEXT.html) structure
- [](#VUID-vkGetMicromapBuildSizesEXT-pSizeInfo-parameter)VUID-vkGetMicromapBuildSizesEXT-pSizeInfo-parameter  
  `pSizeInfo` **must** be a valid pointer to a [VkMicromapBuildSizesInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapBuildSizesInfoEXT.html) structure

## [](#_see_also)See Also

[VK\_EXT\_opacity\_micromap](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_opacity_micromap.html), [VkAccelerationStructureBuildTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildTypeKHR.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkMicromapBuildInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapBuildInfoEXT.html), [VkMicromapBuildSizesInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapBuildSizesInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetMicromapBuildSizesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0