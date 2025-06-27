# VkPresentTimesInfoGOOGLE(3) Manual Page

## Name

VkPresentTimesInfoGOOGLE - The earliest time each image should be
presented



## <a href="#_c_specification" class="anchor"></a>C Specification

When the [`VK_GOOGLE_display_timing`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_GOOGLE_display_timing.html)
extension is enabled, additional fields **can** be specified that allow
an application to specify the earliest time that an image should be
displayed. This allows an application to avoid stutter that is caused by
an image being displayed earlier than planned. Such stuttering can occur
with both fixed and variable-refresh-rate displays, because stuttering
occurs when the geometry is not correctly positioned for when the image
is displayed. An application **can** instruct the presentation engine
that an image should not be displayed earlier than a specified time by
adding a `VkPresentTimesInfoGOOGLE` structure to the `pNext` chain of
the `VkPresentInfoKHR` structure.

The `VkPresentTimesInfoGOOGLE` structure is defined as:

``` c
// Provided by VK_GOOGLE_display_timing
typedef struct VkPresentTimesInfoGOOGLE {
    VkStructureType               sType;
    const void*                   pNext;
    uint32_t                      swapchainCount;
    const VkPresentTimeGOOGLE*    pTimes;
} VkPresentTimesInfoGOOGLE;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `swapchainCount` is the number of swapchains being presented to by
  this command.

- `pTimes` is `NULL` or a pointer to an array of `VkPresentTimeGOOGLE`
  elements with `swapchainCount` entries. If not `NULL`, each element of
  `pTimes` contains the earliest time to present the image corresponding
  to the entry in the `VkPresentInfoKHR`::`pImageIndices` array.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkPresentTimesInfoGOOGLE-swapchainCount-01247"
  id="VUID-VkPresentTimesInfoGOOGLE-swapchainCount-01247"></a>
  VUID-VkPresentTimesInfoGOOGLE-swapchainCount-01247  
  `swapchainCount` **must** be the same value as
  `VkPresentInfoKHR`::`swapchainCount`, where `VkPresentInfoKHR` is
  included in the `pNext` chain of this `VkPresentTimesInfoGOOGLE`
  structure

Valid Usage (Implicit)

- <a href="#VUID-VkPresentTimesInfoGOOGLE-sType-sType"
  id="VUID-VkPresentTimesInfoGOOGLE-sType-sType"></a>
  VUID-VkPresentTimesInfoGOOGLE-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PRESENT_TIMES_INFO_GOOGLE`

- <a href="#VUID-VkPresentTimesInfoGOOGLE-pTimes-parameter"
  id="VUID-VkPresentTimesInfoGOOGLE-pTimes-parameter"></a>
  VUID-VkPresentTimesInfoGOOGLE-pTimes-parameter  
  If `pTimes` is not `NULL`, `pTimes` **must** be a valid pointer to an
  array of `swapchainCount`
  [VkPresentTimeGOOGLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentTimeGOOGLE.html) structures

- <a href="#VUID-VkPresentTimesInfoGOOGLE-swapchainCount-arraylength"
  id="VUID-VkPresentTimesInfoGOOGLE-swapchainCount-arraylength"></a>
  VUID-VkPresentTimesInfoGOOGLE-swapchainCount-arraylength  
  `swapchainCount` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_GOOGLE_display_timing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_GOOGLE_display_timing.html),
[VkPresentTimeGOOGLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentTimeGOOGLE.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPresentTimesInfoGOOGLE"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
