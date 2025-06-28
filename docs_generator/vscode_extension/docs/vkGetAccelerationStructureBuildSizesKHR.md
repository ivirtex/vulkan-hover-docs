# vkGetAccelerationStructureBuildSizesKHR(3) Manual Page

## Name

vkGetAccelerationStructureBuildSizesKHR - Retrieve the required size for an acceleration structure



## [](#_c_specification)C Specification

To get the build sizes for an acceleration structure, call:

```c++
// Provided by VK_KHR_acceleration_structure
void vkGetAccelerationStructureBuildSizesKHR(
    VkDevice                                    device,
    VkAccelerationStructureBuildTypeKHR         buildType,
    const VkAccelerationStructureBuildGeometryInfoKHR* pBuildInfo,
    const uint32_t*                             pMaxPrimitiveCounts,
    VkAccelerationStructureBuildSizesInfoKHR*   pSizeInfo);
```

## [](#_parameters)Parameters

- `device` is the logical device that will be used for creating the acceleration structure.
- `buildType` defines whether host or device operations (or both) are being queried for.
- `pBuildInfo` is a pointer to a [VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html) structure describing parameters of a build operation.
- `pMaxPrimitiveCounts` is a pointer to an array of `pBuildInfo->geometryCount` `uint32_t` values defining the number of primitives built into each geometry.
- `pSizeInfo` is a pointer to a [VkAccelerationStructureBuildSizesInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildSizesInfoKHR.html) structure which returns the size required for an acceleration structure and the sizes required for the scratch buffers, given the build parameters.

## [](#_description)Description

The `srcAccelerationStructure`, `dstAccelerationStructure`, and `mode` members of `pBuildInfo` are ignored. Any [VkDeviceOrHostAddressKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceOrHostAddressKHR.html) or [VkDeviceOrHostAddressConstKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceOrHostAddressConstKHR.html) members of `pBuildInfo` are ignored by this command, except that the `hostAddress` member of [VkAccelerationStructureGeometryTrianglesDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryTrianglesDataKHR.html)::`transformData` will be examined to check if it is `NULL`.

An acceleration structure created with the `accelerationStructureSize` returned by this command supports any build or update with a [VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html) structure and array of [VkAccelerationStructureBuildRangeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildRangeInfoKHR.html) structures subject to the following properties:

