# VkMicromapEXT(3) Manual Page

## Name

VkMicromapEXT - Opaque handle to a micromap object



## [](#_c_specification)C Specification

Micromaps are opaque data structures that are built by the implementation to encode sub-triangle data to be included in an acceleration structure.

Micromaps are represented by `VkMicromapEXT` handles:

```c++
// Provided by VK_EXT_opacity_micromap
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkMicromapEXT)
```

## [](#_see_also)See Also

[VK\_DEFINE\_NON\_DISPATCHABLE\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_DEFINE_NON_DISPATCHABLE_HANDLE.html), [VK\_EXT\_opacity\_micromap](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_opacity_micromap.html), [VkAccelerationStructureTrianglesDisplacementMicromapNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureTrianglesDisplacementMicromapNV.html), [VkAccelerationStructureTrianglesOpacityMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureTrianglesOpacityMicromapEXT.html), [VkCopyMemoryToMicromapInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryToMicromapInfoEXT.html), [VkCopyMicromapInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMicromapInfoEXT.html), [VkCopyMicromapToMemoryInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMicromapToMemoryInfoEXT.html), [VkMicromapBuildInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapBuildInfoEXT.html), [vkCmdWriteMicromapsPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWriteMicromapsPropertiesEXT.html), [vkCreateMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateMicromapEXT.html), [vkDestroyMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyMicromapEXT.html), [vkWriteMicromapsPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkWriteMicromapsPropertiesEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMicromapEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0