# vkGetMicromapBuildSizesEXT(3) Manual Page

## Name

vkGetMicromapBuildSizesEXT - Retrieve the required size for a micromap



## <a href="#_c_specification" class="anchor"></a>C Specification

To get the build sizes for a micromap, call:

``` c
// Provided by VK_EXT_opacity_micromap
void vkGetMicromapBuildSizesEXT(
    VkDevice                                    device,
    VkAccelerationStructureBuildTypeKHR         buildType,
    const VkMicromapBuildInfoEXT*               pBuildInfo,
    VkMicromapBuildSizesInfoEXT*                pSizeInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that will be used for creating the
  micromap.

- `buildType` defines whether host or device operations (or both) are
  being queried for.

- `pBuildInfo` is a pointer to a
  [VkMicromapBuildInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapBuildInfoEXT.html) structure
  describing parameters of a build operation.

- `pSizeInfo` is a pointer to a
  [VkMicromapBuildSizesInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapBuildSizesInfoEXT.html)
  structure which returns the size required for a micromap and the sizes
  required for the scratch buffers, given the build parameters.

## <a href="#_description" class="anchor"></a>Description

The `dstMicromap` and `mode` members of `pBuildInfo` are ignored. Any
[VkDeviceOrHostAddressKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceOrHostAddressKHR.html) members of
`pBuildInfo` are ignored by this command.

A micromap created with the `micromapSize` returned by this command
supports any build with a
[VkMicromapBuildInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapBuildInfoEXT.html) structure subject
to the following properties:

- The build command is a host build command, and `buildType` is
  `VK_ACCELERATION_STRUCTURE_BUILD_TYPE_HOST_KHR` or
  `VK_ACCELERATION_STRUCTURE_BUILD_TYPE_HOST_OR_DEVICE_KHR`

- The build command is a device build command, and `buildType` is
  `VK_ACCELERATION_STRUCTURE_BUILD_TYPE_DEVICE_KHR` or
  `VK_ACCELERATION_STRUCTURE_BUILD_TYPE_HOST_OR_DEVICE_KHR`

- For [VkMicromapBuildInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapBuildInfoEXT.html):

  - Its `type`, and `flags` members are equal to `pBuildInfo->type` and
    `pBuildInfo->flags`, respectively.

  - The sum of usage information in either `pUsageCounts` or
    `ppUsageCounts` is equal to the sum of usage information in either
    `pBuildInfo->pUsageCounts` or `pBuildInfo->ppUsageCounts`.

Similarly, the `buildScratchSize` value will support any build command
specifying the `VK_BUILD_MICROMAP_MODE_BUILD_EXT` `mode` under the above
conditions.

Valid Usage

- <a href="#VUID-vkGetMicromapBuildSizesEXT-dstMicromap-09180"
  id="VUID-vkGetMicromapBuildSizesEXT-dstMicromap-09180"></a>
  VUID-vkGetMicromapBuildSizesEXT-dstMicromap-09180  
  [VkMicromapBuildInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapBuildInfoEXT.html)::`dstMicromap`
  **must** have been created from `device`

- <a href="#VUID-vkGetMicromapBuildSizesEXT-micromap-07439"
  id="VUID-vkGetMicromapBuildSizesEXT-micromap-07439"></a>
  VUID-vkGetMicromapBuildSizesEXT-micromap-07439  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-micromap"
  target="_blank" rel="noopener"><code>micromap</code></a> feature
  **must** be enabled

- <a href="#VUID-vkGetMicromapBuildSizesEXT-device-07440"
  id="VUID-vkGetMicromapBuildSizesEXT-device-07440"></a>
  VUID-vkGetMicromapBuildSizesEXT-device-07440  
  If `device` was created with multiple physical devices, then the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-bufferDeviceAddressMultiDevice"
  target="_blank"
  rel="noopener"><code>bufferDeviceAddressMultiDevice</code></a> feature
  **must** be enabled

Valid Usage (Implicit)

- <a href="#VUID-vkGetMicromapBuildSizesEXT-device-parameter"
  id="VUID-vkGetMicromapBuildSizesEXT-device-parameter"></a>
  VUID-vkGetMicromapBuildSizesEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetMicromapBuildSizesEXT-buildType-parameter"
  id="VUID-vkGetMicromapBuildSizesEXT-buildType-parameter"></a>
  VUID-vkGetMicromapBuildSizesEXT-buildType-parameter  
  `buildType` **must** be a valid
  [VkAccelerationStructureBuildTypeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureBuildTypeKHR.html)
  value

- <a href="#VUID-vkGetMicromapBuildSizesEXT-pBuildInfo-parameter"
  id="VUID-vkGetMicromapBuildSizesEXT-pBuildInfo-parameter"></a>
  VUID-vkGetMicromapBuildSizesEXT-pBuildInfo-parameter  
  `pBuildInfo` **must** be a valid pointer to a valid
  [VkMicromapBuildInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapBuildInfoEXT.html) structure

- <a href="#VUID-vkGetMicromapBuildSizesEXT-pSizeInfo-parameter"
  id="VUID-vkGetMicromapBuildSizesEXT-pSizeInfo-parameter"></a>
  VUID-vkGetMicromapBuildSizesEXT-pSizeInfo-parameter  
  `pSizeInfo` **must** be a valid pointer to a
  [VkMicromapBuildSizesInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapBuildSizesInfoEXT.html)
  structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_opacity_micromap](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_opacity_micromap.html),
[VkAccelerationStructureBuildTypeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureBuildTypeKHR.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkMicromapBuildInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapBuildInfoEXT.html),
[VkMicromapBuildSizesInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapBuildSizesInfoEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetMicromapBuildSizesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