- The build command is a host build command, and `buildType` is `VK_ACCELERATION_STRUCTURE_BUILD_TYPE_HOST_KHR` or `VK_ACCELERATION_STRUCTURE_BUILD_TYPE_HOST_OR_DEVICE_KHR`
- The build command is a device build command, and `buildType` is `VK_ACCELERATION_STRUCTURE_BUILD_TYPE_DEVICE_KHR` or `VK_ACCELERATION_STRUCTURE_BUILD_TYPE_HOST_OR_DEVICE_KHR`
- For [VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html):
  
  - Its `type`, and `flags` members are equal to `pBuildInfo->type` and `pBuildInfo->flags`, respectively.
  - `geometryCount` is less than or equal to `pBuildInfo->geometryCount`.
  - For each element of either `pGeometries` or `ppGeometries` at a given index, its `geometryType` member is equal to `pBuildInfo->geometryType`.
  - For each element of either `pGeometries` or `ppGeometries` at a given index, its `flags` member is equal to the corresponding member of the same element in `pBuildInfo`.
  - For each element of either `pGeometries` or `ppGeometries` at a given index, with a `geometryType` member equal to `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, the `vertexFormat` and `indexType` members of `geometry.triangles` are equal to the corresponding members of the same element in `pBuildInfo`.
  - For each element of either `pGeometries` or `ppGeometries` at a given index, with a `geometryType` member equal to `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, the `maxVertex` member of `geometry.triangles` is less than or equal to the corresponding member of the same element in `pBuildInfo`.
  - For each element of either `pGeometries` or `ppGeometries` at a given index, with a `geometryType` member equal to `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if the applicable address in the `transformData` member of `geometry.triangles` is not `NULL`, the corresponding `transformData.hostAddress` parameter in `pBuildInfo` is not `NULL`.
  - For each element of either `pGeometries` or `ppGeometries` at a given index, with a `geometryType` member equal to `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if the `pNext` chain contains [VkAccelerationStructureTrianglesOpacityMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureTrianglesOpacityMicromapEXT.html) the corresponding member of `pBuildInfo` also contains [VkAccelerationStructureTrianglesOpacityMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureTrianglesOpacityMicromapEXT.html) and with an equivalent `micromap`.
  - For each element of either `pGeometries` or `ppGeometries` at a given index, with a `geometryType` member equal to `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if the `pNext` chain contains [VkAccelerationStructureTrianglesDisplacementMicromapNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureTrianglesDisplacementMicromapNV.html) the corresponding member of `pBuildInfo` also contains [VkAccelerationStructureTrianglesDisplacementMicromapNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureTrianglesDisplacementMicromapNV.html) and with an equivalent `micromap`.
- For each [VkAccelerationStructureBuildRangeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildRangeInfoKHR.html) corresponding to the [VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html):
  
  - Its `primitiveCount` member is less than or equal to the corresponding element of `pMaxPrimitiveCounts`.

Similarly, the `updateScratchSize` value will support any build command specifying the `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR` `mode` under the above conditions, and the `buildScratchSize` value will support any build command specifying the `VK_BUILD_ACCELERATION_STRUCTURE_MODE_BUILD_KHR` `mode` under the above conditions.

Valid Usage

- [](#VUID-vkGetAccelerationStructureBuildSizesKHR-accelerationStructure-08933)VUID-vkGetAccelerationStructureBuildSizesKHR-accelerationStructure-08933  
  The [`VkPhysicalDeviceAccelerationStructureFeaturesKHR`::`accelerationStructure`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-accelerationStructure) feature **must** be enabled
- [](#VUID-vkGetAccelerationStructureBuildSizesKHR-device-03618)VUID-vkGetAccelerationStructureBuildSizesKHR-device-03618  
  If `device` was created with multiple physical devices, then the [`bufferDeviceAddressMultiDevice`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-bufferDeviceAddressMultiDevice) feature **must** be enabled
- [](#VUID-vkGetAccelerationStructureBuildSizesKHR-pBuildInfo-03619)VUID-vkGetAccelerationStructureBuildSizesKHR-pBuildInfo-03619  
  If `pBuildInfo->geometryCount` is not `0`, `pMaxPrimitiveCounts` **must** be a valid pointer to an array of `pBuildInfo->geometryCount` `uint32_t` values
- [](#VUID-vkGetAccelerationStructureBuildSizesKHR-pBuildInfo-03785)VUID-vkGetAccelerationStructureBuildSizesKHR-pBuildInfo-03785  
  If `pBuildInfo->pGeometries` or `pBuildInfo->ppGeometries` has a `geometryType` of `VK_GEOMETRY_TYPE_INSTANCES_KHR`, each `pMaxPrimitiveCounts`\[i] **must** be less than or equal to [VkPhysicalDeviceAccelerationStructurePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceAccelerationStructurePropertiesKHR.html)::`maxInstanceCount`

Valid Usage (Implicit)

- [](#VUID-vkGetAccelerationStructureBuildSizesKHR-device-parameter)VUID-vkGetAccelerationStructureBuildSizesKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetAccelerationStructureBuildSizesKHR-buildType-parameter)VUID-vkGetAccelerationStructureBuildSizesKHR-buildType-parameter  
  `buildType` **must** be a valid [VkAccelerationStructureBuildTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildTypeKHR.html) value
- [](#VUID-vkGetAccelerationStructureBuildSizesKHR-pBuildInfo-parameter)VUID-vkGetAccelerationStructureBuildSizesKHR-pBuildInfo-parameter  
  `pBuildInfo` **must** be a valid pointer to a valid [VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html) structure
- [](#VUID-vkGetAccelerationStructureBuildSizesKHR-pMaxPrimitiveCounts-parameter)VUID-vkGetAccelerationStructureBuildSizesKHR-pMaxPrimitiveCounts-parameter  
  If `pMaxPrimitiveCounts` is not `NULL`, `pMaxPrimitiveCounts` **must** be a valid pointer to an array of `pBuildInfo->geometryCount` `uint32_t` values
- [](#VUID-vkGetAccelerationStructureBuildSizesKHR-pSizeInfo-parameter)VUID-vkGetAccelerationStructureBuildSizesKHR-pSizeInfo-parameter  
  `pSizeInfo` **must** be a valid pointer to a [VkAccelerationStructureBuildSizesInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildSizesInfoKHR.html) structure

## [](#_see_also)See Also

[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html), [VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html), [VkAccelerationStructureBuildSizesInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildSizesInfoKHR.html), [VkAccelerationStructureBuildTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildTypeKHR.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetAccelerationStructureBuildSizesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0