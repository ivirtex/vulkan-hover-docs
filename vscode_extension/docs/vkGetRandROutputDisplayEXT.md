# vkGetRandROutputDisplayEXT(3) Manual Page

## Name

vkGetRandROutputDisplayEXT - Query the VkDisplayKHR corresponding to an X11 RandR Output



## [](#_c_specification)C Specification

When acquiring displays from an X11 server, an application may also wish to enumerate and identify them using a native handle rather than a `VkDisplayKHR` handle. To determine the `VkDisplayKHR` handle corresponding to an X11 RandR Output, call:

```c++
// Provided by VK_EXT_acquire_xlib_display
VkResult vkGetRandROutputDisplayEXT(
    VkPhysicalDevice                            physicalDevice,
    Display*                                    dpy,
    RROutput                                    rrOutput,
    VkDisplayKHR*                               pDisplay);
```

## [](#_parameters)Parameters

- `physicalDevice` The physical device to query the display handle on.
- `dpy` A connection to the X11 server from which `rrOutput` was queried.
- `rrOutput` An X11 RandR output ID.
- `pDisplay` The corresponding [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayKHR.html) handle will be returned here.

## [](#_description)Description

If there is no [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayKHR.html) corresponding to `rrOutput` on `physicalDevice`, [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) **must** be returned in `pDisplay`.

Valid Usage (Implicit)

- [](#VUID-vkGetRandROutputDisplayEXT-physicalDevice-parameter)VUID-vkGetRandROutputDisplayEXT-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetRandROutputDisplayEXT-dpy-parameter)VUID-vkGetRandROutputDisplayEXT-dpy-parameter  
  `dpy` **must** be a valid pointer to a `Display` value
- [](#VUID-vkGetRandROutputDisplayEXT-pDisplay-parameter)VUID-vkGetRandROutputDisplayEXT-pDisplay-parameter  
  `pDisplay` **must** be a valid pointer to a [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayKHR.html) handle

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_EXT\_acquire\_xlib\_display](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_acquire_xlib_display.html), [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayKHR.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetRandROutputDisplayEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0