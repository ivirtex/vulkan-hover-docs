# VkMicromapTriangleEXT(3) Manual Page

## Name

VkMicromapTriangleEXT - Structure specifying the micromap format and data for a triangle



## [](#_c_specification)C Specification

The `VkMicromapTriangleEXT` structure is defined as:

```c++
// Provided by VK_EXT_opacity_micromap
typedef struct VkMicromapTriangleEXT {
    uint32_t    dataOffset;
    uint16_t    subdivisionLevel;
    uint16_t    format;
} VkMicromapTriangleEXT;
```

## [](#_members)Members

- `dataOffset` is the offset in bytes of the start of the data for this triangle. This is a byte aligned value.
- `subdivisionLevel` is the subdivision level of this triangle
- `format` is the format of this triangle

## [](#_description)Description

Valid Usage

- [](#VUID-VkMicromapTriangleEXT-format-07522)VUID-VkMicromapTriangleEXT-format-07522  
  If the [VkMicromapTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapTypeEXT.html) of the micromap is `VK_MICROMAP_TYPE_OPACITY_MICROMAP_EXT` then `format` **must** be `VK_OPACITY_MICROMAP_FORMAT_2_STATE_EXT` or `VK_OPACITY_MICROMAP_FORMAT_4_STATE_EXT`
- [](#VUID-VkMicromapTriangleEXT-format-07523)VUID-VkMicromapTriangleEXT-format-07523  
  If the [VkMicromapTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapTypeEXT.html) of the micromap is `VK_MICROMAP_TYPE_OPACITY_MICROMAP_EXT` and `format` is `VK_OPACITY_MICROMAP_FORMAT_2_STATE_EXT` then `subdivisionLevel` **must** be less than or equal to [VkPhysicalDeviceOpacityMicromapPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceOpacityMicromapPropertiesEXT.html)::`maxOpacity2StateSubdivisionLevel`
- [](#VUID-VkMicromapTriangleEXT-format-07524)VUID-VkMicromapTriangleEXT-format-07524  
  If the [VkMicromapTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapTypeEXT.html) of the micromap is `VK_MICROMAP_TYPE_OPACITY_MICROMAP_EXT` and `format` is `VK_OPACITY_MICROMAP_FORMAT_4_STATE_EXT` then `subdivisionLevel` **must** be less than or equal to [VkPhysicalDeviceOpacityMicromapPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceOpacityMicromapPropertiesEXT.html)::`maxOpacity4StateSubdivisionLevel`
- [](#VUID-VkMicromapTriangleEXT-format-08708)VUID-VkMicromapTriangleEXT-format-08708  
  If the [VkMicromapTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapTypeEXT.html) of the micromap is `VK_MICROMAP_TYPE_DISPLACEMENT_MICROMAP_NV` then `format` **must** be `VK_DISPLACEMENT_MICROMAP_FORMAT_64_TRIANGLES_64_BYTES_NV`, `VK_DISPLACEMENT_MICROMAP_FORMAT_256_TRIANGLES_128_BYTES_NV` or `VK_DISPLACEMENT_MICROMAP_FORMAT_1024_TRIANGLES_128_BYTES_NV`
- [](#VUID-VkMicromapTriangleEXT-subdivisionLevel-08709)VUID-VkMicromapTriangleEXT-subdivisionLevel-08709  
  If the [VkMicromapTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapTypeEXT.html) of the micromap is `VK_MICROMAP_TYPE_DISPLACEMENT_MICROMAP_NV` then `subdivisionLevel` **must** be less than or equal to [VkPhysicalDeviceDisplacementMicromapPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDisplacementMicromapPropertiesNV.html)::`maxDisplacementMicromapSubdivisionLevel`

The `format` is interpreted based on the `type` of the micromap using it.

## [](#_see_also)See Also

[VK\_EXT\_opacity\_micromap](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_opacity_micromap.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMicromapTriangleEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0