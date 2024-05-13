# vkAcquireFullScreenExclusiveModeEXT(3) Manual Page

## Name

vkAcquireFullScreenExclusiveModeEXT - Acquire full-screen exclusive mode
for a swapchain



## <a href="#_c_specification" class="anchor"></a>C Specification

To acquire exclusive full-screen access for a swapchain, call:

``` c
// Provided by VK_EXT_full_screen_exclusive
VkResult vkAcquireFullScreenExclusiveModeEXT(
    VkDevice                                    device,
    VkSwapchainKHR                              swapchain);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device associated with `swapchain`.

- `swapchain` is the swapchain to acquire exclusive full-screen access
  for.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkAcquireFullScreenExclusiveModeEXT-swapchain-02674"
  id="VUID-vkAcquireFullScreenExclusiveModeEXT-swapchain-02674"></a>
  VUID-vkAcquireFullScreenExclusiveModeEXT-swapchain-02674  
  `swapchain` **must** not be in the retired state

- <a href="#VUID-vkAcquireFullScreenExclusiveModeEXT-swapchain-02675"
  id="VUID-vkAcquireFullScreenExclusiveModeEXT-swapchain-02675"></a>
  VUID-vkAcquireFullScreenExclusiveModeEXT-swapchain-02675  
  `swapchain` **must** be a swapchain created with a
  [VkSurfaceFullScreenExclusiveInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceFullScreenExclusiveInfoEXT.html)
  structure, with `fullScreenExclusive` set to
  `VK_FULL_SCREEN_EXCLUSIVE_APPLICATION_CONTROLLED_EXT`

- <a href="#VUID-vkAcquireFullScreenExclusiveModeEXT-swapchain-02676"
  id="VUID-vkAcquireFullScreenExclusiveModeEXT-swapchain-02676"></a>
  VUID-vkAcquireFullScreenExclusiveModeEXT-swapchain-02676  
  `swapchain` **must** not currently have exclusive full-screen access

A return value of `VK_SUCCESS` indicates that the `swapchain`
successfully acquired exclusive full-screen access. The swapchain will
retain this exclusivity until either the application releases exclusive
full-screen access with
[vkReleaseFullScreenExclusiveModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkReleaseFullScreenExclusiveModeEXT.html),
destroys the swapchain, or if any of the swapchain commands return
`VK_ERROR_FULL_SCREEN_EXCLUSIVE_MODE_LOST_EXT` indicating that the mode
was lost because of platform-specific changes.

If the swapchain was unable to acquire exclusive full-screen access to
the display then `VK_ERROR_INITIALIZATION_FAILED` is returned. An
application **can** attempt to acquire exclusive full-screen access
again for the same swapchain even if this command fails, or if
`VK_ERROR_FULL_SCREEN_EXCLUSIVE_MODE_LOST_EXT` has been returned by a
swapchain command.

Valid Usage (Implicit)

- <a href="#VUID-vkAcquireFullScreenExclusiveModeEXT-device-parameter"
  id="VUID-vkAcquireFullScreenExclusiveModeEXT-device-parameter"></a>
  VUID-vkAcquireFullScreenExclusiveModeEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkAcquireFullScreenExclusiveModeEXT-swapchain-parameter"
  id="VUID-vkAcquireFullScreenExclusiveModeEXT-swapchain-parameter"></a>
  VUID-vkAcquireFullScreenExclusiveModeEXT-swapchain-parameter  
  `swapchain` **must** be a valid [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html)
  handle

- <a href="#VUID-vkAcquireFullScreenExclusiveModeEXT-swapchain-parent"
  id="VUID-vkAcquireFullScreenExclusiveModeEXT-swapchain-parent"></a>
  VUID-vkAcquireFullScreenExclusiveModeEXT-swapchain-parent  
  `swapchain` **must** have been created, allocated, or retrieved from
  `device`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_INITIALIZATION_FAILED`

- `VK_ERROR_SURFACE_LOST_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_full_screen_exclusive](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_full_screen_exclusive.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkAcquireFullScreenExclusiveModeEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
