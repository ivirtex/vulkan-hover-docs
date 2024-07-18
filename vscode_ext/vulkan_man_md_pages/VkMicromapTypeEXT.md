# VkMicromapTypeEXT(3) Manual Page

## Name

VkMicromapTypeEXT - Type of micromap



## <a href="#_c_specification" class="anchor"></a>C Specification

Values which **can** be set in
[VkMicromapCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapCreateInfoEXT.html)::`type`
specifying the type of micromap, are:

``` c
// Provided by VK_EXT_opacity_micromap
typedef enum VkMicromapTypeEXT {
    VK_MICROMAP_TYPE_OPACITY_MICROMAP_EXT = 0,
#ifdef VK_ENABLE_BETA_EXTENSIONS
  // Provided by VK_NV_displacement_micromap
    VK_MICROMAP_TYPE_DISPLACEMENT_MICROMAP_NV = 1000397000,
#endif
} VkMicromapTypeEXT;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_MICROMAP_TYPE_OPACITY_MICROMAP_EXT` is a micromap containing data
  to control the opacity of a triangle.

- `VK_MICROMAP_TYPE_DISPLACEMENT_MICROMAP_NV` is a micromap containing
  data to control the displacement of subtriangles within a triangle.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_opacity_micromap](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_opacity_micromap.html),
[VkMicromapBuildInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapBuildInfoEXT.html),
[VkMicromapCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapCreateInfoEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMicromapTypeEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
