# VkMicromapUsageEXT(3) Manual Page

## Name

VkMicromapUsageEXT - Structure specifying the usage information used to build a micromap



## [](#_c_specification)C Specification

The `VkMicromapUsageEXT` structure is defined as:

```c++
// Provided by VK_EXT_opacity_micromap
typedef struct VkMicromapUsageEXT {
    uint32_t    count;
    uint32_t    subdivisionLevel;
    uint32_t    format;
} VkMicromapUsageEXT;
```

## [](#_members)Members

- `count` is the number of triangles in the usage format defined by the `subdivisionLevel` and `format` below in the micromap
- `subdivisionLevel` is the subdivision level of this usage format
- `format` is the format of this usage format

## [](#_description)Description

Valid Usage

- [](#VUID-VkMicromapUsageEXT-format-07519)VUID-VkMicromapUsageEXT-format-07519  
  If the [VkMicromapTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapTypeEXT.html) of the micromap is `VK_MICROMAP_TYPE_OPACITY_MICROMAP_EXT` then `format` **must** be `VK_OPACITY_MICROMAP_FORMAT_2_STATE_EXT` or `VK_OPACITY_MICROMAP_FORMAT_4_STATE_EXT`
- [](#VUID-VkMicromapUsageEXT-format-07520)VUID-VkMicromapUsageEXT-format-07520  
  If the [VkMicromapTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapTypeEXT.html) of the micromap is `VK_MICROMAP_TYPE_OPACITY_MICROMAP_EXT` and `format` is `VK_OPACITY_MICROMAP_FORMAT_2_STATE_EXT` then `subdivisionLevel` **must** be less than or equal to [VkPhysicalDeviceOpacityMicromapPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceOpacityMicromapPropertiesEXT.html)::`maxOpacity2StateSubdivisionLevel`
- [](#VUID-VkMicromapUsageEXT-format-07521)VUID-VkMicromapUsageEXT-format-07521  
  If the [VkMicromapTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapTypeEXT.html) of the micromap is `VK_MICROMAP_TYPE_OPACITY_MICROMAP_EXT` and `format` is `VK_OPACITY_MICROMAP_FORMAT_4_STATE_EXT` then `subdivisionLevel` **must** be less than or equal to [VkPhysicalDeviceOpacityMicromapPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceOpacityMicromapPropertiesEXT.html)::`maxOpacity4StateSubdivisionLevel`
- [](#VUID-VkMicromapUsageEXT-format-08706)VUID-VkMicromapUsageEXT-format-08706  
  If the [VkMicromapTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapTypeEXT.html) of the micromap is `VK_MICROMAP_TYPE_DISPLACEMENT_MICROMAP_NV` then `format` **must** be `VK_DISPLACEMENT_MICROMAP_FORMAT_64_TRIANGLES_64_BYTES_NV`, `VK_DISPLACEMENT_MICROMAP_FORMAT_256_TRIANGLES_128_BYTES_NV` or `VK_DISPLACEMENT_MICROMAP_FORMAT_1024_TRIANGLES_128_BYTES_NV`
- [](#VUID-VkMicromapUsageEXT-subdivisionLevel-08707)VUID-VkMicromapUsageEXT-subdivisionLevel-08707  
  If the [VkMicromapTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapTypeEXT.html) of the micromap is `VK_MICROMAP_TYPE_DISPLACEMENT_MICROMAP_NV` then `subdivisionLevel` **must** be less than or equal to [VkPhysicalDeviceDisplacementMicromapPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDisplacementMicromapPropertiesNV.html)::`maxDisplacementMicromapSubdivisionLevel`

The `format` is interpreted based on the `type` of the micromap using it.

## [](#_see_also)See Also

[VK\_EXT\_opacity\_micromap](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_opacity_micromap.html), [VkAccelerationStructureTrianglesDisplacementMicromapNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureTrianglesDisplacementMicromapNV.html), [VkAccelerationStructureTrianglesOpacityMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureTrianglesOpacityMicromapEXT.html), [VkMicromapBuildInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapBuildInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMicromapUsageEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0