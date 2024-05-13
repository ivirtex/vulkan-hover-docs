# VK_WHOLE_SIZE(3) Manual Page

## Name

VK_WHOLE_SIZE - Sentinel value to use entire remaining array length



## <a href="#_c_specification" class="anchor"></a>C Specification

`VK_WHOLE_SIZE` is a special value indicating that the entire remaining
length of a buffer following a given `offset` should be used. It **can**
be specified for
[VkBufferMemoryBarrier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferMemoryBarrier.html)::`size` and other
structures.

``` c
#define VK_WHOLE_SIZE                     (~0ULL)
```

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_WHOLE_SIZE"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
