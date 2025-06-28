# VK\_MAX\_GLOBAL\_PRIORITY\_SIZE(3) Manual Page

## Name

VK\_MAX\_GLOBAL\_PRIORITY\_SIZE - Length of an array of global queue priorities



## [](#_c_specification)C Specification

`VK_MAX_GLOBAL_PRIORITY_SIZE` is the length of an array of [VkQueueGlobalPriority](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueGlobalPriority.html) enumerants representing supported queue priorities, as returned in [VkQueueFamilyGlobalPriorityProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyGlobalPriorityProperties.html)::`priorities`.

```c++
#define VK_MAX_GLOBAL_PRIORITY_SIZE       16U
```

or the equivalent

```c++
#define VK_MAX_GLOBAL_PRIORITY_SIZE_KHR   VK_MAX_GLOBAL_PRIORITY_SIZE
```

or the equivalent

```c++
#define VK_MAX_GLOBAL_PRIORITY_SIZE_EXT   VK_MAX_GLOBAL_PRIORITY_SIZE
```

## [](#_see_also)See Also

[VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_MAX_GLOBAL_PRIORITY_SIZE)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0