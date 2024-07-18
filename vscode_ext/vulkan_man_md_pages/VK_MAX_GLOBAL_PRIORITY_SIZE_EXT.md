# VK_MAX_GLOBAL_PRIORITY_SIZE_KHR(3) Manual Page

## Name

VK_MAX_GLOBAL_PRIORITY_SIZE_KHR - Length of an array of global queue
priorities



## <a href="#_c_specification" class="anchor"></a>C Specification

`VK_MAX_GLOBAL_PRIORITY_SIZE_KHR` is the length of an array of
[VkQueueGlobalPriorityKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueGlobalPriorityKHR.html) enumerants
representing supported queue priorities, as returned in
[VkQueueFamilyGlobalPriorityPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyGlobalPriorityPropertiesKHR.html)::`priorities`.

``` c
#define VK_MAX_GLOBAL_PRIORITY_SIZE_KHR   16U
```

or the equivalent

``` c
#define VK_MAX_GLOBAL_PRIORITY_SIZE_EXT   VK_MAX_GLOBAL_PRIORITY_SIZE_KHR
```

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_global_priority](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_global_priority.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_MAX_GLOBAL_PRIORITY_SIZE_KHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
