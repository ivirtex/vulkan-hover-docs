# VkMicromapUsageEXT(3) Manual Page

## Name

VkMicromapUsageEXT - Structure specifying the usage information used to
build a micromap



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkMicromapUsageEXT` structure is defined as:

``` c
// Provided by VK_EXT_opacity_micromap
typedef struct VkMicromapUsageEXT {
    uint32_t    count;
    uint32_t    subdivisionLevel;
    uint32_t    format;
} VkMicromapUsageEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `count` is the number of triangles in the usage format defined by the
  `subdivisionLevel` and `format` below in the micromap

- `subdivisionLevel` is the subdivision level of this usage format

- `format` is the format of this usage format

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkMicromapUsageEXT-format-07519"
  id="VUID-VkMicromapUsageEXT-format-07519"></a>
  VUID-VkMicromapUsageEXT-format-07519  
  If the [VkMicromapTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapTypeEXT.html) of the micromap is
  `VK_MICROMAP_TYPE_OPACITY_MICROMAP_EXT` then `format` **must** be
  `VK_OPACITY_MICROMAP_FORMAT_2_STATE_EXT` or
  `VK_OPACITY_MICROMAP_FORMAT_4_STATE_EXT`

- <a href="#VUID-VkMicromapUsageEXT-format-07520"
  id="VUID-VkMicromapUsageEXT-format-07520"></a>
  VUID-VkMicromapUsageEXT-format-07520  
  If the [VkMicromapTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapTypeEXT.html) of the micromap is
  `VK_MICROMAP_TYPE_OPACITY_MICROMAP_EXT` and `format` is
  `VK_OPACITY_MICROMAP_FORMAT_2_STATE_EXT` then `subdivisionLevel`
  **must** be less than or equal to
  [VkPhysicalDeviceOpacityMicromapPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceOpacityMicromapPropertiesEXT.html)::`maxOpacity2StateSubdivisionLevel`

- <a href="#VUID-VkMicromapUsageEXT-format-07521"
  id="VUID-VkMicromapUsageEXT-format-07521"></a>
  VUID-VkMicromapUsageEXT-format-07521  
  If the [VkMicromapTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapTypeEXT.html) of the micromap is
  `VK_MICROMAP_TYPE_OPACITY_MICROMAP_EXT` and `format` is
  `VK_OPACITY_MICROMAP_FORMAT_4_STATE_EXT` then `subdivisionLevel`
  **must** be less than or equal to
  [VkPhysicalDeviceOpacityMicromapPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceOpacityMicromapPropertiesEXT.html)::`maxOpacity4StateSubdivisionLevel`

- <a href="#VUID-VkMicromapUsageEXT-format-08706"
  id="VUID-VkMicromapUsageEXT-format-08706"></a>
  VUID-VkMicromapUsageEXT-format-08706  
  If the [VkMicromapTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapTypeEXT.html) of the micromap is
  `VK_MICROMAP_TYPE_DISPLACEMENT_MICROMAP_NV` then `format` **must** be
  `VK_DISPLACEMENT_MICROMAP_FORMAT_64_TRIANGLES_64_BYTES_NV`,
  `VK_DISPLACEMENT_MICROMAP_FORMAT_256_TRIANGLES_128_BYTES_NV` or
  `VK_DISPLACEMENT_MICROMAP_FORMAT_1024_TRIANGLES_128_BYTES_NV`

- <a href="#VUID-VkMicromapUsageEXT-subdivisionLevel-08707"
  id="VUID-VkMicromapUsageEXT-subdivisionLevel-08707"></a>
  VUID-VkMicromapUsageEXT-subdivisionLevel-08707  
  If the [VkMicromapTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapTypeEXT.html) of the micromap is
  `VK_MICROMAP_TYPE_DISPLACEMENT_MICROMAP_NV` then `subdivisionLevel`
  **must** be less than or equal to
  [VkPhysicalDeviceDisplacementMicromapPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDisplacementMicromapPropertiesNV.html)::`maxDisplacementMicromapSubdivisionLevel`

The `format` is interpreted based on the `type` of the micromap using
it.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_opacity_micromap](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_opacity_micromap.html),
[VkAccelerationStructureTrianglesDisplacementMicromapNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureTrianglesDisplacementMicromapNV.html),
[VkAccelerationStructureTrianglesOpacityMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureTrianglesOpacityMicromapEXT.html),
[VkMicromapBuildInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapBuildInfoEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMicromapUsageEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
