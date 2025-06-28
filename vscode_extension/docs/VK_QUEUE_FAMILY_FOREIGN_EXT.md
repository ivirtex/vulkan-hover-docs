# VK\_QUEUE\_FAMILY\_FOREIGN\_EXT(3) Manual Page

## Name

VK\_QUEUE\_FAMILY\_FOREIGN\_EXT - Foreign queue family index sentinel



## [](#_c_specification)C Specification

The special queue family index `VK_QUEUE_FAMILY_FOREIGN_EXT` represents any queue external to the resource’s current Vulkan instance, regardless of the queue’s underlying physical device or driver version. This includes, for example, queues for fixed-function image processing devices, media codec devices, and display devices, as well as all queues that use the same underlying device group or physical device, and the same driver version as the resource’s [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

```c++
#define VK_QUEUE_FAMILY_FOREIGN_EXT       (~2U)
```

## [](#_see_also)See Also

[VK\_EXT\_queue\_family\_foreign](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_queue_family_foreign.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_QUEUE_FAMILY_FOREIGN_EXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0