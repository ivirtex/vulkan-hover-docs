# VkMicromapTypeEXT(3) Manual Page

## Name

VkMicromapTypeEXT - Type of micromap



## [](#_c_specification)C Specification

Values which **can** be set in [VkMicromapCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapCreateInfoEXT.html)::`type` specifying the type of micromap, are:

```c++
// Provided by VK_EXT_opacity_micromap
typedef enum VkMicromapTypeEXT {
    VK_MICROMAP_TYPE_OPACITY_MICROMAP_EXT = 0,
#ifdef VK_ENABLE_BETA_EXTENSIONS
  // Provided by VK_NV_displacement_micromap
    VK_MICROMAP_TYPE_DISPLACEMENT_MICROMAP_NV = 1000397000,
#endif
} VkMicromapTypeEXT;
```

## [](#_description)Description

- `VK_MICROMAP_TYPE_OPACITY_MICROMAP_EXT` is a micromap containing data to control the opacity of a triangle.
- `VK_MICROMAP_TYPE_DISPLACEMENT_MICROMAP_NV` is a micromap containing data to control the displacement of subtriangles within a triangle.

## [](#_see_also)See Also

[VK\_EXT\_opacity\_micromap](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_opacity_micromap.html), [VkMicromapBuildInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapBuildInfoEXT.html), [VkMicromapCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapCreateInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMicromapTypeEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0