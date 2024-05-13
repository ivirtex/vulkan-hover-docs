# VkRefreshCycleDurationGOOGLE(3) Manual Page

## Name

VkRefreshCycleDurationGOOGLE - Structure containing the RC duration of a
display



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkRefreshCycleDurationGOOGLE` structure is defined as:

``` c
// Provided by VK_GOOGLE_display_timing
typedef struct VkRefreshCycleDurationGOOGLE {
    uint64_t    refreshDuration;
} VkRefreshCycleDurationGOOGLE;
```

## <a href="#_members" class="anchor"></a>Members

- `refreshDuration` is the number of nanoseconds from the start of one
  refresh cycle to the next.

## <a href="#_description" class="anchor"></a>Description

## <a href="#_see_also" class="anchor"></a>See Also

[VK_GOOGLE_display_timing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_GOOGLE_display_timing.html),
[vkGetRefreshCycleDurationGOOGLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetRefreshCycleDurationGOOGLE.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkRefreshCycleDurationGOOGLE"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
