# VkRefreshCycleDurationGOOGLE(3) Manual Page

## Name

VkRefreshCycleDurationGOOGLE - Structure containing the RC duration of a display



## [](#_c_specification)C Specification

The `VkRefreshCycleDurationGOOGLE` structure is defined as:

```c++
// Provided by VK_GOOGLE_display_timing
typedef struct VkRefreshCycleDurationGOOGLE {
    uint64_t    refreshDuration;
} VkRefreshCycleDurationGOOGLE;
```

## [](#_members)Members

- `refreshDuration` is the number of nanoseconds from the start of one refresh cycle to the next.

## [](#_description)Description

## [](#_see_also)See Also

[VK\_GOOGLE\_display\_timing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_GOOGLE_display_timing.html), [vkGetRefreshCycleDurationGOOGLE](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetRefreshCycleDurationGOOGLE.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkRefreshCycleDurationGOOGLE).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0