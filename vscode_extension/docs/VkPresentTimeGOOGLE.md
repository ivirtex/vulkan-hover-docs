# VkPresentTimeGOOGLE(3) Manual Page

## Name

VkPresentTimeGOOGLE - The earliest time image should be presented



## [](#_c_specification)C Specification

The `VkPresentTimeGOOGLE` structure is defined as:

```c++
// Provided by VK_GOOGLE_display_timing
typedef struct VkPresentTimeGOOGLE {
    uint32_t    presentID;
    uint64_t    desiredPresentTime;
} VkPresentTimeGOOGLE;
```

## [](#_members)Members

- `presentID` is an application-provided identification value, that **can** be used with the results of [vkGetPastPresentationTimingGOOGLE](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPastPresentationTimingGOOGLE.html), in order to uniquely identify this present. In order to be useful to the application, it **should** be unique within some period of time that is meaningful to the application.
- `desiredPresentTime` specifies that the image given **should** not be displayed to the user any earlier than this time. `desiredPresentTime` is a time in nanoseconds, relative to a monotonically-increasing clock (e.g. `CLOCK_MONOTONIC` (see clock\_gettime(2)) on Android and Linux). A value of zero specifies that the presentation engine **may** display the image at any time. This is useful when the application desires to provide `presentID`,

## [](#_description)Description

```
but does not need a specific pname:desiredPresentTime.
```

## [](#_see_also)See Also

[VK\_GOOGLE\_display\_timing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_GOOGLE_display_timing.html), [VkPresentTimesInfoGOOGLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentTimesInfoGOOGLE.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPresentTimeGOOGLE).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0