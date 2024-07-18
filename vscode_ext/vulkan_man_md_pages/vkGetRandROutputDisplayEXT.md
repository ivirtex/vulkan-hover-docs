# vkGetRandROutputDisplayEXT(3) Manual Page

## Name

vkGetRandROutputDisplayEXT - Query the VkDisplayKHR corresponding to an
X11 RandR Output



## <a href="#_c_specification" class="anchor"></a>C Specification

When acquiring displays from an X11 server, an application may also wish
to enumerate and identify them using a native handle rather than a
`VkDisplayKHR` handle. To determine the `VkDisplayKHR` handle
corresponding to an X11 RandR Output, call:

``` c
// Provided by VK_EXT_acquire_xlib_display
VkResult vkGetRandROutputDisplayEXT(
    VkPhysicalDevice                            physicalDevice,
    Display*                                    dpy,
    RROutput                                    rrOutput,
    VkDisplayKHR*                               pDisplay);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` The physical device to query the display handle on.

- `dpy` A connection to the X11 server from which `rrOutput` was
  queried.

- `rrOutput` An X11 RandR output ID.

- `pDisplay` The corresponding [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayKHR.html) handle
  will be returned here.

## <a href="#_description" class="anchor"></a>Description

If there is no [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayKHR.html) corresponding to
`rrOutput` on `physicalDevice`, [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)
**must** be returned in `pDisplay`.

Valid Usage (Implicit)

- <a href="#VUID-vkGetRandROutputDisplayEXT-physicalDevice-parameter"
  id="VUID-vkGetRandROutputDisplayEXT-physicalDevice-parameter"></a>
  VUID-vkGetRandROutputDisplayEXT-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a href="#VUID-vkGetRandROutputDisplayEXT-dpy-parameter"
  id="VUID-vkGetRandROutputDisplayEXT-dpy-parameter"></a>
  VUID-vkGetRandROutputDisplayEXT-dpy-parameter  
  `dpy` **must** be a valid pointer to a `Display` value

- <a href="#VUID-vkGetRandROutputDisplayEXT-pDisplay-parameter"
  id="VUID-vkGetRandROutputDisplayEXT-pDisplay-parameter"></a>
  VUID-vkGetRandROutputDisplayEXT-pDisplay-parameter  
  `pDisplay` **must** be a valid pointer to a
  [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayKHR.html) handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_acquire_xlib_display](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_acquire_xlib_display.html),
[VkDisplayKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayKHR.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetRandROutputDisplayEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
