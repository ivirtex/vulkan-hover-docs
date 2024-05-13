# vkGetPastPresentationTimingGOOGLE(3) Manual Page

## Name

vkGetPastPresentationTimingGOOGLE - Obtain timing of a
previously-presented image



## <a href="#_c_specification" class="anchor"></a>C Specification

The implementation will maintain a limited amount of history of timing
information about previous presents. Because of the asynchronous nature
of the presentation engine, the timing information for a given
[vkQueuePresentKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueuePresentKHR.html) command will become
available some time later. These time values can be asynchronously
queried, and will be returned if available. All time values are in
nanoseconds, relative to a monotonically-increasing clock (e.g.
`CLOCK_MONOTONIC` (see clock_gettime(2)) on Android and Linux).

To asynchronously query the presentation engine, for newly-available
timing information about one or more previous presents to a given
swapchain, call:

``` c
// Provided by VK_GOOGLE_display_timing
VkResult vkGetPastPresentationTimingGOOGLE(
    VkDevice                                    device,
    VkSwapchainKHR                              swapchain,
    uint32_t*                                   pPresentationTimingCount,
    VkPastPresentationTimingGOOGLE*             pPresentationTimings);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device associated with `swapchain`.

- `swapchain` is the swapchain to obtain presentation timing information
  duration for.

- `pPresentationTimingCount` is a pointer to an integer related to the
  number of `VkPastPresentationTimingGOOGLE` structures to query, as
  described below.

- `pPresentationTimings` is either `NULL` or a pointer to an array of
  `VkPastPresentationTimingGOOGLE` structures.

## <a href="#_description" class="anchor"></a>Description

If `pPresentationTimings` is `NULL`, then the number of newly-available
timing records for the given `swapchain` is returned in
`pPresentationTimingCount`. Otherwise, `pPresentationTimingCount`
**must** point to a variable set by the user to the number of elements
in the `pPresentationTimings` array, and on return the variable is
overwritten with the number of structures actually written to
`pPresentationTimings`. If the value of `pPresentationTimingCount` is
less than the number of newly-available timing records, at most
`pPresentationTimingCount` structures will be written, and
`VK_INCOMPLETE` will be returned instead of `VK_SUCCESS`, to indicate
that not all the available timing records were returned.

Valid Usage (Implicit)

- <a href="#VUID-vkGetPastPresentationTimingGOOGLE-device-parameter"
  id="VUID-vkGetPastPresentationTimingGOOGLE-device-parameter"></a>
  VUID-vkGetPastPresentationTimingGOOGLE-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetPastPresentationTimingGOOGLE-swapchain-parameter"
  id="VUID-vkGetPastPresentationTimingGOOGLE-swapchain-parameter"></a>
  VUID-vkGetPastPresentationTimingGOOGLE-swapchain-parameter  
  `swapchain` **must** be a valid [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html)
  handle

- <a
  href="#VUID-vkGetPastPresentationTimingGOOGLE-pPresentationTimingCount-parameter"
  id="VUID-vkGetPastPresentationTimingGOOGLE-pPresentationTimingCount-parameter"></a>
  VUID-vkGetPastPresentationTimingGOOGLE-pPresentationTimingCount-parameter  
  `pPresentationTimingCount` **must** be a valid pointer to a `uint32_t`
  value

- <a
  href="#VUID-vkGetPastPresentationTimingGOOGLE-pPresentationTimings-parameter"
  id="VUID-vkGetPastPresentationTimingGOOGLE-pPresentationTimings-parameter"></a>
  VUID-vkGetPastPresentationTimingGOOGLE-pPresentationTimings-parameter  
  If the value referenced by `pPresentationTimingCount` is not `0`, and
  `pPresentationTimings` is not `NULL`, `pPresentationTimings` **must**
  be a valid pointer to an array of `pPresentationTimingCount`
  [VkPastPresentationTimingGOOGLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPastPresentationTimingGOOGLE.html)
  structures

- <a href="#VUID-vkGetPastPresentationTimingGOOGLE-swapchain-parent"
  id="VUID-vkGetPastPresentationTimingGOOGLE-swapchain-parent"></a>
  VUID-vkGetPastPresentationTimingGOOGLE-swapchain-parent  
  `swapchain` **must** have been created, allocated, or retrieved from
  `device`

Host Synchronization

- Host access to `swapchain` **must** be externally synchronized

Return Codes

On success, this command returns  
- `VK_SUCCESS`

- `VK_INCOMPLETE`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_DEVICE_LOST`

- `VK_ERROR_OUT_OF_DATE_KHR`

- `VK_ERROR_SURFACE_LOST_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_GOOGLE_display_timing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_GOOGLE_display_timing.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkPastPresentationTimingGOOGLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPastPresentationTimingGOOGLE.html),
[VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPastPresentationTimingGOOGLE"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
