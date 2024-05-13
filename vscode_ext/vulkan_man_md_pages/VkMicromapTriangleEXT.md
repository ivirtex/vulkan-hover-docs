# VkMicromapTriangleEXT(3) Manual Page

## Name

VkMicromapTriangleEXT - Structure specifying the micromap format and
data for a triangle



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkMicromapTriangleEXT` structure is defined as:

``` c
// Provided by VK_EXT_opacity_micromap
typedef struct VkMicromapTriangleEXT {
    uint32_t    dataOffset;
    uint16_t    subdivisionLevel;
    uint16_t    format;
} VkMicromapTriangleEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `dataOffset` is the offset in bytes of the start of the data for this
  triangle. This is a byte aligned value.

- `subdivisionLevel` is the subdivision level of this triangle

- `format` is the format of this triangle

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkMicromapTriangleEXT-format-07522"
  id="VUID-VkMicromapTriangleEXT-format-07522"></a>
  VUID-VkMicromapTriangleEXT-format-07522  
  If the [VkMicromapTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapTypeEXT.html) of the micromap is
  `VK_MICROMAP_TYPE_OPACITY_MICROMAP_EXT` then `format` **must** be
  `VK_OPACITY_MICROMAP_FORMAT_2_STATE_EXT` or
  `VK_OPACITY_MICROMAP_FORMAT_4_STATE_EXT`

- <a href="#VUID-VkMicromapTriangleEXT-format-07523"
  id="VUID-VkMicromapTriangleEXT-format-07523"></a>
  VUID-VkMicromapTriangleEXT-format-07523  
  If the [VkMicromapTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapTypeEXT.html) of the micromap is
  `VK_MICROMAP_TYPE_OPACITY_MICROMAP_EXT` and `format` is
  `VK_OPACITY_MICROMAP_FORMAT_2_STATE_EXT` then `subdivisionLevel`
  **must** be less than or equal to
  [VkPhysicalDeviceOpacityMicromapPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceOpacityMicromapPropertiesEXT.html)::`maxOpacity2StateSubdivisionLevel`

- <a href="#VUID-VkMicromapTriangleEXT-format-07524"
  id="VUID-VkMicromapTriangleEXT-format-07524"></a>
  VUID-VkMicromapTriangleEXT-format-07524  
  If the [VkMicromapTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapTypeEXT.html) of the micromap is
  `VK_MICROMAP_TYPE_OPACITY_MICROMAP_EXT` and `format` is
  `VK_OPACITY_MICROMAP_FORMAT_4_STATE_EXT` then `subdivisionLevel`
  **must** be less than or equal to
  [VkPhysicalDeviceOpacityMicromapPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceOpacityMicromapPropertiesEXT.html)::`maxOpacity4StateSubdivisionLevel`

- <a href="#VUID-VkMicromapTriangleEXT-format-08708"
  id="VUID-VkMicromapTriangleEXT-format-08708"></a>
  VUID-VkMicromapTriangleEXT-format-08708  
  If the [VkMicromapTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapTypeEXT.html) of the micromap is
  `VK_MICROMAP_TYPE_DISPLACEMENT_MICROMAP_NV` then `format` **must** be
  `VK_DISPLACEMENT_MICROMAP_FORMAT_64_TRIANGLES_64_BYTES_NV`,
  `VK_DISPLACEMENT_MICROMAP_FORMAT_256_TRIANGLES_128_BYTES_NV` or
  `VK_DISPLACEMENT_MICROMAP_FORMAT_1024_TRIANGLES_128_BYTES_NV`

- <a href="#VUID-VkMicromapTriangleEXT-subdivisionLevel-08709"
  id="VUID-VkMicromapTriangleEXT-subdivisionLevel-08709"></a>
  VUID-VkMicromapTriangleEXT-subdivisionLevel-08709  
  If the [VkMicromapTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapTypeEXT.html) of the micromap is
  `VK_MICROMAP_TYPE_DISPLACEMENT_MICROMAP_NV` then `subdivisionLevel`
  **must** be less than or equal to
  [VkPhysicalDeviceDisplacementMicromapPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDisplacementMicromapPropertiesNV.html)::`maxDisplacementMicromapSubdivisionLevel`

The `format` is interpreted based on the `type` of the micromap using
it.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_opacity_micromap](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_opacity_micromap.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMicromapTriangleEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
