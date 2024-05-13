# VK_QUEUE_FAMILY_FOREIGN_EXT(3) Manual Page

## Name

VK_QUEUE_FAMILY_FOREIGN_EXT - Foreign queue family index sentinel



## <a href="#_c_specification" class="anchor"></a>C Specification

The special queue family index `VK_QUEUE_FAMILY_FOREIGN_EXT` represents
any queue external to the resource’s current Vulkan instance, regardless
of the queue’s underlying physical device or driver version. This
includes, for example, queues for fixed-function image processing
devices, media codec devices, and display devices, as well as all queues
that use the same underlying device group or physical device, and the
same driver version as the resource’s [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html).

``` c
#define VK_QUEUE_FAMILY_FOREIGN_EXT       (~2U)
```

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_queue_family_foreign](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_queue_family_foreign.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_QUEUE_FAMILY_FOREIGN_EXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
