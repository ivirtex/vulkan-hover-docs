# vkAcquireXlibDisplayEXT(3) Manual Page

## Name

vkAcquireXlibDisplayEXT - Acquire access to a VkDisplayKHR using Xlib



## <a href="#_c_specification" class="anchor"></a>C Specification

To acquire permission to directly access a display in Vulkan from an X11
server, call:

``` c
// Provided by VK_EXT_acquire_xlib_display
VkResult vkAcquireXlibDisplayEXT(
    VkPhysicalDevice                            physicalDevice,
    Display*                                    dpy,
    VkDisplayKHR                                display);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` The physical device the display is on.

- `dpy` A connection to the X11 server that currently owns `display`.

- `display` The display the caller wishes to control in Vulkan.

## <a href="#_description" class="anchor"></a>Description

All permissions necessary to control the display are granted to the
Vulkan instance associated with `physicalDevice` until the display is
released or the X11 connection specified by `dpy` is terminated.
Permission to access the display **may** be temporarily revoked during
periods when the X11 server from which control was acquired itself loses
access to `display`. During such periods, operations which require
access to the display **must** fail with an appropriate error code. If
the X11 server associated with `dpy` does not own `display`, or if
permission to access it has already been acquired by another entity, the
call **must** return the error code `VK_ERROR_INITIALIZATION_FAILED`.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>One example of when an X11 server loses access to a display is when
it loses ownership of its virtual terminal.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a href="#VUID-vkAcquireXlibDisplayEXT-physicalDevice-parameter"
  id="VUID-vkAcquireXlibDisplayEXT-physicalDevice-parameter"></a>
  VUID-vkAcquireXlibDisplayEXT-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a href="#VUID-vkAcquireXlibDisplayEXT-dpy-parameter"
  id="VUID-vkAcquireXlibDisplayEXT-dpy-parameter"></a>
  VUID-vkAcquireXlibDisplayEXT-dpy-parameter  
  `dpy` **must** be a valid pointer to a `Display` value

- <a href="#VUID-vkAcquireXlibDisplayEXT-display-parameter"
  id="VUID-vkAcquireXlibDisplayEXT-display-parameter"></a>
  VUID-vkAcquireXlibDisplayEXT-display-parameter  
  `display` **must** be a valid [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayKHR.html) handle

- <a href="#VUID-vkAcquireXlibDisplayEXT-display-parent"
  id="VUID-vkAcquireXlibDisplayEXT-display-parent"></a>
  VUID-vkAcquireXlibDisplayEXT-display-parent  
  `display` **must** have been created, allocated, or retrieved from
  `physicalDevice`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_INITIALIZATION_FAILED`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_acquire_xlib_display](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_acquire_xlib_display.html),
[VkDisplayKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayKHR.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkAcquireXlibDisplayEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
