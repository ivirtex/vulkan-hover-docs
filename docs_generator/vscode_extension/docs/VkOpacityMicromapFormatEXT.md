# VkOpacityMicromapFormatEXT(3) Manual Page

## Name

VkOpacityMicromapFormatEXT - Format enum for opacity micromaps



## [](#_c_specification)C Specification

Formats which **can** be set in [VkMicromapUsageEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapUsageEXT.html)::`format` and [VkMicromapTriangleEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapTriangleEXT.html)::`format` for micromap builds, are:

```c++
// Provided by VK_EXT_opacity_micromap
typedef enum VkOpacityMicromapFormatEXT {
    VK_OPACITY_MICROMAP_FORMAT_2_STATE_EXT = 1,
    VK_OPACITY_MICROMAP_FORMAT_4_STATE_EXT = 2,
} VkOpacityMicromapFormatEXT;
```

## [](#_description)Description

- `VK_OPACITY_MICROMAP_FORMAT_2_STATE_EXT` specifies that the given micromap format has one bit per subtriangle encoding either fully opaque or fully transparent.
- `VK_OPACITY_MICROMAP_FORMAT_4_STATE_EXT` specifies that the given micromap format has two bits per subtriangle encoding four modes which can be interpreted as described in [ray traversal](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#ray-opacity-micromap).

Note

For compactness, these values are stored as 16-bit in some structures.

## [](#_see_also)See Also

[VK\_EXT\_opacity\_micromap](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_opacity_micromap.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkOpacityMicromapFormatEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0