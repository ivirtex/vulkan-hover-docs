# vkReleaseFullScreenExclusiveModeEXT(3) Manual Page

## Name

vkReleaseFullScreenExclusiveModeEXT - Release full-screen exclusive mode
from a swapchain



## <a href="#_c_specification" class="anchor"></a>C Specification

To release exclusive full-screen access from a swapchain, call:

``` c
// Provided by VK_EXT_full_screen_exclusive
VkResult vkReleaseFullScreenExclusiveModeEXT(
    VkDevice                                    device,
    VkSwapchainKHR                              swapchain);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device associated with `swapchain`.

- `swapchain` is the swapchain to release exclusive full-screen access
  from.

## <a href="#_description" class="anchor"></a>Description

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>Applications will not be able to present to <code>swapchain</code>
after this call until exclusive full-screen access is reacquired. This
is usually useful to handle when an application is minimized or
otherwise intends to stop presenting for a time.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-vkReleaseFullScreenExclusiveModeEXT-swapchain-02677"
  id="VUID-vkReleaseFullScreenExclusiveModeEXT-swapchain-02677"></a>
  VUID-vkReleaseFullScreenExclusiveModeEXT-swapchain-02677  
  `swapchain` **must** not be in the retired state

- <a href="#VUID-vkReleaseFullScreenExclusiveModeEXT-swapchain-02678"
  id="VUID-vkReleaseFullScreenExclusiveModeEXT-swapchain-02678"></a>
  VUID-vkReleaseFullScreenExclusiveModeEXT-swapchain-02678  
  `swapchain` **must** be a swapchain created with a
  [VkSurfaceFullScreenExclusiveInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceFullScreenExclusiveInfoEXT.html)
  structure, with `fullScreenExclusive` set to
  `VK_FULL_SCREEN_EXCLUSIVE_APPLICATION_CONTROLLED_EXT`

Valid Usage (Implicit)

- <a href="#VUID-vkReleaseFullScreenExclusiveModeEXT-device-parameter"
  id="VUID-vkReleaseFullScreenExclusiveModeEXT-device-parameter"></a>
  VUID-vkReleaseFullScreenExclusiveModeEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkReleaseFullScreenExclusiveModeEXT-swapchain-parameter"
  id="VUID-vkReleaseFullScreenExclusiveModeEXT-swapchain-parameter"></a>
  VUID-vkReleaseFullScreenExclusiveModeEXT-swapchain-parameter  
  `swapchain` **must** be a valid [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html)
  handle

- <a href="#VUID-vkReleaseFullScreenExclusiveModeEXT-swapchain-parent"
  id="VUID-vkReleaseFullScreenExclusiveModeEXT-swapchain-parent"></a>
  VUID-vkReleaseFullScreenExclusiveModeEXT-swapchain-parent  
  `swapchain` **must** have been created, allocated, or retrieved from
  `device`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_SURFACE_LOST_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_full_screen_exclusive](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_full_screen_exclusive.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkReleaseFullScreenExclusiveModeEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
