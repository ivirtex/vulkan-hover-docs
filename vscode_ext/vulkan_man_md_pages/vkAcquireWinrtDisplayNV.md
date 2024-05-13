# vkAcquireWinrtDisplayNV(3) Manual Page

## Name

vkAcquireWinrtDisplayNV - Acquire access to a VkDisplayKHR



## <a href="#_c_specification" class="anchor"></a>C Specification

To acquire permission to directly access a display in Vulkan on Windows
10, call:

``` c
// Provided by VK_NV_acquire_winrt_display
VkResult vkAcquireWinrtDisplayNV(
    VkPhysicalDevice                            physicalDevice,
    VkDisplayKHR                                display);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` The physical device the display is on.

- `display` The display the caller wishes to control in Vulkan.

## <a href="#_description" class="anchor"></a>Description

All permissions necessary to control the display are granted to the
Vulkan instance associated with `physicalDevice` until the display is
released or the application is terminated. Permission to access the
display **may** be revoked by events that cause Windows 10 itself to
lose access to `display`. If this has happened, operations which require
access to the display **must** fail with an appropriate error code. If
permission to access `display` has already been acquired by another
entity, the call **must** return the error code
`VK_ERROR_INITIALIZATION_FAILED`.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>The Vulkan instance acquires control of a <a
href="https://docs.microsoft.com/en-us/uwp/api/windows.devices.display.core.displaytarget">“winrt::Windows::Devices::Display::Core::DisplayTarget”</a>
by performing an operation equivalent to <a
href="https://docs.microsoft.com/en-us/uwp/api/windows.devices.display.core.displaymanager.tryacquiretarget">“winrt::Windows::Devices::Display::Core::DisplayManager.TryAcquireTarget()”</a>
on the “DisplayTarget”.</p></td>
</tr>
</tbody>
</table>

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>One example of when Windows 10 loses access to a display is when the
display is hot-unplugged.</p></td>
</tr>
</tbody>
</table>

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>One example of when a display has already been acquired by another
entity is when the Windows desktop compositor (DWM) is in control of the
display. Beginning with Windows 10 version 2004 it is possible to cause
DWM to release a display by using the “Advanced display settings”
sub-page of the “Display settings” control panel. <a
href="vkAcquireWinrtDisplayNV.html">vkAcquireWinrtDisplayNV</a> does not
itself cause DWM to release a display; this action must be performed
outside of Vulkan.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a href="#VUID-vkAcquireWinrtDisplayNV-physicalDevice-parameter"
  id="VUID-vkAcquireWinrtDisplayNV-physicalDevice-parameter"></a>
  VUID-vkAcquireWinrtDisplayNV-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a href="#VUID-vkAcquireWinrtDisplayNV-display-parameter"
  id="VUID-vkAcquireWinrtDisplayNV-display-parameter"></a>
  VUID-vkAcquireWinrtDisplayNV-display-parameter  
  `display` **must** be a valid [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayKHR.html) handle

- <a href="#VUID-vkAcquireWinrtDisplayNV-display-parent"
  id="VUID-vkAcquireWinrtDisplayNV-display-parent"></a>
  VUID-vkAcquireWinrtDisplayNV-display-parent  
  `display` **must** have been created, allocated, or retrieved from
  `physicalDevice`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_DEVICE_LOST`

- `VK_ERROR_INITIALIZATION_FAILED`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_acquire_winrt_display](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_acquire_winrt_display.html),
[VkDisplayKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayKHR.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkAcquireWinrtDisplayNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
