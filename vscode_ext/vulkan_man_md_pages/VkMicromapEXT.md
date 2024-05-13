# VkMicromapEXT(3) Manual Page

## Name

VkMicromapEXT - Opaque handle to a micromap object



## <a href="#_c_specification" class="anchor"></a>C Specification

Micromaps are opaque data structures that are built by the
implementation to encode sub-triangle data to be included in an
acceleration structure.

Micromaps are represented by `VkMicromapEXT` handles:

``` c
// Provided by VK_EXT_opacity_micromap
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkMicromapEXT)
```

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_opacity_micromap](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_opacity_micromap.html),
[VkAccelerationStructureTrianglesDisplacementMicromapNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureTrianglesDisplacementMicromapNV.html),
[VkAccelerationStructureTrianglesOpacityMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureTrianglesOpacityMicromapEXT.html),
[VkCopyMemoryToMicromapInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMemoryToMicromapInfoEXT.html),
[VkCopyMicromapInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMicromapInfoEXT.html),
[VkCopyMicromapToMemoryInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMicromapToMemoryInfoEXT.html),
[VkMicromapBuildInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapBuildInfoEXT.html),
[vkCmdWriteMicromapsPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWriteMicromapsPropertiesEXT.html),
[vkCreateMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateMicromapEXT.html),
[vkDestroyMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyMicromapEXT.html),
[vkWriteMicromapsPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkWriteMicromapsPropertiesEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMicromapEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
