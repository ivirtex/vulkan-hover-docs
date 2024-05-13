# VkPresentTimeGOOGLE(3) Manual Page

## Name

VkPresentTimeGOOGLE - The earliest time image should be presented



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPresentTimeGOOGLE` structure is defined as:

``` c
// Provided by VK_GOOGLE_display_timing
typedef struct VkPresentTimeGOOGLE {
    uint32_t    presentID;
    uint64_t    desiredPresentTime;
} VkPresentTimeGOOGLE;
```

## <a href="#_members" class="anchor"></a>Members

- `presentID` is an application-provided identification value, that
  **can** be used with the results of
  [vkGetPastPresentationTimingGOOGLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPastPresentationTimingGOOGLE.html),
  in order to uniquely identify this present. In order to be useful to
  the application, it **should** be unique within some period of time
  that is meaningful to the application.

- `desiredPresentTime` specifies that the image given **should** not be
  displayed to the user any earlier than this time. `desiredPresentTime`
  is a time in nanoseconds, relative to a monotonically-increasing clock
  (e.g. `CLOCK_MONOTONIC` (see clock_gettime(2)) on Android and Linux).
  A value of zero specifies that the presentation engine **may** display
  the image at any time. This is useful when the application desires to
  provide `presentID`,

## <a href="#_description" class="anchor"></a>Description

    but does not need a specific pname:desiredPresentTime.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_GOOGLE_display_timing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_GOOGLE_display_timing.html),
[VkPresentTimesInfoGOOGLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentTimesInfoGOOGLE.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPresentTimeGOOGLE"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
